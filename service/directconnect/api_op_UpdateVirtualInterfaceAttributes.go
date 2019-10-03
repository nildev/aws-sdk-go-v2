// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/UpdateVirtualInterfaceAttributesRequest
type UpdateVirtualInterfaceAttributesInput struct {
	_ struct{} `type:"structure"`

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 9001. The default value is 1500.
	Mtu *int64 `locationName:"mtu" type:"integer"`

	// The ID of the virtual private interface.
	//
	// VirtualInterfaceId is a required field
	VirtualInterfaceId *string `locationName:"virtualInterfaceId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateVirtualInterfaceAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVirtualInterfaceAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateVirtualInterfaceAttributesInput"}

	if s.VirtualInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a virtual interface.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/VirtualInterface
type UpdateVirtualInterfaceAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The address family for the BGP peer.
	AddressFamily AddressFamily `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:AddressFamily" locationName:"addressFamily" type:"string" enum:"true"`

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:AmazonAddress" locationName:"amazonAddress" type:"string"`

	// The autonomous system number (ASN) for the Amazon side of the connection.
	AmazonSideAsn *int64 `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:AmazonSideAsn" locationName:"amazonSideAsn" type:"long"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	Asn *int64 `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:Asn" locationName:"asn" type:"integer"`

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:AuthKey" locationName:"authKey" type:"string"`

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDeviceV2 *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:AwsDeviceV2" locationName:"awsDeviceV2" type:"string"`

	// The BGP peers configured on this virtual interface.
	BgpPeers []BGPPeer `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:BgpPeers" locationName:"bgpPeers" type:"list"`

	// The ID of the connection.
	ConnectionId *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:ConnectionId" locationName:"connectionId" type:"string"`

	// The IP address assigned to the customer interface.
	CustomerAddress *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:CustomerAddress" locationName:"customerAddress" type:"string"`

	// The customer router configuration.
	CustomerRouterConfig *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:CustomerRouterConfig" locationName:"customerRouterConfig" type:"string"`

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:DirectConnectGatewayId" locationName:"directConnectGatewayId" type:"string"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:JumboFrameCapable" locationName:"jumboFrameCapable" type:"boolean"`

	// The location of the connection.
	Location *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:Location" locationName:"location" type:"string"`

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 9001. The default value is 1500.
	Mtu *int64 `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:Mtu" locationName:"mtu" type:"integer"`

	// The ID of the AWS account that owns the virtual interface.
	OwnerAccount *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:OwnerAccount" locationName:"ownerAccount" type:"string"`

	// The AWS Region where the virtual interface is located.
	Region *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:Region" locationName:"region" type:"string"`

	// The routes to be advertised to the AWS network in this Region. Applies to
	// public virtual interfaces.
	RouteFilterPrefixes []RouteFilterPrefix `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:RouteFilterPrefixes" locationName:"routeFilterPrefixes" type:"list"`

	// Any tags assigned to the virtual interface.
	Tags []Tag `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:Tags" locationName:"tags" min:"1" type:"list"`

	// The ID of the virtual private gateway. Applies only to private virtual interfaces.
	VirtualGatewayId *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:VirtualGatewayId" locationName:"virtualGatewayId" deprecated:"true" type:"string"`

	// The ID of the virtual interface.
	VirtualInterfaceId *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:VirtualInterfaceId" locationName:"virtualInterfaceId" type:"string"`

	// The name of the virtual interface assigned by the customer network.
	VirtualInterfaceName *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:VirtualInterfaceName" locationName:"virtualInterfaceName" type:"string"`

	// The state of the virtual interface. The following are the possible values:
	//
	//    * confirming: The creation of the virtual interface is pending confirmation
	//    from the virtual interface owner. If the owner of the virtual interface
	//    is different from the owner of the connection on which it is provisioned,
	//    then the virtual interface will remain in this state until it is confirmed
	//    by the virtual interface owner.
	//
	//    * verifying: This state only applies to public virtual interfaces. Each
	//    public virtual interface needs validation before the virtual interface
	//    can be created.
	//
	//    * pending: A virtual interface is in this state from the time that it
	//    is created until the virtual interface is ready to forward traffic.
	//
	//    * available: A virtual interface that is able to forward traffic.
	//
	//    * down: A virtual interface that is BGP down.
	//
	//    * deleting: A virtual interface is in this state immediately after calling
	//    DeleteVirtualInterface until it can no longer forward traffic.
	//
	//    * deleted: A virtual interface that cannot forward traffic.
	//
	//    * rejected: The virtual interface owner has declined creation of the virtual
	//    interface. If a virtual interface in the Confirming state is deleted by
	//    the virtual interface owner, the virtual interface enters the Rejected
	//    state.
	//
	//    * unknown: The state of the virtual interface is not available.
	VirtualInterfaceState VirtualInterfaceState `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:VirtualInterfaceState" locationName:"virtualInterfaceState" type:"string" enum:"true"`

	// The type of virtual interface. The possible values are private and public.
	VirtualInterfaceType *string `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:VirtualInterfaceType" locationName:"virtualInterfaceType" type:"string"`

	// The ID of the VLAN.
	Vlan *int64 `json:"directconnect:UpdateVirtualInterfaceAttributesOutput:Vlan" locationName:"vlan" type:"integer"`
}

