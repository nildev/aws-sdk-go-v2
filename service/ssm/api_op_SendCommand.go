// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/SendCommandRequest
type SendCommandInput struct {
	_ struct{} `type:"structure"`

	// Enables Systems Manager to send Run Command output to Amazon CloudWatch Logs.
	CloudWatchOutputConfig *CloudWatchOutputConfig `type:"structure"`

	// User-specified information about the command, such as a brief description
	// of what the command should do.
	Comment *string `type:"string"`

	// The Sha256 or Sha1 hash created by the system when the document was created.
	//
	// Sha1 hashes have been deprecated.
	DocumentHash *string `type:"string"`

	// Sha256 or Sha1.
	//
	// Sha1 hashes have been deprecated.
	DocumentHashType DocumentHashType `type:"string" enum:"true"`

	// Required. The name of the Systems Manager document to run. This can be a
	// public document or a custom document.
	//
	// DocumentName is a required field
	DocumentName *string `type:"string" required:"true"`

	// The SSM document version to use in the request. You can specify $DEFAULT,
	// $LATEST, or a specific version number. If you run commands by using the AWS
	// CLI, then you must escape the first two options by using a backslash. If
	// you specify a version number, then you don't need to use the backslash. For
	// example:
	//
	// --document-version "\$DEFAULT"
	//
	// --document-version "\$LATEST"
	//
	// --document-version "3"
	DocumentVersion *string `type:"string"`

	// The instance IDs where the command should run. You can specify a maximum
	// of 50 IDs. If you prefer not to list individual instance IDs, you can instead
	// send commands to a fleet of instances using the Targets parameter, which
	// accepts EC2 tags. For more information about how to use targets, see Sending
	// Commands to a Fleet (http://docs.aws.amazon.com/systems-manager/latest/userguide/send-commands-multiple.html)
	// in the AWS Systems Manager User Guide.
	InstanceIds []string `type:"list"`

	// (Optional) The maximum number of instances that are allowed to run the command
	// at the same time. You can specify a number such as 10 or a percentage such
	// as 10%. The default value is 50. For more information about how to use MaxConcurrency,
	// see Using Concurrency Controls (http://docs.aws.amazon.com/systems-manager/latest/userguide/send-commands-multiple.html#send-commands-velocity)
	// in the AWS Systems Manager User Guide.
	MaxConcurrency *string `min:"1" type:"string"`

	// The maximum number of errors allowed without the command failing. When the
	// command fails one more time beyond the value of MaxErrors, the systems stops
	// sending the command to additional targets. You can specify a number like
	// 10 or a percentage like 10%. The default value is 0. For more information
	// about how to use MaxErrors, see Using Error Controls (http://docs.aws.amazon.com/systems-manager/latest/userguide/send-commands-multiple.html#send-commands-maxerrors)
	// in the AWS Systems Manager User Guide.
	MaxErrors *string `min:"1" type:"string"`

	// Configurations for sending notifications.
	NotificationConfig *NotificationConfig `type:"structure"`

	// The name of the S3 bucket where command execution responses should be stored.
	OutputS3BucketName *string `min:"3" type:"string"`

	// The directory structure within the S3 bucket where the responses should be
	// stored.
	OutputS3KeyPrefix *string `type:"string"`

	// (Deprecated) You can no longer specify this parameter. The system ignores
	// it. Instead, Systems Manager automatically determines the Amazon S3 bucket
	// region.
	OutputS3Region *string `min:"3" type:"string"`

	// The required and optional parameters specified in the document being run.
	Parameters map[string][]string `type:"map"`

	// The ARN of the IAM service role to use to publish Amazon Simple Notification
	// Service (Amazon SNS) notifications for Run Command commands.
	ServiceRoleArn *string `type:"string"`

	// (Optional) An array of search criteria that targets instances using a Key,Value
	// combination that you specify. Targets is required if you don't provide one
	// or more instance IDs in the call. For more information about how to use targets,
	// see Sending Commands to a Fleet (http://docs.aws.amazon.com/systems-manager/latest/userguide/send-commands-multiple.html)
	// in the AWS Systems Manager User Guide.
	Targets []Target `type:"list"`

	// If this time is reached and the command has not already started running,
	// it will not run.
	TimeoutSeconds *int64 `min:"30" type:"integer"`
}

// String returns the string representation
func (s SendCommandInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendCommandInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendCommandInput"}

	if s.DocumentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentName"))
	}
	if s.MaxConcurrency != nil && len(*s.MaxConcurrency) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MaxConcurrency", 1))
	}
	if s.MaxErrors != nil && len(*s.MaxErrors) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MaxErrors", 1))
	}
	if s.OutputS3BucketName != nil && len(*s.OutputS3BucketName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("OutputS3BucketName", 3))
	}
	if s.OutputS3Region != nil && len(*s.OutputS3Region) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("OutputS3Region", 3))
	}
	if s.TimeoutSeconds != nil && *s.TimeoutSeconds < 30 {
		invalidParams.Add(aws.NewErrParamMinValue("TimeoutSeconds", 30))
	}
	if s.CloudWatchOutputConfig != nil {
		if err := s.CloudWatchOutputConfig.Validate(); err != nil {
			invalidParams.AddNested("CloudWatchOutputConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/SendCommandResult
type SendCommandOutput struct {
	_ struct{} `type:"structure"`

	// The request as it was received by Systems Manager. Also provides the command
	// ID which can be used future references to this request.
	Command *Command `json:"ssm:SendCommandOutput:Command" type:"structure"`
}

// String returns the string representation
func (s SendCommandOutput) String() string {
	return awsutil.Prettify(s)
}

const opSendCommand = "SendCommand"

// SendCommandRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Runs commands on one or more managed instances.
//
//    // Example sending a request using SendCommandRequest.
//    req := client.SendCommandRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/SendCommand
func (c *Client) SendCommandRequest(input *SendCommandInput) SendCommandRequest {
	op := &aws.Operation{
		Name:       opSendCommand,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendCommandInput{}
	}

	req := c.newRequest(op, input, &SendCommandOutput{})
	return SendCommandRequest{Request: req, Input: input, Copy: c.SendCommandRequest}
}

// SendCommandRequest is the request type for the
// SendCommand API operation.
type SendCommandRequest struct {
	*aws.Request
	Input *SendCommandInput
	Copy  func(*SendCommandInput) SendCommandRequest
}

// Send marshals and sends the SendCommand API request.
func (r SendCommandRequest) Send(ctx context.Context) (*SendCommandResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendCommandResponse{
		SendCommandOutput: r.Request.Data.(*SendCommandOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendCommandResponse is the response type for the
// SendCommand API operation.
type SendCommandResponse struct {
	*SendCommandOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendCommand request.
func (r *SendCommandResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
