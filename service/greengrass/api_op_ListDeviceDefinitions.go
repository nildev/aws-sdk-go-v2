// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListDeviceDefinitionsRequest
type ListDeviceDefinitionsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListDeviceDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeviceDefinitionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListDeviceDefinitionsResponse
type ListDeviceDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	Definitions []DefinitionInformation `json:"greengrass:ListDeviceDefinitionsOutput:Definitions" type:"list"`

	NextToken *string `json:"greengrass:ListDeviceDefinitionsOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s ListDeviceDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeviceDefinitionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Definitions != nil {
		v := s.Definitions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Definitions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDeviceDefinitions = "ListDeviceDefinitions"

// ListDeviceDefinitionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves a list of device definitions.
//
//    // Example sending a request using ListDeviceDefinitionsRequest.
//    req := client.ListDeviceDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListDeviceDefinitions
func (c *Client) ListDeviceDefinitionsRequest(input *ListDeviceDefinitionsInput) ListDeviceDefinitionsRequest {
	op := &aws.Operation{
		Name:       opListDeviceDefinitions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/devices",
	}

	if input == nil {
		input = &ListDeviceDefinitionsInput{}
	}

	req := c.newRequest(op, input, &ListDeviceDefinitionsOutput{})
	return ListDeviceDefinitionsRequest{Request: req, Input: input, Copy: c.ListDeviceDefinitionsRequest}
}

// ListDeviceDefinitionsRequest is the request type for the
// ListDeviceDefinitions API operation.
type ListDeviceDefinitionsRequest struct {
	*aws.Request
	Input *ListDeviceDefinitionsInput
	Copy  func(*ListDeviceDefinitionsInput) ListDeviceDefinitionsRequest
}

// Send marshals and sends the ListDeviceDefinitions API request.
func (r ListDeviceDefinitionsRequest) Send(ctx context.Context) (*ListDeviceDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeviceDefinitionsResponse{
		ListDeviceDefinitionsOutput: r.Request.Data.(*ListDeviceDefinitionsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDeviceDefinitionsResponse is the response type for the
// ListDeviceDefinitions API operation.
type ListDeviceDefinitionsResponse struct {
	*ListDeviceDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeviceDefinitions request.
func (r *ListDeviceDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
