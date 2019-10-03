// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcClassicLinkRequest
type DescribeVpcClassicLinkInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * is-classic-link-enabled - Whether the VPC is enabled for ClassicLink
	//    (true | false).
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more VPCs for which you want to describe the ClassicLink status.
	VpcIds []string `locationName:"VpcId" locationNameList:"VpcId" type:"list"`
}

// String returns the string representation
func (s DescribeVpcClassicLinkInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcClassicLinkResult
type DescribeVpcClassicLinkOutput struct {
	_ struct{} `type:"structure"`

	// The ClassicLink status of one or more VPCs.
	Vpcs []VpcClassicLink `json:"ec2:DescribeVpcClassicLinkOutput:Vpcs" locationName:"vpcSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpcClassicLinkOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVpcClassicLink = "DescribeVpcClassicLink"

// DescribeVpcClassicLinkRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the ClassicLink status of one or more VPCs.
//
//    // Example sending a request using DescribeVpcClassicLinkRequest.
//    req := client.DescribeVpcClassicLinkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcClassicLink
func (c *Client) DescribeVpcClassicLinkRequest(input *DescribeVpcClassicLinkInput) DescribeVpcClassicLinkRequest {
	op := &aws.Operation{
		Name:       opDescribeVpcClassicLink,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpcClassicLinkInput{}
	}

	req := c.newRequest(op, input, &DescribeVpcClassicLinkOutput{})
	return DescribeVpcClassicLinkRequest{Request: req, Input: input, Copy: c.DescribeVpcClassicLinkRequest}
}

// DescribeVpcClassicLinkRequest is the request type for the
// DescribeVpcClassicLink API operation.
type DescribeVpcClassicLinkRequest struct {
	*aws.Request
	Input *DescribeVpcClassicLinkInput
	Copy  func(*DescribeVpcClassicLinkInput) DescribeVpcClassicLinkRequest
}

// Send marshals and sends the DescribeVpcClassicLink API request.
func (r DescribeVpcClassicLinkRequest) Send(ctx context.Context) (*DescribeVpcClassicLinkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVpcClassicLinkResponse{
		DescribeVpcClassicLinkOutput: r.Request.Data.(*DescribeVpcClassicLinkOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVpcClassicLinkResponse is the response type for the
// DescribeVpcClassicLink API operation.
type DescribeVpcClassicLinkResponse struct {
	*DescribeVpcClassicLinkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVpcClassicLink request.
func (r *DescribeVpcClassicLinkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
