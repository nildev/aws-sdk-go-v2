// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DeleteCacheCluster operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteCacheClusterMessage
type DeleteCacheClusterInput struct {
	_ struct{} `type:"structure"`

	// The cluster identifier for the cluster to be deleted. This parameter is not
	// case sensitive.
	//
	// CacheClusterId is a required field
	CacheClusterId *string `type:"string" required:"true"`

	// The user-supplied name of a final cluster snapshot. This is the unique name
	// that identifies the snapshot. ElastiCache creates the snapshot, and then
	// deletes the cluster immediately afterward.
	FinalSnapshotIdentifier *string `type:"string"`
}

// String returns the string representation
func (s DeleteCacheClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCacheClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCacheClusterInput"}

	if s.CacheClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheClusterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteCacheClusterResult
type DeleteCacheClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains all of the attributes of a specific cluster.
	CacheCluster *CacheCluster `type:"structure"`
}

// String returns the string representation
func (s DeleteCacheClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCacheCluster = "DeleteCacheCluster"

// DeleteCacheClusterRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Deletes a previously provisioned cluster. DeleteCacheCluster deletes all
// associated cache nodes, node endpoints and the cluster itself. When you receive
// a successful response from this operation, Amazon ElastiCache immediately
// begins deleting the cluster; you cannot cancel or revert this operation.
//
// This operation is not valid for:
//
//    * Redis (cluster mode enabled) clusters
//
//    * A cluster that is the last read replica of a replication group
//
//    * A node group (shard) that has Multi-AZ mode enabled
//
//    * A cluster from a Redis (cluster mode enabled) replication group
//
//    * A cluster that is not in the available state
//
//    // Example sending a request using DeleteCacheClusterRequest.
//    req := client.DeleteCacheClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteCacheCluster
func (c *Client) DeleteCacheClusterRequest(input *DeleteCacheClusterInput) DeleteCacheClusterRequest {
	op := &aws.Operation{
		Name:       opDeleteCacheCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCacheClusterInput{}
	}

	req := c.newRequest(op, input, &DeleteCacheClusterOutput{})
	return DeleteCacheClusterRequest{Request: req, Input: input, Copy: c.DeleteCacheClusterRequest}
}

// DeleteCacheClusterRequest is the request type for the
// DeleteCacheCluster API operation.
type DeleteCacheClusterRequest struct {
	*aws.Request
	Input *DeleteCacheClusterInput
	Copy  func(*DeleteCacheClusterInput) DeleteCacheClusterRequest
}

// Send marshals and sends the DeleteCacheCluster API request.
func (r DeleteCacheClusterRequest) Send(ctx context.Context) (*DeleteCacheClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCacheClusterResponse{
		DeleteCacheClusterOutput: r.Request.Data.(*DeleteCacheClusterOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCacheClusterResponse is the response type for the
// DeleteCacheCluster API operation.
type DeleteCacheClusterResponse struct {
	*DeleteCacheClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCacheCluster request.
func (r *DeleteCacheClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
