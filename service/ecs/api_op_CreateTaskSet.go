// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/CreateTaskSetRequest
type CreateTaskSetInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. Up to 32 ASCII characters are allowed.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the service to create the task set in.
	//
	// Cluster is a required field
	Cluster *string `locationName:"cluster" type:"string" required:"true"`

	// An optional non-unique tag that identifies this task set in external systems.
	// If the task set is associated with a service discovery registry, the tasks
	// in this task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute
	// set to the provided value.
	ExternalId *string `locationName:"externalId" type:"string"`

	// The launch type that new tasks in the task set will use. For more information,
	// see Amazon ECS Launch Types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html)
	// in the Amazon Elastic Container Service Developer Guide.
	LaunchType LaunchType `locationName:"launchType" type:"string" enum:"true"`

	// A load balancer object representing the load balancer to use with the task
	// set. The supported load balancer types are either an Application Load Balancer
	// or a Network Load Balancer.
	LoadBalancers []LoadBalancer `locationName:"loadBalancers" type:"list"`

	// An object representing the network configuration for a task or service.
	NetworkConfiguration *NetworkConfiguration `locationName:"networkConfiguration" type:"structure"`

	// The platform version that the tasks in the task set should use. A platform
	// version is specified only for tasks using the Fargate launch type. If one
	// isn't specified, the LATEST platform version is used by default.
	PlatformVersion *string `locationName:"platformVersion" type:"string"`

	// A floating-point percentage of the desired number of tasks to place and keep
	// running in the task set.
	Scale *Scale `locationName:"scale" type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the service to create
	// the task set in.
	//
	// Service is a required field
	Service *string `locationName:"service" type:"string" required:"true"`

	// The details of the service discovery registries to assign to this task set.
	// For more information, see Service Discovery (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html).
	ServiceRegistries []ServiceRegistry `locationName:"serviceRegistries" type:"list"`

	// The task definition for the tasks in the task set to use.
	//
	// TaskDefinition is a required field
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTaskSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTaskSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTaskSetInput"}

	if s.Cluster == nil {
		invalidParams.Add(aws.NewErrParamRequired("Cluster"))
	}

	if s.Service == nil {
		invalidParams.Add(aws.NewErrParamRequired("Service"))
	}

	if s.TaskDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskDefinition"))
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			invalidParams.AddNested("NetworkConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/CreateTaskSetResponse
type CreateTaskSetOutput struct {
	_ struct{} `type:"structure"`

	// Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or
	// an EXTERNAL deployment. An Amazon ECS task set includes details such as the
	// desired number of tasks, how many tasks are running, and whether the task
	// set serves production traffic.
	TaskSet *TaskSet `json:"ecs:CreateTaskSetOutput:TaskSet" locationName:"taskSet" type:"structure"`
}

// String returns the string representation
func (s CreateTaskSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTaskSet = "CreateTaskSet"

// CreateTaskSetRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Create a task set in the specified cluster and service. This is used when
// a service uses the EXTERNAL deployment controller type. For more information,
// see Amazon ECS Deployment Types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
//
//    // Example sending a request using CreateTaskSetRequest.
//    req := client.CreateTaskSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/CreateTaskSet
func (c *Client) CreateTaskSetRequest(input *CreateTaskSetInput) CreateTaskSetRequest {
	op := &aws.Operation{
		Name:       opCreateTaskSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTaskSetInput{}
	}

	req := c.newRequest(op, input, &CreateTaskSetOutput{})
	return CreateTaskSetRequest{Request: req, Input: input, Copy: c.CreateTaskSetRequest}
}

// CreateTaskSetRequest is the request type for the
// CreateTaskSet API operation.
type CreateTaskSetRequest struct {
	*aws.Request
	Input *CreateTaskSetInput
	Copy  func(*CreateTaskSetInput) CreateTaskSetRequest
}

// Send marshals and sends the CreateTaskSet API request.
func (r CreateTaskSetRequest) Send(ctx context.Context) (*CreateTaskSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTaskSetResponse{
		CreateTaskSetOutput: r.Request.Data.(*CreateTaskSetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTaskSetResponse is the response type for the
// CreateTaskSet API operation.
type CreateTaskSetResponse struct {
	*CreateTaskSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTaskSet request.
func (r *CreateTaskSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
