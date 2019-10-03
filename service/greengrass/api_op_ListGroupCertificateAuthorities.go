// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListGroupCertificateAuthoritiesRequest
type ListGroupCertificateAuthoritiesInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s ListGroupCertificateAuthoritiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGroupCertificateAuthoritiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGroupCertificateAuthoritiesInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGroupCertificateAuthoritiesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListGroupCertificateAuthoritiesResponse
type ListGroupCertificateAuthoritiesOutput struct {
	_ struct{} `type:"structure"`

	// A list of certificate authorities associated with the group.
	GroupCertificateAuthorities []GroupCertificateAuthorityProperties `json:"greengrass:ListGroupCertificateAuthoritiesOutput:GroupCertificateAuthorities" type:"list"`
}

// String returns the string representation
func (s ListGroupCertificateAuthoritiesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGroupCertificateAuthoritiesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GroupCertificateAuthorities != nil {
		v := s.GroupCertificateAuthorities

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "GroupCertificateAuthorities", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListGroupCertificateAuthorities = "ListGroupCertificateAuthorities"

// ListGroupCertificateAuthoritiesRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves the current CAs for a group.
//
//    // Example sending a request using ListGroupCertificateAuthoritiesRequest.
//    req := client.ListGroupCertificateAuthoritiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListGroupCertificateAuthorities
func (c *Client) ListGroupCertificateAuthoritiesRequest(input *ListGroupCertificateAuthoritiesInput) ListGroupCertificateAuthoritiesRequest {
	op := &aws.Operation{
		Name:       opListGroupCertificateAuthorities,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/groups/{GroupId}/certificateauthorities",
	}

	if input == nil {
		input = &ListGroupCertificateAuthoritiesInput{}
	}

	req := c.newRequest(op, input, &ListGroupCertificateAuthoritiesOutput{})
	return ListGroupCertificateAuthoritiesRequest{Request: req, Input: input, Copy: c.ListGroupCertificateAuthoritiesRequest}
}

// ListGroupCertificateAuthoritiesRequest is the request type for the
// ListGroupCertificateAuthorities API operation.
type ListGroupCertificateAuthoritiesRequest struct {
	*aws.Request
	Input *ListGroupCertificateAuthoritiesInput
	Copy  func(*ListGroupCertificateAuthoritiesInput) ListGroupCertificateAuthoritiesRequest
}

// Send marshals and sends the ListGroupCertificateAuthorities API request.
func (r ListGroupCertificateAuthoritiesRequest) Send(ctx context.Context) (*ListGroupCertificateAuthoritiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupCertificateAuthoritiesResponse{
		ListGroupCertificateAuthoritiesOutput: r.Request.Data.(*ListGroupCertificateAuthoritiesOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGroupCertificateAuthoritiesResponse is the response type for the
// ListGroupCertificateAuthorities API operation.
type ListGroupCertificateAuthoritiesResponse struct {
	*ListGroupCertificateAuthoritiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroupCertificateAuthorities request.
func (r *ListGroupCertificateAuthoritiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
