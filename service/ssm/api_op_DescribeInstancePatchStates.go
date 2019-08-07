// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeInstancePatchStatesRequest
type DescribeInstancePatchStatesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the instance whose patch state information should be retrieved.
	//
	// InstanceIds is a required field
	InstanceIds []string `type:"list" required:"true"`

	// The maximum number of instances to return (per page).
	MaxResults *int64 `min:"10" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstancePatchStatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstancePatchStatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInstancePatchStatesInput"}

	if s.InstanceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceIds"))
	}
	if s.MaxResults != nil && *s.MaxResults < 10 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 10))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeInstancePatchStatesResult
type DescribeInstancePatchStatesOutput struct {
	_ struct{} `type:"structure"`

	// The high-level patch state for the requested instances.
	InstancePatchStates []InstancePatchState `json:"ssm:DescribeInstancePatchStatesOutput:InstancePatchStates" type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `json:"ssm:DescribeInstancePatchStatesOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s DescribeInstancePatchStatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInstancePatchStates = "DescribeInstancePatchStates"

// DescribeInstancePatchStatesRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves the high-level patch state of one or more instances.
//
//    // Example sending a request using DescribeInstancePatchStatesRequest.
//    req := client.DescribeInstancePatchStatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeInstancePatchStates
func (c *Client) DescribeInstancePatchStatesRequest(input *DescribeInstancePatchStatesInput) DescribeInstancePatchStatesRequest {
	op := &aws.Operation{
		Name:       opDescribeInstancePatchStates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstancePatchStatesInput{}
	}

	req := c.newRequest(op, input, &DescribeInstancePatchStatesOutput{})
	return DescribeInstancePatchStatesRequest{Request: req, Input: input, Copy: c.DescribeInstancePatchStatesRequest}
}

// DescribeInstancePatchStatesRequest is the request type for the
// DescribeInstancePatchStates API operation.
type DescribeInstancePatchStatesRequest struct {
	*aws.Request
	Input *DescribeInstancePatchStatesInput
	Copy  func(*DescribeInstancePatchStatesInput) DescribeInstancePatchStatesRequest
}

// Send marshals and sends the DescribeInstancePatchStates API request.
func (r DescribeInstancePatchStatesRequest) Send(ctx context.Context) (*DescribeInstancePatchStatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstancePatchStatesResponse{
		DescribeInstancePatchStatesOutput: r.Request.Data.(*DescribeInstancePatchStatesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInstancePatchStatesResponse is the response type for the
// DescribeInstancePatchStates API operation.
type DescribeInstancePatchStatesResponse struct {
	*DescribeInstancePatchStatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstancePatchStates request.
func (r *DescribeInstancePatchStatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
