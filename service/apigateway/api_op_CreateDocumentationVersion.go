// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Creates a new documentation version of a given API.
type CreateDocumentationVersionInput struct {
	_ struct{} `type:"structure"`

	// A description about the new documentation snapshot.
	Description *string `locationName:"description" type:"string"`

	// [Required] The version identifier of the new snapshot.
	//
	// DocumentationVersion is a required field
	DocumentationVersion *string `locationName:"documentationVersion" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// The stage name to be associated with the new documentation snapshot.
	StageName *string `locationName:"stageName" type:"string"`
}

// String returns the string representation
func (s CreateDocumentationVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDocumentationVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDocumentationVersionInput"}

	if s.DocumentationVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentationVersion"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDocumentationVersionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DocumentationVersion != nil {
		v := *s.DocumentationVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "documentationVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageName != nil {
		v := *s.StageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stageName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A snapshot of the documentation of an API.
//
// Publishing API documentation involves creating a documentation version associated
// with an API stage and exporting the versioned documentation to an external
// (e.g., OpenAPI) file.
//
// Documenting an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationPart, DocumentationVersions
type CreateDocumentationVersionOutput struct {
	_ struct{} `type:"structure"`

	// The date when the API documentation snapshot is created.
	CreatedDate *time.Time `json:"apigateway:CreateDocumentationVersionOutput:CreatedDate" locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// The description of the API documentation snapshot.
	Description *string `json:"apigateway:CreateDocumentationVersionOutput:Description" locationName:"description" type:"string"`

	// The version identifier of the API documentation snapshot.
	Version *string `json:"apigateway:CreateDocumentationVersionOutput:Version" locationName:"version" type:"string"`
}

// String returns the string representation
func (s CreateDocumentationVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDocumentationVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDocumentationVersion = "CreateDocumentationVersion"

// CreateDocumentationVersionRequest returns a request value for making API operation for
// Amazon API Gateway.
//
//    // Example sending a request using CreateDocumentationVersionRequest.
//    req := client.CreateDocumentationVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateDocumentationVersionRequest(input *CreateDocumentationVersionInput) CreateDocumentationVersionRequest {
	op := &aws.Operation{
		Name:       opCreateDocumentationVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/restapis/{restapi_id}/documentation/versions",
	}

	if input == nil {
		input = &CreateDocumentationVersionInput{}
	}

	req := c.newRequest(op, input, &CreateDocumentationVersionOutput{})
	return CreateDocumentationVersionRequest{Request: req, Input: input, Copy: c.CreateDocumentationVersionRequest}
}

// CreateDocumentationVersionRequest is the request type for the
// CreateDocumentationVersion API operation.
type CreateDocumentationVersionRequest struct {
	*aws.Request
	Input *CreateDocumentationVersionInput
	Copy  func(*CreateDocumentationVersionInput) CreateDocumentationVersionRequest
}

// Send marshals and sends the CreateDocumentationVersion API request.
func (r CreateDocumentationVersionRequest) Send(ctx context.Context) (*CreateDocumentationVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDocumentationVersionResponse{
		CreateDocumentationVersionOutput: r.Request.Data.(*CreateDocumentationVersionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDocumentationVersionResponse is the response type for the
// CreateDocumentationVersion API operation.
type CreateDocumentationVersionResponse struct {
	*CreateDocumentationVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDocumentationVersion request.
func (r *CreateDocumentationVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
