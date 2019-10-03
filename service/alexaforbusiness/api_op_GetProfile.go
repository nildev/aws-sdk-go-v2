// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetProfileRequest
type GetProfileInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the room profile for which to request details. Required.
	ProfileArn *string `type:"string"`
}

// String returns the string representation
func (s GetProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetProfileResponse
type GetProfileOutput struct {
	_ struct{} `type:"structure"`

	// The details of the room profile requested. Required.
	Profile *Profile `json:"a4b:GetProfileOutput:Profile" type:"structure"`
}

// String returns the string representation
func (s GetProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetProfile = "GetProfile"

// GetProfileRequest returns a request value for making API operation for
// Alexa For Business.
//
// Gets the details of a room profile by profile ARN.
//
//    // Example sending a request using GetProfileRequest.
//    req := client.GetProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetProfile
func (c *Client) GetProfileRequest(input *GetProfileInput) GetProfileRequest {
	op := &aws.Operation{
		Name:       opGetProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetProfileInput{}
	}

	req := c.newRequest(op, input, &GetProfileOutput{})
	return GetProfileRequest{Request: req, Input: input, Copy: c.GetProfileRequest}
}

// GetProfileRequest is the request type for the
// GetProfile API operation.
type GetProfileRequest struct {
	*aws.Request
	Input *GetProfileInput
	Copy  func(*GetProfileInput) GetProfileRequest
}

// Send marshals and sends the GetProfile API request.
func (r GetProfileRequest) Send(ctx context.Context) (*GetProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetProfileResponse{
		GetProfileOutput: r.Request.Data.(*GetProfileOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetProfileResponse is the response type for the
// GetProfile API operation.
type GetProfileResponse struct {
	*GetProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetProfile request.
func (r *GetProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
