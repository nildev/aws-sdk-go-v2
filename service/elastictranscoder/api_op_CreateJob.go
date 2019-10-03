// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The CreateJobRequest structure.
type CreateJobInput struct {
	_ struct{} `type:"structure"`

	// A section of the request body that provides information about the file that
	// is being transcoded.
	Input *JobInput `type:"structure"`

	// A section of the request body that provides information about the files that
	// are being transcoded.
	Inputs []JobInput `type:"list"`

	// A section of the request body that provides information about the transcoded
	// (target) file. We strongly recommend that you use the Outputs syntax instead
	// of the Output syntax.
	Output *CreateJobOutputResult `type:"structure"`

	// The value, if any, that you want Elastic Transcoder to prepend to the names
	// of all files that this job creates, including output files, thumbnails, and
	// playlists.
	OutputKeyPrefix *string `min:"1" type:"string"`

	// A section of the request body that provides information about the transcoded
	// (target) files. We recommend that you use the Outputs syntax instead of the
	// Output syntax.
	Outputs []CreateJobOutputResult `type:"list"`

	// The Id of the pipeline that you want Elastic Transcoder to use for transcoding.
	// The pipeline determines several settings, including the Amazon S3 bucket
	// from which Elastic Transcoder gets the files to transcode and the bucket
	// into which Elastic Transcoder puts the transcoded files.
	//
	// PipelineId is a required field
	PipelineId *string `type:"string" required:"true"`

	// If you specify a preset in PresetId for which the value of Container is fmp4
	// (Fragmented MP4) or ts (MPEG-TS), Playlists contains information about the
	// master playlists that you want Elastic Transcoder to create.
	//
	// The maximum number of master playlists in a job is 30.
	Playlists []CreateJobPlaylist `type:"list"`

	// User-defined metadata that you want to associate with an Elastic Transcoder
	// job. You specify metadata in key/value pairs, and you can add up to 10 key/value
	// pairs per job. Elastic Transcoder does not guarantee that key/value pairs
	// are returned in the same order in which you specify them.
	UserMetadata map[string]string `type:"map"`
}

// String returns the string representation
func (s CreateJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateJobInput"}
	if s.OutputKeyPrefix != nil && len(*s.OutputKeyPrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OutputKeyPrefix", 1))
	}

	if s.PipelineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineId"))
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			invalidParams.AddNested("Input", err.(aws.ErrInvalidParams))
		}
	}
	if s.Inputs != nil {
		for i, v := range s.Inputs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Inputs", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			invalidParams.AddNested("Output", err.(aws.ErrInvalidParams))
		}
	}
	if s.Outputs != nil {
		for i, v := range s.Outputs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Outputs", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Playlists != nil {
		for i, v := range s.Playlists {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Playlists", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateJobInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Input != nil {
		v := s.Input

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Input", v, metadata)
	}
	if s.Inputs != nil {
		v := s.Inputs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Inputs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Output != nil {
		v := s.Output

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Output", v, metadata)
	}
	if s.OutputKeyPrefix != nil {
		v := *s.OutputKeyPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OutputKeyPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Outputs != nil {
		v := s.Outputs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Outputs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.PipelineId != nil {
		v := *s.PipelineId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PipelineId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Playlists != nil {
		v := s.Playlists

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Playlists", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.UserMetadata != nil {
		v := s.UserMetadata

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "UserMetadata", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// The CreateJobResponse structure.
type CreateJobOutput struct {
	_ struct{} `type:"structure"`

	// A section of the response body that provides information about the job that
	// is created.
	Job *Job `json:"elastictranscoder:CreateJobOutput:Job" type:"structure"`
}

// String returns the string representation
func (s CreateJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Job != nil {
		v := s.Job

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Job", v, metadata)
	}
	return nil
}

const opCreateJob = "CreateJob"

// CreateJobRequest returns a request value for making API operation for
// Amazon Elastic Transcoder.
//
// When you create a job, Elastic Transcoder returns JSON data that includes
// the values that you specified plus information about the job that is created.
//
// If you have specified more than one output for your jobs (for example, one
// output for the Kindle Fire and another output for the Apple iPhone 4s), you
// currently must use the Elastic Transcoder API to list the jobs (as opposed
// to the AWS Console).
//
//    // Example sending a request using CreateJobRequest.
//    req := client.CreateJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateJobRequest(input *CreateJobInput) CreateJobRequest {
	op := &aws.Operation{
		Name:       opCreateJob,
		HTTPMethod: "POST",
		HTTPPath:   "/2012-09-25/jobs",
	}

	if input == nil {
		input = &CreateJobInput{}
	}

	req := c.newRequest(op, input, &CreateJobOutput{})
	return CreateJobRequest{Request: req, Input: input, Copy: c.CreateJobRequest}
}

// CreateJobRequest is the request type for the
// CreateJob API operation.
type CreateJobRequest struct {
	*aws.Request
	Input *CreateJobInput
	Copy  func(*CreateJobInput) CreateJobRequest
}

// Send marshals and sends the CreateJob API request.
func (r CreateJobRequest) Send(ctx context.Context) (*CreateJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateJobResponse{
		CreateJobOutput: r.Request.Data.(*CreateJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateJobResponse is the response type for the
// CreateJob API operation.
type CreateJobResponse struct {
	*CreateJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateJob request.
func (r *CreateJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
