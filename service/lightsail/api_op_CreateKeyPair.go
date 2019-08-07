// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateKeyPairRequest
type CreateKeyPairInput struct {
	_ struct{} `type:"structure"`

	// The name for your new key pair.
	//
	// KeyPairName is a required field
	KeyPairName *string `locationName:"keyPairName" type:"string" required:"true"`

	// The tag keys and optional values to add to the resource during create.
	//
	// To tag a resource after it has been created, see the tag resource operation.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateKeyPairInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateKeyPairInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateKeyPairInput"}

	if s.KeyPairName == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyPairName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateKeyPairResult
type CreateKeyPairOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the new key pair
	// you just created.
	KeyPair *KeyPair `json:"lightsail:CreateKeyPairOutput:KeyPair" locationName:"keyPair" type:"structure"`

	// An array of key-value pairs containing information about the results of your
	// create key pair request.
	Operation *Operation `json:"lightsail:CreateKeyPairOutput:Operation" locationName:"operation" type:"structure"`

	// A base64-encoded RSA private key.
	PrivateKeyBase64 *string `json:"lightsail:CreateKeyPairOutput:PrivateKeyBase64" locationName:"privateKeyBase64" type:"string"`

	// A base64-encoded public key of the ssh-rsa type.
	PublicKeyBase64 *string `json:"lightsail:CreateKeyPairOutput:PublicKeyBase64" locationName:"publicKeyBase64" type:"string"`
}

// String returns the string representation
func (s CreateKeyPairOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateKeyPair = "CreateKeyPair"

// CreateKeyPairRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Creates an SSH key pair.
//
// The create key pair operation supports tag-based access control via request
// tags. For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using CreateKeyPairRequest.
//    req := client.CreateKeyPairRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateKeyPair
func (c *Client) CreateKeyPairRequest(input *CreateKeyPairInput) CreateKeyPairRequest {
	op := &aws.Operation{
		Name:       opCreateKeyPair,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateKeyPairInput{}
	}

	req := c.newRequest(op, input, &CreateKeyPairOutput{})
	return CreateKeyPairRequest{Request: req, Input: input, Copy: c.CreateKeyPairRequest}
}

// CreateKeyPairRequest is the request type for the
// CreateKeyPair API operation.
type CreateKeyPairRequest struct {
	*aws.Request
	Input *CreateKeyPairInput
	Copy  func(*CreateKeyPairInput) CreateKeyPairRequest
}

// Send marshals and sends the CreateKeyPair API request.
func (r CreateKeyPairRequest) Send(ctx context.Context) (*CreateKeyPairResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateKeyPairResponse{
		CreateKeyPairOutput: r.Request.Data.(*CreateKeyPairOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateKeyPairResponse is the response type for the
// CreateKeyPair API operation.
type CreateKeyPairResponse struct {
	*CreateKeyPairOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateKeyPair request.
func (r *CreateKeyPairResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
