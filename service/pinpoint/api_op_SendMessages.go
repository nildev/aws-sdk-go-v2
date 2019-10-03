// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/SendMessagesRequest
type SendMessagesInput struct {
	_ struct{} `type:"structure" payload:"MessageRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the objects that define configuration and other settings for a
	// message.
	//
	// MessageRequest is a required field
	MessageRequest *MessageRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s SendMessagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendMessagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendMessagesInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.MessageRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("MessageRequest"))
	}
	if s.MessageRequest != nil {
		if err := s.MessageRequest.Validate(); err != nil {
			invalidParams.AddNested("MessageRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SendMessagesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MessageRequest != nil {
		v := s.MessageRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "MessageRequest", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/SendMessagesResponse
type SendMessagesOutput struct {
	_ struct{} `type:"structure" payload:"MessageResponse"`

	// Provides information about the results of a request to send a message to
	// an endpoint address.
	//
	// MessageResponse is a required field
	MessageResponse *MessageResponse `json:"pinpoint:SendMessagesOutput:MessageResponse" type:"structure" required:"true"`
}

// String returns the string representation
func (s SendMessagesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SendMessagesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MessageResponse != nil {
		v := s.MessageResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "MessageResponse", v, metadata)
	}
	return nil
}

const opSendMessages = "SendMessages"

// SendMessagesRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates and sends a direct message.
//
//    // Example sending a request using SendMessagesRequest.
//    req := client.SendMessagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/SendMessages
func (c *Client) SendMessagesRequest(input *SendMessagesInput) SendMessagesRequest {
	op := &aws.Operation{
		Name:       opSendMessages,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apps/{application-id}/messages",
	}

	if input == nil {
		input = &SendMessagesInput{}
	}

	req := c.newRequest(op, input, &SendMessagesOutput{})
	return SendMessagesRequest{Request: req, Input: input, Copy: c.SendMessagesRequest}
}

// SendMessagesRequest is the request type for the
// SendMessages API operation.
type SendMessagesRequest struct {
	*aws.Request
	Input *SendMessagesInput
	Copy  func(*SendMessagesInput) SendMessagesRequest
}

// Send marshals and sends the SendMessages API request.
func (r SendMessagesRequest) Send(ctx context.Context) (*SendMessagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendMessagesResponse{
		SendMessagesOutput: r.Request.Data.(*SendMessagesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendMessagesResponse is the response type for the
// SendMessages API operation.
type SendMessagesResponse struct {
	*SendMessagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendMessages request.
func (r *SendMessagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
