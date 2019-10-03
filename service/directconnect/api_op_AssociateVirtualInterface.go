// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AssociateVirtualInterfaceRequest
type AssociateVirtualInterfaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the LAG or connection.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`

	// The ID of the virtual interface.
	//
	// VirtualInterfaceId is a required field
	VirtualInterfaceId *string `locationName:"virtualInterfaceId" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateVirtualInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateVirtualInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateVirtualInterfaceInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

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
type AssociateVirtualInterfaceOutput struct {
	_ struct{} `type:"structure"`

	// The address family for the BGP peer.
	AddressFamily AddressFamily `json:"directconnect:AssociateVirtualInterfaceOutput:AddressFamily" locationName:"addressFamily" type:"string" enum:"true"`

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string `json:"directconnect:AssociateVirtualInterfaceOutput:AmazonAddress" locationName:"amazonAddress" type:"string"`

	// The autonomous system number (ASN) for the Amazon side of the connection.
	AmazonSideAsn *int64 `json:"directconnect:AssociateVirtualInterfaceOutput:AmazonSideAsn" locationName:"amazonSideAsn" type:"long"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	Asn *int64 `json:"directconnect:AssociateVirtualInterfaceOutput:Asn" locationName:"asn" type:"integer"`

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string `json:"directconnect:AssociateVirtualInterfaceOutput:AuthKey" locationName:"authKey" type:"string"`

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDeviceV2 *string `json:"directconnect:AssociateVirtualInterfaceOutput:AwsDeviceV2" locationName:"awsDeviceV2" type:"string"`

	// The BGP peers configured on this virtual interface.
	BgpPeers []BGPPeer `json:"directconnect:AssociateVirtualInterfaceOutput:BgpPeers" locationName:"bgpPeers" type:"list"`

	// The ID of the connection.
	ConnectionId *string `json:"directconnect:AssociateVirtualInterfaceOutput:ConnectionId" locationName:"connectionId" type:"string"`

	// The IP address assigned to the customer interface.
	CustomerAddress *string `json:"directconnect:AssociateVirtualInterfaceOutput:CustomerAddress" locationName:"customerAddress" type:"string"`

	// The customer router configuration.
	CustomerRouterConfig *string `json:"directconnect:AssociateVirtualInterfaceOutput:CustomerRouterConfig" locationName:"customerRouterConfig" type:"string"`

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string `json:"directconnect:AssociateVirtualInterfaceOutput:DirectConnectGatewayId" locationName:"directConnectGatewayId" type:"string"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `json:"directconnect:AssociateVirtualInterfaceOutput:JumboFrameCapable" locationName:"jumboFrameCapable" type:"boolean"`

	// The location of the connection.
	Location *string `json:"directconnect:AssociateVirtualInterfaceOutput:Location" locationName:"location" type:"string"`

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 9001. The default value is 1500.
	Mtu *int64 `json:"directconnect:AssociateVirtualInterfaceOutput:Mtu" locationName:"mtu" type:"integer"`

	// The ID of the AWS account that owns the virtual interface.
	OwnerAccount *string `json:"directconnect:AssociateVirtualInterfaceOutput:OwnerAccount" locationName:"ownerAccount" type:"string"`

	// The AWS Region where the virtual interface is located.
	Region *string `json:"directconnect:AssociateVirtualInterfaceOutput:Region" locationName:"region" type:"string"`

	// The routes to be advertised to the AWS network in this Region. Applies to
	// public virtual interfaces.
	RouteFilterPrefixes []RouteFilterPrefix `json:"directconnect:AssociateVirtualInterfaceOutput:RouteFilterPrefixes" locationName:"routeFilterPrefixes" type:"list"`

	// Any tags assigned to the virtual interface.
	Tags []Tag `json:"directconnect:AssociateVirtualInterfaceOutput:Tags" locationName:"tags" min:"1" type:"list"`

	// The ID of the virtual private gateway. Applies only to private virtual interfaces.
	VirtualGatewayId *string `json:"directconnect:AssociateVirtualInterfaceOutput:VirtualGatewayId" locationName:"virtualGatewayId" deprecated:"true" type:"string"`

	// The ID of the virtual interface.
	VirtualInterfaceId *string `json:"directconnect:AssociateVirtualInterfaceOutput:VirtualInterfaceId" locationName:"virtualInterfaceId" type:"string"`

	// The name of the virtual interface assigned by the customer network.
	VirtualInterfaceName *string `json:"directconnect:AssociateVirtualInterfaceOutput:VirtualInterfaceName" locationName:"virtualInterfaceName" type:"string"`

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
	VirtualInterfaceState VirtualInterfaceState `json:"directconnect:AssociateVirtualInterfaceOutput:VirtualInterfaceState" locationName:"virtualInterfaceState" type:"string" enum:"true"`

	// The type of virtual interface. The possible values are private and public.
	VirtualInterfaceType *string `json:"directconnect:AssociateVirtualInterfaceOutput:VirtualInterfaceType" locationName:"virtualInterfaceType" type:"string"`

	// The ID of the VLAN.
	Vlan *int64 `json:"directconnect:AssociateVirtualInterfaceOutput:Vlan" locationName:"vlan" type:"integer"`
}

// String returns the string representation
func (s AssociateVirtualInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateVirtualInterface = "AssociateVirtualInterface"

// AssociateVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Associates a virtual interface with a specified link aggregation group (LAG)
// or connection. Connectivity to AWS is temporarily interrupted as the virtual
// interface is being migrated. If the target connection or LAG has an associated
// virtual interface with a conflicting VLAN number or a conflicting IP address,
// the operation fails.
//
// Virtual interfaces associated with a hosted connection cannot be associated
// with a LAG; hosted connections must be migrated along with their virtual
// interfaces using AssociateHostedConnection.
//
// To reassociate a virtual interface to a new connection or LAG, the requester
// must own either the virtual interface itself or the connection to which the
// virtual interface is currently associated. Additionally, the requester must
// own the connection or LAG for the association.
//
//    // Example sending a request using AssociateVirtualInterfaceRequest.
//    req := client.AssociateVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AssociateVirtualInterface
func (c *Client) AssociateVirtualInterfaceRequest(input *AssociateVirtualInterfaceInput) AssociateVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opAssociateVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &AssociateVirtualInterfaceOutput{})
	return AssociateVirtualInterfaceRequest{Request: req, Input: input, Copy: c.AssociateVirtualInterfaceRequest}
}

// AssociateVirtualInterfaceRequest is the request type for the
// AssociateVirtualInterface API operation.
type AssociateVirtualInterfaceRequest struct {
	*aws.Request
	Input *AssociateVirtualInterfaceInput
	Copy  func(*AssociateVirtualInterfaceInput) AssociateVirtualInterfaceRequest
}

// Send marshals and sends the AssociateVirtualInterface API request.
func (r AssociateVirtualInterfaceRequest) Send(ctx context.Context) (*AssociateVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateVirtualInterfaceResponse{
		AssociateVirtualInterfaceOutput: r.Request.Data.(*AssociateVirtualInterfaceOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateVirtualInterfaceResponse is the response type for the
// AssociateVirtualInterface API operation.
type AssociateVirtualInterfaceResponse struct {
	*AssociateVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateVirtualInterface request.
func (r *AssociateVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
