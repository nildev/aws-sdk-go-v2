// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetFunctionConfigurationRequest
type GetFunctionConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the Lambda function, version, or alias.
	//
	// Name formats
	//
	//    * Function name - my-function (name-only), my-function:v1 (with alias).
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function
	// name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// Specify a version or alias to get details about a published version of the
	// function.
	Qualifier *string `location:"querystring" locationName:"Qualifier" min:"1" type:"string"`
}

// String returns the string representation
func (s GetFunctionConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFunctionConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFunctionConfigurationInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.Qualifier != nil && len(*s.Qualifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Qualifier", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFunctionConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Qualifier != nil {
		v := *s.Qualifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Qualifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Details about a function's configuration.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/FunctionConfiguration
type GetFunctionConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string `json:"lambda:GetFunctionConfigurationOutput:CodeSha256" type:"string"`

	// The size of the function's deployment package, in bytes.
	CodeSize *int64 `json:"lambda:GetFunctionConfigurationOutput:CodeSize" type:"long"`

	// The function's dead letter queue.
	DeadLetterConfig *DeadLetterConfig `json:"lambda:GetFunctionConfigurationOutput:DeadLetterConfig" type:"structure"`

	// The function's description.
	Description *string `json:"lambda:GetFunctionConfigurationOutput:Description" type:"string"`

	// The function's environment variables.
	Environment *EnvironmentResponse `json:"lambda:GetFunctionConfigurationOutput:Environment" type:"structure"`

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string `json:"lambda:GetFunctionConfigurationOutput:FunctionArn" type:"string"`

	// The name of the function.
	FunctionName *string `json:"lambda:GetFunctionConfigurationOutput:FunctionName" min:"1" type:"string"`

	// The function that Lambda calls to begin executing your function.
	Handler *string `json:"lambda:GetFunctionConfigurationOutput:Handler" type:"string"`

	// The KMS key that's used to encrypt the function's environment variables.
	// This key is only returned if you've configured a customer-managed CMK.
	KMSKeyArn *string `json:"lambda:GetFunctionConfigurationOutput:KMSKeyArn" type:"string"`

	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string `json:"lambda:GetFunctionConfigurationOutput:LastModified" type:"string"`

	// The function's layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	Layers []Layer `json:"lambda:GetFunctionConfigurationOutput:Layers" type:"list"`

	// For Lambda@Edge functions, the ARN of the master function.
	MasterArn *string `json:"lambda:GetFunctionConfigurationOutput:MasterArn" type:"string"`

	// The memory that's allocated to the function.
	MemorySize *int64 `json:"lambda:GetFunctionConfigurationOutput:MemorySize" min:"128" type:"integer"`

	// The latest updated revision of the function or alias.
	RevisionId *string `json:"lambda:GetFunctionConfigurationOutput:RevisionId" type:"string"`

	// The function's execution role.
	Role *string `json:"lambda:GetFunctionConfigurationOutput:Role" type:"string"`

	// The runtime environment for the Lambda function.
	Runtime Runtime `json:"lambda:GetFunctionConfigurationOutput:Runtime" type:"string" enum:"true"`

	// The amount of time that Lambda allows a function to run before stopping it.
	Timeout *int64 `json:"lambda:GetFunctionConfigurationOutput:Timeout" min:"1" type:"integer"`

	// The function's AWS X-Ray tracing configuration.
	TracingConfig *TracingConfigResponse `json:"lambda:GetFunctionConfigurationOutput:TracingConfig" type:"structure"`

	// The version of the Lambda function.
	Version *string `json:"lambda:GetFunctionConfigurationOutput:Version" min:"1" type:"string"`

	// The function's networking configuration.
	VpcConfig *VpcConfigResponse `json:"lambda:GetFunctionConfigurationOutput:VpcConfig" type:"structure"`
}

// String returns the string representation
func (s GetFunctionConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFunctionConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CodeSha256 != nil {
		v := *s.CodeSha256

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CodeSha256", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CodeSize != nil {
		v := *s.CodeSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CodeSize", protocol.Int64Value(v), metadata)
	}
	if s.DeadLetterConfig != nil {
		v := s.DeadLetterConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DeadLetterConfig", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Environment != nil {
		v := s.Environment

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Environment", v, metadata)
	}
	if s.FunctionArn != nil {
		v := *s.FunctionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Handler != nil {
		v := *s.Handler

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Handler", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KMSKeyArn != nil {
		v := *s.KMSKeyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KMSKeyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Layers != nil {
		v := s.Layers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Layers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MasterArn != nil {
		v := *s.MasterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MasterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MemorySize != nil {
		v := *s.MemorySize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MemorySize", protocol.Int64Value(v), metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Role != nil {
		v := *s.Role

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Role", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Runtime) > 0 {
		v := s.Runtime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Runtime", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Timeout != nil {
		v := *s.Timeout

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Timeout", protocol.Int64Value(v), metadata)
	}
	if s.TracingConfig != nil {
		v := s.TracingConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TracingConfig", v, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VpcConfig != nil {
		v := s.VpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VpcConfig", v, metadata)
	}
	return nil
}

const opGetFunctionConfiguration = "GetFunctionConfiguration"

// GetFunctionConfigurationRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns the version-specific settings of a Lambda function or version. The
// output includes only options that can vary between versions of a function.
// To modify these settings, use UpdateFunctionConfiguration.
//
// To get all of a function's details, including function-level settings, use
// GetFunction.
//
//    // Example sending a request using GetFunctionConfigurationRequest.
//    req := client.GetFunctionConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetFunctionConfiguration
func (c *Client) GetFunctionConfigurationRequest(input *GetFunctionConfigurationInput) GetFunctionConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetFunctionConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/configuration",
	}

	if input == nil {
		input = &GetFunctionConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetFunctionConfigurationOutput{})
	return GetFunctionConfigurationRequest{Request: req, Input: input, Copy: c.GetFunctionConfigurationRequest}
}

// GetFunctionConfigurationRequest is the request type for the
// GetFunctionConfiguration API operation.
type GetFunctionConfigurationRequest struct {
	*aws.Request
	Input *GetFunctionConfigurationInput
	Copy  func(*GetFunctionConfigurationInput) GetFunctionConfigurationRequest
}

// Send marshals and sends the GetFunctionConfiguration API request.
func (r GetFunctionConfigurationRequest) Send(ctx context.Context) (*GetFunctionConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFunctionConfigurationResponse{
		GetFunctionConfigurationOutput: r.Request.Data.(*GetFunctionConfigurationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFunctionConfigurationResponse is the response type for the
// GetFunctionConfiguration API operation.
type GetFunctionConfigurationResponse struct {
	*GetFunctionConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFunctionConfiguration request.
func (r *GetFunctionConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
