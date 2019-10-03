// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/TerminateInstanceInAutoScalingGroupType
type TerminateInstanceInAutoScalingGroupInput struct {
	_ struct{} `type:"structure"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `min:"1" type:"string" required:"true"`

	// Indicates whether terminating the instance also decrements the size of the
	// Auto Scaling group.
	//
	// ShouldDecrementDesiredCapacity is a required field
	ShouldDecrementDesiredCapacity *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s TerminateInstanceInAutoScalingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminateInstanceInAutoScalingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminateInstanceInAutoScalingGroupInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if s.ShouldDecrementDesiredCapacity == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShouldDecrementDesiredCapacity"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/ActivityType
type TerminateInstanceInAutoScalingGroupOutput struct {
	_ struct{} `type:"structure"`

	// A scaling activity.
	Activity *Activity `json:"autoscaling:TerminateInstanceInAutoScalingGroupOutput:Activity" type:"structure"`
}

// String returns the string representation
func (s TerminateInstanceInAutoScalingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opTerminateInstanceInAutoScalingGroup = "TerminateInstanceInAutoScalingGroup"

// TerminateInstanceInAutoScalingGroupRequest returns a request value for making API operation for
// Auto Scaling.
//
// Terminates the specified instance and optionally adjusts the desired group
// size.
//
// This call simply makes a termination request. The instance is not terminated
// immediately.
//
//    // Example sending a request using TerminateInstanceInAutoScalingGroupRequest.
//    req := client.TerminateInstanceInAutoScalingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/TerminateInstanceInAutoScalingGroup
func (c *Client) TerminateInstanceInAutoScalingGroupRequest(input *TerminateInstanceInAutoScalingGroupInput) TerminateInstanceInAutoScalingGroupRequest {
	op := &aws.Operation{
		Name:       opTerminateInstanceInAutoScalingGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TerminateInstanceInAutoScalingGroupInput{}
	}

	req := c.newRequest(op, input, &TerminateInstanceInAutoScalingGroupOutput{})
	return TerminateInstanceInAutoScalingGroupRequest{Request: req, Input: input, Copy: c.TerminateInstanceInAutoScalingGroupRequest}
}

// TerminateInstanceInAutoScalingGroupRequest is the request type for the
// TerminateInstanceInAutoScalingGroup API operation.
type TerminateInstanceInAutoScalingGroupRequest struct {
	*aws.Request
	Input *TerminateInstanceInAutoScalingGroupInput
	Copy  func(*TerminateInstanceInAutoScalingGroupInput) TerminateInstanceInAutoScalingGroupRequest
}

// Send marshals and sends the TerminateInstanceInAutoScalingGroup API request.
func (r TerminateInstanceInAutoScalingGroupRequest) Send(ctx context.Context) (*TerminateInstanceInAutoScalingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TerminateInstanceInAutoScalingGroupResponse{
		TerminateInstanceInAutoScalingGroupOutput: r.Request.Data.(*TerminateInstanceInAutoScalingGroupOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TerminateInstanceInAutoScalingGroupResponse is the response type for the
// TerminateInstanceInAutoScalingGroup API operation.
type TerminateInstanceInAutoScalingGroupResponse struct {
	*TerminateInstanceInAutoScalingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TerminateInstanceInAutoScalingGroup request.
func (r *TerminateInstanceInAutoScalingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
