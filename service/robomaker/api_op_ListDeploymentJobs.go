// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListDeploymentJobsRequest
type ListDeploymentJobsInput struct {
	_ struct{} `type:"structure"`

	// Optional filters to limit results.
	//
	// The filter names status and fleetName are supported. When filtering, you
	// must use the complete value of the filtered item. You can use up to three
	// filters, but they must be for the same named item. For example, if you are
	// looking for items with the status InProgress or the status Pending.
	Filters []Filter `locationName:"filters" min:"1" type:"list"`

	// The maximum number of deployment job results returned by ListDeploymentJobs
	// in paginated output. When this parameter is used, ListDeploymentJobs only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListDeploymentJobs request with the returned nextToken value. This
	// value can be between 1 and 100. If this parameter is not used, then ListDeploymentJobs
	// returns up to 100 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListDeploymentJobs
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListDeploymentJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeploymentJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDeploymentJobsInput"}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeploymentJobsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListDeploymentJobsResponse
type ListDeploymentJobsOutput struct {
	_ struct{} `type:"structure"`

	// A list of deployment jobs that meet the criteria of the request.
	DeploymentJobs []DeploymentJob `locationName:"deploymentJobs" type:"list"`

	// The nextToken value to include in a future ListDeploymentJobs request. When
	// the results of a ListDeploymentJobs request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListDeploymentJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDeploymentJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeploymentJobs != nil {
		v := s.DeploymentJobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "deploymentJobs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDeploymentJobs = "ListDeploymentJobs"

// ListDeploymentJobsRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Returns a list of deployment jobs for a fleet. You can optionally provide
// filters to retrieve specific deployment jobs.
//
//    // Example sending a request using ListDeploymentJobsRequest.
//    req := client.ListDeploymentJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListDeploymentJobs
func (c *Client) ListDeploymentJobsRequest(input *ListDeploymentJobsInput) ListDeploymentJobsRequest {
	op := &aws.Operation{
		Name:       opListDeploymentJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/listDeploymentJobs",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDeploymentJobsInput{}
	}

	req := c.newRequest(op, input, &ListDeploymentJobsOutput{})
	return ListDeploymentJobsRequest{Request: req, Input: input, Copy: c.ListDeploymentJobsRequest}
}

// ListDeploymentJobsRequest is the request type for the
// ListDeploymentJobs API operation.
type ListDeploymentJobsRequest struct {
	*aws.Request
	Input *ListDeploymentJobsInput
	Copy  func(*ListDeploymentJobsInput) ListDeploymentJobsRequest
}

// Send marshals and sends the ListDeploymentJobs API request.
func (r ListDeploymentJobsRequest) Send(ctx context.Context) (*ListDeploymentJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeploymentJobsResponse{
		ListDeploymentJobsOutput: r.Request.Data.(*ListDeploymentJobsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDeploymentJobsRequestPaginator returns a paginator for ListDeploymentJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDeploymentJobsRequest(input)
//   p := robomaker.NewListDeploymentJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDeploymentJobsPaginator(req ListDeploymentJobsRequest) ListDeploymentJobsPaginator {
	return ListDeploymentJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDeploymentJobsInput
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

// ListDeploymentJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDeploymentJobsPaginator struct {
	aws.Pager
}

func (p *ListDeploymentJobsPaginator) CurrentPage() *ListDeploymentJobsOutput {
	return p.Pager.CurrentPage().(*ListDeploymentJobsOutput)
}

// ListDeploymentJobsResponse is the response type for the
// ListDeploymentJobs API operation.
type ListDeploymentJobsResponse struct {
	*ListDeploymentJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeploymentJobs request.
func (r *ListDeploymentJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
