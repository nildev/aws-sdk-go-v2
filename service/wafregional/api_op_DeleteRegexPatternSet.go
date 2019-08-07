// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/DeleteRegexPatternSetRequest
type DeleteRegexPatternSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The RegexPatternSetId of the RegexPatternSet that you want to delete. RegexPatternSetId
	// is returned by CreateRegexPatternSet and by ListRegexPatternSets.
	//
	// RegexPatternSetId is a required field
	RegexPatternSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRegexPatternSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRegexPatternSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRegexPatternSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.RegexPatternSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegexPatternSetId"))
	}
	if s.RegexPatternSetId != nil && len(*s.RegexPatternSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RegexPatternSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/DeleteRegexPatternSetResponse
type DeleteRegexPatternSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the DeleteRegexPatternSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `json:"waf-regional:DeleteRegexPatternSetOutput:ChangeToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteRegexPatternSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRegexPatternSet = "DeleteRegexPatternSet"

// DeleteRegexPatternSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Permanently deletes a RegexPatternSet. You can't delete a RegexPatternSet
// if it's still used in any RegexMatchSet or if the RegexPatternSet is not
// empty.
//
//    // Example sending a request using DeleteRegexPatternSetRequest.
//    req := client.DeleteRegexPatternSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/DeleteRegexPatternSet
func (c *Client) DeleteRegexPatternSetRequest(input *DeleteRegexPatternSetInput) DeleteRegexPatternSetRequest {
	op := &aws.Operation{
		Name:       opDeleteRegexPatternSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRegexPatternSetInput{}
	}

	req := c.newRequest(op, input, &DeleteRegexPatternSetOutput{})
	return DeleteRegexPatternSetRequest{Request: req, Input: input, Copy: c.DeleteRegexPatternSetRequest}
}

// DeleteRegexPatternSetRequest is the request type for the
// DeleteRegexPatternSet API operation.
type DeleteRegexPatternSetRequest struct {
	*aws.Request
	Input *DeleteRegexPatternSetInput
	Copy  func(*DeleteRegexPatternSetInput) DeleteRegexPatternSetRequest
}

// Send marshals and sends the DeleteRegexPatternSet API request.
func (r DeleteRegexPatternSetRequest) Send(ctx context.Context) (*DeleteRegexPatternSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRegexPatternSetResponse{
		DeleteRegexPatternSetOutput: r.Request.Data.(*DeleteRegexPatternSetOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRegexPatternSetResponse is the response type for the
// DeleteRegexPatternSet API operation.
type DeleteRegexPatternSetResponse struct {
	*DeleteRegexPatternSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRegexPatternSet request.
func (r *DeleteRegexPatternSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
