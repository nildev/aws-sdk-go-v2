// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeClientVpnTargetNetworksRequest
type DescribeClientVpnTargetNetworksInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the target network associations.
	AssociationIds []string `locationNameList:"item" type:"list"`

	// The ID of the Client VPN endpoint.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. Filter names and values are case-sensitive.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results can be seen by sending another request with the nextToken
	// value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeClientVpnTargetNetworksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeClientVpnTargetNetworksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeClientVpnTargetNetworksInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeClientVpnTargetNetworksResult
type DescribeClientVpnTargetNetworksOutput struct {
	_ struct{} `type:"structure"`

	// Information about the associated target networks.
	ClientVpnTargetNetworks []TargetNetwork `json:"ec2:DescribeClientVpnTargetNetworksOutput:ClientVpnTargetNetworks" locationName:"clientVpnTargetNetworks" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `json:"ec2:DescribeClientVpnTargetNetworksOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeClientVpnTargetNetworksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClientVpnTargetNetworks = "DescribeClientVpnTargetNetworks"

// DescribeClientVpnTargetNetworksRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the target networks associated with the specified Client VPN endpoint.
//
//    // Example sending a request using DescribeClientVpnTargetNetworksRequest.
//    req := client.DescribeClientVpnTargetNetworksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeClientVpnTargetNetworks
func (c *Client) DescribeClientVpnTargetNetworksRequest(input *DescribeClientVpnTargetNetworksInput) DescribeClientVpnTargetNetworksRequest {
	op := &aws.Operation{
		Name:       opDescribeClientVpnTargetNetworks,
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
		input = &DescribeClientVpnTargetNetworksInput{}
	}

	req := c.newRequest(op, input, &DescribeClientVpnTargetNetworksOutput{})
	return DescribeClientVpnTargetNetworksRequest{Request: req, Input: input, Copy: c.DescribeClientVpnTargetNetworksRequest}
}

// DescribeClientVpnTargetNetworksRequest is the request type for the
// DescribeClientVpnTargetNetworks API operation.
type DescribeClientVpnTargetNetworksRequest struct {
	*aws.Request
	Input *DescribeClientVpnTargetNetworksInput
	Copy  func(*DescribeClientVpnTargetNetworksInput) DescribeClientVpnTargetNetworksRequest
}

// Send marshals and sends the DescribeClientVpnTargetNetworks API request.
func (r DescribeClientVpnTargetNetworksRequest) Send(ctx context.Context) (*DescribeClientVpnTargetNetworksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClientVpnTargetNetworksResponse{
		DescribeClientVpnTargetNetworksOutput: r.Request.Data.(*DescribeClientVpnTargetNetworksOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeClientVpnTargetNetworksRequestPaginator returns a paginator for DescribeClientVpnTargetNetworks.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeClientVpnTargetNetworksRequest(input)
//   p := ec2.NewDescribeClientVpnTargetNetworksRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeClientVpnTargetNetworksPaginator(req DescribeClientVpnTargetNetworksRequest) DescribeClientVpnTargetNetworksPaginator {
	return DescribeClientVpnTargetNetworksPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeClientVpnTargetNetworksInput
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

// DescribeClientVpnTargetNetworksPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeClientVpnTargetNetworksPaginator struct {
	aws.Pager
}

func (p *DescribeClientVpnTargetNetworksPaginator) CurrentPage() *DescribeClientVpnTargetNetworksOutput {
	return p.Pager.CurrentPage().(*DescribeClientVpnTargetNetworksOutput)
}

// DescribeClientVpnTargetNetworksResponse is the response type for the
// DescribeClientVpnTargetNetworks API operation.
type DescribeClientVpnTargetNetworksResponse struct {
	*DescribeClientVpnTargetNetworksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClientVpnTargetNetworks request.
func (r *DescribeClientVpnTargetNetworksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
