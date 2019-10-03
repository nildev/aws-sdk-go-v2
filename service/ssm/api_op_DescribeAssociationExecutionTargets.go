// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeAssociationExecutionTargetsRequest
type DescribeAssociationExecutionTargetsInput struct {
	_ struct{} `type:"structure"`

	// The association ID that includes the execution for which you want to view
	// details.
	//
	// AssociationId is a required field
	AssociationId *string `type:"string" required:"true"`

	// The execution ID for which you want to view details.
	//
	// ExecutionId is a required field
	ExecutionId *string `type:"string" required:"true"`

	// Filters for the request. You can specify the following filters and values.
	//
	// Status (EQUAL)
	//
	// ResourceId (EQUAL)
	//
	// ResourceType (EQUAL)
	Filters []AssociationExecutionTargetsFilter `min:"1" type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAssociationExecutionTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAssociationExecutionTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAssociationExecutionTargetsInput"}

	if s.AssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationId"))
	}

	if s.ExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExecutionId"))
	}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeAssociationExecutionTargetsResult
type DescribeAssociationExecutionTargetsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the execution.
	AssociationExecutionTargets []AssociationExecutionTarget `json:"ssm:DescribeAssociationExecutionTargetsOutput:AssociationExecutionTargets" type:"list"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `json:"ssm:DescribeAssociationExecutionTargetsOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s DescribeAssociationExecutionTargetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAssociationExecutionTargets = "DescribeAssociationExecutionTargets"

// DescribeAssociationExecutionTargetsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Use this API action to view information about a specific execution of a specific
// association.
//
//    // Example sending a request using DescribeAssociationExecutionTargetsRequest.
//    req := client.DescribeAssociationExecutionTargetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeAssociationExecutionTargets
func (c *Client) DescribeAssociationExecutionTargetsRequest(input *DescribeAssociationExecutionTargetsInput) DescribeAssociationExecutionTargetsRequest {
	op := &aws.Operation{
		Name:       opDescribeAssociationExecutionTargets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAssociationExecutionTargetsInput{}
	}

	req := c.newRequest(op, input, &DescribeAssociationExecutionTargetsOutput{})
	return DescribeAssociationExecutionTargetsRequest{Request: req, Input: input, Copy: c.DescribeAssociationExecutionTargetsRequest}
}

// DescribeAssociationExecutionTargetsRequest is the request type for the
// DescribeAssociationExecutionTargets API operation.
type DescribeAssociationExecutionTargetsRequest struct {
	*aws.Request
	Input *DescribeAssociationExecutionTargetsInput
	Copy  func(*DescribeAssociationExecutionTargetsInput) DescribeAssociationExecutionTargetsRequest
}

// Send marshals and sends the DescribeAssociationExecutionTargets API request.
func (r DescribeAssociationExecutionTargetsRequest) Send(ctx context.Context) (*DescribeAssociationExecutionTargetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAssociationExecutionTargetsResponse{
		DescribeAssociationExecutionTargetsOutput: r.Request.Data.(*DescribeAssociationExecutionTargetsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAssociationExecutionTargetsResponse is the response type for the
// DescribeAssociationExecutionTargets API operation.
type DescribeAssociationExecutionTargetsResponse struct {
	*DescribeAssociationExecutionTargetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAssociationExecutionTargets request.
func (r *DescribeAssociationExecutionTargetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
