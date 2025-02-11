// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/StartSpeechSynthesisTaskInput
type StartSpeechSynthesisTaskInput struct {
	_ struct{} `type:"structure"`

	// Optional language code for the Speech Synthesis request. This is only necessary
	// if using a bilingual voice, such as Aditi, which can be used for either Indian
	// English (en-IN) or Hindi (hi-IN).
	//
	// If a bilingual voice is used and no language code is specified, Amazon Polly
	// will use the default language of the bilingual voice. The default language
	// for any voice is the one returned by the DescribeVoices (https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html)
	// operation for the LanguageCode parameter. For example, if no language code
	// is specified, Aditi will use Indian English rather than Hindi.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// List of one or more pronunciation lexicon names you want the service to apply
	// during synthesis. Lexicons are applied only if the language of the lexicon
	// is the same as the language of the voice.
	LexiconNames []string `type:"list"`

	// The format in which the returned output will be encoded. For audio stream,
	// this will be mp3, ogg_vorbis, or pcm. For speech marks, this will be json.
	//
	// OutputFormat is a required field
	OutputFormat OutputFormat `type:"string" required:"true" enum:"true"`

	// Amazon S3 bucket name to which the output file will be saved.
	//
	// OutputS3BucketName is a required field
	OutputS3BucketName *string `type:"string" required:"true"`

	// The Amazon S3 key prefix for the output speech file.
	OutputS3KeyPrefix *string `type:"string"`

	// The audio frequency specified in Hz.
	//
	// The valid values for mp3 and ogg_vorbis are "8000", "16000", and "22050".
	// The default value is "22050".
	//
	// Valid values for pcm are "8000" and "16000" The default value is "16000".
	SampleRate *string `type:"string"`

	// ARN for the SNS topic optionally used for providing status notification for
	// a speech synthesis task.
	SnsTopicArn *string `type:"string"`

	// The type of speech marks returned for the input text.
	SpeechMarkTypes []SpeechMarkType `type:"list"`

	// The input text to synthesize. If you specify ssml as the TextType, follow
	// the SSML format for the input text.
	//
	// Text is a required field
	Text *string `type:"string" required:"true"`

	// Specifies whether the input text is plain text or SSML. The default value
	// is plain text.
	TextType TextType `type:"string" enum:"true"`

	// Voice ID to use for the synthesis.
	//
	// VoiceId is a required field
	VoiceId VoiceId `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s StartSpeechSynthesisTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSpeechSynthesisTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSpeechSynthesisTaskInput"}
	if len(s.OutputFormat) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("OutputFormat"))
	}

	if s.OutputS3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputS3BucketName"))
	}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}
	if len(s.VoiceId) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("VoiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartSpeechSynthesisTaskInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.LanguageCode) > 0 {
		v := s.LanguageCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LanguageCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.LexiconNames != nil {
		v := s.LexiconNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "LexiconNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.OutputFormat) > 0 {
		v := s.OutputFormat

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OutputFormat", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.OutputS3BucketName != nil {
		v := *s.OutputS3BucketName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OutputS3BucketName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OutputS3KeyPrefix != nil {
		v := *s.OutputS3KeyPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OutputS3KeyPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SampleRate != nil {
		v := *s.SampleRate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SampleRate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SnsTopicArn != nil {
		v := *s.SnsTopicArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SnsTopicArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SpeechMarkTypes != nil {
		v := s.SpeechMarkTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SpeechMarkTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Text != nil {
		v := *s.Text

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Text", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.TextType) > 0 {
		v := s.TextType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TextType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.VoiceId) > 0 {
		v := s.VoiceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VoiceId", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/StartSpeechSynthesisTaskOutput
type StartSpeechSynthesisTaskOutput struct {
	_ struct{} `type:"structure"`

	// SynthesisTask object that provides information and attributes about a newly
	// submitted speech synthesis task.
	SynthesisTask *SynthesisTask `type:"structure"`
}

// String returns the string representation
func (s StartSpeechSynthesisTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartSpeechSynthesisTaskOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SynthesisTask != nil {
		v := s.SynthesisTask

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SynthesisTask", v, metadata)
	}
	return nil
}

const opStartSpeechSynthesisTask = "StartSpeechSynthesisTask"

// StartSpeechSynthesisTaskRequest returns a request value for making API operation for
// Amazon Polly.
//
// Allows the creation of an asynchronous synthesis task, by starting a new
// SpeechSynthesisTask. This operation requires all the standard information
// needed for speech synthesis, plus the name of an Amazon S3 bucket for the
// service to store the output of the synthesis task and two optional parameters
// (OutputS3KeyPrefix and SnsTopicArn). Once the synthesis task is created,
// this operation will return a SpeechSynthesisTask object, which will include
// an identifier of this task as well as the current status.
//
//    // Example sending a request using StartSpeechSynthesisTaskRequest.
//    req := client.StartSpeechSynthesisTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/StartSpeechSynthesisTask
func (c *Client) StartSpeechSynthesisTaskRequest(input *StartSpeechSynthesisTaskInput) StartSpeechSynthesisTaskRequest {
	op := &aws.Operation{
		Name:       opStartSpeechSynthesisTask,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/synthesisTasks",
	}

	if input == nil {
		input = &StartSpeechSynthesisTaskInput{}
	}

	req := c.newRequest(op, input, &StartSpeechSynthesisTaskOutput{})
	return StartSpeechSynthesisTaskRequest{Request: req, Input: input, Copy: c.StartSpeechSynthesisTaskRequest}
}

// StartSpeechSynthesisTaskRequest is the request type for the
// StartSpeechSynthesisTask API operation.
type StartSpeechSynthesisTaskRequest struct {
	*aws.Request
	Input *StartSpeechSynthesisTaskInput
	Copy  func(*StartSpeechSynthesisTaskInput) StartSpeechSynthesisTaskRequest
}

// Send marshals and sends the StartSpeechSynthesisTask API request.
func (r StartSpeechSynthesisTaskRequest) Send(ctx context.Context) (*StartSpeechSynthesisTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartSpeechSynthesisTaskResponse{
		StartSpeechSynthesisTaskOutput: r.Request.Data.(*StartSpeechSynthesisTaskOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartSpeechSynthesisTaskResponse is the response type for the
// StartSpeechSynthesisTask API operation.
type StartSpeechSynthesisTaskResponse struct {
	*StartSpeechSynthesisTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartSpeechSynthesisTask request.
func (r *StartSpeechSynthesisTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
