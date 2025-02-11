// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to list all of the predictive inbox placement tests that you've
// performed.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/ListDeliverabilityTestReportsRequest
type ListDeliverabilityTestReportsInput struct {
	_ struct{} `type:"structure"`

	// A token returned from a previous call to ListDeliverabilityTestReports to
	// indicate the position in the list of predictive inbox placement tests.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	// The number of results to show in a single call to ListDeliverabilityTestReports.
	// If the number of results is larger than the number you specified in this
	// parameter, then the response includes a NextToken element, which you can
	// use to obtain additional results.
	//
	// The value you specify has to be at least 0, and can be no more than 1000.
	PageSize *int64 `location:"querystring" locationName:"PageSize" type:"integer"`
}

// String returns the string representation
func (s ListDeliverabilityTestReportsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeliverabilityTestReportsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "PageSize", protocol.Int64Value(v), metadata)
	}
	return nil
}

// A list of the predictive inbox placement test reports that are available
// for your account, regardless of whether or not those tests are complete.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/ListDeliverabilityTestReportsResponse
type ListDeliverabilityTestReportsOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains a lists of predictive inbox placement tests that
	// you've performed.
	//
	// DeliverabilityTestReports is a required field
	DeliverabilityTestReports []DeliverabilityTestReport `type:"list" required:"true"`

	// A token that indicates that there are additional predictive inbox placement
	// tests to list. To view additional predictive inbox placement tests, issue
	// another request to ListDeliverabilityTestReports, and pass this token in
	// the NextToken parameter.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDeliverabilityTestReportsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeliverabilityTestReportsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeliverabilityTestReports != nil {
		v := s.DeliverabilityTestReports

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "DeliverabilityTestReports", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDeliverabilityTestReports = "ListDeliverabilityTestReports"

// ListDeliverabilityTestReportsRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Show a list of the predictive inbox placement tests that you've performed,
// regardless of their statuses. For predictive inbox placement tests that are
// complete, you can use the GetDeliverabilityTestReport operation to view the
// results.
//
//    // Example sending a request using ListDeliverabilityTestReportsRequest.
//    req := client.ListDeliverabilityTestReportsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/ListDeliverabilityTestReports
func (c *Client) ListDeliverabilityTestReportsRequest(input *ListDeliverabilityTestReportsInput) ListDeliverabilityTestReportsRequest {
	op := &aws.Operation{
		Name:       opListDeliverabilityTestReports,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/deliverability-dashboard/test-reports",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDeliverabilityTestReportsInput{}
	}

	req := c.newRequest(op, input, &ListDeliverabilityTestReportsOutput{})
	return ListDeliverabilityTestReportsRequest{Request: req, Input: input, Copy: c.ListDeliverabilityTestReportsRequest}
}

// ListDeliverabilityTestReportsRequest is the request type for the
// ListDeliverabilityTestReports API operation.
type ListDeliverabilityTestReportsRequest struct {
	*aws.Request
	Input *ListDeliverabilityTestReportsInput
	Copy  func(*ListDeliverabilityTestReportsInput) ListDeliverabilityTestReportsRequest
}

// Send marshals and sends the ListDeliverabilityTestReports API request.
func (r ListDeliverabilityTestReportsRequest) Send(ctx context.Context) (*ListDeliverabilityTestReportsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeliverabilityTestReportsResponse{
		ListDeliverabilityTestReportsOutput: r.Request.Data.(*ListDeliverabilityTestReportsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDeliverabilityTestReportsRequestPaginator returns a paginator for ListDeliverabilityTestReports.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDeliverabilityTestReportsRequest(input)
//   p := pinpointemail.NewListDeliverabilityTestReportsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDeliverabilityTestReportsPaginator(req ListDeliverabilityTestReportsRequest) ListDeliverabilityTestReportsPaginator {
	return ListDeliverabilityTestReportsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDeliverabilityTestReportsInput
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

// ListDeliverabilityTestReportsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDeliverabilityTestReportsPaginator struct {
	aws.Pager
}

func (p *ListDeliverabilityTestReportsPaginator) CurrentPage() *ListDeliverabilityTestReportsOutput {
	return p.Pager.CurrentPage().(*ListDeliverabilityTestReportsOutput)
}

// ListDeliverabilityTestReportsResponse is the response type for the
// ListDeliverabilityTestReports API operation.
type ListDeliverabilityTestReportsResponse struct {
	*ListDeliverabilityTestReportsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeliverabilityTestReports request.
func (r *ListDeliverabilityTestReportsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
