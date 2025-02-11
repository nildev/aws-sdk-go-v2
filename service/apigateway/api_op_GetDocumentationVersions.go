// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Gets the documentation versions of an API.
type GetDocumentationVersionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDocumentationVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDocumentationVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDocumentationVersionsInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDocumentationVersionsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
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

// The collection of documentation snapshots of an API.
//
// Use the DocumentationVersions to manage documentation snapshots associated
// with various API stages.
//
// Documenting an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationPart, DocumentationVersion
type GetDocumentationVersionsOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []DocumentationVersion `locationName:"item" type:"list"`

	Position *string `locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetDocumentationVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDocumentationVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opGetDocumentationVersions = "GetDocumentationVersions"

// GetDocumentationVersionsRequest returns a request value for making API operation for
// Amazon API Gateway.
//
//    // Example sending a request using GetDocumentationVersionsRequest.
//    req := client.GetDocumentationVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetDocumentationVersionsRequest(input *GetDocumentationVersionsInput) GetDocumentationVersionsRequest {
	op := &aws.Operation{
		Name:       opGetDocumentationVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/documentation/versions",
	}

	if input == nil {
		input = &GetDocumentationVersionsInput{}
	}

	req := c.newRequest(op, input, &GetDocumentationVersionsOutput{})
	return GetDocumentationVersionsRequest{Request: req, Input: input, Copy: c.GetDocumentationVersionsRequest}
}

// GetDocumentationVersionsRequest is the request type for the
// GetDocumentationVersions API operation.
type GetDocumentationVersionsRequest struct {
	*aws.Request
	Input *GetDocumentationVersionsInput
	Copy  func(*GetDocumentationVersionsInput) GetDocumentationVersionsRequest
}

// Send marshals and sends the GetDocumentationVersions API request.
func (r GetDocumentationVersionsRequest) Send(ctx context.Context) (*GetDocumentationVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDocumentationVersionsResponse{
		GetDocumentationVersionsOutput: r.Request.Data.(*GetDocumentationVersionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDocumentationVersionsResponse is the response type for the
// GetDocumentationVersions API operation.
type GetDocumentationVersionsResponse struct {
	*GetDocumentationVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDocumentationVersions request.
func (r *GetDocumentationVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
