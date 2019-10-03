// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetPersonTrackingInput struct {
	_ struct{} `type:"structure"`

	// The identifier for a job that tracks persons in a video. You get the JobId
	// from a call to StartPersonTracking.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`

	// Maximum number of results to return per paginated call. The largest value
	// you can specify is 1000. If you specify a value greater than 1000, a maximum
	// of 1000 results is returned. The default value is 1000.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous response was incomplete (because there are more persons to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of persons.
	NextToken *string `type:"string"`

	// Sort to use for elements in the Persons array. Use TIMESTAMP to sort array
	// elements by the time persons are detected. Use INDEX to sort by the tracked
	// persons. If you sort by INDEX, the array elements for each person are sorted
	// by detection confidence. The default sort is by TIMESTAMP.
	SortBy PersonTrackingSortBy `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetPersonTrackingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPersonTrackingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPersonTrackingInput"}

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

type GetPersonTrackingOutput struct {
	_ struct{} `type:"structure"`

	// The current status of the person tracking job.
	JobStatus VideoJobStatus `json:"rekognition:GetPersonTrackingOutput:JobStatus" type:"string" enum:"true"`

	// If the response is truncated, Amazon Rekognition Video returns this token
	// that you can use in the subsequent request to retrieve the next set of persons.
	NextToken *string `json:"rekognition:GetPersonTrackingOutput:NextToken" type:"string"`

	// An array of the persons detected in the video and the time(s) their path
	// was tracked throughout the video. An array element will exist for each time
	// a person's path is tracked.
	Persons []PersonDetection `json:"rekognition:GetPersonTrackingOutput:Persons" type:"list"`

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string `json:"rekognition:GetPersonTrackingOutput:StatusMessage" type:"string"`

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition
	// Video operation.
	VideoMetadata *VideoMetadata `json:"rekognition:GetPersonTrackingOutput:VideoMetadata" type:"structure"`
}

// String returns the string representation
func (s GetPersonTrackingOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPersonTracking = "GetPersonTracking"

// GetPersonTrackingRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Gets the path tracking results of a Amazon Rekognition Video analysis started
// by StartPersonTracking.
//
// The person path tracking operation is started by a call to StartPersonTracking
// which returns a job identifier (JobId). When the operation finishes, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple Notification
// Service topic registered in the initial call to StartPersonTracking.
//
// To get the results of the person path tracking operation, first check that
// the status value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetPersonTracking and pass the job identifier (JobId) from the initial call
// to StartPersonTracking.
//
// GetPersonTracking returns an array, Persons, of tracked persons and the time(s)
// their paths were tracked in the video.
//
// GetPersonTracking only returns the default facial attributes (BoundingBox,
// Confidence, Landmarks, Pose, and Quality). The other facial attributes listed
// in the Face object of the following response syntax are not returned.
//
// For more information, see FaceDetail in the Amazon Rekognition Developer
// Guide.
//
// By default, the array is sorted by the time(s) a person's path is tracked
// in the video. You can sort by tracked persons by specifying INDEX for the
// SortBy input parameter.
//
// Use the MaxResults parameter to limit the number of items returned. If there
// are more results than specified in MaxResults, the value of NextToken in
// the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetPersonTracking and populate
// the NextToken request parameter with the token value returned from the previous
// call to GetPersonTracking.
//
//    // Example sending a request using GetPersonTrackingRequest.
//    req := client.GetPersonTrackingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetPersonTrackingRequest(input *GetPersonTrackingInput) GetPersonTrackingRequest {
	op := &aws.Operation{
		Name:       opGetPersonTracking,
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
		input = &GetPersonTrackingInput{}
	}

	req := c.newRequest(op, input, &GetPersonTrackingOutput{})
	return GetPersonTrackingRequest{Request: req, Input: input, Copy: c.GetPersonTrackingRequest}
}

// GetPersonTrackingRequest is the request type for the
// GetPersonTracking API operation.
type GetPersonTrackingRequest struct {
	*aws.Request
	Input *GetPersonTrackingInput
	Copy  func(*GetPersonTrackingInput) GetPersonTrackingRequest
}

// Send marshals and sends the GetPersonTracking API request.
func (r GetPersonTrackingRequest) Send(ctx context.Context) (*GetPersonTrackingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPersonTrackingResponse{
		GetPersonTrackingOutput: r.Request.Data.(*GetPersonTrackingOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetPersonTrackingRequestPaginator returns a paginator for GetPersonTracking.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetPersonTrackingRequest(input)
//   p := rekognition.NewGetPersonTrackingRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetPersonTrackingPaginator(req GetPersonTrackingRequest) GetPersonTrackingPaginator {
	return GetPersonTrackingPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetPersonTrackingInput
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

// GetPersonTrackingPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetPersonTrackingPaginator struct {
	aws.Pager
}

func (p *GetPersonTrackingPaginator) CurrentPage() *GetPersonTrackingOutput {
	return p.Pager.CurrentPage().(*GetPersonTrackingOutput)
}

// GetPersonTrackingResponse is the response type for the
// GetPersonTracking API operation.
type GetPersonTrackingResponse struct {
	*GetPersonTrackingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPersonTracking request.
func (r *GetPersonTrackingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
