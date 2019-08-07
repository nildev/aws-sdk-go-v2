// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/ListServicesRequest
type ListServicesInput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains specifications for the namespaces that you want
	// to list services for.
	//
	// If you specify more than one filter, an operation must match all filters
	// to be returned by ListServices.
	Filters []ServiceFilter `type:"list"`

	// The maximum number of services that you want AWS Cloud Map to return in the
	// response to a ListServices request. If you don't specify a value for MaxResults,
	// AWS Cloud Map returns up to 100 services.
	MaxResults *int64 `min:"1" type:"integer"`

	// For the first ListServices request, omit this value.
	//
	// If the response contains NextToken, submit another ListServices request to
	// get the next group of results. Specify the value of NextToken from the previous
	// response in the next request.
	//
	// AWS Cloud Map gets MaxResults services and then filters them based on the
	// specified criteria. It's possible that no services in the first MaxResults
	// services matched the specified criteria but that subsequent groups of MaxResults
	// services do contain services that match the criteria.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListServicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListServicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListServicesInput"}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/ListServicesResponse
type ListServicesOutput struct {
	_ struct{} `type:"structure"`

	// If the response contains NextToken, submit another ListServices request to
	// get the next group of results. Specify the value of NextToken from the previous
	// response in the next request.
	//
	// AWS Cloud Map gets MaxResults services and then filters them based on the
	// specified criteria. It's possible that no services in the first MaxResults
	// services matched the specified criteria but that subsequent groups of MaxResults
	// services do contain services that match the criteria.
	NextToken *string `json:"servicediscovery:ListServicesOutput:NextToken" type:"string"`

	// An array that contains one ServiceSummary object for each service that matches
	// the specified filter criteria.
	Services []ServiceSummary `json:"servicediscovery:ListServicesOutput:Services" type:"list"`
}

// String returns the string representation
func (s ListServicesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListServices = "ListServices"

// ListServicesRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Lists summary information for all the services that are associated with one
// or more specified namespaces.
//
//    // Example sending a request using ListServicesRequest.
//    req := client.ListServicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/ListServices
func (c *Client) ListServicesRequest(input *ListServicesInput) ListServicesRequest {
	op := &aws.Operation{
		Name:       opListServices,
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
		input = &ListServicesInput{}
	}

	req := c.newRequest(op, input, &ListServicesOutput{})
	return ListServicesRequest{Request: req, Input: input, Copy: c.ListServicesRequest}
}

// ListServicesRequest is the request type for the
// ListServices API operation.
type ListServicesRequest struct {
	*aws.Request
	Input *ListServicesInput
	Copy  func(*ListServicesInput) ListServicesRequest
}

// Send marshals and sends the ListServices API request.
func (r ListServicesRequest) Send(ctx context.Context) (*ListServicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServicesResponse{
		ListServicesOutput: r.Request.Data.(*ListServicesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListServicesRequestPaginator returns a paginator for ListServices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListServicesRequest(input)
//   p := servicediscovery.NewListServicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListServicesPaginator(req ListServicesRequest) ListServicesPaginator {
	return ListServicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListServicesInput
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

// ListServicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListServicesPaginator struct {
	aws.Pager
}

func (p *ListServicesPaginator) CurrentPage() *ListServicesOutput {
	return p.Pager.CurrentPage().(*ListServicesOutput)
}

// ListServicesResponse is the response type for the
// ListServices API operation.
type ListServicesResponse struct {
	*ListServicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServices request.
func (r *ListServicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
