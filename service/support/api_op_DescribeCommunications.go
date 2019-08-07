// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeCommunicationsRequest
type DescribeCommunicationsInput struct {
	_ struct{} `type:"structure"`

	// The start date for a filtered date search on support case communications.
	// Case communications are available for 12 months after creation.
	AfterTime *string `locationName:"afterTime" type:"string"`

	// The end date for a filtered date search on support case communications. Case
	// communications are available for 12 months after creation.
	BeforeTime *string `locationName:"beforeTime" type:"string"`

	// The AWS Support case ID requested or returned in the call. The case ID is
	// an alphanumeric string formatted as shown in this example: case-12345678910-2013-c4c1d2bf33c5cf47
	//
	// CaseId is a required field
	CaseId *string `locationName:"caseId" type:"string" required:"true"`

	// The maximum number of results to return before paginating.
	MaxResults *int64 `locationName:"maxResults" min:"10" type:"integer"`

	// A resumption point for pagination.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeCommunicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCommunicationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCommunicationsInput"}

	if s.CaseId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CaseId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 10 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 10))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The communications returned by the DescribeCommunications operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeCommunicationsResponse
type DescribeCommunicationsOutput struct {
	_ struct{} `type:"structure"`

	// The communications for the case.
	Communications []Communication `json:"support:DescribeCommunicationsOutput:Communications" locationName:"communications" type:"list"`

	// A resumption point for pagination.
	NextToken *string `json:"support:DescribeCommunicationsOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeCommunicationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCommunications = "DescribeCommunications"

// DescribeCommunicationsRequest returns a request value for making API operation for
// AWS Support.
//
// Returns communications (and attachments) for one or more support cases. You
// can use the afterTime and beforeTime parameters to filter by date. You can
// use the caseId parameter to restrict the results to a particular case.
//
// Case data is available for 12 months after creation. If a case was created
// more than 12 months ago, a request for data might cause an error.
//
// You can use the maxResults and nextToken parameters to control the pagination
// of the result set. Set maxResults to the number of cases you want displayed
// on each page, and use nextToken to specify the resumption of pagination.
//
//    // Example sending a request using DescribeCommunicationsRequest.
//    req := client.DescribeCommunicationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeCommunications
func (c *Client) DescribeCommunicationsRequest(input *DescribeCommunicationsInput) DescribeCommunicationsRequest {
	op := &aws.Operation{
		Name:       opDescribeCommunications,
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
		input = &DescribeCommunicationsInput{}
	}

	req := c.newRequest(op, input, &DescribeCommunicationsOutput{})
	return DescribeCommunicationsRequest{Request: req, Input: input, Copy: c.DescribeCommunicationsRequest}
}

// DescribeCommunicationsRequest is the request type for the
// DescribeCommunications API operation.
type DescribeCommunicationsRequest struct {
	*aws.Request
	Input *DescribeCommunicationsInput
	Copy  func(*DescribeCommunicationsInput) DescribeCommunicationsRequest
}

// Send marshals and sends the DescribeCommunications API request.
func (r DescribeCommunicationsRequest) Send(ctx context.Context) (*DescribeCommunicationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCommunicationsResponse{
		DescribeCommunicationsOutput: r.Request.Data.(*DescribeCommunicationsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCommunicationsRequestPaginator returns a paginator for DescribeCommunications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCommunicationsRequest(input)
//   p := support.NewDescribeCommunicationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCommunicationsPaginator(req DescribeCommunicationsRequest) DescribeCommunicationsPaginator {
	return DescribeCommunicationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeCommunicationsInput
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

// DescribeCommunicationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCommunicationsPaginator struct {
	aws.Pager
}

func (p *DescribeCommunicationsPaginator) CurrentPage() *DescribeCommunicationsOutput {
	return p.Pager.CurrentPage().(*DescribeCommunicationsOutput)
}

// DescribeCommunicationsResponse is the response type for the
// DescribeCommunications API operation.
type DescribeCommunicationsResponse struct {
	*DescribeCommunicationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCommunications request.
func (r *DescribeCommunicationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
