// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/BatchEnableStandardsRequest
type BatchEnableStandardsInput struct {
	_ struct{} `type:"structure"`

	// The list of standards compliance checks to enable.
	//
	// In this release, Security Hub supports only the CIS AWS Foundations standard.
	//
	// The ARN for the standard is arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0.
	//
	// StandardsSubscriptionRequests is a required field
	StandardsSubscriptionRequests []StandardsSubscriptionRequest `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchEnableStandardsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchEnableStandardsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchEnableStandardsInput"}

	if s.StandardsSubscriptionRequests == nil {
		invalidParams.Add(aws.NewErrParamRequired("StandardsSubscriptionRequests"))
	}
	if s.StandardsSubscriptionRequests != nil && len(s.StandardsSubscriptionRequests) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StandardsSubscriptionRequests", 1))
	}
	if s.StandardsSubscriptionRequests != nil {
		for i, v := range s.StandardsSubscriptionRequests {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "StandardsSubscriptionRequests", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchEnableStandardsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.StandardsSubscriptionRequests != nil {
		v := s.StandardsSubscriptionRequests

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "StandardsSubscriptionRequests", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/BatchEnableStandardsResponse
type BatchEnableStandardsOutput struct {
	_ struct{} `type:"structure"`

	// The details of the standards subscriptions that were enabled.
	StandardsSubscriptions []StandardsSubscription `json:"securityhub:BatchEnableStandardsOutput:StandardsSubscriptions" type:"list"`
}

// String returns the string representation
func (s BatchEnableStandardsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchEnableStandardsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.StandardsSubscriptions != nil {
		v := s.StandardsSubscriptions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "StandardsSubscriptions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchEnableStandards = "BatchEnableStandards"

// BatchEnableStandardsRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Enables the standards specified by the provided standardsArn. In this release,
// only CIS AWS Foundations standards are supported. For more information, see
// Standards Supported in AWS Security Hub (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards.html).
//
//    // Example sending a request using BatchEnableStandardsRequest.
//    req := client.BatchEnableStandardsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/BatchEnableStandards
func (c *Client) BatchEnableStandardsRequest(input *BatchEnableStandardsInput) BatchEnableStandardsRequest {
	op := &aws.Operation{
		Name:       opBatchEnableStandards,
		HTTPMethod: "POST",
		HTTPPath:   "/standards/register",
	}

	if input == nil {
		input = &BatchEnableStandardsInput{}
	}

	req := c.newRequest(op, input, &BatchEnableStandardsOutput{})
	return BatchEnableStandardsRequest{Request: req, Input: input, Copy: c.BatchEnableStandardsRequest}
}

// BatchEnableStandardsRequest is the request type for the
// BatchEnableStandards API operation.
type BatchEnableStandardsRequest struct {
	*aws.Request
	Input *BatchEnableStandardsInput
	Copy  func(*BatchEnableStandardsInput) BatchEnableStandardsRequest
}

// Send marshals and sends the BatchEnableStandards API request.
func (r BatchEnableStandardsRequest) Send(ctx context.Context) (*BatchEnableStandardsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchEnableStandardsResponse{
		BatchEnableStandardsOutput: r.Request.Data.(*BatchEnableStandardsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchEnableStandardsResponse is the response type for the
// BatchEnableStandards API operation.
type BatchEnableStandardsResponse struct {
	*BatchEnableStandardsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchEnableStandards request.
func (r *BatchEnableStandardsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
