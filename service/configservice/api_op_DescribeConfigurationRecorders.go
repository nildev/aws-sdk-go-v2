// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the DescribeConfigurationRecorders action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeConfigurationRecordersRequest
type DescribeConfigurationRecordersInput struct {
	_ struct{} `type:"structure"`

	// A list of configuration recorder names.
	ConfigurationRecorderNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeConfigurationRecordersInput) String() string {
	return awsutil.Prettify(s)
}

// The output for the DescribeConfigurationRecorders action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeConfigurationRecordersResponse
type DescribeConfigurationRecordersOutput struct {
	_ struct{} `type:"structure"`

	// A list that contains the descriptions of the specified configuration recorders.
	ConfigurationRecorders []ConfigurationRecorder `json:"config:DescribeConfigurationRecordersOutput:ConfigurationRecorders" type:"list"`
}

// String returns the string representation
func (s DescribeConfigurationRecordersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeConfigurationRecorders = "DescribeConfigurationRecorders"

// DescribeConfigurationRecordersRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the details for the specified configuration recorders. If the configuration
// recorder is not specified, this action returns the details for all configuration
// recorders associated with the account.
//
// Currently, you can specify only one configuration recorder per region in
// your account.
//
//    // Example sending a request using DescribeConfigurationRecordersRequest.
//    req := client.DescribeConfigurationRecordersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeConfigurationRecorders
func (c *Client) DescribeConfigurationRecordersRequest(input *DescribeConfigurationRecordersInput) DescribeConfigurationRecordersRequest {
	op := &aws.Operation{
		Name:       opDescribeConfigurationRecorders,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConfigurationRecordersInput{}
	}

	req := c.newRequest(op, input, &DescribeConfigurationRecordersOutput{})
	return DescribeConfigurationRecordersRequest{Request: req, Input: input, Copy: c.DescribeConfigurationRecordersRequest}
}

// DescribeConfigurationRecordersRequest is the request type for the
// DescribeConfigurationRecorders API operation.
type DescribeConfigurationRecordersRequest struct {
	*aws.Request
	Input *DescribeConfigurationRecordersInput
	Copy  func(*DescribeConfigurationRecordersInput) DescribeConfigurationRecordersRequest
}

// Send marshals and sends the DescribeConfigurationRecorders API request.
func (r DescribeConfigurationRecordersRequest) Send(ctx context.Context) (*DescribeConfigurationRecordersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConfigurationRecordersResponse{
		DescribeConfigurationRecordersOutput: r.Request.Data.(*DescribeConfigurationRecordersOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConfigurationRecordersResponse is the response type for the
// DescribeConfigurationRecorders API operation.
type DescribeConfigurationRecordersResponse struct {
	*DescribeConfigurationRecordersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConfigurationRecorders request.
func (r *DescribeConfigurationRecordersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
