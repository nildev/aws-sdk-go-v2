// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/GetCloudFormationTemplateRequest
type GetCloudFormationTemplateInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"applicationId" type:"string" required:"true"`

	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"templateId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCloudFormationTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCloudFormationTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCloudFormationTemplateInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.TemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCloudFormationTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "templateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/GetCloudFormationTemplateResponse
type GetCloudFormationTemplateOutput struct {
	_ struct{} `type:"structure"`

	ApplicationId *string `json:"serverlessrepo:GetCloudFormationTemplateOutput:ApplicationId" locationName:"applicationId" type:"string"`

	CreationTime *string `json:"serverlessrepo:GetCloudFormationTemplateOutput:CreationTime" locationName:"creationTime" type:"string"`

	ExpirationTime *string `json:"serverlessrepo:GetCloudFormationTemplateOutput:ExpirationTime" locationName:"expirationTime" type:"string"`

	SemanticVersion *string `json:"serverlessrepo:GetCloudFormationTemplateOutput:SemanticVersion" locationName:"semanticVersion" type:"string"`

	Status Status `json:"serverlessrepo:GetCloudFormationTemplateOutput:Status" locationName:"status" type:"string" enum:"true"`

	TemplateId *string `json:"serverlessrepo:GetCloudFormationTemplateOutput:TemplateId" locationName:"templateId" type:"string"`

	TemplateUrl *string `json:"serverlessrepo:GetCloudFormationTemplateOutput:TemplateUrl" locationName:"templateUrl" type:"string"`
}

// String returns the string representation
func (s GetCloudFormationTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCloudFormationTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.ExpirationTime != nil {
		v := *s.ExpirationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expirationTime", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SemanticVersion != nil {
		v := *s.SemanticVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "semanticVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateUrl != nil {
		v := *s.TemplateUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetCloudFormationTemplate = "GetCloudFormationTemplate"

// GetCloudFormationTemplateRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Gets the specified AWS CloudFormation template.
//
//    // Example sending a request using GetCloudFormationTemplateRequest.
//    req := client.GetCloudFormationTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/GetCloudFormationTemplate
func (c *Client) GetCloudFormationTemplateRequest(input *GetCloudFormationTemplateInput) GetCloudFormationTemplateRequest {
	op := &aws.Operation{
		Name:       opGetCloudFormationTemplate,
		HTTPMethod: "GET",
		HTTPPath:   "/applications/{applicationId}/templates/{templateId}",
	}

	if input == nil {
		input = &GetCloudFormationTemplateInput{}
	}

	req := c.newRequest(op, input, &GetCloudFormationTemplateOutput{})
	return GetCloudFormationTemplateRequest{Request: req, Input: input, Copy: c.GetCloudFormationTemplateRequest}
}

// GetCloudFormationTemplateRequest is the request type for the
// GetCloudFormationTemplate API operation.
type GetCloudFormationTemplateRequest struct {
	*aws.Request
	Input *GetCloudFormationTemplateInput
	Copy  func(*GetCloudFormationTemplateInput) GetCloudFormationTemplateRequest
}

// Send marshals and sends the GetCloudFormationTemplate API request.
func (r GetCloudFormationTemplateRequest) Send(ctx context.Context) (*GetCloudFormationTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCloudFormationTemplateResponse{
		GetCloudFormationTemplateOutput: r.Request.Data.(*GetCloudFormationTemplateOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCloudFormationTemplateResponse is the response type for the
// GetCloudFormationTemplate API operation.
type GetCloudFormationTemplateResponse struct {
	*GetCloudFormationTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCloudFormationTemplate request.
func (r *GetCloudFormationTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
