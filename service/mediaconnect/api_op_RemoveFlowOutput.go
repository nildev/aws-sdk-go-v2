// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/RemoveFlowOutputRequest
type RemoveFlowOutputInput struct {
	_ struct{} `type:"structure"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`

	// OutputArn is a required field
	OutputArn *string `location:"uri" locationName:"outputArn" type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveFlowOutputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveFlowOutputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveFlowOutputInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if s.OutputArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveFlowOutputInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OutputArn != nil {
		v := *s.OutputArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "outputArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful RemoveFlowOutput request including the flow ARN
// and the output ARN that was removed.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/RemoveFlowOutputResponse
type RemoveFlowOutputOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the flow that is associated with the output you removed.
	FlowArn *string `json:"mediaconnect:RemoveFlowOutputOutput:FlowArn" locationName:"flowArn" type:"string"`

	// The ARN of the output that was removed.
	OutputArn *string `json:"mediaconnect:RemoveFlowOutputOutput:OutputArn" locationName:"outputArn" type:"string"`
}

// String returns the string representation
func (s RemoveFlowOutputOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveFlowOutputOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OutputArn != nil {
		v := *s.OutputArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "outputArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opRemoveFlowOutput = "RemoveFlowOutput"

// RemoveFlowOutputRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Removes an output from an existing flow. This request can be made only on
// an output that does not have an entitlement associated with it. If the output
// has an entitlement, you must revoke the entitlement instead. When an entitlement
// is revoked from a flow, the service automatically removes the associated
// output.
//
//    // Example sending a request using RemoveFlowOutputRequest.
//    req := client.RemoveFlowOutputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/RemoveFlowOutput
func (c *Client) RemoveFlowOutputRequest(input *RemoveFlowOutputInput) RemoveFlowOutputRequest {
	op := &aws.Operation{
		Name:       opRemoveFlowOutput,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/flows/{flowArn}/outputs/{outputArn}",
	}

	if input == nil {
		input = &RemoveFlowOutputInput{}
	}

	req := c.newRequest(op, input, &RemoveFlowOutputOutput{})
	return RemoveFlowOutputRequest{Request: req, Input: input, Copy: c.RemoveFlowOutputRequest}
}

// RemoveFlowOutputRequest is the request type for the
// RemoveFlowOutput API operation.
type RemoveFlowOutputRequest struct {
	*aws.Request
	Input *RemoveFlowOutputInput
	Copy  func(*RemoveFlowOutputInput) RemoveFlowOutputRequest
}

// Send marshals and sends the RemoveFlowOutput API request.
func (r RemoveFlowOutputRequest) Send(ctx context.Context) (*RemoveFlowOutputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveFlowOutputResponse{
		RemoveFlowOutputOutput: r.Request.Data.(*RemoveFlowOutputOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveFlowOutputResponse is the response type for the
// RemoveFlowOutput API operation.
type RemoveFlowOutputResponse struct {
	*RemoveFlowOutputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveFlowOutput request.
func (r *RemoveFlowOutputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
