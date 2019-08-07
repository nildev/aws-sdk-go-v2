// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the VPC that you want to disassociate
// from a specified private hosted zone.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/DisassociateVPCFromHostedZoneRequest
type DisassociateVPCFromHostedZoneInput struct {
	_ struct{} `locationName:"DisassociateVPCFromHostedZoneRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// Optional: A comment about the disassociation request.
	Comment *string `type:"string"`

	// The ID of the private hosted zone that you want to disassociate a VPC from.
	//
	// HostedZoneId is a required field
	HostedZoneId *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// A complex type that contains information about the VPC that you're disassociating
	// from the specified hosted zone.
	//
	// VPC is a required field
	VPC *VPC `type:"structure" required:"true"`
}

// String returns the string representation
func (s DisassociateVPCFromHostedZoneInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateVPCFromHostedZoneInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateVPCFromHostedZoneInput"}

	if s.HostedZoneId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HostedZoneId"))
	}

	if s.VPC == nil {
		invalidParams.Add(aws.NewErrParamRequired("VPC"))
	}
	if s.VPC != nil {
		if err := s.VPC.Validate(); err != nil {
			invalidParams.AddNested("VPC", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateVPCFromHostedZoneInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "DisassociateVPCFromHostedZoneRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.Comment != nil {
			v := *s.Comment

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "Comment", protocol.StringValue(v), metadata)
		}
		if s.VPC != nil {
			v := s.VPC

			metadata := protocol.Metadata{}
			e.SetFields(protocol.BodyTarget, "VPC", v, metadata)
		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	if s.HostedZoneId != nil {
		v := *s.HostedZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response information for the disassociate
// request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/DisassociateVPCFromHostedZoneResponse
type DisassociateVPCFromHostedZoneOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that describes the changes made to the specified private hosted
	// zone.
	//
	// ChangeInfo is a required field
	ChangeInfo *ChangeInfo `json:"route53:DisassociateVPCFromHostedZoneOutput:ChangeInfo" type:"structure" required:"true"`
}

// String returns the string representation
func (s DisassociateVPCFromHostedZoneOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateVPCFromHostedZoneOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ChangeInfo != nil {
		v := s.ChangeInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ChangeInfo", v, metadata)
	}
	return nil
}

const opDisassociateVPCFromHostedZone = "DisassociateVPCFromHostedZone"

// DisassociateVPCFromHostedZoneRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Disassociates a VPC from a Amazon Route 53 private hosted zone. Note the
// following:
//
//    * You can't disassociate the last VPC from a private hosted zone.
//
//    * You can't convert a private hosted zone into a public hosted zone.
//
//    * You can submit a DisassociateVPCFromHostedZone request using either
//    the account that created the hosted zone or the account that created the
//    VPC.
//
//    // Example sending a request using DisassociateVPCFromHostedZoneRequest.
//    req := client.DisassociateVPCFromHostedZoneRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/DisassociateVPCFromHostedZone
func (c *Client) DisassociateVPCFromHostedZoneRequest(input *DisassociateVPCFromHostedZoneInput) DisassociateVPCFromHostedZoneRequest {
	op := &aws.Operation{
		Name:       opDisassociateVPCFromHostedZone,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/hostedzone/{Id}/disassociatevpc",
	}

	if input == nil {
		input = &DisassociateVPCFromHostedZoneInput{}
	}

	req := c.newRequest(op, input, &DisassociateVPCFromHostedZoneOutput{})
	return DisassociateVPCFromHostedZoneRequest{Request: req, Input: input, Copy: c.DisassociateVPCFromHostedZoneRequest}
}

// DisassociateVPCFromHostedZoneRequest is the request type for the
// DisassociateVPCFromHostedZone API operation.
type DisassociateVPCFromHostedZoneRequest struct {
	*aws.Request
	Input *DisassociateVPCFromHostedZoneInput
	Copy  func(*DisassociateVPCFromHostedZoneInput) DisassociateVPCFromHostedZoneRequest
}

// Send marshals and sends the DisassociateVPCFromHostedZone API request.
func (r DisassociateVPCFromHostedZoneRequest) Send(ctx context.Context) (*DisassociateVPCFromHostedZoneResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateVPCFromHostedZoneResponse{
		DisassociateVPCFromHostedZoneOutput: r.Request.Data.(*DisassociateVPCFromHostedZoneOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateVPCFromHostedZoneResponse is the response type for the
// DisassociateVPCFromHostedZone API operation.
type DisassociateVPCFromHostedZoneResponse struct {
	*DisassociateVPCFromHostedZoneOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateVPCFromHostedZone request.
func (r *DisassociateVPCFromHostedZoneResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
