// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeMyUserProfileInput
type DescribeMyUserProfileInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeMyUserProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a DescribeMyUserProfile request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeMyUserProfileResult
type DescribeMyUserProfileOutput struct {
	_ struct{} `type:"structure"`

	// A UserProfile object that describes the user's SSH information.
	UserProfile *SelfUserProfile `json:"opsworks:DescribeMyUserProfileOutput:UserProfile" type:"structure"`
}

// String returns the string representation
func (s DescribeMyUserProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMyUserProfile = "DescribeMyUserProfile"

// DescribeMyUserProfileRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Describes a user's SSH information.
//
// Required Permissions: To use this action, an IAM user must have self-management
// enabled or an attached policy that explicitly grants permissions. For more
// information about user permissions, see Managing User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeMyUserProfileRequest.
//    req := client.DescribeMyUserProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeMyUserProfile
func (c *Client) DescribeMyUserProfileRequest(input *DescribeMyUserProfileInput) DescribeMyUserProfileRequest {
	op := &aws.Operation{
		Name:       opDescribeMyUserProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMyUserProfileInput{}
	}

	req := c.newRequest(op, input, &DescribeMyUserProfileOutput{})
	return DescribeMyUserProfileRequest{Request: req, Input: input, Copy: c.DescribeMyUserProfileRequest}
}

// DescribeMyUserProfileRequest is the request type for the
// DescribeMyUserProfile API operation.
type DescribeMyUserProfileRequest struct {
	*aws.Request
	Input *DescribeMyUserProfileInput
	Copy  func(*DescribeMyUserProfileInput) DescribeMyUserProfileRequest
}

// Send marshals and sends the DescribeMyUserProfile API request.
func (r DescribeMyUserProfileRequest) Send(ctx context.Context) (*DescribeMyUserProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMyUserProfileResponse{
		DescribeMyUserProfileOutput: r.Request.Data.(*DescribeMyUserProfileOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMyUserProfileResponse is the response type for the
// DescribeMyUserProfile API operation.
type DescribeMyUserProfileResponse struct {
	*DescribeMyUserProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMyUserProfile request.
func (r *DescribeMyUserProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
