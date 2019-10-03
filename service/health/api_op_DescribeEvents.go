// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package health

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/health-2016-08-04/DescribeEventsRequest
type DescribeEventsInput struct {
	_ struct{} `type:"structure"`

	// Values to narrow the results returned.
	Filter *EventFilter `locationName:"filter" type:"structure"`

	// The locale (language) to return information in. English (en) is the default
	// and the only supported value at this time.
	Locale *string `locationName:"locale" min:"2" type:"string"`

	// The maximum number of items to return in one batch, between 10 and 100, inclusive.
	MaxResults *int64 `locationName:"maxResults" min:"10" type:"integer"`

	// If the results of a search are large, only a portion of the results are returned,
	// and a nextToken pagination token is returned in the response. To retrieve
	// the next batch of results, reissue the search request and include the returned
	// token. When all results have been returned, the response does not contain
	// a pagination token value.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventsInput"}
	if s.Locale != nil && len(*s.Locale) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Locale", 2))
	}
	if s.MaxResults != nil && *s.MaxResults < 10 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 10))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/health-2016-08-04/DescribeEventsResponse
type DescribeEventsOutput struct {
	_ struct{} `type:"structure"`

	// The events that match the specified filter criteria.
	Events []Event `json:"health:DescribeEventsOutput:Events" locationName:"events" type:"list"`

	// If the results of a search are large, only a portion of the results are returned,
	// and a nextToken pagination token is returned in the response. To retrieve
	// the next batch of results, reissue the search request and include the returned
	// token. When all results have been returned, the response does not contain
	// a pagination token value.
	NextToken *string `json:"health:DescribeEventsOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeEventsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEvents = "DescribeEvents"

// DescribeEventsRequest returns a request value for making API operation for
// AWS Health APIs and Notifications.
//
// Returns information about events that meet the specified filter criteria.
// Events are returned in a summary form and do not include the detailed description,
// any additional metadata that depends on the event type, or any affected resources.
// To retrieve that information, use the DescribeEventDetails and DescribeAffectedEntities
// operations.
//
// If no filter criteria are specified, all events are returned. Results are
// sorted by lastModifiedTime, starting with the most recent.
//
//    // Example sending a request using DescribeEventsRequest.
//    req := client.DescribeEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/health-2016-08-04/DescribeEvents
func (c *Client) DescribeEventsRequest(input *DescribeEventsInput) DescribeEventsRequest {
	op := &aws.Operation{
		Name:       opDescribeEvents,
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
		input = &DescribeEventsInput{}
	}

	req := c.newRequest(op, input, &DescribeEventsOutput{})
	return DescribeEventsRequest{Request: req, Input: input, Copy: c.DescribeEventsRequest}
}

// DescribeEventsRequest is the request type for the
// DescribeEvents API operation.
type DescribeEventsRequest struct {
	*aws.Request
	Input *DescribeEventsInput
	Copy  func(*DescribeEventsInput) DescribeEventsRequest
}

// Send marshals and sends the DescribeEvents API request.
func (r DescribeEventsRequest) Send(ctx context.Context) (*DescribeEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventsResponse{
		DescribeEventsOutput: r.Request.Data.(*DescribeEventsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEventsRequestPaginator returns a paginator for DescribeEvents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEventsRequest(input)
//   p := health.NewDescribeEventsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEventsPaginator(req DescribeEventsRequest) DescribeEventsPaginator {
	return DescribeEventsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEventsInput
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

// DescribeEventsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEventsPaginator struct {
	aws.Pager
}

func (p *DescribeEventsPaginator) CurrentPage() *DescribeEventsOutput {
	return p.Pager.CurrentPage().(*DescribeEventsOutput)
}

// DescribeEventsResponse is the response type for the
// DescribeEvents API operation.
type DescribeEventsResponse struct {
	*DescribeEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEvents request.
func (r *DescribeEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
