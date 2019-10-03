// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetNodeInput
type GetNodeInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the member that owns the node.
	//
	// MemberId is a required field
	MemberId *string `location:"uri" locationName:"memberId" min:"1" type:"string" required:"true"`

	// The unique identifier of the network to which the node belongs.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`

	// The unique identifier of the node.
	//
	// NodeId is a required field
	NodeId *string `location:"uri" locationName:"nodeId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetNodeInput"}

	if s.MemberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberId"))
	}
	if s.MemberId != nil && len(*s.MemberId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberId", 1))
	}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if s.NodeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeId"))
	}
	if s.NodeId != nil && len(*s.NodeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NodeId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetNodeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MemberId != nil {
		v := *s.MemberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "memberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NodeId != nil {
		v := *s.NodeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "nodeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetNodeOutput
type GetNodeOutput struct {
	_ struct{} `type:"structure"`

	// Properties of the node configuration.
	Node *Node `json:"managedblockchain:GetNodeOutput:Node" type:"structure"`
}

// String returns the string representation
func (s GetNodeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetNodeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Node != nil {
		v := s.Node

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Node", v, metadata)
	}
	return nil
}

const opGetNode = "GetNode"

// GetNodeRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Returns detailed information about a peer node.
//
//    // Example sending a request using GetNodeRequest.
//    req := client.GetNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetNode
func (c *Client) GetNodeRequest(input *GetNodeInput) GetNodeRequest {
	op := &aws.Operation{
		Name:       opGetNode,
		HTTPMethod: "GET",
		HTTPPath:   "/networks/{networkId}/members/{memberId}/nodes/{nodeId}",
	}

	if input == nil {
		input = &GetNodeInput{}
	}

	req := c.newRequest(op, input, &GetNodeOutput{})
	return GetNodeRequest{Request: req, Input: input, Copy: c.GetNodeRequest}
}

// GetNodeRequest is the request type for the
// GetNode API operation.
type GetNodeRequest struct {
	*aws.Request
	Input *GetNodeInput
	Copy  func(*GetNodeInput) GetNodeRequest
}

// Send marshals and sends the GetNode API request.
func (r GetNodeRequest) Send(ctx context.Context) (*GetNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetNodeResponse{
		GetNodeOutput: r.Request.Data.(*GetNodeOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetNodeResponse is the response type for the
// GetNode API operation.
type GetNodeResponse struct {
	*GetNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetNode request.
func (r *GetNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
