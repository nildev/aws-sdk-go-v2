// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListTagsForResourceInput
type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of tag results returned by ListTagsForResource in paginated
	// output. When this parameter is used, ListTagsForResource returns only limit
	// results in a single page along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another ListTagsForResource
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If you don't use this parameter, ListTagsForResource returns up to 100
	// results and a nextToken value if applicable.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated ListTagsForResource
	// request where limit was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The Amazon Resource Name (ARN) that identifies the resource to list the tags
	// for.
	//
	// ResourceArn is a required field
	ResourceArn *string `location:"querystring" locationName:"resourceArn" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForResourceInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsForResourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListTagsForResourceOutput
type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListTagsForResource request. When
	// the results of a ListTagsForResource request exceed limit, you can use this
	// value to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The tags for the resource.
	//
	// Tags is a required field
	Tags []TagRef `locationName:"tags" type:"list" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsForResourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListTagsForResource = "ListTagsForResource"

// ListTagsForResourceRequest returns a request value for making API operation for
// AWS App Mesh.
//
// List the tags for an App Mesh resource.
//
//    // Example sending a request using ListTagsForResourceRequest.
//    req := client.ListTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListTagsForResource
func (c *Client) ListTagsForResourceRequest(input *ListTagsForResourceInput) ListTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opListTagsForResource,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/tags",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &ListTagsForResourceOutput{})
	return ListTagsForResourceRequest{Request: req, Input: input, Copy: c.ListTagsForResourceRequest}
}

// ListTagsForResourceRequest is the request type for the
// ListTagsForResource API operation.
type ListTagsForResourceRequest struct {
	*aws.Request
	Input *ListTagsForResourceInput
	Copy  func(*ListTagsForResourceInput) ListTagsForResourceRequest
}

// Send marshals and sends the ListTagsForResource API request.
func (r ListTagsForResourceRequest) Send(ctx context.Context) (*ListTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForResourceResponse{
		ListTagsForResourceOutput: r.Request.Data.(*ListTagsForResourceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTagsForResourceRequestPaginator returns a paginator for ListTagsForResource.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTagsForResourceRequest(input)
//   p := appmesh.NewListTagsForResourceRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTagsForResourcePaginator(req ListTagsForResourceRequest) ListTagsForResourcePaginator {
	return ListTagsForResourcePaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTagsForResourceInput
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

// ListTagsForResourcePaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTagsForResourcePaginator struct {
	aws.Pager
}

func (p *ListTagsForResourcePaginator) CurrentPage() *ListTagsForResourceOutput {
	return p.Pager.CurrentPage().(*ListTagsForResourceOutput)
}

// ListTagsForResourceResponse is the response type for the
// ListTagsForResource API operation.
type ListTagsForResourceResponse struct {
	*ListTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForResource request.
func (r *ListTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
