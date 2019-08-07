// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateCloudFormationChangeSetRequest
type CreateCloudFormationChangeSetInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"applicationId" type:"string" required:"true"`

	Capabilities []string `locationName:"capabilities" type:"list"`

	ChangeSetName *string `locationName:"changeSetName" type:"string"`

	ClientToken *string `locationName:"clientToken" type:"string"`

	Description *string `locationName:"description" type:"string"`

	NotificationArns []string `locationName:"notificationArns" type:"list"`

	ParameterOverrides []ParameterValue `locationName:"parameterOverrides" type:"list"`

	ResourceTypes []string `locationName:"resourceTypes" type:"list"`

	// This property corresponds to the AWS CloudFormation RollbackConfiguration
	// (https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/RollbackConfiguration)
	// Data Type.
	RollbackConfiguration *RollbackConfiguration `locationName:"rollbackConfiguration" type:"structure"`

	SemanticVersion *string `locationName:"semanticVersion" type:"string"`

	// StackName is a required field
	StackName *string `locationName:"stackName" type:"string" required:"true"`

	Tags []Tag `locationName:"tags" type:"list"`

	TemplateId *string `locationName:"templateId" type:"string"`
}

// String returns the string representation
func (s CreateCloudFormationChangeSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCloudFormationChangeSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCloudFormationChangeSetInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.ParameterOverrides != nil {
		for i, v := range s.ParameterOverrides {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ParameterOverrides", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RollbackConfiguration != nil {
		if err := s.RollbackConfiguration.Validate(); err != nil {
			invalidParams.AddNested("RollbackConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateCloudFormationChangeSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Capabilities != nil {
		v := s.Capabilities

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "capabilities", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ChangeSetName != nil {
		v := *s.ChangeSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "changeSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NotificationArns != nil {
		v := s.NotificationArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "notificationArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ParameterOverrides != nil {
		v := s.ParameterOverrides

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "parameterOverrides", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ResourceTypes != nil {
		v := s.ResourceTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.RollbackConfiguration != nil {
		v := s.RollbackConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "rollbackConfiguration", v, metadata)
	}
	if s.SemanticVersion != nil {
		v := *s.SemanticVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "semanticVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StackName != nil {
		v := *s.StackName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stackName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateCloudFormationChangeSetResponse
type CreateCloudFormationChangeSetOutput struct {
	_ struct{} `type:"structure"`

	ApplicationId *string `json:"serverlessrepo:CreateCloudFormationChangeSetOutput:ApplicationId" locationName:"applicationId" type:"string"`

	ChangeSetId *string `json:"serverlessrepo:CreateCloudFormationChangeSetOutput:ChangeSetId" locationName:"changeSetId" type:"string"`

	SemanticVersion *string `json:"serverlessrepo:CreateCloudFormationChangeSetOutput:SemanticVersion" locationName:"semanticVersion" type:"string"`

	StackId *string `json:"serverlessrepo:CreateCloudFormationChangeSetOutput:StackId" locationName:"stackId" type:"string"`
}

// String returns the string representation
func (s CreateCloudFormationChangeSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateCloudFormationChangeSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ChangeSetId != nil {
		v := *s.ChangeSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "changeSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SemanticVersion != nil {
		v := *s.SemanticVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "semanticVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StackId != nil {
		v := *s.StackId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stackId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateCloudFormationChangeSet = "CreateCloudFormationChangeSet"

// CreateCloudFormationChangeSetRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Creates an AWS CloudFormation change set for the given application.
//
//    // Example sending a request using CreateCloudFormationChangeSetRequest.
//    req := client.CreateCloudFormationChangeSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateCloudFormationChangeSet
func (c *Client) CreateCloudFormationChangeSetRequest(input *CreateCloudFormationChangeSetInput) CreateCloudFormationChangeSetRequest {
	op := &aws.Operation{
		Name:       opCreateCloudFormationChangeSet,
		HTTPMethod: "POST",
		HTTPPath:   "/applications/{applicationId}/changesets",
	}

	if input == nil {
		input = &CreateCloudFormationChangeSetInput{}
	}

	req := c.newRequest(op, input, &CreateCloudFormationChangeSetOutput{})
	return CreateCloudFormationChangeSetRequest{Request: req, Input: input, Copy: c.CreateCloudFormationChangeSetRequest}
}

// CreateCloudFormationChangeSetRequest is the request type for the
// CreateCloudFormationChangeSet API operation.
type CreateCloudFormationChangeSetRequest struct {
	*aws.Request
	Input *CreateCloudFormationChangeSetInput
	Copy  func(*CreateCloudFormationChangeSetInput) CreateCloudFormationChangeSetRequest
}

// Send marshals and sends the CreateCloudFormationChangeSet API request.
func (r CreateCloudFormationChangeSetRequest) Send(ctx context.Context) (*CreateCloudFormationChangeSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCloudFormationChangeSetResponse{
		CreateCloudFormationChangeSetOutput: r.Request.Data.(*CreateCloudFormationChangeSetOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCloudFormationChangeSetResponse is the response type for the
// CreateCloudFormationChangeSet API operation.
type CreateCloudFormationChangeSetResponse struct {
	*CreateCloudFormationChangeSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCloudFormationChangeSet request.
func (r *CreateCloudFormationChangeSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
