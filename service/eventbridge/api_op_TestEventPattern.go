// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/TestEventPatternRequest
type TestEventPatternInput struct {
	_ struct{} `type:"structure"`

	// The event, in JSON format, to test against the event pattern.
	//
	// Event is a required field
	Event *string `type:"string" required:"true"`

	// The event pattern. For more information, see Event Patterns (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)
	// in the Amazon EventBridge User Guide.
	//
	// EventPattern is a required field
	EventPattern *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TestEventPatternInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestEventPatternInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TestEventPatternInput"}

	if s.Event == nil {
		invalidParams.Add(aws.NewErrParamRequired("Event"))
	}

	if s.EventPattern == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventPattern"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/TestEventPatternResponse
type TestEventPatternOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the event matches the event pattern.
	Result *bool `json:"events:TestEventPatternOutput:Result" type:"boolean"`
}

// String returns the string representation
func (s TestEventPatternOutput) String() string {
	return awsutil.Prettify(s)
}

const opTestEventPattern = "TestEventPattern"

// TestEventPatternRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Tests whether the specified event pattern matches the provided event.
//
// Most services in AWS treat : or / as the same character in Amazon Resource
// Names (ARNs). However, EventBridge uses an exact match in event patterns
// and rules. Be sure to use the correct ARN characters when creating event
// patterns so that they match the ARN syntax in the event that you want to
// match.
//
//    // Example sending a request using TestEventPatternRequest.
//    req := client.TestEventPatternRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/TestEventPattern
func (c *Client) TestEventPatternRequest(input *TestEventPatternInput) TestEventPatternRequest {
	op := &aws.Operation{
		Name:       opTestEventPattern,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TestEventPatternInput{}
	}

	req := c.newRequest(op, input, &TestEventPatternOutput{})
	return TestEventPatternRequest{Request: req, Input: input, Copy: c.TestEventPatternRequest}
}

// TestEventPatternRequest is the request type for the
// TestEventPattern API operation.
type TestEventPatternRequest struct {
	*aws.Request
	Input *TestEventPatternInput
	Copy  func(*TestEventPatternInput) TestEventPatternRequest
}

// Send marshals and sends the TestEventPattern API request.
func (r TestEventPatternRequest) Send(ctx context.Context) (*TestEventPatternResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestEventPatternResponse{
		TestEventPatternOutput: r.Request.Data.(*TestEventPatternOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestEventPatternResponse is the response type for the
// TestEventPattern API operation.
type TestEventPatternResponse struct {
	*TestEventPatternOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestEventPattern request.
func (r *TestEventPatternResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
