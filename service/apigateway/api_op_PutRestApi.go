// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A PUT request to update an existing API, with external API definitions specified
// as the request body.
type PutRestApiInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// [Required] The PUT request body containing external API definitions. Currently,
	// only OpenAPI definition JSON/YAML files are supported. The maximum size of
	// the API definition file is 2MB.
	//
	// Body is a required field
	Body []byte `locationName:"body" type:"blob" required:"true"`

	// A query parameter to indicate whether to rollback the API update (true) or
	// not (false) when a warning is encountered. The default value is false.
	FailOnWarnings *bool `location:"querystring" locationName:"failonwarnings" type:"boolean"`

	// The mode query parameter to specify the update mode. Valid values are "merge"
	// and "overwrite". By default, the update mode is "merge".
	Mode PutMode `location:"querystring" locationName:"mode" type:"string" enum:"true"`

	// Custom header parameters as part of the request. For example, to exclude
	// DocumentationParts from an imported API, set ignore=documentation as a parameters
	// value, as in the AWS CLI command of aws apigateway import-rest-api --parameters
	// ignore=documentation --body 'file:///path/to/imported-api-body.json'.
	Parameters map[string]string `location:"querystring" locationName:"parameters" type:"map"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s PutRestApiInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRestApiInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRestApiInput"}

	if s.Body == nil {
		invalidParams.Add(aws.NewErrParamRequired("Body"))
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
func (s PutRestApiInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "body", protocol.BytesStream(v), metadata)
	}
	if s.FailOnWarnings != nil {
		v := *s.FailOnWarnings

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "failonwarnings", protocol.BoolValue(v), metadata)
	}
	if len(s.Mode) > 0 {
		v := s.Mode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "mode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Parameters != nil {
		v := s.Parameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.QueryTarget, "parameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Represents a REST API.
//
// Create an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type PutRestApiOutput struct {
	_ struct{} `type:"structure"`

	// The source of the API key for metering requests according to a usage plan.
	// Valid values are:
	//    * HEADER to read the API key from the X-API-Key header of a request.
	//
	//    * AUTHORIZER to read the API key from the UsageIdentifierKey from a custom
	//    authorizer.
	ApiKeySource ApiKeySourceType `locationName:"apiKeySource" type:"string" enum:"true"`

	// The list of binary media types supported by the RestApi. By default, the
	// RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []string `locationName:"binaryMediaTypes" type:"list"`

	// The timestamp when the API was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// The API's description.
	Description *string `locationName:"description" type:"string"`

	// The endpoint configuration of this RestApi showing the endpoint types of
	// the API.
	EndpointConfiguration *EndpointConfiguration `locationName:"endpointConfiguration" type:"structure"`

	// The API's identifier. This identifier is unique across all of your APIs in
	// API Gateway.
	Id *string `locationName:"id" type:"string"`

	// A nullable integer that is used to enable compression (with non-negative
	// between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with
	// a null value) on an API. When compression is enabled, compression or decompression
	// is not applied on the payload if the payload size is smaller than this value.
	// Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *int64 `locationName:"minimumCompressionSize" type:"integer"`

	// The API's name.
	Name *string `locationName:"name" type:"string"`

	// A stringified JSON policy document that applies to this RestApi regardless
	// of the caller and Method configuration.
	Policy *string `locationName:"policy" type:"string"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// A version identifier for the API.
	Version *string `locationName:"version" type:"string"`

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []string `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s PutRestApiOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutRestApiOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ApiKeySource) > 0 {
		v := s.ApiKeySource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeySource", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.BinaryMediaTypes != nil {
		v := s.BinaryMediaTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "binaryMediaTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
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
	if s.EndpointConfiguration != nil {
		v := s.EndpointConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "endpointConfiguration", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MinimumCompressionSize != nil {
		v := *s.MinimumCompressionSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "minimumCompressionSize", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policy != nil {
		v := *s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Warnings != nil {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opPutRestApi = "PutRestApi"

// PutRestApiRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// A feature of the API Gateway control service for updating an existing API
// with an input of external API definitions. The update can take the form of
// merging the supplied definition into the existing API or overwriting the
// existing API.
//
//    // Example sending a request using PutRestApiRequest.
//    req := client.PutRestApiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) PutRestApiRequest(input *PutRestApiInput) PutRestApiRequest {
	op := &aws.Operation{
		Name:       opPutRestApi,
		HTTPMethod: "PUT",
		HTTPPath:   "/restapis/{restapi_id}",
	}

	if input == nil {
		input = &PutRestApiInput{}
	}

	req := c.newRequest(op, input, &PutRestApiOutput{})
	return PutRestApiRequest{Request: req, Input: input, Copy: c.PutRestApiRequest}
}

// PutRestApiRequest is the request type for the
// PutRestApi API operation.
type PutRestApiRequest struct {
	*aws.Request
	Input *PutRestApiInput
	Copy  func(*PutRestApiInput) PutRestApiRequest
}

// Send marshals and sends the PutRestApi API request.
func (r PutRestApiRequest) Send(ctx context.Context) (*PutRestApiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRestApiResponse{
		PutRestApiOutput: r.Request.Data.(*PutRestApiOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRestApiResponse is the response type for the
// PutRestApi API operation.
type PutRestApiResponse struct {
	*PutRestApiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRestApi request.
func (r *PutRestApiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
