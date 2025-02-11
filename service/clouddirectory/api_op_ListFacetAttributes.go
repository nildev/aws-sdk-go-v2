// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListFacetAttributesRequest
type ListFacetAttributesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to retrieve.
	MaxResults *int64 `min:"1" type:"integer"`

	// The name of the facet whose attributes will be retrieved.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The ARN of the schema where the facet resides.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s ListFacetAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFacetAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFacetAttributesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFacetAttributesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListFacetAttributesResponse
type ListFacetAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The attributes attached to the facet.
	Attributes []FacetAttribute `type:"list"`

	// The pagination token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListFacetAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFacetAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Attributes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListFacetAttributes = "ListFacetAttributes"

// ListFacetAttributesRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Retrieves attributes attached to the facet.
//
//    // Example sending a request using ListFacetAttributesRequest.
//    req := client.ListFacetAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListFacetAttributes
func (c *Client) ListFacetAttributesRequest(input *ListFacetAttributesInput) ListFacetAttributesRequest {
	op := &aws.Operation{
		Name:       opListFacetAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/facet/attributes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFacetAttributesInput{}
	}

	req := c.newRequest(op, input, &ListFacetAttributesOutput{})
	return ListFacetAttributesRequest{Request: req, Input: input, Copy: c.ListFacetAttributesRequest}
}

// ListFacetAttributesRequest is the request type for the
// ListFacetAttributes API operation.
type ListFacetAttributesRequest struct {
	*aws.Request
	Input *ListFacetAttributesInput
	Copy  func(*ListFacetAttributesInput) ListFacetAttributesRequest
}

// Send marshals and sends the ListFacetAttributes API request.
func (r ListFacetAttributesRequest) Send(ctx context.Context) (*ListFacetAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFacetAttributesResponse{
		ListFacetAttributesOutput: r.Request.Data.(*ListFacetAttributesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFacetAttributesRequestPaginator returns a paginator for ListFacetAttributes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFacetAttributesRequest(input)
//   p := clouddirectory.NewListFacetAttributesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFacetAttributesPaginator(req ListFacetAttributesRequest) ListFacetAttributesPaginator {
	return ListFacetAttributesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFacetAttributesInput
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

// ListFacetAttributesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFacetAttributesPaginator struct {
	aws.Pager
}

func (p *ListFacetAttributesPaginator) CurrentPage() *ListFacetAttributesOutput {
	return p.Pager.CurrentPage().(*ListFacetAttributesOutput)
}

// ListFacetAttributesResponse is the response type for the
// ListFacetAttributes API operation.
type ListFacetAttributesResponse struct {
	*ListFacetAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFacetAttributes request.
func (r *ListFacetAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
