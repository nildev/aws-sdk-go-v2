// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetCommandInvocationRequest
type GetCommandInvocationInput struct {
	_ struct{} `type:"structure"`

	// (Required) The parent command ID of the invocation plugin.
	//
	// CommandId is a required field
	CommandId *string `min:"36" type:"string" required:"true"`

	// (Required) The ID of the managed instance targeted by the command. A managed
	// instance can be an Amazon EC2 instance or an instance in your hybrid environment
	// that is configured for Systems Manager.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// (Optional) The name of the plugin for which you want detailed results. If
	// the document contains only one plugin, the name can be omitted and the details
	// will be returned.
	PluginName *string `min:"4" type:"string"`
}

// String returns the string representation
func (s GetCommandInvocationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCommandInvocationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCommandInvocationInput"}

	if s.CommandId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CommandId"))
	}
	if s.CommandId != nil && len(*s.CommandId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("CommandId", 36))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.PluginName != nil && len(*s.PluginName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("PluginName", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetCommandInvocationResult
type GetCommandInvocationOutput struct {
	_ struct{} `type:"structure"`

	// CloudWatch Logs information where Systems Manager sent the command output.
	CloudWatchOutputConfig *CloudWatchOutputConfig `json:"ssm:GetCommandInvocationOutput:CloudWatchOutputConfig" type:"structure"`

	// The parent command ID of the invocation plugin.
	CommandId *string `json:"ssm:GetCommandInvocationOutput:CommandId" min:"36" type:"string"`

	// The comment text for the command.
	Comment *string `json:"ssm:GetCommandInvocationOutput:Comment" type:"string"`

	// The name of the document that was run. For example, AWS-RunShellScript.
	DocumentName *string `json:"ssm:GetCommandInvocationOutput:DocumentName" type:"string"`

	// The SSM document version used in the request.
	DocumentVersion *string `json:"ssm:GetCommandInvocationOutput:DocumentVersion" type:"string"`

	// Duration since ExecutionStartDateTime.
	ExecutionElapsedTime *string `json:"ssm:GetCommandInvocationOutput:ExecutionElapsedTime" type:"string"`

	// The date and time the plugin was finished running. Date and time are written
	// in ISO 8601 format. For example, June 7, 2017 is represented as 2017-06-7.
	// The following sample AWS CLI command uses the InvokedAfter filter.
	//
	// aws ssm list-commands --filters key=InvokedAfter,value=2017-06-07T00:00:00Z
	//
	// If the plugin has not started to run, the string is empty.
	ExecutionEndDateTime *string `json:"ssm:GetCommandInvocationOutput:ExecutionEndDateTime" type:"string"`

	// The date and time the plugin started running. Date and time are written in
	// ISO 8601 format. For example, June 7, 2017 is represented as 2017-06-7. The
	// following sample AWS CLI command uses the InvokedBefore filter.
	//
	// aws ssm list-commands --filters key=InvokedBefore,value=2017-06-07T00:00:00Z
	//
	// If the plugin has not started to run, the string is empty.
	ExecutionStartDateTime *string `json:"ssm:GetCommandInvocationOutput:ExecutionStartDateTime" type:"string"`

	// The ID of the managed instance targeted by the command. A managed instance
	// can be an Amazon EC2 instance or an instance in your hybrid environment that
	// is configured for Systems Manager.
	InstanceId *string `json:"ssm:GetCommandInvocationOutput:InstanceId" type:"string"`

	// The name of the plugin for which you want detailed results. For example,
	// aws:RunShellScript is a plugin.
	PluginName *string `json:"ssm:GetCommandInvocationOutput:PluginName" min:"4" type:"string"`

	// The error level response code for the plugin script. If the response code
	// is -1, then the command has not started running on the instance, or it was
	// not received by the instance.
	ResponseCode *int64 `json:"ssm:GetCommandInvocationOutput:ResponseCode" type:"integer"`

	// The first 8,000 characters written by the plugin to stderr. If the command
	// has not finished running, then this string is empty.
	StandardErrorContent *string `json:"ssm:GetCommandInvocationOutput:StandardErrorContent" type:"string"`

	// The URL for the complete text written by the plugin to stderr. If the command
	// has not finished running, then this string is empty.
	StandardErrorUrl *string `json:"ssm:GetCommandInvocationOutput:StandardErrorUrl" type:"string"`

	// The first 24,000 characters written by the plugin to stdout. If the command
	// has not finished running, if ExecutionStatus is neither Succeeded nor Failed,
	// then this string is empty.
	StandardOutputContent *string `json:"ssm:GetCommandInvocationOutput:StandardOutputContent" type:"string"`

	// The URL for the complete text written by the plugin to stdout in Amazon S3.
	// If an Amazon S3 bucket was not specified, then this string is empty.
	StandardOutputUrl *string `json:"ssm:GetCommandInvocationOutput:StandardOutputUrl" type:"string"`

	// The status of this invocation plugin. This status can be different than StatusDetails.
	Status CommandInvocationStatus `json:"ssm:GetCommandInvocationOutput:Status" type:"string" enum:"true"`

	// A detailed status of the command execution for an invocation. StatusDetails
	// includes more information than Status because it includes states resulting
	// from error and concurrency control parameters. StatusDetails can show different
	// results than Status. For more information about these statuses, see Understanding
	// Command Statuses (http://docs.aws.amazon.com/systems-manager/latest/userguide/monitor-commands.html)
	// in the AWS Systems Manager User Guide. StatusDetails can be one of the following
	// values:
	//
	//    * Pending: The command has not been sent to the instance.
	//
	//    * In Progress: The command has been sent to the instance but has not reached
	//    a terminal state.
	//
	//    * Delayed: The system attempted to send the command to the target, but
	//    the target was not available. The instance might not be available because
	//    of network issues, the instance was stopped, etc. The system will try
	//    to deliver the command again.
	//
	//    * Success: The command or plugin was run successfully. This is a terminal
	//    state.
	//
	//    * Delivery Timed Out: The command was not delivered to the instance before
	//    the delivery timeout expired. Delivery timeouts do not count against the
	//    parent command's MaxErrors limit, but they do contribute to whether the
	//    parent command status is Success or Incomplete. This is a terminal state.
	//
	//    * Execution Timed Out: The command started to run on the instance, but
	//    the execution was not complete before the timeout expired. Execution timeouts
	//    count against the MaxErrors limit of the parent command. This is a terminal
	//    state.
	//
	//    * Failed: The command wasn't run successfully on the instance. For a plugin,
	//    this indicates that the result code was not zero. For a command invocation,
	//    this indicates that the result code for one or more plugins was not zero.
	//    Invocation failures count against the MaxErrors limit of the parent command.
	//    This is a terminal state.
	//
	//    * Canceled: The command was terminated before it was completed. This is
	//    a terminal state.
	//
	//    * Undeliverable: The command can't be delivered to the instance. The instance
	//    might not exist or might not be responding. Undeliverable invocations
	//    don't count against the parent command's MaxErrors limit and don't contribute
	//    to whether the parent command status is Success or Incomplete. This is
	//    a terminal state.
	//
	//    * Terminated: The parent command exceeded its MaxErrors limit and subsequent
	//    command invocations were canceled by the system. This is a terminal state.
	StatusDetails *string `json:"ssm:GetCommandInvocationOutput:StatusDetails" type:"string"`
}

// String returns the string representation
func (s GetCommandInvocationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCommandInvocation = "GetCommandInvocation"

// GetCommandInvocationRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Returns detailed information about command execution for an invocation or
// plugin.
//
//    // Example sending a request using GetCommandInvocationRequest.
//    req := client.GetCommandInvocationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetCommandInvocation
func (c *Client) GetCommandInvocationRequest(input *GetCommandInvocationInput) GetCommandInvocationRequest {
	op := &aws.Operation{
		Name:       opGetCommandInvocation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCommandInvocationInput{}
	}

	req := c.newRequest(op, input, &GetCommandInvocationOutput{})
	return GetCommandInvocationRequest{Request: req, Input: input, Copy: c.GetCommandInvocationRequest}
}

// GetCommandInvocationRequest is the request type for the
// GetCommandInvocation API operation.
type GetCommandInvocationRequest struct {
	*aws.Request
	Input *GetCommandInvocationInput
	Copy  func(*GetCommandInvocationInput) GetCommandInvocationRequest
}

// Send marshals and sends the GetCommandInvocation API request.
func (r GetCommandInvocationRequest) Send(ctx context.Context) (*GetCommandInvocationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCommandInvocationResponse{
		GetCommandInvocationOutput: r.Request.Data.(*GetCommandInvocationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCommandInvocationResponse is the response type for the
// GetCommandInvocation API operation.
type GetCommandInvocationResponse struct {
	*GetCommandInvocationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCommandInvocation request.
func (r *GetCommandInvocationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
