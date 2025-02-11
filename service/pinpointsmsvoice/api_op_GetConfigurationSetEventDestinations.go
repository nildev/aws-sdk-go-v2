// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/GetConfigurationSetEventDestinationsRequest
type GetConfigurationSetEventDestinationsInput struct {
	_ struct{} `type:"structure"`

	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetConfigurationSetEventDestinationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConfigurationSetEventDestinationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConfigurationSetEventDestinationsInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetConfigurationSetEventDestinationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object that contains information about an event destination.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/GetConfigurationSetEventDestinationsResponse
type GetConfigurationSetEventDestinationsOutput struct {
	_ struct{} `type:"structure"`

	// An array of EventDestination objects. Each EventDestination object includes
	// ARNs and other information that define an event destination.
	EventDestinations []EventDestination `type:"list"`
}

// String returns the string representation
func (s GetConfigurationSetEventDestinationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetConfigurationSetEventDestinationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EventDestinations != nil {
		v := s.EventDestinations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "EventDestinations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetConfigurationSetEventDestinations = "GetConfigurationSetEventDestinations"

// GetConfigurationSetEventDestinationsRequest returns a request value for making API operation for
// Amazon Pinpoint SMS and Voice Service.
//
// Obtain information about an event destination, including the types of events
// it reports, the Amazon Resource Name (ARN) of the destination, and the name
// of the event destination.
//
//    // Example sending a request using GetConfigurationSetEventDestinationsRequest.
//    req := client.GetConfigurationSetEventDestinationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/GetConfigurationSetEventDestinations
func (c *Client) GetConfigurationSetEventDestinationsRequest(input *GetConfigurationSetEventDestinationsInput) GetConfigurationSetEventDestinationsRequest {
	op := &aws.Operation{
		Name:       opGetConfigurationSetEventDestinations,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/sms-voice/configuration-sets/{ConfigurationSetName}/event-destinations",
	}

	if input == nil {
		input = &GetConfigurationSetEventDestinationsInput{}
	}

	req := c.newRequest(op, input, &GetConfigurationSetEventDestinationsOutput{})
	return GetConfigurationSetEventDestinationsRequest{Request: req, Input: input, Copy: c.GetConfigurationSetEventDestinationsRequest}
}

// GetConfigurationSetEventDestinationsRequest is the request type for the
// GetConfigurationSetEventDestinations API operation.
type GetConfigurationSetEventDestinationsRequest struct {
	*aws.Request
	Input *GetConfigurationSetEventDestinationsInput
	Copy  func(*GetConfigurationSetEventDestinationsInput) GetConfigurationSetEventDestinationsRequest
}

// Send marshals and sends the GetConfigurationSetEventDestinations API request.
func (r GetConfigurationSetEventDestinationsRequest) Send(ctx context.Context) (*GetConfigurationSetEventDestinationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConfigurationSetEventDestinationsResponse{
		GetConfigurationSetEventDestinationsOutput: r.Request.Data.(*GetConfigurationSetEventDestinationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConfigurationSetEventDestinationsResponse is the response type for the
// GetConfigurationSetEventDestinations API operation.
type GetConfigurationSetEventDestinationsResponse struct {
	*GetConfigurationSetEventDestinationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConfigurationSetEventDestinations request.
func (r *GetConfigurationSetEventDestinationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
