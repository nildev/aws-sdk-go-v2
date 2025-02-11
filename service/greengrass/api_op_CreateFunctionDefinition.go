// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateFunctionDefinitionRequest
type CreateFunctionDefinitionInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	// Information about a function definition version.
	InitialVersion *FunctionDefinitionVersion `type:"structure"`

	Name *string `type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateFunctionDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFunctionDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateFunctionDefinitionResponse
type CreateFunctionDefinitionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `type:"string"`

	CreationTimestamp *string `type:"string"`

	Id *string `type:"string"`

	LastUpdatedTimestamp *string `type:"string"`

	LatestVersion *string `type:"string"`

	LatestVersionArn *string `type:"string"`

	Name *string `type:"string"`
}

// String returns the string representation
func (s CreateFunctionDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFunctionDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opCreateFunctionDefinition = "CreateFunctionDefinition"

// CreateFunctionDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a Lambda function definition which contains a list of Lambda functions
// and their configurations to be used in a group. You can create an initial
// version of the definition by providing a list of Lambda functions and their
// configurations now, or use ''CreateFunctionDefinitionVersion'' later.
//
//    // Example sending a request using CreateFunctionDefinitionRequest.
//    req := client.CreateFunctionDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateFunctionDefinition
func (c *Client) CreateFunctionDefinitionRequest(input *CreateFunctionDefinitionInput) CreateFunctionDefinitionRequest {
	op := &aws.Operation{
		Name:       opCreateFunctionDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/definition/functions",
	}

	if input == nil {
		input = &CreateFunctionDefinitionInput{}
	}

	req := c.newRequest(op, input, &CreateFunctionDefinitionOutput{})
	return CreateFunctionDefinitionRequest{Request: req, Input: input, Copy: c.CreateFunctionDefinitionRequest}
}

// CreateFunctionDefinitionRequest is the request type for the
// CreateFunctionDefinition API operation.
type CreateFunctionDefinitionRequest struct {
	*aws.Request
	Input *CreateFunctionDefinitionInput
	Copy  func(*CreateFunctionDefinitionInput) CreateFunctionDefinitionRequest
}

// Send marshals and sends the CreateFunctionDefinition API request.
func (r CreateFunctionDefinitionRequest) Send(ctx context.Context) (*CreateFunctionDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFunctionDefinitionResponse{
		CreateFunctionDefinitionOutput: r.Request.Data.(*CreateFunctionDefinitionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFunctionDefinitionResponse is the response type for the
// CreateFunctionDefinition API operation.
type CreateFunctionDefinitionResponse struct {
	*CreateFunctionDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFunctionDefinition request.
func (r *CreateFunctionDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
