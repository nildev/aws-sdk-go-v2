// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackSetInput
type DescribeStackSetInput struct {
	_ struct{} `type:"structure"`

	// The name or unique ID of the stack set whose description you want.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStackSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackSetInput"}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackSetOutput
type DescribeStackSetOutput struct {
	_ struct{} `type:"structure"`

	// The specified stack set.
	StackSet *StackSet `json:"cloudformation:DescribeStackSetOutput:StackSet" type:"structure"`
}

// String returns the string representation
func (s DescribeStackSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStackSet = "DescribeStackSet"

// DescribeStackSetRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns the description of the specified stack set.
//
//    // Example sending a request using DescribeStackSetRequest.
//    req := client.DescribeStackSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackSet
func (c *Client) DescribeStackSetRequest(input *DescribeStackSetInput) DescribeStackSetRequest {
	op := &aws.Operation{
		Name:       opDescribeStackSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStackSetInput{}
	}

	req := c.newRequest(op, input, &DescribeStackSetOutput{})
	return DescribeStackSetRequest{Request: req, Input: input, Copy: c.DescribeStackSetRequest}
}

// DescribeStackSetRequest is the request type for the
// DescribeStackSet API operation.
type DescribeStackSetRequest struct {
	*aws.Request
	Input *DescribeStackSetInput
	Copy  func(*DescribeStackSetInput) DescribeStackSetRequest
}

// Send marshals and sends the DescribeStackSet API request.
func (r DescribeStackSetRequest) Send(ctx context.Context) (*DescribeStackSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStackSetResponse{
		DescribeStackSetOutput: r.Request.Data.(*DescribeStackSetOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStackSetResponse is the response type for the
// DescribeStackSet API operation.
type DescribeStackSetResponse struct {
	*DescribeStackSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStackSet request.
func (r *DescribeStackSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
