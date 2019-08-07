// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloud9

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DescribeEnvironmentsRequest
type DescribeEnvironmentsInput struct {
	_ struct{} `type:"structure"`

	// The IDs of individual environments to get information about.
	//
	// EnvironmentIds is a required field
	EnvironmentIds []string `locationName:"environmentIds" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeEnvironmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEnvironmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEnvironmentsInput"}

	if s.EnvironmentIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentIds"))
	}
	if s.EnvironmentIds != nil && len(s.EnvironmentIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DescribeEnvironmentsResult
type DescribeEnvironmentsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the environments that are returned.
	Environments []Environment `json:"cloud9:DescribeEnvironmentsOutput:Environments" locationName:"environments" type:"list"`
}

// String returns the string representation
func (s DescribeEnvironmentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEnvironments = "DescribeEnvironments"

// DescribeEnvironmentsRequest returns a request value for making API operation for
// AWS Cloud9.
//
// Gets information about AWS Cloud9 development environments.
//
//    // Example sending a request using DescribeEnvironmentsRequest.
//    req := client.DescribeEnvironmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DescribeEnvironments
func (c *Client) DescribeEnvironmentsRequest(input *DescribeEnvironmentsInput) DescribeEnvironmentsRequest {
	op := &aws.Operation{
		Name:       opDescribeEnvironments,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEnvironmentsInput{}
	}

	req := c.newRequest(op, input, &DescribeEnvironmentsOutput{})
	return DescribeEnvironmentsRequest{Request: req, Input: input, Copy: c.DescribeEnvironmentsRequest}
}

// DescribeEnvironmentsRequest is the request type for the
// DescribeEnvironments API operation.
type DescribeEnvironmentsRequest struct {
	*aws.Request
	Input *DescribeEnvironmentsInput
	Copy  func(*DescribeEnvironmentsInput) DescribeEnvironmentsRequest
}

// Send marshals and sends the DescribeEnvironments API request.
func (r DescribeEnvironmentsRequest) Send(ctx context.Context) (*DescribeEnvironmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEnvironmentsResponse{
		DescribeEnvironmentsOutput: r.Request.Data.(*DescribeEnvironmentsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEnvironmentsResponse is the response type for the
// DescribeEnvironments API operation.
type DescribeEnvironmentsResponse struct {
	*DescribeEnvironmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEnvironments request.
func (r *DescribeEnvironmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
