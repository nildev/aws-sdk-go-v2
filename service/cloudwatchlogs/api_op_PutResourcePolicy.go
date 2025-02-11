// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/PutResourcePolicyRequest
type PutResourcePolicyInput struct {
	_ struct{} `type:"structure"`

	// Details of the new policy, including the identity of the principal that is
	// enabled to put logs to this account. This is formatted as a JSON string.
	// This parameter is required.
	//
	// The following example creates a resource policy enabling the Route 53 service
	// to put DNS query logs in to the specified log group. Replace "logArn" with
	// the ARN of your CloudWatch Logs resource, such as a log group or log stream.
	//
	// { "Version": "2012-10-17", "Statement": [ { "Sid": "Route53LogsToCloudWatchLogs",
	// "Effect": "Allow", "Principal": { "Service": [ "route53.amazonaws.com" ]
	// }, "Action":"logs:PutLogEvents", "Resource": "logArn" } ] }
	PolicyDocument *string `locationName:"policyDocument" min:"1" type:"string"`

	// Name of the new policy. This parameter is required.
	PolicyName *string `locationName:"policyName" type:"string"`
}

// String returns the string representation
func (s PutResourcePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutResourcePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutResourcePolicyInput"}
	if s.PolicyDocument != nil && len(*s.PolicyDocument) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyDocument", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/PutResourcePolicyResponse
type PutResourcePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The new policy.
	ResourcePolicy *ResourcePolicy `locationName:"resourcePolicy" type:"structure"`
}

// String returns the string representation
func (s PutResourcePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutResourcePolicy = "PutResourcePolicy"

// PutResourcePolicyRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Creates or updates a resource policy allowing other AWS services to put log
// events to this account, such as Amazon Route 53. An account can have up to
// 10 resource policies per region.
//
//    // Example sending a request using PutResourcePolicyRequest.
//    req := client.PutResourcePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/PutResourcePolicy
func (c *Client) PutResourcePolicyRequest(input *PutResourcePolicyInput) PutResourcePolicyRequest {
	op := &aws.Operation{
		Name:       opPutResourcePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutResourcePolicyInput{}
	}

	req := c.newRequest(op, input, &PutResourcePolicyOutput{})
	return PutResourcePolicyRequest{Request: req, Input: input, Copy: c.PutResourcePolicyRequest}
}

// PutResourcePolicyRequest is the request type for the
// PutResourcePolicy API operation.
type PutResourcePolicyRequest struct {
	*aws.Request
	Input *PutResourcePolicyInput
	Copy  func(*PutResourcePolicyInput) PutResourcePolicyRequest
}

// Send marshals and sends the PutResourcePolicy API request.
func (r PutResourcePolicyRequest) Send(ctx context.Context) (*PutResourcePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutResourcePolicyResponse{
		PutResourcePolicyOutput: r.Request.Data.(*PutResourcePolicyOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutResourcePolicyResponse is the response type for the
// PutResourcePolicy API operation.
type PutResourcePolicyResponse struct {
	*PutResourcePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutResourcePolicy request.
func (r *PutResourcePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
