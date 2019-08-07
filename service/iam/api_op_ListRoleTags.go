// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListRoleTagsRequest
type ListRoleTagsInput struct {
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

	// The name of the IAM role for which you want to see the list of tags.
	//
	// This parameter accepts (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that consist of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListRoleTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRoleTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRoleTagsInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListRoleTagsResponse
type ListRoleTagsOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can use the Marker request parameter to make a subsequent
	// pagination request that retrieves more items. Note that IAM might return
	// fewer than the MaxItems number of results even when more results are available.
	// Check IsTruncated after every call to ensure that you receive all of your
	// results.
	IsTruncated *bool `json:"iam:ListRoleTagsOutput:IsTruncated" type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `json:"iam:ListRoleTagsOutput:Marker" type:"string"`

	// The list of tags currently that is attached to the role. Each tag consists
	// of a key name and an associated value. If no tags are attached to the specified
	// role, the response contains an empty list.
	//
	// Tags is a required field
	Tags []Tag `json:"iam:ListRoleTagsOutput:Tags" type:"list" required:"true"`
}

// String returns the string representation
func (s ListRoleTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRoleTags = "ListRoleTags"

// ListRoleTagsRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the tags that are attached to the specified role. The returned list
// of tags is sorted by tag key. For more information about tagging, see Tagging
// IAM Identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
// in the IAM User Guide.
//
//    // Example sending a request using ListRoleTagsRequest.
//    req := client.ListRoleTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListRoleTags
func (c *Client) ListRoleTagsRequest(input *ListRoleTagsInput) ListRoleTagsRequest {
	op := &aws.Operation{
		Name:       opListRoleTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRoleTagsInput{}
	}

	req := c.newRequest(op, input, &ListRoleTagsOutput{})
	return ListRoleTagsRequest{Request: req, Input: input, Copy: c.ListRoleTagsRequest}
}

// ListRoleTagsRequest is the request type for the
// ListRoleTags API operation.
type ListRoleTagsRequest struct {
	*aws.Request
	Input *ListRoleTagsInput
	Copy  func(*ListRoleTagsInput) ListRoleTagsRequest
}

// Send marshals and sends the ListRoleTags API request.
func (r ListRoleTagsRequest) Send(ctx context.Context) (*ListRoleTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRoleTagsResponse{
		ListRoleTagsOutput: r.Request.Data.(*ListRoleTagsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRoleTagsResponse is the response type for the
// ListRoleTags API operation.
type ListRoleTagsResponse struct {
	*ListRoleTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRoleTags request.
func (r *ListRoleTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
