// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/UpdateVocabularyRequest
type UpdateVocabularyInput struct {
	_ struct{} `type:"structure"`

	// The language code of the vocabulary entries.
	//
	// LanguageCode is a required field
	LanguageCode LanguageCode `type:"string" required:"true" enum:"true"`

	// An array of strings containing the vocabulary entries.
	Phrases []string `type:"list"`

	// The S3 location of the text file that contains the definition of the custom
	// vocabulary. The URI must be in the same region as the API endpoint that you
	// are calling. The general form is
	//
	// https://s3-<aws-region>.amazonaws.com/<bucket-name>/<keyprefix>/<objectkey>
	//
	// For example:
	//
	// https://s3-us-east-1.amazonaws.com/examplebucket/vocab.txt
	//
	// For more information about S3 object names, see Object Keys (http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#object-keys)
	// in the Amazon S3 Developer Guide.
	//
	// For more information about custom vocabularies, see Custom Vocabularies (http://docs.aws.amazon.com/transcribe/latest/dg/how-it-works.html#how-vocabulary).
	VocabularyFileUri *string `min:"1" type:"string"`

	// The name of the vocabulary to update. The name is case-sensitive.
	//
	// VocabularyName is a required field
	VocabularyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateVocabularyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVocabularyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateVocabularyInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}
	if s.VocabularyFileUri != nil && len(*s.VocabularyFileUri) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyFileUri", 1))
	}

	if s.VocabularyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VocabularyName"))
	}
	if s.VocabularyName != nil && len(*s.VocabularyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/UpdateVocabularyResponse
type UpdateVocabularyOutput struct {
	_ struct{} `type:"structure"`

	// The language code of the vocabulary entries.
	LanguageCode LanguageCode `json:"transcribe:UpdateVocabularyOutput:LanguageCode" type:"string" enum:"true"`

	// The date and time that the vocabulary was updated.
	LastModifiedTime *time.Time `json:"transcribe:UpdateVocabularyOutput:LastModifiedTime" type:"timestamp" timestampFormat:"unix"`

	// The name of the vocabulary that was updated.
	VocabularyName *string `json:"transcribe:UpdateVocabularyOutput:VocabularyName" min:"1" type:"string"`

	// The processing state of the vocabulary. When the VocabularyState field contains
	// READY the vocabulary is ready to be used in a StartTranscriptionJob request.
	VocabularyState VocabularyState `json:"transcribe:UpdateVocabularyOutput:VocabularyState" type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateVocabularyOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateVocabulary = "UpdateVocabulary"

// UpdateVocabularyRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Updates an existing vocabulary with new values. The UpdateVocabulary operation
// overwrites all of the existing information with the values that you provide
// in the request.
//
//    // Example sending a request using UpdateVocabularyRequest.
//    req := client.UpdateVocabularyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/UpdateVocabulary
func (c *Client) UpdateVocabularyRequest(input *UpdateVocabularyInput) UpdateVocabularyRequest {
	op := &aws.Operation{
		Name:       opUpdateVocabulary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateVocabularyInput{}
	}

	req := c.newRequest(op, input, &UpdateVocabularyOutput{})
	return UpdateVocabularyRequest{Request: req, Input: input, Copy: c.UpdateVocabularyRequest}
}

// UpdateVocabularyRequest is the request type for the
// UpdateVocabulary API operation.
type UpdateVocabularyRequest struct {
	*aws.Request
	Input *UpdateVocabularyInput
	Copy  func(*UpdateVocabularyInput) UpdateVocabularyRequest
}

// Send marshals and sends the UpdateVocabulary API request.
func (r UpdateVocabularyRequest) Send(ctx context.Context) (*UpdateVocabularyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVocabularyResponse{
		UpdateVocabularyOutput: r.Request.Data.(*UpdateVocabularyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVocabularyResponse is the response type for the
// UpdateVocabulary API operation.
type UpdateVocabularyResponse struct {
	*UpdateVocabularyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVocabulary request.
func (r *UpdateVocabularyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
