// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeSpotPriceHistory.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeSpotPriceHistoryRequest
type DescribeSpotPriceHistoryInput struct {
	_ struct{} `type:"structure"`

	// Filters the results by the specified Availability Zone.
	AvailabilityZone *string `locationName:"availabilityZone" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The date and time, up to the current date, from which to stop retrieving
	// the price history data, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime *time.Time `locationName:"endTime" type:"timestamp" timestampFormat:"iso8601"`

	// One or more filters.
	//
	//    * availability-zone - The Availability Zone for which prices should be
	//    returned.
	//
	//    * instance-type - The type of instance (for example, m3.medium).
	//
	//    * product-description - The product description for the Spot price (Linux/UNIX
	//    | SUSE Linux | Windows | Linux/UNIX (Amazon VPC) | SUSE Linux (Amazon
	//    VPC) | Windows (Amazon VPC)).
	//
	//    * spot-price - The Spot price. The value must match exactly (or use wildcards;
	//    greater than or less than comparison is not supported).
	//
	//    * timestamp - The time stamp of the Spot price history, in UTC format
	//    (for example, YYYY-MM-DDTHH:MM:SSZ). You can use wildcards (* and ?).
	//    Greater than or less than comparison is not supported.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// Filters the results by the specified instance types.
	InstanceTypes []InstanceType `locationName:"InstanceType" type:"list"`

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Filters the results by the specified basic product descriptions.
	ProductDescriptions []string `locationName:"ProductDescription" type:"list"`

	// The date and time, up to the past 90 days, from which to start retrieving
	// the price history data, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time `locationName:"startTime" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s DescribeSpotPriceHistoryInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output of DescribeSpotPriceHistory.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeSpotPriceHistoryResult
type DescribeSpotPriceHistoryOutput struct {
	_ struct{} `type:"structure"`

	// The token required to retrieve the next set of results. This value is null
	// or an empty string when there are no more results to return.
	NextToken *string `json:"ec2:DescribeSpotPriceHistoryOutput:NextToken" locationName:"nextToken" type:"string"`

	// The historical Spot prices.
	SpotPriceHistory []SpotPrice `json:"ec2:DescribeSpotPriceHistoryOutput:SpotPriceHistory" locationName:"spotPriceHistorySet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeSpotPriceHistoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSpotPriceHistory = "DescribeSpotPriceHistory"

// DescribeSpotPriceHistoryRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the Spot price history. For more information, see Spot Instance
// Pricing History (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances-history.html)
// in the Amazon EC2 User Guide for Linux Instances.
//
// When you specify a start and end time, this operation returns the prices
// of the instance types within the time range that you specified and the time
// when the price changed. The price is valid within the time period that you
// specified; the response merely indicates the last time that the price changed.
//
//    // Example sending a request using DescribeSpotPriceHistoryRequest.
//    req := client.DescribeSpotPriceHistoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeSpotPriceHistory
func (c *Client) DescribeSpotPriceHistoryRequest(input *DescribeSpotPriceHistoryInput) DescribeSpotPriceHistoryRequest {
	op := &aws.Operation{
		Name:       opDescribeSpotPriceHistory,
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
		input = &DescribeSpotPriceHistoryInput{}
	}

	req := c.newRequest(op, input, &DescribeSpotPriceHistoryOutput{})
	return DescribeSpotPriceHistoryRequest{Request: req, Input: input, Copy: c.DescribeSpotPriceHistoryRequest}
}

// DescribeSpotPriceHistoryRequest is the request type for the
// DescribeSpotPriceHistory API operation.
type DescribeSpotPriceHistoryRequest struct {
	*aws.Request
	Input *DescribeSpotPriceHistoryInput
	Copy  func(*DescribeSpotPriceHistoryInput) DescribeSpotPriceHistoryRequest
}

// Send marshals and sends the DescribeSpotPriceHistory API request.
func (r DescribeSpotPriceHistoryRequest) Send(ctx context.Context) (*DescribeSpotPriceHistoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSpotPriceHistoryResponse{
		DescribeSpotPriceHistoryOutput: r.Request.Data.(*DescribeSpotPriceHistoryOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeSpotPriceHistoryRequestPaginator returns a paginator for DescribeSpotPriceHistory.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeSpotPriceHistoryRequest(input)
//   p := ec2.NewDescribeSpotPriceHistoryRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeSpotPriceHistoryPaginator(req DescribeSpotPriceHistoryRequest) DescribeSpotPriceHistoryPaginator {
	return DescribeSpotPriceHistoryPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeSpotPriceHistoryInput
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

// DescribeSpotPriceHistoryPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeSpotPriceHistoryPaginator struct {
	aws.Pager
}

func (p *DescribeSpotPriceHistoryPaginator) CurrentPage() *DescribeSpotPriceHistoryOutput {
	return p.Pager.CurrentPage().(*DescribeSpotPriceHistoryOutput)
}

// DescribeSpotPriceHistoryResponse is the response type for the
// DescribeSpotPriceHistory API operation.
type DescribeSpotPriceHistoryResponse struct {
	*DescribeSpotPriceHistoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSpotPriceHistory request.
func (r *DescribeSpotPriceHistoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
