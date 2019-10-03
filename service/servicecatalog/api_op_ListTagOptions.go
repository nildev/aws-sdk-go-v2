// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListTagOptionsInput
type ListTagOptionsInput struct {
	_ struct{} `type:"structure"`

	// The search filters. If no search filters are specified, the output includes
	// all TagOptions.
	Filters *ListTagOptionsFilters `type:"structure"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`
}

// String returns the string representation
func (s ListTagOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagOptionsInput"}
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			invalidParams.AddNested("Filters", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListTagOptionsOutput
type ListTagOptionsOutput struct {
	_ struct{} `type:"structure"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `json:"servicecatalog:ListTagOptionsOutput:PageToken" type:"string"`

	// Information about the TagOptions.
	TagOptionDetails []TagOptionDetail `json:"servicecatalog:ListTagOptionsOutput:TagOptionDetails" type:"list"`
}

// String returns the string representation
func (s ListTagOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagOptions = "ListTagOptions"

// ListTagOptionsRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists the specified TagOptions or all TagOptions.
//
//    // Example sending a request using ListTagOptionsRequest.
//    req := client.ListTagOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListTagOptions
func (c *Client) ListTagOptionsRequest(input *ListTagOptionsInput) ListTagOptionsRequest {
	op := &aws.Operation{
		Name:       opListTagOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"PageToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTagOptionsInput{}
	}

	req := c.newRequest(op, input, &ListTagOptionsOutput{})
	return ListTagOptionsRequest{Request: req, Input: input, Copy: c.ListTagOptionsRequest}
}

// ListTagOptionsRequest is the request type for the
// ListTagOptions API operation.
type ListTagOptionsRequest struct {
	*aws.Request
	Input *ListTagOptionsInput
	Copy  func(*ListTagOptionsInput) ListTagOptionsRequest
}

// Send marshals and sends the ListTagOptions API request.
func (r ListTagOptionsRequest) Send(ctx context.Context) (*ListTagOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagOptionsResponse{
		ListTagOptionsOutput: r.Request.Data.(*ListTagOptionsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTagOptionsRequestPaginator returns a paginator for ListTagOptions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTagOptionsRequest(input)
//   p := servicecatalog.NewListTagOptionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTagOptionsPaginator(req ListTagOptionsRequest) ListTagOptionsPaginator {
	return ListTagOptionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTagOptionsInput
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

// ListTagOptionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTagOptionsPaginator struct {
	aws.Pager
}

func (p *ListTagOptionsPaginator) CurrentPage() *ListTagOptionsOutput {
	return p.Pager.CurrentPage().(*ListTagOptionsOutput)
}

// ListTagOptionsResponse is the response type for the
// ListTagOptions API operation.
type ListTagOptionsResponse struct {
	*ListTagOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagOptions request.
func (r *ListTagOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
