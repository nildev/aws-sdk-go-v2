// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/UpdateGraphqlApiRequest
type UpdateGraphqlApiInput struct {
	_ struct{} `type:"structure"`

	// A list of additional authentication providers for the GraphqlApi API.
	AdditionalAuthenticationProviders []AdditionalAuthenticationProvider `locationName:"additionalAuthenticationProviders" type:"list"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The new authentication type for the GraphqlApi object.
	AuthenticationType AuthenticationType `locationName:"authenticationType" type:"string" enum:"true"`

	// The Amazon CloudWatch Logs configuration for the GraphqlApi object.
	LogConfig *LogConfig `locationName:"logConfig" type:"structure"`

	// The new name for the GraphqlApi object.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The OpenID Connect configuration for the GraphqlApi object.
	OpenIDConnectConfig *OpenIDConnectConfig `locationName:"openIDConnectConfig" type:"structure"`

	// The new Amazon Cognito user pool configuration for the GraphqlApi object.
	UserPoolConfig *UserPoolConfig `locationName:"userPoolConfig" type:"structure"`
}

// String returns the string representation
func (s UpdateGraphqlApiInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGraphqlApiInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGraphqlApiInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.AdditionalAuthenticationProviders != nil {
		for i, v := range s.AdditionalAuthenticationProviders {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AdditionalAuthenticationProviders", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.LogConfig != nil {
		if err := s.LogConfig.Validate(); err != nil {
			invalidParams.AddNested("LogConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.OpenIDConnectConfig != nil {
		if err := s.OpenIDConnectConfig.Validate(); err != nil {
			invalidParams.AddNested("OpenIDConnectConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.UserPoolConfig != nil {
		if err := s.UserPoolConfig.Validate(); err != nil {
			invalidParams.AddNested("UserPoolConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGraphqlApiInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AdditionalAuthenticationProviders != nil {
		v := s.AdditionalAuthenticationProviders

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "additionalAuthenticationProviders", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.AuthenticationType) > 0 {
		v := s.AuthenticationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authenticationType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.LogConfig != nil {
		v := s.LogConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "logConfig", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OpenIDConnectConfig != nil {
		v := s.OpenIDConnectConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "openIDConnectConfig", v, metadata)
	}
	if s.UserPoolConfig != nil {
		v := s.UserPoolConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "userPoolConfig", v, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/UpdateGraphqlApiResponse
type UpdateGraphqlApiOutput struct {
	_ struct{} `type:"structure"`

	// The updated GraphqlApi object.
	GraphqlApi *GraphqlApi `json:"appsync:UpdateGraphqlApiOutput:GraphqlApi" locationName:"graphqlApi" type:"structure"`
}

// String returns the string representation
func (s UpdateGraphqlApiOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGraphqlApiOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GraphqlApi != nil {
		v := s.GraphqlApi

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "graphqlApi", v, metadata)
	}
	return nil
}

const opUpdateGraphqlApi = "UpdateGraphqlApi"

// UpdateGraphqlApiRequest returns a request value for making API operation for
// AWS AppSync.
//
// Updates a GraphqlApi object.
//
//    // Example sending a request using UpdateGraphqlApiRequest.
//    req := client.UpdateGraphqlApiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/UpdateGraphqlApi
func (c *Client) UpdateGraphqlApiRequest(input *UpdateGraphqlApiInput) UpdateGraphqlApiRequest {
	op := &aws.Operation{
		Name:       opUpdateGraphqlApi,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}",
	}

	if input == nil {
		input = &UpdateGraphqlApiInput{}
	}

	req := c.newRequest(op, input, &UpdateGraphqlApiOutput{})
	return UpdateGraphqlApiRequest{Request: req, Input: input, Copy: c.UpdateGraphqlApiRequest}
}

// UpdateGraphqlApiRequest is the request type for the
// UpdateGraphqlApi API operation.
type UpdateGraphqlApiRequest struct {
	*aws.Request
	Input *UpdateGraphqlApiInput
	Copy  func(*UpdateGraphqlApiInput) UpdateGraphqlApiRequest
}

// Send marshals and sends the UpdateGraphqlApi API request.
func (r UpdateGraphqlApiRequest) Send(ctx context.Context) (*UpdateGraphqlApiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateGraphqlApiResponse{
		UpdateGraphqlApiOutput: r.Request.Data.(*UpdateGraphqlApiOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateGraphqlApiResponse is the response type for the
// UpdateGraphqlApi API operation.
type UpdateGraphqlApiResponse struct {
	*UpdateGraphqlApiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateGraphqlApi request.
func (r *UpdateGraphqlApiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
