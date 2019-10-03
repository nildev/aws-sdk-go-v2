// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetUserRequest
type GetUserInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The user ID.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUserInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetUserInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "userId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetUserResponse
type GetUserOutput struct {
	_ struct{} `type:"structure"`

	// The user details.
	User *User `json:"chime:GetUserOutput:User" type:"structure"`
}

// String returns the string representation
func (s GetUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.User != nil {
		v := s.User

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "User", v, metadata)
	}
	return nil
}

const opGetUser = "GetUser"

// GetUserRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves details for the specified user ID, such as primary email address,
// license type, and personal meeting PIN.
//
// To retrieve user details with an email address instead of a user ID, use
// the ListUsers action, and then filter by email address.
//
//    // Example sending a request using GetUserRequest.
//    req := client.GetUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetUser
func (c *Client) GetUserRequest(input *GetUserInput) GetUserRequest {
	op := &aws.Operation{
		Name:       opGetUser,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{accountId}/users/{userId}",
	}

	if input == nil {
		input = &GetUserInput{}
	}

	req := c.newRequest(op, input, &GetUserOutput{})
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
