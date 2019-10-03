// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/SetIpAddressTypeInput
type SetIpAddressTypeInput struct {
	_ struct{} `type:"structure"`

	// The IP address type. The possible values are ipv4 (for IPv4 addresses) and
	// dualstack (for IPv4 and IPv6 addresses). Internal load balancers must use
	// ipv4. Network Load Balancers must use ipv4.
	//
	// IpAddressType is a required field
	IpAddressType IpAddressType `type:"string" required:"true" enum:"true"`

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// LoadBalancerArn is a required field
	LoadBalancerArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SetIpAddressTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetIpAddressTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetIpAddressTypeInput"}
	if len(s.IpAddressType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("IpAddressType"))
	}

	if s.LoadBalancerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/SetIpAddressTypeOutput
type SetIpAddressTypeOutput struct {
	_ struct{} `type:"structure"`

	// The IP address type.
	IpAddressType IpAddressType `json:"elasticloadbalancing:SetIpAddressTypeOutput:IpAddressType" type:"string" enum:"true"`
}

// String returns the string representation
func (s SetIpAddressTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetIpAddressType = "SetIpAddressType"

// SetIpAddressTypeRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Sets the type of IP addresses used by the subnets of the specified Application
// Load Balancer or Network Load Balancer.
//
//    // Example sending a request using SetIpAddressTypeRequest.
//    req := client.SetIpAddressTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/SetIpAddressType
func (c *Client) SetIpAddressTypeRequest(input *SetIpAddressTypeInput) SetIpAddressTypeRequest {
	op := &aws.Operation{
		Name:       opSetIpAddressType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetIpAddressTypeInput{}
	}

	req := c.newRequest(op, input, &SetIpAddressTypeOutput{})
	return SetIpAddressTypeRequest{Request: req, Input: input, Copy: c.SetIpAddressTypeRequest}
}

// SetIpAddressTypeRequest is the request type for the
// SetIpAddressType API operation.
type SetIpAddressTypeRequest struct {
	*aws.Request
	Input *SetIpAddressTypeInput
	Copy  func(*SetIpAddressTypeInput) SetIpAddressTypeRequest
}

// Send marshals and sends the SetIpAddressType API request.
func (r SetIpAddressTypeRequest) Send(ctx context.Context) (*SetIpAddressTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetIpAddressTypeResponse{
		SetIpAddressTypeOutput: r.Request.Data.(*SetIpAddressTypeOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetIpAddressTypeResponse is the response type for the
// SetIpAddressType API operation.
type SetIpAddressTypeResponse struct {
	*SetIpAddressTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetIpAddressType request.
func (r *SetIpAddressTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
