// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/ListWorkGroupsInput
type ListWorkGroupsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of workgroups to return in this request.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to be used by the next request if this request is truncated.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListWorkGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListWorkGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListWorkGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/ListWorkGroupsOutput
type ListWorkGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A token to be used by the next request if this request is truncated.
	NextToken *string `json:"athena:ListWorkGroupsOutput:NextToken" min:"1" type:"string"`

	// The list of workgroups, including their names, descriptions, creation times,
	// and states.
	WorkGroups []WorkGroupSummary `json:"athena:ListWorkGroupsOutput:WorkGroups" type:"list"`
}

// String returns the string representation
func (s ListWorkGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListWorkGroups = "ListWorkGroups"

// ListWorkGroupsRequest returns a request value for making API operation for
// Amazon Athena.
//
// Lists available workgroups for the account.
//
//    // Example sending a request using ListWorkGroupsRequest.
//    req := client.ListWorkGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/ListWorkGroups
func (c *Client) ListWorkGroupsRequest(input *ListWorkGroupsInput) ListWorkGroupsRequest {
	op := &aws.Operation{
		Name:       opListWorkGroups,
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
		input = &ListWorkGroupsInput{}
	}

	req := c.newRequest(op, input, &ListWorkGroupsOutput{})
	return ListWorkGroupsRequest{Request: req, Input: input, Copy: c.ListWorkGroupsRequest}
}

// ListWorkGroupsRequest is the request type for the
// ListWorkGroups API operation.
type ListWorkGroupsRequest struct {
	*aws.Request
	Input *ListWorkGroupsInput
	Copy  func(*ListWorkGroupsInput) ListWorkGroupsRequest
}

// Send marshals and sends the ListWorkGroups API request.
func (r ListWorkGroupsRequest) Send(ctx context.Context) (*ListWorkGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListWorkGroupsResponse{
		ListWorkGroupsOutput: r.Request.Data.(*ListWorkGroupsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListWorkGroupsRequestPaginator returns a paginator for ListWorkGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListWorkGroupsRequest(input)
//   p := athena.NewListWorkGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListWorkGroupsPaginator(req ListWorkGroupsRequest) ListWorkGroupsPaginator {
	return ListWorkGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListWorkGroupsInput
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

// ListWorkGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListWorkGroupsPaginator struct {
	aws.Pager
}

func (p *ListWorkGroupsPaginator) CurrentPage() *ListWorkGroupsOutput {
	return p.Pager.CurrentPage().(*ListWorkGroupsOutput)
}

// ListWorkGroupsResponse is the response type for the
// ListWorkGroups API operation.
type ListWorkGroupsResponse struct {
	*ListWorkGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListWorkGroups request.
func (r *ListWorkGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
