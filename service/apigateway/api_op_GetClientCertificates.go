// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to get information about a collection of ClientCertificate resources.
type GetClientCertificatesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetClientCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetClientCertificatesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a collection of ClientCertificate resources.
//
// Use Client-Side Certificate (https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-client-side-ssl-authentication.html)
type GetClientCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []Certificate `json:"apigateway:GetClientCertificatesOutput:Items" locationName:"item" type:"list"`

	Position *string `json:"apigateway:GetClientCertificatesOutput:Position" locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetClientCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetClientCertificatesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "item", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetClientCertificates = "GetClientCertificates"

// GetClientCertificatesRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets a collection of ClientCertificate resources.
//
//    // Example sending a request using GetClientCertificatesRequest.
//    req := client.GetClientCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetClientCertificatesRequest(input *GetClientCertificatesInput) GetClientCertificatesRequest {
	op := &aws.Operation{
		Name:       opGetClientCertificates,
		HTTPMethod: "GET",
		HTTPPath:   "/clientcertificates",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"position"},
			OutputTokens:    []string{"position"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetClientCertificatesInput{}
	}

	req := c.newRequest(op, input, &GetClientCertificatesOutput{})
	return GetClientCertificatesRequest{Request: req, Input: input, Copy: c.GetClientCertificatesRequest}
}

// GetClientCertificatesRequest is the request type for the
// GetClientCertificates API operation.
type GetClientCertificatesRequest struct {
	*aws.Request
	Input *GetClientCertificatesInput
	Copy  func(*GetClientCertificatesInput) GetClientCertificatesRequest
}

// Send marshals and sends the GetClientCertificates API request.
func (r GetClientCertificatesRequest) Send(ctx context.Context) (*GetClientCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetClientCertificatesResponse{
		GetClientCertificatesOutput: r.Request.Data.(*GetClientCertificatesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetClientCertificatesRequestPaginator returns a paginator for GetClientCertificates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetClientCertificatesRequest(input)
//   p := apigateway.NewGetClientCertificatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetClientCertificatesPaginator(req GetClientCertificatesRequest) GetClientCertificatesPaginator {
	return GetClientCertificatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetClientCertificatesInput
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

// GetClientCertificatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetClientCertificatesPaginator struct {
	aws.Pager
}

func (p *GetClientCertificatesPaginator) CurrentPage() *GetClientCertificatesOutput {
	return p.Pager.CurrentPage().(*GetClientCertificatesOutput)
}

// GetClientCertificatesResponse is the response type for the
// GetClientCertificates API operation.
type GetClientCertificatesResponse struct {
	*GetClientCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetClientCertificates request.
func (r *GetClientCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
