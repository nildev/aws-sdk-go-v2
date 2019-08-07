// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTrafficMirrorTargetsRequest
type DescribeTrafficMirrorTargetsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * description: The Traffic Mirror target description.
	//
	//    * network-interface-id: The ID of the Traffic Mirror session network interface.
	//
	//    * network-load-balancer-arn: The Amazon Resource Name (ARN) of the Network
	//    Load Balancer that is associated with the session.
	//
	//    * owner-id: The ID of the account that owns the Traffic Mirror session.
	//
	//    * traffic-mirror-target-id: The ID of the Traffic Mirror target.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The ID of the Traffic Mirror targets.
	TrafficMirrorTargetIds []string `locationName:"TrafficMirrorTargetId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeTrafficMirrorTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTrafficMirrorTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTrafficMirrorTargetsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTrafficMirrorTargetsResult
type DescribeTrafficMirrorTargetsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. The value is null
	// when there are no more results to return.
	NextToken *string `json:"ec2:DescribeTrafficMirrorTargetsOutput:NextToken" locationName:"nextToken" type:"string"`

	// Information about one or more Traffic Mirror targets.
	TrafficMirrorTargets []TrafficMirrorTarget `json:"ec2:DescribeTrafficMirrorTargetsOutput:TrafficMirrorTargets" locationName:"trafficMirrorTargetSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeTrafficMirrorTargetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTrafficMirrorTargets = "DescribeTrafficMirrorTargets"

// DescribeTrafficMirrorTargetsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Information about one or more Traffic Mirror targets.
//
//    // Example sending a request using DescribeTrafficMirrorTargetsRequest.
//    req := client.DescribeTrafficMirrorTargetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTrafficMirrorTargets
func (c *Client) DescribeTrafficMirrorTargetsRequest(input *DescribeTrafficMirrorTargetsInput) DescribeTrafficMirrorTargetsRequest {
	op := &aws.Operation{
		Name:       opDescribeTrafficMirrorTargets,
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
		input = &DescribeTrafficMirrorTargetsInput{}
	}

	req := c.newRequest(op, input, &DescribeTrafficMirrorTargetsOutput{})
	return DescribeTrafficMirrorTargetsRequest{Request: req, Input: input, Copy: c.DescribeTrafficMirrorTargetsRequest}
}

// DescribeTrafficMirrorTargetsRequest is the request type for the
// DescribeTrafficMirrorTargets API operation.
type DescribeTrafficMirrorTargetsRequest struct {
	*aws.Request
	Input *DescribeTrafficMirrorTargetsInput
	Copy  func(*DescribeTrafficMirrorTargetsInput) DescribeTrafficMirrorTargetsRequest
}

// Send marshals and sends the DescribeTrafficMirrorTargets API request.
func (r DescribeTrafficMirrorTargetsRequest) Send(ctx context.Context) (*DescribeTrafficMirrorTargetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTrafficMirrorTargetsResponse{
		DescribeTrafficMirrorTargetsOutput: r.Request.Data.(*DescribeTrafficMirrorTargetsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTrafficMirrorTargetsRequestPaginator returns a paginator for DescribeTrafficMirrorTargets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTrafficMirrorTargetsRequest(input)
//   p := ec2.NewDescribeTrafficMirrorTargetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTrafficMirrorTargetsPaginator(req DescribeTrafficMirrorTargetsRequest) DescribeTrafficMirrorTargetsPaginator {
	return DescribeTrafficMirrorTargetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeTrafficMirrorTargetsInput
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

// DescribeTrafficMirrorTargetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTrafficMirrorTargetsPaginator struct {
	aws.Pager
}

func (p *DescribeTrafficMirrorTargetsPaginator) CurrentPage() *DescribeTrafficMirrorTargetsOutput {
	return p.Pager.CurrentPage().(*DescribeTrafficMirrorTargetsOutput)
}

// DescribeTrafficMirrorTargetsResponse is the response type for the
// DescribeTrafficMirrorTargets API operation.
type DescribeTrafficMirrorTargetsResponse struct {
	*DescribeTrafficMirrorTargetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTrafficMirrorTargets request.
func (r *DescribeTrafficMirrorTargetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
