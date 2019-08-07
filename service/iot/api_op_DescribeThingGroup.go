// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeThingGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing group.
	//
	// ThingGroupName is a required field
	ThingGroupName *string `location:"uri" locationName:"thingGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeThingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeThingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeThingGroupInput"}

	if s.ThingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingGroupName"))
	}
	if s.ThingGroupName != nil && len(*s.ThingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeThingGroupInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ThingGroupName != nil {
		v := *s.ThingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeThingGroupOutput struct {
	_ struct{} `type:"structure"`

	// The dynamic thing group index name.
	IndexName *string `json:"iot:DescribeThingGroupOutput:IndexName" locationName:"indexName" min:"1" type:"string"`

	// The dynamic thing group search query string.
	QueryString *string `json:"iot:DescribeThingGroupOutput:QueryString" locationName:"queryString" min:"1" type:"string"`

	// The dynamic thing group query version.
	QueryVersion *string `json:"iot:DescribeThingGroupOutput:QueryVersion" locationName:"queryVersion" type:"string"`

	// The dynamic thing group status.
	Status DynamicGroupStatus `json:"iot:DescribeThingGroupOutput:Status" locationName:"status" type:"string" enum:"true"`

	// The thing group ARN.
	ThingGroupArn *string `json:"iot:DescribeThingGroupOutput:ThingGroupArn" locationName:"thingGroupArn" type:"string"`

	// The thing group ID.
	ThingGroupId *string `json:"iot:DescribeThingGroupOutput:ThingGroupId" locationName:"thingGroupId" min:"1" type:"string"`

	// Thing group metadata.
	ThingGroupMetadata *ThingGroupMetadata `json:"iot:DescribeThingGroupOutput:ThingGroupMetadata" locationName:"thingGroupMetadata" type:"structure"`

	// The name of the thing group.
	ThingGroupName *string `json:"iot:DescribeThingGroupOutput:ThingGroupName" locationName:"thingGroupName" min:"1" type:"string"`

	// The thing group properties.
	ThingGroupProperties *ThingGroupProperties `json:"iot:DescribeThingGroupOutput:ThingGroupProperties" locationName:"thingGroupProperties" type:"structure"`

	// The version of the thing group.
	Version *int64 `json:"iot:DescribeThingGroupOutput:Version" locationName:"version" type:"long"`
}

// String returns the string representation
func (s DescribeThingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeThingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IndexName != nil {
		v := *s.IndexName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "indexName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.QueryString != nil {
		v := *s.QueryString

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queryString", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.QueryVersion != nil {
		v := *s.QueryVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "queryVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ThingGroupArn != nil {
		v := *s.ThingGroupArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingGroupArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingGroupId != nil {
		v := *s.ThingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingGroupMetadata != nil {
		v := s.ThingGroupMetadata

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingGroupMetadata", v, metadata)
	}
	if s.ThingGroupName != nil {
		v := *s.ThingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingGroupProperties != nil {
		v := s.ThingGroupProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingGroupProperties", v, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opDescribeThingGroup = "DescribeThingGroup"

// DescribeThingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Describe a thing group.
//
//    // Example sending a request using DescribeThingGroupRequest.
//    req := client.DescribeThingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeThingGroupRequest(input *DescribeThingGroupInput) DescribeThingGroupRequest {
	op := &aws.Operation{
		Name:       opDescribeThingGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/thing-groups/{thingGroupName}",
	}

	if input == nil {
		input = &DescribeThingGroupInput{}
	}

	req := c.newRequest(op, input, &DescribeThingGroupOutput{})
	return DescribeThingGroupRequest{Request: req, Input: input, Copy: c.DescribeThingGroupRequest}
}

// DescribeThingGroupRequest is the request type for the
// DescribeThingGroup API operation.
type DescribeThingGroupRequest struct {
	*aws.Request
	Input *DescribeThingGroupInput
	Copy  func(*DescribeThingGroupInput) DescribeThingGroupRequest
}

// Send marshals and sends the DescribeThingGroup API request.
func (r DescribeThingGroupRequest) Send(ctx context.Context) (*DescribeThingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeThingGroupResponse{
		DescribeThingGroupOutput: r.Request.Data.(*DescribeThingGroupOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeThingGroupResponse is the response type for the
// DescribeThingGroup API operation.
type DescribeThingGroupResponse struct {
	*DescribeThingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeThingGroup request.
func (r *DescribeThingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
