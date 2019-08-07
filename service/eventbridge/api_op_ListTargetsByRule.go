// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/ListTargetsByRuleRequest
type ListTargetsByRuleInput struct {
	_ struct{} `type:"structure"`

	// The event bus associated with the rule. If you omit this, the default event
	// bus is used.
	EventBusName *string `min:"1" type:"string"`

	// The maximum number of results to return.
	Limit *int64 `min:"1" type:"integer"`

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`

	// The name of the rule.
	//
	// Rule is a required field
	Rule *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTargetsByRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTargetsByRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTargetsByRuleInput"}
	if s.EventBusName != nil && len(*s.EventBusName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventBusName", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.Rule == nil {
		invalidParams.Add(aws.NewErrParamRequired("Rule"))
	}
	if s.Rule != nil && len(*s.Rule) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Rule", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/ListTargetsByRuleResponse
type ListTargetsByRuleOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether there are additional results to retrieve. If there are
	// no more results, the value is null.
	NextToken *string `json:"events:ListTargetsByRuleOutput:NextToken" min:"1" type:"string"`

	// The targets assigned to the rule.
	Targets []Target `json:"events:ListTargetsByRuleOutput:Targets" min:"1" type:"list"`
}

// String returns the string representation
func (s ListTargetsByRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTargetsByRule = "ListTargetsByRule"

// ListTargetsByRuleRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Lists the targets assigned to the specified rule.
//
//    // Example sending a request using ListTargetsByRuleRequest.
//    req := client.ListTargetsByRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/ListTargetsByRule
func (c *Client) ListTargetsByRuleRequest(input *ListTargetsByRuleInput) ListTargetsByRuleRequest {
	op := &aws.Operation{
		Name:       opListTargetsByRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTargetsByRuleInput{}
	}

	req := c.newRequest(op, input, &ListTargetsByRuleOutput{})
	return ListTargetsByRuleRequest{Request: req, Input: input, Copy: c.ListTargetsByRuleRequest}
}

// ListTargetsByRuleRequest is the request type for the
// ListTargetsByRule API operation.
type ListTargetsByRuleRequest struct {
	*aws.Request
	Input *ListTargetsByRuleInput
	Copy  func(*ListTargetsByRuleInput) ListTargetsByRuleRequest
}

// Send marshals and sends the ListTargetsByRule API request.
func (r ListTargetsByRuleRequest) Send(ctx context.Context) (*ListTargetsByRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTargetsByRuleResponse{
		ListTargetsByRuleOutput: r.Request.Data.(*ListTargetsByRuleOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTargetsByRuleResponse is the response type for the
// ListTargetsByRule API operation.
type ListTargetsByRuleResponse struct {
	*ListTargetsByRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTargetsByRule request.
func (r *ListTargetsByRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
