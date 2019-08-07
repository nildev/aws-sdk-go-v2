// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Retrieve deliverability data for all the campaigns that used a specific domain
// to send email during a specified time range. This data is available for a
// domain only if you enabled the Deliverability dashboard (PutDeliverabilityDashboardOption
// operation) for the domain.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/ListDomainDeliverabilityCampaignsRequest
type ListDomainDeliverabilityCampaignsInput struct {
	_ struct{} `type:"structure"`

	// The last day, in Unix time format, that you want to obtain deliverability
	// data for. This value has to be less than or equal to 30 days after the value
	// of the StartDate parameter.
	//
	// EndDate is a required field
	EndDate *time.Time `location:"querystring" locationName:"EndDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// A token that’s returned from a previous call to the ListDomainDeliverabilityCampaigns
	// operation. This token indicates the position of a campaign in the list of
	// campaigns.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	// The maximum number of results to include in response to a single call to
	// the ListDomainDeliverabilityCampaigns operation. If the number of results
	// is larger than the number that you specify in this parameter, the response
	// includes a NextToken element, which you can use to obtain additional results.
	PageSize *int64 `location:"querystring" locationName:"PageSize" type:"integer"`

	// The first day, in Unix time format, that you want to obtain deliverability
	// data for.
	//
	// StartDate is a required field
	StartDate *time.Time `location:"querystring" locationName:"StartDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The domain to obtain deliverability data for.
	//
	// SubscribedDomain is a required field
	SubscribedDomain *string `location:"uri" locationName:"SubscribedDomain" type:"string" required:"true"`
}

// String returns the string representation
func (s ListDomainDeliverabilityCampaignsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDomainDeliverabilityCampaignsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDomainDeliverabilityCampaignsInput"}

	if s.EndDate == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndDate"))
	}

	if s.StartDate == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartDate"))
	}

	if s.SubscribedDomain == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscribedDomain"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDomainDeliverabilityCampaignsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.SubscribedDomain != nil {
		v := *s.SubscribedDomain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "SubscribedDomain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndDate != nil {
		v := *s.EndDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "EndDate", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
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
	if s.StartDate != nil {
		v := *s.StartDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "StartDate", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	return nil
}

// An array of objects that provide deliverability data for all the campaigns
// that used a specific domain to send email during a specified time range.
// This data is available for a domain only if you enabled the Deliverability
// dashboard (PutDeliverabilityDashboardOption operation) for the domain.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/ListDomainDeliverabilityCampaignsResponse
type ListDomainDeliverabilityCampaignsOutput struct {
	_ struct{} `type:"structure"`

	// An array of responses, one for each campaign that used the domain to send
	// email during the specified time range.
	//
	// DomainDeliverabilityCampaigns is a required field
	DomainDeliverabilityCampaigns []DomainDeliverabilityCampaign `json:"email:ListDomainDeliverabilityCampaignsOutput:DomainDeliverabilityCampaigns" type:"list" required:"true"`

	// A token that’s returned from a previous call to the ListDomainDeliverabilityCampaigns
	// operation. This token indicates the position of the campaign in the list
	// of campaigns.
	NextToken *string `json:"email:ListDomainDeliverabilityCampaignsOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s ListDomainDeliverabilityCampaignsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDomainDeliverabilityCampaignsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainDeliverabilityCampaigns != nil {
		v := s.DomainDeliverabilityCampaigns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "DomainDeliverabilityCampaigns", metadata)
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

const opListDomainDeliverabilityCampaigns = "ListDomainDeliverabilityCampaigns"

// ListDomainDeliverabilityCampaignsRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Retrieve deliverability data for all the campaigns that used a specific domain
// to send email during a specified time range. This data is available for a
// domain only if you enabled the Deliverability dashboard (PutDeliverabilityDashboardOption
// operation) for the domain.
//
//    // Example sending a request using ListDomainDeliverabilityCampaignsRequest.
//    req := client.ListDomainDeliverabilityCampaignsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/ListDomainDeliverabilityCampaigns
func (c *Client) ListDomainDeliverabilityCampaignsRequest(input *ListDomainDeliverabilityCampaignsInput) ListDomainDeliverabilityCampaignsRequest {
	op := &aws.Operation{
		Name:       opListDomainDeliverabilityCampaigns,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/deliverability-dashboard/domains/{SubscribedDomain}/campaigns",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDomainDeliverabilityCampaignsInput{}
	}

	req := c.newRequest(op, input, &ListDomainDeliverabilityCampaignsOutput{})
	return ListDomainDeliverabilityCampaignsRequest{Request: req, Input: input, Copy: c.ListDomainDeliverabilityCampaignsRequest}
}

// ListDomainDeliverabilityCampaignsRequest is the request type for the
// ListDomainDeliverabilityCampaigns API operation.
type ListDomainDeliverabilityCampaignsRequest struct {
	*aws.Request
	Input *ListDomainDeliverabilityCampaignsInput
	Copy  func(*ListDomainDeliverabilityCampaignsInput) ListDomainDeliverabilityCampaignsRequest
}

// Send marshals and sends the ListDomainDeliverabilityCampaigns API request.
func (r ListDomainDeliverabilityCampaignsRequest) Send(ctx context.Context) (*ListDomainDeliverabilityCampaignsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDomainDeliverabilityCampaignsResponse{
		ListDomainDeliverabilityCampaignsOutput: r.Request.Data.(*ListDomainDeliverabilityCampaignsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDomainDeliverabilityCampaignsRequestPaginator returns a paginator for ListDomainDeliverabilityCampaigns.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDomainDeliverabilityCampaignsRequest(input)
//   p := pinpointemail.NewListDomainDeliverabilityCampaignsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDomainDeliverabilityCampaignsPaginator(req ListDomainDeliverabilityCampaignsRequest) ListDomainDeliverabilityCampaignsPaginator {
	return ListDomainDeliverabilityCampaignsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDomainDeliverabilityCampaignsInput
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

// ListDomainDeliverabilityCampaignsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDomainDeliverabilityCampaignsPaginator struct {
	aws.Pager
}

func (p *ListDomainDeliverabilityCampaignsPaginator) CurrentPage() *ListDomainDeliverabilityCampaignsOutput {
	return p.Pager.CurrentPage().(*ListDomainDeliverabilityCampaignsOutput)
}

// ListDomainDeliverabilityCampaignsResponse is the response type for the
// ListDomainDeliverabilityCampaigns API operation.
type ListDomainDeliverabilityCampaignsResponse struct {
	*ListDomainDeliverabilityCampaignsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDomainDeliverabilityCampaigns request.
func (r *ListDomainDeliverabilityCampaignsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
