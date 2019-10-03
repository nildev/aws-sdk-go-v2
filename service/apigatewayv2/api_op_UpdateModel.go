// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateModelRequest
type UpdateModelInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// A string with a length between [1-256].
	ContentType *string `locationName:"contentType" type:"string"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	// ModelId is a required field
	ModelId *string `location:"uri" locationName:"modelId" type:"string" required:"true"`

	// A string with a length between [1-128].
	Name *string `locationName:"name" type:"string"`

	// A string with a length between [0-32768].
	Schema *string `locationName:"schema" type:"string"`
}

// String returns the string representation
func (s UpdateModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateModelInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.ModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateModelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contentType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Schema != nil {
		v := *s.Schema

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "schema", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ModelId != nil {
		v := *s.ModelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "modelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateModelResponse
type UpdateModelOutput struct {
	_ struct{} `type:"structure"`

	// A string with a length between [1-256].
	ContentType *string `json:"apigateway:UpdateModelOutput:ContentType" locationName:"contentType" type:"string"`

	// A string with a length between [0-1024].
	Description *string `json:"apigateway:UpdateModelOutput:Description" locationName:"description" type:"string"`

	// The identifier.
	ModelId *string `json:"apigateway:UpdateModelOutput:ModelId" locationName:"modelId" type:"string"`

	// A string with a length between [1-128].
	Name *string `json:"apigateway:UpdateModelOutput:Name" locationName:"name" type:"string"`

	// A string with a length between [0-32768].
	Schema *string `json:"apigateway:UpdateModelOutput:Schema" locationName:"schema" type:"string"`
}

// String returns the string representation
func (s UpdateModelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateModelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contentType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ModelId != nil {
		v := *s.ModelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "modelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Schema != nil {
		v := *s.Schema

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "schema", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateModel = "UpdateModel"

// UpdateModelRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Updates a Model.
//
//    // Example sending a request using UpdateModelRequest.
//    req := client.UpdateModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateModel
func (c *Client) UpdateModelRequest(input *UpdateModelInput) UpdateModelRequest {
	op := &aws.Operation{
		Name:       opUpdateModel,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v2/apis/{apiId}/models/{modelId}",
	}

	if input == nil {
		input = &UpdateModelInput{}
	}

	req := c.newRequest(op, input, &UpdateModelOutput{})
	return UpdateModelRequest{Request: req, Input: input, Copy: c.UpdateModelRequest}
}

// UpdateModelRequest is the request type for the
// UpdateModel API operation.
type UpdateModelRequest struct {
	*aws.Request
	Input *UpdateModelInput
	Copy  func(*UpdateModelInput) UpdateModelRequest
}

// Send marshals and sends the UpdateModel API request.
func (r UpdateModelRequest) Send(ctx context.Context) (*UpdateModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateModelResponse{
		UpdateModelOutput: r.Request.Data.(*UpdateModelOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateModelResponse is the response type for the
// UpdateModel API operation.
type UpdateModelResponse struct {
	*UpdateModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateModel request.
func (r *UpdateModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
