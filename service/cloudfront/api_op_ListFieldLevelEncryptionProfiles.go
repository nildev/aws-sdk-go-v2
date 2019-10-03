// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListFieldLevelEncryptionProfilesRequest
type ListFieldLevelEncryptionProfilesInput struct {
	_ struct{} `type:"structure"`

	// Use this when paginating results to indicate where to begin in your list
	// of profiles. The results include profiles in the list that occur after the
	// marker. To get the next page of results, set the Marker to the value of the
	// NextMarker from the current page's response (which is also the ID of the
	// last profile on that page).
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The maximum number of field-level encryption profiles you want in the response
	// body.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`
}

// String returns the string representation
func (s ListFieldLevelEncryptionProfilesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFieldLevelEncryptionProfilesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListFieldLevelEncryptionProfilesResult
type ListFieldLevelEncryptionProfilesOutput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryptionProfileList"`

	// Returns a list of the field-level encryption profiles that have been created
	// in CloudFront for this account.
	FieldLevelEncryptionProfileList *FieldLevelEncryptionProfileList `json:"cloudfront:ListFieldLevelEncryptionProfilesOutput:FieldLevelEncryptionProfileList" type:"structure"`
}

// String returns the string representation
func (s ListFieldLevelEncryptionProfilesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFieldLevelEncryptionProfilesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FieldLevelEncryptionProfileList != nil {
		v := s.FieldLevelEncryptionProfileList

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryptionProfileList", v, metadata)
	}
	return nil
}

const opListFieldLevelEncryptionProfiles = "ListFieldLevelEncryptionProfiles2019_03_26"

// ListFieldLevelEncryptionProfilesRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Request a list of field-level encryption profiles that have been created
// in CloudFront for this account.
//
//    // Example sending a request using ListFieldLevelEncryptionProfilesRequest.
//    req := client.ListFieldLevelEncryptionProfilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListFieldLevelEncryptionProfiles
func (c *Client) ListFieldLevelEncryptionProfilesRequest(input *ListFieldLevelEncryptionProfilesInput) ListFieldLevelEncryptionProfilesRequest {
	op := &aws.Operation{
		Name:       opListFieldLevelEncryptionProfiles,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/field-level-encryption-profile",
	}

	if input == nil {
		input = &ListFieldLevelEncryptionProfilesInput{}
	}

	req := c.newRequest(op, input, &ListFieldLevelEncryptionProfilesOutput{})
	return ListFieldLevelEncryptionProfilesRequest{Request: req, Input: input, Copy: c.ListFieldLevelEncryptionProfilesRequest}
}

// ListFieldLevelEncryptionProfilesRequest is the request type for the
// ListFieldLevelEncryptionProfiles API operation.
type ListFieldLevelEncryptionProfilesRequest struct {
	*aws.Request
	Input *ListFieldLevelEncryptionProfilesInput
	Copy  func(*ListFieldLevelEncryptionProfilesInput) ListFieldLevelEncryptionProfilesRequest
}

// Send marshals and sends the ListFieldLevelEncryptionProfiles API request.
func (r ListFieldLevelEncryptionProfilesRequest) Send(ctx context.Context) (*ListFieldLevelEncryptionProfilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFieldLevelEncryptionProfilesResponse{
		ListFieldLevelEncryptionProfilesOutput: r.Request.Data.(*ListFieldLevelEncryptionProfilesOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListFieldLevelEncryptionProfilesResponse is the response type for the
// ListFieldLevelEncryptionProfiles API operation.
type ListFieldLevelEncryptionProfilesResponse struct {
	*ListFieldLevelEncryptionProfilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFieldLevelEncryptionProfiles request.
func (r *ListFieldLevelEncryptionProfilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
