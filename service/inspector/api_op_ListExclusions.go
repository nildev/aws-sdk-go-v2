// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/ListExclusionsRequest
type ListExclusionsInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the assessment run that generated the exclusions that you want
	// to list.
	//
	// AssessmentRunArn is a required field
	AssessmentRunArn *string `locationName:"assessmentRunArn" min:"1" type:"string" required:"true"`

	// You can use this parameter to indicate the maximum number of items you want
	// in the response. The default value is 100. The maximum value is 500.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListExclusionsRequest action.
	// Subsequent calls to the action fill nextToken in the request with the value
	// of nextToken from the previous response to continue listing data.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListExclusionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListExclusionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListExclusionsInput"}

	if s.AssessmentRunArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentRunArn"))
	}
	if s.AssessmentRunArn != nil && len(*s.AssessmentRunArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentRunArn", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/ListExclusionsResponse
type ListExclusionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of exclusions' ARNs returned by the action.
	//
	// ExclusionArns is a required field
	ExclusionArns []string `json:"inspector:ListExclusionsOutput:ExclusionArns" locationName:"exclusionArns" type:"list" required:"true"`

	// When a response is generated, if there is more data to be listed, this parameters
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to
	// be listed, this parameter is set to null.
	NextToken *string `json:"inspector:ListExclusionsOutput:NextToken" locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListExclusionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListExclusions = "ListExclusions"

// ListExclusionsRequest returns a request value for making API operation for
// Amazon Inspector.
//
// List exclusions that are generated by the assessment run.
//
//    // Example sending a request using ListExclusionsRequest.
//    req := client.ListExclusionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/ListExclusions
func (c *Client) ListExclusionsRequest(input *ListExclusionsInput) ListExclusionsRequest {
	op := &aws.Operation{
		Name:       opListExclusions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListExclusionsInput{}
	}

	req := c.newRequest(op, input, &ListExclusionsOutput{})
	return ListExclusionsRequest{Request: req, Input: input, Copy: c.ListExclusionsRequest}
}

// ListExclusionsRequest is the request type for the
// ListExclusions API operation.
type ListExclusionsRequest struct {
	*aws.Request
	Input *ListExclusionsInput
	Copy  func(*ListExclusionsInput) ListExclusionsRequest
}

// Send marshals and sends the ListExclusions API request.
func (r ListExclusionsRequest) Send(ctx context.Context) (*ListExclusionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListExclusionsResponse{
		ListExclusionsOutput: r.Request.Data.(*ListExclusionsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListExclusionsRequestPaginator returns a paginator for ListExclusions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListExclusionsRequest(input)
//   p := inspector.NewListExclusionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListExclusionsPaginator(req ListExclusionsRequest) ListExclusionsPaginator {
	return ListExclusionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListExclusionsInput
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

// ListExclusionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListExclusionsPaginator struct {
	aws.Pager
}

func (p *ListExclusionsPaginator) CurrentPage() *ListExclusionsOutput {
	return p.Pager.CurrentPage().(*ListExclusionsOutput)
}

// ListExclusionsResponse is the response type for the
// ListExclusions API operation.
type ListExclusionsResponse struct {
	*ListExclusionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListExclusions request.
func (r *ListExclusionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
