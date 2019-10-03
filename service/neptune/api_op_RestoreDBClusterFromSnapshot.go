// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/RestoreDBClusterFromSnapshotMessage
type RestoreDBClusterFromSnapshotInput struct {
	_ struct{} `type:"structure"`

	// Provides the list of EC2 Availability Zones that instances in the restored
	// DB cluster can be created in.
	AvailabilityZones []string `locationNameList:"AvailabilityZone" type:"list"`

	// The name of the DB cluster to create from the DB snapshot or DB cluster snapshot.
	// This parameter isn't case-sensitive.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-snapshot-id
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the DB cluster parameter group to associate with the new DB cluster.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBClusterParameterGroup.
	DBClusterParameterGroupName *string `type:"string"`

	// The name of the DB subnet group to use for the new DB cluster.
	//
	// Constraints: If supplied, must match the name of an existing DBSubnetGroup.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// Not supported.
	DatabaseName *string `type:"string"`

	// The list of logs that the restored DB cluster is to export to Amazon CloudWatch
	// Logs.
	EnableCloudwatchLogsExports []string `type:"list"`

	// True to enable mapping of AWS Identity and Access Management (IAM) accounts
	// to database accounts, and otherwise false.
	//
	// Default: false
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The database engine to use for the new DB cluster.
	//
	// Default: The same as source
	//
	// Constraint: Must be compatible with the engine of the source
	//
	// Engine is a required field
	Engine *string `type:"string" required:"true"`

	// The version of the database engine to use for the new DB cluster.
	EngineVersion *string `type:"string"`

	// The AWS KMS key identifier to use when restoring an encrypted DB cluster
	// from a DB snapshot or DB cluster snapshot.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are restoring a DB cluster with the same AWS account that owns
	// the KMS encryption key used to encrypt the new DB cluster, then you can use
	// the KMS key alias instead of the ARN for the KMS encryption key.
	//
	// If you do not specify a value for the KmsKeyId parameter, then the following
	// will occur:
	//
	//    * If the DB snapshot or DB cluster snapshot in SnapshotIdentifier is encrypted,
	//    then the restored DB cluster is encrypted using the KMS key that was used
	//    to encrypt the DB snapshot or DB cluster snapshot.
	//
	//    * If the DB snapshot or DB cluster snapshot in SnapshotIdentifier is not
	//    encrypted, then the restored DB cluster is not encrypted.
	KmsKeyId *string `type:"string"`

	// The name of the option group to use for the restored DB cluster.
	OptionGroupName *string `type:"string"`

	// The port number on which the new DB cluster accepts connections.
	//
	// Constraints: Value must be 1150-65535
	//
	// Default: The same port as the original DB cluster.
	Port *int64 `type:"integer"`

	// The identifier for the DB snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify
	// a DB cluster snapshot. However, you can use only the ARN to specify a DB
	// snapshot.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing Snapshot.
	//
	// SnapshotIdentifier is a required field
	SnapshotIdentifier *string `type:"string" required:"true"`

	// The tags to be assigned to the restored DB cluster.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A list of VPC security groups that the new DB cluster will belong to.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s RestoreDBClusterFromSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreDBClusterFromSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreDBClusterFromSnapshotInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.Engine == nil {
		invalidParams.Add(aws.NewErrParamRequired("Engine"))
	}

	if s.SnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/RestoreDBClusterFromSnapshotResult
type RestoreDBClusterFromSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Neptune DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters action.
	DBCluster *DBCluster `json:"rds:RestoreDBClusterFromSnapshotOutput:DBCluster" type:"structure"`
}

// String returns the string representation
func (s RestoreDBClusterFromSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreDBClusterFromSnapshot = "RestoreDBClusterFromSnapshot"

// RestoreDBClusterFromSnapshotRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
//
// If a DB snapshot is specified, the target DB cluster is created from the
// source DB snapshot with a default configuration and default security group.
//
// If a DB cluster snapshot is specified, the target DB cluster is created from
// the source DB cluster restore point with the same configuration as the original
// source DB cluster, except that the new DB cluster is created with the default
// security group.
//
//    // Example sending a request using RestoreDBClusterFromSnapshotRequest.
//    req := client.RestoreDBClusterFromSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/RestoreDBClusterFromSnapshot
func (c *Client) RestoreDBClusterFromSnapshotRequest(input *RestoreDBClusterFromSnapshotInput) RestoreDBClusterFromSnapshotRequest {
	op := &aws.Operation{
		Name:       opRestoreDBClusterFromSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreDBClusterFromSnapshotInput{}
	}

	req := c.newRequest(op, input, &RestoreDBClusterFromSnapshotOutput{})
	return RestoreDBClusterFromSnapshotRequest{Request: req, Input: input, Copy: c.RestoreDBClusterFromSnapshotRequest}
}

// RestoreDBClusterFromSnapshotRequest is the request type for the
// RestoreDBClusterFromSnapshot API operation.
type RestoreDBClusterFromSnapshotRequest struct {
	*aws.Request
	Input *RestoreDBClusterFromSnapshotInput
	Copy  func(*RestoreDBClusterFromSnapshotInput) RestoreDBClusterFromSnapshotRequest
}

// Send marshals and sends the RestoreDBClusterFromSnapshot API request.
func (r RestoreDBClusterFromSnapshotRequest) Send(ctx context.Context) (*RestoreDBClusterFromSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBClusterFromSnapshotResponse{
		RestoreDBClusterFromSnapshotOutput: r.Request.Data.(*RestoreDBClusterFromSnapshotOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBClusterFromSnapshotResponse is the response type for the
// RestoreDBClusterFromSnapshot API operation.
type RestoreDBClusterFromSnapshotResponse struct {
	*RestoreDBClusterFromSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBClusterFromSnapshot request.
func (r *RestoreDBClusterFromSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
