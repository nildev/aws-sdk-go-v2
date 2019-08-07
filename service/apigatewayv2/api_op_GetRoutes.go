// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetRoutesRequest
type GetRoutesInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"maxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetRoutesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRoutesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRoutesInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRoutesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetRoutesResponse
type GetRoutesOutput struct {
	_ struct{} `type:"structure"`

	Items []Route `json:"apigateway:GetRoutesOutput:Items" locationName:"items" type:"list"`

	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string `json:"apigateway:GetRoutesOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetRoutesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRoutesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetRoutes = "GetRoutes"

// GetRoutesRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets the Routes for an API.
//
//    // Example sending a request using GetRoutesRequest.
//    req := client.GetRoutesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetRoutes
func (c *Client) GetRoutesRequest(input *GetRoutesInput) GetRoutesRequest {
	op := &aws.Operation{
		Name:       opGetRoutes,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/routes",
	}

	if input == nil {
		input = &GetRoutesInput{}
	}

	req := c.newRequest(op, input, &GetRoutesOutput{})
	return GetRoutesRequest{Request: req, Input: input, Copy: c.GetRoutesRequest}
}

// GetRoutesRequest is the request type for the
// GetRoutes API operation.
type GetRoutesRequest struct {
	*aws.Request
	Input *GetRoutesInput
	Copy  func(*GetRoutesInput) GetRoutesRequest
}

// Send marshals and sends the GetRoutes API request.
func (r GetRoutesRequest) Send(ctx context.Context) (*GetRoutesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRoutesResponse{
		GetRoutesOutput: r.Request.Data.(*GetRoutesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRoutesResponse is the response type for the
// GetRoutes API operation.
type GetRoutesResponse struct {
	*GetRoutesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRoutes request.
func (r *GetRoutesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
