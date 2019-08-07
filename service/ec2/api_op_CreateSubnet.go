// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateSubnetRequest
type CreateSubnetInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone for the subnet.
	//
	// Default: AWS selects one for you. If you create more than one subnet in your
	// VPC, we may not necessarily select a different zone for each subnet.
	AvailabilityZone *string `type:"string"`

	// The AZ ID of the subnet.
	AvailabilityZoneId *string `type:"string"`

	// The IPv4 network range for the subnet, in CIDR notation. For example, 10.0.0.0/24.
	//
	// CidrBlock is a required field
	CidrBlock *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The IPv6 network range for the subnet, in CIDR notation. The subnet size
	// must use a /64 prefix length.
	Ipv6CidrBlock *string `type:"string"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSubnetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSubnetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSubnetInput"}

	if s.CidrBlock == nil {
		invalidParams.Add(aws.NewErrParamRequired("CidrBlock"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateSubnetResult
type CreateSubnetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the subnet.
	Subnet *Subnet `json:"ec2:CreateSubnetOutput:Subnet" locationName:"subnet" type:"structure"`
}

// String returns the string representation
func (s CreateSubnetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSubnet = "CreateSubnet"

// CreateSubnetRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a subnet in an existing VPC.
//
// When you create each subnet, you provide the VPC ID and IPv4 CIDR block for
// the subnet. After you create a subnet, you can't change its CIDR block. The
// size of the subnet's IPv4 CIDR block can be the same as a VPC's IPv4 CIDR
// block, or a subset of a VPC's IPv4 CIDR block. If you create more than one
// subnet in a VPC, the subnets' CIDR blocks must not overlap. The smallest
// IPv4 subnet (and VPC) you can create uses a /28 netmask (16 IPv4 addresses),
// and the largest uses a /16 netmask (65,536 IPv4 addresses).
//
// If you've associated an IPv6 CIDR block with your VPC, you can create a subnet
// with an IPv6 CIDR block that uses a /64 prefix length.
//
// AWS reserves both the first four and the last IPv4 address in each subnet's
// CIDR block. They're not available for use.
//
// If you add more than one subnet to a VPC, they're set up in a star topology
// with a logical router in the middle.
//
// If you launch an instance in a VPC using an Amazon EBS-backed AMI, the IP
// address doesn't change if you stop and restart the instance (unlike a similar
// instance launched outside a VPC, which gets a new IP address when restarted).
// It's therefore possible to have a subnet with no running instances (they're
// all stopped), but no remaining IP addresses available.
//
// For more information about subnets, see Your VPC and Subnets (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateSubnetRequest.
//    req := client.CreateSubnetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateSubnet
func (c *Client) CreateSubnetRequest(input *CreateSubnetInput) CreateSubnetRequest {
	op := &aws.Operation{
		Name:       opCreateSubnet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSubnetInput{}
	}

	req := c.newRequest(op, input, &CreateSubnetOutput{})
	return CreateSubnetRequest{Request: req, Input: input, Copy: c.CreateSubnetRequest}
}

// CreateSubnetRequest is the request type for the
// CreateSubnet API operation.
type CreateSubnetRequest struct {
	*aws.Request
	Input *CreateSubnetInput
	Copy  func(*CreateSubnetInput) CreateSubnetRequest
}

// Send marshals and sends the CreateSubnet API request.
func (r CreateSubnetRequest) Send(ctx context.Context) (*CreateSubnetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSubnetResponse{
		CreateSubnetOutput: r.Request.Data.(*CreateSubnetOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSubnetResponse is the response type for the
// CreateSubnet API operation.
type CreateSubnetResponse struct {
	*CreateSubnetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSubnet request.
func (r *CreateSubnetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
