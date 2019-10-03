// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetWorkflowRunsRequest
type GetWorkflowRunsInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether to include the workflow graph in response or not.
	IncludeGraph *bool `type:"boolean"`

	// The maximum number of workflow runs to be included in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// Name of the workflow whose metadata of runs should be returned.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The maximum size of the response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetWorkflowRunsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetWorkflowRunsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetWorkflowRunsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetWorkflowRunsResponse
type GetWorkflowRunsOutput struct {
	_ struct{} `type:"structure"`

	// A continuation token, if not all requested workflow runs have been returned.
	NextToken *string `json:"glue:GetWorkflowRunsOutput:NextToken" type:"string"`

	// A list of workflow run metadata objects.
	Runs []WorkflowRun `json:"glue:GetWorkflowRunsOutput:Runs" min:"1" type:"list"`
}

// String returns the string representation
func (s GetWorkflowRunsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetWorkflowRuns = "GetWorkflowRuns"

// GetWorkflowRunsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves metadata for all runs of a given workflow.
//
//    // Example sending a request using GetWorkflowRunsRequest.
//    req := client.GetWorkflowRunsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetWorkflowRuns
func (c *Client) GetWorkflowRunsRequest(input *GetWorkflowRunsInput) GetWorkflowRunsRequest {
	op := &aws.Operation{
		Name:       opGetWorkflowRuns,
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
		input = &GetWorkflowRunsInput{}
	}

	req := c.newRequest(op, input, &GetWorkflowRunsOutput{})
	return GetWorkflowRunsRequest{Request: req, Input: input, Copy: c.GetWorkflowRunsRequest}
}

// GetWorkflowRunsRequest is the request type for the
// GetWorkflowRuns API operation.
type GetWorkflowRunsRequest struct {
	*aws.Request
	Input *GetWorkflowRunsInput
	Copy  func(*GetWorkflowRunsInput) GetWorkflowRunsRequest
}

// Send marshals and sends the GetWorkflowRuns API request.
func (r GetWorkflowRunsRequest) Send(ctx context.Context) (*GetWorkflowRunsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetWorkflowRunsResponse{
		GetWorkflowRunsOutput: r.Request.Data.(*GetWorkflowRunsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetWorkflowRunsRequestPaginator returns a paginator for GetWorkflowRuns.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetWorkflowRunsRequest(input)
//   p := glue.NewGetWorkflowRunsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetWorkflowRunsPaginator(req GetWorkflowRunsRequest) GetWorkflowRunsPaginator {
	return GetWorkflowRunsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetWorkflowRunsInput
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

// GetWorkflowRunsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetWorkflowRunsPaginator struct {
	aws.Pager
}

func (p *GetWorkflowRunsPaginator) CurrentPage() *GetWorkflowRunsOutput {
	return p.Pager.CurrentPage().(*GetWorkflowRunsOutput)
}

// GetWorkflowRunsResponse is the response type for the
// GetWorkflowRuns API operation.
type GetWorkflowRunsResponse struct {
	*GetWorkflowRunsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetWorkflowRuns request.
func (r *GetWorkflowRunsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
