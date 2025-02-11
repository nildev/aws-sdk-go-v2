// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListDevelopmentSchemaArnsRequest
type ListDevelopmentSchemaArnsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to retrieve.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDevelopmentSchemaArnsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDevelopmentSchemaArnsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDevelopmentSchemaArnsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDevelopmentSchemaArnsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListDevelopmentSchemaArnsResponse
type ListDevelopmentSchemaArnsOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The ARNs of retrieved development schemas.
	SchemaArns []string `type:"list"`
}

// String returns the string representation
func (s ListDevelopmentSchemaArnsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDevelopmentSchemaArnsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaArns != nil {
		v := s.SchemaArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SchemaArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opListDevelopmentSchemaArns = "ListDevelopmentSchemaArns"

// ListDevelopmentSchemaArnsRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Retrieves each Amazon Resource Name (ARN) of schemas in the development state.
//
//    // Example sending a request using ListDevelopmentSchemaArnsRequest.
//    req := client.ListDevelopmentSchemaArnsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListDevelopmentSchemaArns
func (c *Client) ListDevelopmentSchemaArnsRequest(input *ListDevelopmentSchemaArnsInput) ListDevelopmentSchemaArnsRequest {
	op := &aws.Operation{
		Name:       opListDevelopmentSchemaArns,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/schema/development",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDevelopmentSchemaArnsInput{}
	}

	req := c.newRequest(op, input, &ListDevelopmentSchemaArnsOutput{})
	return ListDevelopmentSchemaArnsRequest{Request: req, Input: input, Copy: c.ListDevelopmentSchemaArnsRequest}
}

// ListDevelopmentSchemaArnsRequest is the request type for the
// ListDevelopmentSchemaArns API operation.
type ListDevelopmentSchemaArnsRequest struct {
	*aws.Request
	Input *ListDevelopmentSchemaArnsInput
	Copy  func(*ListDevelopmentSchemaArnsInput) ListDevelopmentSchemaArnsRequest
}

// Send marshals and sends the ListDevelopmentSchemaArns API request.
func (r ListDevelopmentSchemaArnsRequest) Send(ctx context.Context) (*ListDevelopmentSchemaArnsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDevelopmentSchemaArnsResponse{
		ListDevelopmentSchemaArnsOutput: r.Request.Data.(*ListDevelopmentSchemaArnsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDevelopmentSchemaArnsRequestPaginator returns a paginator for ListDevelopmentSchemaArns.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDevelopmentSchemaArnsRequest(input)
//   p := clouddirectory.NewListDevelopmentSchemaArnsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDevelopmentSchemaArnsPaginator(req ListDevelopmentSchemaArnsRequest) ListDevelopmentSchemaArnsPaginator {
	return ListDevelopmentSchemaArnsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDevelopmentSchemaArnsInput
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

// ListDevelopmentSchemaArnsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDevelopmentSchemaArnsPaginator struct {
	aws.Pager
}

func (p *ListDevelopmentSchemaArnsPaginator) CurrentPage() *ListDevelopmentSchemaArnsOutput {
	return p.Pager.CurrentPage().(*ListDevelopmentSchemaArnsOutput)
}

// ListDevelopmentSchemaArnsResponse is the response type for the
// ListDevelopmentSchemaArns API operation.
type ListDevelopmentSchemaArnsResponse struct {
	*ListDevelopmentSchemaArnsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDevelopmentSchemaArns request.
func (r *ListDevelopmentSchemaArnsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
