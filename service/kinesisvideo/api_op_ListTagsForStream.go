// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/ListTagsForStreamInput
type ListTagsForStreamInput struct {
	_ struct{} `type:"structure"`

	// If you specify this parameter and the result of a ListTagsForStream call
	// is truncated, the response includes a token that you can use in the next
	// request to fetch the next batch of tags.
	NextToken *string `type:"string"`

	// The Amazon Resource Name (ARN) of the stream that you want to list tags for.
	StreamARN *string `min:"1" type:"string"`

	// The name of the stream that you want to list tags for.
	StreamName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListTagsForStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForStreamInput"}
	if s.StreamARN != nil && len(*s.StreamARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamARN", 1))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsForStreamInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamARN != nil {
		v := *s.StreamARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamName != nil {
		v := *s.StreamName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/ListTagsForStreamOutput
type ListTagsForStreamOutput struct {
	_ struct{} `type:"structure"`

	// If you specify this parameter and the result of a ListTags call is truncated,
	// the response includes a token that you can use in the next request to fetch
	// the next set of tags.
	NextToken *string `type:"string"`

	// A map of tag keys and values associated with the specified stream.
	Tags map[string]string `min:"1" type:"map"`
}

// String returns the string representation
func (s ListTagsForStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsForStreamOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opListTagsForStream = "ListTagsForStream"

// ListTagsForStreamRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Returns a list of tags associated with the specified stream.
//
// In the request, you must specify either the StreamName or the StreamARN.
//
//    // Example sending a request using ListTagsForStreamRequest.
//    req := client.ListTagsForStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/ListTagsForStream
func (c *Client) ListTagsForStreamRequest(input *ListTagsForStreamInput) ListTagsForStreamRequest {
	op := &aws.Operation{
		Name:       opListTagsForStream,
		HTTPMethod: "POST",
		HTTPPath:   "/listTagsForStream",
	}

	if input == nil {
		input = &ListTagsForStreamInput{}
	}

	req := c.newRequest(op, input, &ListTagsForStreamOutput{})
	return ListTagsForStreamRequest{Request: req, Input: input, Copy: c.ListTagsForStreamRequest}
}

// ListTagsForStreamRequest is the request type for the
// ListTagsForStream API operation.
type ListTagsForStreamRequest struct {
	*aws.Request
	Input *ListTagsForStreamInput
	Copy  func(*ListTagsForStreamInput) ListTagsForStreamRequest
}

// Send marshals and sends the ListTagsForStream API request.
func (r ListTagsForStreamRequest) Send(ctx context.Context) (*ListTagsForStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForStreamResponse{
		ListTagsForStreamOutput: r.Request.Data.(*ListTagsForStreamOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForStreamResponse is the response type for the
// ListTagsForStream API operation.
type ListTagsForStreamResponse struct {
	*ListTagsForStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForStream request.
func (r *ListTagsForStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
