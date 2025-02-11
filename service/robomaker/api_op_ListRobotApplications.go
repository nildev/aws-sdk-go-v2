// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListRobotApplicationsRequest
type ListRobotApplicationsInput struct {
	_ struct{} `type:"structure"`

	// Optional filters to limit results.
	//
	// The filter name name is supported. When filtering, you must use the complete
	// value of the filtered item. You can use up to three filters.
	Filters []Filter `locationName:"filters" min:"1" type:"list"`

	// The maximum number of deployment job results returned by ListRobotApplications
	// in paginated output. When this parameter is used, ListRobotApplications only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListRobotApplications request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// ListRobotApplications returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListRobotApplications
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The version qualifier of the robot application.
	VersionQualifier *string `locationName:"versionQualifier" type:"string"`
}

// String returns the string representation
func (s ListRobotApplicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRobotApplicationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRobotApplicationsInput"}
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
func (s ListRobotApplicationsInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.VersionQualifier != nil {
		v := *s.VersionQualifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "versionQualifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListRobotApplicationsResponse
type ListRobotApplicationsOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListRobotApplications request.
	// When the results of a ListRobotApplications request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// A list of robot application summaries that meet the criteria of the request.
	RobotApplicationSummaries []RobotApplicationSummary `locationName:"robotApplicationSummaries" type:"list"`
}

// String returns the string representation
func (s ListRobotApplicationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListRobotApplicationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RobotApplicationSummaries != nil {
		v := s.RobotApplicationSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "robotApplicationSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListRobotApplications = "ListRobotApplications"

// ListRobotApplicationsRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Returns a list of robot application. You can optionally provide filters to
// retrieve specific robot applications.
//
//    // Example sending a request using ListRobotApplicationsRequest.
//    req := client.ListRobotApplicationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListRobotApplications
func (c *Client) ListRobotApplicationsRequest(input *ListRobotApplicationsInput) ListRobotApplicationsRequest {
	op := &aws.Operation{
		Name:       opListRobotApplications,
		HTTPMethod: "POST",
		HTTPPath:   "/listRobotApplications",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListRobotApplicationsInput{}
	}

	req := c.newRequest(op, input, &ListRobotApplicationsOutput{})
	return ListRobotApplicationsRequest{Request: req, Input: input, Copy: c.ListRobotApplicationsRequest}
}

// ListRobotApplicationsRequest is the request type for the
// ListRobotApplications API operation.
type ListRobotApplicationsRequest struct {
	*aws.Request
	Input *ListRobotApplicationsInput
	Copy  func(*ListRobotApplicationsInput) ListRobotApplicationsRequest
}

// Send marshals and sends the ListRobotApplications API request.
func (r ListRobotApplicationsRequest) Send(ctx context.Context) (*ListRobotApplicationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRobotApplicationsResponse{
		ListRobotApplicationsOutput: r.Request.Data.(*ListRobotApplicationsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRobotApplicationsRequestPaginator returns a paginator for ListRobotApplications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRobotApplicationsRequest(input)
//   p := robomaker.NewListRobotApplicationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRobotApplicationsPaginator(req ListRobotApplicationsRequest) ListRobotApplicationsPaginator {
	return ListRobotApplicationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListRobotApplicationsInput
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

// ListRobotApplicationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRobotApplicationsPaginator struct {
	aws.Pager
}

func (p *ListRobotApplicationsPaginator) CurrentPage() *ListRobotApplicationsOutput {
	return p.Pager.CurrentPage().(*ListRobotApplicationsOutput)
}

// ListRobotApplicationsResponse is the response type for the
// ListRobotApplications API operation.
type ListRobotApplicationsResponse struct {
	*ListRobotApplicationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRobotApplications request.
func (r *ListRobotApplicationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
