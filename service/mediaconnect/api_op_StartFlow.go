// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/StartFlowRequest
type StartFlowInput struct {
	_ struct{} `type:"structure"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`
}

// String returns the string representation
func (s StartFlowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartFlowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartFlowInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartFlowInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful StartFlow request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/StartFlowResponse
type StartFlowOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the flow that you started.
	FlowArn *string `json:"mediaconnect:StartFlowOutput:FlowArn" locationName:"flowArn" type:"string"`

	// The status of the flow when the StartFlow process begins.
	Status Status `json:"mediaconnect:StartFlowOutput:Status" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s StartFlowOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartFlowOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opStartFlow = "StartFlow"

// StartFlowRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Starts a flow.
//
//    // Example sending a request using StartFlowRequest.
//    req := client.StartFlowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/StartFlow
func (c *Client) StartFlowRequest(input *StartFlowInput) StartFlowRequest {
	op := &aws.Operation{
		Name:       opStartFlow,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/flows/start/{flowArn}",
	}

	if input == nil {
		input = &StartFlowInput{}
	}

	req := c.newRequest(op, input, &StartFlowOutput{})
	return StartFlowRequest{Request: req, Input: input, Copy: c.StartFlowRequest}
}

// StartFlowRequest is the request type for the
// StartFlow API operation.
type StartFlowRequest struct {
	*aws.Request
	Input *StartFlowInput
	Copy  func(*StartFlowInput) StartFlowRequest
}

// Send marshals and sends the StartFlow API request.
func (r StartFlowRequest) Send(ctx context.Context) (*StartFlowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartFlowResponse{
		StartFlowOutput: r.Request.Data.(*StartFlowOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartFlowResponse is the response type for the
// StartFlow API operation.
type StartFlowResponse struct {
	*StartFlowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartFlow request.
func (r *StartFlowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
