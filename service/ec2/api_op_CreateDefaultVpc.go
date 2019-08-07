// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateDefaultVpcRequest
type CreateDefaultVpcInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s CreateDefaultVpcInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateDefaultVpcResult
type CreateDefaultVpcOutput struct {
	_ struct{} `type:"structure"`

	// Information about the VPC.
	Vpc *Vpc `json:"ec2:CreateDefaultVpcOutput:Vpc" locationName:"vpc" type:"structure"`
}

// String returns the string representation
func (s CreateDefaultVpcOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDefaultVpc = "CreateDefaultVpc"

// CreateDefaultVpcRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a default VPC with a size /16 IPv4 CIDR block and a default subnet
// in each Availability Zone. For more information about the components of a
// default VPC, see Default VPC and Default Subnets (https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html)
// in the Amazon Virtual Private Cloud User Guide. You cannot specify the components
// of the default VPC yourself.
//
// If you deleted your previous default VPC, you can create a default VPC. You
// cannot have more than one default VPC per Region.
//
// If your account supports EC2-Classic, you cannot use this action to create
// a default VPC in a Region that supports EC2-Classic. If you want a default
// VPC in a Region that supports EC2-Classic, see "I really want a default VPC
// for my existing EC2 account. Is that possible?" in the Default VPCs FAQ (http://aws.amazon.com/vpc/faqs/#Default_VPCs).
//
//    // Example sending a request using CreateDefaultVpcRequest.
//    req := client.CreateDefaultVpcRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateDefaultVpc
func (c *Client) CreateDefaultVpcRequest(input *CreateDefaultVpcInput) CreateDefaultVpcRequest {
	op := &aws.Operation{
		Name:       opCreateDefaultVpc,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDefaultVpcInput{}
	}

	req := c.newRequest(op, input, &CreateDefaultVpcOutput{})
	return CreateDefaultVpcRequest{Request: req, Input: input, Copy: c.CreateDefaultVpcRequest}
}

// CreateDefaultVpcRequest is the request type for the
// CreateDefaultVpc API operation.
type CreateDefaultVpcRequest struct {
	*aws.Request
	Input *CreateDefaultVpcInput
	Copy  func(*CreateDefaultVpcInput) CreateDefaultVpcRequest
}

// Send marshals and sends the CreateDefaultVpc API request.
func (r CreateDefaultVpcRequest) Send(ctx context.Context) (*CreateDefaultVpcResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDefaultVpcResponse{
		CreateDefaultVpcOutput: r.Request.Data.(*CreateDefaultVpcOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDefaultVpcResponse is the response type for the
// CreateDefaultVpc API operation.
type CreateDefaultVpcResponse struct {
	*CreateDefaultVpcOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDefaultVpc request.
func (r *CreateDefaultVpcResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
