// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/UpdateChannelRequest
type UpdateChannelInput struct {
	_ struct{} `type:"structure"`

	// ChannelId is a required field
	ChannelId *string `location:"uri" locationName:"channelId" type:"string" required:"true"`

	Destinations []OutputDestination `locationName:"destinations" type:"list"`

	// Encoder Settings
	EncoderSettings *EncoderSettings `locationName:"encoderSettings" type:"structure"`

	InputAttachments []InputAttachment `locationName:"inputAttachments" type:"list"`

	InputSpecification *InputSpecification `locationName:"inputSpecification" type:"structure"`

	// The log level the user wants for their channel.
	LogLevel LogLevel `locationName:"logLevel" type:"string" enum:"true"`

	Name *string `locationName:"name" type:"string"`

	RoleArn *string `locationName:"roleArn" type:"string"`
}

// String returns the string representation
func (s UpdateChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateChannelInput"}

	if s.ChannelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelId"))
	}
	if s.Destinations != nil {
		for i, v := range s.Destinations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Destinations", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.EncoderSettings != nil {
		if err := s.EncoderSettings.Validate(); err != nil {
			invalidParams.AddNested("EncoderSettings", err.(aws.ErrInvalidParams))
		}
	}
	if s.InputAttachments != nil {
		for i, v := range s.InputAttachments {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "InputAttachments", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Destinations != nil {
		v := s.Destinations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "destinations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.EncoderSettings != nil {
		v := s.EncoderSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encoderSettings", v, metadata)
	}
	if s.InputAttachments != nil {
		v := s.InputAttachments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "inputAttachments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.InputSpecification != nil {
		v := s.InputSpecification

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "inputSpecification", v, metadata)
	}
	if len(s.LogLevel) > 0 {
		v := s.LogLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "logLevel", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ChannelId != nil {
		v := *s.ChannelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "channelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/UpdateChannelResponse
type UpdateChannelOutput struct {
	_ struct{} `type:"structure"`

	Channel *Channel `json:"medialive:UpdateChannelOutput:Channel" locationName:"channel" type:"structure"`
}

// String returns the string representation
func (s UpdateChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Channel != nil {
		v := s.Channel

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "channel", v, metadata)
	}
	return nil
}

const opUpdateChannel = "UpdateChannel"

// UpdateChannelRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Updates a channel.
//
//    // Example sending a request using UpdateChannelRequest.
//    req := client.UpdateChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/UpdateChannel
func (c *Client) UpdateChannelRequest(input *UpdateChannelInput) UpdateChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/prod/channels/{channelId}",
	}

	if input == nil {
		input = &UpdateChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateChannelOutput{})
	return UpdateChannelRequest{Request: req, Input: input, Copy: c.UpdateChannelRequest}
}

// UpdateChannelRequest is the request type for the
// UpdateChannel API operation.
type UpdateChannelRequest struct {
	*aws.Request
	Input *UpdateChannelInput
	Copy  func(*UpdateChannelInput) UpdateChannelRequest
}

// Send marshals and sends the UpdateChannel API request.
func (r UpdateChannelRequest) Send(ctx context.Context) (*UpdateChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateChannelResponse{
		UpdateChannelOutput: r.Request.Data.(*UpdateChannelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateChannelResponse is the response type for the
// UpdateChannel API operation.
type UpdateChannelResponse struct {
	*UpdateChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateChannel request.
func (r *UpdateChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
