// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/ListTagsForResourceRequest
type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the AWS CloudHSM resource.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/ListTagsForResourceResponse
type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// One or more tags.
	//
	// TagList is a required field
	TagList []Tag `json:"cloudhsm:ListTagsForResourceOutput:TagList" type:"list" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagsForResource = "ListTagsForResource"

// ListTagsForResourceRequest returns a request value for making API operation for
// Amazon CloudHSM.
//
// This is documentation for AWS CloudHSM Classic. For more information, see
// AWS CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/),
// the AWS CloudHSM Classic User Guide (http://docs.aws.amazon.com/cloudhsm/classic/userguide/),
// and the AWS CloudHSM Classic API Reference (http://docs.aws.amazon.com/cloudhsm/classic/APIReference/).
//
// For information about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide (http://docs.aws.amazon.com/cloudhsm/latest/userguide/),
// and the AWS CloudHSM API Reference (http://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
//
// Returns a list of all tags for the specified AWS CloudHSM resource.
//
//    // Example sending a request using ListTagsForResourceRequest.
//    req := client.ListTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/ListTagsForResource
func (c *Client) ListTagsForResourceRequest(input *ListTagsForResourceInput) ListTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opListTagsForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &ListTagsForResourceOutput{})
	return ListTagsForResourceRequest{Request: req, Input: input, Copy: c.ListTagsForResourceRequest}
}

// ListTagsForResourceRequest is the request type for the
// ListTagsForResource API operation.
type ListTagsForResourceRequest struct {
	*aws.Request
	Input *ListTagsForResourceInput
	Copy  func(*ListTagsForResourceInput) ListTagsForResourceRequest
}

// Send marshals and sends the ListTagsForResource API request.
func (r ListTagsForResourceRequest) Send(ctx context.Context) (*ListTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForResourceResponse{
		ListTagsForResourceOutput: r.Request.Data.(*ListTagsForResourceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForResourceResponse is the response type for the
// ListTagsForResource API operation.
type ListTagsForResourceResponse struct {
	*ListTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForResource request.
func (r *ListTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
