// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreateDirectConnectGatewayRequest
type CreateDirectConnectGatewayInput struct {
	_ struct{} `type:"structure"`

	// The autonomous system number (ASN) for Border Gateway Protocol (BGP) to be
	// configured on the Amazon side of the connection. The ASN must be in the private
	// range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294. The default
	// is 64512.
	AmazonSideAsn *int64 `locationName:"amazonSideAsn" type:"long"`

	// The name of the Direct Connect gateway.
	//
	// DirectConnectGatewayName is a required field
	DirectConnectGatewayName *string `locationName:"directConnectGatewayName" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDirectConnectGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDirectConnectGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDirectConnectGatewayInput"}

	if s.DirectConnectGatewayName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectConnectGatewayName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreateDirectConnectGatewayResult
type CreateDirectConnectGatewayOutput struct {
	_ struct{} `type:"structure"`

	// The Direct Connect gateway.
	DirectConnectGateway *DirectConnectGateway `json:"directconnect:CreateDirectConnectGatewayOutput:DirectConnectGateway" locationName:"directConnectGateway" type:"structure"`
}

// String returns the string representation
func (s CreateDirectConnectGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDirectConnectGateway = "CreateDirectConnectGateway"

// CreateDirectConnectGatewayRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Creates a Direct Connect gateway, which is an intermediate object that enables
// you to connect a set of virtual interfaces and virtual private gateways.
// A Direct Connect gateway is global and visible in any AWS Region after it
// is created. The virtual interfaces and virtual private gateways that are
// connected through a Direct Connect gateway can be in different AWS Regions.
// This enables you to connect to a VPC in any Region, regardless of the Region
// in which the virtual interfaces are located, and pass traffic between them.
//
//    // Example sending a request using CreateDirectConnectGatewayRequest.
//    req := client.CreateDirectConnectGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreateDirectConnectGateway
func (c *Client) CreateDirectConnectGatewayRequest(input *CreateDirectConnectGatewayInput) CreateDirectConnectGatewayRequest {
	op := &aws.Operation{
		Name:       opCreateDirectConnectGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDirectConnectGatewayInput{}
	}

	req := c.newRequest(op, input, &CreateDirectConnectGatewayOutput{})
	return CreateDirectConnectGatewayRequest{Request: req, Input: input, Copy: c.CreateDirectConnectGatewayRequest}
}

// CreateDirectConnectGatewayRequest is the request type for the
// CreateDirectConnectGateway API operation.
type CreateDirectConnectGatewayRequest struct {
	*aws.Request
	Input *CreateDirectConnectGatewayInput
	Copy  func(*CreateDirectConnectGatewayInput) CreateDirectConnectGatewayRequest
}

// Send marshals and sends the CreateDirectConnectGateway API request.
func (r CreateDirectConnectGatewayRequest) Send(ctx context.Context) (*CreateDirectConnectGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDirectConnectGatewayResponse{
		CreateDirectConnectGatewayOutput: r.Request.Data.(*CreateDirectConnectGatewayOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDirectConnectGatewayResponse is the response type for the
// CreateDirectConnectGateway API operation.
type CreateDirectConnectGatewayResponse struct {
	*CreateDirectConnectGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDirectConnectGateway request.
func (r *CreateDirectConnectGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
