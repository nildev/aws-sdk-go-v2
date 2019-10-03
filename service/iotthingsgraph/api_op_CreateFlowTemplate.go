// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/CreateFlowTemplateRequest
type CreateFlowTemplateInput struct {
	_ struct{} `type:"structure"`

	// The namespace version in which the workflow is to be created.
	//
	// If no value is specified, the latest version is used by default.
	CompatibleNamespaceVersion *int64 `locationName:"compatibleNamespaceVersion" type:"long"`

	// The workflow DefinitionDocument.
	//
	// Definition is a required field
	Definition *DefinitionDocument `locationName:"definition" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateFlowTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFlowTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFlowTemplateInput"}

	if s.Definition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Definition"))
	}
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			invalidParams.AddNested("Definition", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/CreateFlowTemplateResponse
type CreateFlowTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The summary object that describes the created workflow.
	Summary *FlowTemplateSummary `json:"iotthingsgraph:CreateFlowTemplateOutput:Summary" locationName:"summary" type:"structure"`
}

// String returns the string representation
func (s CreateFlowTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateFlowTemplate = "CreateFlowTemplate"

// CreateFlowTemplateRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Creates a workflow template. Workflows can be created only in the user's
// namespace. (The public namespace contains only entities.) The workflow can
// contain only entities in the specified namespace. The workflow is validated
// against the entities in the latest version of the user's namespace unless
// another namespace version is specified in the request.
//
//    // Example sending a request using CreateFlowTemplateRequest.
//    req := client.CreateFlowTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/CreateFlowTemplate
func (c *Client) CreateFlowTemplateRequest(input *CreateFlowTemplateInput) CreateFlowTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateFlowTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFlowTemplateInput{}
	}

	req := c.newRequest(op, input, &CreateFlowTemplateOutput{})
	return CreateFlowTemplateRequest{Request: req, Input: input, Copy: c.CreateFlowTemplateRequest}
}

// CreateFlowTemplateRequest is the request type for the
// CreateFlowTemplate API operation.
type CreateFlowTemplateRequest struct {
	*aws.Request
	Input *CreateFlowTemplateInput
	Copy  func(*CreateFlowTemplateInput) CreateFlowTemplateRequest
}

// Send marshals and sends the CreateFlowTemplate API request.
func (r CreateFlowTemplateRequest) Send(ctx context.Context) (*CreateFlowTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFlowTemplateResponse{
		CreateFlowTemplateOutput: r.Request.Data.(*CreateFlowTemplateOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFlowTemplateResponse is the response type for the
// CreateFlowTemplate API operation.
type CreateFlowTemplateResponse struct {
	*CreateFlowTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFlowTemplate request.
func (r *CreateFlowTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
