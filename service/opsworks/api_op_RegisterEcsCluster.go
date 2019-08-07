// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterEcsClusterRequest
type RegisterEcsClusterInput struct {
	_ struct{} `type:"structure"`

	// The cluster's ARN.
	//
	// EcsClusterArn is a required field
	EcsClusterArn *string `type:"string" required:"true"`

	// The stack ID.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterEcsClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterEcsClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterEcsClusterInput"}

	if s.EcsClusterArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EcsClusterArn"))
	}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a RegisterEcsCluster request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterEcsClusterResult
type RegisterEcsClusterOutput struct {
	_ struct{} `type:"structure"`

	// The cluster's ARN.
	EcsClusterArn *string `json:"opsworks:RegisterEcsClusterOutput:EcsClusterArn" type:"string"`
}

// String returns the string representation
func (s RegisterEcsClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterEcsCluster = "RegisterEcsCluster"

// RegisterEcsClusterRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Registers a specified Amazon ECS cluster with a stack. You can register only
// one cluster with a stack. A cluster can be registered with only one stack.
// For more information, see Resource Management (https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-ecscluster.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using RegisterEcsClusterRequest.
//    req := client.RegisterEcsClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterEcsCluster
func (c *Client) RegisterEcsClusterRequest(input *RegisterEcsClusterInput) RegisterEcsClusterRequest {
	op := &aws.Operation{
		Name:       opRegisterEcsCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterEcsClusterInput{}
	}

	req := c.newRequest(op, input, &RegisterEcsClusterOutput{})
	return RegisterEcsClusterRequest{Request: req, Input: input, Copy: c.RegisterEcsClusterRequest}
}

// RegisterEcsClusterRequest is the request type for the
// RegisterEcsCluster API operation.
type RegisterEcsClusterRequest struct {
	*aws.Request
	Input *RegisterEcsClusterInput
	Copy  func(*RegisterEcsClusterInput) RegisterEcsClusterRequest
}

// Send marshals and sends the RegisterEcsCluster API request.
func (r RegisterEcsClusterRequest) Send(ctx context.Context) (*RegisterEcsClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterEcsClusterResponse{
		RegisterEcsClusterOutput: r.Request.Data.(*RegisterEcsClusterOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterEcsClusterResponse is the response type for the
// RegisterEcsCluster API operation.
type RegisterEcsClusterResponse struct {
	*RegisterEcsClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterEcsCluster request.
func (r *RegisterEcsClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
