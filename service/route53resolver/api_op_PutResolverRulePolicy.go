// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/PutResolverRulePolicyRequest
type PutResolverRulePolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the account that you want to grant permissions
	// to.
	//
	// Arn is a required field
	Arn *string `min:"1" type:"string" required:"true"`

	// An AWS Identity and Access Management policy statement that lists the permissions
	// that you want to grant to another AWS account.
	//
	// ResolverRulePolicy is a required field
	ResolverRulePolicy *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutResolverRulePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutResolverRulePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutResolverRulePolicyInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 1))
	}

	if s.ResolverRulePolicy == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverRulePolicy"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response to a PutResolverRulePolicy request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/PutResolverRulePolicyResponse
type PutResolverRulePolicyOutput struct {
	_ struct{} `type:"structure"`

	// Whether the PutResolverRulePolicy request was successful.
	ReturnValue *bool `json:"route53resolver:PutResolverRulePolicyOutput:ReturnValue" type:"boolean"`
}

// String returns the string representation
func (s PutResolverRulePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutResolverRulePolicy = "PutResolverRulePolicy"

// PutResolverRulePolicyRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Specifies the Resolver operations and resources that you want to allow another
// AWS account to be able to use.
//
//    // Example sending a request using PutResolverRulePolicyRequest.
//    req := client.PutResolverRulePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/PutResolverRulePolicy
func (c *Client) PutResolverRulePolicyRequest(input *PutResolverRulePolicyInput) PutResolverRulePolicyRequest {
	op := &aws.Operation{
		Name:       opPutResolverRulePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutResolverRulePolicyInput{}
	}

	req := c.newRequest(op, input, &PutResolverRulePolicyOutput{})
	return PutResolverRulePolicyRequest{Request: req, Input: input, Copy: c.PutResolverRulePolicyRequest}
}

// PutResolverRulePolicyRequest is the request type for the
// PutResolverRulePolicy API operation.
type PutResolverRulePolicyRequest struct {
	*aws.Request
	Input *PutResolverRulePolicyInput
	Copy  func(*PutResolverRulePolicyInput) PutResolverRulePolicyRequest
}

// Send marshals and sends the PutResolverRulePolicy API request.
func (r PutResolverRulePolicyRequest) Send(ctx context.Context) (*PutResolverRulePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutResolverRulePolicyResponse{
		PutResolverRulePolicyOutput: r.Request.Data.(*PutResolverRulePolicyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutResolverRulePolicyResponse is the response type for the
// PutResolverRulePolicy API operation.
type PutResolverRulePolicyResponse struct {
	*PutResolverRulePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutResolverRulePolicy request.
func (r *PutResolverRulePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
