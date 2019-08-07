// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetPermissionPolicyRequest
type GetPermissionPolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the RuleGroup for which you want to get
	// the policy.
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPermissionPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPermissionPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPermissionPolicyInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetPermissionPolicyResponse
type GetPermissionPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The IAM policy attached to the specified RuleGroup.
	Policy *string `json:"waf-regional:GetPermissionPolicyOutput:Policy" min:"1" type:"string"`
}

// String returns the string representation
func (s GetPermissionPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPermissionPolicy = "GetPermissionPolicy"

// GetPermissionPolicyRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the IAM policy attached to the RuleGroup.
//
//    // Example sending a request using GetPermissionPolicyRequest.
//    req := client.GetPermissionPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetPermissionPolicy
func (c *Client) GetPermissionPolicyRequest(input *GetPermissionPolicyInput) GetPermissionPolicyRequest {
	op := &aws.Operation{
		Name:       opGetPermissionPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPermissionPolicyInput{}
	}

	req := c.newRequest(op, input, &GetPermissionPolicyOutput{})
	return GetPermissionPolicyRequest{Request: req, Input: input, Copy: c.GetPermissionPolicyRequest}
}

// GetPermissionPolicyRequest is the request type for the
// GetPermissionPolicy API operation.
type GetPermissionPolicyRequest struct {
	*aws.Request
	Input *GetPermissionPolicyInput
	Copy  func(*GetPermissionPolicyInput) GetPermissionPolicyRequest
}

// Send marshals and sends the GetPermissionPolicy API request.
func (r GetPermissionPolicyRequest) Send(ctx context.Context) (*GetPermissionPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPermissionPolicyResponse{
		GetPermissionPolicyOutput: r.Request.Data.(*GetPermissionPolicyOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPermissionPolicyResponse is the response type for the
// GetPermissionPolicy API operation.
type GetPermissionPolicyResponse struct {
	*GetPermissionPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPermissionPolicy request.
func (r *GetPermissionPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
