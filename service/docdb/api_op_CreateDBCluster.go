// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to CreateDBCluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/CreateDBClusterMessage
type CreateDBClusterInput struct {
	_ struct{} `type:"structure"`

	// A list of Amazon EC2 Availability Zones that instances in the DB cluster
	// can be created in.
	AvailabilityZones []string `locationNameList:"AvailabilityZone" type:"list"`

	// The number of days for which automated backups are retained. You must specify
	// a minimum value of 1.
	//
	// Default: 1
	//
	// Constraints:
	//
	//    * Must be a value from 1 to 35.
	BackupRetentionPeriod *int64 `type:"integer"`

	// The DB cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//    * The first character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the DB cluster parameter group to associate with this DB cluster.
	DBClusterParameterGroupName *string `type:"string"`

	// A DB subnet group to associate with this DB cluster.
	//
	// Constraints: Must match the name of an existing DBSubnetGroup. Must not be
	// default.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// Specifies whether this cluster can be deleted. If DeletionProtection is enabled,
	// the cluster cannot be deleted unless it is modified and DeletionProtection
	// is disabled. DeletionProtection protects clusters from being accidentally
	// deleted.
	DeletionProtection *bool `type:"boolean"`

	// A list of log types that need to be enabled for exporting to Amazon CloudWatch
	// Logs.
	EnableCloudwatchLogsExports []string `type:"list"`

	// The name of the database engine to be used for this DB cluster.
	//
	// Valid values: docdb
	//
	// Engine is a required field
	Engine *string `type:"string" required:"true"`

	// The version number of the database engine to use.
	EngineVersion *string `type:"string"`

	// The AWS KMS key identifier for an encrypted DB cluster.
	//
	// The AWS KMS key identifier is the Amazon Resource Name (ARN) for the AWS
	// KMS encryption key. If you are creating a DB cluster using the same AWS account
	// that owns the AWS KMS encryption key that is used to encrypt the new DB cluster,
	// you can use the AWS KMS key alias instead of the ARN for the AWS KMS encryption
	// key.
	//
	// If an encryption key is not specified in KmsKeyId:
	//
	//    * If ReplicationSourceIdentifier identifies an encrypted source, then
	//    Amazon DocumentDB uses the encryption key that is used to encrypt the
	//    source. Otherwise, Amazon DocumentDB uses your default encryption key.
	//
	//    * If the StorageEncrypted parameter is true and ReplicationSourceIdentifier
	//    is not specified, Amazon DocumentDB uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account. Your AWS
	// account has a different default encryption key for each AWS Region.
	//
	// If you create a replica of an encrypted DB cluster in another AWS Region,
	// you must set KmsKeyId to a KMS key ID that is valid in the destination AWS
	// Region. This key is used to encrypt the replica in that AWS Region.
	KmsKeyId *string `type:"string"`

	// The password for the master database user. This password can contain any
	// printable ASCII character except forward slash (/), double quote ("), or
	// the "at" symbol (@).
	//
	// Constraints: Must contain from 8 to 41 characters.
	//
	// MasterUserPassword is a required field
	MasterUserPassword *string `type:"string" required:"true"`

	// The name of the master user for the DB cluster.
	//
	// Constraints:
	//
	//    * Must be from 1 to 16 letters or numbers.
	//
	//    * The first character must be a letter.
	//
	//    * Cannot be a reserved word for the chosen database engine.
	//
	// MasterUsername is a required field
	MasterUsername *string `type:"string" required:"true"`

	// The port number on which the instances in the DB cluster accept connections.
	Port *int64 `type:"integer"`

	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region.
	//
	// Constraints:
	//
	//    * Must be in the format hh24:mi-hh24:mi.
	//
	//    * Must be in Universal Coordinated Time (UTC).
	//
	//    * Must not conflict with the preferred maintenance window.
	//
	//    * Must be at least 30 minutes.
	PreferredBackupWindow *string `type:"string"`

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region, occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `type:"string"`

	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted *bool `type:"boolean"`

	// The tags to be assigned to the DB cluster.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A list of EC2 VPC security groups to associate with this DB cluster.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s CreateDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.Engine == nil {
		invalidParams.Add(aws.NewErrParamRequired("Engine"))
	}

	if s.MasterUserPassword == nil {
		invalidParams.Add(aws.NewErrParamRequired("MasterUserPassword"))
	}

	if s.MasterUsername == nil {
		invalidParams.Add(aws.NewErrParamRequired("MasterUsername"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/CreateDBClusterResult
type CreateDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about a DB cluster.
	DBCluster *DBCluster `json:"rds:CreateDBClusterOutput:DBCluster" type:"structure"`
}

// String returns the string representation
func (s CreateDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBCluster = "CreateDBCluster"

// CreateDBClusterRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Creates a new Amazon DocumentDB DB cluster.
//
//    // Example sending a request using CreateDBClusterRequest.
//    req := client.CreateDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/CreateDBCluster
func (c *Client) CreateDBClusterRequest(input *CreateDBClusterInput) CreateDBClusterRequest {
	op := &aws.Operation{
		Name:       opCreateDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBClusterInput{}
	}

	req := c.newRequest(op, input, &CreateDBClusterOutput{})
	return CreateDBClusterRequest{Request: req, Input: input, Copy: c.CreateDBClusterRequest}
}

// CreateDBClusterRequest is the request type for the
// CreateDBCluster API operation.
type CreateDBClusterRequest struct {
	*aws.Request
	Input *CreateDBClusterInput
	Copy  func(*CreateDBClusterInput) CreateDBClusterRequest
}

// Send marshals and sends the CreateDBCluster API request.
func (r CreateDBClusterRequest) Send(ctx context.Context) (*CreateDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBClusterResponse{
		CreateDBClusterOutput: r.Request.Data.(*CreateDBClusterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBClusterResponse is the response type for the
// CreateDBCluster API operation.
type CreateDBClusterResponse struct {
	*CreateDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBCluster request.
func (r *CreateDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
