// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DescribeConfigurationRequest
type DescribeConfigurationInput struct {
	_ struct{} `type:"structure"`

	// ConfigurationId is a required field
	ConfigurationId *string `location:"uri" locationName:"configuration-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConfigurationInput"}

	if s.ConfigurationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ConfigurationId != nil {
		v := *s.ConfigurationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "configuration-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DescribeConfigurationResponse
type DescribeConfigurationOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	Created *time.Time `locationName:"created" type:"timestamp" timestampFormat:"unix"`

	Description *string `locationName:"description" type:"string"`

	// The type of broker engine. Note: Currently, Amazon MQ supports only ActiveMQ.
	EngineType EngineType `locationName:"engineType" type:"string" enum:"true"`

	EngineVersion *string `locationName:"engineVersion" type:"string"`

	Id *string `locationName:"id" type:"string"`

	// Returns information about the specified configuration revision.
	LatestRevision *ConfigurationRevision `locationName:"latestRevision" type:"structure"`

	Name *string `locationName:"name" type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DescribeConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Created != nil {
		v := *s.Created

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "created", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.EngineType) > 0 {
		v := s.EngineType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "engineType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.EngineVersion != nil {
		v := *s.EngineVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "engineVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestRevision != nil {
		v := s.LatestRevision

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "latestRevision", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opDescribeConfiguration = "DescribeConfiguration"

// DescribeConfigurationRequest returns a request value for making API operation for
// AmazonMQ.
//
// Returns information about the specified configuration.
//
//    // Example sending a request using DescribeConfigurationRequest.
//    req := client.DescribeConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DescribeConfiguration
func (c *Client) DescribeConfigurationRequest(input *DescribeConfigurationInput) DescribeConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/configurations/{configuration-id}",
	}

	if input == nil {
		input = &DescribeConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeConfigurationOutput{})
	return DescribeConfigurationRequest{Request: req, Input: input, Copy: c.DescribeConfigurationRequest}
}

// DescribeConfigurationRequest is the request type for the
// DescribeConfiguration API operation.
type DescribeConfigurationRequest struct {
	*aws.Request
	Input *DescribeConfigurationInput
	Copy  func(*DescribeConfigurationInput) DescribeConfigurationRequest
}

// Send marshals and sends the DescribeConfiguration API request.
func (r DescribeConfigurationRequest) Send(ctx context.Context) (*DescribeConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConfigurationResponse{
		DescribeConfigurationOutput: r.Request.Data.(*DescribeConfigurationOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConfigurationResponse is the response type for the
// DescribeConfiguration API operation.
type DescribeConfigurationResponse struct {
	*DescribeConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConfiguration request.
func (r *DescribeConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
