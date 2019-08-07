// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRealtimeEndpointInput struct {
	_ struct{} `type:"structure"`

	// The ID assigned to the MLModel during creation.
	//
	// MLModelId is a required field
	MLModelId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRealtimeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRealtimeEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRealtimeEndpointInput"}

	if s.MLModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MLModelId"))
	}
	if s.MLModelId != nil && len(*s.MLModelId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MLModelId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of an DeleteRealtimeEndpoint operation.
//
// The result contains the MLModelId and the endpoint information for the MLModel.
type DeleteRealtimeEndpointOutput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the MLModel. This value should
	// be identical to the value of the MLModelId in the request.
	MLModelId *string `json:"machinelearning:DeleteRealtimeEndpointOutput:MLModelId" min:"1" type:"string"`

	// The endpoint information of the MLModel
	RealtimeEndpointInfo *RealtimeEndpointInfo `json:"machinelearning:DeleteRealtimeEndpointOutput:RealtimeEndpointInfo" type:"structure"`
}

// String returns the string representation
func (s DeleteRealtimeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRealtimeEndpoint = "DeleteRealtimeEndpoint"

// DeleteRealtimeEndpointRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Deletes a real time endpoint of an MLModel.
//
//    // Example sending a request using DeleteRealtimeEndpointRequest.
//    req := client.DeleteRealtimeEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteRealtimeEndpointRequest(input *DeleteRealtimeEndpointInput) DeleteRealtimeEndpointRequest {
	op := &aws.Operation{
		Name:       opDeleteRealtimeEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRealtimeEndpointInput{}
	}

	req := c.newRequest(op, input, &DeleteRealtimeEndpointOutput{})
	return DeleteRealtimeEndpointRequest{Request: req, Input: input, Copy: c.DeleteRealtimeEndpointRequest}
}

// DeleteRealtimeEndpointRequest is the request type for the
// DeleteRealtimeEndpoint API operation.
type DeleteRealtimeEndpointRequest struct {
	*aws.Request
	Input *DeleteRealtimeEndpointInput
	Copy  func(*DeleteRealtimeEndpointInput) DeleteRealtimeEndpointRequest
}

// Send marshals and sends the DeleteRealtimeEndpoint API request.
func (r DeleteRealtimeEndpointRequest) Send(ctx context.Context) (*DeleteRealtimeEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRealtimeEndpointResponse{
		DeleteRealtimeEndpointOutput: r.Request.Data.(*DeleteRealtimeEndpointOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRealtimeEndpointResponse is the response type for the
// DeleteRealtimeEndpoint API operation.
type DeleteRealtimeEndpointResponse struct {
	*DeleteRealtimeEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRealtimeEndpoint request.
func (r *DeleteRealtimeEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
