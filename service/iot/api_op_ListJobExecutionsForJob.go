// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListJobExecutionsForJobInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier you assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`

	// The maximum number of results to be returned per request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The status of the job.
	Status JobExecutionStatus `location:"querystring" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListJobExecutionsForJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListJobExecutionsForJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListJobExecutionsForJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJobExecutionsForJobInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type ListJobExecutionsForJobOutput struct {
	_ struct{} `type:"structure"`

	// A list of job execution summaries.
	ExecutionSummaries []JobExecutionSummaryForJob `json:"iot:ListJobExecutionsForJobOutput:ExecutionSummaries" locationName:"executionSummaries" type:"list"`

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string `json:"iot:ListJobExecutionsForJobOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListJobExecutionsForJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJobExecutionsForJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExecutionSummaries != nil {
		v := s.ExecutionSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "executionSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListJobExecutionsForJob = "ListJobExecutionsForJob"

// ListJobExecutionsForJobRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the job executions for a job.
//
//    // Example sending a request using ListJobExecutionsForJobRequest.
//    req := client.ListJobExecutionsForJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListJobExecutionsForJobRequest(input *ListJobExecutionsForJobInput) ListJobExecutionsForJobRequest {
	op := &aws.Operation{
		Name:       opListJobExecutionsForJob,
		HTTPMethod: "GET",
		HTTPPath:   "/jobs/{jobId}/things",
	}

	if input == nil {
		input = &ListJobExecutionsForJobInput{}
	}

	req := c.newRequest(op, input, &ListJobExecutionsForJobOutput{})
	return ListJobExecutionsForJobRequest{Request: req, Input: input, Copy: c.ListJobExecutionsForJobRequest}
}

// ListJobExecutionsForJobRequest is the request type for the
// ListJobExecutionsForJob API operation.
type ListJobExecutionsForJobRequest struct {
	*aws.Request
	Input *ListJobExecutionsForJobInput
	Copy  func(*ListJobExecutionsForJobInput) ListJobExecutionsForJobRequest
}

// Send marshals and sends the ListJobExecutionsForJob API request.
func (r ListJobExecutionsForJobRequest) Send(ctx context.Context) (*ListJobExecutionsForJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListJobExecutionsForJobResponse{
		ListJobExecutionsForJobOutput: r.Request.Data.(*ListJobExecutionsForJobOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListJobExecutionsForJobResponse is the response type for the
// ListJobExecutionsForJob API operation.
type ListJobExecutionsForJobResponse struct {
	*ListJobExecutionsForJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListJobExecutionsForJob request.
func (r *ListJobExecutionsForJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
