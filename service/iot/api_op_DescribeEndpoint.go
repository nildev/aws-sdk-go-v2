// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the DescribeEndpoint operation.
type DescribeEndpointInput struct {
	_ struct{} `type:"structure"`

	// The endpoint type. Valid endpoint types include:
	//
	//    * iot:Data - Returns a VeriSign signed data endpoint.
	//
	//    * iot:Data-ATS - Returns an ATS signed data endpoint.
	//
	//    * iot:CredentialProvider - Returns an AWS IoT credentials provider API
	//    endpoint.
	//
	//    * iot:Jobs - Returns an AWS IoT device management Jobs API endpoint.
	EndpointType *string `location:"querystring" locationName:"endpointType" type:"string"`
}

// String returns the string representation
func (s DescribeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeEndpointInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.EndpointType != nil {
		v := *s.EndpointType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "endpointType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the DescribeEndpoint operation.
type DescribeEndpointOutput struct {
	_ struct{} `type:"structure"`

	// The endpoint. The format of the endpoint is as follows: identifier.iot.region.amazonaws.com.
	EndpointAddress *string `json:"iot:DescribeEndpointOutput:EndpointAddress" locationName:"endpointAddress" type:"string"`
}

// String returns the string representation
func (s DescribeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeEndpointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EndpointAddress != nil {
		v := *s.EndpointAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endpointAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeEndpoint = "DescribeEndpoint"

// DescribeEndpointRequest returns a request value for making API operation for
// AWS IoT.
//
// Returns a unique endpoint specific to the AWS account making the call.
//
//    // Example sending a request using DescribeEndpointRequest.
//    req := client.DescribeEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeEndpointRequest(input *DescribeEndpointInput) DescribeEndpointRequest {
	op := &aws.Operation{
		Name:       opDescribeEndpoint,
		HTTPMethod: "GET",
		HTTPPath:   "/endpoint",
	}

	if input == nil {
		input = &DescribeEndpointInput{}
	}

	req := c.newRequest(op, input, &DescribeEndpointOutput{})
	return DescribeEndpointRequest{Request: req, Input: input, Copy: c.DescribeEndpointRequest}
}

// DescribeEndpointRequest is the request type for the
// DescribeEndpoint API operation.
type DescribeEndpointRequest struct {
	*aws.Request
	Input *DescribeEndpointInput
	Copy  func(*DescribeEndpointInput) DescribeEndpointRequest
}

// Send marshals and sends the DescribeEndpoint API request.
func (r DescribeEndpointRequest) Send(ctx context.Context) (*DescribeEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEndpointResponse{
		DescribeEndpointOutput: r.Request.Data.(*DescribeEndpointOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEndpointResponse is the response type for the
// DescribeEndpoint API operation.
type DescribeEndpointResponse struct {
	*DescribeEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEndpoint request.
func (r *DescribeEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
