// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/ListDevEndpointsRequest
type ListDevEndpointsInput struct {
	_ struct{} `type:"structure"`

	// The maximum size of a list to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation request.
	NextToken *string `type:"string"`

	// Specifies to return only these tagged resources.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s ListDevEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDevEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDevEndpointsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/ListDevEndpointsResponse
type ListDevEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// The names of all the DevEndpoints in the account, or the DevEndpoints with
	// the specified tags.
	DevEndpointNames []string `json:"glue:ListDevEndpointsOutput:DevEndpointNames" type:"list"`

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string `json:"glue:ListDevEndpointsOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s ListDevEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDevEndpoints = "ListDevEndpoints"

// ListDevEndpointsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the names of all DevEndpoint resources in this AWS account, or
// the resources with the specified tag. This operation allows you to see which
// resources are available in your account, and their names.
//
// This operation takes the optional Tags field, which you can use as a filter
// on the response so that tagged resources can be retrieved as a group. If
// you choose to use tags filtering, only resources with the tag are retrieved.
//
//    // Example sending a request using ListDevEndpointsRequest.
//    req := client.ListDevEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/ListDevEndpoints
func (c *Client) ListDevEndpointsRequest(input *ListDevEndpointsInput) ListDevEndpointsRequest {
	op := &aws.Operation{
		Name:       opListDevEndpoints,
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
		input = &ListDevEndpointsInput{}
	}

	req := c.newRequest(op, input, &ListDevEndpointsOutput{})
	return ListDevEndpointsRequest{Request: req, Input: input, Copy: c.ListDevEndpointsRequest}
}

// ListDevEndpointsRequest is the request type for the
// ListDevEndpoints API operation.
type ListDevEndpointsRequest struct {
	*aws.Request
	Input *ListDevEndpointsInput
	Copy  func(*ListDevEndpointsInput) ListDevEndpointsRequest
}

// Send marshals and sends the ListDevEndpoints API request.
func (r ListDevEndpointsRequest) Send(ctx context.Context) (*ListDevEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDevEndpointsResponse{
		ListDevEndpointsOutput: r.Request.Data.(*ListDevEndpointsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDevEndpointsRequestPaginator returns a paginator for ListDevEndpoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDevEndpointsRequest(input)
//   p := glue.NewListDevEndpointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDevEndpointsPaginator(req ListDevEndpointsRequest) ListDevEndpointsPaginator {
	return ListDevEndpointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDevEndpointsInput
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

// ListDevEndpointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDevEndpointsPaginator struct {
	aws.Pager
}

func (p *ListDevEndpointsPaginator) CurrentPage() *ListDevEndpointsOutput {
	return p.Pager.CurrentPage().(*ListDevEndpointsOutput)
}

// ListDevEndpointsResponse is the response type for the
// ListDevEndpoints API operation.
type ListDevEndpointsResponse struct {
	*ListDevEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDevEndpoints request.
func (r *ListDevEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
