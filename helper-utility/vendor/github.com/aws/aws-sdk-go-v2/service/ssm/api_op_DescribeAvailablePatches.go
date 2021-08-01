// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all patches eligible to be included in a patch baseline.
func (c *Client) DescribeAvailablePatches(ctx context.Context, params *DescribeAvailablePatchesInput, optFns ...func(*Options)) (*DescribeAvailablePatchesOutput, error) {
	if params == nil {
		params = &DescribeAvailablePatchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAvailablePatches", params, optFns, c.addOperationDescribeAvailablePatchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAvailablePatchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAvailablePatchesInput struct {

	// Each element in the array is a structure containing a key-value pair. Windows
	// Server Supported keys for Windows Server instance patches include the
	// following:
	//
	// * PATCH_SET Sample values: OS | APPLICATION
	//
	// * PRODUCT Sample
	// values: WindowsServer2012 | Office 2010 | MicrosoftDefenderAntivirus
	//
	// *
	// PRODUCT_FAMILY Sample values: Windows | Office
	//
	// * MSRC_SEVERITY Sample values:
	// ServicePacks | Important | Moderate
	//
	// * CLASSIFICATION Sample values:
	// ServicePacks | SecurityUpdates | DefinitionUpdates
	//
	// * PATCH_ID Sample values:
	// KB123456 | KB4516046
	//
	// Linux When specifying filters for Linux patches, you must
	// specify a key-pair for PRODUCT. For example, using the Command Line Interface
	// (CLI), the following command fails: aws ssm describe-available-patches --filters
	// Key=CVE_ID,Values=CVE-2018-3615 However, the following command succeeds: aws ssm
	// describe-available-patches --filters Key=PRODUCT,Values=AmazonLinux2018.03
	// Key=CVE_ID,Values=CVE-2018-3615 Supported keys for Linux instance patches
	// include the following:
	//
	// * PRODUCT Sample values: AmazonLinux2018.03 |
	// AmazonLinux2.0
	//
	// * NAME Sample values: kernel-headers | samba-python | php
	//
	// *
	// SEVERITY Sample values: Critical | Important | Medium | Low
	//
	// * EPOCH Sample
	// values: 0 | 1
	//
	// * VERSION Sample values: 78.6.1 | 4.10.16
	//
	// * RELEASE Sample
	// values: 9.56.amzn1 | 1.amzn2
	//
	// * ARCH Sample values: i686 | x86_64
	//
	// * REPOSITORY
	// Sample values: Core | Updates
	//
	// * ADVISORY_ID Sample values: ALAS-2018-1058 |
	// ALAS2-2021-1594
	//
	// * CVE_ID Sample values: CVE-2018-3615 | CVE-2020-1472
	//
	// *
	// BUGZILLA_ID Sample values: 1463241
	Filters []types.PatchOrchestratorFilter

	// The maximum number of patches to return (per page).
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeAvailablePatchesOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// An array of patches. Each entry in the array is a patch structure.
	Patches []types.Patch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeAvailablePatchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAvailablePatches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAvailablePatches{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAvailablePatches(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeAvailablePatchesAPIClient is a client that implements the
// DescribeAvailablePatches operation.
type DescribeAvailablePatchesAPIClient interface {
	DescribeAvailablePatches(context.Context, *DescribeAvailablePatchesInput, ...func(*Options)) (*DescribeAvailablePatchesOutput, error)
}

var _ DescribeAvailablePatchesAPIClient = (*Client)(nil)

// DescribeAvailablePatchesPaginatorOptions is the paginator options for
// DescribeAvailablePatches
type DescribeAvailablePatchesPaginatorOptions struct {
	// The maximum number of patches to return (per page).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAvailablePatchesPaginator is a paginator for DescribeAvailablePatches
type DescribeAvailablePatchesPaginator struct {
	options   DescribeAvailablePatchesPaginatorOptions
	client    DescribeAvailablePatchesAPIClient
	params    *DescribeAvailablePatchesInput
	nextToken *string
	firstPage bool
}

// NewDescribeAvailablePatchesPaginator returns a new
// DescribeAvailablePatchesPaginator
func NewDescribeAvailablePatchesPaginator(client DescribeAvailablePatchesAPIClient, params *DescribeAvailablePatchesInput, optFns ...func(*DescribeAvailablePatchesPaginatorOptions)) *DescribeAvailablePatchesPaginator {
	if params == nil {
		params = &DescribeAvailablePatchesInput{}
	}

	options := DescribeAvailablePatchesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAvailablePatchesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAvailablePatchesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeAvailablePatches page.
func (p *DescribeAvailablePatchesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAvailablePatchesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeAvailablePatches(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeAvailablePatches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeAvailablePatches",
	}
}
