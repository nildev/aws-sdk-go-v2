// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListResourceDefinitionVersionsRequest
type ListResourceDefinitionVersionsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	// ResourceDefinitionId is a required field
	ResourceDefinitionId *string `location:"uri" locationName:"ResourceDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s ListResourceDefinitionVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResourceDefinitionVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResourceDefinitionVersionsInput"}

	if s.ResourceDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListResourceDefinitionVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ResourceDefinitionId != nil {
		v := *s.ResourceDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ResourceDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListResourceDefinitionVersionsResponse
type ListResourceDefinitionVersionsOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `type:"string"`

	Versions []VersionInformation `type:"list"`
}

// String returns the string representation
func (s ListResourceDefinitionVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListResourceDefinitionVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Versions != nil {
		v := s.Versions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Versions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListResourceDefinitionVersions = "ListResourceDefinitionVersions"

// ListResourceDefinitionVersionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Lists the versions of a resource definition.
//
//    // Example sending a request using ListResourceDefinitionVersionsRequest.
//    req := client.ListResourceDefinitionVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListResourceDefinitionVersions
func (c *Client) ListResourceDefinitionVersionsRequest(input *ListResourceDefinitionVersionsInput) ListResourceDefinitionVersionsRequest {
	op := &aws.Operation{
		Name:       opListResourceDefinitionVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/resources/{ResourceDefinitionId}/versions",
	}

	if input == nil {
		input = &ListResourceDefinitionVersionsInput{}
	}

	req := c.newRequest(op, input, &ListResourceDefinitionVersionsOutput{})
	return ListResourceDefinitionVersionsRequest{Request: req, Input: input, Copy: c.ListResourceDefinitionVersionsRequest}
}

// ListResourceDefinitionVersionsRequest is the request type for the
// ListResourceDefinitionVersions API operation.
type ListResourceDefinitionVersionsRequest struct {
	*aws.Request
	Input *ListResourceDefinitionVersionsInput
	Copy  func(*ListResourceDefinitionVersionsInput) ListResourceDefinitionVersionsRequest
}

// Send marshals and sends the ListResourceDefinitionVersions API request.
func (r ListResourceDefinitionVersionsRequest) Send(ctx context.Context) (*ListResourceDefinitionVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResourceDefinitionVersionsResponse{
		ListResourceDefinitionVersionsOutput: r.Request.Data.(*ListResourceDefinitionVersionsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListResourceDefinitionVersionsResponse is the response type for the
// ListResourceDefinitionVersions API operation.
type ListResourceDefinitionVersionsResponse struct {
	*ListResourceDefinitionVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResourceDefinitionVersions request.
func (r *ListResourceDefinitionVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
