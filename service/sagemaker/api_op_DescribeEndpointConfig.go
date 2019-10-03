// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeEndpointConfigInput
type DescribeEndpointConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of the endpoint configuration.
	//
	// EndpointConfigName is a required field
	EndpointConfigName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEndpointConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEndpointConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEndpointConfigInput"}

	if s.EndpointConfigName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointConfigName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeEndpointConfigOutput
type DescribeEndpointConfigOutput struct {
	_ struct{} `type:"structure"`

	// A timestamp that shows when the endpoint configuration was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `json:"api.sagemaker:DescribeEndpointConfigOutput:CreationTime" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The Amazon Resource Name (ARN) of the endpoint configuration.
	//
	// EndpointConfigArn is a required field
	EndpointConfigArn *string `json:"api.sagemaker:DescribeEndpointConfigOutput:EndpointConfigArn" min:"20" type:"string" required:"true"`

	// Name of the Amazon SageMaker endpoint configuration.
	//
	// EndpointConfigName is a required field
	EndpointConfigName *string `json:"api.sagemaker:DescribeEndpointConfigOutput:EndpointConfigName" type:"string" required:"true"`

	// AWS KMS key ID Amazon SageMaker uses to encrypt data when storing it on the
	// ML storage volume attached to the instance.
	KmsKeyId *string `json:"api.sagemaker:DescribeEndpointConfigOutput:KmsKeyId" type:"string"`

	// An array of ProductionVariant objects, one for each model that you want to
	// host at this endpoint.
	//
	// ProductionVariants is a required field
	ProductionVariants []ProductionVariant `json:"api.sagemaker:DescribeEndpointConfigOutput:ProductionVariants" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeEndpointConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEndpointConfig = "DescribeEndpointConfig"

// DescribeEndpointConfigRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns the description of an endpoint configuration created using the CreateEndpointConfig
// API.
//
//    // Example sending a request using DescribeEndpointConfigRequest.
//    req := client.DescribeEndpointConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeEndpointConfig
func (c *Client) DescribeEndpointConfigRequest(input *DescribeEndpointConfigInput) DescribeEndpointConfigRequest {
	op := &aws.Operation{
		Name:       opDescribeEndpointConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEndpointConfigInput{}
	}

	req := c.newRequest(op, input, &DescribeEndpointConfigOutput{})
	return DescribeEndpointConfigRequest{Request: req, Input: input, Copy: c.DescribeEndpointConfigRequest}
}

// DescribeEndpointConfigRequest is the request type for the
// DescribeEndpointConfig API operation.
type DescribeEndpointConfigRequest struct {
	*aws.Request
	Input *DescribeEndpointConfigInput
	Copy  func(*DescribeEndpointConfigInput) DescribeEndpointConfigRequest
}

// Send marshals and sends the DescribeEndpointConfig API request.
func (r DescribeEndpointConfigRequest) Send(ctx context.Context) (*DescribeEndpointConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEndpointConfigResponse{
		DescribeEndpointConfigOutput: r.Request.Data.(*DescribeEndpointConfigOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEndpointConfigResponse is the response type for the
// DescribeEndpointConfig API operation.
type DescribeEndpointConfigResponse struct {
	*DescribeEndpointConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEndpointConfig request.
func (r *DescribeEndpointConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
