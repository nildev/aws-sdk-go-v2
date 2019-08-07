// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeTags.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DescribeTagsInput
type DescribeTagsInput struct {
	_ struct{} `type:"structure"`

	// The names of the load balancers.
	//
	// LoadBalancerNames is a required field
	LoadBalancerNames []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTagsInput"}

	if s.LoadBalancerNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerNames"))
	}
	if s.LoadBalancerNames != nil && len(s.LoadBalancerNames) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LoadBalancerNames", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output for DescribeTags.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DescribeTagsOutput
type DescribeTagsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the tags.
	TagDescriptions []TagDescription `json:"elasticloadbalancing:DescribeTagsOutput:TagDescriptions" type:"list"`
}

// String returns the string representation
func (s DescribeTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTags = "DescribeTags"

// DescribeTagsRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Describes the tags associated with the specified load balancers.
//
//    // Example sending a request using DescribeTagsRequest.
//    req := client.DescribeTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DescribeTags
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
