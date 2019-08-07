// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/RemoveAttributesRequest
type RemoveAttributesInput struct {
	_ struct{} `type:"structure" payload:"UpdateAttributesRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// AttributeType is a required field
	AttributeType *string `location:"uri" locationName:"attribute-type" type:"string" required:"true"`

	// Specifies one or more attributes to remove from all the endpoints that are
	// associated with an application.
	//
	// UpdateAttributesRequest is a required field
	UpdateAttributesRequest *UpdateAttributesRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s RemoveAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveAttributesInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.AttributeType == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttributeType"))
	}

	if s.UpdateAttributesRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("UpdateAttributesRequest"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AttributeType != nil {
		v := *s.AttributeType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "attribute-type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpdateAttributesRequest != nil {
		v := s.UpdateAttributesRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "UpdateAttributesRequest", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/RemoveAttributesResponse
type RemoveAttributesOutput struct {
	_ struct{} `type:"structure" payload:"AttributesResource"`

	// Provides information about the type and the names of attributes that were
	// removed from all the endpoints that are associated with an application.
	//
	// AttributesResource is a required field
	AttributesResource *AttributesResource `json:"pinpoint:RemoveAttributesOutput:AttributesResource" type:"structure" required:"true"`
}

// String returns the string representation
func (s RemoveAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AttributesResource != nil {
		v := s.AttributesResource

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "AttributesResource", v, metadata)
	}
	return nil
}

const opRemoveAttributes = "RemoveAttributes"

// RemoveAttributesRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Removes one or more attributes, of the same attribute type, from all the
// endpoints that are associated with an application.
//
//    // Example sending a request using RemoveAttributesRequest.
//    req := client.RemoveAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/RemoveAttributes
func (c *Client) RemoveAttributesRequest(input *RemoveAttributesInput) RemoveAttributesRequest {
	op := &aws.Operation{
		Name:       opRemoveAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/attributes/{attribute-type}",
	}

	if input == nil {
		input = &RemoveAttributesInput{}
	}

	req := c.newRequest(op, input, &RemoveAttributesOutput{})
	return RemoveAttributesRequest{Request: req, Input: input, Copy: c.RemoveAttributesRequest}
}

// RemoveAttributesRequest is the request type for the
// RemoveAttributes API operation.
type RemoveAttributesRequest struct {
	*aws.Request
	Input *RemoveAttributesInput
	Copy  func(*RemoveAttributesInput) RemoveAttributesRequest
}

// Send marshals and sends the RemoveAttributes API request.
func (r RemoveAttributesRequest) Send(ctx context.Context) (*RemoveAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveAttributesResponse{
		RemoveAttributesOutput: r.Request.Data.(*RemoveAttributesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveAttributesResponse is the response type for the
// RemoveAttributes API operation.
type RemoveAttributesResponse struct {
	*RemoveAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveAttributes request.
func (r *RemoveAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
