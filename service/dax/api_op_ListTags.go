// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/ListTagsRequest
type ListTagsInput struct {
	_ struct{} `type:"structure"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token.
	NextToken *string `type:"string"`

	// The name of the DAX resource to which the tags belong.
	//
	// ResourceName is a required field
	ResourceName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsInput"}

	if s.ResourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/ListTagsResponse
type ListTagsOutput struct {
	_ struct{} `type:"structure"`

	// If this value is present, there are additional results to be displayed. To
	// retrieve them, call ListTags again, with NextToken set to this value.
	NextToken *string `json:"dax:ListTagsOutput:NextToken" type:"string"`

	// A list of tags currently associated with the DAX cluster.
	Tags []Tag `json:"dax:ListTagsOutput:Tags" type:"list"`
}

// String returns the string representation
func (s ListTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTags = "ListTags"

// ListTagsRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// List all of the tags for a DAX cluster. You can call ListTags up to 10 times
// per second, per account.
//
//    // Example sending a request using ListTagsRequest.
//    req := client.ListTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/ListTags
func (c *Client) ListTagsRequest(input *ListTagsInput) ListTagsRequest {
	op := &aws.Operation{
		Name:       opListTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsInput{}
	}

	req := c.newRequest(op, input, &ListTagsOutput{})
	return ListTagsRequest{Request: req, Input: input, Copy: c.ListTagsRequest}
}

// ListTagsRequest is the request type for the
// ListTags API operation.
type ListTagsRequest struct {
	*aws.Request
	Input *ListTagsInput
	Copy  func(*ListTagsInput) ListTagsRequest
}

// Send marshals and sends the ListTags API request.
func (r ListTagsRequest) Send(ctx context.Context) (*ListTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsResponse{
		ListTagsOutput: r.Request.Data.(*ListTagsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsResponse is the response type for the
// ListTags API operation.
type ListTagsResponse struct {
	*ListTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTags request.
func (r *ListTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
