// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ListDeviceEventsRequest
type ListDeviceEventsInput struct {
	_ struct{} `type:"structure"`

	// The ARN of a device.
	//
	// DeviceArn is a required field
	DeviceArn *string `type:"string" required:"true"`

	// The event type to filter device events. If EventType isn't specified, this
	// returns a list of all device events in reverse chronological order. If EventType
	// is specified, this returns a list of device events for that EventType in
	// reverse chronological order.
	EventType DeviceEventType `type:"string" enum:"true"`

	// The maximum number of results to include in the response. The default value
	// is 50. If more results exist than the specified MaxResults value, a token
	// is included in the response so that the remaining results can be retrieved.
	MaxResults *int64 `min:"1" type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// only includes results beyond the token, up to the value specified by MaxResults.
	// When the end of results is reached, the response has a value of null.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDeviceEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeviceEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDeviceEventsInput"}

	if s.DeviceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ListDeviceEventsResponse
type ListDeviceEventsOutput struct {
	_ struct{} `type:"structure"`

	// The device events requested for the device ARN.
	DeviceEvents []DeviceEvent `json:"a4b:ListDeviceEventsOutput:DeviceEvents" type:"list"`

	// The token returned to indicate that there is more data available.
	NextToken *string `json:"a4b:ListDeviceEventsOutput:NextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListDeviceEventsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDeviceEvents = "ListDeviceEvents"

// ListDeviceEventsRequest returns a request value for making API operation for
// Alexa For Business.
//
// Lists the device event history, including device connection status, for up
// to 30 days.
//
//    // Example sending a request using ListDeviceEventsRequest.
//    req := client.ListDeviceEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ListDeviceEvents
func (c *Client) ListDeviceEventsRequest(input *ListDeviceEventsInput) ListDeviceEventsRequest {
	op := &aws.Operation{
		Name:       opListDeviceEvents,
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
		input = &ListDeviceEventsInput{}
	}

	req := c.newRequest(op, input, &ListDeviceEventsOutput{})
	return ListDeviceEventsRequest{Request: req, Input: input, Copy: c.ListDeviceEventsRequest}
}

// ListDeviceEventsRequest is the request type for the
// ListDeviceEvents API operation.
type ListDeviceEventsRequest struct {
	*aws.Request
	Input *ListDeviceEventsInput
	Copy  func(*ListDeviceEventsInput) ListDeviceEventsRequest
}

// Send marshals and sends the ListDeviceEvents API request.
func (r ListDeviceEventsRequest) Send(ctx context.Context) (*ListDeviceEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeviceEventsResponse{
		ListDeviceEventsOutput: r.Request.Data.(*ListDeviceEventsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDeviceEventsRequestPaginator returns a paginator for ListDeviceEvents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDeviceEventsRequest(input)
//   p := alexaforbusiness.NewListDeviceEventsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDeviceEventsPaginator(req ListDeviceEventsRequest) ListDeviceEventsPaginator {
	return ListDeviceEventsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDeviceEventsInput
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

// ListDeviceEventsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDeviceEventsPaginator struct {
	aws.Pager
}

func (p *ListDeviceEventsPaginator) CurrentPage() *ListDeviceEventsOutput {
	return p.Pager.CurrentPage().(*ListDeviceEventsOutput)
}

// ListDeviceEventsResponse is the response type for the
// ListDeviceEvents API operation.
type ListDeviceEventsResponse struct {
	*ListDeviceEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeviceEvents request.
func (r *ListDeviceEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
