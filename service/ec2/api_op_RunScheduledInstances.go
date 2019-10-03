// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for RunScheduledInstances.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RunScheduledInstancesRequest
type RunScheduledInstancesInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that ensures the idempotency of the request.
	// For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The number of instances.
	//
	// Default: 1
	InstanceCount *int64 `type:"integer"`

	// The launch specification. You must match the instance type, Availability
	// Zone, network, and platform of the schedule that you purchased.
	//
	// LaunchSpecification is a required field
	LaunchSpecification *ScheduledInstancesLaunchSpecification `type:"structure" required:"true"`

	// The Scheduled Instance ID.
	//
	// ScheduledInstanceId is a required field
	ScheduledInstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RunScheduledInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RunScheduledInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RunScheduledInstancesInput"}

	if s.LaunchSpecification == nil {
		invalidParams.Add(aws.NewErrParamRequired("LaunchSpecification"))
	}

	if s.ScheduledInstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledInstanceId"))
	}
	if s.LaunchSpecification != nil {
		if err := s.LaunchSpecification.Validate(); err != nil {
			invalidParams.AddNested("LaunchSpecification", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of RunScheduledInstances.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RunScheduledInstancesResult
type RunScheduledInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the newly launched instances.
	InstanceIdSet []string `json:"ec2:RunScheduledInstancesOutput:InstanceIdSet" locationName:"instanceIdSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s RunScheduledInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opRunScheduledInstances = "RunScheduledInstances"

// RunScheduledInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Launches the specified Scheduled Instances.
//
// Before you can launch a Scheduled Instance, you must purchase it and obtain
// an identifier using PurchaseScheduledInstances.
//
// You must launch a Scheduled Instance during its scheduled time period. You
// can't stop or reboot a Scheduled Instance, but you can terminate it as needed.
// If you terminate a Scheduled Instance before the current scheduled time period
// ends, you can launch it again after a few minutes. For more information,
// see Scheduled Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-scheduled-instances.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using RunScheduledInstancesRequest.
//    req := client.RunScheduledInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RunScheduledInstances
func (c *Client) RunScheduledInstancesRequest(input *RunScheduledInstancesInput) RunScheduledInstancesRequest {
	op := &aws.Operation{
		Name:       opRunScheduledInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RunScheduledInstancesInput{}
	}

	req := c.newRequest(op, input, &RunScheduledInstancesOutput{})
	return RunScheduledInstancesRequest{Request: req, Input: input, Copy: c.RunScheduledInstancesRequest}
}

// RunScheduledInstancesRequest is the request type for the
// RunScheduledInstances API operation.
type RunScheduledInstancesRequest struct {
	*aws.Request
	Input *RunScheduledInstancesInput
	Copy  func(*RunScheduledInstancesInput) RunScheduledInstancesRequest
}

// Send marshals and sends the RunScheduledInstances API request.
func (r RunScheduledInstancesRequest) Send(ctx context.Context) (*RunScheduledInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RunScheduledInstancesResponse{
		RunScheduledInstancesOutput: r.Request.Data.(*RunScheduledInstancesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RunScheduledInstancesResponse is the response type for the
// RunScheduledInstances API operation.
type RunScheduledInstancesResponse struct {
	*RunScheduledInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RunScheduledInstances request.
func (r *RunScheduledInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
