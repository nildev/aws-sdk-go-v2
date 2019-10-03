// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateSecurityProfileInput struct {
	_ struct{} `type:"structure"`

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors but it is also retained for
	// any metric specified here.
	AdditionalMetricsToRetain []string `locationName:"additionalMetricsToRetain" type:"list"`

	// Specifies the destinations to which alerts are sent. (Alerts are always sent
	// to the console.) Alerts are generated when a device (thing) violates a behavior.
	AlertTargets map[string]AlertTarget `locationName:"alertTargets" type:"map"`

	// Specifies the behaviors that, when violated by a device (thing), cause an
	// alert.
	Behaviors []Behavior `locationName:"behaviors" type:"list"`

	// A description of the security profile.
	SecurityProfileDescription *string `locationName:"securityProfileDescription" type:"string"`

	// The name you are giving to the security profile.
	//
	// SecurityProfileName is a required field
	SecurityProfileName *string `location:"uri" locationName:"securityProfileName" min:"1" type:"string" required:"true"`

	// Metadata which can be used to manage the security profile.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateSecurityProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSecurityProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSecurityProfileInput"}

	if s.SecurityProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityProfileName"))
	}
	if s.SecurityProfileName != nil && len(*s.SecurityProfileName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityProfileName", 1))
	}
	if s.AlertTargets != nil {
		for i, v := range s.AlertTargets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AlertTargets", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Behaviors != nil {
		for i, v := range s.Behaviors {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Behaviors", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSecurityProfileInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AdditionalMetricsToRetain != nil {
		v := s.AdditionalMetricsToRetain

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "additionalMetricsToRetain", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.AlertTargets != nil {
		v := s.AlertTargets

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "alertTargets", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.Behaviors != nil {
		v := s.Behaviors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "behaviors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SecurityProfileDescription != nil {
		v := *s.SecurityProfileDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityProfileDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.SecurityProfileName != nil {
		v := *s.SecurityProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "securityProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateSecurityProfileOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the security profile.
	SecurityProfileArn *string `json:"iot:CreateSecurityProfileOutput:SecurityProfileArn" locationName:"securityProfileArn" type:"string"`

	// The name you gave to the security profile.
	SecurityProfileName *string `json:"iot:CreateSecurityProfileOutput:SecurityProfileName" locationName:"securityProfileName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateSecurityProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSecurityProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SecurityProfileArn != nil {
		v := *s.SecurityProfileArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityProfileArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecurityProfileName != nil {
		v := *s.SecurityProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateSecurityProfile = "CreateSecurityProfile"

// CreateSecurityProfileRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a Device Defender security profile.
//
//    // Example sending a request using CreateSecurityProfileRequest.
//    req := client.CreateSecurityProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateSecurityProfileRequest(input *CreateSecurityProfileInput) CreateSecurityProfileRequest {
	op := &aws.Operation{
		Name:       opCreateSecurityProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/security-profiles/{securityProfileName}",
	}

	if input == nil {
		input = &CreateSecurityProfileInput{}
	}

	req := c.newRequest(op, input, &CreateSecurityProfileOutput{})
	return CreateSecurityProfileRequest{Request: req, Input: input, Copy: c.CreateSecurityProfileRequest}
}

// CreateSecurityProfileRequest is the request type for the
// CreateSecurityProfile API operation.
type CreateSecurityProfileRequest struct {
	*aws.Request
	Input *CreateSecurityProfileInput
	Copy  func(*CreateSecurityProfileInput) CreateSecurityProfileRequest
}

// Send marshals and sends the CreateSecurityProfile API request.
func (r CreateSecurityProfileRequest) Send(ctx context.Context) (*CreateSecurityProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSecurityProfileResponse{
		CreateSecurityProfileOutput: r.Request.Data.(*CreateSecurityProfileOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSecurityProfileResponse is the response type for the
// CreateSecurityProfile API operation.
type CreateSecurityProfileResponse struct {
	*CreateSecurityProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSecurityProfile request.
func (r *CreateSecurityProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
