// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListUserTagsRequest
type ListUserTagsInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// (Optional) Use this only when paginating results to indicate the maximum
	// number of items that you want in the response. If additional items exist
	// beyond the maximum that you specify, the IsTruncated response element is
	// true.
	//
	// If you do not include this parameter, it defaults to 100. Note that IAM might
	// return fewer results, even when more results are available. In that case,
	// the IsTruncated response element returns true, and Marker contains a value
	// to include in the subsequent call that tells the service where to continue
	// from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The name of the IAM user whose tags you want to see.
	//
	// This parameter accepts (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that consist of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: =,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListUserTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUserTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUserTagsInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListUserTagsResponse
type ListUserTagsOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can use the Marker request parameter to make a subsequent
	// pagination request that retrieves more items. Note that IAM might return
	// fewer than the MaxItems number of results even when more results are available.
	// Check IsTruncated after every call to ensure that you receive all of your
	// results.
	IsTruncated *bool `json:"iam:ListUserTagsOutput:IsTruncated" type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `json:"iam:ListUserTagsOutput:Marker" type:"string"`

	// The list of tags that are currently attached to the user. Each tag consists
	// of a key name and an associated value. If no tags are attached to the specified
	// user, the response contains an empty list.
	//
	// Tags is a required field
	Tags []Tag `json:"iam:ListUserTagsOutput:Tags" type:"list" required:"true"`
}

// String returns the string representation
func (s ListUserTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListUserTags = "ListUserTags"

// ListUserTagsRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the tags that are attached to the specified user. The returned list
// of tags is sorted by tag key. For more information about tagging, see Tagging
// IAM Identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
// in the IAM User Guide.
//
//    // Example sending a request using ListUserTagsRequest.
//    req := client.ListUserTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListUserTags
func (c *Client) ListUserTagsRequest(input *ListUserTagsInput) ListUserTagsRequest {
	op := &aws.Operation{
		Name:       opListUserTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListUserTagsInput{}
	}

	req := c.newRequest(op, input, &ListUserTagsOutput{})
	return ListUserTagsRequest{Request: req, Input: input, Copy: c.ListUserTagsRequest}
}

// ListUserTagsRequest is the request type for the
// ListUserTags API operation.
type ListUserTagsRequest struct {
	*aws.Request
	Input *ListUserTagsInput
	Copy  func(*ListUserTagsInput) ListUserTagsRequest
}

// Send marshals and sends the ListUserTags API request.
func (r ListUserTagsRequest) Send(ctx context.Context) (*ListUserTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUserTagsResponse{
		ListUserTagsOutput: r.Request.Data.(*ListUserTagsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListUserTagsResponse is the response type for the
// ListUserTags API operation.
type ListUserTagsResponse struct {
	*ListUserTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUserTags request.
func (r *ListUserTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
