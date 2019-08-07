// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetSecurityConfigurationRequest
type GetSecurityConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the security configuration to retrieve.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSecurityConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSecurityConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSecurityConfigurationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetSecurityConfigurationResponse
type GetSecurityConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The requested security configuration.
	SecurityConfiguration *SecurityConfiguration `json:"glue:GetSecurityConfigurationOutput:SecurityConfiguration" type:"structure"`
}

// String returns the string representation
func (s GetSecurityConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSecurityConfiguration = "GetSecurityConfiguration"

// GetSecurityConfigurationRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves a specified security configuration.
//
//    // Example sending a request using GetSecurityConfigurationRequest.
//    req := client.GetSecurityConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetSecurityConfiguration
func (c *Client) GetSecurityConfigurationRequest(input *GetSecurityConfigurationInput) GetSecurityConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetSecurityConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSecurityConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetSecurityConfigurationOutput{})
	return GetSecurityConfigurationRequest{Request: req, Input: input, Copy: c.GetSecurityConfigurationRequest}
}

// GetSecurityConfigurationRequest is the request type for the
// GetSecurityConfiguration API operation.
type GetSecurityConfigurationRequest struct {
	*aws.Request
	Input *GetSecurityConfigurationInput
	Copy  func(*GetSecurityConfigurationInput) GetSecurityConfigurationRequest
}

// Send marshals and sends the GetSecurityConfiguration API request.
func (r GetSecurityConfigurationRequest) Send(ctx context.Context) (*GetSecurityConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSecurityConfigurationResponse{
		GetSecurityConfigurationOutput: r.Request.Data.(*GetSecurityConfigurationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSecurityConfigurationResponse is the response type for the
// GetSecurityConfiguration API operation.
type GetSecurityConfigurationResponse struct {
	*GetSecurityConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSecurityConfiguration request.
func (r *GetSecurityConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
