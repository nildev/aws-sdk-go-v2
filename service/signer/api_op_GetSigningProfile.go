// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/GetSigningProfileRequest
type GetSigningProfileInput struct {
	_ struct{} `type:"structure"`

	// The name of the target signing profile.
	//
	// ProfileName is a required field
	ProfileName *string `location:"uri" locationName:"profileName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSigningProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSigningProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSigningProfileInput"}

	if s.ProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfileName"))
	}
	if s.ProfileName != nil && len(*s.ProfileName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfileName", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSigningProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ProfileName != nil {
		v := *s.ProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/GetSigningProfileResponse
type GetSigningProfileOutput struct {
	_ struct{} `type:"structure"`

	// A list of overrides applied by the target signing profile for signing operations.
	Overrides *SigningPlatformOverrides `json:"signer:GetSigningProfileOutput:Overrides" locationName:"overrides" type:"structure"`

	// The ID of the platform that is used by the target signing profile.
	PlatformId *string `json:"signer:GetSigningProfileOutput:PlatformId" locationName:"platformId" type:"string"`

	// The name of the target signing profile.
	ProfileName *string `json:"signer:GetSigningProfileOutput:ProfileName" locationName:"profileName" min:"2" type:"string"`

	// The ARN of the certificate that the target profile uses for signing operations.
	SigningMaterial *SigningMaterial `json:"signer:GetSigningProfileOutput:SigningMaterial" locationName:"signingMaterial" type:"structure"`

	// A map of key-value pairs for signing operations that is attached to the target
	// signing profile.
	SigningParameters map[string]string `json:"signer:GetSigningProfileOutput:SigningParameters" locationName:"signingParameters" type:"map"`

	// The status of the target signing profile.
	Status SigningProfileStatus `json:"signer:GetSigningProfileOutput:Status" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetSigningProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSigningProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Overrides != nil {
		v := s.Overrides

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "overrides", v, metadata)
	}
	if s.PlatformId != nil {
		v := *s.PlatformId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "platformId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProfileName != nil {
		v := *s.ProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "profileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SigningMaterial != nil {
		v := s.SigningMaterial

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "signingMaterial", v, metadata)
	}
	if s.SigningParameters != nil {
		v := s.SigningParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "signingParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opGetSigningProfile = "GetSigningProfile"

// GetSigningProfileRequest returns a request value for making API operation for
// AWS Signer.
//
// Returns information on a specific signing profile.
//
//    // Example sending a request using GetSigningProfileRequest.
//    req := client.GetSigningProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/GetSigningProfile
func (c *Client) GetSigningProfileRequest(input *GetSigningProfileInput) GetSigningProfileRequest {
	op := &aws.Operation{
		Name:       opGetSigningProfile,
		HTTPMethod: "GET",
		HTTPPath:   "/signing-profiles/{profileName}",
	}

	if input == nil {
		input = &GetSigningProfileInput{}
	}

	req := c.newRequest(op, input, &GetSigningProfileOutput{})
	return GetSigningProfileRequest{Request: req, Input: input, Copy: c.GetSigningProfileRequest}
}

// GetSigningProfileRequest is the request type for the
// GetSigningProfile API operation.
type GetSigningProfileRequest struct {
	*aws.Request
	Input *GetSigningProfileInput
	Copy  func(*GetSigningProfileInput) GetSigningProfileRequest
}

// Send marshals and sends the GetSigningProfile API request.
func (r GetSigningProfileRequest) Send(ctx context.Context) (*GetSigningProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSigningProfileResponse{
		GetSigningProfileOutput: r.Request.Data.(*GetSigningProfileOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSigningProfileResponse is the response type for the
// GetSigningProfile API operation.
type GetSigningProfileResponse struct {
	*GetSigningProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSigningProfile request.
func (r *GetSigningProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
