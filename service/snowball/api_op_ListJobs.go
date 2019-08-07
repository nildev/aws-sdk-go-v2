// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/ListJobsRequest
type ListJobsInput struct {
	_ struct{} `type:"structure"`

	// The number of JobListEntry objects to return.
	MaxResults *int64 `type:"integer"`

	// HTTP requests are stateless. To identify what object comes "next" in the
	// list of JobListEntry objects, you have the option of specifying NextToken
	// as the starting point for your returned list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListJobsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/ListJobsResult
type ListJobsOutput struct {
	_ struct{} `type:"structure"`

	// Each JobListEntry object contains a job's state, a job's ID, and a value
	// that indicates whether the job is a job part, in the case of export jobs.
	JobListEntries []JobListEntry `json:"snowball:ListJobsOutput:JobListEntries" type:"list"`

	// HTTP requests are stateless. If you use this automatically generated NextToken
	// value in your next ListJobs call, your returned JobListEntry objects will
	// start from this point in the array.
	NextToken *string `json:"snowball:ListJobsOutput:NextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListJobs = "ListJobs"

// ListJobsRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// Returns an array of JobListEntry objects of the specified length. Each JobListEntry
// object contains a job's state, a job's ID, and a value that indicates whether
// the job is a job part, in the case of export jobs. Calling this API action
// in one of the US regions will return jobs from the list of all jobs associated
// with this account in all US regions.
//
//    // Example sending a request using ListJobsRequest.
//    req := client.ListJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/ListJobs
func (c *Client) ListJobsRequest(input *ListJobsInput) ListJobsRequest {
	op := &aws.Operation{
		Name:       opListJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
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
//   p := snowball.NewListJobsRequestPaginator(req)
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
