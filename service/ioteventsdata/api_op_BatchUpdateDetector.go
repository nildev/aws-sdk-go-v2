// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ioteventsdata

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-data-2018-10-23/BatchUpdateDetectorRequest
type BatchUpdateDetectorInput struct {
	_ struct{} `type:"structure"`

	// The list of detectors (instances) to update, along with the values to update.
	//
	// Detectors is a required field
	Detectors []UpdateDetectorRequest `locationName:"detectors" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchUpdateDetectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchUpdateDetectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchUpdateDetectorInput"}

	if s.Detectors == nil {
		invalidParams.Add(aws.NewErrParamRequired("Detectors"))
	}
	if s.Detectors != nil && len(s.Detectors) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Detectors", 1))
	}
	if s.Detectors != nil {
		for i, v := range s.Detectors {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Detectors", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchUpdateDetectorInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Detectors != nil {
		v := s.Detectors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "detectors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-data-2018-10-23/BatchUpdateDetectorResponse
type BatchUpdateDetectorOutput struct {
	_ struct{} `type:"structure"`

	// A list of those detector updates that resulted in errors. (If an error is
	// listed here, the specific update did not occur.)
	BatchUpdateDetectorErrorEntries []BatchUpdateDetectorErrorEntry `json:"data.iotevents:BatchUpdateDetectorOutput:BatchUpdateDetectorErrorEntries" locationName:"batchUpdateDetectorErrorEntries" type:"list"`
}

// String returns the string representation
func (s BatchUpdateDetectorOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchUpdateDetectorOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BatchUpdateDetectorErrorEntries != nil {
		v := s.BatchUpdateDetectorErrorEntries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "batchUpdateDetectorErrorEntries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchUpdateDetector = "BatchUpdateDetector"

// BatchUpdateDetectorRequest returns a request value for making API operation for
// AWS IoT Events Data.
//
// Updates the state, variable values, and timer settings of one or more detectors
// (instances) of a specified detector model.
//
//    // Example sending a request using BatchUpdateDetectorRequest.
//    req := client.BatchUpdateDetectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-data-2018-10-23/BatchUpdateDetector
func (c *Client) BatchUpdateDetectorRequest(input *BatchUpdateDetectorInput) BatchUpdateDetectorRequest {
	op := &aws.Operation{
		Name:       opBatchUpdateDetector,
		HTTPMethod: "POST",
		HTTPPath:   "/detectors",
	}

	if input == nil {
		input = &BatchUpdateDetectorInput{}
	}

	req := c.newRequest(op, input, &BatchUpdateDetectorOutput{})
	return BatchUpdateDetectorRequest{Request: req, Input: input, Copy: c.BatchUpdateDetectorRequest}
}

// BatchUpdateDetectorRequest is the request type for the
// BatchUpdateDetector API operation.
type BatchUpdateDetectorRequest struct {
	*aws.Request
	Input *BatchUpdateDetectorInput
	Copy  func(*BatchUpdateDetectorInput) BatchUpdateDetectorRequest
}

// Send marshals and sends the BatchUpdateDetector API request.
func (r BatchUpdateDetectorRequest) Send(ctx context.Context) (*BatchUpdateDetectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchUpdateDetectorResponse{
		BatchUpdateDetectorOutput: r.Request.Data.(*BatchUpdateDetectorOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchUpdateDetectorResponse is the response type for the
// BatchUpdateDetector API operation.
type BatchUpdateDetectorResponse struct {
	*BatchUpdateDetectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchUpdateDetector request.
func (r *BatchUpdateDetectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
