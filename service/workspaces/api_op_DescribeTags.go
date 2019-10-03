// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeTagsRequest
type DescribeTagsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpaces resource. The supported resource types are
	// WorkSpaces, registered directories, images, custom bundles, and IP access
	// control groups.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTagsInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeTagsResult
type DescribeTagsOutput struct {
	_ struct{} `type:"structure"`

	// The tags.
	TagList []Tag `json:"workspaces:DescribeTagsOutput:TagList" type:"list"`
}

// String returns the string representation
func (s DescribeTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTags = "DescribeTags"

// DescribeTagsRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Describes the specified tags for the specified WorkSpaces resource.
//
//    // Example sending a request using DescribeTagsRequest.
//    req := client.DescribeTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeTags
func (c *Client) DescribeTagsRequest(input *DescribeTagsInput) DescribeTagsRequest {
	op := &aws.Operation{
		Name:       opDescribeTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTagsInput{}
	}

	req := c.newRequest(op, input, &DescribeTagsOutput{})
	return DescribeTagsRequest{Request: req, Input: input, Copy: c.DescribeTagsRequest}
}

// DescribeTagsRequest is the request type for the
// DescribeTags API operation.
type DescribeTagsRequest struct {
	*aws.Request
	Input *DescribeTagsInput
	Copy  func(*DescribeTagsInput) DescribeTagsRequest
}

// Send marshals and sends the DescribeTags API request.
func (r DescribeTagsRequest) Send(ctx context.Context) (*DescribeTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTagsResponse{
		DescribeTagsOutput: r.Request.Data.(*DescribeTagsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTagsResponse is the response type for the
// DescribeTags API operation.
type DescribeTagsResponse struct {
	*DescribeTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTags request.
func (r *DescribeTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
