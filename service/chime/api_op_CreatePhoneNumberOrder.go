// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/CreatePhoneNumberOrderRequest
type CreatePhoneNumberOrderInput struct {
	_ struct{} `type:"structure"`

	// List of phone numbers, in E.164 format.
	//
	// E164PhoneNumbers is a required field
	E164PhoneNumbers []string `type:"list" required:"true"`

	// The phone number product type.
	//
	// ProductType is a required field
	ProductType PhoneNumberProductType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreatePhoneNumberOrderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePhoneNumberOrderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePhoneNumberOrderInput"}

	if s.E164PhoneNumbers == nil {
		invalidParams.Add(aws.NewErrParamRequired("E164PhoneNumbers"))
	}
	if len(s.ProductType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ProductType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePhoneNumberOrderInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.E164PhoneNumbers != nil {
		v := s.E164PhoneNumbers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "E164PhoneNumbers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.ProductType) > 0 {
		v := s.ProductType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ProductType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/CreatePhoneNumberOrderResponse
type CreatePhoneNumberOrderOutput struct {
	_ struct{} `type:"structure"`

	// The phone number order details.
	PhoneNumberOrder *PhoneNumberOrder `type:"structure"`
}

// String returns the string representation
func (s CreatePhoneNumberOrderOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePhoneNumberOrderOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PhoneNumberOrder != nil {
		v := s.PhoneNumberOrder

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PhoneNumberOrder", v, metadata)
	}
	return nil
}

const opCreatePhoneNumberOrder = "CreatePhoneNumberOrder"

// CreatePhoneNumberOrderRequest returns a request value for making API operation for
// Amazon Chime.
//
// Creates an order for phone numbers to be provisioned. Choose from Amazon
// Chime Business Calling and Amazon Chime Voice Connector product types. For
// toll-free numbers, you can use only the Amazon Chime Voice Connector product
// type.
//
//    // Example sending a request using CreatePhoneNumberOrderRequest.
//    req := client.CreatePhoneNumberOrderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/CreatePhoneNumberOrder
func (c *Client) CreatePhoneNumberOrderRequest(input *CreatePhoneNumberOrderInput) CreatePhoneNumberOrderRequest {
	op := &aws.Operation{
		Name:       opCreatePhoneNumberOrder,
		HTTPMethod: "POST",
		HTTPPath:   "/phone-number-orders",
	}

	if input == nil {
		input = &CreatePhoneNumberOrderInput{}
	}

	req := c.newRequest(op, input, &CreatePhoneNumberOrderOutput{})
	return CreatePhoneNumberOrderRequest{Request: req, Input: input, Copy: c.CreatePhoneNumberOrderRequest}
}

// CreatePhoneNumberOrderRequest is the request type for the
// CreatePhoneNumberOrder API operation.
type CreatePhoneNumberOrderRequest struct {
	*aws.Request
	Input *CreatePhoneNumberOrderInput
	Copy  func(*CreatePhoneNumberOrderInput) CreatePhoneNumberOrderRequest
}

// Send marshals and sends the CreatePhoneNumberOrder API request.
func (r CreatePhoneNumberOrderRequest) Send(ctx context.Context) (*CreatePhoneNumberOrderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePhoneNumberOrderResponse{
		CreatePhoneNumberOrderOutput: r.Request.Data.(*CreatePhoneNumberOrderOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePhoneNumberOrderResponse is the response type for the
// CreatePhoneNumberOrder API operation.
type CreatePhoneNumberOrderResponse struct {
	*CreatePhoneNumberOrderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePhoneNumberOrder request.
func (r *CreatePhoneNumberOrderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
