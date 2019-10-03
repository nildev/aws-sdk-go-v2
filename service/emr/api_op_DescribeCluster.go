// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// This input determines which cluster to describe.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/DescribeClusterInput
type DescribeClusterInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the cluster to describe.
	//
	// ClusterId is a required field
	ClusterId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeClusterInput"}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This output contains the description of the cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/DescribeClusterOutput
type DescribeClusterOutput struct {
	_ struct{} `type:"structure"`

	// This output contains the details for the requested cluster.
	Cluster *Cluster `json:"elasticmapreduce:DescribeClusterOutput:Cluster" type:"structure"`
}

// String returns the string representation
func (s DescribeClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCluster = "DescribeCluster"

// DescribeClusterRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Provides cluster-level details including status, hardware and software configuration,
// VPC settings, and so on.
//
//    // Example sending a request using DescribeClusterRequest.
//    req := client.DescribeClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/DescribeCluster
func (c *Client) DescribeClusterRequest(input *DescribeClusterInput) DescribeClusterRequest {
	op := &aws.Operation{
		Name:       opDescribeCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeClusterInput{}
	}

	req := c.newRequest(op, input, &DescribeClusterOutput{})
	return DescribeClusterRequest{Request: req, Input: input, Copy: c.DescribeClusterRequest}
}

// DescribeClusterRequest is the request type for the
// DescribeCluster API operation.
type DescribeClusterRequest struct {
	*aws.Request
	Input *DescribeClusterInput
	Copy  func(*DescribeClusterInput) DescribeClusterRequest
}

// Send marshals and sends the DescribeCluster API request.
func (r DescribeClusterRequest) Send(ctx context.Context) (*DescribeClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClusterResponse{
		DescribeClusterOutput: r.Request.Data.(*DescribeClusterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeClusterResponse is the response type for the
// DescribeCluster API operation.
type DescribeClusterResponse struct {
	*DescribeClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCluster request.
func (r *DescribeClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
