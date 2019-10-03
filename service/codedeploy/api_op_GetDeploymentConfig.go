// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetDeploymentConfig operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetDeploymentConfigInput
type GetDeploymentConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of a deployment configuration associated with the IAM user or AWS
	// account.
	//
	// DeploymentConfigName is a required field
	DeploymentConfigName *string `locationName:"deploymentConfigName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeploymentConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeploymentConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeploymentConfigInput"}

	if s.DeploymentConfigName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentConfigName"))
	}
	if s.DeploymentConfigName != nil && len(*s.DeploymentConfigName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentConfigName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetDeploymentConfig operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetDeploymentConfigOutput
type GetDeploymentConfigOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deployment configuration.
	DeploymentConfigInfo *DeploymentConfigInfo `json:"codedeploy:GetDeploymentConfigOutput:DeploymentConfigInfo" locationName:"deploymentConfigInfo" type:"structure"`
}

// String returns the string representation
func (s GetDeploymentConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDeploymentConfig = "GetDeploymentConfig"

// GetDeploymentConfigRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets information about a deployment configuration.
//
//    // Example sending a request using GetDeploymentConfigRequest.
//    req := client.GetDeploymentConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetDeploymentConfig
func (c *Client) GetDeploymentConfigRequest(input *GetDeploymentConfigInput) GetDeploymentConfigRequest {
	op := &aws.Operation{
		Name:       opGetDeploymentConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDeploymentConfigInput{}
	}

	req := c.newRequest(op, input, &GetDeploymentConfigOutput{})
	return GetDeploymentConfigRequest{Request: req, Input: input, Copy: c.GetDeploymentConfigRequest}
}

// GetDeploymentConfigRequest is the request type for the
// GetDeploymentConfig API operation.
type GetDeploymentConfigRequest struct {
	*aws.Request
	Input *GetDeploymentConfigInput
	Copy  func(*GetDeploymentConfigInput) GetDeploymentConfigRequest
}

// Send marshals and sends the GetDeploymentConfig API request.
func (r GetDeploymentConfigRequest) Send(ctx context.Context) (*GetDeploymentConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeploymentConfigResponse{
		GetDeploymentConfigOutput: r.Request.Data.(*GetDeploymentConfigOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeploymentConfigResponse is the response type for the
// GetDeploymentConfig API operation.
type GetDeploymentConfigResponse struct {
	*GetDeploymentConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeploymentConfig request.
func (r *GetDeploymentConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
