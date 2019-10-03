// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListAttachedPoliciesInput struct {
	_ struct{} `type:"structure"`

	// The token to retrieve the next set of results.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The maximum number of results to be returned per request.
	PageSize *int64 `location:"querystring" locationName:"pageSize" min:"1" type:"integer"`

	// When true, recursively list attached policies.
	Recursive *bool `location:"querystring" locationName:"recursive" type:"boolean"`

	// The group for which the policies will be listed.
	//
	// Target is a required field
	Target *string `location:"uri" locationName:"target" type:"string" required:"true"`
}

// String returns the string representation
func (s ListAttachedPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAttachedPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAttachedPoliciesInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if s.Target == nil {
		invalidParams.Add(aws.NewErrParamRequired("Target"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAttachedPoliciesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Target != nil {
		v := *s.Target

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "target", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "pageSize", protocol.Int64Value(v), metadata)
	}
	if s.Recursive != nil {
		v := *s.Recursive

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "recursive", protocol.BoolValue(v), metadata)
	}
	return nil
}

type ListAttachedPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The token to retrieve the next set of results, or ``null`` if there are no
	// more results.
	NextMarker *string `json:"iot:ListAttachedPoliciesOutput:NextMarker" locationName:"nextMarker" type:"string"`

	// The policies.
	Policies []Policy `json:"iot:ListAttachedPoliciesOutput:Policies" locationName:"policies" type:"list"`
}

// String returns the string representation
func (s ListAttachedPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAttachedPoliciesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policies != nil {
		v := s.Policies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "policies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListAttachedPolicies = "ListAttachedPolicies"

// ListAttachedPoliciesRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the policies attached to the specified thing group.
//
//    // Example sending a request using ListAttachedPoliciesRequest.
//    req := client.ListAttachedPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListAttachedPoliciesRequest(input *ListAttachedPoliciesInput) ListAttachedPoliciesRequest {
	op := &aws.Operation{
		Name:       opListAttachedPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/attached-policies/{target}",
	}

	if input == nil {
		input = &ListAttachedPoliciesInput{}
	}

	req := c.newRequest(op, input, &ListAttachedPoliciesOutput{})
	return ListAttachedPoliciesRequest{Request: req, Input: input, Copy: c.ListAttachedPoliciesRequest}
}

// ListAttachedPoliciesRequest is the request type for the
// ListAttachedPolicies API operation.
type ListAttachedPoliciesRequest struct {
	*aws.Request
	Input *ListAttachedPoliciesInput
	Copy  func(*ListAttachedPoliciesInput) ListAttachedPoliciesRequest
}

// Send marshals and sends the ListAttachedPolicies API request.
func (r ListAttachedPoliciesRequest) Send(ctx context.Context) (*ListAttachedPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAttachedPoliciesResponse{
		ListAttachedPoliciesOutput: r.Request.Data.(*ListAttachedPoliciesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAttachedPoliciesResponse is the response type for the
// ListAttachedPolicies API operation.
type ListAttachedPoliciesResponse struct {
	*ListAttachedPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAttachedPolicies request.
func (r *ListAttachedPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
