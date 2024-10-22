// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the tasks in a maintenance window.
//
// For maintenance window tasks without a specified target, you can't supply
// values for --max-errors and --max-concurrency . Instead, the system inserts a
// placeholder value of 1 , which may be reported in the response to this command.
// These values don't affect the running of your task and can be ignored.
func (c *Client) DescribeMaintenanceWindowTasks(ctx context.Context, params *DescribeMaintenanceWindowTasksInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowTasksOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindowTasks", params, optFns, c.addOperationDescribeMaintenanceWindowTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowTasksInput struct {

	// The ID of the maintenance window whose tasks should be retrieved.
	//
	// This member is required.
	WindowId *string

	// Optional filters used to narrow down the scope of the returned tasks. The
	// supported filter keys are WindowTaskId , TaskArn , Priority , and TaskType .
	Filters []types.MaintenanceWindowFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeMaintenanceWindowTasksOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the tasks in the maintenance window.
	Tasks []types.MaintenanceWindowTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMaintenanceWindowTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMaintenanceWindowTasks"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeMaintenanceWindowTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowTasks(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// DescribeMaintenanceWindowTasksPaginatorOptions is the paginator options for
// DescribeMaintenanceWindowTasks
type DescribeMaintenanceWindowTasksPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMaintenanceWindowTasksPaginator is a paginator for
// DescribeMaintenanceWindowTasks
type DescribeMaintenanceWindowTasksPaginator struct {
	options   DescribeMaintenanceWindowTasksPaginatorOptions
	client    DescribeMaintenanceWindowTasksAPIClient
	params    *DescribeMaintenanceWindowTasksInput
	nextToken *string
	firstPage bool
}

// NewDescribeMaintenanceWindowTasksPaginator returns a new
// DescribeMaintenanceWindowTasksPaginator
func NewDescribeMaintenanceWindowTasksPaginator(client DescribeMaintenanceWindowTasksAPIClient, params *DescribeMaintenanceWindowTasksInput, optFns ...func(*DescribeMaintenanceWindowTasksPaginatorOptions)) *DescribeMaintenanceWindowTasksPaginator {
	if params == nil {
		params = &DescribeMaintenanceWindowTasksInput{}
	}

	options := DescribeMaintenanceWindowTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMaintenanceWindowTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMaintenanceWindowTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMaintenanceWindowTasks page.
func (p *DescribeMaintenanceWindowTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMaintenanceWindowTasksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeMaintenanceWindowTasks(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeMaintenanceWindowTasksAPIClient is a client that implements the
// DescribeMaintenanceWindowTasks operation.
type DescribeMaintenanceWindowTasksAPIClient interface {
	DescribeMaintenanceWindowTasks(context.Context, *DescribeMaintenanceWindowTasksInput, ...func(*Options)) (*DescribeMaintenanceWindowTasksOutput, error)
}

var _ DescribeMaintenanceWindowTasksAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMaintenanceWindowTasks",
	}
}
