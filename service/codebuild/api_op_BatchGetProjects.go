// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/BatchGetProjectsInput
type BatchGetProjectsInput struct {
	_ struct{} `type:"structure"`

	// The names of the build projects.
	//
	// Names is a required field
	Names []string `locationName:"names" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetProjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetProjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetProjectsInput"}

	if s.Names == nil {
		invalidParams.Add(aws.NewErrParamRequired("Names"))
	}
	if s.Names != nil && len(s.Names) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Names", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/BatchGetProjectsOutput
type BatchGetProjectsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the requested build projects.
	Projects []Project `json:"codebuild:BatchGetProjectsOutput:Projects" locationName:"projects" type:"list"`

	// The names of build projects for which information could not be found.
	ProjectsNotFound []string `json:"codebuild:BatchGetProjectsOutput:ProjectsNotFound" locationName:"projectsNotFound" min:"1" type:"list"`
}

// String returns the string representation
func (s BatchGetProjectsOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetProjects = "BatchGetProjects"

// BatchGetProjectsRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Gets information about build projects.
//
//    // Example sending a request using BatchGetProjectsRequest.
//    req := client.BatchGetProjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/BatchGetProjects
func (c *Client) BatchGetProjectsRequest(input *BatchGetProjectsInput) BatchGetProjectsRequest {
	op := &aws.Operation{
		Name:       opBatchGetProjects,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetProjectsInput{}
	}

	req := c.newRequest(op, input, &BatchGetProjectsOutput{})
	return BatchGetProjectsRequest{Request: req, Input: input, Copy: c.BatchGetProjectsRequest}
}

// BatchGetProjectsRequest is the request type for the
// BatchGetProjects API operation.
type BatchGetProjectsRequest struct {
	*aws.Request
	Input *BatchGetProjectsInput
	Copy  func(*BatchGetProjectsInput) BatchGetProjectsRequest
}

// Send marshals and sends the BatchGetProjects API request.
func (r BatchGetProjectsRequest) Send(ctx context.Context) (*BatchGetProjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetProjectsResponse{
		BatchGetProjectsOutput: r.Request.Data.(*BatchGetProjectsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetProjectsResponse is the response type for the
// BatchGetProjects API operation.
type BatchGetProjectsResponse struct {
	*BatchGetProjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetProjects request.
func (r *BatchGetProjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
