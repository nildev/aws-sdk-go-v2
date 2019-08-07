// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// You can send list jobs requests with an empty body. Optionally, you can filter
// the response by queue and/or job status by specifying them in your request
// body. You can also optionally specify the maximum number, up to twenty, of
// jobs to be returned.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/ListJobsRequest
type ListJobsInput struct {
	_ struct{} `type:"structure"`

	// Optional. Number of jobs, up to twenty, that will be returned at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Use this string, provided with the response to a previous request, to request
	// the next batch of jobs.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// When you request lists of resources, you can optionally specify whether they
	// are sorted in ASCENDING or DESCENDING order. Default varies by resource.
	Order Order `location:"querystring" locationName:"order" type:"string" enum:"true"`

	// Provide a queue name to get back only jobs from that queue.
	Queue *string `location:"querystring" locationName:"queue" type:"string"`

	// A job's status can be SUBMITTED, PROGRESSING, COMPLETE, CANCELED, or ERROR.
	Status JobStatus `location:"querystring" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJobsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

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
	if len(s.Order) > 0 {
		v := s.Order

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "order", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Queue != nil {
		v := *s.Queue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "queue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Successful list jobs requests return a JSON array of jobs. If you don't specify
// how they are ordered, you will receive the most recently created first.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/ListJobsResponse
type ListJobsOutput struct {
	_ struct{} `type:"structure"`

	// List of jobs
	Jobs []Job `json:"mediaconvert:ListJobsOutput:Jobs" locationName:"jobs" type:"list"`

	// Use this string to request the next batch of jobs.
	NextToken *string `json:"mediaconvert:ListJobsOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Jobs != nil {
		v := s.Jobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "jobs", metadata)
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

const opListJobs = "ListJobs"

// ListJobsRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Retrieve a JSON array of up to twenty of your most recently created jobs.
// This array includes in-process, completed, and errored jobs. This will return
// the jobs themselves, not just a list of the jobs. To retrieve the twenty
// next most recent jobs, use the nextToken string returned with the array.
//
//    // Example sending a request using ListJobsRequest.
//    req := client.ListJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/ListJobs
func (c *Client) ListJobsRequest(input *ListJobsInput) ListJobsRequest {
	op := &aws.Operation{
		Name:       opListJobs,
		HTTPMethod: "GET",
		HTTPPath:   "/2017-08-29/jobs",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListJobsInput{}
	}

	req := c.newRequest(op, input, &ListJobsOutput{})
	return ListJobsRequest{Request: req, Input: input, Copy: c.ListJobsRequest}
}

// ListJobsRequest is the request type for the
// ListJobs API operation.
type ListJobsRequest struct {
	*aws.Request
	Input *ListJobsInput
	Copy  func(*ListJobsInput) ListJobsRequest
}

// Send marshals and sends the ListJobs API request.
func (r ListJobsRequest) Send(ctx context.Context) (*ListJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListJobsResponse{
		ListJobsOutput: r.Request.Data.(*ListJobsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListJobsRequestPaginator returns a paginator for ListJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListJobsRequest(input)
//   p := mediaconvert.NewListJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListJobsPaginator(req ListJobsRequest) ListJobsPaginator {
	return ListJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListJobsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListJobsPaginator struct {
	aws.Pager
}

func (p *ListJobsPaginator) CurrentPage() *ListJobsOutput {
	return p.Pager.CurrentPage().(*ListJobsOutput)
}

// ListJobsResponse is the response type for the
// ListJobs API operation.
type ListJobsResponse struct {
	*ListJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListJobs request.
func (r *ListJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
