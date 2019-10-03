// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListEndpointsInput
type ListEndpointsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only endpoints with a creation time greater than or
	// equal to the specified time (timestamp).
	CreationTimeAfter *time.Time `type:"timestamp" timestampFormat:"unix"`

	// A filter that returns only endpoints that were created before the specified
	// time (timestamp).
	CreationTimeBefore *time.Time `type:"timestamp" timestampFormat:"unix"`

	// A filter that returns only endpoints that were modified after the specified
	// timestamp.
	LastModifiedTimeAfter *time.Time `type:"timestamp" timestampFormat:"unix"`

	// A filter that returns only endpoints that were modified before the specified
	// timestamp.
	LastModifiedTimeBefore *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The maximum number of endpoints to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in endpoint names. This filter returns only endpoints whose name
	// contains the specified string.
	NameContains *string `type:"string"`

	// If the result of a ListEndpoints request was truncated, the response includes
	// a NextToken. To retrieve the next set of endpoints, use the token in the
	// next request.
	NextToken *string `type:"string"`

	// Sorts the list of results. The default is CreationTime.
	SortBy EndpointSortKey `type:"string" enum:"true"`

	// The sort order for results. The default is Descending.
	SortOrder OrderKey `type:"string" enum:"true"`

	// A filter that returns only endpoints with the specified status.
	StatusEquals EndpointStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEndpointsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListEndpointsOutput
type ListEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// An array or endpoint objects.
	//
	// Endpoints is a required field
	Endpoints []EndpointSummary `json:"api.sagemaker:ListEndpointsOutput:Endpoints" type:"list" required:"true"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of training jobs, use it in the subsequent request.
	NextToken *string `json:"api.sagemaker:ListEndpointsOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s ListEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListEndpoints = "ListEndpoints"

// ListEndpointsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists endpoints.
//
//    // Example sending a request using ListEndpointsRequest.
//    req := client.ListEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListEndpoints
func (c *Client) ListEndpointsRequest(input *ListEndpointsInput) ListEndpointsRequest {
	op := &aws.Operation{
		Name:       opListEndpoints,
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
		input = &ListEndpointsInput{}
	}

	req := c.newRequest(op, input, &ListEndpointsOutput{})
	return ListEndpointsRequest{Request: req, Input: input, Copy: c.ListEndpointsRequest}
}

// ListEndpointsRequest is the request type for the
// ListEndpoints API operation.
type ListEndpointsRequest struct {
	*aws.Request
	Input *ListEndpointsInput
	Copy  func(*ListEndpointsInput) ListEndpointsRequest
}

// Send marshals and sends the ListEndpoints API request.
func (r ListEndpointsRequest) Send(ctx context.Context) (*ListEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEndpointsResponse{
		ListEndpointsOutput: r.Request.Data.(*ListEndpointsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEndpointsRequestPaginator returns a paginator for ListEndpoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEndpointsRequest(input)
//   p := sagemaker.NewListEndpointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEndpointsPaginator(req ListEndpointsRequest) ListEndpointsPaginator {
	return ListEndpointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListEndpointsInput
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

// ListEndpointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEndpointsPaginator struct {
	aws.Pager
}

func (p *ListEndpointsPaginator) CurrentPage() *ListEndpointsOutput {
	return p.Pager.CurrentPage().(*ListEndpointsOutput)
}

// ListEndpointsResponse is the response type for the
// ListEndpoints API operation.
type ListEndpointsResponse struct {
	*ListEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEndpoints request.
func (r *ListEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
