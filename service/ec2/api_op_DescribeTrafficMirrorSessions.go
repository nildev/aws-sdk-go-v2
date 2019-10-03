// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTrafficMirrorSessionsRequest
type DescribeTrafficMirrorSessionsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * description: The Traffic Mirror session description.
	//
	//    * network-interface-id: The ID of the Traffic Mirror session network interface.
	//
	//    * owner-id: The ID of the account that owns the Traffic Mirror session.
	//
	//    * packet-length: The assigned number of packets to mirror.
	//
	//    * session-number: The assigned session number.
	//
	//    * traffic-mirror-filter-id: The ID of the Traffic Mirror filter.
	//
	//    * traffic-mirror-session-id: The ID of the Traffic Mirror session.
	//
	//    * traffic-mirror-target-id: The ID of the Traffic Mirror target.
	//
	//    * virtual-network-id: The virtual network ID of the Traffic Mirror session.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The ID of the Traffic Mirror session.
	TrafficMirrorSessionIds []string `locationName:"TrafficMirrorSessionId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeTrafficMirrorSessionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTrafficMirrorSessionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTrafficMirrorSessionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTrafficMirrorSessionsResult
type DescribeTrafficMirrorSessionsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. The value is null
	// when there are no more results to return.
	NextToken *string `json:"ec2:DescribeTrafficMirrorSessionsOutput:NextToken" locationName:"nextToken" type:"string"`

	// Describes one or more Traffic Mirror sessions. By default, all Traffic Mirror
	// sessions are described. Alternatively, you can filter the results.
	TrafficMirrorSessions []TrafficMirrorSession `json:"ec2:DescribeTrafficMirrorSessionsOutput:TrafficMirrorSessions" locationName:"trafficMirrorSessionSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeTrafficMirrorSessionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTrafficMirrorSessions = "DescribeTrafficMirrorSessions"

// DescribeTrafficMirrorSessionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more Traffic Mirror sessions. By default, all Traffic Mirror
// sessions are described. Alternatively, you can filter the results.
//
//    // Example sending a request using DescribeTrafficMirrorSessionsRequest.
//    req := client.DescribeTrafficMirrorSessionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTrafficMirrorSessions
func (c *Client) DescribeTrafficMirrorSessionsRequest(input *DescribeTrafficMirrorSessionsInput) DescribeTrafficMirrorSessionsRequest {
	op := &aws.Operation{
		Name:       opDescribeTrafficMirrorSessions,
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
		input = &DescribeTrafficMirrorSessionsInput{}
	}

	req := c.newRequest(op, input, &DescribeTrafficMirrorSessionsOutput{})
	return DescribeTrafficMirrorSessionsRequest{Request: req, Input: input, Copy: c.DescribeTrafficMirrorSessionsRequest}
}

// DescribeTrafficMirrorSessionsRequest is the request type for the
// DescribeTrafficMirrorSessions API operation.
type DescribeTrafficMirrorSessionsRequest struct {
	*aws.Request
	Input *DescribeTrafficMirrorSessionsInput
	Copy  func(*DescribeTrafficMirrorSessionsInput) DescribeTrafficMirrorSessionsRequest
}

// Send marshals and sends the DescribeTrafficMirrorSessions API request.
func (r DescribeTrafficMirrorSessionsRequest) Send(ctx context.Context) (*DescribeTrafficMirrorSessionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTrafficMirrorSessionsResponse{
		DescribeTrafficMirrorSessionsOutput: r.Request.Data.(*DescribeTrafficMirrorSessionsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTrafficMirrorSessionsRequestPaginator returns a paginator for DescribeTrafficMirrorSessions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTrafficMirrorSessionsRequest(input)
//   p := ec2.NewDescribeTrafficMirrorSessionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTrafficMirrorSessionsPaginator(req DescribeTrafficMirrorSessionsRequest) DescribeTrafficMirrorSessionsPaginator {
	return DescribeTrafficMirrorSessionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeTrafficMirrorSessionsInput
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

// DescribeTrafficMirrorSessionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTrafficMirrorSessionsPaginator struct {
	aws.Pager
}

func (p *DescribeTrafficMirrorSessionsPaginator) CurrentPage() *DescribeTrafficMirrorSessionsOutput {
	return p.Pager.CurrentPage().(*DescribeTrafficMirrorSessionsOutput)
}

// DescribeTrafficMirrorSessionsResponse is the response type for the
// DescribeTrafficMirrorSessions API operation.
type DescribeTrafficMirrorSessionsResponse struct {
	*DescribeTrafficMirrorSessionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTrafficMirrorSessions request.
func (r *DescribeTrafficMirrorSessionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
