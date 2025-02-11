// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetEndpointRequest
type GetEndpointInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// EndpointId is a required field
	EndpointId *string `location:"uri" locationName:"endpoint-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEndpointInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.EndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEndpointInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndpointId != nil {
		v := *s.EndpointId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "endpoint-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetEndpointResponse
type GetEndpointOutput struct {
	_ struct{} `type:"structure" payload:"EndpointResponse"`

	// Provides information about the channel type and other settings for an endpoint.
	//
	// EndpointResponse is a required field
	EndpointResponse *EndpointResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEndpointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EndpointResponse != nil {
		v := s.EndpointResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "EndpointResponse", v, metadata)
	}
	return nil
}

const opGetEndpoint = "GetEndpoint"

// GetEndpointRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the settings and attributes of a specific endpoint
// for an application.
//
//    // Example sending a request using GetEndpointRequest.
//    req := client.GetEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetEndpoint
func (c *Client) GetEndpointRequest(input *GetEndpointInput) GetEndpointRequest {
	op := &aws.Operation{
		Name:       opGetEndpoint,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/endpoints/{endpoint-id}",
	}

	if input == nil {
		input = &GetEndpointInput{}
	}

	req := c.newRequest(op, input, &GetEndpointOutput{})
	return GetEndpointRequest{Request: req, Input: input, Copy: c.GetEndpointRequest}
}

// GetEndpointRequest is the request type for the
// GetEndpoint API operation.
type GetEndpointRequest struct {
	*aws.Request
	Input *GetEndpointInput
	Copy  func(*GetEndpointInput) GetEndpointRequest
}

// Send marshals and sends the GetEndpoint API request.
func (r GetEndpointRequest) Send(ctx context.Context) (*GetEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEndpointResponse{
		GetEndpointOutput: r.Request.Data.(*GetEndpointOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEndpointResponse is the response type for the
// GetEndpoint API operation.
type GetEndpointResponse struct {
	*GetEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEndpoint request.
func (r *GetEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
