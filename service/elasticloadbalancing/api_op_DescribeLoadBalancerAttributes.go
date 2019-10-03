// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeLoadBalancerAttributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DescribeLoadBalancerAttributesInput
type DescribeLoadBalancerAttributesInput struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeLoadBalancerAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLoadBalancerAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLoadBalancerAttributesInput"}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of DescribeLoadBalancerAttributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DescribeLoadBalancerAttributesOutput
type DescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the load balancer attributes.
	LoadBalancerAttributes *LoadBalancerAttributes `json:"elasticloadbalancing:DescribeLoadBalancerAttributesOutput:LoadBalancerAttributes" type:"structure"`
}

// String returns the string representation
func (s DescribeLoadBalancerAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLoadBalancerAttributes = "DescribeLoadBalancerAttributes"

// DescribeLoadBalancerAttributesRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Describes the attributes for the specified load balancer.
//
//    // Example sending a request using DescribeLoadBalancerAttributesRequest.
//    req := client.DescribeLoadBalancerAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DescribeLoadBalancerAttributes
func (c *Client) DescribeLoadBalancerAttributesRequest(input *DescribeLoadBalancerAttributesInput) DescribeLoadBalancerAttributesRequest {
	op := &aws.Operation{
		Name:       opDescribeLoadBalancerAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLoadBalancerAttributesInput{}
	}

	req := c.newRequest(op, input, &DescribeLoadBalancerAttributesOutput{})
	return DescribeLoadBalancerAttributesRequest{Request: req, Input: input, Copy: c.DescribeLoadBalancerAttributesRequest}
}

// DescribeLoadBalancerAttributesRequest is the request type for the
// DescribeLoadBalancerAttributes API operation.
type DescribeLoadBalancerAttributesRequest struct {
	*aws.Request
	Input *DescribeLoadBalancerAttributesInput
	Copy  func(*DescribeLoadBalancerAttributesInput) DescribeLoadBalancerAttributesRequest
}

// Send marshals and sends the DescribeLoadBalancerAttributes API request.
func (r DescribeLoadBalancerAttributesRequest) Send(ctx context.Context) (*DescribeLoadBalancerAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLoadBalancerAttributesResponse{
		DescribeLoadBalancerAttributesOutput: r.Request.Data.(*DescribeLoadBalancerAttributesOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLoadBalancerAttributesResponse is the response type for the
// DescribeLoadBalancerAttributes API operation.
type DescribeLoadBalancerAttributesResponse struct {
	*DescribeLoadBalancerAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLoadBalancerAttributes request.
func (r *DescribeLoadBalancerAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
