// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateSubscriptionDefinitionRequest
type CreateSubscriptionDefinitionInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	// Information about a subscription definition version.
	InitialVersion *SubscriptionDefinitionVersion `type:"structure"`

	Name *string `type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateSubscriptionDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSubscriptionDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.InitialVersion != nil {
		v := s.InitialVersion

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "InitialVersion", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.AmznClientToken != nil {
		v := *s.AmznClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-Client-Token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateSubscriptionDefinitionResponse
type CreateSubscriptionDefinitionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `json:"greengrass:CreateSubscriptionDefinitionOutput:Arn" type:"string"`

	CreationTimestamp *string `json:"greengrass:CreateSubscriptionDefinitionOutput:CreationTimestamp" type:"string"`

	Id *string `json:"greengrass:CreateSubscriptionDefinitionOutput:Id" type:"string"`

	LastUpdatedTimestamp *string `json:"greengrass:CreateSubscriptionDefinitionOutput:LastUpdatedTimestamp" type:"string"`

	LatestVersion *string `json:"greengrass:CreateSubscriptionDefinitionOutput:LatestVersion" type:"string"`

	LatestVersionArn *string `json:"greengrass:CreateSubscriptionDefinitionOutput:LatestVersionArn" type:"string"`

	Name *string `json:"greengrass:CreateSubscriptionDefinitionOutput:Name" type:"string"`
}

// String returns the string representation
func (s CreateSubscriptionDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSubscriptionDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTimestamp != nil {
		v := *s.CreationTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedTimestamp != nil {
		v := *s.LastUpdatedTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastUpdatedTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestVersion != nil {
		v := *s.LatestVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LatestVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestVersionArn != nil {
		v := *s.LatestVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LatestVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateSubscriptionDefinition = "CreateSubscriptionDefinition"

// CreateSubscriptionDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a subscription definition. You may provide the initial version of
// the subscription definition now or use ''CreateSubscriptionDefinitionVersion''
// at a later time.
//
//    // Example sending a request using CreateSubscriptionDefinitionRequest.
//    req := client.CreateSubscriptionDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateSubscriptionDefinition
func (c *Client) CreateSubscriptionDefinitionRequest(input *CreateSubscriptionDefinitionInput) CreateSubscriptionDefinitionRequest {
	op := &aws.Operation{
		Name:       opCreateSubscriptionDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/definition/subscriptions",
	}

	if input == nil {
		input = &CreateSubscriptionDefinitionInput{}
	}

	req := c.newRequest(op, input, &CreateSubscriptionDefinitionOutput{})
	return CreateSubscriptionDefinitionRequest{Request: req, Input: input, Copy: c.CreateSubscriptionDefinitionRequest}
}

// CreateSubscriptionDefinitionRequest is the request type for the
// CreateSubscriptionDefinition API operation.
type CreateSubscriptionDefinitionRequest struct {
	*aws.Request
	Input *CreateSubscriptionDefinitionInput
	Copy  func(*CreateSubscriptionDefinitionInput) CreateSubscriptionDefinitionRequest
}

// Send marshals and sends the CreateSubscriptionDefinition API request.
func (r CreateSubscriptionDefinitionRequest) Send(ctx context.Context) (*CreateSubscriptionDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSubscriptionDefinitionResponse{
		CreateSubscriptionDefinitionOutput: r.Request.Data.(*CreateSubscriptionDefinitionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSubscriptionDefinitionResponse is the response type for the
// CreateSubscriptionDefinition API operation.
type CreateSubscriptionDefinitionResponse struct {
	*CreateSubscriptionDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSubscriptionDefinition request.
func (r *CreateSubscriptionDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
