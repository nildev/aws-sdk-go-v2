// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackInstancesInput
type ListStackInstancesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous request didn't return all of the remaining results, the response's
	// NextToken parameter value is set to a token. To retrieve the next set of
	// results, call ListStackInstances again and assign that token to the request
	// object's NextToken parameter. If there are no remaining results, the previous
	// response object's NextToken parameter is set to null.
	NextToken *string `min:"1" type:"string"`

	// The name of the AWS account that you want to list stack instances for.
	StackInstanceAccount *string `type:"string"`

	// The name of the region where you want to list stack instances.
	StackInstanceRegion *string `type:"string"`

	// The name or unique ID of the stack set that you want to list stack instances
	// for.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListStackInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStackInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStackInstancesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackInstancesOutput
type ListStackInstancesOutput struct {
	_ struct{} `type:"structure"`

	// If the request doesn't return all of the remaining results, NextToken is
	// set to a token. To retrieve the next set of results, call ListStackInstances
	// again and assign that token to the request object's NextToken parameter.
	// If the request returns all results, NextToken is set to null.
	NextToken *string `json:"cloudformation:ListStackInstancesOutput:NextToken" min:"1" type:"string"`

	// A list of StackInstanceSummary structures that contain information about
	// the specified stack instances.
	Summaries []StackInstanceSummary `json:"cloudformation:ListStackInstancesOutput:Summaries" type:"list"`
}

// String returns the string representation
func (s ListStackInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListStackInstances = "ListStackInstances"

// ListStackInstancesRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns summary information about stack instances that are associated with
// the specified stack set. You can filter for stack instances that are associated
// with a specific AWS account name or region.
//
//    // Example sending a request using ListStackInstancesRequest.
//    req := client.ListStackInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListStackInstances
func (c *Client) ListStackInstancesRequest(input *ListStackInstancesInput) ListStackInstancesRequest {
	op := &aws.Operation{
		Name:       opListStackInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListStackInstancesInput{}
	}

	req := c.newRequest(op, input, &ListStackInstancesOutput{})
	return ListStackInstancesRequest{Request: req, Input: input, Copy: c.ListStackInstancesRequest}
}

// ListStackInstancesRequest is the request type for the
// ListStackInstances API operation.
type ListStackInstancesRequest struct {
	*aws.Request
	Input *ListStackInstancesInput
	Copy  func(*ListStackInstancesInput) ListStackInstancesRequest
}

// Send marshals and sends the ListStackInstances API request.
func (r ListStackInstancesRequest) Send(ctx context.Context) (*ListStackInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStackInstancesResponse{
		ListStackInstancesOutput: r.Request.Data.(*ListStackInstancesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListStackInstancesResponse is the response type for the
// ListStackInstances API operation.
type ListStackInstancesResponse struct {
	*ListStackInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStackInstances request.
func (r *ListStackInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
