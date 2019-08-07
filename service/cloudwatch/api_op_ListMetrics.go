// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/ListMetricsInput
type ListMetricsInput struct {
	_ struct{} `type:"structure"`

	// The dimensions to filter against.
	Dimensions []DimensionFilter `type:"list"`

	// The name of the metric to filter against.
	MetricName *string `min:"1" type:"string"`

	// The namespace to filter against.
	Namespace *string `min:"1" type:"string"`

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListMetricsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMetricsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMetricsInput"}
	if s.MetricName != nil && len(*s.MetricName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricName", 1))
	}
	if s.Namespace != nil && len(*s.Namespace) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Namespace", 1))
	}
	if s.Dimensions != nil {
		for i, v := range s.Dimensions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Dimensions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/ListMetricsOutput
type ListMetricsOutput struct {
	_ struct{} `type:"structure"`

	// The metrics.
	Metrics []Metric `json:"monitoring:ListMetricsOutput:Metrics" type:"list"`

	// The token that marks the start of the next batch of returned results.
	NextToken *string `json:"monitoring:ListMetricsOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s ListMetricsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListMetrics = "ListMetrics"

// ListMetricsRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// List the specified metrics. You can use the returned metrics with GetMetricData
// or GetMetricStatistics to obtain statistical data.
//
// Up to 500 results are returned for any one call. To retrieve additional results,
// use the returned token with subsequent calls.
//
// After you create a metric, allow up to fifteen minutes before the metric
// appears. Statistics about the metric, however, are available sooner using
// GetMetricData or GetMetricStatistics.
//
//    // Example sending a request using ListMetricsRequest.
//    req := client.ListMetricsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/ListMetrics
func (c *Client) ListMetricsRequest(input *ListMetricsInput) ListMetricsRequest {
	op := &aws.Operation{
		Name:       opListMetrics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListMetricsInput{}
	}

	req := c.newRequest(op, input, &ListMetricsOutput{})
	return ListMetricsRequest{Request: req, Input: input, Copy: c.ListMetricsRequest}
}

// ListMetricsRequest is the request type for the
// ListMetrics API operation.
type ListMetricsRequest struct {
	*aws.Request
	Input *ListMetricsInput
	Copy  func(*ListMetricsInput) ListMetricsRequest
}

// Send marshals and sends the ListMetrics API request.
func (r ListMetricsRequest) Send(ctx context.Context) (*ListMetricsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMetricsResponse{
		ListMetricsOutput: r.Request.Data.(*ListMetricsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMetricsRequestPaginator returns a paginator for ListMetrics.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMetricsRequest(input)
//   p := cloudwatch.NewListMetricsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMetricsPaginator(req ListMetricsRequest) ListMetricsPaginator {
	return ListMetricsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMetricsInput
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

// ListMetricsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMetricsPaginator struct {
	aws.Pager
}

func (p *ListMetricsPaginator) CurrentPage() *ListMetricsOutput {
	return p.Pager.CurrentPage().(*ListMetricsOutput)
}

// ListMetricsResponse is the response type for the
// ListMetrics API operation.
type ListMetricsResponse struct {
	*ListMetricsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMetrics request.
func (r *ListMetricsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
