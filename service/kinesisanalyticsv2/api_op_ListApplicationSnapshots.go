// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/ListApplicationSnapshotsRequest
type ListApplicationSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// The name of an existing application.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The maximum number of application snapshots to list.
	Limit *int64 `min:"1" type:"integer"`

	// Use this parameter if you receive a NextToken response in a previous request
	// that indicates that there is more output available. Set it to the value of
	// the previous call's NextToken response to indicate where the output should
	// continue from.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListApplicationSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListApplicationSnapshotsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListApplicationSnapshotsInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/ListApplicationSnapshotsResponse
type ListApplicationSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string `json:"kinesisanalytics:ListApplicationSnapshotsOutput:NextToken" min:"1" type:"string"`

	// A collection of objects containing information about the application snapshots.
	SnapshotSummaries []SnapshotDetails `json:"kinesisanalytics:ListApplicationSnapshotsOutput:SnapshotSummaries" type:"list"`
}

// String returns the string representation
func (s ListApplicationSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListApplicationSnapshots = "ListApplicationSnapshots"

// ListApplicationSnapshotsRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Lists information about the current application snapshots.
//
//    // Example sending a request using ListApplicationSnapshotsRequest.
//    req := client.ListApplicationSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/ListApplicationSnapshots
func (c *Client) ListApplicationSnapshotsRequest(input *ListApplicationSnapshotsInput) ListApplicationSnapshotsRequest {
	op := &aws.Operation{
		Name:       opListApplicationSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListApplicationSnapshotsInput{}
	}

	req := c.newRequest(op, input, &ListApplicationSnapshotsOutput{})
	return ListApplicationSnapshotsRequest{Request: req, Input: input, Copy: c.ListApplicationSnapshotsRequest}
}

// ListApplicationSnapshotsRequest is the request type for the
// ListApplicationSnapshots API operation.
type ListApplicationSnapshotsRequest struct {
	*aws.Request
	Input *ListApplicationSnapshotsInput
	Copy  func(*ListApplicationSnapshotsInput) ListApplicationSnapshotsRequest
}

// Send marshals and sends the ListApplicationSnapshots API request.
func (r ListApplicationSnapshotsRequest) Send(ctx context.Context) (*ListApplicationSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListApplicationSnapshotsResponse{
		ListApplicationSnapshotsOutput: r.Request.Data.(*ListApplicationSnapshotsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListApplicationSnapshotsResponse is the response type for the
// ListApplicationSnapshots API operation.
type ListApplicationSnapshotsResponse struct {
	*ListApplicationSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListApplicationSnapshots request.
func (r *ListApplicationSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
