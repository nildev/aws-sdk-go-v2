// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains the information about the request to list the
// traffic policies that are associated with the current AWS account.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListTrafficPoliciesRequest
type ListTrafficPoliciesInput struct {
	_ struct{} `type:"structure"`

	// (Optional) The maximum number of traffic policies that you want Amazon Route
	// 53 to return in response to this request. If you have more than MaxItems
	// traffic policies, the value of IsTruncated in the response is true, and the
	// value of TrafficPolicyIdMarker is the ID of the first traffic policy that
	// Route 53 will return if you submit another request.
	MaxItems *string `location:"querystring" locationName:"maxitems" type:"string"`

	// (Conditional) For your first request to ListTrafficPolicies, don't include
	// the TrafficPolicyIdMarker parameter.
	//
	// If you have more traffic policies than the value of MaxItems, ListTrafficPolicies
	// returns only the first MaxItems traffic policies. To get the next group of
	// policies, submit another request to ListTrafficPolicies. For the value of
	// TrafficPolicyIdMarker, specify the value of TrafficPolicyIdMarker that was
	// returned in the previous response.
	TrafficPolicyIdMarker *string `location:"querystring" locationName:"trafficpolicyid" min:"1" type:"string"`
}

// String returns the string representation
func (s ListTrafficPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTrafficPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTrafficPoliciesInput"}
	if s.TrafficPolicyIdMarker != nil && len(*s.TrafficPolicyIdMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrafficPolicyIdMarker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTrafficPoliciesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxitems", protocol.StringValue(v), metadata)
	}
	if s.TrafficPolicyIdMarker != nil {
		v := *s.TrafficPolicyIdMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "trafficpolicyid", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains the response information for the request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListTrafficPoliciesResponse
type ListTrafficPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more traffic policies to be listed.
	// If the response was truncated, you can get the next group of traffic policies
	// by submitting another ListTrafficPolicies request and specifying the value
	// of TrafficPolicyIdMarker in the TrafficPolicyIdMarker request parameter.
	//
	// IsTruncated is a required field
	IsTruncated *bool `type:"boolean" required:"true"`

	// The value that you specified for the MaxItems parameter in the ListTrafficPolicies
	// request that produced the current response.
	//
	// MaxItems is a required field
	MaxItems *string `type:"string" required:"true"`

	// If the value of IsTruncated is true, TrafficPolicyIdMarker is the ID of the
	// first traffic policy in the next group of MaxItems traffic policies.
	//
	// TrafficPolicyIdMarker is a required field
	TrafficPolicyIdMarker *string `min:"1" type:"string" required:"true"`

	// A list that contains one TrafficPolicySummary element for each traffic policy
	// that was created by the current AWS account.
	//
	// TrafficPolicySummaries is a required field
	TrafficPolicySummaries []TrafficPolicySummary `locationNameList:"TrafficPolicySummary" type:"list" required:"true"`
}

// String returns the string representation
func (s ListTrafficPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTrafficPoliciesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxItems", protocol.StringValue(v), metadata)
	}
	if s.TrafficPolicyIdMarker != nil {
		v := *s.TrafficPolicyIdMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TrafficPolicyIdMarker", protocol.StringValue(v), metadata)
	}
	if s.TrafficPolicySummaries != nil {
		v := s.TrafficPolicySummaries

		metadata := protocol.Metadata{ListLocationName: "TrafficPolicySummary"}
		ls0 := e.List(protocol.BodyTarget, "TrafficPolicySummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListTrafficPolicies = "ListTrafficPolicies"

// ListTrafficPoliciesRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets information about the latest version for every traffic policy that is
// associated with the current AWS account. Policies are listed in the order
// that they were created in.
//
//    // Example sending a request using ListTrafficPoliciesRequest.
//    req := client.ListTrafficPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListTrafficPolicies
func (c *Client) ListTrafficPoliciesRequest(input *ListTrafficPoliciesInput) ListTrafficPoliciesRequest {
	op := &aws.Operation{
		Name:       opListTrafficPolicies,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/trafficpolicies",
	}

	if input == nil {
		input = &ListTrafficPoliciesInput{}
	}

	req := c.newRequest(op, input, &ListTrafficPoliciesOutput{})
	return ListTrafficPoliciesRequest{Request: req, Input: input, Copy: c.ListTrafficPoliciesRequest}
}

// ListTrafficPoliciesRequest is the request type for the
// ListTrafficPolicies API operation.
type ListTrafficPoliciesRequest struct {
	*aws.Request
	Input *ListTrafficPoliciesInput
	Copy  func(*ListTrafficPoliciesInput) ListTrafficPoliciesRequest
}

// Send marshals and sends the ListTrafficPolicies API request.
func (r ListTrafficPoliciesRequest) Send(ctx context.Context) (*ListTrafficPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTrafficPoliciesResponse{
		ListTrafficPoliciesOutput: r.Request.Data.(*ListTrafficPoliciesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTrafficPoliciesResponse is the response type for the
// ListTrafficPolicies API operation.
type ListTrafficPoliciesResponse struct {
	*ListTrafficPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTrafficPolicies request.
func (r *ListTrafficPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
