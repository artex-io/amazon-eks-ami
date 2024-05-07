// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modify the instance metadata parameters on a running or stopped instance. When
// you modify the parameters on a stopped instance, they are applied when the
// instance is started. When you modify the parameters on a running instance, the
// API responds with a state of “pending”. After the parameter modifications are
// successfully applied to the instance, the state of the modifications changes
// from “pending” to “applied” in subsequent describe-instances API calls. For more
// information, see Instance metadata and user data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)
// in the Amazon EC2 User Guide.
func (c *Client) ModifyInstanceMetadataOptions(ctx context.Context, params *ModifyInstanceMetadataOptionsInput, optFns ...func(*Options)) (*ModifyInstanceMetadataOptionsOutput, error) {
	if params == nil {
		params = &ModifyInstanceMetadataOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceMetadataOptions", params, optFns, c.addOperationModifyInstanceMetadataOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceMetadataOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceMetadataOptionsInput struct {

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Enables or disables the HTTP metadata endpoint on your instances. If this
	// parameter is not specified, the existing state is maintained. If you specify a
	// value of disabled , you cannot access your instance metadata.
	HttpEndpoint types.InstanceMetadataEndpointState

	// Enables or disables the IPv6 endpoint for the instance metadata service.
	// Applies only if you enabled the HTTP metadata endpoint.
	HttpProtocolIpv6 types.InstanceMetadataProtocolState

	// The desired HTTP PUT response hop limit for instance metadata requests. The
	// larger the number, the further instance metadata requests can travel. If no
	// parameter is specified, the existing state is maintained. Possible values:
	// Integers from 1 to 64
	HttpPutResponseHopLimit *int32

	// Indicates whether IMDSv2 is required.
	//   - optional - IMDSv2 is optional. You can choose whether to send a session
	//   token in your instance metadata retrieval requests. If you retrieve IAM role
	//   credentials without a session token, you receive the IMDSv1 role credentials. If
	//   you retrieve IAM role credentials using a valid session token, you receive the
	//   IMDSv2 role credentials.
	//   - required - IMDSv2 is required. You must send a session token in your
	//   instance metadata retrieval requests. With this option, retrieving the IAM role
	//   credentials always returns IMDSv2 credentials; IMDSv1 credentials are not
	//   available.
	// Default: If the value of ImdsSupport for the Amazon Machine Image (AMI) for
	// your instance is v2.0 , the default is required .
	HttpTokens types.HttpTokensState

	// Set to enabled to allow access to instance tags from the instance metadata. Set
	// to disabled to turn off access to instance tags from the instance metadata. For
	// more information, see Work with instance tags using the instance metadata (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#work-with-tags-in-IMDS)
	// . Default: disabled
	InstanceMetadataTags types.InstanceMetadataTagsState

	noSmithyDocumentSerde
}

type ModifyInstanceMetadataOptionsOutput struct {

	// The ID of the instance.
	InstanceId *string

	// The metadata options for the instance.
	InstanceMetadataOptions *types.InstanceMetadataOptionsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyInstanceMetadataOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyInstanceMetadataOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyInstanceMetadataOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyInstanceMetadataOptions"); err != nil {
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
	if err = addOpModifyInstanceMetadataOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceMetadataOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyInstanceMetadataOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyInstanceMetadataOptions",
	}
}
