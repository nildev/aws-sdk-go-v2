// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/DisassociateResolverRuleRequest
type DisassociateResolverRuleInput struct {
	_ struct{} `type:"structure"`

	// The ID of the resolver rule that you want to disassociate from the specified
	// VPC.
	//
	// ResolverRuleId is a required field
	ResolverRuleId *string `min:"1" type:"string" required:"true"`

	// The ID of the VPC that you want to disassociate the resolver rule from.
	//
	// VPCId is a required field
	VPCId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateResolverRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateResolverRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateResolverRuleInput"}

	if s.ResolverRuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverRuleId"))
	}
	if s.ResolverRuleId != nil && len(*s.ResolverRuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverRuleId", 1))
	}

	if s.VPCId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VPCId"))
	}
	if s.VPCId != nil && len(*s.VPCId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VPCId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/DisassociateResolverRuleResponse
type DisassociateResolverRuleOutput struct {
	_ struct{} `type:"structure"`

	// Information about the DisassociateResolverRule request, including the status
	// of the request.
	ResolverRuleAssociation *ResolverRuleAssociation `json:"route53resolver:DisassociateResolverRuleOutput:ResolverRuleAssociation" type:"structure"`
}

// String returns the string representation
func (s DisassociateResolverRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateResolverRule = "DisassociateResolverRule"

// DisassociateResolverRuleRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Removes the association between a specified resolver rule and a specified
// VPC.
//
// If you disassociate a resolver rule from a VPC, Resolver stops forwarding
// DNS queries for the domain name that you specified in the resolver rule.
//
//    // Example sending a request using DisassociateResolverRuleRequest.
//    req := client.DisassociateResolverRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/DisassociateResolverRule
func (c *Client) DisassociateResolverRuleRequest(input *DisassociateResolverRuleInput) DisassociateResolverRuleRequest {
	op := &aws.Operation{
		Name:       opDisassociateResolverRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateResolverRuleInput{}
	}

	req := c.newRequest(op, input, &DisassociateResolverRuleOutput{})
	return DisassociateResolverRuleRequest{Request: req, Input: input, Copy: c.DisassociateResolverRuleRequest}
}

// DisassociateResolverRuleRequest is the request type for the
// DisassociateResolverRule API operation.
type DisassociateResolverRuleRequest struct {
	*aws.Request
	Input *DisassociateResolverRuleInput
	Copy  func(*DisassociateResolverRuleInput) DisassociateResolverRuleRequest
}

// Send marshals and sends the DisassociateResolverRule API request.
func (r DisassociateResolverRuleRequest) Send(ctx context.Context) (*DisassociateResolverRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateResolverRuleResponse{
		DisassociateResolverRuleOutput: r.Request.Data.(*DisassociateResolverRuleOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateResolverRuleResponse is the response type for the
// DisassociateResolverRule API operation.
type DisassociateResolverRuleResponse struct {
	*DisassociateResolverRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateResolverRule request.
func (r *DisassociateResolverRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
