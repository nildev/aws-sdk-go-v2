// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickdevicesservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/FinalizeDeviceClaimRequest
type FinalizeDeviceClaimInput struct {
	_ struct{} `type:"structure"`

	// DeviceId is a required field
	DeviceId *string `location:"uri" locationName:"deviceId" type:"string" required:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s FinalizeDeviceClaimInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FinalizeDeviceClaimInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FinalizeDeviceClaimInput"}

	if s.DeviceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s FinalizeDeviceClaimInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.DeviceId != nil {
		v := *s.DeviceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "deviceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/FinalizeDeviceClaimResponse
type FinalizeDeviceClaimOutput struct {
	_ struct{} `type:"structure"`

	State *string `json:"devices.iot1click:FinalizeDeviceClaimOutput:State" locationName:"state" type:"string"`
}

// String returns the string representation
func (s FinalizeDeviceClaimOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s FinalizeDeviceClaimOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.State != nil {
		v := *s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "state", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opFinalizeDeviceClaim = "FinalizeDeviceClaim"

// FinalizeDeviceClaimRequest returns a request value for making API operation for
// AWS IoT 1-Click Devices Service.
//
// Given a device ID, finalizes the claim request for the associated device.
//
// Claiming a device consists of initiating a claim, then publishing a device
// event, and finalizing the claim. For a device of type button, a device event
// can be published by simply clicking the device.
//
//    // Example sending a request using FinalizeDeviceClaimRequest.
//    req := client.FinalizeDeviceClaimRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/FinalizeDeviceClaim
func (c *Client) FinalizeDeviceClaimRequest(input *FinalizeDeviceClaimInput) FinalizeDeviceClaimRequest {
	op := &aws.Operation{
		Name:       opFinalizeDeviceClaim,
		HTTPMethod: "PUT",
		HTTPPath:   "/devices/{deviceId}/finalize-claim",
	}

	if input == nil {
		input = &FinalizeDeviceClaimInput{}
	}

	req := c.newRequest(op, input, &FinalizeDeviceClaimOutput{})
	return FinalizeDeviceClaimRequest{Request: req, Input: input, Copy: c.FinalizeDeviceClaimRequest}
}

// FinalizeDeviceClaimRequest is the request type for the
// FinalizeDeviceClaim API operation.
type FinalizeDeviceClaimRequest struct {
	*aws.Request
	Input *FinalizeDeviceClaimInput
	Copy  func(*FinalizeDeviceClaimInput) FinalizeDeviceClaimRequest
}

// Send marshals and sends the FinalizeDeviceClaim API request.
func (r FinalizeDeviceClaimRequest) Send(ctx context.Context) (*FinalizeDeviceClaimResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &FinalizeDeviceClaimResponse{
		FinalizeDeviceClaimOutput: r.Request.Data.(*FinalizeDeviceClaimOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// FinalizeDeviceClaimResponse is the response type for the
// FinalizeDeviceClaim API operation.
type FinalizeDeviceClaimResponse struct {
	*FinalizeDeviceClaimOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// FinalizeDeviceClaim request.
func (r *FinalizeDeviceClaimResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
