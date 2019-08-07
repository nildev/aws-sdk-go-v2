// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ListOrganizationsRequest
type ListOrganizationsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to use to retrieve the next page of results. The first call does
	// not contain any tokens.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListOrganizationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListOrganizationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListOrganizationsInput"}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ListOrganizationsResponse
type ListOrganizationsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. The value is "null"
	// when there are no more results to return.
	NextToken *string `json:"workmail:ListOrganizationsOutput:NextToken" min:"1" type:"string"`

	// The overview of owned organizations presented as a list of organization summaries.
	OrganizationSummaries []OrganizationSummary `json:"workmail:ListOrganizationsOutput:OrganizationSummaries" type:"list"`
}

// String returns the string representation
func (s ListOrganizationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListOrganizations = "ListOrganizations"

// ListOrganizationsRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Returns summaries of the customer's non-deleted organizations.
//
//    // Example sending a request using ListOrganizationsRequest.
//    req := client.ListOrganizationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ListOrganizations
func (c *Client) ListOrganizationsRequest(input *ListOrganizationsInput) ListOrganizationsRequest {
	op := &aws.Operation{
		Name:       opListOrganizations,
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
		input = &ListOrganizationsInput{}
	}

	req := c.newRequest(op, input, &ListOrganizationsOutput{})
	return ListOrganizationsRequest{Request: req, Input: input, Copy: c.ListOrganizationsRequest}
}

// ListOrganizationsRequest is the request type for the
// ListOrganizations API operation.
type ListOrganizationsRequest struct {
	*aws.Request
	Input *ListOrganizationsInput
	Copy  func(*ListOrganizationsInput) ListOrganizationsRequest
}

// Send marshals and sends the ListOrganizations API request.
func (r ListOrganizationsRequest) Send(ctx context.Context) (*ListOrganizationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListOrganizationsResponse{
		ListOrganizationsOutput: r.Request.Data.(*ListOrganizationsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListOrganizationsRequestPaginator returns a paginator for ListOrganizations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListOrganizationsRequest(input)
//   p := workmail.NewListOrganizationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListOrganizationsPaginator(req ListOrganizationsRequest) ListOrganizationsPaginator {
	return ListOrganizationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListOrganizationsInput
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

// ListOrganizationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListOrganizationsPaginator struct {
	aws.Pager
}

func (p *ListOrganizationsPaginator) CurrentPage() *ListOrganizationsOutput {
	return p.Pager.CurrentPage().(*ListOrganizationsOutput)
}

// ListOrganizationsResponse is the response type for the
// ListOrganizations API operation.
type ListOrganizationsResponse struct {
	*ListOrganizationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListOrganizations request.
func (r *ListOrganizationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
