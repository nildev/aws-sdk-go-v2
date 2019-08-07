// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeEventsRequest
type DescribeEventsInput struct {
	_ struct{} `type:"structure"`

	// The number of minutes' worth of events to retrieve.
	Duration *int64 `type:"integer"`

	// The end of the time interval for which to retrieve events, specified in ISO
	// 8601 format.
	EndTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	//
	// The value for MaxResults must be between 20 and 100.
	MaxResults *int64 `type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `type:"string"`

	// The identifier of the event source for which events will be returned. If
	// not specified, then all sources are included in the response.
	SourceName *string `type:"string"`

	// The event source to retrieve events for. If no value is specified, all events
	// are returned.
	SourceType SourceType `type:"string" enum:"true"`

	// The beginning of the time interval to retrieve events for, specified in ISO
	// 8601 format.
	StartTime *time.Time `type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s DescribeEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeEventsResponse
type DescribeEventsOutput struct {
	_ struct{} `type:"structure"`

	// An array of events. Each element in the array represents one event.
	Events []Event `json:"dax:DescribeEventsOutput:Events" type:"list"`

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string `json:"dax:DescribeEventsOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s DescribeEventsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEvents = "DescribeEvents"

// DescribeEventsRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Returns events related to DAX clusters and parameter groups. You can obtain
// events specific to a particular DAX cluster or parameter group by providing
// the name as a parameter.
//
// By default, only the events occurring within the last hour are returned;
// however, you can retrieve up to 14 days' worth of events if necessary.
//
//    // Example sending a request using DescribeEventsRequest.
//    req := client.DescribeEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeEvents
func (c *Client) DescribeEventsRequest(input *DescribeEventsInput) DescribeEventsRequest {
	op := &aws.Operation{
		Name:       opDescribeEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEventsInput{}
	}

	req := c.newRequest(op, input, &DescribeEventsOutput{})
	return DescribeEventsRequest{Request: req, Input: input, Copy: c.DescribeEventsRequest}
}

// DescribeEventsRequest is the request type for the
// DescribeEvents API operation.
type DescribeEventsRequest struct {
	*aws.Request
	Input *DescribeEventsInput
	Copy  func(*DescribeEventsInput) DescribeEventsRequest
}

// Send marshals and sends the DescribeEvents API request.
func (r DescribeEventsRequest) Send(ctx context.Context) (*DescribeEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventsResponse{
		DescribeEventsOutput: r.Request.Data.(*DescribeEventsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEventsResponse is the response type for the
// DescribeEvents API operation.
type DescribeEventsResponse struct {
	*DescribeEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEvents request.
func (r *DescribeEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
