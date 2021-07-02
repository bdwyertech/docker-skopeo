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

// Describes details about the activation, such as the date and time the activation
// was created, its expiration date, the IAM role assigned to the instances in the
// activation, and the number of instances registered by using this activation.
func (c *Client) DescribeActivations(ctx context.Context, params *DescribeActivationsInput, optFns ...func(*Options)) (*DescribeActivationsOutput, error) {
	if params == nil {
		params = &DescribeActivationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeActivations", params, optFns, c.addOperationDescribeActivationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeActivationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeActivationsInput struct {

	// A filter to view information about your activations.
	Filters []types.DescribeActivationsFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
}

type DescribeActivationsOutput struct {

	// A list of activations for your AWS account.
	ActivationList []types.Activation

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeActivationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeActivations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeActivations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActivations(options.Region), middleware.Before); err != nil {
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

// DescribeActivationsAPIClient is a client that implements the DescribeActivations
// operation.
type DescribeActivationsAPIClient interface {
	DescribeActivations(context.Context, *DescribeActivationsInput, ...func(*Options)) (*DescribeActivationsOutput, error)
}

var _ DescribeActivationsAPIClient = (*Client)(nil)

// DescribeActivationsPaginatorOptions is the paginator options for
// DescribeActivations
type DescribeActivationsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeActivationsPaginator is a paginator for DescribeActivations
type DescribeActivationsPaginator struct {
	options   DescribeActivationsPaginatorOptions
	client    DescribeActivationsAPIClient
	params    *DescribeActivationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeActivationsPaginator returns a new DescribeActivationsPaginator
func NewDescribeActivationsPaginator(client DescribeActivationsAPIClient, params *DescribeActivationsInput, optFns ...func(*DescribeActivationsPaginatorOptions)) *DescribeActivationsPaginator {
	if params == nil {
		params = &DescribeActivationsInput{}
	}

	options := DescribeActivationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeActivationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeActivationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeActivations page.
func (p *DescribeActivationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeActivationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeActivations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeActivations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeActivations",
	}
}
