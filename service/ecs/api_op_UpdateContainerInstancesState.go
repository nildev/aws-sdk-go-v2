// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/UpdateContainerInstancesStateRequest
type UpdateContainerInstancesStateInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instance to update. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// A list of container instance IDs or full ARN entries.
	//
	// ContainerInstances is a required field
	ContainerInstances []string `locationName:"containerInstances" type:"list" required:"true"`

	// The container instance state with which to update the container instance.
	// The only valid values for this action are ACTIVE and DRAINING. A container
	// instance can only be updated to DRAINING status once it has reached an ACTIVE
	// state. If a container instance is in REGISTERING, DEREGISTERING, or REGISTRATION_FAILED
	// state you can describe the container instance but will be unable to update
	// the container instance state.
	//
	// Status is a required field
	Status ContainerInstanceStatus `locationName:"status" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s UpdateContainerInstancesStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateContainerInstancesStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateContainerInstancesStateInput"}

	if s.ContainerInstances == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerInstances"))
	}
	if len(s.Status) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Status"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/UpdateContainerInstancesStateResponse
type UpdateContainerInstancesStateOutput struct {
	_ struct{} `type:"structure"`

	// The list of container instances.
	ContainerInstances []ContainerInstance `locationName:"containerInstances" type:"list"`

	// Any failures associated with the call.
	Failures []Failure `locationName:"failures" type:"list"`
}

// String returns the string representation
func (s UpdateContainerInstancesStateOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateContainerInstancesState = "UpdateContainerInstancesState"

// UpdateContainerInstancesStateRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Modifies the status of an Amazon ECS container instance.
//
// Once a container instance has reached an ACTIVE state, you can change the
// status of a container instance to DRAINING to manually remove an instance
// from a cluster, for example to perform system updates, update the Docker
// daemon, or scale down the cluster size.
//
// A container instance cannot be changed to DRAINING until it has reached an
// ACTIVE status. If the instance is in any other status, an error will be received.
//
// When you set a container instance to DRAINING, Amazon ECS prevents new tasks
// from being scheduled for placement on the container instance and replacement
// service tasks are started on other container instances in the cluster if
// the resources are available. Service tasks on the container instance that
// are in the PENDING state are stopped immediately.
//
// Service tasks on the container instance that are in the RUNNING state are
// stopped and replaced according to the service's deployment configuration
// parameters, minimumHealthyPercent and maximumPercent. You can change the
// deployment configuration of your service using UpdateService.
//
//    * If minimumHealthyPercent is below 100%, the scheduler can ignore desiredCount
//    temporarily during task replacement. For example, desiredCount is four
//    tasks, a minimum of 50% allows the scheduler to stop two existing tasks
//    before starting two new tasks. If the minimum is 100%, the service scheduler
//    can't remove existing tasks until the replacement tasks are considered
//    healthy. Tasks for services that do not use a load balancer are considered
//    healthy if they are in the RUNNING state. Tasks for services that use
//    a load balancer are considered healthy if they are in the RUNNING state
//    and the container instance they are hosted on is reported as healthy by
//    the load balancer.
//
//    * The maximumPercent parameter represents an upper limit on the number
//    of running tasks during task replacement, which enables you to define
//    the replacement batch size. For example, if desiredCount is four tasks,
//    a maximum of 200% starts four new tasks before stopping the four tasks
//    to be drained, provided that the cluster resources required to do this
//    are available. If the maximum is 100%, then replacement tasks can't start
//    until the draining tasks have stopped.
//
// Any PENDING or RUNNING tasks that do not belong to a service are not affected.
// You must wait for them to finish or stop them manually.
//
// A container instance has completed draining when it has no more RUNNING tasks.
// You can verify this using ListTasks.
//
// When a container instance has been drained, you can set a container instance
// to ACTIVE status and once it has reached that status the Amazon ECS scheduler
// can begin scheduling tasks on the instance again.
//
//    // Example sending a request using UpdateContainerInstancesStateRequest.
//    req := client.UpdateContainerInstancesStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/UpdateContainerInstancesState
func (c *Client) UpdateContainerInstancesStateRequest(input *UpdateContainerInstancesStateInput) UpdateContainerInstancesStateRequest {
	op := &aws.Operation{
		Name:       opUpdateContainerInstancesState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateContainerInstancesStateInput{}
	}

	req := c.newRequest(op, input, &UpdateContainerInstancesStateOutput{})
	return UpdateContainerInstancesStateRequest{Request: req, Input: input, Copy: c.UpdateContainerInstancesStateRequest}
}

// UpdateContainerInstancesStateRequest is the request type for the
// UpdateContainerInstancesState API operation.
type UpdateContainerInstancesStateRequest struct {
	*aws.Request
	Input *UpdateContainerInstancesStateInput
	Copy  func(*UpdateContainerInstancesStateInput) UpdateContainerInstancesStateRequest
}

// Send marshals and sends the UpdateContainerInstancesState API request.
func (r UpdateContainerInstancesStateRequest) Send(ctx context.Context) (*UpdateContainerInstancesStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateContainerInstancesStateResponse{
		UpdateContainerInstancesStateOutput: r.Request.Data.(*UpdateContainerInstancesStateOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateContainerInstancesStateResponse is the response type for the
// UpdateContainerInstancesState API operation.
type UpdateContainerInstancesStateResponse struct {
	*UpdateContainerInstancesStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateContainerInstancesState request.
func (r *UpdateContainerInstancesStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
