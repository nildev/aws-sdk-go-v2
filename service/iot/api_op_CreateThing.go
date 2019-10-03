// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the CreateThing operation.
type CreateThingInput struct {
	_ struct{} `type:"structure"`

	// The attribute payload, which consists of up to three name/value pairs in
	// a JSON document. For example:
	//
	// {\"attributes\":{\"string1\":\"string2\"}}
	AttributePayload *AttributePayload `locationName:"attributePayload" type:"structure"`

	// The name of the billing group the thing will be added to.
	BillingGroupName *string `locationName:"billingGroupName" min:"1" type:"string"`

	// The name of the thing to create.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`

	// The name of the thing type associated with the new thing.
	ThingTypeName *string `locationName:"thingTypeName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateThingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateThingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateThingInput"}
	if s.BillingGroupName != nil && len(*s.BillingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BillingGroupName", 1))
	}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}
	if s.ThingTypeName != nil && len(*s.ThingTypeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingTypeName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateThingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AttributePayload != nil {
		v := s.AttributePayload

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "attributePayload", v, metadata)
	}
	if s.BillingGroupName != nil {
		v := *s.BillingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "billingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingTypeName != nil {
		v := *s.ThingTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output of the CreateThing operation.
type CreateThingOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the new thing.
	ThingArn *string `json:"iot:CreateThingOutput:ThingArn" locationName:"thingArn" type:"string"`

	// The thing ID.
	ThingId *string `json:"iot:CreateThingOutput:ThingId" locationName:"thingId" type:"string"`

	// The name of the new thing.
	ThingName *string `json:"iot:CreateThingOutput:ThingName" locationName:"thingName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateThingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateThingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ThingArn != nil {
		v := *s.ThingArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingId != nil {
		v := *s.ThingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateThing = "CreateThing"

// CreateThingRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a thing record in the registry. If this call is made multiple times
// using the same thing name and configuration, the call will succeed. If this
// call is made with the same thing name but different configuration a ResourceAlreadyExistsException
// is thrown.
//
// This is a control plane operation. See Authorization (https://docs.aws.amazon.com/iot/latest/developerguide/authorization.html)
// for information about authorizing control plane actions.
//
//    // Example sending a request using CreateThingRequest.
//    req := client.CreateThingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateThingRequest(input *CreateThingInput) CreateThingRequest {
	op := &aws.Operation{
		Name:       opCreateThing,
		HTTPMethod: "POST",
		HTTPPath:   "/things/{thingName}",
	}

	if input == nil {
		input = &CreateThingInput{}
	}

	req := c.newRequest(op, input, &CreateThingOutput{})
	return CreateThingRequest{Request: req, Input: input, Copy: c.CreateThingRequest}
}

// CreateThingRequest is the request type for the
// CreateThing API operation.
type CreateThingRequest struct {
	*aws.Request
	Input *CreateThingInput
	Copy  func(*CreateThingInput) CreateThingRequest
}

// Send marshals and sends the CreateThing API request.
func (r CreateThingRequest) Send(ctx context.Context) (*CreateThingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateThingResponse{
		CreateThingOutput: r.Request.Data.(*CreateThingOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateThingResponse is the response type for the
// CreateThing API operation.
type CreateThingResponse struct {
	*CreateThingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateThing request.
func (r *CreateThingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
