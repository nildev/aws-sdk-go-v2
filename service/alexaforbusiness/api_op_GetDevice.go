// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetDeviceRequest
type GetDeviceInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the device for which to request details. Required.
	DeviceArn *string `type:"string"`
}

// String returns the string representation
func (s GetDeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetDeviceResponse
type GetDeviceOutput struct {
	_ struct{} `type:"structure"`

	// The details of the device requested. Required.
	Device *Device `json:"a4b:GetDeviceOutput:Device" type:"structure"`
}

// String returns the string representation
func (s GetDeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDevice = "GetDevice"

// GetDeviceRequest returns a request value for making API operation for
// Alexa For Business.
//
// Gets the details of a device by device ARN.
//
//    // Example sending a request using GetDeviceRequest.
//    req := client.GetDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetDevice
func (c *Client) GetDeviceRequest(input *GetDeviceInput) GetDeviceRequest {
	op := &aws.Operation{
		Name:       opGetDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDeviceInput{}
	}

	req := c.newRequest(op, input, &GetDeviceOutput{})
	return GetDeviceRequest{Request: req, Input: input, Copy: c.GetDeviceRequest}
}

// GetDeviceRequest is the request type for the
// GetDevice API operation.
type GetDeviceRequest struct {
	*aws.Request
	Input *GetDeviceInput
	Copy  func(*GetDeviceInput) GetDeviceRequest
}

// Send marshals and sends the GetDevice API request.
func (r GetDeviceRequest) Send(ctx context.Context) (*GetDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeviceResponse{
		GetDeviceOutput: r.Request.Data.(*GetDeviceOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeviceResponse is the response type for the
// GetDevice API operation.
type GetDeviceResponse struct {
	*GetDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDevice request.
func (r *GetDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
