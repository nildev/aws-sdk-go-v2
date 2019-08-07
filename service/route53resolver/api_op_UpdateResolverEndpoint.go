// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/UpdateResolverEndpointRequest
type UpdateResolverEndpointInput struct {
	_ struct{} `type:"structure"`

	// The name of the resolver endpoint that you want to update.
	Name *string `type:"string"`

	// The ID of the resolver endpoint that you want to update.
	//
	// ResolverEndpointId is a required field
	ResolverEndpointId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateResolverEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateResolverEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateResolverEndpointInput"}

	if s.ResolverEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverEndpointId"))
	}
	if s.ResolverEndpointId != nil && len(*s.ResolverEndpointId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverEndpointId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/UpdateResolverEndpointResponse
type UpdateResolverEndpointOutput struct {
	_ struct{} `type:"structure"`

	// The response to an UpdateResolverEndpoint request.
	ResolverEndpoint *ResolverEndpoint `json:"route53resolver:UpdateResolverEndpointOutput:ResolverEndpoint" type:"structure"`
}

// String returns the string representation
func (s UpdateResolverEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateResolverEndpoint = "UpdateResolverEndpoint"

// UpdateResolverEndpointRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Updates the name of an inbound or an outbound resolver endpoint.
//
//    // Example sending a request using UpdateResolverEndpointRequest.
//    req := client.UpdateResolverEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/UpdateResolverEndpoint
func (c *Client) UpdateResolverEndpointRequest(input *UpdateResolverEndpointInput) UpdateResolverEndpointRequest {
	op := &aws.Operation{
		Name:       opUpdateResolverEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateResolverEndpointInput{}
	}

	req := c.newRequest(op, input, &UpdateResolverEndpointOutput{})
	return UpdateResolverEndpointRequest{Request: req, Input: input, Copy: c.UpdateResolverEndpointRequest}
}

// UpdateResolverEndpointRequest is the request type for the
// UpdateResolverEndpoint API operation.
type UpdateResolverEndpointRequest struct {
	*aws.Request
	Input *UpdateResolverEndpointInput
	Copy  func(*UpdateResolverEndpointInput) UpdateResolverEndpointRequest
}

// Send marshals and sends the UpdateResolverEndpoint API request.
func (r UpdateResolverEndpointRequest) Send(ctx context.Context) (*UpdateResolverEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateResolverEndpointResponse{
		UpdateResolverEndpointOutput: r.Request.Data.(*UpdateResolverEndpointOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateResolverEndpointResponse is the response type for the
// UpdateResolverEndpoint API operation.
type UpdateResolverEndpointResponse struct {
	*UpdateResolverEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateResolverEndpoint request.
func (r *UpdateResolverEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
