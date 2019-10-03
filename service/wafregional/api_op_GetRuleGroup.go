// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetRuleGroupRequest
type GetRuleGroupInput struct {
	_ struct{} `type:"structure"`

	// The RuleGroupId of the RuleGroup that you want to get. RuleGroupId is returned
	// by CreateRuleGroup and by ListRuleGroups.
	//
	// RuleGroupId is a required field
	RuleGroupId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRuleGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRuleGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRuleGroupInput"}

	if s.RuleGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleGroupId"))
	}
	if s.RuleGroupId != nil && len(*s.RuleGroupId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleGroupId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetRuleGroupResponse
type GetRuleGroupOutput struct {
	_ struct{} `type:"structure"`

	// Information about the RuleGroup that you specified in the GetRuleGroup request.
	RuleGroup *waf.RuleGroup `json:"waf-regional:GetRuleGroupOutput:RuleGroup" type:"structure"`
}

// String returns the string representation
func (s GetRuleGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRuleGroup = "GetRuleGroup"

// GetRuleGroupRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the RuleGroup that is specified by the RuleGroupId that you included
// in the GetRuleGroup request.
//
// To view the rules in a rule group, use ListActivatedRulesInRuleGroup.
//
//    // Example sending a request using GetRuleGroupRequest.
//    req := client.GetRuleGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetRuleGroup
func (c *Client) GetRuleGroupRequest(input *GetRuleGroupInput) GetRuleGroupRequest {
	op := &aws.Operation{
		Name:       opGetRuleGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRuleGroupInput{}
	}

	req := c.newRequest(op, input, &GetRuleGroupOutput{})
	return GetRuleGroupRequest{Request: req, Input: input, Copy: c.GetRuleGroupRequest}
}

// GetRuleGroupRequest is the request type for the
// GetRuleGroup API operation.
type GetRuleGroupRequest struct {
	*aws.Request
	Input *GetRuleGroupInput
	Copy  func(*GetRuleGroupInput) GetRuleGroupRequest
}

// Send marshals and sends the GetRuleGroup API request.
func (r GetRuleGroupRequest) Send(ctx context.Context) (*GetRuleGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRuleGroupResponse{
		GetRuleGroupOutput: r.Request.Data.(*GetRuleGroupOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRuleGroupResponse is the response type for the
// GetRuleGroup API operation.
type GetRuleGroupResponse struct {
	*GetRuleGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRuleGroup request.
func (r *GetRuleGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
