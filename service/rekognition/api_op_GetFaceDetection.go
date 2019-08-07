// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetFaceDetectionInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for the face detection job. The JobId is returned from
	// StartFaceDetection.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`

	// Maximum number of results to return per paginated call. The largest value
	// you can specify is 1000. If you specify a value greater than 1000, a maximum
	// of 1000 results is returned. The default value is 1000.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous response was incomplete (because there are more faces to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of faces.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetFaceDetectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFaceDetectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFaceDetectionInput"}

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

type GetFaceDetectionOutput struct {
	_ struct{} `type:"structure"`

	// An array of faces detected in the video. Each element contains a detected
	// face's details and the time, in milliseconds from the start of the video,
	// the face was detected.
	Faces []FaceDetection `json:"rekognition:GetFaceDetectionOutput:Faces" type:"list"`

	// The current status of the face detection job.
	JobStatus VideoJobStatus `json:"rekognition:GetFaceDetectionOutput:JobStatus" type:"string" enum:"true"`

	// If the response is truncated, Amazon Rekognition returns this token that
	// you can use in the subsequent request to retrieve the next set of faces.
	NextToken *string `json:"rekognition:GetFaceDetectionOutput:NextToken" type:"string"`

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string `json:"rekognition:GetFaceDetectionOutput:StatusMessage" type:"string"`

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition
	// video operation.
	VideoMetadata *VideoMetadata `json:"rekognition:GetFaceDetectionOutput:VideoMetadata" type:"structure"`
}

// String returns the string representation
func (s GetFaceDetectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetFaceDetection = "GetFaceDetection"

// GetFaceDetectionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Gets face detection results for a Amazon Rekognition Video analysis started
// by StartFaceDetection.
//
// Face detection with Amazon Rekognition Video is an asynchronous operation.
// You start face detection by calling StartFaceDetection which returns a job
// identifier (JobId). When the face detection operation finishes, Amazon Rekognition
// Video publishes a completion status to the Amazon Simple Notification Service
// topic registered in the initial call to StartFaceDetection. To get the results
// of the face detection operation, first check that the status value published
// to the Amazon SNS topic is SUCCEEDED. If so, call GetFaceDetection and pass
// the job identifier (JobId) from the initial call to StartFaceDetection.
//
// GetFaceDetection returns an array of detected faces (Faces) sorted by the
// time the faces were detected.
//
// Use MaxResults parameter to limit the number of labels returned. If there
// are more results than specified in MaxResults, the value of NextToken in
// the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetFaceDetection and populate
// the NextToken request parameter with the token value returned from the previous
// call to GetFaceDetection.
//
//    // Example sending a request using GetFaceDetectionRequest.
//    req := client.GetFaceDetectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetFaceDetectionRequest(input *GetFaceDetectionInput) GetFaceDetectionRequest {
	op := &aws.Operation{
		Name:       opGetFaceDetection,
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
		input = &GetFaceDetectionInput{}
	}

	req := c.newRequest(op, input, &GetFaceDetectionOutput{})
	return GetFaceDetectionRequest{Request: req, Input: input, Copy: c.GetFaceDetectionRequest}
}

// GetFaceDetectionRequest is the request type for the
// GetFaceDetection API operation.
type GetFaceDetectionRequest struct {
	*aws.Request
	Input *GetFaceDetectionInput
	Copy  func(*GetFaceDetectionInput) GetFaceDetectionRequest
}

// Send marshals and sends the GetFaceDetection API request.
func (r GetFaceDetectionRequest) Send(ctx context.Context) (*GetFaceDetectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFaceDetectionResponse{
		GetFaceDetectionOutput: r.Request.Data.(*GetFaceDetectionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetFaceDetectionRequestPaginator returns a paginator for GetFaceDetection.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetFaceDetectionRequest(input)
//   p := rekognition.NewGetFaceDetectionRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetFaceDetectionPaginator(req GetFaceDetectionRequest) GetFaceDetectionPaginator {
	return GetFaceDetectionPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetFaceDetectionInput
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

// GetFaceDetectionPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetFaceDetectionPaginator struct {
	aws.Pager
}

func (p *GetFaceDetectionPaginator) CurrentPage() *GetFaceDetectionOutput {
	return p.Pager.CurrentPage().(*GetFaceDetectionOutput)
}

// GetFaceDetectionResponse is the response type for the
// GetFaceDetection API operation.
type GetFaceDetectionResponse struct {
	*GetFaceDetectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFaceDetection request.
func (r *GetFaceDetectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
