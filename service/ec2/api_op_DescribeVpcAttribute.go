// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcAttributeRequest
type DescribeVpcAttributeInput struct {
	_ struct{} `type:"structure"`

	// The VPC attribute.
	//
	// Attribute is a required field
	Attribute VpcAttributeName `type:"string" required:"true" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeVpcAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpcAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeVpcAttributeInput"}
	if len(s.Attribute) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Attribute"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcAttributeResult
type DescribeVpcAttributeOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the instances launched in the VPC get DNS hostnames. If
	// this attribute is true, instances in the VPC get DNS hostnames; otherwise,
	// they do not.
	EnableDnsHostnames *AttributeBooleanValue `json:"ec2:DescribeVpcAttributeOutput:EnableDnsHostnames" locationName:"enableDnsHostnames" type:"structure"`

	// Indicates whether DNS resolution is enabled for the VPC. If this attribute
	// is true, the Amazon DNS server resolves DNS hostnames for your instances
	// to their corresponding IP addresses; otherwise, it does not.
	EnableDnsSupport *AttributeBooleanValue `json:"ec2:DescribeVpcAttributeOutput:EnableDnsSupport" locationName:"enableDnsSupport" type:"structure"`

	// The ID of the VPC.
	VpcId *string `json:"ec2:DescribeVpcAttributeOutput:VpcId" locationName:"vpcId" type:"string"`
}

// String returns the string representation
func (s DescribeVpcAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVpcAttribute = "DescribeVpcAttribute"

// DescribeVpcAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified attribute of the specified VPC. You can specify only
// one attribute at a time.
//
//    // Example sending a request using DescribeVpcAttributeRequest.
//    req := client.DescribeVpcAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcAttribute
func (c *Client) DescribeVpcAttributeRequest(input *DescribeVpcAttributeInput) DescribeVpcAttributeRequest {
	op := &aws.Operation{
		Name:       opDescribeVpcAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpcAttributeInput{}
	}

	req := c.newRequest(op, input, &DescribeVpcAttributeOutput{})
	return DescribeVpcAttributeRequest{Request: req, Input: input, Copy: c.DescribeVpcAttributeRequest}
}

// DescribeVpcAttributeRequest is the request type for the
// DescribeVpcAttribute API operation.
type DescribeVpcAttributeRequest struct {
	*aws.Request
	Input *DescribeVpcAttributeInput
	Copy  func(*DescribeVpcAttributeInput) DescribeVpcAttributeRequest
}

// Send marshals and sends the DescribeVpcAttribute API request.
func (r DescribeVpcAttributeRequest) Send(ctx context.Context) (*DescribeVpcAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVpcAttributeResponse{
		DescribeVpcAttributeOutput: r.Request.Data.(*DescribeVpcAttributeOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVpcAttributeResponse is the response type for the
// DescribeVpcAttribute API operation.
type DescribeVpcAttributeResponse struct {
	*DescribeVpcAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVpcAttribute request.
func (r *DescribeVpcAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
