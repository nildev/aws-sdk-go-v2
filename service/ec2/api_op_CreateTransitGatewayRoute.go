// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTransitGatewayRouteRequest
type CreateTransitGatewayRouteInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether to drop traffic that matches this route.
	Blackhole *bool `type:"boolean"`

	// The CIDR range used for destination matches. Routing decisions are based
	// on the most specific match.
	//
	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the attachment.
	TransitGatewayAttachmentId *string `type:"string"`

	// The ID of the transit gateway route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTransitGatewayRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTransitGatewayRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTransitGatewayRouteInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTransitGatewayRouteResult
type CreateTransitGatewayRouteOutput struct {
	_ struct{} `type:"structure"`

	// Information about the route.
	Route *TransitGatewayRoute `json:"ec2:CreateTransitGatewayRouteOutput:Route" locationName:"route" type:"structure"`
}

// String returns the string representation
func (s CreateTransitGatewayRouteOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTransitGatewayRoute = "CreateTransitGatewayRoute"

// CreateTransitGatewayRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a static route for the specified transit gateway route table.
//
//    // Example sending a request using CreateTransitGatewayRouteRequest.
//    req := client.CreateTransitGatewayRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTransitGatewayRoute
func (c *Client) CreateTransitGatewayRouteRequest(input *CreateTransitGatewayRouteInput) CreateTransitGatewayRouteRequest {
	op := &aws.Operation{
		Name:       opCreateTransitGatewayRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTransitGatewayRouteInput{}
	}

	req := c.newRequest(op, input, &CreateTransitGatewayRouteOutput{})
	return CreateTransitGatewayRouteRequest{Request: req, Input: input, Copy: c.CreateTransitGatewayRouteRequest}
}

// CreateTransitGatewayRouteRequest is the request type for the
// CreateTransitGatewayRoute API operation.
type CreateTransitGatewayRouteRequest struct {
	*aws.Request
	Input *CreateTransitGatewayRouteInput
	Copy  func(*CreateTransitGatewayRouteInput) CreateTransitGatewayRouteRequest
}

// Send marshals and sends the CreateTransitGatewayRoute API request.
func (r CreateTransitGatewayRouteRequest) Send(ctx context.Context) (*CreateTransitGatewayRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTransitGatewayRouteResponse{
		CreateTransitGatewayRouteOutput: r.Request.Data.(*CreateTransitGatewayRouteOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTransitGatewayRouteResponse is the response type for the
// CreateTransitGatewayRoute API operation.
type CreateTransitGatewayRouteResponse struct {
	*CreateTransitGatewayRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTransitGatewayRoute request.
func (r *CreateTransitGatewayRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
