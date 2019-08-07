// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PollForActivityTaskInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that contains the task lists being polled.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// Identity of the worker making the request, recorded in the ActivityTaskStarted
	// event in the workflow history. This enables diagnostic tracing when problems
	// arise. The form of this identity is user defined.
	Identity *string `locationName:"identity" type:"string"`

	// Specifies the task list to poll for activity tasks.
	//
	// The specified string must not start or end with whitespace. It must not contain
	// a : (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f
	// | \u007f-\u009f). Also, it must not be the literal string arn.
	//
	// TaskList is a required field
	TaskList *TaskList `locationName:"taskList" type:"structure" required:"true"`
}

// String returns the string representation
func (s PollForActivityTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PollForActivityTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PollForActivityTaskInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.TaskList == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskList"))
	}
	if s.TaskList != nil {
		if err := s.TaskList.Validate(); err != nil {
			invalidParams.AddNested("TaskList", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Unit of work sent to an activity worker.
type PollForActivityTaskOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the task.
	//
	// ActivityId is a required field
	ActivityId *string `json:"swf:PollForActivityTaskOutput:ActivityId" locationName:"activityId" min:"1" type:"string" required:"true"`

	// The type of this activity task.
	//
	// ActivityType is a required field
	ActivityType *ActivityType `json:"swf:PollForActivityTaskOutput:ActivityType" locationName:"activityType" type:"structure" required:"true"`

	// The inputs provided when the activity task was scheduled. The form of the
	// input is user defined and should be meaningful to the activity implementation.
	Input *string `json:"swf:PollForActivityTaskOutput:Input" locationName:"input" type:"string"`

	// The ID of the ActivityTaskStarted event recorded in the history.
	//
	// StartedEventId is a required field
	StartedEventId *int64 `json:"swf:PollForActivityTaskOutput:StartedEventId" locationName:"startedEventId" type:"long" required:"true"`

	// The opaque string used as a handle on the task. This token is used by workers
	// to communicate progress and response information back to the system about
	// the task.
	//
	// TaskToken is a required field
	TaskToken *string `json:"swf:PollForActivityTaskOutput:TaskToken" locationName:"taskToken" min:"1" type:"string" required:"true"`

	// The workflow execution that started this activity task.
	//
	// WorkflowExecution is a required field
	WorkflowExecution *WorkflowExecution `json:"swf:PollForActivityTaskOutput:WorkflowExecution" locationName:"workflowExecution" type:"structure" required:"true"`
}

// String returns the string representation
func (s PollForActivityTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opPollForActivityTask = "PollForActivityTask"

// PollForActivityTaskRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Used by workers to get an ActivityTask from the specified activity taskList.
// This initiates a long poll, where the service holds the HTTP connection open
// and responds as soon as a task becomes available. The maximum time the service
// holds on to the request before responding is 60 seconds. If no task is available
// within 60 seconds, the poll returns an empty result. An empty result, in
// this context, means that an ActivityTask is returned, but that the value
// of taskToken is an empty string. If a task is returned, the worker should
// use its type to identify and process it correctly.
//
// Workers should set their client side socket timeout to at least 70 seconds
// (10 seconds higher than the maximum time service may hold the poll request).
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * Constrain the taskList.name parameter by using a Condition element with
//    the swf:taskList.name key to allow the action to access only certain task
//    lists.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using PollForActivityTaskRequest.
//    req := client.PollForActivityTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) PollForActivityTaskRequest(input *PollForActivityTaskInput) PollForActivityTaskRequest {
	op := &aws.Operation{
		Name:       opPollForActivityTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PollForActivityTaskInput{}
	}

	req := c.newRequest(op, input, &PollForActivityTaskOutput{})
	return PollForActivityTaskRequest{Request: req, Input: input, Copy: c.PollForActivityTaskRequest}
}

// PollForActivityTaskRequest is the request type for the
// PollForActivityTask API operation.
type PollForActivityTaskRequest struct {
	*aws.Request
	Input *PollForActivityTaskInput
	Copy  func(*PollForActivityTaskInput) PollForActivityTaskRequest
}

// Send marshals and sends the PollForActivityTask API request.
func (r PollForActivityTaskRequest) Send(ctx context.Context) (*PollForActivityTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PollForActivityTaskResponse{
		PollForActivityTaskOutput: r.Request.Data.(*PollForActivityTaskOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PollForActivityTaskResponse is the response type for the
// PollForActivityTask API operation.
type PollForActivityTaskResponse struct {
	*PollForActivityTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PollForActivityTask request.
func (r *PollForActivityTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
