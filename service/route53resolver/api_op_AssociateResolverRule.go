// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/AssociateResolverRuleRequest
type AssociateResolverRuleInput struct {
	_ struct{} `type:"structure"`

	// A name for the association that you're creating between a resolver rule and
	// a VPC.
	Name *string `type:"string"`

	// The ID of the resolver rule that you want to associate with the VPC. To list
	// the existing resolver rules, use ListResolverRules.
	//
	// ResolverRuleId is a required field
	ResolverRuleId *string `min:"1" type:"string" required:"true"`

	// The ID of the VPC that you want to associate the resolver rule with.
	//
	// VPCId is a required field
	VPCId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateResolverRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateResolverRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateResolverRuleInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/AssociateResolverRuleResponse
type AssociateResolverRuleOutput struct {
	_ struct{} `type:"structure"`

	// Information about the AssociateResolverRule request, including the status
	// of the request.
	ResolverRuleAssociation *ResolverRuleAssociation `json:"route53resolver:AssociateResolverRuleOutput:ResolverRuleAssociation" type:"structure"`
}

// String returns the string representation
func (s AssociateResolverRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateResolverRule = "AssociateResolverRule"

// AssociateResolverRuleRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Associates a resolver rule with a VPC. When you associate a rule with a VPC,
// Resolver forwards all DNS queries for the domain name that is specified in
// the rule and that originate in the VPC. The queries are forwarded to the
// IP addresses for the DNS resolvers that are specified in the rule. For more
// information about rules, see CreateResolverRule.
//
//    // Example sending a request using AssociateResolverRuleRequest.
//    req := client.AssociateResolverRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/AssociateResolverRule
func (c *Client) AssociateResolverRuleRequest(input *AssociateResolverRuleInput) AssociateResolverRuleRequest {
	op := &aws.Operation{
		Name:       opAssociateResolverRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateResolverRuleInput{}
	}

	req := c.newRequest(op, input, &AssociateResolverRuleOutput{})
	return AssociateResolverRuleRequest{Request: req, Input: input, Copy: c.AssociateResolverRuleRequest}
}

// AssociateResolverRuleRequest is the request type for the
// AssociateResolverRule API operation.
type AssociateResolverRuleRequest struct {
	*aws.Request
	Input *AssociateResolverRuleInput
	Copy  func(*AssociateResolverRuleInput) AssociateResolverRuleRequest
}

// Send marshals and sends the AssociateResolverRule API request.
func (r AssociateResolverRuleRequest) Send(ctx context.Context) (*AssociateResolverRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateResolverRuleResponse{
		AssociateResolverRuleOutput: r.Request.Data.(*AssociateResolverRuleOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateResolverRuleResponse is the response type for the
// AssociateResolverRule API operation.
type AssociateResolverRuleResponse struct {
	*AssociateResolverRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateResolverRule request.
func (r *AssociateResolverRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
