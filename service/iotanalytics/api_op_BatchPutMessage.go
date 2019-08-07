// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/BatchPutMessageRequest
type BatchPutMessageInput struct {
	_ struct{} `type:"structure"`

	// The name of the channel where the messages are sent.
	//
	// ChannelName is a required field
	ChannelName *string `locationName:"channelName" min:"1" type:"string" required:"true"`

	// The list of messages to be sent. Each message has format: '{ "messageId":
	// "string", "payload": "string"}'.
	//
	// Note that the field names of message payloads (data) that you send to AWS
	// IoT Analytics:
	//
	//    * Must contain only alphanumeric characters and undescores (_); no other
	//    special characters are allowed.
	//
	//    * Must begin with an alphabetic character or single underscore (_).
	//
	//    * Cannot contain hyphens (-).
	//
	//    * In regular expression terms: "^[A-Za-z_]([A-Za-z0-9]*|[A-Za-z0-9][A-Za-z0-9_]*)$".
	//
	//    * Cannot be greater than 255 characters.
	//
	//    * Are case-insensitive. (Fields named "foo" and "FOO" in the same payload
	//    are considered duplicates.)
	//
	// For example, {"temp_01": 29} or {"_temp_01": 29} are valid, but {"temp-01":
	// 29}, {"01_temp": 29} or {"__temp_01": 29} are invalid in message payloads.
	//
	// Messages is a required field
	Messages []Message `locationName:"messages" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchPutMessageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchPutMessageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchPutMessageInput"}

	if s.ChannelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelName"))
	}
	if s.ChannelName != nil && len(*s.ChannelName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelName", 1))
	}

	if s.Messages == nil {
		invalidParams.Add(aws.NewErrParamRequired("Messages"))
	}
	if s.Messages != nil {
		for i, v := range s.Messages {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Messages", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchPutMessageInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ChannelName != nil {
		v := *s.ChannelName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "channelName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Messages != nil {
		v := s.Messages

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "messages", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/BatchPutMessageResponse
type BatchPutMessageOutput struct {
	_ struct{} `type:"structure"`

	// A list of any errors encountered when sending the messages to the channel.
	BatchPutMessageErrorEntries []BatchPutMessageErrorEntry `json:"iotanalytics:BatchPutMessageOutput:BatchPutMessageErrorEntries" locationName:"batchPutMessageErrorEntries" type:"list"`
}

// String returns the string representation
func (s BatchPutMessageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchPutMessageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BatchPutMessageErrorEntries != nil {
		v := s.BatchPutMessageErrorEntries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "batchPutMessageErrorEntries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchPutMessage = "BatchPutMessage"

// BatchPutMessageRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Sends messages to a channel.
//
//    // Example sending a request using BatchPutMessageRequest.
//    req := client.BatchPutMessageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/BatchPutMessage
func (c *Client) BatchPutMessageRequest(input *BatchPutMessageInput) BatchPutMessageRequest {
	op := &aws.Operation{
		Name:       opBatchPutMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/messages/batch",
	}

	if input == nil {
		input = &BatchPutMessageInput{}
	}

	req := c.newRequest(op, input, &BatchPutMessageOutput{})
	return BatchPutMessageRequest{Request: req, Input: input, Copy: c.BatchPutMessageRequest}
}

// BatchPutMessageRequest is the request type for the
// BatchPutMessage API operation.
type BatchPutMessageRequest struct {
	*aws.Request
	Input *BatchPutMessageInput
	Copy  func(*BatchPutMessageInput) BatchPutMessageRequest
}

// Send marshals and sends the BatchPutMessage API request.
func (r BatchPutMessageRequest) Send(ctx context.Context) (*BatchPutMessageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchPutMessageResponse{
		BatchPutMessageOutput: r.Request.Data.(*BatchPutMessageOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchPutMessageResponse is the response type for the
// BatchPutMessage API operation.
type BatchPutMessageResponse struct {
	*BatchPutMessageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchPutMessage request.
func (r *BatchPutMessageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
