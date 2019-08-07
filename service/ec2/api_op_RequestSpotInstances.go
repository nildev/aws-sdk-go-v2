// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for RequestSpotInstances.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RequestSpotInstancesRequest
type RequestSpotInstancesInput struct {
	_ struct{} `type:"structure"`

	// The user-specified name for a logical grouping of requests.
	//
	// When you specify an Availability Zone group in a Spot Instance request, all
	// Spot Instances in the request are launched in the same Availability Zone.
	// Instance proximity is maintained with this parameter, but the choice of Availability
	// Zone is not. The group applies only to requests for Spot Instances of the
	// same instance type. Any additional Spot Instance requests that are specified
	// with the same Availability Zone group name are launched in that same Availability
	// Zone, as long as at least one instance from the group is still active.
	//
	// If there is no active instance running in the Availability Zone group that
	// you specify for a new Spot Instance request (all instances are terminated,
	// the request is expired, or the maximum price you specified falls below current
	// Spot price), then Amazon EC2 launches the instance in any Availability Zone
	// where the constraint can be met. Consequently, the subsequent set of Spot
	// Instances could be placed in a different zone from the original request,
	// even if you specified the same Availability Zone group.
	//
	// Default: Instances are launched in any available Availability Zone.
	AvailabilityZoneGroup *string `locationName:"availabilityZoneGroup" type:"string"`

	// The required duration for the Spot Instances (also known as Spot blocks),
	// in minutes. This value must be a multiple of 60 (60, 120, 180, 240, 300,
	// or 360).
	//
	// The duration period starts as soon as your Spot Instance receives its instance
	// ID. At the end of the duration period, Amazon EC2 marks the Spot Instance
	// for termination and provides a Spot Instance termination notice, which gives
	// the instance a two-minute warning before it terminates.
	//
	// You can't specify an Availability Zone group or a launch group if you specify
	// a duration.
	BlockDurationMinutes *int64 `locationName:"blockDurationMinutes" type:"integer"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html)
	// in the Amazon EC2 User Guide for Linux Instances.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The maximum number of Spot Instances to launch.
	//
	// Default: 1
	InstanceCount *int64 `locationName:"instanceCount" type:"integer"`

	// The behavior when a Spot Instance is interrupted. The default is terminate.
	InstanceInterruptionBehavior InstanceInterruptionBehavior `type:"string" enum:"true"`

	// The instance launch group. Launch groups are Spot Instances that launch together
	// and terminate together.
	//
	// Default: Instances are launched and terminated individually
	LaunchGroup *string `locationName:"launchGroup" type:"string"`

	// The launch specification.
	LaunchSpecification *RequestSpotLaunchSpecification `type:"structure"`

	// The maximum price per hour that you are willing to pay for a Spot Instance.
	// The default is the On-Demand price.
	SpotPrice *string `locationName:"spotPrice" type:"string"`

	// The Spot Instance request type.
	//
	// Default: one-time
	Type SpotInstanceType `locationName:"type" type:"string" enum:"true"`

	// The start date of the request. If this is a one-time request, the request
	// becomes active at this date and time and remains active until all instances
	// launch, the request expires, or the request is canceled. If the request is
	// persistent, the request becomes active at this date and time and remains
	// active until it expires or is canceled.
	ValidFrom *time.Time `locationName:"validFrom" type:"timestamp" timestampFormat:"iso8601"`

	// The end date of the request. If this is a one-time request, the request remains
	// active until all instances launch, the request is canceled, or this date
	// is reached. If the request is persistent, it remains active until it is canceled
	// or this date is reached. The default end date is 7 days from the current
	// date.
	ValidUntil *time.Time `locationName:"validUntil" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s RequestSpotInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestSpotInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestSpotInstancesInput"}
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

// Contains the output of RequestSpotInstances.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RequestSpotInstancesResult
type RequestSpotInstancesOutput struct {
	_ struct{} `type:"structure"`

	// One or more Spot Instance requests.
	SpotInstanceRequests []SpotInstanceRequest `json:"ec2:RequestSpotInstancesOutput:SpotInstanceRequests" locationName:"spotInstanceRequestSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s RequestSpotInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opRequestSpotInstances = "RequestSpotInstances"

// RequestSpotInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a Spot Instance request.
//
// For more information, see Spot Instance Requests (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-requests.html)
// in the Amazon EC2 User Guide for Linux Instances.
//
//    // Example sending a request using RequestSpotInstancesRequest.
//    req := client.RequestSpotInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RequestSpotInstances
func (c *Client) RequestSpotInstancesRequest(input *RequestSpotInstancesInput) RequestSpotInstancesRequest {
	op := &aws.Operation{
		Name:       opRequestSpotInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RequestSpotInstancesInput{}
	}

	req := c.newRequest(op, input, &RequestSpotInstancesOutput{})
	return RequestSpotInstancesRequest{Request: req, Input: input, Copy: c.RequestSpotInstancesRequest}
}

// RequestSpotInstancesRequest is the request type for the
// RequestSpotInstances API operation.
type RequestSpotInstancesRequest struct {
	*aws.Request
	Input *RequestSpotInstancesInput
	Copy  func(*RequestSpotInstancesInput) RequestSpotInstancesRequest
}

// Send marshals and sends the RequestSpotInstances API request.
func (r RequestSpotInstancesRequest) Send(ctx context.Context) (*RequestSpotInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RequestSpotInstancesResponse{
		RequestSpotInstancesOutput: r.Request.Data.(*RequestSpotInstancesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RequestSpotInstancesResponse is the response type for the
// RequestSpotInstances API operation.
type RequestSpotInstancesResponse struct {
	*RequestSpotInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RequestSpotInstances request.
func (r *RequestSpotInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
