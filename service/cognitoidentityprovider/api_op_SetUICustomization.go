// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/SetUICustomizationRequest
type SetUICustomizationInput struct {
	_ struct{} `type:"structure"`

	// The CSS values in the UI customization.
	CSS *string `type:"string"`

	// The client ID for the client app.
	ClientId *string `min:"1" type:"string"`

	// The uploaded logo image for the UI customization.
	//
	// ImageFile is automatically base64 encoded/decoded by the SDK.
	ImageFile []byte `type:"blob"`

	// The user pool ID for the user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SetUICustomizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetUICustomizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetUICustomizationInput"}
	if s.ClientId != nil && len(*s.ClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientId", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/SetUICustomizationResponse
type SetUICustomizationOutput struct {
	_ struct{} `type:"structure"`

	// The UI customization information.
	//
	// UICustomization is a required field
	UICustomization *UICustomizationType `json:"cognito-idp:SetUICustomizationOutput:UICustomization" type:"structure" required:"true"`
}

// String returns the string representation
func (s SetUICustomizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetUICustomization = "SetUICustomization"

// SetUICustomizationRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Sets the UI customization information for a user pool's built-in app UI.
//
// You can specify app UI customization settings for a single client (with a
// specific clientId) or for all clients (by setting the clientId to ALL). If
// you specify ALL, the default configuration will be used for every client
// that has no UI customization set previously. If you specify UI customization
// settings for a particular client, it will no longer fall back to the ALL
// configuration.
//
// To use this API, your user pool must have a domain associated with it. Otherwise,
// there is no place to host the app's pages, and the service will throw an
// error.
//
//    // Example sending a request using SetUICustomizationRequest.
//    req := client.SetUICustomizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/SetUICustomization
func (c *Client) SetUICustomizationRequest(input *SetUICustomizationInput) SetUICustomizationRequest {
	op := &aws.Operation{
		Name:       opSetUICustomization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetUICustomizationInput{}
	}

	req := c.newRequest(op, input, &SetUICustomizationOutput{})
	return SetUICustomizationRequest{Request: req, Input: input, Copy: c.SetUICustomizationRequest}
}

// SetUICustomizationRequest is the request type for the
// SetUICustomization API operation.
type SetUICustomizationRequest struct {
	*aws.Request
	Input *SetUICustomizationInput
	Copy  func(*SetUICustomizationInput) SetUICustomizationRequest
}

// Send marshals and sends the SetUICustomization API request.
func (r SetUICustomizationRequest) Send(ctx context.Context) (*SetUICustomizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetUICustomizationResponse{
		SetUICustomizationOutput: r.Request.Data.(*SetUICustomizationOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetUICustomizationResponse is the response type for the
// SetUICustomization API operation.
type SetUICustomizationResponse struct {
	*SetUICustomizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetUICustomization request.
func (r *SetUICustomizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
