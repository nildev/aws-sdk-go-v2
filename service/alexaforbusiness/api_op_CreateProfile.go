// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateProfileRequest
type CreateProfileInput struct {
	_ struct{} `type:"structure"`

	// The valid address for the room.
	//
	// Address is a required field
	Address *string `min:"1" type:"string" required:"true"`

	// The user-specified token that is used during the creation of a profile.
	ClientRequestToken *string `min:"10" type:"string" idempotencyToken:"true"`

	// The distance unit to be used by devices in the profile.
	//
	// DistanceUnit is a required field
	DistanceUnit DistanceUnit `type:"string" required:"true" enum:"true"`

	// The maximum volume limit for a room profile.
	MaxVolumeLimit *int64 `type:"integer"`

	// Whether PSTN calling is enabled.
	PSTNEnabled *bool `type:"boolean"`

	// The name of a room profile.
	//
	// ProfileName is a required field
	ProfileName *string `min:"1" type:"string" required:"true"`

	// Whether room profile setup is enabled.
	SetupModeDisabled *bool `type:"boolean"`

	// The temperature unit to be used by devices in the profile.
	//
	// TemperatureUnit is a required field
	TemperatureUnit TemperatureUnit `type:"string" required:"true" enum:"true"`

	// The time zone used by a room profile.
	//
	// Timezone is a required field
	Timezone *string `min:"1" type:"string" required:"true"`

	// A wake word for Alexa, Echo, Amazon, or a computer.
	//
	// WakeWord is a required field
	WakeWord WakeWord `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProfileInput"}

	if s.Address == nil {
		invalidParams.Add(aws.NewErrParamRequired("Address"))
	}
	if s.Address != nil && len(*s.Address) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Address", 1))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 10))
	}
	if len(s.DistanceUnit) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DistanceUnit"))
	}

	if s.ProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfileName"))
	}
	if s.ProfileName != nil && len(*s.ProfileName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfileName", 1))
	}
	if len(s.TemperatureUnit) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TemperatureUnit"))
	}

	if s.Timezone == nil {
		invalidParams.Add(aws.NewErrParamRequired("Timezone"))
	}
	if s.Timezone != nil && len(*s.Timezone) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Timezone", 1))
	}
	if len(s.WakeWord) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("WakeWord"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateProfileResponse
type CreateProfileOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the newly created room profile in the response.
	ProfileArn *string `json:"a4b:CreateProfileOutput:ProfileArn" type:"string"`
}

// String returns the string representation
func (s CreateProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateProfile = "CreateProfile"

// CreateProfileRequest returns a request value for making API operation for
// Alexa For Business.
//
// Creates a new room profile with the specified details.
//
//    // Example sending a request using CreateProfileRequest.
//    req := client.CreateProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateProfile
func (c *Client) CreateProfileRequest(input *CreateProfileInput) CreateProfileRequest {
	op := &aws.Operation{
		Name:       opCreateProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateProfileInput{}
	}

	req := c.newRequest(op, input, &CreateProfileOutput{})
	return CreateProfileRequest{Request: req, Input: input, Copy: c.CreateProfileRequest}
}

// CreateProfileRequest is the request type for the
// CreateProfile API operation.
type CreateProfileRequest struct {
	*aws.Request
	Input *CreateProfileInput
	Copy  func(*CreateProfileInput) CreateProfileRequest
}

// Send marshals and sends the CreateProfile API request.
func (r CreateProfileRequest) Send(ctx context.Context) (*CreateProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProfileResponse{
		CreateProfileOutput: r.Request.Data.(*CreateProfileOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProfileResponse is the response type for the
// CreateProfile API operation.
type CreateProfileResponse struct {
	*CreateProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProfile request.
func (r *CreateProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
