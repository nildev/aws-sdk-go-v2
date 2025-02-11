// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteUserEndpointsRequest
type DeleteUserEndpointsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// UserId is a required field
	UserId *string `location:"uri" locationName:"user-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUserEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserEndpointsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteUserEndpointsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "user-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteUserEndpointsResponse
type DeleteUserEndpointsOutput struct {
	_ struct{} `type:"structure" payload:"EndpointsResponse"`

	// Provides information about all the endpoints that are associated with a user
	// ID.
	//
	// EndpointsResponse is a required field
	EndpointsResponse *EndpointsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteUserEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteUserEndpointsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EndpointsResponse != nil {
		v := s.EndpointsResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "EndpointsResponse", v, metadata)
	}
	return nil
}

const opDeleteUserEndpoints = "DeleteUserEndpoints"

// DeleteUserEndpointsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Deletes all the endpoints that are associated with a specific user ID.
//
//    // Example sending a request using DeleteUserEndpointsRequest.
//    req := client.DeleteUserEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteUserEndpoints
func (c *Client) DeleteUserEndpointsRequest(input *DeleteUserEndpointsInput) DeleteUserEndpointsRequest {
	op := &aws.Operation{
		Name:       opDeleteUserEndpoints,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/users/{user-id}",
	}

	if input == nil {
		input = &DeleteUserEndpointsInput{}
	}

	req := c.newRequest(op, input, &DeleteUserEndpointsOutput{})
	return DeleteUserEndpointsRequest{Request: req, Input: input, Copy: c.DeleteUserEndpointsRequest}
}

// DeleteUserEndpointsRequest is the request type for the
// DeleteUserEndpoints API operation.
type DeleteUserEndpointsRequest struct {
	*aws.Request
	Input *DeleteUserEndpointsInput
	Copy  func(*DeleteUserEndpointsInput) DeleteUserEndpointsRequest
}

// Send marshals and sends the DeleteUserEndpoints API request.
func (r DeleteUserEndpointsRequest) Send(ctx context.Context) (*DeleteUserEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserEndpointsResponse{
		DeleteUserEndpointsOutput: r.Request.Data.(*DeleteUserEndpointsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserEndpointsResponse is the response type for the
// DeleteUserEndpoints API operation.
type DeleteUserEndpointsResponse struct {
	*DeleteUserEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUserEndpoints request.
func (r *DeleteUserEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
