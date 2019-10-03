// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Initiates the verification of an existing trust relationship between an AWS
// Managed Microsoft AD directory and an external domain.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/VerifyTrustRequest
type VerifyTrustInput struct {
	_ struct{} `type:"structure"`

	// The unique Trust ID of the trust relationship to verify.
	//
	// TrustId is a required field
	TrustId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s VerifyTrustInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VerifyTrustInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "VerifyTrustInput"}

	if s.TrustId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrustId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result of a VerifyTrust request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/VerifyTrustResult
type VerifyTrustOutput struct {
	_ struct{} `type:"structure"`

	// The unique Trust ID of the trust relationship that was verified.
	TrustId *string `json:"ds:VerifyTrustOutput:TrustId" type:"string"`
}

// String returns the string representation
func (s VerifyTrustOutput) String() string {
	return awsutil.Prettify(s)
}

const opVerifyTrust = "VerifyTrust"

// VerifyTrustRequest returns a request value for making API operation for
// AWS Directory Service.
//
// AWS Directory Service for Microsoft Active Directory allows you to configure
// and verify trust relationships.
//
// This action verifies a trust relationship between your AWS Managed Microsoft
// AD directory and an external domain.
//
//    // Example sending a request using VerifyTrustRequest.
//    req := client.VerifyTrustRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/VerifyTrust
func (c *Client) VerifyTrustRequest(input *VerifyTrustInput) VerifyTrustRequest {
	op := &aws.Operation{
		Name:       opVerifyTrust,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &VerifyTrustInput{}
	}

	req := c.newRequest(op, input, &VerifyTrustOutput{})
	return VerifyTrustRequest{Request: req, Input: input, Copy: c.VerifyTrustRequest}
}

// VerifyTrustRequest is the request type for the
// VerifyTrust API operation.
type VerifyTrustRequest struct {
	*aws.Request
	Input *VerifyTrustInput
	Copy  func(*VerifyTrustInput) VerifyTrustRequest
}

// Send marshals and sends the VerifyTrust API request.
func (r VerifyTrustRequest) Send(ctx context.Context) (*VerifyTrustResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &VerifyTrustResponse{
		VerifyTrustOutput: r.Request.Data.(*VerifyTrustOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// VerifyTrustResponse is the response type for the
// VerifyTrust API operation.
type VerifyTrustResponse struct {
	*VerifyTrustOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// VerifyTrust request.
func (r *VerifyTrustResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
