package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v3"

	"github.com/TylerBrock/colorjson"
	"github.com/araddon/dateparse"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecrtypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/fatih/color"
)

type ScanConfig struct {
	Severity []string `yaml:"severity"`
	Excluded []struct {
		ID          string `yaml:"id"`
		Expires     string `yaml:"expires,omitempty"`
		Description string `yaml:"description,omitempty"`
	} `yaml:"excluded"`
}

func init() {
	if os.Getenv("GITLAB_CI") != "" {
		color.NoColor = false
	}
	if v := os.Getenv("ECR_SCANNER_DEBUG"); len(v) != 0 {
		log.SetLevel(log.DebugLevel)
	}
	if v := os.Getenv("ECR_SCANNER_TRACE"); len(v) != 0 {
		log.SetLevel(log.TraceLevel)
		log.SetReportCaller(true)
	}
}

func main() {
	if v := os.Getenv("ECR_SCANNER_BYPASS"); len(v) != 0 {
		log.Println("ECR_SCANNER_BYPASS detected... Skipping")
		os.Exit(0)
	}

	if len(os.Args) <= 1 {
		log.Fatal("Must supply image to scan as an argument!")
	}

	ecrUrl, err := url.Parse("https://" + os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	hostPart := strings.Split(ecrUrl.Host, ".")
	if len(hostPart) != 6 {
		log.Fatalln("Unknown host portion of ECR URL:", ecrUrl.Host)
	}
	ecrAccount := hostPart[0]
	ecrRegion := hostPart[3]

	imagePart := strings.Split(ecrUrl.Path, ":")
	if len(imagePart) != 2 {
		log.Fatal("Unable to detect image:tag in supplied URL")
	}
	image := strings.TrimPrefix(imagePart[0], "/")
	tag := imagePart[1]

	// AWS Session
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(ecrRegion))
	if err != nil {
		log.Fatal(err)
	}

	if log.IsLevelEnabled(log.DebugLevel) {
		stsclient := sts.NewFromConfig(cfg)
		id, err := stsclient.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("CallerIdentity: %#v", id)
	}

	// ECR Client
	ecrclient := ecr.NewFromConfig(cfg)

	input := &ecr.DescribeImageScanFindingsInput{
		RegistryId:     aws.String(ecrAccount),
		RepositoryName: aws.String(image),
		ImageId: &ecrtypes.ImageIdentifier{
			ImageTag: aws.String(tag),
		},
	}

	findings := make([]ecrtypes.ImageScanFinding, 0)
	findingsPaginator := ecr.NewDescribeImageScanFindingsPaginator(ecrclient, input)
	for findingsPaginator.HasMorePages() {
		res, err := findingsPaginator.NextPage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		if res.ImageScanFindings != nil {
			findings = append(findings, res.ImageScanFindings.Findings...)
		}
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	yamlFile, err := os.ReadFile(path.Join(pwd, ".ecr-scanner.yml"))
	if err != nil {
		log.Fatal(err)
	}

	var scanConfig ScanConfig
	err = yaml.Unmarshal(yamlFile, &scanConfig)
	if err != nil {
		log.Fatal(err)
	}

	matching := make([]ecrtypes.ImageScanFinding, 0)
	for _, finding := range findings {
		func() {
			for _, excluded := range scanConfig.Excluded {
				if *finding.Name == excluded.ID {
					if excluded.Expires != "" {
						t, err := dateparse.ParseLocal(excluded.Expires)
						if err != nil {
							log.Fatal(err)
						}
						if time.Now().After(t) {
							log.Warnln("Exception has expired!", excluded)
							break
						}
					}
					log.Println("Skipping", excluded.ID)
					return
				}
			}
			for _, severity := range scanConfig.Severity {
				if string(finding.Severity) == strings.ToUpper(severity) {
					matching = append(matching, finding)
					return
				}
			}
		}()
	}

	if len(matching) > 0 {
		jsonBytes, err := json.Marshal(matching)
		if err != nil {
			log.Fatal(err)
		}
		var findings interface{}
		err = json.Unmarshal(jsonBytes, &findings)
		if err != nil {
			log.Fatal(err)
		}
		f := colorjson.NewFormatter()
		f.Indent = 2
		out, err := f.Marshal(findings)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
		log.Fatal("Findings detected!")
	}

	log.Infoln("No findings detected!")
}
