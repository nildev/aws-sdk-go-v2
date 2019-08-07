// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetContactRequest
type GetContactInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the contact for which to request details.
	//
	// ContactArn is a required field
	ContactArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetContactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetContactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetContactInput"}

	if s.ContactArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContactArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetContactResponse
type GetContactOutput struct {
	_ struct{} `type:"structure"`

	// The details of the requested contact.
	Contact *Contact `json:"a4b:GetContactOutput:Contact" type:"structure"`
}

// String returns the string representation
func (s GetContactOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetContact = "GetContact"

// GetContactRequest returns a request value for making API operation for
// Alexa For Business.
//
// Gets the contact details by the contact ARN.
//
//    // Example sending a request using GetContactRequest.
//    req := client.GetContactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetContact
func (c *Client) GetContactRequest(input *GetContactInput) GetContactRequest {
	op := &aws.Operation{
		Name:       opGetContact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetContactInput{}
	}

	req := c.newRequest(op, input, &GetContactOutput{})
	return GetContactRequest{Request: req, Input: input, Copy: c.GetContactRequest}
}

// GetContactRequest is the request type for the
// GetContact API operation.
type GetContactRequest struct {
	*aws.Request
	Input *GetContactInput
	Copy  func(*GetContactInput) GetContactRequest
}

// Send marshals and sends the GetContact API request.
func (r GetContactRequest) Send(ctx context.Context) (*GetContactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetContactResponse{
		GetContactOutput: r.Request.Data.(*GetContactOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetContactResponse is the response type for the
// GetContact API operation.
type GetContactResponse struct {
	*GetContactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetContact request.
func (r *GetContactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
