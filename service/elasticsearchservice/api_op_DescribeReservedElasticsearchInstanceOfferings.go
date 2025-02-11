// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Container for parameters to DescribeReservedElasticsearchInstanceOfferings
type DescribeReservedElasticsearchInstanceOfferingsInput struct {
	_ struct{} `type:"structure"`

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// NextToken should be sent in case if earlier API call produced result containing
	// NextToken. It is used for pagination.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The offering identifier filter value. Use this parameter to show only the
	// available offering that matches the specified reservation identifier.
	ReservedElasticsearchInstanceOfferingId *string `location:"querystring" locationName:"offeringId" type:"string"`
}

// String returns the string representation
func (s DescribeReservedElasticsearchInstanceOfferingsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeReservedElasticsearchInstanceOfferingsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ReservedElasticsearchInstanceOfferingId != nil {
		v := *s.ReservedElasticsearchInstanceOfferingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "offeringId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Container for results from DescribeReservedElasticsearchInstanceOfferings
type DescribeReservedElasticsearchInstanceOfferingsOutput struct {
	_ struct{} `type:"structure"`

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string `type:"string"`

	// List of reserved Elasticsearch instance offerings
	ReservedElasticsearchInstanceOfferings []ReservedElasticsearchInstanceOffering `type:"list"`
}

// String returns the string representation
func (s DescribeReservedElasticsearchInstanceOfferingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeReservedElasticsearchInstanceOfferingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ReservedElasticsearchInstanceOfferings != nil {
		v := s.ReservedElasticsearchInstanceOfferings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ReservedElasticsearchInstanceOfferings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribeReservedElasticsearchInstanceOfferings = "DescribeReservedElasticsearchInstanceOfferings"

// DescribeReservedElasticsearchInstanceOfferingsRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Lists available reserved Elasticsearch instance offerings.
//
//    // Example sending a request using DescribeReservedElasticsearchInstanceOfferingsRequest.
//    req := client.DescribeReservedElasticsearchInstanceOfferingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeReservedElasticsearchInstanceOfferingsRequest(input *DescribeReservedElasticsearchInstanceOfferingsInput) DescribeReservedElasticsearchInstanceOfferingsRequest {
	op := &aws.Operation{
		Name:       opDescribeReservedElasticsearchInstanceOfferings,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/es/reservedInstanceOfferings",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeReservedElasticsearchInstanceOfferingsInput{}
	}

	req := c.newRequest(op, input, &DescribeReservedElasticsearchInstanceOfferingsOutput{})
	return DescribeReservedElasticsearchInstanceOfferingsRequest{Request: req, Input: input, Copy: c.DescribeReservedElasticsearchInstanceOfferingsRequest}
}

// DescribeReservedElasticsearchInstanceOfferingsRequest is the request type for the
// DescribeReservedElasticsearchInstanceOfferings API operation.
type DescribeReservedElasticsearchInstanceOfferingsRequest struct {
	*aws.Request
	Input *DescribeReservedElasticsearchInstanceOfferingsInput
	Copy  func(*DescribeReservedElasticsearchInstanceOfferingsInput) DescribeReservedElasticsearchInstanceOfferingsRequest
}

// Send marshals and sends the DescribeReservedElasticsearchInstanceOfferings API request.
func (r DescribeReservedElasticsearchInstanceOfferingsRequest) Send(ctx context.Context) (*DescribeReservedElasticsearchInstanceOfferingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReservedElasticsearchInstanceOfferingsResponse{
		DescribeReservedElasticsearchInstanceOfferingsOutput: r.Request.Data.(*DescribeReservedElasticsearchInstanceOfferingsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReservedElasticsearchInstanceOfferingsRequestPaginator returns a paginator for DescribeReservedElasticsearchInstanceOfferings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReservedElasticsearchInstanceOfferingsRequest(input)
//   p := elasticsearchservice.NewDescribeReservedElasticsearchInstanceOfferingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReservedElasticsearchInstanceOfferingsPaginator(req DescribeReservedElasticsearchInstanceOfferingsRequest) DescribeReservedElasticsearchInstanceOfferingsPaginator {
	return DescribeReservedElasticsearchInstanceOfferingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeReservedElasticsearchInstanceOfferingsInput
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

// DescribeReservedElasticsearchInstanceOfferingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReservedElasticsearchInstanceOfferingsPaginator struct {
	aws.Pager
}

func (p *DescribeReservedElasticsearchInstanceOfferingsPaginator) CurrentPage() *DescribeReservedElasticsearchInstanceOfferingsOutput {
	return p.Pager.CurrentPage().(*DescribeReservedElasticsearchInstanceOfferingsOutput)
}

// DescribeReservedElasticsearchInstanceOfferingsResponse is the response type for the
// DescribeReservedElasticsearchInstanceOfferings API operation.
type DescribeReservedElasticsearchInstanceOfferingsResponse struct {
	*DescribeReservedElasticsearchInstanceOfferingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReservedElasticsearchInstanceOfferings request.
func (r *DescribeReservedElasticsearchInstanceOfferingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
