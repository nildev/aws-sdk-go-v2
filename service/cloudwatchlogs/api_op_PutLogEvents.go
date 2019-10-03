// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/PutLogEventsRequest
type PutLogEventsInput struct {
	_ struct{} `type:"structure"`

	// The log events.
	//
	// LogEvents is a required field
	LogEvents []InputLogEvent `locationName:"logEvents" min:"1" type:"list" required:"true"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// The name of the log stream.
	//
	// LogStreamName is a required field
	LogStreamName *string `locationName:"logStreamName" min:"1" type:"string" required:"true"`

	// The sequence token obtained from the response of the previous PutLogEvents
	// call. An upload in a newly created log stream does not require a sequence
	// token. You can also get the sequence token using DescribeLogStreams. If you
	// call PutLogEvents twice within a narrow time period using the same value
	// for sequenceToken, both calls may be successful, or one may be rejected.
	SequenceToken *string `locationName:"sequenceToken" min:"1" type:"string"`
}

// String returns the string representation
func (s PutLogEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutLogEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutLogEventsInput"}

	if s.LogEvents == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogEvents"))
	}
	if s.LogEvents != nil && len(s.LogEvents) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogEvents", 1))
	}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if s.LogStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogStreamName"))
	}
	if s.LogStreamName != nil && len(*s.LogStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogStreamName", 1))
	}
	if s.SequenceToken != nil && len(*s.SequenceToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SequenceToken", 1))
	}
	if s.LogEvents != nil {
		for i, v := range s.LogEvents {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "LogEvents", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/PutLogEventsResponse
type PutLogEventsOutput struct {
	_ struct{} `type:"structure"`

	// The next sequence token.
	NextSequenceToken *string `json:"logs:PutLogEventsOutput:NextSequenceToken" locationName:"nextSequenceToken" min:"1" type:"string"`

	// The rejected events.
	RejectedLogEventsInfo *RejectedLogEventsInfo `json:"logs:PutLogEventsOutput:RejectedLogEventsInfo" locationName:"rejectedLogEventsInfo" type:"structure"`
}

// String returns the string representation
func (s PutLogEventsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutLogEvents = "PutLogEvents"

// PutLogEventsRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Uploads a batch of log events to the specified log stream.
//
// You must include the sequence token obtained from the response of the previous
// call. An upload in a newly created log stream does not require a sequence
// token. You can also get the sequence token using DescribeLogStreams. If you
// call PutLogEvents twice within a narrow time period using the same value
// for sequenceToken, both calls may be successful, or one may be rejected.
//
// The batch of events must satisfy the following constraints:
//
//    * The maximum batch size is 1,048,576 bytes, and this size is calculated
//    as the sum of all event messages in UTF-8, plus 26 bytes for each log
//    event.
//
//    * None of the log events in the batch can be more than 2 hours in the
//    future.
//
//    * None of the log events in the batch can be older than 14 days or older
//    than the retention period of the log group.
//
//    * The log events in the batch must be in chronological ordered by their
//    timestamp. The timestamp is the time the event occurred, expressed as
//    the number of milliseconds after Jan 1, 1970 00:00:00 UTC. (In AWS Tools
//    for PowerShell and the AWS SDK for .NET, the timestamp is specified in
//    .NET format: yyyy-mm-ddThh:mm:ss. For example, 2017-09-15T13:45:30.)
//
//    * The maximum number of log events in a batch is 10,000.
//
//    * A batch of log events in a single request cannot span more than 24 hours.
//    Otherwise, the operation fails.
//
// If a call to PutLogEvents returns "UnrecognizedClientException" the most
// likely cause is an invalid AWS access key ID or secret key.
//
//    // Example sending a request using PutLogEventsRequest.
//    req := client.PutLogEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/PutLogEvents
func (c *Client) PutLogEventsRequest(input *PutLogEventsInput) PutLogEventsRequest {
	op := &aws.Operation{
		Name:       opPutLogEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutLogEventsInput{}
	}

	req := c.newRequest(op, input, &PutLogEventsOutput{})
	return PutLogEventsRequest{Request: req, Input: input, Copy: c.PutLogEventsRequest}
}

// PutLogEventsRequest is the request type for the
// PutLogEvents API operation.
type PutLogEventsRequest struct {
	*aws.Request
	Input *PutLogEventsInput
	Copy  func(*PutLogEventsInput) PutLogEventsRequest
}

// Send marshals and sends the PutLogEvents API request.
func (r PutLogEventsRequest) Send(ctx context.Context) (*PutLogEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutLogEventsResponse{
		PutLogEventsOutput: r.Request.Data.(*PutLogEventsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutLogEventsResponse is the response type for the
// PutLogEvents API operation.
type PutLogEventsResponse struct {
	*PutLogEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutLogEvents request.
func (r *PutLogEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
