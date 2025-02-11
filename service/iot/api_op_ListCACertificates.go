// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Input for the ListCACertificates operation.
type ListCACertificatesInput struct {
	_ struct{} `type:"structure"`

	// Determines the order of the results.
	AscendingOrder *bool `location:"querystring" locationName:"isAscendingOrder" type:"boolean"`

	// The marker for the next set of results.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The result page size.
	PageSize *int64 `location:"querystring" locationName:"pageSize" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListCACertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCACertificatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCACertificatesInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListCACertificatesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AscendingOrder != nil {
		v := *s.AscendingOrder

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "isAscendingOrder", protocol.BoolValue(v), metadata)
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
	return nil
}

// The output from the ListCACertificates operation.
type ListCACertificatesOutput struct {
	_ struct{} `type:"structure"`

	// The CA certificates registered in your AWS account.
	Certificates []CACertificate `locationName:"certificates" type:"list"`

	// The current position within the list of CA certificates.
	NextMarker *string `locationName:"nextMarker" type:"string"`
}

// String returns the string representation
func (s ListCACertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListCACertificatesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Certificates != nil {
		v := s.Certificates

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "certificates", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListCACertificates = "ListCACertificates"

// ListCACertificatesRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the CA certificates registered for your AWS account.
//
// The results are paginated with a default page size of 25. You can use the
// returned marker to retrieve additional results.
//
//    // Example sending a request using ListCACertificatesRequest.
//    req := client.ListCACertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListCACertificatesRequest(input *ListCACertificatesInput) ListCACertificatesRequest {
	op := &aws.Operation{
		Name:       opListCACertificates,
		HTTPMethod: "GET",
		HTTPPath:   "/cacertificates",
	}

	if input == nil {
		input = &ListCACertificatesInput{}
	}

	req := c.newRequest(op, input, &ListCACertificatesOutput{})
	return ListCACertificatesRequest{Request: req, Input: input, Copy: c.ListCACertificatesRequest}
}

// ListCACertificatesRequest is the request type for the
// ListCACertificates API operation.
type ListCACertificatesRequest struct {
	*aws.Request
	Input *ListCACertificatesInput
	Copy  func(*ListCACertificatesInput) ListCACertificatesRequest
}

// Send marshals and sends the ListCACertificates API request.
func (r ListCACertificatesRequest) Send(ctx context.Context) (*ListCACertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCACertificatesResponse{
		ListCACertificatesOutput: r.Request.Data.(*ListCACertificatesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListCACertificatesResponse is the response type for the
// ListCACertificates API operation.
type ListCACertificatesResponse struct {
	*ListCACertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCACertificates request.
func (r *ListCACertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
