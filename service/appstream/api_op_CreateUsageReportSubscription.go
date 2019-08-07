// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateUsageReportSubscriptionRequest
type CreateUsageReportSubscriptionInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateUsageReportSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateUsageReportSubscriptionResult
type CreateUsageReportSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 bucket where generated reports are stored.
	//
	// If you enabled on-instance session scripts and Amazon S3 logging for your
	// session script configuration, AppStream 2.0 created an S3 bucket to store
	// the script output. The bucket is unique to your account and Region. When
	// you enable usage reporting in this case, AppStream 2.0 uses the same bucket
	// to store your usage reports. If you haven't already enabled on-instance session
	// scripts, when you enable usage reports, AppStream 2.0 creates a new S3 bucket.
	S3BucketName *string `json:"appstream2:CreateUsageReportSubscriptionOutput:S3BucketName" min:"1" type:"string"`

	// The schedule for generating usage reports.
	Schedule UsageReportSchedule `json:"appstream2:CreateUsageReportSubscriptionOutput:Schedule" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateUsageReportSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateUsageReportSubscription = "CreateUsageReportSubscription"

// CreateUsageReportSubscriptionRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Creates a usage report subscription. Usage reports are generated daily.
//
//    // Example sending a request using CreateUsageReportSubscriptionRequest.
//    req := client.CreateUsageReportSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateUsageReportSubscription
func (c *Client) CreateUsageReportSubscriptionRequest(input *CreateUsageReportSubscriptionInput) CreateUsageReportSubscriptionRequest {
	op := &aws.Operation{
		Name:       opCreateUsageReportSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateUsageReportSubscriptionInput{}
	}

	req := c.newRequest(op, input, &CreateUsageReportSubscriptionOutput{})
	return CreateUsageReportSubscriptionRequest{Request: req, Input: input, Copy: c.CreateUsageReportSubscriptionRequest}
}

// CreateUsageReportSubscriptionRequest is the request type for the
// CreateUsageReportSubscription API operation.
type CreateUsageReportSubscriptionRequest struct {
	*aws.Request
	Input *CreateUsageReportSubscriptionInput
	Copy  func(*CreateUsageReportSubscriptionInput) CreateUsageReportSubscriptionRequest
}

// Send marshals and sends the CreateUsageReportSubscription API request.
func (r CreateUsageReportSubscriptionRequest) Send(ctx context.Context) (*CreateUsageReportSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUsageReportSubscriptionResponse{
		CreateUsageReportSubscriptionOutput: r.Request.Data.(*CreateUsageReportSubscriptionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUsageReportSubscriptionResponse is the response type for the
// CreateUsageReportSubscription API operation.
type CreateUsageReportSubscriptionResponse struct {
	*CreateUsageReportSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUsageReportSubscription request.
func (r *CreateUsageReportSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
