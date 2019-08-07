// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a PutActionRevision action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/PutActionRevisionInput
type PutActionRevisionInput struct {
	_ struct{} `type:"structure"`

	// The name of the action that will process the revision.
	//
	// ActionName is a required field
	ActionName *string `locationName:"actionName" min:"1" type:"string" required:"true"`

	// Represents information about the version (or revision) of an action.
	//
	// ActionRevision is a required field
	ActionRevision *ActionRevision `locationName:"actionRevision" type:"structure" required:"true"`

	// The name of the pipeline that will start processing the revision to the source.
	//
	// PipelineName is a required field
	PipelineName *string `locationName:"pipelineName" min:"1" type:"string" required:"true"`

	// The name of the stage that contains the action that will act upon the revision.
	//
	// StageName is a required field
	StageName *string `locationName:"stageName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutActionRevisionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutActionRevisionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutActionRevisionInput"}

	if s.ActionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionName"))
	}
	if s.ActionName != nil && len(*s.ActionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ActionName", 1))
	}

	if s.ActionRevision == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionRevision"))
	}

	if s.PipelineName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineName"))
	}
	if s.PipelineName != nil && len(*s.PipelineName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineName", 1))
	}

	if s.StageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StageName"))
	}
	if s.StageName != nil && len(*s.StageName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StageName", 1))
	}
	if s.ActionRevision != nil {
		if err := s.ActionRevision.Validate(); err != nil {
			invalidParams.AddNested("ActionRevision", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a PutActionRevision action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/PutActionRevisionOutput
type PutActionRevisionOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the artifact revision was previously used in an execution
	// of the specified pipeline.
	NewRevision *bool `json:"codepipeline:PutActionRevisionOutput:NewRevision" locationName:"newRevision" type:"boolean"`

	// The ID of the current workflow state of the pipeline.
	PipelineExecutionId *string `json:"codepipeline:PutActionRevisionOutput:PipelineExecutionId" locationName:"pipelineExecutionId" type:"string"`
}

// String returns the string representation
func (s PutActionRevisionOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutActionRevision = "PutActionRevision"

// PutActionRevisionRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Provides information to AWS CodePipeline about new revisions to a source.
//
//    // Example sending a request using PutActionRevisionRequest.
//    req := client.PutActionRevisionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/PutActionRevision
func (c *Client) PutActionRevisionRequest(input *PutActionRevisionInput) PutActionRevisionRequest {
	op := &aws.Operation{
		Name:       opPutActionRevision,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutActionRevisionInput{}
	}

	req := c.newRequest(op, input, &PutActionRevisionOutput{})
	return PutActionRevisionRequest{Request: req, Input: input, Copy: c.PutActionRevisionRequest}
}

// PutActionRevisionRequest is the request type for the
// PutActionRevision API operation.
type PutActionRevisionRequest struct {
	*aws.Request
	Input *PutActionRevisionInput
	Copy  func(*PutActionRevisionInput) PutActionRevisionRequest
}

// Send marshals and sends the PutActionRevision API request.
func (r PutActionRevisionRequest) Send(ctx context.Context) (*PutActionRevisionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutActionRevisionResponse{
		PutActionRevisionOutput: r.Request.Data.(*PutActionRevisionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutActionRevisionResponse is the response type for the
// PutActionRevision API operation.
type PutActionRevisionResponse struct {
	*PutActionRevisionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutActionRevision request.
func (r *PutActionRevisionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
