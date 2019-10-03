// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to get information about the user.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetUserRequest
type GetUserInput struct {
	_ struct{} `type:"structure"`

	// The access token returned by the server response to get information about
	// the user.
	//
	// AccessToken is a required field
	AccessToken *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUserInput"}

	if s.AccessToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessToken"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server from the request to get information
// about the user.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetUserResponse
type GetUserOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the options for MFA (e.g., email or phone number).
	MFAOptions []MFAOptionType `json:"cognito-idp:GetUserOutput:MFAOptions" type:"list"`

	// The user's preferred MFA setting.
	PreferredMfaSetting *string `json:"cognito-idp:GetUserOutput:PreferredMfaSetting" type:"string"`

	// An array of name-value pairs representing user attributes.
	//
	// For custom attributes, you must prepend the custom: prefix to the attribute
	// name.
	//
	// UserAttributes is a required field
	UserAttributes []AttributeType `json:"cognito-idp:GetUserOutput:UserAttributes" type:"list" required:"true"`

	// The list of the user's MFA settings.
	UserMFASettingList []string `json:"cognito-idp:GetUserOutput:UserMFASettingList" type:"list"`

	// The user name of the user you wish to retrieve from the get user request.
	//
	// Username is a required field
	Username *string `json:"cognito-idp:GetUserOutput:Username" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetUser = "GetUser"

// GetUserRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Gets the user attributes and metadata for a user.
//
//    // Example sending a request using GetUserRequest.
//    req := client.GetUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetUser
func (c *Client) GetUserRequest(input *GetUserInput) GetUserRequest {
	op := &aws.Operation{
		Name:       opGetUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetUserInput{}
	}

	req := c.newRequest(op, input, &GetUserOutput{})
	req.Config.Credentials = aws.AnonymousCredentials
	return GetUserRequest{Request: req, Input: input, Copy: c.GetUserRequest}
}

// GetUserRequest is the request type for the
// GetUser API operation.
type GetUserRequest struct {
	*aws.Request
	Input *GetUserInput
	Copy  func(*GetUserInput) GetUserRequest
}

// Send marshals and sends the GetUser API request.
func (r GetUserRequest) Send(ctx context.Context) (*GetUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUserResponse{
		GetUserOutput: r.Request.Data.(*GetUserOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUserResponse is the response type for the
// GetUser API operation.
type GetUserResponse struct {
	*GetUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUser request.
func (r *GetUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