// String returns the string representation
func (s UpdateVirtualInterfaceAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateVirtualInterfaceAttributes = "UpdateVirtualInterfaceAttributes"

// UpdateVirtualInterfaceAttributesRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Updates the specified attributes of the specified virtual private interface.
//
// Setting the MTU of a virtual interface to 9001 (jumbo frames) can cause an
// update to the underlying physical connection if it wasn't updated to support
// jumbo frames. Updating the connection disrupts network connectivity for all
// virtual interfaces associated with the connection for up to 30 seconds. To
// check whether your connection supports jumbo frames, call DescribeConnections.
// To check whether your virtual interface supports jumbo frames, call DescribeVirtualInterfaces.
//
//    // Example sending a request using UpdateVirtualInterfaceAttributesRequest.
//    req := client.UpdateVirtualInterfaceAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/UpdateVirtualInterfaceAttributes
func (c *Client) UpdateVirtualInterfaceAttributesRequest(input *UpdateVirtualInterfaceAttributesInput) UpdateVirtualInterfaceAttributesRequest {
	op := &aws.Operation{
		Name:       opUpdateVirtualInterfaceAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateVirtualInterfaceAttributesInput{}
	}

	req := c.newRequest(op, input, &UpdateVirtualInterfaceAttributesOutput{})
	return UpdateVirtualInterfaceAttributesRequest{Request: req, Input: input, Copy: c.UpdateVirtualInterfaceAttributesRequest}
}

// UpdateVirtualInterfaceAttributesRequest is the request type for the
// UpdateVirtualInterfaceAttributes API operation.
type UpdateVirtualInterfaceAttributesRequest struct {
	*aws.Request
	Input *UpdateVirtualInterfaceAttributesInput
	Copy  func(*UpdateVirtualInterfaceAttributesInput) UpdateVirtualInterfaceAttributesRequest
}

// Send marshals and sends the UpdateVirtualInterfaceAttributes API request.
func (r UpdateVirtualInterfaceAttributesRequest) Send(ctx context.Context) (*UpdateVirtualInterfaceAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVirtualInterfaceAttributesResponse{
		UpdateVirtualInterfaceAttributesOutput: r.Request.Data.(*UpdateVirtualInterfaceAttributesOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVirtualInterfaceAttributesResponse is the response type for the
// UpdateVirtualInterfaceAttributes API operation.
type UpdateVirtualInterfaceAttributesResponse struct {
	*UpdateVirtualInterfaceAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVirtualInterfaceAttributes request.
func (r *UpdateVirtualInterfaceAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
