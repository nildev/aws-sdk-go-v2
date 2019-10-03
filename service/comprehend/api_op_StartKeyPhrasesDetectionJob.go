// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/StartKeyPhrasesDetectionJobRequest
type StartKeyPhrasesDetectionJobInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the request. If you don't set the client request
	// token, Amazon Comprehend generates one.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management
	// (IAM) role that grants Amazon Comprehend read access to your input data.
	// For more information, see https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions.html#auth-role-permissions
	// (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions.html#auth-role-permissions).
	//
	// DataAccessRoleArn is a required field
	DataAccessRoleArn *string `min:"20" type:"string" required:"true"`

	// Specifies the format and location of the input data for the job.
	//
	// InputDataConfig is a required field
	InputDataConfig *InputDataConfig `type:"structure" required:"true"`

	// The identifier of the job.
	JobName *string `min:"1" type:"string"`

	// The language of the input documents. You can specify any of the primary languages
	// supported by Amazon Comprehend: German ("de"), English ("en"), Spanish ("es"),
	// French ("fr"), Italian ("it"), or Portuguese ("pt"). All documents must be
	// in the same language.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// Specifies where to send the output files.
	//
	// OutputDataConfig is a required field
	OutputDataConfig *OutputDataConfig `type:"structure" required:"true"`

	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses
	// to encrypt data on the storage volume attached to the ML compute instance(s)
	// that process the analysis job. The VolumeKmsKeyId can be either of the following
	// formats:
	//
	//    * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	//    * Amazon Resource Name (ARN) of a KMS Key: "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	VolumeKmsKeyId *string `type:"string"`

	// Configuration parameters for an optional private Virtual Private Cloud (VPC)
	// containing the resources you are using for your key phrases detection job.
	// For more information, see Amazon VPC (https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).
	VpcConfig *VpcConfig `type:"structure"`
}

// String returns the string representation
func (s StartKeyPhrasesDetectionJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartKeyPhrasesDetectionJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartKeyPhrasesDetectionJobInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.DataAccessRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataAccessRoleArn"))
	}
	if s.DataAccessRoleArn != nil && len(*s.DataAccessRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("DataAccessRoleArn", 20))
	}

	if s.InputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputDataConfig"))
	}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.OutputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputDataConfig"))
	}
	if s.InputDataConfig != nil {
		if err := s.InputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.OutputDataConfig != nil {
		if err := s.OutputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			invalidParams.AddNested("VpcConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/StartKeyPhrasesDetectionJobResponse
type StartKeyPhrasesDetectionJobOutput struct {
	_ struct{} `type:"structure"`

	// The identifier generated for the job. To get the status of a job, use this
	// identifier with the operation.
	JobId *string `json:"comprehend:StartKeyPhrasesDetectionJobOutput:JobId" min:"1" type:"string"`

	// The status of the job.
	//
	//    * SUBMITTED - The job has been received and is queued for processing.
	//
	//    * IN_PROGRESS - Amazon Comprehend is processing the job.
	//
	//    * COMPLETED - The job was successfully completed and the output is available.
	//
	//    * FAILED - The job did not complete. To get details, use the operation.
	JobStatus JobStatus `json:"comprehend:StartKeyPhrasesDetectionJobOutput:JobStatus" type:"string" enum:"true"`
}

// String returns the string representation
func (s StartKeyPhrasesDetectionJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartKeyPhrasesDetectionJob = "StartKeyPhrasesDetectionJob"

// StartKeyPhrasesDetectionJobRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Starts an asynchronous key phrase detection job for a collection of documents.
// Use the operation to track the status of a job.
//
//    // Example sending a request using StartKeyPhrasesDetectionJobRequest.
//    req := client.StartKeyPhrasesDetectionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/StartKeyPhrasesDetectionJob
func (c *Client) StartKeyPhrasesDetectionJobRequest(input *StartKeyPhrasesDetectionJobInput) StartKeyPhrasesDetectionJobRequest {
	op := &aws.Operation{
		Name:       opStartKeyPhrasesDetectionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartKeyPhrasesDetectionJobInput{}
	}

	req := c.newRequest(op, input, &StartKeyPhrasesDetectionJobOutput{})
	return StartKeyPhrasesDetectionJobRequest{Request: req, Input: input, Copy: c.StartKeyPhrasesDetectionJobRequest}
}

// StartKeyPhrasesDetectionJobRequest is the request type for the
// StartKeyPhrasesDetectionJob API operation.
type StartKeyPhrasesDetectionJobRequest struct {
	*aws.Request
	Input *StartKeyPhrasesDetectionJobInput
	Copy  func(*StartKeyPhrasesDetectionJobInput) StartKeyPhrasesDetectionJobRequest
}

// Send marshals and sends the StartKeyPhrasesDetectionJob API request.
func (r StartKeyPhrasesDetectionJobRequest) Send(ctx context.Context) (*StartKeyPhrasesDetectionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartKeyPhrasesDetectionJobResponse{
		StartKeyPhrasesDetectionJobOutput: r.Request.Data.(*StartKeyPhrasesDetectionJobOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartKeyPhrasesDetectionJobResponse is the response type for the
// StartKeyPhrasesDetectionJob API operation.
type StartKeyPhrasesDetectionJobResponse struct {
	*StartKeyPhrasesDetectionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartKeyPhrasesDetectionJob request.
func (r *StartKeyPhrasesDetectionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
