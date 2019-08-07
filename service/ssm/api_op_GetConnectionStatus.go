// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetConnectionStatusRequest
type GetConnectionStatusInput struct {
	_ struct{} `type:"structure"`

	// The ID of the instance.
	//
	// Target is a required field
	Target *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetConnectionStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConnectionStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConnectionStatusInput"}

	if s.Target == nil {
		invalidParams.Add(aws.NewErrParamRequired("Target"))
	}
	if s.Target != nil && len(*s.Target) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Target", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetConnectionStatusResponse
type GetConnectionStatusOutput struct {
	_ struct{} `type:"structure"`

	// The status of the connection to the instance. For example, 'Connected' or
	// 'Not Connected'.
	Status ConnectionStatus `json:"ssm:GetConnectionStatusOutput:Status" type:"string" enum:"true"`

	// The ID of the instance to check connection status.
	Target *string `json:"ssm:GetConnectionStatusOutput:Target" min:"1" type:"string"`
}

// String returns the string representation
func (s GetConnectionStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetConnectionStatus = "GetConnectionStatus"

// GetConnectionStatusRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves the Session Manager connection status for an instance to determine
// whether it is connected and ready to receive Session Manager connections.
//
//    // Example sending a request using GetConnectionStatusRequest.
//    req := client.GetConnectionStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetConnectionStatus
func (c *Client) GetConnectionStatusRequest(input *GetConnectionStatusInput) GetConnectionStatusRequest {
	op := &aws.Operation{
		Name:       opGetConnectionStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConnectionStatusInput{}
	}

	req := c.newRequest(op, input, &GetConnectionStatusOutput{})
	return GetConnectionStatusRequest{Request: req, Input: input, Copy: c.GetConnectionStatusRequest}
}

// GetConnectionStatusRequest is the request type for the
// GetConnectionStatus API operation.
type GetConnectionStatusRequest struct {
	*aws.Request
	Input *GetConnectionStatusInput
	Copy  func(*GetConnectionStatusInput) GetConnectionStatusRequest
}

// Send marshals and sends the GetConnectionStatus API request.
func (r GetConnectionStatusRequest) Send(ctx context.Context) (*GetConnectionStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConnectionStatusResponse{
		GetConnectionStatusOutput: r.Request.Data.(*GetConnectionStatusOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConnectionStatusResponse is the response type for the
// GetConnectionStatus API operation.
type GetConnectionStatusResponse struct {
	*GetConnectionStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConnectionStatus request.
func (r *GetConnectionStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
