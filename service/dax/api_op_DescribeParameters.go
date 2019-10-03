// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeParametersRequest
type DescribeParametersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	//
	// The value for MaxResults must be between 20 and 100.
	MaxResults *int64 `type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `type:"string"`

	// The name of the parameter group.
	//
	// ParameterGroupName is a required field
	ParameterGroupName *string `type:"string" required:"true"`

	// How the parameter is defined. For example, system denotes a system-defined
	// parameter.
	Source *string `type:"string"`
}

// String returns the string representation
func (s DescribeParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeParametersInput"}

	if s.ParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeParametersResponse
type DescribeParametersOutput struct {
	_ struct{} `type:"structure"`

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string `json:"dax:DescribeParametersOutput:NextToken" type:"string"`

	// A list of parameters within a parameter group. Each element in the list represents
	// one parameter.
	Parameters []Parameter `json:"dax:DescribeParametersOutput:Parameters" type:"list"`
}

// String returns the string representation
func (s DescribeParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeParameters = "DescribeParameters"

// DescribeParametersRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Returns the detailed parameter list for a particular parameter group.
//
//    // Example sending a request using DescribeParametersRequest.
//    req := client.DescribeParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DescribeParameters
func (c *Client) DescribeParametersRequest(input *DescribeParametersInput) DescribeParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeParametersOutput{})
	return DescribeParametersRequest{Request: req, Input: input, Copy: c.DescribeParametersRequest}
}

// DescribeParametersRequest is the request type for the
// DescribeParameters API operation.
type DescribeParametersRequest struct {
	*aws.Request
	Input *DescribeParametersInput
	Copy  func(*DescribeParametersInput) DescribeParametersRequest
}

// Send marshals and sends the DescribeParameters API request.
func (r DescribeParametersRequest) Send(ctx context.Context) (*DescribeParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeParametersResponse{
		DescribeParametersOutput: r.Request.Data.(*DescribeParametersOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeParametersResponse is the response type for the
// DescribeParameters API operation.
type DescribeParametersResponse struct {
	*DescribeParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeParameters request.
func (r *DescribeParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
