// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/UpdateSizeConstraintSetRequest
type UpdateSizeConstraintSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The SizeConstraintSetId of the SizeConstraintSet that you want to update.
	// SizeConstraintSetId is returned by CreateSizeConstraintSet and by ListSizeConstraintSets.
	//
	// SizeConstraintSetId is a required field
	SizeConstraintSetId *string `min:"1" type:"string" required:"true"`

	// An array of SizeConstraintSetUpdate objects that you want to insert into
	// or delete from a SizeConstraintSet. For more information, see the applicable
	// data types:
	//
	//    * SizeConstraintSetUpdate: Contains Action and SizeConstraint
	//
	//    * SizeConstraint: Contains FieldToMatch, TextTransformation, ComparisonOperator,
	//    and Size
	//
	//    * FieldToMatch: Contains Data and Type
	//
	// Updates is a required field
	Updates []SizeConstraintSetUpdate `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateSizeConstraintSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSizeConstraintSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSizeConstraintSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.SizeConstraintSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SizeConstraintSetId"))
	}
	if s.SizeConstraintSetId != nil && len(*s.SizeConstraintSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SizeConstraintSetId", 1))
	}

	if s.Updates == nil {
		invalidParams.Add(aws.NewErrParamRequired("Updates"))
	}
	if s.Updates != nil && len(s.Updates) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Updates", 1))
	}
	if s.Updates != nil {
		for i, v := range s.Updates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Updates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/UpdateSizeConstraintSetResponse
type UpdateSizeConstraintSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the UpdateSizeConstraintSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `json:"waf:UpdateSizeConstraintSetOutput:ChangeToken" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateSizeConstraintSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateSizeConstraintSet = "UpdateSizeConstraintSet"

// UpdateSizeConstraintSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Inserts or deletes SizeConstraint objects (filters) in a SizeConstraintSet.
// For each SizeConstraint object, you specify the following values:
//
//    * Whether to insert or delete the object from the array. If you want to
//    change a SizeConstraintSetUpdate object, you delete the existing object
//    and add a new one.
//
//    * The part of a web request that you want AWS WAF to evaluate, such as
//    the length of a query string or the length of the User-Agent header.
//
//    * Whether to perform any transformations on the request, such as converting
//    it to lowercase, before checking its length. Note that transformations
//    of the request body are not supported because the AWS resource forwards
//    only the first 8192 bytes of your request to AWS WAF. You can only specify
//    a single type of TextTransformation.
//
//    * A ComparisonOperator used for evaluating the selected part of the request
//    against the specified Size, such as equals, greater than, less than, and
//    so on.
//
//    * The length, in bytes, that you want AWS WAF to watch for in selected
//    part of the request. The length is computed after applying the transformation.
//
// For example, you can add a SizeConstraintSetUpdate object that matches web
// requests in which the length of the User-Agent header is greater than 100
// bytes. You can then configure AWS WAF to block those requests.
//
// To create and configure a SizeConstraintSet, perform the following steps:
//
// Create a SizeConstraintSet. For more information, see CreateSizeConstraintSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateSizeConstraintSet request.
//
// Submit an UpdateSizeConstraintSet request to specify the part of the request
// that you want AWS WAF to inspect (for example, the header or the URI) and
// the value that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateSizeConstraintSetRequest.
//    req := client.UpdateSizeConstraintSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/UpdateSizeConstraintSet
func (c *Client) UpdateSizeConstraintSetRequest(input *UpdateSizeConstraintSetInput) UpdateSizeConstraintSetRequest {
	op := &aws.Operation{
		Name:       opUpdateSizeConstraintSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSizeConstraintSetInput{}
	}

	req := c.newRequest(op, input, &UpdateSizeConstraintSetOutput{})
	return UpdateSizeConstraintSetRequest{Request: req, Input: input, Copy: c.UpdateSizeConstraintSetRequest}
}

// UpdateSizeConstraintSetRequest is the request type for the
// UpdateSizeConstraintSet API operation.
type UpdateSizeConstraintSetRequest struct {
	*aws.Request
	Input *UpdateSizeConstraintSetInput
	Copy  func(*UpdateSizeConstraintSetInput) UpdateSizeConstraintSetRequest
}

// Send marshals and sends the UpdateSizeConstraintSet API request.
func (r UpdateSizeConstraintSetRequest) Send(ctx context.Context) (*UpdateSizeConstraintSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSizeConstraintSetResponse{
		UpdateSizeConstraintSetOutput: r.Request.Data.(*UpdateSizeConstraintSetOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSizeConstraintSetResponse is the response type for the
// UpdateSizeConstraintSet API operation.
type UpdateSizeConstraintSetResponse struct {
	*UpdateSizeConstraintSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSizeConstraintSet request.
func (r *UpdateSizeConstraintSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
