// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateStreamProcessorInput struct {
	_ struct{} `type:"structure"`

	// Kinesis video stream stream that provides the source streaming video. If
	// you are using the AWS CLI, the parameter name is StreamProcessorInput.
	//
	// Input is a required field
	Input *StreamProcessorInput `type:"structure" required:"true"`

	// An identifier you assign to the stream processor. You can use Name to manage
	// the stream processor. For example, you can get the current status of the
	// stream processor by calling DescribeStreamProcessor. Name is idempotent.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Kinesis data stream stream to which Amazon Rekognition Video puts the analysis
	// results. If you are using the AWS CLI, the parameter name is StreamProcessorOutput.
	//
	// Output is a required field
	Output *StreamProcessorOutput `type:"structure" required:"true"`

	// ARN of the IAM role that allows access to the stream processor.
	//
	// RoleArn is a required field
	RoleArn *string `type:"string" required:"true"`

	// Face recognition input parameters to be used by the stream processor. Includes
	// the collection to use for face recognition and the face attributes to detect.
	//
	// Settings is a required field
	Settings *StreamProcessorSettings `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateStreamProcessorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateStreamProcessorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateStreamProcessorInput"}

	if s.Input == nil {
		invalidParams.Add(aws.NewErrParamRequired("Input"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Output == nil {
		invalidParams.Add(aws.NewErrParamRequired("Output"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if s.Settings == nil {
		invalidParams.Add(aws.NewErrParamRequired("Settings"))
	}
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			invalidParams.AddNested("Settings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateStreamProcessorOutput struct {
	_ struct{} `type:"structure"`

	// ARN for the newly create stream processor.
	StreamProcessorArn *string `json:"rekognition:CreateStreamProcessorOutput:StreamProcessorArn" type:"string"`
}

// String returns the string representation
func (s CreateStreamProcessorOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateStreamProcessor = "CreateStreamProcessor"

// CreateStreamProcessorRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Creates an Amazon Rekognition stream processor that you can use to detect
// and recognize faces in a streaming video.
//
// Amazon Rekognition Video is a consumer of live video from Amazon Kinesis
// Video Streams. Amazon Rekognition Video sends analysis results to Amazon
// Kinesis Data Streams.
//
// You provide as input a Kinesis video stream (Input) and a Kinesis data stream
// (Output) stream. You also specify the face recognition criteria in Settings.
// For example, the collection containing faces that you want to recognize.
// Use Name to assign an identifier for the stream processor. You use Name to
// manage the stream processor. For example, you can start processing the source
// video by calling StartStreamProcessor with the Name field.
//
// After you have finished analyzing a streaming video, use StopStreamProcessor
// to stop processing. You can delete the stream processor by calling DeleteStreamProcessor.
//
//    // Example sending a request using CreateStreamProcessorRequest.
//    req := client.CreateStreamProcessorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateStreamProcessorRequest(input *CreateStreamProcessorInput) CreateStreamProcessorRequest {
	op := &aws.Operation{
		Name:       opCreateStreamProcessor,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateStreamProcessorInput{}
	}

	req := c.newRequest(op, input, &CreateStreamProcessorOutput{})
	return CreateStreamProcessorRequest{Request: req, Input: input, Copy: c.CreateStreamProcessorRequest}
}

// CreateStreamProcessorRequest is the request type for the
// CreateStreamProcessor API operation.
type CreateStreamProcessorRequest struct {
	*aws.Request
	Input *CreateStreamProcessorInput
	Copy  func(*CreateStreamProcessorInput) CreateStreamProcessorRequest
}

// Send marshals and sends the CreateStreamProcessor API request.
func (r CreateStreamProcessorRequest) Send(ctx context.Context) (*CreateStreamProcessorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStreamProcessorResponse{
		CreateStreamProcessorOutput: r.Request.Data.(*CreateStreamProcessorOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStreamProcessorResponse is the response type for the
// CreateStreamProcessor API operation.
type CreateStreamProcessorResponse struct {
	*CreateStreamProcessorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStreamProcessor request.
func (r *CreateStreamProcessorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
