// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datapipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for GetPipelineDefinition.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/GetPipelineDefinitionInput
type GetPipelineDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the pipeline.
	//
	// PipelineId is a required field
	PipelineId *string `locationName:"pipelineId" min:"1" type:"string" required:"true"`

	// The version of the pipeline definition to retrieve. Set this parameter to
	// latest (default) to use the last definition saved to the pipeline or active
	// to use the last definition that was activated.
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s GetPipelineDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPipelineDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPipelineDefinitionInput"}

	if s.PipelineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineId"))
	}
	if s.PipelineId != nil && len(*s.PipelineId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of GetPipelineDefinition.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/GetPipelineDefinitionOutput
type GetPipelineDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The parameter objects used in the pipeline definition.
	ParameterObjects []ParameterObject `json:"datapipeline:GetPipelineDefinitionOutput:ParameterObjects" locationName:"parameterObjects" type:"list"`

	// The parameter values used in the pipeline definition.
	ParameterValues []ParameterValue `json:"datapipeline:GetPipelineDefinitionOutput:ParameterValues" locationName:"parameterValues" type:"list"`

	// The objects defined in the pipeline.
	PipelineObjects []PipelineObject `json:"datapipeline:GetPipelineDefinitionOutput:PipelineObjects" locationName:"pipelineObjects" type:"list"`
}

// String returns the string representation
func (s GetPipelineDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPipelineDefinition = "GetPipelineDefinition"

// GetPipelineDefinitionRequest returns a request value for making API operation for
// AWS Data Pipeline.
//
// Gets the definition of the specified pipeline. You can call GetPipelineDefinition
// to retrieve the pipeline definition that you provided using PutPipelineDefinition.
//
//    // Example sending a request using GetPipelineDefinitionRequest.
//    req := client.GetPipelineDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/GetPipelineDefinition
func (c *Client) GetPipelineDefinitionRequest(input *GetPipelineDefinitionInput) GetPipelineDefinitionRequest {
	op := &aws.Operation{
		Name:       opGetPipelineDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPipelineDefinitionInput{}
	}

	req := c.newRequest(op, input, &GetPipelineDefinitionOutput{})
	return GetPipelineDefinitionRequest{Request: req, Input: input, Copy: c.GetPipelineDefinitionRequest}
}

// GetPipelineDefinitionRequest is the request type for the
// GetPipelineDefinition API operation.
type GetPipelineDefinitionRequest struct {
	*aws.Request
	Input *GetPipelineDefinitionInput
	Copy  func(*GetPipelineDefinitionInput) GetPipelineDefinitionRequest
}

// Send marshals and sends the GetPipelineDefinition API request.
func (r GetPipelineDefinitionRequest) Send(ctx context.Context) (*GetPipelineDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPipelineDefinitionResponse{
		GetPipelineDefinitionOutput: r.Request.Data.(*GetPipelineDefinitionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPipelineDefinitionResponse is the response type for the
// GetPipelineDefinition API operation.
type GetPipelineDefinitionResponse struct {
	*GetPipelineDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPipelineDefinition request.
func (r *GetPipelineDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
