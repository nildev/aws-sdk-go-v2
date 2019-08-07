// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateAddressBookRequest
type CreateAddressBookInput struct {
	_ struct{} `type:"structure"`

	// A unique, user-specified identifier for the request that ensures idempotency.
	ClientRequestToken *string `min:"10" type:"string" idempotencyToken:"true"`

	// The description of the address book.
	Description *string `min:"1" type:"string"`

	// The name of the address book.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateAddressBookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAddressBookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAddressBookInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 10))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateAddressBookResponse
type CreateAddressBookOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the newly created address book.
	AddressBookArn *string `json:"a4b:CreateAddressBookOutput:AddressBookArn" type:"string"`
}

// String returns the string representation
func (s CreateAddressBookOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateAddressBook = "CreateAddressBook"

// CreateAddressBookRequest returns a request value for making API operation for
// Alexa For Business.
//
// Creates an address book with the specified details.
//
//    // Example sending a request using CreateAddressBookRequest.
//    req := client.CreateAddressBookRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateAddressBook
func (c *Client) CreateAddressBookRequest(input *CreateAddressBookInput) CreateAddressBookRequest {
	op := &aws.Operation{
		Name:       opCreateAddressBook,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAddressBookInput{}
	}

	req := c.newRequest(op, input, &CreateAddressBookOutput{})
	return CreateAddressBookRequest{Request: req, Input: input, Copy: c.CreateAddressBookRequest}
}

// CreateAddressBookRequest is the request type for the
// CreateAddressBook API operation.
type CreateAddressBookRequest struct {
	*aws.Request
	Input *CreateAddressBookInput
	Copy  func(*CreateAddressBookInput) CreateAddressBookRequest
}

// Send marshals and sends the CreateAddressBook API request.
func (r CreateAddressBookRequest) Send(ctx context.Context) (*CreateAddressBookResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAddressBookResponse{
		CreateAddressBookOutput: r.Request.Data.(*CreateAddressBookOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAddressBookResponse is the response type for the
// CreateAddressBook API operation.
type CreateAddressBookResponse struct {
	*CreateAddressBookOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAddressBook request.
func (r *CreateAddressBookResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
