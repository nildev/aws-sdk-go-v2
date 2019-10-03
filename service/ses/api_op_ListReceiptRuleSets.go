// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to list the receipt rule sets that exist under your
// AWS account. You use receipt rule sets to receive email with Amazon SES.
// For more information, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListReceiptRuleSetsRequest
type ListReceiptRuleSetsInput struct {
	_ struct{} `type:"structure"`

	// A token returned from a previous call to ListReceiptRuleSets to indicate
	// the position in the receipt rule set list.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListReceiptRuleSetsInput) String() string {
	return awsutil.Prettify(s)
}

// A list of receipt rule sets that exist under your AWS account.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListReceiptRuleSetsResponse
type ListReceiptRuleSetsOutput struct {
	_ struct{} `type:"structure"`

	// A token indicating that there are additional receipt rule sets available
	// to be listed. Pass this token to successive calls of ListReceiptRuleSets
	// to retrieve up to 100 receipt rule sets at a time.
	NextToken *string `json:"email:ListReceiptRuleSetsOutput:NextToken" type:"string"`

	// The metadata for the currently active receipt rule set. The metadata consists
	// of the rule set name and the timestamp of when the rule set was created.
	RuleSets []ReceiptRuleSetMetadata `json:"email:ListReceiptRuleSetsOutput:RuleSets" type:"list"`
}

// String returns the string representation
func (s ListReceiptRuleSetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListReceiptRuleSets = "ListReceiptRuleSets"

// ListReceiptRuleSetsRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Lists the receipt rule sets that exist under your AWS account in the current
// AWS Region. If there are additional receipt rule sets to be retrieved, you
// will receive a NextToken that you can provide to the next call to ListReceiptRuleSets
// to retrieve the additional entries.
//
// For information about managing receipt rule sets, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rule-sets.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using ListReceiptRuleSetsRequest.
//    req := client.ListReceiptRuleSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListReceiptRuleSets
func (c *Client) ListReceiptRuleSetsRequest(input *ListReceiptRuleSetsInput) ListReceiptRuleSetsRequest {
	op := &aws.Operation{
		Name:       opListReceiptRuleSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListReceiptRuleSetsInput{}
	}

	req := c.newRequest(op, input, &ListReceiptRuleSetsOutput{})
	return ListReceiptRuleSetsRequest{Request: req, Input: input, Copy: c.ListReceiptRuleSetsRequest}
}

// ListReceiptRuleSetsRequest is the request type for the
// ListReceiptRuleSets API operation.
type ListReceiptRuleSetsRequest struct {
	*aws.Request
	Input *ListReceiptRuleSetsInput
	Copy  func(*ListReceiptRuleSetsInput) ListReceiptRuleSetsRequest
}

// Send marshals and sends the ListReceiptRuleSets API request.
func (r ListReceiptRuleSetsRequest) Send(ctx context.Context) (*ListReceiptRuleSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListReceiptRuleSetsResponse{
		ListReceiptRuleSetsOutput: r.Request.Data.(*ListReceiptRuleSetsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListReceiptRuleSetsResponse is the response type for the
// ListReceiptRuleSets API operation.
type ListReceiptRuleSetsResponse struct {
	*ListReceiptRuleSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListReceiptRuleSets request.
func (r *ListReceiptRuleSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
