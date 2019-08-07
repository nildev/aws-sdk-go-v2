// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/CreateClusterRequest
type CreateClusterInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zones (AZs) in which the cluster nodes will be created.
	// All nodes belonging to the cluster are placed in these Availability Zones.
	// Use this parameter if you want to distribute the nodes across multiple AZs.
	AvailabilityZones []string `type:"list"`

	// The cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * A name must contain from 1 to 20 alphanumeric characters or hyphens.
	//
	//    * The first character must be a letter.
	//
	//    * A name cannot end with a hyphen or contain two consecutive hyphens.
	//
	// ClusterName is a required field
	ClusterName *string `type:"string" required:"true"`

	// A description of the cluster.
	Description *string `type:"string"`

	// A valid Amazon Resource Name (ARN) that identifies an IAM role. At runtime,
	// DAX will assume this role and use the role's permissions to access DynamoDB
	// on your behalf.
	//
	// IamRoleArn is a required field
	IamRoleArn *string `type:"string" required:"true"`

	// The compute and memory capacity of the nodes in the cluster.
	//
	// NodeType is a required field
	NodeType *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications
	// will be sent.
	//
	// The Amazon SNS topic owner must be same as the DAX cluster owner.
	NotificationTopicArn *string `type:"string"`

	// The parameter group to be associated with the DAX cluster.
	ParameterGroupName *string `type:"string"`

	// Specifies the weekly time range during which maintenance on the DAX cluster
	// is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Valid
	// values for ddd are:
	//
	//    * sun
	//
	//    * mon
	//
	//    * tue
	//
	//    * wed
	//
	//    * thu
	//
	//    * fri
	//
	//    * sat
	//
	// Example: sun:05:00-sun:09:00
	//
	// If you don't specify a preferred maintenance window when you create or modify
	// a cache cluster, DAX assigns a 60-minute maintenance window on a randomly
	// selected day of the week.
	PreferredMaintenanceWindow *string `type:"string"`

	// The number of nodes in the DAX cluster. A replication factor of 1 will create
	// a single-node cluster, without any read replicas. For additional fault tolerance,
	// you can create a multiple node cluster with one or more read replicas. To
	// do this, set ReplicationFactor to 2 or more.
	//
	// AWS recommends that you have at least two read replicas per cluster.
	//
	// ReplicationFactor is a required field
	ReplicationFactor *int64 `type:"integer" required:"true"`

	// Represents the settings used to enable server-side encryption on the cluster.
	SSESpecification *SSESpecification `type:"structure"`

	// A list of security group IDs to be assigned to each node in the DAX cluster.
	// (Each of the security group ID is system-generated.)
	//
	// If this parameter is not specified, DAX assigns the default VPC security
	// group to each node.
	SecurityGroupIds []string `type:"list"`

	// The name of the subnet group to be used for the replication group.
	//
	// DAX clusters can only run in an Amazon VPC environment. All of the subnets
	// that you specify in a subnet group must exist in the same VPC.
	SubnetGroupName *string `type:"string"`

	// A set of tags to associate with the DAX cluster.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if s.IamRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("IamRoleArn"))
	}

	if s.NodeType == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeType"))
	}

	if s.ReplicationFactor == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationFactor"))
	}
	if s.SSESpecification != nil {
		if err := s.SSESpecification.Validate(); err != nil {
			invalidParams.AddNested("SSESpecification", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/CreateClusterResponse
type CreateClusterOutput struct {
	_ struct{} `type:"structure"`

	// A description of the DAX cluster that you have created.
	Cluster *Cluster `json:"dax:CreateClusterOutput:Cluster" type:"structure"`
}

// String returns the string representation
func (s CreateClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCluster = "CreateCluster"

// CreateClusterRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Creates a DAX cluster. All nodes in the cluster run the same DAX caching
// software.
//
//    // Example sending a request using CreateClusterRequest.
//    req := client.CreateClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/CreateCluster
func (c *Client) CreateClusterRequest(input *CreateClusterInput) CreateClusterRequest {
	op := &aws.Operation{
		Name:       opCreateCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterInput{}
	}

	req := c.newRequest(op, input, &CreateClusterOutput{})
	return CreateClusterRequest{Request: req, Input: input, Copy: c.CreateClusterRequest}
}

// CreateClusterRequest is the request type for the
// CreateCluster API operation.
type CreateClusterRequest struct {
	*aws.Request
	Input *CreateClusterInput
	Copy  func(*CreateClusterInput) CreateClusterRequest
}

// Send marshals and sends the CreateCluster API request.
func (r CreateClusterRequest) Send(ctx context.Context) (*CreateClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterResponse{
		CreateClusterOutput: r.Request.Data.(*CreateClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterResponse is the response type for the
// CreateCluster API operation.
type CreateClusterResponse struct {
	*CreateClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCluster request.
func (r *CreateClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
