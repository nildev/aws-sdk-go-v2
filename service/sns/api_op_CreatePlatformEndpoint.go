// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for CreatePlatformEndpoint action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/CreatePlatformEndpointInput
type CreatePlatformEndpointInput struct {
	_ struct{} `type:"structure"`

	// For a list of attributes, see SetEndpointAttributes (https://docs.aws.amazon.com/sns/latest/api/API_SetEndpointAttributes.html).
	Attributes map[string]string `type:"map"`

	// Arbitrary user data to associate with the endpoint. Amazon SNS does not use
	// this data. The data must be in UTF-8 format and less than 2KB.
	CustomUserData *string `type:"string"`

	// PlatformApplicationArn returned from CreatePlatformApplication is used to
	// create a an endpoint.
	//
	// PlatformApplicationArn is a required field
	PlatformApplicationArn *string `type:"string" required:"true"`

	// Unique identifier created by the notification service for an app on a device.
	// The specific name for Token will vary, depending on which notification service
	// is being used. For example, when using APNS as the notification service,
	// you need the device token. Alternatively, when using GCM or ADM, the device
	// token equivalent is called the registration ID.
	//
	// Token is a required field
	Token *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePlatformEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePlatformEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePlatformEndpointInput"}

	if s.PlatformApplicationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlatformApplicationArn"))
	}

	if s.Token == nil {
		invalidParams.Add(aws.NewErrParamRequired("Token"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response from CreateEndpoint action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/CreateEndpointResponse
type CreatePlatformEndpointOutput struct {
	_ struct{} `type:"structure"`

	// EndpointArn returned from CreateEndpoint action.
	EndpointArn *string `json:"sns:CreatePlatformEndpointOutput:EndpointArn" type:"string"`
}

// String returns the string representation
func (s CreatePlatformEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePlatformEndpoint = "CreatePlatformEndpoint"

// CreatePlatformEndpointRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Creates an endpoint for a device and mobile app on one of the supported push
// notification services, such as GCM and APNS. CreatePlatformEndpoint requires
// the PlatformApplicationArn that is returned from CreatePlatformApplication.
// The EndpointArn that is returned when using CreatePlatformEndpoint can then
// be used by the Publish action to send a message to a mobile app or by the
// Subscribe action for subscription to a topic. The CreatePlatformEndpoint
// action is idempotent, so if the requester already owns an endpoint with the
// same device token and attributes, that endpoint's ARN is returned without
// creating a new endpoint. For more information, see Using Amazon SNS Mobile
// Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
//
// When using CreatePlatformEndpoint with Baidu, two attributes must be provided:
// ChannelId and UserId. The token field must also contain the ChannelId. For
// more information, see Creating an Amazon SNS Endpoint for Baidu (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePushBaiduEndpoint.html).
//
//    // Example sending a request using CreatePlatformEndpointRequest.
//    req := client.CreatePlatformEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/CreatePlatformEndpoint
func (c *Client) CreatePlatformEndpointRequest(input *CreatePlatformEndpointInput) CreatePlatformEndpointRequest {
	op := &aws.Operation{
		Name:       opCreatePlatformEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePlatformEndpointInput{}
	}

	req := c.newRequest(op, input, &CreatePlatformEndpointOutput{})
	return CreatePlatformEndpointRequest{Request: req, Input: input, Copy: c.CreatePlatformEndpointRequest}
}

// CreatePlatformEndpointRequest is the request type for the
// CreatePlatformEndpoint API operation.
type CreatePlatformEndpointRequest struct {
	*aws.Request
	Input *CreatePlatformEndpointInput
	Copy  func(*CreatePlatformEndpointInput) CreatePlatformEndpointRequest
}

// Send marshals and sends the CreatePlatformEndpoint API request.
func (r CreatePlatformEndpointRequest) Send(ctx context.Context) (*CreatePlatformEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePlatformEndpointResponse{
		CreatePlatformEndpointOutput: r.Request.Data.(*CreatePlatformEndpointOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePlatformEndpointResponse is the response type for the
// CreatePlatformEndpoint API operation.
type CreatePlatformEndpointResponse struct {
	*CreatePlatformEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePlatformEndpoint request.
func (r *CreatePlatformEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
