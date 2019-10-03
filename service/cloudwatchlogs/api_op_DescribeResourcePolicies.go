// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DescribeResourcePoliciesRequest
type DescribeResourcePoliciesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of resource policies to be displayed with one call of
	// this API.
	Limit *int64 `locationName:"limit" min:"1" type:"integer"`

	// The token for the next set of items to return. The token expires after 24
	// hours.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeResourcePoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeResourcePoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeResourcePoliciesInput"}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DescribeResourcePoliciesResponse
type DescribeResourcePoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of items to return. The token expires after 24
	// hours.
	NextToken *string `json:"logs:DescribeResourcePoliciesOutput:NextToken" locationName:"nextToken" min:"1" type:"string"`

	// The resource policies that exist in this account.
	ResourcePolicies []ResourcePolicy `json:"logs:DescribeResourcePoliciesOutput:ResourcePolicies" locationName:"resourcePolicies" type:"list"`
}

// String returns the string representation
func (s DescribeResourcePoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeResourcePolicies = "DescribeResourcePolicies"

// DescribeResourcePoliciesRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Lists the resource policies in this account.
//
//    // Example sending a request using DescribeResourcePoliciesRequest.
//    req := client.DescribeResourcePoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DescribeResourcePolicies
func (c *Client) DescribeResourcePoliciesRequest(input *DescribeResourcePoliciesInput) DescribeResourcePoliciesRequest {
	op := &aws.Operation{
		Name:       opDescribeResourcePolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeResourcePoliciesInput{}
	}

	req := c.newRequest(op, input, &DescribeResourcePoliciesOutput{})
	return DescribeResourcePoliciesRequest{Request: req, Input: input, Copy: c.DescribeResourcePoliciesRequest}
}

// DescribeResourcePoliciesRequest is the request type for the
// DescribeResourcePolicies API operation.
type DescribeResourcePoliciesRequest struct {
	*aws.Request
	Input *DescribeResourcePoliciesInput
	Copy  func(*DescribeResourcePoliciesInput) DescribeResourcePoliciesRequest
}

// Send marshals and sends the DescribeResourcePolicies API request.
func (r DescribeResourcePoliciesRequest) Send(ctx context.Context) (*DescribeResourcePoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeResourcePoliciesResponse{
		DescribeResourcePoliciesOutput: r.Request.Data.(*DescribeResourcePoliciesOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeResourcePoliciesResponse is the response type for the
// DescribeResourcePolicies API operation.
type DescribeResourcePoliciesResponse struct {
	*DescribeResourcePoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeResourcePolicies request.
func (r *DescribeResourcePoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
