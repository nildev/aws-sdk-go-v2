// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RecordActivityTaskHeartbeatInput struct {
	_ struct{} `type:"structure"`

	// If specified, contains details about the progress of the task.
	Details *string `locationName:"details" type:"string"`

	// The taskToken of the ActivityTask.
	//
	// taskToken is generated by the service and should be treated as an opaque
	// value. If the task is passed to another process, its taskToken must also
	// be passed. This enables it to provide its progress and respond with results.
	//
	// TaskToken is a required field
	TaskToken *string `locationName:"taskToken" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RecordActivityTaskHeartbeatInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RecordActivityTaskHeartbeatInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RecordActivityTaskHeartbeatInput"}

	if s.TaskToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskToken"))
	}
	if s.TaskToken != nil && len(*s.TaskToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Status information about an activity task.
type RecordActivityTaskHeartbeatOutput struct {
	_ struct{} `type:"structure"`

	// Set to true if cancellation of the task is requested.
	//
	// CancelRequested is a required field
	CancelRequested *bool `json:"swf:RecordActivityTaskHeartbeatOutput:CancelRequested" locationName:"cancelRequested" type:"boolean" required:"true"`
}

// String returns the string representation
func (s RecordActivityTaskHeartbeatOutput) String() string {
	return awsutil.Prettify(s)
}

const opRecordActivityTaskHeartbeat = "RecordActivityTaskHeartbeat"

// RecordActivityTaskHeartbeatRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Used by activity workers to report to the service that the ActivityTask represented
// by the specified taskToken is still making progress. The worker can also
// specify details of the progress, for example percent complete, using the
// details parameter. This action can also be used by the worker as a mechanism
// to check if cancellation is being requested for the activity task. If a cancellation
// is being attempted for the specified task, then the boolean cancelRequested
// flag returned by the service is set to true.
//
// This action resets the taskHeartbeatTimeout clock. The taskHeartbeatTimeout
// is specified in RegisterActivityType.
//
// This action doesn't in itself create an event in the workflow execution history.
// However, if the task times out, the workflow execution history contains a
// ActivityTaskTimedOut event that contains the information from the last heartbeat
// generated by the activity worker.
//
// The taskStartToCloseTimeout of an activity type is the maximum duration of
// an activity task, regardless of the number of RecordActivityTaskHeartbeat
// requests received. The taskStartToCloseTimeout is also specified in RegisterActivityType.
//
// This operation is only useful for long-lived activities to report liveliness
// of the task and to determine if a cancellation is being attempted.
//
// If the cancelRequested flag returns true, a cancellation is being attempted.
// If the worker can cancel the activity, it should respond with RespondActivityTaskCanceled.
// Otherwise, it should ignore the cancellation request.
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
//    * You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using RecordActivityTaskHeartbeatRequest.
//    req := client.RecordActivityTaskHeartbeatRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RecordActivityTaskHeartbeatRequest(input *RecordActivityTaskHeartbeatInput) RecordActivityTaskHeartbeatRequest {
	op := &aws.Operation{
		Name:       opRecordActivityTaskHeartbeat,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RecordActivityTaskHeartbeatInput{}
	}

	req := c.newRequest(op, input, &RecordActivityTaskHeartbeatOutput{})
	return RecordActivityTaskHeartbeatRequest{Request: req, Input: input, Copy: c.RecordActivityTaskHeartbeatRequest}
}

// RecordActivityTaskHeartbeatRequest is the request type for the
// RecordActivityTaskHeartbeat API operation.
type RecordActivityTaskHeartbeatRequest struct {
	*aws.Request
	Input *RecordActivityTaskHeartbeatInput
	Copy  func(*RecordActivityTaskHeartbeatInput) RecordActivityTaskHeartbeatRequest
}

// Send marshals and sends the RecordActivityTaskHeartbeat API request.
func (r RecordActivityTaskHeartbeatRequest) Send(ctx context.Context) (*RecordActivityTaskHeartbeatResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RecordActivityTaskHeartbeatResponse{
		RecordActivityTaskHeartbeatOutput: r.Request.Data.(*RecordActivityTaskHeartbeatOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RecordActivityTaskHeartbeatResponse is the response type for the
// RecordActivityTaskHeartbeat API operation.
type RecordActivityTaskHeartbeatResponse struct {
	*RecordActivityTaskHeartbeatOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RecordActivityTaskHeartbeat request.
func (r *RecordActivityTaskHeartbeatResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
