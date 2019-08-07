// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/UpdateVirtualRouterInput
type UpdateVirtualRouterInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of therequest. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// The name of the service mesh that the virtual router resides in.
	//
	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// The new virtual router specification to apply. This overwrites the existing
	// data.
	//
	// Spec is a required field
	Spec *VirtualRouterSpec `locationName:"spec" type:"structure" required:"true"`

	// The name of the virtual router to update.
	//
	// VirtualRouterName is a required field
	VirtualRouterName *string `location:"uri" locationName:"virtualRouterName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateVirtualRouterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVirtualRouterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateVirtualRouterInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.Spec == nil {
		invalidParams.Add(aws.NewErrParamRequired("Spec"))
	}

	if s.VirtualRouterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualRouterName"))
	}
	if s.VirtualRouterName != nil && len(*s.VirtualRouterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualRouterName", 1))
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			invalidParams.AddNested("Spec", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateVirtualRouterInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Spec != nil {
		v := s.Spec

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "spec", v, metadata)
	}
	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualRouterName != nil {
		v := *s.VirtualRouterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "virtualRouterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/UpdateVirtualRouterOutput
type UpdateVirtualRouterOutput struct {
	_ struct{} `type:"structure" payload:"VirtualRouter"`

	// A full description of the virtual router that was updated.
	//
	// VirtualRouter is a required field
	VirtualRouter *VirtualRouterData `json:"appmesh:UpdateVirtualRouterOutput:VirtualRouter" locationName:"virtualRouter" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateVirtualRouterOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateVirtualRouterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VirtualRouter != nil {
		v := s.VirtualRouter

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "virtualRouter", v, metadata)
	}
	return nil
}

const opUpdateVirtualRouter = "UpdateVirtualRouter"

// UpdateVirtualRouterRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Updates an existing virtual router in a specified service mesh.
//
//    // Example sending a request using UpdateVirtualRouterRequest.
//    req := client.UpdateVirtualRouterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/UpdateVirtualRouter
func (c *Client) UpdateVirtualRouterRequest(input *UpdateVirtualRouterInput) UpdateVirtualRouterRequest {
	op := &aws.Operation{
		Name:       opUpdateVirtualRouter,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualRouters/{virtualRouterName}",
	}

	if input == nil {
		input = &UpdateVirtualRouterInput{}
	}

	req := c.newRequest(op, input, &UpdateVirtualRouterOutput{})
	return UpdateVirtualRouterRequest{Request: req, Input: input, Copy: c.UpdateVirtualRouterRequest}
}

// UpdateVirtualRouterRequest is the request type for the
// UpdateVirtualRouter API operation.
type UpdateVirtualRouterRequest struct {
	*aws.Request
	Input *UpdateVirtualRouterInput
	Copy  func(*UpdateVirtualRouterInput) UpdateVirtualRouterRequest
}

// Send marshals and sends the UpdateVirtualRouter API request.
func (r UpdateVirtualRouterRequest) Send(ctx context.Context) (*UpdateVirtualRouterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVirtualRouterResponse{
		UpdateVirtualRouterOutput: r.Request.Data.(*UpdateVirtualRouterOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVirtualRouterResponse is the response type for the
// UpdateVirtualRouter API operation.
type UpdateVirtualRouterResponse struct {
	*UpdateVirtualRouterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVirtualRouter request.
func (r *UpdateVirtualRouterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
