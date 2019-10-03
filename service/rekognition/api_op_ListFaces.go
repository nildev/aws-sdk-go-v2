// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListFacesInput struct {
	_ struct{} `type:"structure"`

	// ID of the collection from which to list the faces.
	//
	// CollectionId is a required field
	CollectionId *string `min:"1" type:"string" required:"true"`

	// Maximum number of faces to return.
	MaxResults *int64 `type:"integer"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Rekognition returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of faces.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListFacesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFacesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFacesInput"}

	if s.CollectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CollectionId"))
	}
	if s.CollectionId != nil && len(*s.CollectionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CollectionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListFacesOutput struct {
	_ struct{} `type:"structure"`

	// Version number of the face detection model associated with the input collection
	// (CollectionId).
	FaceModelVersion *string `json:"rekognition:ListFacesOutput:FaceModelVersion" type:"string"`

	// An array of Face objects.
	Faces []Face `json:"rekognition:ListFacesOutput:Faces" type:"list"`

	// If the response is truncated, Amazon Rekognition returns this token that
	// you can use in the subsequent request to retrieve the next set of faces.
	NextToken *string `json:"rekognition:ListFacesOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s ListFacesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListFaces = "ListFaces"

// ListFacesRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Returns metadata for faces in the specified collection. This metadata includes
// information such as the bounding box coordinates, the confidence (that the
// bounding box contains a face), and face ID. For an example, see Listing Faces
// in a Collection in the Amazon Rekognition Developer Guide.
//
// This operation requires permissions to perform the rekognition:ListFaces
// action.
//
//    // Example sending a request using ListFacesRequest.
//    req := client.ListFacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListFacesRequest(input *ListFacesInput) ListFacesRequest {
	op := &aws.Operation{
		Name:       opListFaces,
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
		input = &ListFacesInput{}
	}

	req := c.newRequest(op, input, &ListFacesOutput{})
	return ListFacesRequest{Request: req, Input: input, Copy: c.ListFacesRequest}
}

// ListFacesRequest is the request type for the
// ListFaces API operation.
type ListFacesRequest struct {
	*aws.Request
	Input *ListFacesInput
	Copy  func(*ListFacesInput) ListFacesRequest
}

// Send marshals and sends the ListFaces API request.
func (r ListFacesRequest) Send(ctx context.Context) (*ListFacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFacesResponse{
		ListFacesOutput: r.Request.Data.(*ListFacesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFacesRequestPaginator returns a paginator for ListFaces.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFacesRequest(input)
//   p := rekognition.NewListFacesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFacesPaginator(req ListFacesRequest) ListFacesPaginator {
	return ListFacesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFacesInput
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

// ListFacesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFacesPaginator struct {
	aws.Pager
}

func (p *ListFacesPaginator) CurrentPage() *ListFacesOutput {
	return p.Pager.CurrentPage().(*ListFacesOutput)
}

// ListFacesResponse is the response type for the
// ListFaces API operation.
type ListFacesResponse struct {
	*ListFacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFaces request.
func (r *ListFacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
