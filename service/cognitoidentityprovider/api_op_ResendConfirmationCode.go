// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to resend the confirmation code.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ResendConfirmationCodeRequest
type ResendConfirmationCodeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Pinpoint analytics metadata for collecting metrics for ResendConfirmationCode
	// calls.
	AnalyticsMetadata *AnalyticsMetadataType `type:"structure"`

	// The ID of the client associated with the user pool.
	//
	// ClientId is a required field
	ClientId *string `min:"1" type:"string" required:"true"`

	// A keyed-hash message authentication code (HMAC) calculated using the secret
	// key of a user pool client and username plus the client ID in the message.
	SecretHash *string `min:"1" type:"string"`

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	UserContextData *UserContextDataType `type:"structure"`

	// The user name of the user to whom you wish to resend a confirmation code.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ResendConfirmationCodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResendConfirmationCodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResendConfirmationCodeInput"}

	if s.ClientId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientId"))
	}
	if s.ClientId != nil && len(*s.ClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientId", 1))
	}
	if s.SecretHash != nil && len(*s.SecretHash) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretHash", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response from the server when the Amazon Cognito Your User Pools service
// makes the request to resend a confirmation code.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ResendConfirmationCodeResponse
type ResendConfirmationCodeOutput struct {
	_ struct{} `type:"structure"`

	// The code delivery details returned by the server in response to the request
	// to resend the confirmation code.
	CodeDeliveryDetails *CodeDeliveryDetailsType `json:"cognito-idp:ResendConfirmationCodeOutput:CodeDeliveryDetails" type:"structure"`
}

// String returns the string representation
func (s ResendConfirmationCodeOutput) String() string {
	return awsutil.Prettify(s)
}

const opResendConfirmationCode = "ResendConfirmationCode"

// ResendConfirmationCodeRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Resends the confirmation (for confirmation of registration) to a specific
// user in the user pool.
//
//    // Example sending a request using ResendConfirmationCodeRequest.
//    req := client.ResendConfirmationCodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ResendConfirmationCode
func (c *Client) ResendConfirmationCodeRequest(input *ResendConfirmationCodeInput) ResendConfirmationCodeRequest {
	op := &aws.Operation{
		Name:       opResendConfirmationCode,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResendConfirmationCodeInput{}
	}

	req := c.newRequest(op, input, &ResendConfirmationCodeOutput{})
	req.Config.Credentials = aws.AnonymousCredentials
	return ResendConfirmationCodeRequest{Request: req, Input: input, Copy: c.ResendConfirmationCodeRequest}
}

// ResendConfirmationCodeRequest is the request type for the
// ResendConfirmationCode API operation.
type ResendConfirmationCodeRequest struct {
	*aws.Request
	Input *ResendConfirmationCodeInput
	Copy  func(*ResendConfirmationCodeInput) ResendConfirmationCodeRequest
}

// Send marshals and sends the ResendConfirmationCode API request.
func (r ResendConfirmationCodeRequest) Send(ctx context.Context) (*ResendConfirmationCodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResendConfirmationCodeResponse{
		ResendConfirmationCodeOutput: r.Request.Data.(*ResendConfirmationCodeOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResendConfirmationCodeResponse is the response type for the
// ResendConfirmationCode API operation.
type ResendConfirmationCodeResponse struct {
	*ResendConfirmationCodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResendConfirmationCode request.
func (r *ResendConfirmationCodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
