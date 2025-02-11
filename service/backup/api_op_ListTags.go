// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListTagsInput
type ListTagsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format
	// of the ARN depends on the type of resource. Valid targets for ListTags are
	// recovery points, backup plans, and backup vaults.
	//
	// ResourceArn is a required field
	ResourceArn *string `location:"uri" locationName:"resourceArn" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
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
func (s ListTagsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListTagsOutput
type ListTagsOutput struct {
	_ struct{} `type:"structure"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `type:"string"`

	// To help organize your resources, you can assign your own metadata to the
	// resources you create. Each tag is a key-value pair.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s ListTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTagsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opListTags = "ListTags"

// ListTagsRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns a list of key-value pairs assigned to a target recovery point, backup
// plan, or backup vault.
//
//    // Example sending a request using ListTagsRequest.
//    req := client.ListTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListTags
func (c *Client) ListTagsRequest(input *ListTagsInput) ListTagsRequest {
	op := &aws.Operation{
		Name:       opListTags,
		HTTPMethod: "GET",
		HTTPPath:   "/tags/{resourceArn}/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTagsInput{}
	}

	req := c.newRequest(op, input, &ListTagsOutput{})
	return ListTagsRequest{Request: req, Input: input, Copy: c.ListTagsRequest}
}

// ListTagsRequest is the request type for the
// ListTags API operation.
type ListTagsRequest struct {
	*aws.Request
	Input *ListTagsInput
	Copy  func(*ListTagsInput) ListTagsRequest
}

// Send marshals and sends the ListTags API request.
func (r ListTagsRequest) Send(ctx context.Context) (*ListTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsResponse{
		ListTagsOutput: r.Request.Data.(*ListTagsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTagsRequestPaginator returns a paginator for ListTags.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTagsRequest(input)
//   p := backup.NewListTagsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTagsPaginator(req ListTagsRequest) ListTagsPaginator {
	return ListTagsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTagsInput
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

// ListTagsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTagsPaginator struct {
	aws.Pager
}

func (p *ListTagsPaginator) CurrentPage() *ListTagsOutput {
	return p.Pager.CurrentPage().(*ListTagsOutput)
}

// ListTagsResponse is the response type for the
// ListTags API operation.
type ListTagsResponse struct {
	*ListTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTags request.
func (r *ListTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
