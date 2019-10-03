// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datapipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for QueryObjects.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/QueryObjectsInput
type QueryObjectsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of object names that QueryObjects will return in a single
	// call. The default value is 100.
	Limit *int64 `locationName:"limit" type:"integer"`

	// The starting point for the results to be returned. For the first call, this
	// value should be empty. As long as there are more results, continue to call
	// QueryObjects with the marker value from the previous call to retrieve the
	// next set of results.
	Marker *string `locationName:"marker" type:"string"`

	// The ID of the pipeline.
	//
	// PipelineId is a required field
	PipelineId *string `locationName:"pipelineId" min:"1" type:"string" required:"true"`

	// The query that defines the objects to be returned. The Query object can contain
	// a maximum of ten selectors. The conditions in the query are limited to top-level
	// String fields in the object. These filters can be applied to components,
	// instances, and attempts.
	Query *Query `locationName:"query" type:"structure"`

	// Indicates whether the query applies to components or instances. The possible
	// values are: COMPONENT, INSTANCE, and ATTEMPT.
	//
	// Sphere is a required field
	Sphere *string `locationName:"sphere" type:"string" required:"true"`
}

// String returns the string representation
func (s QueryObjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryObjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "QueryObjectsInput"}

	if s.PipelineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineId"))
	}
	if s.PipelineId != nil && len(*s.PipelineId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineId", 1))
	}

	if s.Sphere == nil {
		invalidParams.Add(aws.NewErrParamRequired("Sphere"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of QueryObjects.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/QueryObjectsOutput
type QueryObjectsOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether there are more results that can be obtained by a subsequent
	// call.
	HasMoreResults *bool `json:"datapipeline:QueryObjectsOutput:HasMoreResults" locationName:"hasMoreResults" type:"boolean"`

	// The identifiers that match the query selectors.
	Ids []string `json:"datapipeline:QueryObjectsOutput:Ids" locationName:"ids" type:"list"`

	// The starting point for the next page of results. To view the next page of
	// results, call QueryObjects again with this marker value. If the value is
	// null, there are no more results.
	Marker *string `json:"datapipeline:QueryObjectsOutput:Marker" locationName:"marker" type:"string"`
}

// String returns the string representation
func (s QueryObjectsOutput) String() string {
	return awsutil.Prettify(s)
}

const opQueryObjects = "QueryObjects"

// QueryObjectsRequest returns a request value for making API operation for
// AWS Data Pipeline.
//
// Queries the specified pipeline for the names of objects that match the specified
// set of conditions.
//
//    // Example sending a request using QueryObjectsRequest.
//    req := client.QueryObjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/QueryObjects
func (c *Client) QueryObjectsRequest(input *QueryObjectsInput) QueryObjectsRequest {
	op := &aws.Operation{
		Name:       opQueryObjects,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"marker"},
			OutputTokens:    []string{"marker"},
			LimitToken:      "limit",
			TruncationToken: "hasMoreResults",
		},
	}

	if input == nil {
		input = &QueryObjectsInput{}
	}

	req := c.newRequest(op, input, &QueryObjectsOutput{})
	return QueryObjectsRequest{Request: req, Input: input, Copy: c.QueryObjectsRequest}
}

// QueryObjectsRequest is the request type for the
// QueryObjects API operation.
type QueryObjectsRequest struct {
	*aws.Request
	Input *QueryObjectsInput
	Copy  func(*QueryObjectsInput) QueryObjectsRequest
}

// Send marshals and sends the QueryObjects API request.
func (r QueryObjectsRequest) Send(ctx context.Context) (*QueryObjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &QueryObjectsResponse{
		QueryObjectsOutput: r.Request.Data.(*QueryObjectsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewQueryObjectsRequestPaginator returns a paginator for QueryObjects.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.QueryObjectsRequest(input)
//   p := datapipeline.NewQueryObjectsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewQueryObjectsPaginator(req QueryObjectsRequest) QueryObjectsPaginator {
	return QueryObjectsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *QueryObjectsInput
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

// QueryObjectsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type QueryObjectsPaginator struct {
	aws.Pager
}

func (p *QueryObjectsPaginator) CurrentPage() *QueryObjectsOutput {
	return p.Pager.CurrentPage().(*QueryObjectsOutput)
}

// QueryObjectsResponse is the response type for the
// QueryObjects API operation.
type QueryObjectsResponse struct {
	*QueryObjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// QueryObjects request.
func (r *QueryObjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
