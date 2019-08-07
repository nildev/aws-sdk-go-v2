// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/PeerVpcRequest
type PeerVpcInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PeerVpcInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/PeerVpcResult
type PeerVpcOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the request operation.
	Operation *Operation `json:"lightsail:PeerVpcOutput:Operation" locationName:"operation" type:"structure"`
}

// String returns the string representation
func (s PeerVpcOutput) String() string {
	return awsutil.Prettify(s)
}

const opPeerVpc = "PeerVpc"

// PeerVpcRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Tries to peer the Lightsail VPC with the user's default VPC.
//
//    // Example sending a request using PeerVpcRequest.
//    req := client.PeerVpcRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/PeerVpc
func (c *Client) PeerVpcRequest(input *PeerVpcInput) PeerVpcRequest {
	op := &aws.Operation{
		Name:       opPeerVpc,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PeerVpcInput{}
	}

	req := c.newRequest(op, input, &PeerVpcOutput{})
	return PeerVpcRequest{Request: req, Input: input, Copy: c.PeerVpcRequest}
}

// PeerVpcRequest is the request type for the
// PeerVpc API operation.
type PeerVpcRequest struct {
	*aws.Request
	Input *PeerVpcInput
	Copy  func(*PeerVpcInput) PeerVpcRequest
}

// Send marshals and sends the PeerVpc API request.
func (r PeerVpcRequest) Send(ctx context.Context) (*PeerVpcResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PeerVpcResponse{
		PeerVpcOutput: r.Request.Data.(*PeerVpcOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PeerVpcResponse is the response type for the
// PeerVpc API operation.
type PeerVpcResponse struct {
	*PeerVpcOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PeerVpc request.
func (r *PeerVpcResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
