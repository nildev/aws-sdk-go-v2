// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteFleetInput
type DeleteFleetInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet to be deleted.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFleetInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteFleetOutput
type DeleteFleetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFleetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteFleet = "DeleteFleet"

// DeleteFleetRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Deletes everything related to a fleet. Before deleting a fleet, you must
// set the fleet's desired capacity to zero. See UpdateFleetCapacity.
//
// If the fleet being deleted has a VPC peering connection, you first need to
// get a valid authorization (good for 24 hours) by calling CreateVpcPeeringAuthorization.
// You do not need to explicitly delete the VPC peering connection--this is
// done as part of the delete fleet process.
//
// This action removes the fleet's resources and the fleet record. Once a fleet
// is deleted, you can no longer use that fleet.
//
// Learn more
//
//  Working with Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html).
//
// Related operations
//
//    * CreateFleet
//
//    * ListFleets
//
//    * DeleteFleet
//
//    * Describe fleets: DescribeFleetAttributes DescribeFleetCapacity DescribeFleetPortSettings
//    DescribeFleetUtilization DescribeRuntimeConfiguration DescribeEC2InstanceLimits
//    DescribeFleetEvents
//
//    * Update fleets: UpdateFleetAttributes UpdateFleetCapacity UpdateFleetPortSettings
//    UpdateRuntimeConfiguration
//
//    * Manage fleet actions: StartFleetActions StopFleetActions
//
//    // Example sending a request using DeleteFleetRequest.
//    req := client.DeleteFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteFleet
func (c *Client) DeleteFleetRequest(input *DeleteFleetInput) DeleteFleetRequest {
	op := &aws.Operation{
		Name:       opDeleteFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteFleetInput{}
	}

	req := c.newRequest(op, input, &DeleteFleetOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteFleetRequest{Request: req, Input: input, Copy: c.DeleteFleetRequest}
}

// DeleteFleetRequest is the request type for the
// DeleteFleet API operation.
type DeleteFleetRequest struct {
	*aws.Request
	Input *DeleteFleetInput
	Copy  func(*DeleteFleetInput) DeleteFleetRequest
}

// Send marshals and sends the DeleteFleet API request.
func (r DeleteFleetRequest) Send(ctx context.Context) (*DeleteFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFleetResponse{
		DeleteFleetOutput: r.Request.Data.(*DeleteFleetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFleetResponse is the response type for the
// DeleteFleet API operation.
type DeleteFleetResponse struct {
	*DeleteFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFleet request.
func (r *DeleteFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
