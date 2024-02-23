// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Remove a label or labels from a parameter.
func (c *Client) UnlabelParameterVersion(ctx context.Context, params *UnlabelParameterVersionInput, optFns ...func(*Options)) (*UnlabelParameterVersionOutput, error) {
	if params == nil {
		params = &UnlabelParameterVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnlabelParameterVersion", params, optFns, c.addOperationUnlabelParameterVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnlabelParameterVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnlabelParameterVersionInput struct {

	// One or more labels to delete from the specified parameter version.
	//
	// This member is required.
	Labels []string

	// The name of the parameter from which you want to delete one or more labels. You
	// can't enter the Amazon Resource Name (ARN) for a parameter, only the parameter
	// name itself.
	//
	// This member is required.
	Name *string

	// The specific version of the parameter which you want to delete one or more
	// labels from. If it isn't present, the call will fail.
	//
	// This member is required.
	ParameterVersion *int64

	noSmithyDocumentSerde
}

type UnlabelParameterVersionOutput struct {

	// The labels that aren't attached to the given parameter version.
	InvalidLabels []string

	// A list of all labels deleted from the parameter.
	RemovedLabels []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUnlabelParameterVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUnlabelParameterVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUnlabelParameterVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UnlabelParameterVersion"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpUnlabelParameterVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUnlabelParameterVersion(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opUnlabelParameterVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UnlabelParameterVersion",
	}
}
