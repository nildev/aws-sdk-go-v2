// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the DescribeThingType operation.
type DescribeThingTypeInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing type.
	//
	// ThingTypeName is a required field
	ThingTypeName *string `location:"uri" locationName:"thingTypeName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeThingTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeThingTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeThingTypeInput"}

	if s.ThingTypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingTypeName"))
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
func (s DescribeThingTypeInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ThingTypeName != nil {
		v := *s.ThingTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output for the DescribeThingType operation.
type DescribeThingTypeOutput struct {
	_ struct{} `type:"structure"`

	// The thing type ARN.
	ThingTypeArn *string `json:"iot:DescribeThingTypeOutput:ThingTypeArn" locationName:"thingTypeArn" type:"string"`

	// The thing type ID.
	ThingTypeId *string `json:"iot:DescribeThingTypeOutput:ThingTypeId" locationName:"thingTypeId" type:"string"`

	// The ThingTypeMetadata contains additional information about the thing type
	// including: creation date and time, a value indicating whether the thing type
	// is deprecated, and a date and time when it was deprecated.
	ThingTypeMetadata *ThingTypeMetadata `json:"iot:DescribeThingTypeOutput:ThingTypeMetadata" locationName:"thingTypeMetadata" type:"structure"`

	// The name of the thing type.
	ThingTypeName *string `json:"iot:DescribeThingTypeOutput:ThingTypeName" locationName:"thingTypeName" min:"1" type:"string"`

	// The ThingTypeProperties contains information about the thing type including
	// description, and a list of searchable thing attribute names.
	ThingTypeProperties *ThingTypeProperties `json:"iot:DescribeThingTypeOutput:ThingTypeProperties" locationName:"thingTypeProperties" type:"structure"`
}

// String returns the string representation
func (s DescribeThingTypeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeThingTypeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ThingTypeArn != nil {
		v := *s.ThingTypeArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingTypeArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingTypeId != nil {
		v := *s.ThingTypeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingTypeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingTypeMetadata != nil {
		v := s.ThingTypeMetadata

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingTypeMetadata", v, metadata)
	}
	if s.ThingTypeName != nil {
		v := *s.ThingTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingTypeProperties != nil {
		v := s.ThingTypeProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingTypeProperties", v, metadata)
	}
	return nil
}

const opDescribeThingType = "DescribeThingType"

// DescribeThingTypeRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about the specified thing type.
//
//    // Example sending a request using DescribeThingTypeRequest.
//    req := client.DescribeThingTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeThingTypeRequest(input *DescribeThingTypeInput) DescribeThingTypeRequest {
	op := &aws.Operation{
		Name:       opDescribeThingType,
		HTTPMethod: "GET",
		HTTPPath:   "/thing-types/{thingTypeName}",
	}

	if input == nil {
		input = &DescribeThingTypeInput{}
	}

	req := c.newRequest(op, input, &DescribeThingTypeOutput{})
	return DescribeThingTypeRequest{Request: req, Input: input, Copy: c.DescribeThingTypeRequest}
}

// DescribeThingTypeRequest is the request type for the
// DescribeThingType API operation.
type DescribeThingTypeRequest struct {
	*aws.Request
	Input *DescribeThingTypeInput
	Copy  func(*DescribeThingTypeInput) DescribeThingTypeRequest
}

// Send marshals and sends the DescribeThingType API request.
func (r DescribeThingTypeRequest) Send(ctx context.Context) (*DescribeThingTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeThingTypeResponse{
		DescribeThingTypeOutput: r.Request.Data.(*DescribeThingTypeOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeThingTypeResponse is the response type for the
// DescribeThingType API operation.
type DescribeThingTypeResponse struct {
	*DescribeThingTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeThingType request.
func (r *DescribeThingTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
