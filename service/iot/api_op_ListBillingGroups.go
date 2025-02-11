// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListBillingGroupsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return per request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Limit the results to billing groups whose names have the given prefix.
	NamePrefixFilter *string `location:"querystring" locationName:"namePrefixFilter" min:"1" type:"string"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBillingGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBillingGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBillingGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NamePrefixFilter != nil && len(*s.NamePrefixFilter) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NamePrefixFilter", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBillingGroupsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NamePrefixFilter != nil {
		v := *s.NamePrefixFilter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "namePrefixFilter", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListBillingGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The list of billing groups.
	BillingGroups []GroupNameAndArn `locationName:"billingGroups" type:"list"`

	// The token used to get the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBillingGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBillingGroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BillingGroups != nil {
		v := s.BillingGroups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "billingGroups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListBillingGroups = "ListBillingGroups"

// ListBillingGroupsRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the billing groups you have created.
//
//    // Example sending a request using ListBillingGroupsRequest.
//    req := client.ListBillingGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListBillingGroupsRequest(input *ListBillingGroupsInput) ListBillingGroupsRequest {
	op := &aws.Operation{
		Name:       opListBillingGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/billing-groups",
	}

	if input == nil {
		input = &ListBillingGroupsInput{}
	}

	req := c.newRequest(op, input, &ListBillingGroupsOutput{})
	return ListBillingGroupsRequest{Request: req, Input: input, Copy: c.ListBillingGroupsRequest}
}

// ListBillingGroupsRequest is the request type for the
// ListBillingGroups API operation.
type ListBillingGroupsRequest struct {
	*aws.Request
	Input *ListBillingGroupsInput
	Copy  func(*ListBillingGroupsInput) ListBillingGroupsRequest
}

// Send marshals and sends the ListBillingGroups API request.
func (r ListBillingGroupsRequest) Send(ctx context.Context) (*ListBillingGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBillingGroupsResponse{
		ListBillingGroupsOutput: r.Request.Data.(*ListBillingGroupsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListBillingGroupsResponse is the response type for the
// ListBillingGroups API operation.
type ListBillingGroupsResponse struct {
	*ListBillingGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBillingGroups request.
func (r *ListBillingGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
