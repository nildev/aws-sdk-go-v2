// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to list completed and failed managed actions.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeEnvironmentManagedActionHistoryRequest
type DescribeEnvironmentManagedActionHistoryInput struct {
	_ struct{} `type:"structure"`

	// The environment ID of the target environment.
	EnvironmentId *string `type:"string"`

	// The name of the target environment.
	EnvironmentName *string `min:"4" type:"string"`

	// The maximum number of items to return for a single request.
	MaxItems *int64 `type:"integer"`

	// The pagination token returned by a previous request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeEnvironmentManagedActionHistoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEnvironmentManagedActionHistoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEnvironmentManagedActionHistoryInput"}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A result message containing a list of completed and failed managed actions.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeEnvironmentManagedActionHistoryResult
type DescribeEnvironmentManagedActionHistoryOutput struct {
	_ struct{} `type:"structure"`

	// A list of completed and failed managed actions.
	ManagedActionHistoryItems []ManagedActionHistoryItem `json:"elasticbeanstalk:DescribeEnvironmentManagedActionHistoryOutput:ManagedActionHistoryItems" min:"1" type:"list"`

	// A pagination token that you pass to DescribeEnvironmentManagedActionHistory
	// to get the next page of results.
	NextToken *string `json:"elasticbeanstalk:DescribeEnvironmentManagedActionHistoryOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s DescribeEnvironmentManagedActionHistoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEnvironmentManagedActionHistory = "DescribeEnvironmentManagedActionHistory"

// DescribeEnvironmentManagedActionHistoryRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Lists an environment's completed and failed managed actions.
//
//    // Example sending a request using DescribeEnvironmentManagedActionHistoryRequest.
//    req := client.DescribeEnvironmentManagedActionHistoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeEnvironmentManagedActionHistory
func (c *Client) DescribeEnvironmentManagedActionHistoryRequest(input *DescribeEnvironmentManagedActionHistoryInput) DescribeEnvironmentManagedActionHistoryRequest {
	op := &aws.Operation{
		Name:       opDescribeEnvironmentManagedActionHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEnvironmentManagedActionHistoryInput{}
	}

	req := c.newRequest(op, input, &DescribeEnvironmentManagedActionHistoryOutput{})
	return DescribeEnvironmentManagedActionHistoryRequest{Request: req, Input: input, Copy: c.DescribeEnvironmentManagedActionHistoryRequest}
}

// DescribeEnvironmentManagedActionHistoryRequest is the request type for the
// DescribeEnvironmentManagedActionHistory API operation.
type DescribeEnvironmentManagedActionHistoryRequest struct {
	*aws.Request
	Input *DescribeEnvironmentManagedActionHistoryInput
	Copy  func(*DescribeEnvironmentManagedActionHistoryInput) DescribeEnvironmentManagedActionHistoryRequest
}

// Send marshals and sends the DescribeEnvironmentManagedActionHistory API request.
func (r DescribeEnvironmentManagedActionHistoryRequest) Send(ctx context.Context) (*DescribeEnvironmentManagedActionHistoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEnvironmentManagedActionHistoryResponse{
		DescribeEnvironmentManagedActionHistoryOutput: r.Request.Data.(*DescribeEnvironmentManagedActionHistoryOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEnvironmentManagedActionHistoryResponse is the response type for the
// DescribeEnvironmentManagedActionHistory API operation.
type DescribeEnvironmentManagedActionHistoryResponse struct {
	*DescribeEnvironmentManagedActionHistoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEnvironmentManagedActionHistory request.
func (r *DescribeEnvironmentManagedActionHistoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
