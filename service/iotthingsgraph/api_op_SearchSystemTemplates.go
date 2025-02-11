// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchSystemTemplatesRequest
type SearchSystemTemplatesInput struct {
	_ struct{} `type:"structure"`

	// An array of filters that limit the result set. The only valid filter is FLOW_TEMPLATE_ID.
	Filters []SystemTemplateFilter `locationName:"filters" type:"list"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s SearchSystemTemplatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchSystemTemplatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchSystemTemplatesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchSystemTemplatesResponse
type SearchSystemTemplatesOutput struct {
	_ struct{} `type:"structure"`

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// An array of objects that contain summary information about each system deployment
	// in the result set.
	Summaries []SystemTemplateSummary `locationName:"summaries" type:"list"`
}

// String returns the string representation
func (s SearchSystemTemplatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchSystemTemplates = "SearchSystemTemplates"

// SearchSystemTemplatesRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Searches for summary information about systems in the user's account. You
// can filter by the ID of a workflow to return only systems that use the specified
// workflow.
//
//    // Example sending a request using SearchSystemTemplatesRequest.
//    req := client.SearchSystemTemplatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchSystemTemplates
func (c *Client) SearchSystemTemplatesRequest(input *SearchSystemTemplatesInput) SearchSystemTemplatesRequest {
	op := &aws.Operation{
		Name:       opSearchSystemTemplates,
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
		input = &SearchSystemTemplatesInput{}
	}

	req := c.newRequest(op, input, &SearchSystemTemplatesOutput{})
	return SearchSystemTemplatesRequest{Request: req, Input: input, Copy: c.SearchSystemTemplatesRequest}
}

// SearchSystemTemplatesRequest is the request type for the
// SearchSystemTemplates API operation.
type SearchSystemTemplatesRequest struct {
	*aws.Request
	Input *SearchSystemTemplatesInput
	Copy  func(*SearchSystemTemplatesInput) SearchSystemTemplatesRequest
}

// Send marshals and sends the SearchSystemTemplates API request.
func (r SearchSystemTemplatesRequest) Send(ctx context.Context) (*SearchSystemTemplatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchSystemTemplatesResponse{
		SearchSystemTemplatesOutput: r.Request.Data.(*SearchSystemTemplatesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchSystemTemplatesRequestPaginator returns a paginator for SearchSystemTemplates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchSystemTemplatesRequest(input)
//   p := iotthingsgraph.NewSearchSystemTemplatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchSystemTemplatesPaginator(req SearchSystemTemplatesRequest) SearchSystemTemplatesPaginator {
	return SearchSystemTemplatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *SearchSystemTemplatesInput
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

// SearchSystemTemplatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchSystemTemplatesPaginator struct {
	aws.Pager
}

func (p *SearchSystemTemplatesPaginator) CurrentPage() *SearchSystemTemplatesOutput {
	return p.Pager.CurrentPage().(*SearchSystemTemplatesOutput)
}

// SearchSystemTemplatesResponse is the response type for the
// SearchSystemTemplates API operation.
type SearchSystemTemplatesResponse struct {
	*SearchSystemTemplatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchSystemTemplates request.
func (r *SearchSystemTemplatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
