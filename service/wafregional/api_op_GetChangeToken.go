// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetChangeTokenRequest
type GetChangeTokenInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetChangeTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetChangeTokenResponse
type GetChangeTokenOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used in the request. Use this value in a GetChangeTokenStatus
	// request to get the current status of the request.
	ChangeToken *string `json:"waf-regional:GetChangeTokenOutput:ChangeToken" min:"1" type:"string"`
}

// String returns the string representation
func (s GetChangeTokenOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetChangeToken = "GetChangeToken"

// GetChangeTokenRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// When you want to create, update, or delete AWS WAF objects, get a change
// token and include the change token in the create, update, or delete request.
// Change tokens ensure that your application doesn't submit conflicting requests
// to AWS WAF.
//
// Each create, update, or delete request must use a unique change token. If
// your application submits a GetChangeToken request and then submits a second
// GetChangeToken request before submitting a create, update, or delete request,
// the second GetChangeToken request returns the same value as the first GetChangeToken
// request.
//
// When you use a change token in a create, update, or delete request, the status
// of the change token changes to PENDING, which indicates that AWS WAF is propagating
// the change to all AWS WAF servers. Use GetChangeTokenStatus to determine
// the status of your change token.
//
//    // Example sending a request using GetChangeTokenRequest.
//    req := client.GetChangeTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetChangeToken
func (c *Client) GetChangeTokenRequest(input *GetChangeTokenInput) GetChangeTokenRequest {
	op := &aws.Operation{
		Name:       opGetChangeToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetChangeTokenInput{}
	}

	req := c.newRequest(op, input, &GetChangeTokenOutput{})
	return GetChangeTokenRequest{Request: req, Input: input, Copy: c.GetChangeTokenRequest}
}

// GetChangeTokenRequest is the request type for the
// GetChangeToken API operation.
type GetChangeTokenRequest struct {
	*aws.Request
	Input *GetChangeTokenInput
	Copy  func(*GetChangeTokenInput) GetChangeTokenRequest
}

// Send marshals and sends the GetChangeToken API request.
func (r GetChangeTokenRequest) Send(ctx context.Context) (*GetChangeTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetChangeTokenResponse{
		GetChangeTokenOutput: r.Request.Data.(*GetChangeTokenOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetChangeTokenResponse is the response type for the
// GetChangeToken API operation.
type GetChangeTokenResponse struct {
	*GetChangeTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetChangeToken request.
func (r *GetChangeTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
