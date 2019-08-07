// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetLabelDetectionInput struct {
	_ struct{} `type:"structure"`

	// Job identifier for the label detection operation for which you want results
	// returned. You get the job identifer from an initial call to StartlabelDetection.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`

	// Maximum number of results to return per paginated call. The largest value
	// you can specify is 1000. If you specify a value greater than 1000, a maximum
	// of 1000 results is returned. The default value is 1000.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous response was incomplete (because there are more labels to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of labels.
	NextToken *string `type:"string"`

	// Sort to use for elements in the Labels array. Use TIMESTAMP to sort array
	// elements by the time labels are detected. Use NAME to alphabetically group
	// elements for a label together. Within each label group, the array element
	// are sorted by detection confidence. The default sort is by TIMESTAMP.
	SortBy LabelDetectionSortBy `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetLabelDetectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLabelDetectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLabelDetectionInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLabelDetectionOutput struct {
	_ struct{} `type:"structure"`

	// The current status of the label detection job.
	JobStatus VideoJobStatus `json:"rekognition:GetLabelDetectionOutput:JobStatus" type:"string" enum:"true"`

	// Version number of the label detection model that was used to detect labels.
	LabelModelVersion *string `json:"rekognition:GetLabelDetectionOutput:LabelModelVersion" type:"string"`

	// An array of labels detected in the video. Each element contains the detected
	// label and the time, in milliseconds from the start of the video, that the
	// label was detected.
	Labels []LabelDetection `json:"rekognition:GetLabelDetectionOutput:Labels" type:"list"`

	// If the response is truncated, Amazon Rekognition Video returns this token
	// that you can use in the subsequent request to retrieve the next set of labels.
	NextToken *string `json:"rekognition:GetLabelDetectionOutput:NextToken" type:"string"`

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string `json:"rekognition:GetLabelDetectionOutput:StatusMessage" type:"string"`

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition
	// video operation.
	VideoMetadata *VideoMetadata `json:"rekognition:GetLabelDetectionOutput:VideoMetadata" type:"structure"`
}

// String returns the string representation
func (s GetLabelDetectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetLabelDetection = "GetLabelDetection"

// GetLabelDetectionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Gets the label detection results of a Amazon Rekognition Video analysis started
// by StartLabelDetection.
//
// The label detection operation is started by a call to StartLabelDetection
// which returns a job identifier (JobId). When the label detection operation
// finishes, Amazon Rekognition publishes a completion status to the Amazon
// Simple Notification Service topic registered in the initial call to StartlabelDetection.
// To get the results of the label detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetLabelDetection and pass the job identifier (JobId) from the initial call
// to StartLabelDetection.
//
// GetLabelDetection returns an array of detected labels (Labels) sorted by
// the time the labels were detected. You can also sort by the label name by
// specifying NAME for the SortBy input parameter.
//
// The labels returned include the label name, the percentage confidence in
// the accuracy of the detected label, and the time the label was detected in
// the video.
//
// The returned labels also include bounding box information for common objects,
// a hierarchical taxonomy of detected labels, and the version of the label
// model used for detection.
//
// Use MaxResults parameter to limit the number of labels returned. If there
// are more results than specified in MaxResults, the value of NextToken in
// the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetlabelDetection and populate
// the NextToken request parameter with the token value returned from the previous
// call to GetLabelDetection.
//
//    // Example sending a request using GetLabelDetectionRequest.
//    req := client.GetLabelDetectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetLabelDetectionRequest(input *GetLabelDetectionInput) GetLabelDetectionRequest {
	op := &aws.Operation{
		Name:       opGetLabelDetection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetLabelDetectionInput{}
	}

	req := c.newRequest(op, input, &GetLabelDetectionOutput{})
	return GetLabelDetectionRequest{Request: req, Input: input, Copy: c.GetLabelDetectionRequest}
}

// GetLabelDetectionRequest is the request type for the
// GetLabelDetection API operation.
type GetLabelDetectionRequest struct {
	*aws.Request
	Input *GetLabelDetectionInput
	Copy  func(*GetLabelDetectionInput) GetLabelDetectionRequest
}

// Send marshals and sends the GetLabelDetection API request.
func (r GetLabelDetectionRequest) Send(ctx context.Context) (*GetLabelDetectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLabelDetectionResponse{
		GetLabelDetectionOutput: r.Request.Data.(*GetLabelDetectionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetLabelDetectionRequestPaginator returns a paginator for GetLabelDetection.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetLabelDetectionRequest(input)
//   p := rekognition.NewGetLabelDetectionRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetLabelDetectionPaginator(req GetLabelDetectionRequest) GetLabelDetectionPaginator {
	return GetLabelDetectionPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetLabelDetectionInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetLabelDetectionPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetLabelDetectionPaginator struct {
	aws.Pager
}

func (p *GetLabelDetectionPaginator) CurrentPage() *GetLabelDetectionOutput {
	return p.Pager.CurrentPage().(*GetLabelDetectionOutput)
}

// GetLabelDetectionResponse is the response type for the
// GetLabelDetection API operation.
type GetLabelDetectionResponse struct {
	*GetLabelDetectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLabelDetection request.
func (r *GetLabelDetectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
