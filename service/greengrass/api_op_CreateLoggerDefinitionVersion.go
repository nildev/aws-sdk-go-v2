// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateLoggerDefinitionVersionRequest
type CreateLoggerDefinitionVersionInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	// LoggerDefinitionId is a required field
	LoggerDefinitionId *string `location:"uri" locationName:"LoggerDefinitionId" type:"string" required:"true"`

	Loggers []Logger `type:"list"`
}

// String returns the string representation
func (s CreateLoggerDefinitionVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLoggerDefinitionVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLoggerDefinitionVersionInput"}

	if s.LoggerDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoggerDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateLoggerDefinitionVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Loggers != nil {
		v := s.Loggers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Loggers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AmznClientToken != nil {
		v := *s.AmznClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-Client-Token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LoggerDefinitionId != nil {
		v := *s.LoggerDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "LoggerDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateLoggerDefinitionVersionResponse
type CreateLoggerDefinitionVersionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `json:"greengrass:CreateLoggerDefinitionVersionOutput:Arn" type:"string"`

	CreationTimestamp *string `json:"greengrass:CreateLoggerDefinitionVersionOutput:CreationTimestamp" type:"string"`

	Id *string `json:"greengrass:CreateLoggerDefinitionVersionOutput:Id" type:"string"`

	Version *string `json:"greengrass:CreateLoggerDefinitionVersionOutput:Version" type:"string"`
}

// String returns the string representation
func (s CreateLoggerDefinitionVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateLoggerDefinitionVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTimestamp != nil {
		v := *s.CreationTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateLoggerDefinitionVersion = "CreateLoggerDefinitionVersion"

// CreateLoggerDefinitionVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a version of a logger definition that has already been defined.
//
//    // Example sending a request using CreateLoggerDefinitionVersionRequest.
//    req := client.CreateLoggerDefinitionVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateLoggerDefinitionVersion
func (c *Client) CreateLoggerDefinitionVersionRequest(input *CreateLoggerDefinitionVersionInput) CreateLoggerDefinitionVersionRequest {
	op := &aws.Operation{
		Name:       opCreateLoggerDefinitionVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/definition/loggers/{LoggerDefinitionId}/versions",
	}

	if input == nil {
		input = &CreateLoggerDefinitionVersionInput{}
	}

	req := c.newRequest(op, input, &CreateLoggerDefinitionVersionOutput{})
	return CreateLoggerDefinitionVersionRequest{Request: req, Input: input, Copy: c.CreateLoggerDefinitionVersionRequest}
}

// CreateLoggerDefinitionVersionRequest is the request type for the
// CreateLoggerDefinitionVersion API operation.
type CreateLoggerDefinitionVersionRequest struct {
	*aws.Request
	Input *CreateLoggerDefinitionVersionInput
	Copy  func(*CreateLoggerDefinitionVersionInput) CreateLoggerDefinitionVersionRequest
}

// Send marshals and sends the CreateLoggerDefinitionVersion API request.
func (r CreateLoggerDefinitionVersionRequest) Send(ctx context.Context) (*CreateLoggerDefinitionVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLoggerDefinitionVersionResponse{
		CreateLoggerDefinitionVersionOutput: r.Request.Data.(*CreateLoggerDefinitionVersionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLoggerDefinitionVersionResponse is the response type for the
// CreateLoggerDefinitionVersion API operation.
type CreateLoggerDefinitionVersionResponse struct {
	*CreateLoggerDefinitionVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLoggerDefinitionVersion request.
func (r *CreateLoggerDefinitionVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
