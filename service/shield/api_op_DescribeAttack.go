// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DescribeAttackRequest
type DescribeAttackInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) for the attack that to be described.
	//
	// AttackId is a required field
	AttackId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAttackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAttackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAttackInput"}

	if s.AttackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttackId"))
	}
	if s.AttackId != nil && len(*s.AttackId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AttackId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DescribeAttackResponse
type DescribeAttackOutput struct {
	_ struct{} `type:"structure"`

	// The attack that is described.
	Attack *AttackDetail `json:"shield:DescribeAttackOutput:Attack" type:"structure"`
}

// String returns the string representation
func (s DescribeAttackOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAttack = "DescribeAttack"

// DescribeAttackRequest returns a request value for making API operation for
// AWS Shield.
//
// Describes the details of a DDoS attack.
//
//    // Example sending a request using DescribeAttackRequest.
//    req := client.DescribeAttackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DescribeAttack
func (c *Client) DescribeAttackRequest(input *DescribeAttackInput) DescribeAttackRequest {
	op := &aws.Operation{
		Name:       opDescribeAttack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAttackInput{}
	}

	req := c.newRequest(op, input, &DescribeAttackOutput{})
	return DescribeAttackRequest{Request: req, Input: input, Copy: c.DescribeAttackRequest}
}

// DescribeAttackRequest is the request type for the
// DescribeAttack API operation.
type DescribeAttackRequest struct {
	*aws.Request
	Input *DescribeAttackInput
	Copy  func(*DescribeAttackInput) DescribeAttackRequest
}

// Send marshals and sends the DescribeAttack API request.
func (r DescribeAttackRequest) Send(ctx context.Context) (*DescribeAttackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAttackResponse{
		DescribeAttackOutput: r.Request.Data.(*DescribeAttackOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAttackResponse is the response type for the
// DescribeAttack API operation.
type DescribeAttackResponse struct {
	*DescribeAttackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAttack request.
func (r *DescribeAttackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
