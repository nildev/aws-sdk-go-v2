// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/ListPartnerEventSourceAccountsRequest
type ListPartnerEventSourceAccountsInput struct {
	_ struct{} `type:"structure"`

	// The name of the partner event source to display account information about.
	//
	// EventSourceName is a required field
	EventSourceName *string `min:"1" type:"string" required:"true"`

	// Specifying this limits the number of results returned by this operation.
	// The operation also returns a NextToken that you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit *int64 `min:"1" type:"integer"`

	// The token returned by a previous call to this operation. Specifying this
	// retrieves the next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListPartnerEventSourceAccountsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPartnerEventSourceAccountsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPartnerEventSourceAccountsInput"}

	if s.EventSourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventSourceName"))
	}
	if s.EventSourceName != nil && len(*s.EventSourceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventSourceName", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/ListPartnerEventSourceAccountsResponse
type ListPartnerEventSourceAccountsOutput struct {
	_ struct{} `type:"structure"`

	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string `json:"events:ListPartnerEventSourceAccountsOutput:NextToken" min:"1" type:"string"`

	// The list of partner event sources returned by the operation.
	PartnerEventSourceAccounts []PartnerEventSourceAccount `json:"events:ListPartnerEventSourceAccountsOutput:PartnerEventSourceAccounts" type:"list"`
}

// String returns the string representation
func (s ListPartnerEventSourceAccountsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPartnerEventSourceAccounts = "ListPartnerEventSourceAccounts"

// ListPartnerEventSourceAccountsRequest returns a request value for making API operation for
// Amazon CloudWatch Events.
//
// An SaaS partner can use this operation to display the AWS account ID that
// a particular partner event source name is associated with.
//
// This operation is used by SaaS partners, not by AWS customers.
//
//    // Example sending a request using ListPartnerEventSourceAccountsRequest.
//    req := client.ListPartnerEventSourceAccountsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/ListPartnerEventSourceAccounts
func (c *Client) ListPartnerEventSourceAccountsRequest(input *ListPartnerEventSourceAccountsInput) ListPartnerEventSourceAccountsRequest {
	op := &aws.Operation{
		Name:       opListPartnerEventSourceAccounts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListPartnerEventSourceAccountsInput{}
	}

	req := c.newRequest(op, input, &ListPartnerEventSourceAccountsOutput{})
	return ListPartnerEventSourceAccountsRequest{Request: req, Input: input, Copy: c.ListPartnerEventSourceAccountsRequest}
}

// ListPartnerEventSourceAccountsRequest is the request type for the
// ListPartnerEventSourceAccounts API operation.
type ListPartnerEventSourceAccountsRequest struct {
	*aws.Request
	Input *ListPartnerEventSourceAccountsInput
	Copy  func(*ListPartnerEventSourceAccountsInput) ListPartnerEventSourceAccountsRequest
}

// Send marshals and sends the ListPartnerEventSourceAccounts API request.
func (r ListPartnerEventSourceAccountsRequest) Send(ctx context.Context) (*ListPartnerEventSourceAccountsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPartnerEventSourceAccountsResponse{
		ListPartnerEventSourceAccountsOutput: r.Request.Data.(*ListPartnerEventSourceAccountsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPartnerEventSourceAccountsResponse is the response type for the
// ListPartnerEventSourceAccounts API operation.
type ListPartnerEventSourceAccountsResponse struct {
	*ListPartnerEventSourceAccountsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPartnerEventSourceAccounts request.
func (r *ListPartnerEventSourceAccountsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
