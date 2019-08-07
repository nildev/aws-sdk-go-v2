// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/ListDatasetImportJobsRequest
type ListDatasetImportJobsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset to list the dataset import
	// jobs for.
	DatasetArn *string `locationName:"datasetArn" type:"string"`

	// The maximum number of dataset import jobs to return.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// A token returned from the previous call to ListDatasetImportJobs for getting
	// the next set of dataset import jobs (if they exist).
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatasetImportJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDatasetImportJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDatasetImportJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/ListDatasetImportJobsResponse
type ListDatasetImportJobsOutput struct {
	_ struct{} `type:"structure"`

	// The list of dataset import jobs.
	DatasetImportJobs []DatasetImportJobSummary `json:"personalize:ListDatasetImportJobsOutput:DatasetImportJobs" locationName:"datasetImportJobs" type:"list"`

	// A token for getting the next set of dataset import jobs (if they exist).
	NextToken *string `json:"personalize:ListDatasetImportJobsOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatasetImportJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDatasetImportJobs = "ListDatasetImportJobs"

// ListDatasetImportJobsRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Returns a list of dataset import jobs that use the given dataset. When a
// dataset is not specified, all the dataset import jobs associated with the
// account are listed. The response provides the properties for each dataset
// import job, including the Amazon Resource Name (ARN). For more information
// on dataset import jobs, see CreateDatasetImportJob. For more information
// on datasets, see CreateDataset.
//
//    // Example sending a request using ListDatasetImportJobsRequest.
//    req := client.ListDatasetImportJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/ListDatasetImportJobs
func (c *Client) ListDatasetImportJobsRequest(input *ListDatasetImportJobsInput) ListDatasetImportJobsRequest {
	op := &aws.Operation{
		Name:       opListDatasetImportJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDatasetImportJobsInput{}
	}

	req := c.newRequest(op, input, &ListDatasetImportJobsOutput{})
	return ListDatasetImportJobsRequest{Request: req, Input: input, Copy: c.ListDatasetImportJobsRequest}
}

// ListDatasetImportJobsRequest is the request type for the
// ListDatasetImportJobs API operation.
type ListDatasetImportJobsRequest struct {
	*aws.Request
	Input *ListDatasetImportJobsInput
	Copy  func(*ListDatasetImportJobsInput) ListDatasetImportJobsRequest
}

// Send marshals and sends the ListDatasetImportJobs API request.
func (r ListDatasetImportJobsRequest) Send(ctx context.Context) (*ListDatasetImportJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDatasetImportJobsResponse{
		ListDatasetImportJobsOutput: r.Request.Data.(*ListDatasetImportJobsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDatasetImportJobsRequestPaginator returns a paginator for ListDatasetImportJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDatasetImportJobsRequest(input)
//   p := personalize.NewListDatasetImportJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDatasetImportJobsPaginator(req ListDatasetImportJobsRequest) ListDatasetImportJobsPaginator {
	return ListDatasetImportJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDatasetImportJobsInput
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

// ListDatasetImportJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDatasetImportJobsPaginator struct {
	aws.Pager
}

func (p *ListDatasetImportJobsPaginator) CurrentPage() *ListDatasetImportJobsOutput {
	return p.Pager.CurrentPage().(*ListDatasetImportJobsOutput)
}

// ListDatasetImportJobsResponse is the response type for the
// ListDatasetImportJobs API operation.
type ListDatasetImportJobsResponse struct {
	*ListDatasetImportJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDatasetImportJobs request.
func (r *ListDatasetImportJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
