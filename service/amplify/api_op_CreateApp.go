// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure used to create Apps in Amplify.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/CreateAppRequest
type CreateAppInput struct {
	_ struct{} `type:"structure"`

	// Personal Access token for 3rd party source control system for an Amplify
	// App, used to create webhook and read-only deploy key. Token is not stored.
	AccessToken *string `locationName:"accessToken" min:"1" type:"string"`

	// Automated branch creation config for the Amplify App.
	AutoBranchCreationConfig *AutoBranchCreationConfig `locationName:"autoBranchCreationConfig" type:"structure"`

	// Automated branch creation glob patterns for the Amplify App.
	AutoBranchCreationPatterns []string `locationName:"autoBranchCreationPatterns" type:"list"`

	// Credentials for Basic Authorization for an Amplify App.
	BasicAuthCredentials *string `locationName:"basicAuthCredentials" type:"string"`

	// BuildSpec for an Amplify App
	BuildSpec *string `locationName:"buildSpec" min:"1" type:"string"`

	// Custom rewrite / redirect rules for an Amplify App.
	CustomRules []CustomRule `locationName:"customRules" type:"list"`

	// Description for an Amplify App
	Description *string `locationName:"description" type:"string"`

	// Enables automated branch creation for the Amplify App.
	EnableAutoBranchCreation *bool `locationName:"enableAutoBranchCreation" type:"boolean"`

	// Enable Basic Authorization for an Amplify App, this will apply to all branches
	// part of this App.
	EnableBasicAuth *bool `locationName:"enableBasicAuth" type:"boolean"`

	// Enable the auto building of branches for an Amplify App.
	EnableBranchAutoBuild *bool `locationName:"enableBranchAutoBuild" type:"boolean"`

	// Environment variables map for an Amplify App.
	EnvironmentVariables map[string]string `locationName:"environmentVariables" type:"map"`

	// AWS IAM service role for an Amplify App
	IamServiceRoleArn *string `locationName:"iamServiceRoleArn" min:"1" type:"string"`

	// Name for the Amplify App
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// OAuth token for 3rd party source control system for an Amplify App, used
	// to create webhook and read-only deploy key. OAuth token is not stored.
	OauthToken *string `locationName:"oauthToken" type:"string"`

	// Platform / framework for an Amplify App
	Platform Platform `locationName:"platform" type:"string" enum:"true"`

	// Repository for an Amplify App
	Repository *string `locationName:"repository" type:"string"`

	// Tag for an Amplify App
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s CreateAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAppInput"}
	if s.AccessToken != nil && len(*s.AccessToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessToken", 1))
	}
	if s.BuildSpec != nil && len(*s.BuildSpec) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BuildSpec", 1))
	}
	if s.IamServiceRoleArn != nil && len(*s.IamServiceRoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IamServiceRoleArn", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.AutoBranchCreationConfig != nil {
		if err := s.AutoBranchCreationConfig.Validate(); err != nil {
			invalidParams.AddNested("AutoBranchCreationConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.CustomRules != nil {
		for i, v := range s.CustomRules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "CustomRules", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAppInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AccessToken != nil {
		v := *s.AccessToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "accessToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AutoBranchCreationConfig != nil {
		v := s.AutoBranchCreationConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "autoBranchCreationConfig", v, metadata)
	}
	if s.AutoBranchCreationPatterns != nil {
		v := s.AutoBranchCreationPatterns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "autoBranchCreationPatterns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.BasicAuthCredentials != nil {
		v := *s.BasicAuthCredentials

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "basicAuthCredentials", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BuildSpec != nil {
		v := *s.BuildSpec

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "buildSpec", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CustomRules != nil {
		v := s.CustomRules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "customRules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnableAutoBranchCreation != nil {
		v := *s.EnableAutoBranchCreation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableAutoBranchCreation", protocol.BoolValue(v), metadata)
	}
	if s.EnableBasicAuth != nil {
		v := *s.EnableBasicAuth

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableBasicAuth", protocol.BoolValue(v), metadata)
	}
	if s.EnableBranchAutoBuild != nil {
		v := *s.EnableBranchAutoBuild

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableBranchAutoBuild", protocol.BoolValue(v), metadata)
	}
	if s.EnvironmentVariables != nil {
		v := s.EnvironmentVariables

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "environmentVariables", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.IamServiceRoleArn != nil {
		v := *s.IamServiceRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "iamServiceRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OauthToken != nil {
		v := *s.OauthToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "oauthToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Platform) > 0 {
		v := s.Platform

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "platform", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Repository != nil {
		v := *s.Repository

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "repository", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/CreateAppResult
type CreateAppOutput struct {
	_ struct{} `type:"structure"`

	// Amplify App represents different branches of a repository for building, deploying,
	// and hosting.
	//
	// App is a required field
	App *App `json:"amplify:CreateAppOutput:App" locationName:"app" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateAppOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAppOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.App != nil {
		v := s.App

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "app", v, metadata)
	}
	return nil
}

const opCreateApp = "CreateApp"

// CreateAppRequest returns a request value for making API operation for
// AWS Amplify.
//
// Creates a new Amplify App.
//
//    // Example sending a request using CreateAppRequest.
//    req := client.CreateAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/CreateApp
func (c *Client) CreateAppRequest(input *CreateAppInput) CreateAppRequest {
	op := &aws.Operation{
		Name:       opCreateApp,
		HTTPMethod: "POST",
		HTTPPath:   "/apps",
	}

	if input == nil {
		input = &CreateAppInput{}
	}

	req := c.newRequest(op, input, &CreateAppOutput{})
	return CreateAppRequest{Request: req, Input: input, Copy: c.CreateAppRequest}
}

// CreateAppRequest is the request type for the
// CreateApp API operation.
type CreateAppRequest struct {
	*aws.Request
	Input *CreateAppInput
	Copy  func(*CreateAppInput) CreateAppRequest
}

// Send marshals and sends the CreateApp API request.
func (r CreateAppRequest) Send(ctx context.Context) (*CreateAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAppResponse{
		CreateAppOutput: r.Request.Data.(*CreateAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAppResponse is the response type for the
// CreateApp API operation.
type CreateAppResponse struct {
	*CreateAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApp request.
func (r *CreateAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
