// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The ListDomains request includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/ListDomainsRequest
type ListDomainsInput struct {
	_ struct{} `type:"structure"`

	// For an initial request for a list of domains, omit this element. If the number
	// of domains that are associated with the current AWS account is greater than
	// the value that you specified for MaxItems, you can use Marker to return additional
	// domains. Get the value of NextPageMarker from the previous response, and
	// submit another request that includes the value of NextPageMarker in the Marker
	// element.
	//
	// Constraints: The marker must match the value specified in the previous request.
	Marker *string `type:"string"`

	// Number of domains to be returned.
	//
	// Default: 20
	MaxItems *int64 `type:"integer"`
}

// String returns the string representation
func (s ListDomainsInput) String() string {
	return awsutil.Prettify(s)
}

// The ListDomains response includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/ListDomainsResponse
type ListDomainsOutput struct {
	_ struct{} `type:"structure"`

	// A summary of domains.
	//
	// Domains is a required field
	Domains []DomainSummary `json:"route53domains:ListDomainsOutput:Domains" type:"list" required:"true"`

	// If there are more domains than you specified for MaxItems in the request,
	// submit another request and include the value of NextPageMarker in the value
	// of Marker.
	NextPageMarker *string `json:"route53domains:ListDomainsOutput:NextPageMarker" type:"string"`
}

// String returns the string representation
func (s ListDomainsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDomains = "ListDomains"

// ListDomainsRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation returns all the domain names registered with Amazon Route
// 53 for the current AWS account.
//
//    // Example sending a request using ListDomainsRequest.
//    req := client.ListDomainsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/ListDomains
func (c *Client) ListDomainsRequest(input *ListDomainsInput) ListDomainsRequest {
	op := &aws.Operation{
		Name:       opListDomains,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextPageMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDomainsInput{}
	}

	req := c.newRequest(op, input, &ListDomainsOutput{})
	return ListDomainsRequest{Request: req, Input: input, Copy: c.ListDomainsRequest}
}

// ListDomainsRequest is the request type for the
// ListDomains API operation.
type ListDomainsRequest struct {
	*aws.Request
	Input *ListDomainsInput
	Copy  func(*ListDomainsInput) ListDomainsRequest
}

// Send marshals and sends the ListDomains API request.
func (r ListDomainsRequest) Send(ctx context.Context) (*ListDomainsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDomainsResponse{
		ListDomainsOutput: r.Request.Data.(*ListDomainsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDomainsRequestPaginator returns a paginator for ListDomains.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDomainsRequest(input)
//   p := route53domains.NewListDomainsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDomainsPaginator(req ListDomainsRequest) ListDomainsPaginator {
	return ListDomainsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDomainsInput
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

// ListDomainsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDomainsPaginator struct {
	aws.Pager
}

func (p *ListDomainsPaginator) CurrentPage() *ListDomainsOutput {
	return p.Pager.CurrentPage().(*ListDomainsOutput)
}

// ListDomainsResponse is the response type for the
// ListDomains API operation.
type ListDomainsResponse struct {
	*ListDomainsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDomains request.
func (r *ListDomainsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
