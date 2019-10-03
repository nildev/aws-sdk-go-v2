// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetPipeline action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetPipelineInput
type GetPipelineInput struct {
	_ struct{} `type:"structure"`

	// The name of the pipeline for which you want to get information. Pipeline
	// names must be unique under an Amazon Web Services (AWS) user account.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The version number of the pipeline. If you do not specify a version, defaults
	// to the most current version.
	Version *int64 `locationName:"version" min:"1" type:"integer"`
}

// String returns the string representation
func (s GetPipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPipelineInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Version != nil && *s.Version < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Version", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetPipeline action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetPipelineOutput
type GetPipelineOutput struct {
	_ struct{} `type:"structure"`

	// Represents the pipeline metadata information returned as part of the output
	// of a GetPipeline action.
	Metadata *PipelineMetadata `json:"codepipeline:GetPipelineOutput:Metadata" locationName:"metadata" type:"structure"`

	// Represents the structure of actions and stages to be performed in the pipeline.
	Pipeline *PipelineDeclaration `json:"codepipeline:GetPipelineOutput:Pipeline" locationName:"pipeline" type:"structure"`
}

// String returns the string representation
func (s GetPipelineOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPipeline = "GetPipeline"

// GetPipelineRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Returns the metadata, structure, stages, and actions of a pipeline. Can be
// used to return the entire structure of a pipeline in JSON format, which can
// then be modified and used to update the pipeline structure with UpdatePipeline.
//
//    // Example sending a request using GetPipelineRequest.
//    req := client.GetPipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetPipeline
func (c *Client) GetPipelineRequest(input *GetPipelineInput) GetPipelineRequest {
	op := &aws.Operation{
		Name:       opGetPipeline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPipelineInput{}
	}

	req := c.newRequest(op, input, &GetPipelineOutput{})
	return GetPipelineRequest{Request: req, Input: input, Copy: c.GetPipelineRequest}
}

// GetPipelineRequest is the request type for the
// GetPipeline API operation.
type GetPipelineRequest struct {
	*aws.Request
	Input *GetPipelineInput
	Copy  func(*GetPipelineInput) GetPipelineRequest
}

// Send marshals and sends the GetPipeline API request.
func (r GetPipelineRequest) Send(ctx context.Context) (*GetPipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPipelineResponse{
		GetPipelineOutput: r.Request.Data.(*GetPipelineOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPipelineResponse is the response type for the
// GetPipeline API operation.
type GetPipelineResponse struct {
	*GetPipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPipeline request.
func (r *GetPipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
