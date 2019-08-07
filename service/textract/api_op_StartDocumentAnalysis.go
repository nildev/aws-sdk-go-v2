// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package textract

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/textract-2018-06-27/StartDocumentAnalysisRequest
type StartDocumentAnalysisInput struct {
	_ struct{} `type:"structure"`

	// The idempotent token that you use to identify the start request. If you use
	// the same token with multiple StartDocumentAnalysis requests, the same JobId
	// is returned. Use ClientRequestToken to prevent the same job from being accidentally
	// started more than once.
	ClientRequestToken *string `min:"1" type:"string"`

	// The location of the document to be processed.
	//
	// DocumentLocation is a required field
	DocumentLocation *DocumentLocation `type:"structure" required:"true"`

	// A list of the types of analysis to perform. Add TABLES to the list to return
	// information about the tables that are detected in the input document. Add
	// FORMS to return detected fields and the associated text. To perform both
	// types of analysis, add TABLES and FORMS to FeatureTypes. All selectable elements
	// (SELECTION_ELEMENT) that are detected are returned, whatever the value of
	// FeatureTypes.
	//
	// FeatureTypes is a required field
	FeatureTypes []FeatureType `type:"list" required:"true"`

	// An identifier you specify that's included in the completion notification
	// that's published to the Amazon SNS topic. For example, you can use JobTag
	// to identify the type of document, such as a tax form or a receipt, that the
	// completion notification corresponds to.
	JobTag *string `min:"1" type:"string"`

	// The Amazon SNS topic ARN that you want Amazon Textract to publish the completion
	// status of the operation to.
	NotificationChannel *NotificationChannel `type:"structure"`
}

// String returns the string representation
func (s StartDocumentAnalysisInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDocumentAnalysisInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartDocumentAnalysisInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.DocumentLocation == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentLocation"))
	}

	if s.FeatureTypes == nil {
		invalidParams.Add(aws.NewErrParamRequired("FeatureTypes"))
	}
	if s.JobTag != nil && len(*s.JobTag) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobTag", 1))
	}
	if s.DocumentLocation != nil {
		if err := s.DocumentLocation.Validate(); err != nil {
			invalidParams.AddNested("DocumentLocation", err.(aws.ErrInvalidParams))
		}
	}
	if s.NotificationChannel != nil {
		if err := s.NotificationChannel.Validate(); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/textract-2018-06-27/StartDocumentAnalysisResponse
type StartDocumentAnalysisOutput struct {
	_ struct{} `type:"structure"`

	// The identifier for the document text detection job. Use JobId to identify
	// the job in a subsequent call to GetDocumentAnalysis.
	JobId *string `json:"textract:StartDocumentAnalysisOutput:JobId" min:"1" type:"string"`
}

// String returns the string representation
func (s StartDocumentAnalysisOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartDocumentAnalysis = "StartDocumentAnalysis"

// StartDocumentAnalysisRequest returns a request value for making API operation for
// Amazon Textract.
//
// Starts asynchronous analysis of an input document for relationships between
// detected items such as key and value pairs, tables, and selection elements.
//
// StartDocumentAnalysis can analyze text in documents that are in JPG, PNG,
// and PDF format. The documents are stored in an Amazon S3 bucket. Use DocumentLocation
// to specify the bucket name and file name of the document.
//
// StartDocumentAnalysis returns a job identifier (JobId) that you use to get
// the results of the operation. When text analysis is finished, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that you specify in NotificationChannel. To get the results of
// the text analysis operation, first check that the status value published
// to the Amazon SNS topic is SUCCEEDED. If so, call GetDocumentAnalysis, and
// pass the job identifier (JobId) from the initial call to StartDocumentAnalysis.
//
// For more information, see Document Text Analysis (https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html).
//
//    // Example sending a request using StartDocumentAnalysisRequest.
//    req := client.StartDocumentAnalysisRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/textract-2018-06-27/StartDocumentAnalysis
func (c *Client) StartDocumentAnalysisRequest(input *StartDocumentAnalysisInput) StartDocumentAnalysisRequest {
	op := &aws.Operation{
		Name:       opStartDocumentAnalysis,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartDocumentAnalysisInput{}
	}

	req := c.newRequest(op, input, &StartDocumentAnalysisOutput{})
	return StartDocumentAnalysisRequest{Request: req, Input: input, Copy: c.StartDocumentAnalysisRequest}
}

// StartDocumentAnalysisRequest is the request type for the
// StartDocumentAnalysis API operation.
type StartDocumentAnalysisRequest struct {
	*aws.Request
	Input *StartDocumentAnalysisInput
	Copy  func(*StartDocumentAnalysisInput) StartDocumentAnalysisRequest
}

// Send marshals and sends the StartDocumentAnalysis API request.
func (r StartDocumentAnalysisRequest) Send(ctx context.Context) (*StartDocumentAnalysisResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDocumentAnalysisResponse{
		StartDocumentAnalysisOutput: r.Request.Data.(*StartDocumentAnalysisOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDocumentAnalysisResponse is the response type for the
// StartDocumentAnalysis API operation.
type StartDocumentAnalysisResponse struct {
	*StartDocumentAnalysisOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDocumentAnalysis request.
func (r *StartDocumentAnalysisResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
