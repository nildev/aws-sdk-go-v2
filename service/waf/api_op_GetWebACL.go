// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetWebACLRequest
type GetWebACLInput struct {
	_ struct{} `type:"structure"`

	// The WebACLId of the WebACL that you want to get. WebACLId is returned by
	// CreateWebACL and by ListWebACLs.
	//
	// WebACLId is a required field
	WebACLId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetWebACLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetWebACLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetWebACLInput"}

	if s.WebACLId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLId"))
	}
	if s.WebACLId != nil && len(*s.WebACLId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebACLId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetWebACLResponse
type GetWebACLOutput struct {
	_ struct{} `type:"structure"`

	// Information about the WebACL that you specified in the GetWebACL request.
	// For more information, see the following topics:
	//
	//    * WebACL: Contains DefaultAction, MetricName, Name, an array of Rule objects,
	//    and WebACLId
	//
	//    * DefaultAction (Data type is WafAction): Contains Type
	//
	//    * Rules: Contains an array of ActivatedRule objects, which contain Action,
	//    Priority, and RuleId
	//
	//    * Action: Contains Type
	WebACL *WebACL `json:"waf:GetWebACLOutput:WebACL" type:"structure"`
}

// String returns the string representation
func (s GetWebACLOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetWebACL = "GetWebACL"

// GetWebACLRequest returns a request value for making API operation for
// AWS WAF.
//
// Returns the WebACL that is specified by WebACLId.
//
//    // Example sending a request using GetWebACLRequest.
//    req := client.GetWebACLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetWebACL
func (c *Client) GetWebACLRequest(input *GetWebACLInput) GetWebACLRequest {
	op := &aws.Operation{
		Name:       opGetWebACL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetWebACLInput{}
	}

	req := c.newRequest(op, input, &GetWebACLOutput{})
	return GetWebACLRequest{Request: req, Input: input, Copy: c.GetWebACLRequest}
}

// GetWebACLRequest is the request type for the
// GetWebACL API operation.
type GetWebACLRequest struct {
	*aws.Request
	Input *GetWebACLInput
	Copy  func(*GetWebACLInput) GetWebACLRequest
}

// Send marshals and sends the GetWebACL API request.
func (r GetWebACLRequest) Send(ctx context.Context) (*GetWebACLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetWebACLResponse{
		GetWebACLOutput: r.Request.Data.(*GetWebACLOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetWebACLResponse is the response type for the
// GetWebACL API operation.
type GetWebACLResponse struct {
	*GetWebACLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetWebACL request.
func (r *GetWebACLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
