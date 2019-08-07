// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/StopDBClusterMessage
type StopDBClusterInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the cluster to stop. Example: docdb-2019-05-28-15-24-52
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StopDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/StopDBClusterResult
type StopDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about a DB cluster.
	DBCluster *DBCluster `json:"rds:StopDBClusterOutput:DBCluster" type:"structure"`
}

// String returns the string representation
func (s StopDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopDBCluster = "StopDBCluster"

// StopDBClusterRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Stops the running cluster that is specified by DBClusterIdentifier. The cluster
// must be in the available state. For more information, see Stopping and Starting
// an Amazon DocumentDB Cluster (https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-stop-start.html).
//
//    // Example sending a request using StopDBClusterRequest.
//    req := client.StopDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/StopDBCluster
func (c *Client) StopDBClusterRequest(input *StopDBClusterInput) StopDBClusterRequest {
	op := &aws.Operation{
		Name:       opStopDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopDBClusterInput{}
	}

	req := c.newRequest(op, input, &StopDBClusterOutput{})
	return StopDBClusterRequest{Request: req, Input: input, Copy: c.StopDBClusterRequest}
}

// StopDBClusterRequest is the request type for the
// StopDBCluster API operation.
type StopDBClusterRequest struct {
	*aws.Request
	Input *StopDBClusterInput
	Copy  func(*StopDBClusterInput) StopDBClusterRequest
}

// Send marshals and sends the StopDBCluster API request.
func (r StopDBClusterRequest) Send(ctx context.Context) (*StopDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopDBClusterResponse{
		StopDBClusterOutput: r.Request.Data.(*StopDBClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopDBClusterResponse is the response type for the
// StopDBCluster API operation.
type StopDBClusterResponse struct {
	*StopDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopDBCluster request.
func (r *StopDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
