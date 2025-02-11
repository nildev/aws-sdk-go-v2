// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeUserProfilesRequest
type DescribeUserProfilesInput struct {
	_ struct{} `type:"structure"`

	// An array of IAM or federated user ARNs that identify the users to be described.
	IamUserArns []string `type:"list"`
}

// String returns the string representation
func (s DescribeUserProfilesInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a DescribeUserProfiles request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeUserProfilesResult
type DescribeUserProfilesOutput struct {
	_ struct{} `type:"structure"`

	// A Users object that describes the specified users.
	UserProfiles []UserProfile `type:"list"`
}

// String returns the string representation
func (s DescribeUserProfilesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeUserProfiles = "DescribeUserProfiles"

// DescribeUserProfilesRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Describe specified users.
//
// Required Permissions: To use this action, an IAM user must have an attached
// policy that explicitly grants permissions. For more information about user
// permissions, see Managing User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeUserProfilesRequest.
//    req := client.DescribeUserProfilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeUserProfiles
func (c *Client) DescribeUserProfilesRequest(input *DescribeUserProfilesInput) DescribeUserProfilesRequest {
	op := &aws.Operation{
		Name:       opDescribeUserProfiles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeUserProfilesInput{}
	}

	req := c.newRequest(op, input, &DescribeUserProfilesOutput{})
	return DescribeUserProfilesRequest{Request: req, Input: input, Copy: c.DescribeUserProfilesRequest}
}

// DescribeUserProfilesRequest is the request type for the
// DescribeUserProfiles API operation.
type DescribeUserProfilesRequest struct {
	*aws.Request
	Input *DescribeUserProfilesInput
	Copy  func(*DescribeUserProfilesInput) DescribeUserProfilesRequest
}

// Send marshals and sends the DescribeUserProfiles API request.
func (r DescribeUserProfilesRequest) Send(ctx context.Context) (*DescribeUserProfilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserProfilesResponse{
		DescribeUserProfilesOutput: r.Request.Data.(*DescribeUserProfilesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserProfilesResponse is the response type for the
// DescribeUserProfiles API operation.
type DescribeUserProfilesResponse struct {
	*DescribeUserProfilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUserProfiles request.
func (r *DescribeUserProfilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
