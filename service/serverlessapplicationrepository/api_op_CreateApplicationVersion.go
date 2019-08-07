// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateApplicationVersionRequest
type CreateApplicationVersionInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"applicationId" type:"string" required:"true"`

	// SemanticVersion is a required field
	SemanticVersion *string `location:"uri" locationName:"semanticVersion" type:"string" required:"true"`

	SourceCodeArchiveUrl *string `locationName:"sourceCodeArchiveUrl" type:"string"`

	SourceCodeUrl *string `locationName:"sourceCodeUrl" type:"string"`

	TemplateBody *string `locationName:"templateBody" type:"string"`

	TemplateUrl *string `locationName:"templateUrl" type:"string"`
}

// String returns the string representation
func (s CreateApplicationVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateApplicationVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateApplicationVersionInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.SemanticVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("SemanticVersion"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateApplicationVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.SourceCodeArchiveUrl != nil {
		v := *s.SourceCodeArchiveUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceCodeArchiveUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceCodeUrl != nil {
		v := *s.SourceCodeUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceCodeUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateBody != nil {
		v := *s.TemplateBody

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateBody", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateUrl != nil {
		v := *s.TemplateUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SemanticVersion != nil {
		v := *s.SemanticVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "semanticVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateApplicationVersionResponse
type CreateApplicationVersionOutput struct {
	_ struct{} `type:"structure"`

	ApplicationId *string `json:"serverlessrepo:CreateApplicationVersionOutput:ApplicationId" locationName:"applicationId" type:"string"`

	CreationTime *string `json:"serverlessrepo:CreateApplicationVersionOutput:CreationTime" locationName:"creationTime" type:"string"`

	ParameterDefinitions []ParameterDefinition `json:"serverlessrepo:CreateApplicationVersionOutput:ParameterDefinitions" locationName:"parameterDefinitions" type:"list"`

	RequiredCapabilities []Capability `json:"serverlessrepo:CreateApplicationVersionOutput:RequiredCapabilities" locationName:"requiredCapabilities" type:"list"`

	ResourcesSupported *bool `json:"serverlessrepo:CreateApplicationVersionOutput:ResourcesSupported" locationName:"resourcesSupported" type:"boolean"`

	SemanticVersion *string `json:"serverlessrepo:CreateApplicationVersionOutput:SemanticVersion" locationName:"semanticVersion" type:"string"`

	SourceCodeArchiveUrl *string `json:"serverlessrepo:CreateApplicationVersionOutput:SourceCodeArchiveUrl" locationName:"sourceCodeArchiveUrl" type:"string"`

	SourceCodeUrl *string `json:"serverlessrepo:CreateApplicationVersionOutput:SourceCodeUrl" locationName:"sourceCodeUrl" type:"string"`

	TemplateUrl *string `json:"serverlessrepo:CreateApplicationVersionOutput:TemplateUrl" locationName:"templateUrl" type:"string"`
}

// String returns the string representation
func (s CreateApplicationVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateApplicationVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ParameterDefinitions != nil {
		v := s.ParameterDefinitions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "parameterDefinitions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RequiredCapabilities != nil {
		v := s.RequiredCapabilities

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "requiredCapabilities", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ResourcesSupported != nil {
		v := *s.ResourcesSupported

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourcesSupported", protocol.BoolValue(v), metadata)
	}
	if s.SemanticVersion != nil {
		v := *s.SemanticVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "semanticVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceCodeArchiveUrl != nil {
		v := *s.SourceCodeArchiveUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceCodeArchiveUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceCodeUrl != nil {
		v := *s.SourceCodeUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceCodeUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateUrl != nil {
		v := *s.TemplateUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateApplicationVersion = "CreateApplicationVersion"

// CreateApplicationVersionRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Creates an application version.
//
//    // Example sending a request using CreateApplicationVersionRequest.
//    req := client.CreateApplicationVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateApplicationVersion
func (c *Client) CreateApplicationVersionRequest(input *CreateApplicationVersionInput) CreateApplicationVersionRequest {
	op := &aws.Operation{
		Name:       opCreateApplicationVersion,
		HTTPMethod: "PUT",
		HTTPPath:   "/applications/{applicationId}/versions/{semanticVersion}",
	}

	if input == nil {
		input = &CreateApplicationVersionInput{}
	}

	req := c.newRequest(op, input, &CreateApplicationVersionOutput{})
	return CreateApplicationVersionRequest{Request: req, Input: input, Copy: c.CreateApplicationVersionRequest}
}

// CreateApplicationVersionRequest is the request type for the
// CreateApplicationVersion API operation.
type CreateApplicationVersionRequest struct {
	*aws.Request
	Input *CreateApplicationVersionInput
	Copy  func(*CreateApplicationVersionInput) CreateApplicationVersionRequest
}

// Send marshals and sends the CreateApplicationVersion API request.
func (r CreateApplicationVersionRequest) Send(ctx context.Context) (*CreateApplicationVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApplicationVersionResponse{
		CreateApplicationVersionOutput: r.Request.Data.(*CreateApplicationVersionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApplicationVersionResponse is the response type for the
// CreateApplicationVersion API operation.
type CreateApplicationVersionResponse struct {
	*CreateApplicationVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApplicationVersion request.
func (r *CreateApplicationVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
