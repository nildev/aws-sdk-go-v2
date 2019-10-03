// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datapipeline

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for ValidatePipelineDefinition.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ValidatePipelineDefinitionInput
type ValidatePipelineDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The parameter objects used with the pipeline.
	ParameterObjects []ParameterObject `locationName:"parameterObjects" type:"list"`

	// The parameter values used with the pipeline.
	ParameterValues []ParameterValue `locationName:"parameterValues" type:"list"`

	// The ID of the pipeline.
	//
	// PipelineId is a required field
	PipelineId *string `locationName:"pipelineId" min:"1" type:"string" required:"true"`

	// The objects that define the pipeline changes to validate against the pipeline.
	//
	// PipelineObjects is a required field
	PipelineObjects []PipelineObject `locationName:"pipelineObjects" type:"list" required:"true"`
}

// String returns the string representation
func (s ValidatePipelineDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ValidatePipelineDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ValidatePipelineDefinitionInput"}

	if s.PipelineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineId"))
	}
	if s.PipelineId != nil && len(*s.PipelineId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineId", 1))
	}

	if s.PipelineObjects == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineObjects"))
	}
	if s.ParameterObjects != nil {
		for i, v := range s.ParameterObjects {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ParameterObjects", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ParameterValues != nil {
		for i, v := range s.ParameterValues {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ParameterValues", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.PipelineObjects != nil {
		for i, v := range s.PipelineObjects {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PipelineObjects", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of ValidatePipelineDefinition.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ValidatePipelineDefinitionOutput
type ValidatePipelineDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether there were validation errors.
	//
	// Errored is a required field
	Errored *bool `json:"datapipeline:ValidatePipelineDefinitionOutput:Errored" locationName:"errored" type:"boolean" required:"true"`

	// Any validation errors that were found.
	ValidationErrors []ValidationError `json:"datapipeline:ValidatePipelineDefinitionOutput:ValidationErrors" locationName:"validationErrors" type:"list"`

	// Any validation warnings that were found.
	ValidationWarnings []ValidationWarning `json:"datapipeline:ValidatePipelineDefinitionOutput:ValidationWarnings" locationName:"validationWarnings" type:"list"`
}

// String returns the string representation
func (s ValidatePipelineDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opValidatePipelineDefinition = "ValidatePipelineDefinition"

// ValidatePipelineDefinitionRequest returns a request value for making API operation for
// AWS Data Pipeline.
//
// Validates the specified pipeline definition to ensure that it is well formed
// and can be run without error.
//
//    // Example sending a request using ValidatePipelineDefinitionRequest.
//    req := client.ValidatePipelineDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ValidatePipelineDefinition
func (c *Client) ValidatePipelineDefinitionRequest(input *ValidatePipelineDefinitionInput) ValidatePipelineDefinitionRequest {
	op := &aws.Operation{
		Name:       opValidatePipelineDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ValidatePipelineDefinitionInput{}
	}

	req := c.newRequest(op, input, &ValidatePipelineDefinitionOutput{})
	return ValidatePipelineDefinitionRequest{Request: req, Input: input, Copy: c.ValidatePipelineDefinitionRequest}
}

// ValidatePipelineDefinitionRequest is the request type for the
// ValidatePipelineDefinition API operation.
type ValidatePipelineDefinitionRequest struct {
	*aws.Request
	Input *ValidatePipelineDefinitionInput
	Copy  func(*ValidatePipelineDefinitionInput) ValidatePipelineDefinitionRequest
}

// Send marshals and sends the ValidatePipelineDefinition API request.
func (r ValidatePipelineDefinitionRequest) Send(ctx context.Context) (*ValidatePipelineDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ValidatePipelineDefinitionResponse{
		ValidatePipelineDefinitionOutput: r.Request.Data.(*ValidatePipelineDefinitionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ValidatePipelineDefinitionResponse is the response type for the
// ValidatePipelineDefinition API operation.
type ValidatePipelineDefinitionResponse struct {
	*ValidatePipelineDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ValidatePipelineDefinition request.
func (r *ValidatePipelineDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
