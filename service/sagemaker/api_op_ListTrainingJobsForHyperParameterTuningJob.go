// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListTrainingJobsForHyperParameterTuningJobRequest
type ListTrainingJobsForHyperParameterTuningJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the tuning job whose training jobs you want to list.
	//
	// HyperParameterTuningJobName is a required field
	HyperParameterTuningJobName *string `min:"1" type:"string" required:"true"`

	// The maximum number of training jobs to return. The default value is 10.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the result of the previous ListTrainingJobsForHyperParameterTuningJob
	// request was truncated, the response includes a NextToken. To retrieve the
	// next set of training jobs, use the token in the next request.
	NextToken *string `type:"string"`

	// The field to sort results by. The default is Name.
	//
	// If the value of this field is FinalObjectiveMetricValue, any training jobs
	// that did not return an objective metric are not listed.
	SortBy TrainingJobSortByOptions `type:"string" enum:"true"`

	// The sort order for results. The default is Ascending.
	SortOrder SortOrder `type:"string" enum:"true"`

	// A filter that returns only training jobs with the specified status.
	StatusEquals TrainingJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListTrainingJobsForHyperParameterTuningJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTrainingJobsForHyperParameterTuningJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTrainingJobsForHyperParameterTuningJobInput"}

	if s.HyperParameterTuningJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("HyperParameterTuningJobName"))
	}
	if s.HyperParameterTuningJobName != nil && len(*s.HyperParameterTuningJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HyperParameterTuningJobName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListTrainingJobsForHyperParameterTuningJobResponse
type ListTrainingJobsForHyperParameterTuningJobOutput struct {
	_ struct{} `type:"structure"`

	// If the result of this ListTrainingJobsForHyperParameterTuningJob request
	// was truncated, the response includes a NextToken. To retrieve the next set
	// of training jobs, use the token in the next request.
	NextToken *string `json:"api.sagemaker:ListTrainingJobsForHyperParameterTuningJobOutput:NextToken" type:"string"`

	// A list of TrainingJobSummary objects that describe the training jobs that
	// the ListTrainingJobsForHyperParameterTuningJob request returned.
	//
	// TrainingJobSummaries is a required field
	TrainingJobSummaries []HyperParameterTrainingJobSummary `json:"api.sagemaker:ListTrainingJobsForHyperParameterTuningJobOutput:TrainingJobSummaries" type:"list" required:"true"`
}

// String returns the string representation
func (s ListTrainingJobsForHyperParameterTuningJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTrainingJobsForHyperParameterTuningJob = "ListTrainingJobsForHyperParameterTuningJob"

// ListTrainingJobsForHyperParameterTuningJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Gets a list of TrainingJobSummary objects that describe the training jobs
// that a hyperparameter tuning job launched.
//
//    // Example sending a request using ListTrainingJobsForHyperParameterTuningJobRequest.
//    req := client.ListTrainingJobsForHyperParameterTuningJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListTrainingJobsForHyperParameterTuningJob
func (c *Client) ListTrainingJobsForHyperParameterTuningJobRequest(input *ListTrainingJobsForHyperParameterTuningJobInput) ListTrainingJobsForHyperParameterTuningJobRequest {
	op := &aws.Operation{
		Name:       opListTrainingJobsForHyperParameterTuningJob,
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
		input = &ListTrainingJobsForHyperParameterTuningJobInput{}
	}

	req := c.newRequest(op, input, &ListTrainingJobsForHyperParameterTuningJobOutput{})
	return ListTrainingJobsForHyperParameterTuningJobRequest{Request: req, Input: input, Copy: c.ListTrainingJobsForHyperParameterTuningJobRequest}
}

// ListTrainingJobsForHyperParameterTuningJobRequest is the request type for the
// ListTrainingJobsForHyperParameterTuningJob API operation.
type ListTrainingJobsForHyperParameterTuningJobRequest struct {
	*aws.Request
	Input *ListTrainingJobsForHyperParameterTuningJobInput
	Copy  func(*ListTrainingJobsForHyperParameterTuningJobInput) ListTrainingJobsForHyperParameterTuningJobRequest
}

// Send marshals and sends the ListTrainingJobsForHyperParameterTuningJob API request.
func (r ListTrainingJobsForHyperParameterTuningJobRequest) Send(ctx context.Context) (*ListTrainingJobsForHyperParameterTuningJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTrainingJobsForHyperParameterTuningJobResponse{
		ListTrainingJobsForHyperParameterTuningJobOutput: r.Request.Data.(*ListTrainingJobsForHyperParameterTuningJobOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTrainingJobsForHyperParameterTuningJobRequestPaginator returns a paginator for ListTrainingJobsForHyperParameterTuningJob.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTrainingJobsForHyperParameterTuningJobRequest(input)
//   p := sagemaker.NewListTrainingJobsForHyperParameterTuningJobRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTrainingJobsForHyperParameterTuningJobPaginator(req ListTrainingJobsForHyperParameterTuningJobRequest) ListTrainingJobsForHyperParameterTuningJobPaginator {
	return ListTrainingJobsForHyperParameterTuningJobPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTrainingJobsForHyperParameterTuningJobInput
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

// ListTrainingJobsForHyperParameterTuningJobPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTrainingJobsForHyperParameterTuningJobPaginator struct {
	aws.Pager
}

func (p *ListTrainingJobsForHyperParameterTuningJobPaginator) CurrentPage() *ListTrainingJobsForHyperParameterTuningJobOutput {
	return p.Pager.CurrentPage().(*ListTrainingJobsForHyperParameterTuningJobOutput)
}

// ListTrainingJobsForHyperParameterTuningJobResponse is the response type for the
// ListTrainingJobsForHyperParameterTuningJob API operation.
type ListTrainingJobsForHyperParameterTuningJobResponse struct {
	*ListTrainingJobsForHyperParameterTuningJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTrainingJobsForHyperParameterTuningJob request.
func (r *ListTrainingJobsForHyperParameterTuningJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
