// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/GetTypedLinkFacetInformationRequest
type GetTypedLinkFacetInformationInput struct {
	_ struct{} `type:"structure"`

	// The unique name of the typed link facet.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTypedLinkFacetInformationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTypedLinkFacetInformationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTypedLinkFacetInformationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTypedLinkFacetInformationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/GetTypedLinkFacetInformationResponse
type GetTypedLinkFacetInformationOutput struct {
	_ struct{} `type:"structure"`

	// The order of identity attributes for the facet, from most significant to
	// least significant. The ability to filter typed links considers the order
	// that the attributes are defined on the typed link facet. When providing ranges
	// to typed link selection, any inexact ranges must be specified at the end.
	// Any attributes that do not have a range specified are presumed to match the
	// entire range. Filters are interpreted in the order of the attributes on the
	// typed link facet, not the order in which they are supplied to any API calls.
	// For more information about identity attributes, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
	IdentityAttributeOrder []string `json:"clouddirectory:GetTypedLinkFacetInformationOutput:IdentityAttributeOrder" type:"list"`
}

// String returns the string representation
func (s GetTypedLinkFacetInformationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTypedLinkFacetInformationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IdentityAttributeOrder != nil {
		v := s.IdentityAttributeOrder

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "IdentityAttributeOrder", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opGetTypedLinkFacetInformation = "GetTypedLinkFacetInformation"

// GetTypedLinkFacetInformationRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Returns the identity attribute order for a specific TypedLinkFacet. For more
// information, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
//
//    // Example sending a request using GetTypedLinkFacetInformationRequest.
//    req := client.GetTypedLinkFacetInformationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/GetTypedLinkFacetInformation
func (c *Client) GetTypedLinkFacetInformationRequest(input *GetTypedLinkFacetInformationInput) GetTypedLinkFacetInformationRequest {
	op := &aws.Operation{
		Name:       opGetTypedLinkFacetInformation,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/typedlink/facet/get",
	}

	if input == nil {
		input = &GetTypedLinkFacetInformationInput{}
	}

	req := c.newRequest(op, input, &GetTypedLinkFacetInformationOutput{})
	return GetTypedLinkFacetInformationRequest{Request: req, Input: input, Copy: c.GetTypedLinkFacetInformationRequest}
}

// GetTypedLinkFacetInformationRequest is the request type for the
// GetTypedLinkFacetInformation API operation.
type GetTypedLinkFacetInformationRequest struct {
	*aws.Request
	Input *GetTypedLinkFacetInformationInput
	Copy  func(*GetTypedLinkFacetInformationInput) GetTypedLinkFacetInformationRequest
}

// Send marshals and sends the GetTypedLinkFacetInformation API request.
func (r GetTypedLinkFacetInformationRequest) Send(ctx context.Context) (*GetTypedLinkFacetInformationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTypedLinkFacetInformationResponse{
		GetTypedLinkFacetInformationOutput: r.Request.Data.(*GetTypedLinkFacetInformationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTypedLinkFacetInformationResponse is the response type for the
// GetTypedLinkFacetInformation API operation.
type GetTypedLinkFacetInformationResponse struct {
	*GetTypedLinkFacetInformationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTypedLinkFacetInformation request.
func (r *GetTypedLinkFacetInformationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
