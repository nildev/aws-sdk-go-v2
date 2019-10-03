// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/DescribeDetectorModelRequest
type DescribeDetectorModelInput struct {
	_ struct{} `type:"structure"`

	// The name of the detector model.
	//
	// DetectorModelName is a required field
	DetectorModelName *string `location:"uri" locationName:"detectorModelName" min:"1" type:"string" required:"true"`

	// The version of the detector model.
	DetectorModelVersion *string `location:"querystring" locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDetectorModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDetectorModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDetectorModelInput"}

	if s.DetectorModelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorModelName"))
	}
	if s.DetectorModelName != nil && len(*s.DetectorModelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorModelName", 1))
	}
	if s.DetectorModelVersion != nil && len(*s.DetectorModelVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorModelVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDetectorModelInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DetectorModelName != nil {
		v := *s.DetectorModelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorModelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DetectorModelVersion != nil {
		v := *s.DetectorModelVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/DescribeDetectorModelResponse
type DescribeDetectorModelOutput struct {
	_ struct{} `type:"structure"`

	// Information about the detector model.
	DetectorModel *DetectorModel `json:"iotevents:DescribeDetectorModelOutput:DetectorModel" locationName:"detectorModel" type:"structure"`
}

// String returns the string representation
func (s DescribeDetectorModelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDetectorModelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DetectorModel != nil {
		v := s.DetectorModel

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "detectorModel", v, metadata)
	}
	return nil
}

const opDescribeDetectorModel = "DescribeDetectorModel"

// DescribeDetectorModelRequest returns a request value for making API operation for
// AWS IoT Events.
//
// Describes a detector model. If the "version" parameter is not specified,
// information about the latest version is returned.
//
//    // Example sending a request using DescribeDetectorModelRequest.
//    req := client.DescribeDetectorModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/DescribeDetectorModel
func (c *Client) DescribeDetectorModelRequest(input *DescribeDetectorModelInput) DescribeDetectorModelRequest {
	op := &aws.Operation{
		Name:       opDescribeDetectorModel,
		HTTPMethod: "GET",
		HTTPPath:   "/detector-models/{detectorModelName}",
	}

	if input == nil {
		input = &DescribeDetectorModelInput{}
	}

	req := c.newRequest(op, input, &DescribeDetectorModelOutput{})
	return DescribeDetectorModelRequest{Request: req, Input: input, Copy: c.DescribeDetectorModelRequest}
}

// DescribeDetectorModelRequest is the request type for the
// DescribeDetectorModel API operation.
type DescribeDetectorModelRequest struct {
	*aws.Request
	Input *DescribeDetectorModelInput
	Copy  func(*DescribeDetectorModelInput) DescribeDetectorModelRequest
}

// Send marshals and sends the DescribeDetectorModel API request.
func (r DescribeDetectorModelRequest) Send(ctx context.Context) (*DescribeDetectorModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDetectorModelResponse{
		DescribeDetectorModelOutput: r.Request.Data.(*DescribeDetectorModelOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDetectorModelResponse is the response type for the
// DescribeDetectorModel API operation.
type DescribeDetectorModelResponse struct {
	*DescribeDetectorModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDetectorModel request.
func (r *DescribeDetectorModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
