// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/GetJobManifestRequest
type GetJobManifestInput struct {
	_ struct{} `type:"structure"`

	// The ID for a job that you want to get the manifest file for, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	//
	// JobId is a required field
	JobId *string `min:"39" type:"string" required:"true"`
}

// String returns the string representation
func (s GetJobManifestInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJobManifestInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJobManifestInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 39 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 39))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/GetJobManifestResult
type GetJobManifestOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 presigned URL for the manifest file associated with the specified
	// JobId value.
	ManifestURI *string `json:"snowball:GetJobManifestOutput:ManifestURI" min:"1" type:"string"`
}

// String returns the string representation
func (s GetJobManifestOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetJobManifest = "GetJobManifest"

// GetJobManifestRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// Returns a link to an Amazon S3 presigned URL for the manifest file associated
// with the specified JobId value. You can access the manifest file for up to
// 60 minutes after this request has been made. To access the manifest file
// after 60 minutes have passed, you'll have to make another call to the GetJobManifest
// action.
//
// The manifest is an encrypted file that you can download after your job enters
// the WithCustomer status. The manifest is decrypted by using the UnlockCode
// code value, when you pass both values to the Snowball through the Snowball
// client when the client is started for the first time.
//
// As a best practice, we recommend that you don't save a copy of an UnlockCode
// value in the same location as the manifest file for that job. Saving these
// separately helps prevent unauthorized parties from gaining access to the
// Snowball associated with that job.
//
// The credentials of a given job, including its manifest file and unlock code,
// expire 90 days after the job is created.
//
//    // Example sending a request using GetJobManifestRequest.
//    req := client.GetJobManifestRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/GetJobManifest
func (c *Client) GetJobManifestRequest(input *GetJobManifestInput) GetJobManifestRequest {
	op := &aws.Operation{
		Name:       opGetJobManifest,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetJobManifestInput{}
	}

	req := c.newRequest(op, input, &GetJobManifestOutput{})
	return GetJobManifestRequest{Request: req, Input: input, Copy: c.GetJobManifestRequest}
}

// GetJobManifestRequest is the request type for the
// GetJobManifest API operation.
type GetJobManifestRequest struct {
	*aws.Request
	Input *GetJobManifestInput
	Copy  func(*GetJobManifestInput) GetJobManifestRequest
}

// Send marshals and sends the GetJobManifest API request.
func (r GetJobManifestRequest) Send(ctx context.Context) (*GetJobManifestResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJobManifestResponse{
		GetJobManifestOutput: r.Request.Data.(*GetJobManifestOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJobManifestResponse is the response type for the
// GetJobManifest API operation.
type GetJobManifestResponse struct {
	*GetJobManifestOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJobManifest request.
func (r *GetJobManifestResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
