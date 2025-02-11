// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AddFacetToObjectRequest
type AddFacetToObjectInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// the object resides. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// Attributes on the facet that you are adding to the object.
	ObjectAttributeList []AttributeKeyAndValue `type:"list"`

	// A reference to the object you are adding the specified facet to.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`

	// Identifiers for the facet that you are adding to the object. See SchemaFacet
	// for details.
	//
	// SchemaFacet is a required field
	SchemaFacet *SchemaFacet `type:"structure" required:"true"`
}

// String returns the string representation
func (s AddFacetToObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddFacetToObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddFacetToObjectInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if s.SchemaFacet == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaFacet"))
	}
	if s.ObjectAttributeList != nil {
		for i, v := range s.ObjectAttributeList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ObjectAttributeList", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SchemaFacet != nil {
		if err := s.SchemaFacet.Validate(); err != nil {
			invalidParams.AddNested("SchemaFacet", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AddFacetToObjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ObjectAttributeList != nil {
		v := s.ObjectAttributeList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ObjectAttributeList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if s.SchemaFacet != nil {
		v := s.SchemaFacet

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SchemaFacet", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AddFacetToObjectResponse
type AddFacetToObjectOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddFacetToObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AddFacetToObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAddFacetToObject = "AddFacetToObject"

// AddFacetToObjectRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Adds a new Facet to an object. An object can have more than one facet applied
// on it.
//
//    // Example sending a request using AddFacetToObjectRequest.
//    req := client.AddFacetToObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AddFacetToObject
func (c *Client) AddFacetToObjectRequest(input *AddFacetToObjectInput) AddFacetToObjectRequest {
	op := &aws.Operation{
		Name:       opAddFacetToObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/facets",
	}

	if input == nil {
		input = &AddFacetToObjectInput{}
	}

	req := c.newRequest(op, input, &AddFacetToObjectOutput{})
	return AddFacetToObjectRequest{Request: req, Input: input, Copy: c.AddFacetToObjectRequest}
}

// AddFacetToObjectRequest is the request type for the
// AddFacetToObject API operation.
type AddFacetToObjectRequest struct {
	*aws.Request
	Input *AddFacetToObjectInput
	Copy  func(*AddFacetToObjectInput) AddFacetToObjectRequest
}

// Send marshals and sends the AddFacetToObject API request.
func (r AddFacetToObjectRequest) Send(ctx context.Context) (*AddFacetToObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddFacetToObjectResponse{
		AddFacetToObjectOutput: r.Request.Data.(*AddFacetToObjectOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddFacetToObjectResponse is the response type for the
// AddFacetToObject API operation.
type AddFacetToObjectResponse struct {
	*AddFacetToObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddFacetToObject request.
func (r *AddFacetToObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
