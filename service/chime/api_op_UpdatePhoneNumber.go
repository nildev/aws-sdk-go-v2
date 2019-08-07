// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdatePhoneNumberRequest
type UpdatePhoneNumberInput struct {
	_ struct{} `type:"structure"`

	// The phone number ID.
	//
	// PhoneNumberId is a required field
	PhoneNumberId *string `location:"uri" locationName:"phoneNumberId" type:"string" required:"true"`

	// The product type.
	ProductType PhoneNumberProductType `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdatePhoneNumberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePhoneNumberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePhoneNumberInput"}

	if s.PhoneNumberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PhoneNumberId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePhoneNumberInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.ProductType) > 0 {
		v := s.ProductType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ProductType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.PhoneNumberId != nil {
		v := *s.PhoneNumberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "phoneNumberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdatePhoneNumberResponse
type UpdatePhoneNumberOutput struct {
	_ struct{} `type:"structure"`

	// The updated phone number details.
	PhoneNumber *PhoneNumber `json:"chime:UpdatePhoneNumberOutput:PhoneNumber" type:"structure"`
}

// String returns the string representation
func (s UpdatePhoneNumberOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePhoneNumberOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PhoneNumber != nil {
		v := s.PhoneNumber

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PhoneNumber", v, metadata)
	}
	return nil
}

const opUpdatePhoneNumber = "UpdatePhoneNumber"

// UpdatePhoneNumberRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates phone number details, such as product type, for the specified phone
// number ID. For toll-free numbers, you can use only the Amazon Chime Voice
// Connector product type.
//
//    // Example sending a request using UpdatePhoneNumberRequest.
//    req := client.UpdatePhoneNumberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdatePhoneNumber
func (c *Client) UpdatePhoneNumberRequest(input *UpdatePhoneNumberInput) UpdatePhoneNumberRequest {
	op := &aws.Operation{
		Name:       opUpdatePhoneNumber,
		HTTPMethod: "POST",
		HTTPPath:   "/phone-numbers/{phoneNumberId}",
	}

	if input == nil {
		input = &UpdatePhoneNumberInput{}
	}

	req := c.newRequest(op, input, &UpdatePhoneNumberOutput{})
	return UpdatePhoneNumberRequest{Request: req, Input: input, Copy: c.UpdatePhoneNumberRequest}
}

// UpdatePhoneNumberRequest is the request type for the
// UpdatePhoneNumber API operation.
type UpdatePhoneNumberRequest struct {
	*aws.Request
	Input *UpdatePhoneNumberInput
	Copy  func(*UpdatePhoneNumberInput) UpdatePhoneNumberRequest
}

// Send marshals and sends the UpdatePhoneNumber API request.
func (r UpdatePhoneNumberRequest) Send(ctx context.Context) (*UpdatePhoneNumberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePhoneNumberResponse{
		UpdatePhoneNumberOutput: r.Request.Data.(*UpdatePhoneNumberOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePhoneNumberResponse is the response type for the
// UpdatePhoneNumber API operation.
type UpdatePhoneNumberResponse struct {
	*UpdatePhoneNumberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePhoneNumber request.
func (r *UpdatePhoneNumberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
