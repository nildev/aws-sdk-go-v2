// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/GetQueueUrlRequest
type GetQueueUrlInput struct {
	_ struct{} `type:"structure"`

	// The name of the queue whose URL must be fetched. Maximum 80 characters. Valid
	// values: alphanumeric characters, hyphens (-), and underscores (_).
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueName is a required field
	QueueName *string `type:"string" required:"true"`

	// The AWS account ID of the account that created the queue.
	QueueOwnerAWSAccountId *string `type:"string"`
}

// String returns the string representation
func (s GetQueueUrlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetQueueUrlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetQueueUrlInput"}

	if s.QueueName == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// For more information, see Interpreting Responses (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-api-responses.html)
// in the Amazon Simple Queue Service Developer Guide.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/GetQueueUrlResult
type GetQueueUrlOutput struct {
	_ struct{} `type:"structure"`

	// The URL of the queue.
	QueueUrl *string `json:"sqs:GetQueueUrlOutput:QueueUrl" type:"string"`
}

// String returns the string representation
func (s GetQueueUrlOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetQueueUrl = "GetQueueUrl"

// GetQueueUrlRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Returns the URL of an existing Amazon SQS queue.
//
// To access a queue that belongs to another AWS account, use the QueueOwnerAWSAccountId
// parameter to specify the account ID of the queue's owner. The queue's owner
// must grant you permission to access the queue. For more information about
// shared queue access, see AddPermission or see Allow Developers to Write Messages
// to a Shared Queue (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue)
// in the Amazon Simple Queue Service Developer Guide.
//
//    // Example sending a request using GetQueueUrlRequest.
//    req := client.GetQueueUrlRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/GetQueueUrl
func (c *Client) GetQueueUrlRequest(input *GetQueueUrlInput) GetQueueUrlRequest {
	op := &aws.Operation{
		Name:       opGetQueueUrl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetQueueUrlInput{}
	}

	req := c.newRequest(op, input, &GetQueueUrlOutput{})
	return GetQueueUrlRequest{Request: req, Input: input, Copy: c.GetQueueUrlRequest}
}

// GetQueueUrlRequest is the request type for the
// GetQueueUrl API operation.
type GetQueueUrlRequest struct {
	*aws.Request
	Input *GetQueueUrlInput
	Copy  func(*GetQueueUrlInput) GetQueueUrlRequest
}

// Send marshals and sends the GetQueueUrl API request.
func (r GetQueueUrlRequest) Send(ctx context.Context) (*GetQueueUrlResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetQueueUrlResponse{
		GetQueueUrlOutput: r.Request.Data.(*GetQueueUrlOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetQueueUrlResponse is the response type for the
// GetQueueUrl API operation.
type GetQueueUrlResponse struct {
	*GetQueueUrlOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetQueueUrl request.
func (r *GetQueueUrlResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
