// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreatePrivateVirtualInterfaceRequest
type CreatePrivateVirtualInterfaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the connection.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`

	// Information about the private virtual interface.
	//
	// NewPrivateVirtualInterface is a required field
	NewPrivateVirtualInterface *NewPrivateVirtualInterface `locationName:"newPrivateVirtualInterface" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreatePrivateVirtualInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePrivateVirtualInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePrivateVirtualInterfaceInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if s.NewPrivateVirtualInterface == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewPrivateVirtualInterface"))
	}
	if s.NewPrivateVirtualInterface != nil {
		if err := s.NewPrivateVirtualInterface.Validate(); err != nil {
			invalidParams.AddNested("NewPrivateVirtualInterface", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a virtual interface.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/VirtualInterface
type CreatePrivateVirtualInterfaceOutput struct {
	_ struct{} `type:"structure"`

	// The address family for the BGP peer.
	AddressFamily AddressFamily `json:"directconnect:CreatePrivateVirtualInterfaceOutput:AddressFamily" locationName:"addressFamily" type:"string" enum:"true"`

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:AmazonAddress" locationName:"amazonAddress" type:"string"`

	// The autonomous system number (ASN) for the Amazon side of the connection.
	AmazonSideAsn *int64 `json:"directconnect:CreatePrivateVirtualInterfaceOutput:AmazonSideAsn" locationName:"amazonSideAsn" type:"long"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	Asn *int64 `json:"directconnect:CreatePrivateVirtualInterfaceOutput:Asn" locationName:"asn" type:"integer"`

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:AuthKey" locationName:"authKey" type:"string"`

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDeviceV2 *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:AwsDeviceV2" locationName:"awsDeviceV2" type:"string"`

	// The BGP peers configured on this virtual interface.
	BgpPeers []BGPPeer `json:"directconnect:CreatePrivateVirtualInterfaceOutput:BgpPeers" locationName:"bgpPeers" type:"list"`

	// The ID of the connection.
	ConnectionId *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:ConnectionId" locationName:"connectionId" type:"string"`

	// The IP address assigned to the customer interface.
	CustomerAddress *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:CustomerAddress" locationName:"customerAddress" type:"string"`

	// The customer router configuration.
	CustomerRouterConfig *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:CustomerRouterConfig" locationName:"customerRouterConfig" type:"string"`

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:DirectConnectGatewayId" locationName:"directConnectGatewayId" type:"string"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `json:"directconnect:CreatePrivateVirtualInterfaceOutput:JumboFrameCapable" locationName:"jumboFrameCapable" type:"boolean"`

	// The location of the connection.
	Location *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:Location" locationName:"location" type:"string"`

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 9001. The default value is 1500.
	Mtu *int64 `json:"directconnect:CreatePrivateVirtualInterfaceOutput:Mtu" locationName:"mtu" type:"integer"`

	// The ID of the AWS account that owns the virtual interface.
	OwnerAccount *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:OwnerAccount" locationName:"ownerAccount" type:"string"`

	// The AWS Region where the virtual interface is located.
	Region *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:Region" locationName:"region" type:"string"`

	// The routes to be advertised to the AWS network in this Region. Applies to
	// public virtual interfaces.
	RouteFilterPrefixes []RouteFilterPrefix `json:"directconnect:CreatePrivateVirtualInterfaceOutput:RouteFilterPrefixes" locationName:"routeFilterPrefixes" type:"list"`

	// Any tags assigned to the virtual interface.
	Tags []Tag `json:"directconnect:CreatePrivateVirtualInterfaceOutput:Tags" locationName:"tags" min:"1" type:"list"`

	// The ID of the virtual private gateway. Applies only to private virtual interfaces.
	VirtualGatewayId *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:VirtualGatewayId" locationName:"virtualGatewayId" deprecated:"true" type:"string"`

	// The ID of the virtual interface.
	VirtualInterfaceId *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:VirtualInterfaceId" locationName:"virtualInterfaceId" type:"string"`

	// The name of the virtual interface assigned by the customer network.
	VirtualInterfaceName *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:VirtualInterfaceName" locationName:"virtualInterfaceName" type:"string"`

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
	VirtualInterfaceState VirtualInterfaceState `json:"directconnect:CreatePrivateVirtualInterfaceOutput:VirtualInterfaceState" locationName:"virtualInterfaceState" type:"string" enum:"true"`

	// The type of virtual interface. The possible values are private and public.
	VirtualInterfaceType *string `json:"directconnect:CreatePrivateVirtualInterfaceOutput:VirtualInterfaceType" locationName:"virtualInterfaceType" type:"string"`

	// The ID of the VLAN.
	Vlan *int64 `json:"directconnect:CreatePrivateVirtualInterfaceOutput:Vlan" locationName:"vlan" type:"integer"`
}

// String returns the string representation
func (s CreatePrivateVirtualInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePrivateVirtualInterface = "CreatePrivateVirtualInterface"

// CreatePrivateVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Creates a private virtual interface. A virtual interface is the VLAN that
// transports AWS Direct Connect traffic. A private virtual interface can be
// connected to either a Direct Connect gateway or a Virtual Private Gateway
// (VGW). Connecting the private virtual interface to a Direct Connect gateway
// enables the possibility for connecting to multiple VPCs, including VPCs in
// different AWS Regions. Connecting the private virtual interface to a VGW
// only provides access to a single VPC within the same Region.
//
//    // Example sending a request using CreatePrivateVirtualInterfaceRequest.
//    req := client.CreatePrivateVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreatePrivateVirtualInterface
func (c *Client) CreatePrivateVirtualInterfaceRequest(input *CreatePrivateVirtualInterfaceInput) CreatePrivateVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opCreatePrivateVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePrivateVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &CreatePrivateVirtualInterfaceOutput{})
	return CreatePrivateVirtualInterfaceRequest{Request: req, Input: input, Copy: c.CreatePrivateVirtualInterfaceRequest}
}

// CreatePrivateVirtualInterfaceRequest is the request type for the
// CreatePrivateVirtualInterface API operation.
type CreatePrivateVirtualInterfaceRequest struct {
	*aws.Request
	Input *CreatePrivateVirtualInterfaceInput
	Copy  func(*CreatePrivateVirtualInterfaceInput) CreatePrivateVirtualInterfaceRequest
}

// Send marshals and sends the CreatePrivateVirtualInterface API request.
func (r CreatePrivateVirtualInterfaceRequest) Send(ctx context.Context) (*CreatePrivateVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePrivateVirtualInterfaceResponse{
		CreatePrivateVirtualInterfaceOutput: r.Request.Data.(*CreatePrivateVirtualInterfaceOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePrivateVirtualInterfaceResponse is the response type for the
// CreatePrivateVirtualInterface API operation.
type CreatePrivateVirtualInterfaceResponse struct {
	*CreatePrivateVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePrivateVirtualInterface request.
func (r *CreatePrivateVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
