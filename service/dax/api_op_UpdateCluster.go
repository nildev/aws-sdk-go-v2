// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/UpdateClusterRequest
type UpdateClusterInput struct {
	_ struct{} `type:"structure"`

	// The name of the DAX cluster to be modified.
	//
	// ClusterName is a required field
	ClusterName *string `type:"string" required:"true"`

	// A description of the changes being made to the cluster.
	Description *string `type:"string"`

	// The Amazon Resource Name (ARN) that identifies the topic.
	NotificationTopicArn *string `type:"string"`

	// The current state of the topic.
	NotificationTopicStatus *string `type:"string"`

	// The name of a parameter group for this cluster.
	ParameterGroupName *string `type:"string"`

	// A range of time when maintenance of DAX cluster software will be performed.
	// For example: sun:01:00-sun:09:00. Cluster maintenance normally takes less
	// than 30 minutes, and is performed automatically within the maintenance window.
	PreferredMaintenanceWindow *string `type:"string"`

	// A list of user-specified security group IDs to be assigned to each node in
	// the DAX cluster. If this parameter is not specified, DAX assigns the default
	// VPC security group to each node.
	SecurityGroupIds []string `type:"list"`
}

// String returns the string representation
func (s UpdateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateClusterInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/UpdateClusterResponse
type UpdateClusterOutput struct {
	_ struct{} `type:"structure"`

	// A description of the DAX cluster, after it has been modified.
	Cluster *Cluster `json:"dax:UpdateClusterOutput:Cluster" type:"structure"`
}

// String returns the string representation
func (s UpdateClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateCluster = "UpdateCluster"

// UpdateClusterRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Modifies the settings for a DAX cluster. You can use this action to change
// one or more cluster configuration parameters by specifying the parameters
// and the new values.
//
//    // Example sending a request using UpdateClusterRequest.
//    req := client.UpdateClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/UpdateCluster
func (c *Client) UpdateClusterRequest(input *UpdateClusterInput) UpdateClusterRequest {
	op := &aws.Operation{
		Name:       opUpdateCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateClusterInput{}
	}

	req := c.newRequest(op, input, &UpdateClusterOutput{})
	return UpdateClusterRequest{Request: req, Input: input, Copy: c.UpdateClusterRequest}
}

// UpdateClusterRequest is the request type for the
// UpdateCluster API operation.
type UpdateClusterRequest struct {
	*aws.Request
	Input *UpdateClusterInput
	Copy  func(*UpdateClusterInput) UpdateClusterRequest
}

// Send marshals and sends the UpdateCluster API request.
func (r UpdateClusterRequest) Send(ctx context.Context) (*UpdateClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClusterResponse{
		UpdateClusterOutput: r.Request.Data.(*UpdateClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClusterResponse is the response type for the
// UpdateCluster API operation.
type UpdateClusterResponse struct {
	*UpdateClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCluster request.
func (r *UpdateClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
