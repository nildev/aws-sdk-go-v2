// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeTrustedAdvisorCheckResultRequest
type DescribeTrustedAdvisorCheckResultInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the Trusted Advisor check.
	//
	// CheckId is a required field
	CheckId *string `locationName:"checkId" type:"string" required:"true"`

	// The ISO 639-1 code for the language in which AWS provides support. AWS Support
	// currently supports English ("en") and Japanese ("ja"). Language parameters
	// must be passed explicitly for operations that take them.
	Language *string `locationName:"language" type:"string"`
}

// String returns the string representation
func (s DescribeTrustedAdvisorCheckResultInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTrustedAdvisorCheckResultInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTrustedAdvisorCheckResultInput"}

	if s.CheckId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CheckId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of the Trusted Advisor check returned by the DescribeTrustedAdvisorCheckResult
// operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeTrustedAdvisorCheckResultResponse
type DescribeTrustedAdvisorCheckResultOutput struct {
	_ struct{} `type:"structure"`

	// The detailed results of the Trusted Advisor check.
	Result *TrustedAdvisorCheckResult `json:"support:DescribeTrustedAdvisorCheckResultOutput:Result" locationName:"result" type:"structure"`
}

// String returns the string representation
func (s DescribeTrustedAdvisorCheckResultOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTrustedAdvisorCheckResult = "DescribeTrustedAdvisorCheckResult"

// DescribeTrustedAdvisorCheckResultRequest returns a request value for making API operation for
// AWS Support.
//
// Returns the results of the Trusted Advisor check that has the specified check
// ID. Check IDs can be obtained by calling DescribeTrustedAdvisorChecks.
//
// The response contains a TrustedAdvisorCheckResult object, which contains
// these three objects:
//
//    * TrustedAdvisorCategorySpecificSummary
//
//    * TrustedAdvisorResourceDetail
//
//    * TrustedAdvisorResourcesSummary
//
// In addition, the response contains these fields:
//
//    * status. The alert status of the check: "ok" (green), "warning" (yellow),
//    "error" (red), or "not_available".
//
//    * timestamp. The time of the last refresh of the check.
//
//    * checkId. The unique identifier for the check.
//
//    // Example sending a request using DescribeTrustedAdvisorCheckResultRequest.
//    req := client.DescribeTrustedAdvisorCheckResultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeTrustedAdvisorCheckResult
func (c *Client) DescribeTrustedAdvisorCheckResultRequest(input *DescribeTrustedAdvisorCheckResultInput) DescribeTrustedAdvisorCheckResultRequest {
	op := &aws.Operation{
		Name:       opDescribeTrustedAdvisorCheckResult,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTrustedAdvisorCheckResultInput{}
	}

	req := c.newRequest(op, input, &DescribeTrustedAdvisorCheckResultOutput{})
	return DescribeTrustedAdvisorCheckResultRequest{Request: req, Input: input, Copy: c.DescribeTrustedAdvisorCheckResultRequest}
}

// DescribeTrustedAdvisorCheckResultRequest is the request type for the
// DescribeTrustedAdvisorCheckResult API operation.
type DescribeTrustedAdvisorCheckResultRequest struct {
	*aws.Request
	Input *DescribeTrustedAdvisorCheckResultInput
	Copy  func(*DescribeTrustedAdvisorCheckResultInput) DescribeTrustedAdvisorCheckResultRequest
}

// Send marshals and sends the DescribeTrustedAdvisorCheckResult API request.
func (r DescribeTrustedAdvisorCheckResultRequest) Send(ctx context.Context) (*DescribeTrustedAdvisorCheckResultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTrustedAdvisorCheckResultResponse{
		DescribeTrustedAdvisorCheckResultOutput: r.Request.Data.(*DescribeTrustedAdvisorCheckResultOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTrustedAdvisorCheckResultResponse is the response type for the
// DescribeTrustedAdvisorCheckResult API operation.
type DescribeTrustedAdvisorCheckResultResponse struct {
	*DescribeTrustedAdvisorCheckResultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTrustedAdvisorCheckResult request.
func (r *DescribeTrustedAdvisorCheckResultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
