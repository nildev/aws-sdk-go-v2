// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// StartTaskExecutionRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/StartTaskExecutionRequest
type StartTaskExecutionInput struct {
	_ struct{} `type:"structure"`

	// A filter that determines which files to include in the transfer during a
	// task execution based on the specified pattern in the filter. When multiple
	// include filters are set, they are interpreted as an OR.
	Includes []FilterRule `type:"list"`

	// Represents the options that are available to control the behavior of a StartTaskExecution
	// operation. Behavior includes preserving metadata such as user ID (UID), group
	// ID (GID), and file permissions, and also overwriting files in the destination,
	// data integrity verification, and so on.
	//
	// A task has a set of default options associated with it. If you don't specify
	// an option in StartTaskExecution, the default value is used. You can override
	// the defaults options on each task execution by specifying an overriding Options
	// value to StartTaskExecution.
	OverrideOptions *Options `type:"structure"`

	// The Amazon Resource Name (ARN) of the task to start.
	//
	// TaskArn is a required field
	TaskArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartTaskExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartTaskExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartTaskExecutionInput"}

	if s.TaskArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskArn"))
	}
	if s.OverrideOptions != nil {
		if err := s.OverrideOptions.Validate(); err != nil {
			invalidParams.AddNested("OverrideOptions", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// StartTaskExecutionResponse
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/StartTaskExecutionResponse
type StartTaskExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the specific task execution that was started.
	TaskExecutionArn *string `json:"datasync:StartTaskExecutionOutput:TaskExecutionArn" type:"string"`
}

// String returns the string representation
func (s StartTaskExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartTaskExecution = "StartTaskExecution"

// StartTaskExecutionRequest returns a request value for making API operation for
// AWS DataSync.
//
// Starts a specific invocation of a task. A TaskExecution value represents
// an individual run of a task. Each task can have at most one TaskExecution
// at a time.
//
// TaskExecution has the following transition phases: INITIALIZING | PREPARING
// | TRANSFERRING | VERIFYING | SUCCESS/FAILURE.
//
// For detailed information, see Task Execution in "https://docs.aws.amazon.com/datasync/latest/userguide/how-datasync-works.html#terminology"
// (Components and Terminology) in the AWS DataSync User Guide.
//
//    // Example sending a request using StartTaskExecutionRequest.
//    req := client.StartTaskExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/StartTaskExecution
func (c *Client) StartTaskExecutionRequest(input *StartTaskExecutionInput) StartTaskExecutionRequest {
	op := &aws.Operation{
		Name:       opStartTaskExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartTaskExecutionInput{}
	}

	req := c.newRequest(op, input, &StartTaskExecutionOutput{})
	return StartTaskExecutionRequest{Request: req, Input: input, Copy: c.StartTaskExecutionRequest}
}

// StartTaskExecutionRequest is the request type for the
// StartTaskExecution API operation.
type StartTaskExecutionRequest struct {
	*aws.Request
	Input *StartTaskExecutionInput
	Copy  func(*StartTaskExecutionInput) StartTaskExecutionRequest
}

// Send marshals and sends the StartTaskExecution API request.
func (r StartTaskExecutionRequest) Send(ctx context.Context) (*StartTaskExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartTaskExecutionResponse{
		StartTaskExecutionOutput: r.Request.Data.(*StartTaskExecutionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartTaskExecutionResponse is the response type for the
// StartTaskExecution API operation.
type StartTaskExecutionResponse struct {
	*StartTaskExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartTaskExecution request.
func (r *StartTaskExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
