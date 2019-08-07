// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayRouteRequest
type DeleteTransitGatewayRouteInput struct {
	_ struct{} `type:"structure"`

	// The CIDR range for the route. This must match the CIDR for the route exactly.
	//
	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the transit gateway route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitGatewayRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitGatewayRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTransitGatewayRouteInput"}

	if s.DestinationCidrBlock == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCidrBlock"))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayRouteResult
type DeleteTransitGatewayRouteOutput struct {
	_ struct{} `type:"structure"`

	// Information about the route.
	Route *TransitGatewayRoute `json:"ec2:DeleteTransitGatewayRouteOutput:Route" locationName:"route" type:"structure"`
}

// String returns the string representation
func (s DeleteTransitGatewayRouteOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTransitGatewayRoute = "DeleteTransitGatewayRoute"

// DeleteTransitGatewayRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified route from the specified transit gateway route table.
//
//    // Example sending a request using DeleteTransitGatewayRouteRequest.
//    req := client.DeleteTransitGatewayRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayRoute
func (c *Client) DeleteTransitGatewayRouteRequest(input *DeleteTransitGatewayRouteInput) DeleteTransitGatewayRouteRequest {
	op := &aws.Operation{
		Name:       opDeleteTransitGatewayRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTransitGatewayRouteInput{}
	}

	req := c.newRequest(op, input, &DeleteTransitGatewayRouteOutput{})
	return DeleteTransitGatewayRouteRequest{Request: req, Input: input, Copy: c.DeleteTransitGatewayRouteRequest}
}

// DeleteTransitGatewayRouteRequest is the request type for the
// DeleteTransitGatewayRoute API operation.
type DeleteTransitGatewayRouteRequest struct {
	*aws.Request
	Input *DeleteTransitGatewayRouteInput
	Copy  func(*DeleteTransitGatewayRouteInput) DeleteTransitGatewayRouteRequest
}

// Send marshals and sends the DeleteTransitGatewayRoute API request.
func (r DeleteTransitGatewayRouteRequest) Send(ctx context.Context) (*DeleteTransitGatewayRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTransitGatewayRouteResponse{
		DeleteTransitGatewayRouteOutput: r.Request.Data.(*DeleteTransitGatewayRouteOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTransitGatewayRouteResponse is the response type for the
// DeleteTransitGatewayRoute API operation.
type DeleteTransitGatewayRouteResponse struct {
	*DeleteTransitGatewayRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTransitGatewayRoute request.
func (r *DeleteTransitGatewayRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
