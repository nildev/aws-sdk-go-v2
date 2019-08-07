// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// For List operations, an optional specification to return a subset of objects,
// such as resolver endpoints or resolver rules.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/Filter
type Filter struct {
	_ struct{} `type:"structure"`

	// When you're using a List operation and you want the operation to return a
	// subset of objects, such as resolver endpoints or resolver rules, the name
	// of the parameter that you want to use to filter objects. For example, to
	// list only inbound resolver endpoints, specify Direction for the value of
	// Name.
	Name *string `json:"route53resolver:Filter:Name" min:"1" type:"string"`

	// When you're using a List operation and you want the operation to return a
	// subset of objects, such as resolver endpoints or resolver rules, the value
	// of the parameter that you want to use to filter objects. For example, to
	// list only inbound resolver endpoints, specify INBOUND for the value of Values.
	Values []string `json:"route53resolver:Filter:Values" type:"list"`
}

// String returns the string representation
func (s Filter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Filter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Filter"}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// In an CreateResolverEndpoint request, a subnet and IP address that you want
// to use for DNS queries.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/IpAddressRequest
type IpAddressRequest struct {
	_ struct{} `type:"structure"`

	// The IP address that you want to use for DNS queries.
	Ip *string `json:"route53resolver:IpAddressRequest:Ip" min:"7" type:"string"`

	// The subnet that contains the IP address.
	//
	// SubnetId is a required field
	SubnetId *string `json:"route53resolver:IpAddressRequest:SubnetId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s IpAddressRequest) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *IpAddressRequest) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "IpAddressRequest"}
	if s.Ip != nil && len(*s.Ip) < 7 {
		invalidParams.Add(aws.NewErrParamMinLen("Ip", 7))
	}

	if s.SubnetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetId"))
	}
	if s.SubnetId != nil && len(*s.SubnetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SubnetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// In the response to a GetResolverEndpoint request, information about the IP
// addresses that the resolver endpoint uses for DNS queries.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/IpAddressResponse
type IpAddressResponse struct {
	_ struct{} `type:"structure"`

	// The date and time that the IP address was created, in Unix time format and
	// Coordinated Universal Time (UTC).
	CreationTime *string `json:"route53resolver:IpAddressResponse:CreationTime" min:"20" type:"string"`

	// One IP address that the resolver endpoint uses for DNS queries.
	Ip *string `json:"route53resolver:IpAddressResponse:Ip" min:"7" type:"string"`

	// The ID of one IP address.
	IpId *string `json:"route53resolver:IpAddressResponse:IpId" min:"1" type:"string"`

	// The date and time that the IP address was last modified, in Unix time format
	// and Coordinated Universal Time (UTC).
	ModificationTime *string `json:"route53resolver:IpAddressResponse:ModificationTime" min:"20" type:"string"`

	// A status code that gives the current status of the request.
	Status IpAddressStatus `json:"route53resolver:IpAddressResponse:Status" type:"string" enum:"true"`

	// A message that provides additional information about the status of the request.
	StatusMessage *string `json:"route53resolver:IpAddressResponse:StatusMessage" type:"string"`

	// The ID of one subnet.
	SubnetId *string `json:"route53resolver:IpAddressResponse:SubnetId" min:"1" type:"string"`
}

// String returns the string representation
func (s IpAddressResponse) String() string {
	return awsutil.Prettify(s)
}

// In an UpdateResolverEndpoint request, information about an IP address to
// update.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/IpAddressUpdate
type IpAddressUpdate struct {
	_ struct{} `type:"structure"`

	// The new IP address.
	Ip *string `json:"route53resolver:IpAddressUpdate:Ip" min:"7" type:"string"`

	// Only when removing an IP address from a resolver endpoint: The ID of the
	// IP address that you want to remove. To get this ID, use GetResolverEndpoint.
	IpId *string `json:"route53resolver:IpAddressUpdate:IpId" min:"1" type:"string"`

	// The ID of the subnet that includes the IP address that you want to update.
	// To get this ID, use GetResolverEndpoint.
	SubnetId *string `json:"route53resolver:IpAddressUpdate:SubnetId" min:"1" type:"string"`
}

// String returns the string representation
func (s IpAddressUpdate) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *IpAddressUpdate) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "IpAddressUpdate"}
	if s.Ip != nil && len(*s.Ip) < 7 {
		invalidParams.Add(aws.NewErrParamMinLen("Ip", 7))
	}
	if s.IpId != nil && len(*s.IpId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IpId", 1))
	}
	if s.SubnetId != nil && len(*s.SubnetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SubnetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// In the response to a CreateResolverEndpoint, DeleteResolverEndpoint, GetResolverEndpoint,
// ListResolverEndpoints, or UpdateResolverEndpoint request, a complex type
// that contains settings for an existing inbound or outbound resolver endpoint.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ResolverEndpoint
type ResolverEndpoint struct {
	_ struct{} `type:"structure"`

	// The ARN (Amazon Resource Name) for the resolver endpoint.
	Arn *string `json:"route53resolver:ResolverEndpoint:Arn" min:"1" type:"string"`

	// The date and time that the endpoint was created, in Unix time format and
	// Coordinated Universal Time (UTC).
	CreationTime *string `json:"route53resolver:ResolverEndpoint:CreationTime" min:"20" type:"string"`

	// A unique string that identifies the request that created the resolver endpoint.
	// The CreatorRequestId allows failed requests to be retried without the risk
	// of executing the operation twice.
	CreatorRequestId *string `json:"route53resolver:ResolverEndpoint:CreatorRequestId" min:"1" type:"string"`

	// Indicates whether the resolver endpoint allows inbound or outbound DNS queries:
	//
	//    * INBOUND: allows DNS queries to your VPC from your network or another
	//    VPC
	//
	//    * OUTBOUND: allows DNS queries from your VPC to your network or another
	//    VPC
	Direction ResolverEndpointDirection `json:"route53resolver:ResolverEndpoint:Direction" type:"string" enum:"true"`

	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVPCId *string `json:"route53resolver:ResolverEndpoint:HostVPCId" min:"1" type:"string"`

	// The ID of the resolver endpoint.
	Id *string `json:"route53resolver:ResolverEndpoint:Id" min:"1" type:"string"`

	// The number of IP addresses that the resolver endpoint can use for DNS queries.
	IpAddressCount *int64 `json:"route53resolver:ResolverEndpoint:IpAddressCount" type:"integer"`

	// The date and time that the endpoint was last modified, in Unix time format
	// and Coordinated Universal Time (UTC).
	ModificationTime *string `json:"route53resolver:ResolverEndpoint:ModificationTime" min:"20" type:"string"`

	// The name that you assigned to the resolver endpoint when you submitted a
	// CreateResolverEndpoint request.
	Name *string `json:"route53resolver:ResolverEndpoint:Name" type:"string"`

	// The ID of one or more security groups that control access to this VPC. The
	// security group must include one or more inbound resolver rules.
	SecurityGroupIds []string `json:"route53resolver:ResolverEndpoint:SecurityGroupIds" type:"list"`

	// A code that specifies the current status of the resolver endpoint.
	Status ResolverEndpointStatus `json:"route53resolver:ResolverEndpoint:Status" type:"string" enum:"true"`

	// A detailed description of the status of the resolver endpoint.
	StatusMessage *string `json:"route53resolver:ResolverEndpoint:StatusMessage" type:"string"`
}

// String returns the string representation
func (s ResolverEndpoint) String() string {
	return awsutil.Prettify(s)
}

// For queries that originate in your VPC, detailed information about a resolver
// rule, which specifies how to route DNS queries out of the VPC. The ResolverRule
// parameter appears in the response to a CreateResolverRule, DeleteResolverRule,
// GetResolverRule, ListResolverRules, or UpdateResolverRule request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ResolverRule
type ResolverRule struct {
	_ struct{} `type:"structure"`

	// The ARN (Amazon Resource Name) for the resolver rule specified by Id.
	Arn *string `json:"route53resolver:ResolverRule:Arn" min:"1" type:"string"`

	// A unique string that you specified when you created the resolver rule. CreatorRequestIdidentifies
	// the request and allows failed requests to be retried without the risk of
	// executing the operation twice.
	CreatorRequestId *string `json:"route53resolver:ResolverRule:CreatorRequestId" min:"1" type:"string"`

	// DNS queries for this domain name are forwarded to the IP addresses that are
	// specified in TargetIps. If a query matches multiple resolver rules (example.com
	// and www.example.com), the query is routed using the resolver rule that contains
	// the most specific domain name (www.example.com).
	DomainName *string `json:"route53resolver:ResolverRule:DomainName" min:"1" type:"string"`

	// The ID that Resolver assigned to the resolver rule when you created it.
	Id *string `json:"route53resolver:ResolverRule:Id" min:"1" type:"string"`

	// The name for the resolver rule, which you specified when you created the
	// resolver rule.
	Name *string `json:"route53resolver:ResolverRule:Name" type:"string"`

	// When a rule is shared with another AWS account, the account ID of the account
	// that the rule is shared with.
	OwnerId *string `json:"route53resolver:ResolverRule:OwnerId" min:"12" type:"string"`

	// The ID of the endpoint that the rule is associated with.
	ResolverEndpointId *string `json:"route53resolver:ResolverRule:ResolverEndpointId" min:"1" type:"string"`

	// This value is always FORWARD. Other resolver rule types aren't supported.
	RuleType RuleTypeOption `json:"route53resolver:ResolverRule:RuleType" type:"string" enum:"true"`

	// Whether the rules is shared and, if so, whether the current account is sharing
	// the rule with another account, or another account is sharing the rule with
	// the current account.
	ShareStatus ShareStatus `json:"route53resolver:ResolverRule:ShareStatus" type:"string" enum:"true"`

	// A code that specifies the current status of the resolver rule.
	Status ResolverRuleStatus `json:"route53resolver:ResolverRule:Status" type:"string" enum:"true"`

	// A detailed description of the status of a resolver rule.
	StatusMessage *string `json:"route53resolver:ResolverRule:StatusMessage" type:"string"`

	// An array that contains the IP addresses and ports that you want to forward
	TargetIps []TargetAddress `json:"route53resolver:ResolverRule:TargetIps" min:"1" type:"list"`
}

// String returns the string representation
func (s ResolverRule) String() string {
	return awsutil.Prettify(s)
}

// In the response to an AssociateResolverRule, DisassociateResolverRule, or
// ListResolverRuleAssociations request, information about an association between
// a resolver rule and a VPC.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ResolverRuleAssociation
type ResolverRuleAssociation struct {
	_ struct{} `type:"structure"`

	// The ID of the association between a resolver rule and a VPC. Resolver assigns
	// this value when you submit an AssociateResolverRule request.
	Id *string `json:"route53resolver:ResolverRuleAssociation:Id" min:"1" type:"string"`

	// The name of an association between a resolver rule and a VPC.
	Name *string `json:"route53resolver:ResolverRuleAssociation:Name" type:"string"`

	// The ID of the resolver rule that you associated with the VPC that is specified
	// by VPCId.
	ResolverRuleId *string `json:"route53resolver:ResolverRuleAssociation:ResolverRuleId" min:"1" type:"string"`

	// A code that specifies the current status of the association between a resolver
	// rule and a VPC.
	Status ResolverRuleAssociationStatus `json:"route53resolver:ResolverRuleAssociation:Status" type:"string" enum:"true"`

	// A detailed description of the status of the association between a resolver
	// rule and a VPC.
	StatusMessage *string `json:"route53resolver:ResolverRuleAssociation:StatusMessage" type:"string"`

	// The ID of the VPC that you associated the resolver rule with.
	VPCId *string `json:"route53resolver:ResolverRuleAssociation:VPCId" min:"1" type:"string"`
}

// String returns the string representation
func (s ResolverRuleAssociation) String() string {
	return awsutil.Prettify(s)
}

// In an UpdateResolverRule request, information about the changes that you
// want to make.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ResolverRuleConfig
type ResolverRuleConfig struct {
	_ struct{} `type:"structure"`

	// The new name for the resolver rule. The name that you specify appears in
	// the Resolver dashboard in the Route 53 console.
	Name *string `json:"route53resolver:ResolverRuleConfig:Name" type:"string"`

	// The ID of the new outbound resolver endpoint that you want to use to route
	// DNS queries to the IP addresses that you specify in TargetIps.
	ResolverEndpointId *string `json:"route53resolver:ResolverRuleConfig:ResolverEndpointId" min:"1" type:"string"`

	// For DNS queries that originate in your VPC, the new IP addresses that you
	// want to route outbound DNS queries to.
	TargetIps []TargetAddress `json:"route53resolver:ResolverRuleConfig:TargetIps" min:"1" type:"list"`
}

// String returns the string representation
func (s ResolverRuleConfig) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResolverRuleConfig) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResolverRuleConfig"}
	if s.ResolverEndpointId != nil && len(*s.ResolverEndpointId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverEndpointId", 1))
	}
	if s.TargetIps != nil && len(s.TargetIps) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetIps", 1))
	}
	if s.TargetIps != nil {
		for i, v := range s.TargetIps {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TargetIps", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// One tag that you want to add to the specified resource. A tag consists of
// a Key (a name for the tag) and a Value.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// The name for the tag. For example, if you want to associate Resolver resources
	// with the account IDs of your customers for billing purposes, the value of
	// Key might be account-id.
	Key *string `json:"route53resolver:Tag:Key" type:"string"`

	// The value for the tag. For example, if Key is account-id, then Value might
	// be the ID of the customer account that you're creating the resource for.
	Value *string `json:"route53resolver:Tag:Value" type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// In a CreateResolverRule request, an array of the IPs that you want to forward
// DNS queries to.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/TargetAddress
type TargetAddress struct {
	_ struct{} `type:"structure"`

	// One IP address that you want to forward DNS queries to. You can specify only
	// IPv4 addresses.
	//
	// Ip is a required field
	Ip *string `json:"route53resolver:TargetAddress:Ip" min:"7" type:"string" required:"true"`

	// The port at Ip that you want to forward DNS queries to.
	Port *int64 `json:"route53resolver:TargetAddress:Port" type:"integer"`
}

// String returns the string representation
func (s TargetAddress) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TargetAddress) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TargetAddress"}

	if s.Ip == nil {
		invalidParams.Add(aws.NewErrParamRequired("Ip"))
	}
	if s.Ip != nil && len(*s.Ip) < 7 {
		invalidParams.Add(aws.NewErrParamMinLen("Ip", 7))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
