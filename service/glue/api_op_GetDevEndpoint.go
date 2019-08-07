// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetDevEndpointRequest
type GetDevEndpointInput struct {
	_ struct{} `type:"structure"`

	// Name of the DevEndpoint to retrieve information for.
	//
	// EndpointName is a required field
	EndpointName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetDevEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDevEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDevEndpointInput"}

	if s.EndpointName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetDevEndpointResponse
type GetDevEndpointOutput struct {
	_ struct{} `type:"structure"`

	// A DevEndpoint definition.
	DevEndpoint *DevEndpoint `json:"glue:GetDevEndpointOutput:DevEndpoint" type:"structure"`
}

// String returns the string representation
func (s GetDevEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDevEndpoint = "GetDevEndpoint"

// GetDevEndpointRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves information about a specified development endpoint.
//
// When you create a development endpoint in a virtual private cloud (VPC),
// AWS Glue returns only a private IP address, and the public IP address field
// is not populated. When you create a non-VPC development endpoint, AWS Glue
// returns only a public IP address.
//
//    // Example sending a request using GetDevEndpointRequest.
//    req := client.GetDevEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetDevEndpoint
func (c *Client) GetDevEndpointRequest(input *GetDevEndpointInput) GetDevEndpointRequest {
	op := &aws.Operation{
		Name:       opGetDevEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDevEndpointInput{}
	}

	req := c.newRequest(op, input, &GetDevEndpointOutput{})
	return GetDevEndpointRequest{Request: req, Input: input, Copy: c.GetDevEndpointRequest}
}

// GetDevEndpointRequest is the request type for the
// GetDevEndpoint API operation.
type GetDevEndpointRequest struct {
	*aws.Request
	Input *GetDevEndpointInput
	Copy  func(*GetDevEndpointInput) GetDevEndpointRequest
}

// Send marshals and sends the GetDevEndpoint API request.
func (r GetDevEndpointRequest) Send(ctx context.Context) (*GetDevEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDevEndpointResponse{
		GetDevEndpointOutput: r.Request.Data.(*GetDevEndpointOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDevEndpointResponse is the response type for the
// GetDevEndpoint API operation.
type GetDevEndpointResponse struct {
	*GetDevEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDevEndpoint request.
func (r *GetDevEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
