// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartCelebrityRecognitionInput struct {
	_ struct{} `type:"structure"`

	// Idempotent token used to identify the start request. If you use the same
	// token with multiple StartCelebrityRecognition requests, the same JobId is
	// returned. Use ClientRequestToken to prevent the same job from being accidently
	// started more than once.
	ClientRequestToken *string `min:"1" type:"string"`

	// Unique identifier you specify to identify the job in the completion status
	// published to the Amazon Simple Notification Service topic.
	JobTag *string `min:"1" type:"string"`

	// The Amazon SNS topic ARN that you want Amazon Rekognition Video to publish
	// the completion status of the celebrity recognition analysis to.
	NotificationChannel *NotificationChannel `type:"structure"`

	// The video in which you want to recognize celebrities. The video must be stored
	// in an Amazon S3 bucket.
	//
	// Video is a required field
	Video *Video `type:"structure" required:"true"`
}

// String returns the string representation
func (s StartCelebrityRecognitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartCelebrityRecognitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartCelebrityRecognitionInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.JobTag != nil && len(*s.JobTag) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobTag", 1))
	}

	if s.Video == nil {
		invalidParams.Add(aws.NewErrParamRequired("Video"))
	}
	if s.NotificationChannel != nil {
		if err := s.NotificationChannel.Validate(); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(aws.ErrInvalidParams))
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			invalidParams.AddNested("Video", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartCelebrityRecognitionOutput struct {
	_ struct{} `type:"structure"`

	// The identifier for the celebrity recognition analysis job. Use JobId to identify
	// the job in a subsequent call to GetCelebrityRecognition.
	JobId *string `json:"rekognition:StartCelebrityRecognitionOutput:JobId" min:"1" type:"string"`
}

// String returns the string representation
func (s StartCelebrityRecognitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartCelebrityRecognition = "StartCelebrityRecognition"

// StartCelebrityRecognitionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Starts asynchronous recognition of celebrities in a stored video.
//
// Amazon Rekognition Video can detect celebrities in a video must be stored
// in an Amazon S3 bucket. Use Video to specify the bucket name and the filename
// of the video. StartCelebrityRecognition returns a job identifier (JobId)
// which you use to get the results of the analysis. When celebrity recognition
// analysis is finished, Amazon Rekognition Video publishes a completion status
// to the Amazon Simple Notification Service topic that you specify in NotificationChannel.
// To get the results of the celebrity recognition analysis, first check that
// the status value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetCelebrityRecognition and pass the job identifier (JobId) from the initial
// call to StartCelebrityRecognition.
//
// For more information, see Recognizing Celebrities in the Amazon Rekognition
// Developer Guide.
//
//    // Example sending a request using StartCelebrityRecognitionRequest.
//    req := client.StartCelebrityRecognitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StartCelebrityRecognitionRequest(input *StartCelebrityRecognitionInput) StartCelebrityRecognitionRequest {
	op := &aws.Operation{
		Name:       opStartCelebrityRecognition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartCelebrityRecognitionInput{}
	}

	req := c.newRequest(op, input, &StartCelebrityRecognitionOutput{})
	return StartCelebrityRecognitionRequest{Request: req, Input: input, Copy: c.StartCelebrityRecognitionRequest}
}

// StartCelebrityRecognitionRequest is the request type for the
// StartCelebrityRecognition API operation.
type StartCelebrityRecognitionRequest struct {
	*aws.Request
	Input *StartCelebrityRecognitionInput
	Copy  func(*StartCelebrityRecognitionInput) StartCelebrityRecognitionRequest
}

// Send marshals and sends the StartCelebrityRecognition API request.
func (r StartCelebrityRecognitionRequest) Send(ctx context.Context) (*StartCelebrityRecognitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartCelebrityRecognitionResponse{
		StartCelebrityRecognitionOutput: r.Request.Data.(*StartCelebrityRecognitionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartCelebrityRecognitionResponse is the response type for the
// StartCelebrityRecognition API operation.
type StartCelebrityRecognitionResponse struct {
	*StartCelebrityRecognitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartCelebrityRecognition request.
func (r *StartCelebrityRecognitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
