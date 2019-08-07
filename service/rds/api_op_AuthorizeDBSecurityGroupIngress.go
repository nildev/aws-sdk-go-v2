// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/AuthorizeDBSecurityGroupIngressMessage
type AuthorizeDBSecurityGroupIngressInput struct {
	_ struct{} `type:"structure"`

	// The IP range to authorize.
	CIDRIP *string `type:"string"`

	// The name of the DB security group to add authorization to.
	//
	// DBSecurityGroupName is a required field
	DBSecurityGroupName *string `type:"string" required:"true"`

	// Id of the EC2 security group to authorize. For VPC DB security groups, EC2SecurityGroupId
	// must be provided. Otherwise, EC2SecurityGroupOwnerId and either EC2SecurityGroupName
	// or EC2SecurityGroupId must be provided.
	EC2SecurityGroupId *string `type:"string"`

	// Name of the EC2 security group to authorize. For VPC DB security groups,
	// EC2SecurityGroupId must be provided. Otherwise, EC2SecurityGroupOwnerId and
	// either EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupName *string `type:"string"`

	// AWS account number of the owner of the EC2 security group specified in the
	// EC2SecurityGroupName parameter. The AWS Access Key ID is not an acceptable
	// value. For VPC DB security groups, EC2SecurityGroupId must be provided. Otherwise,
	// EC2SecurityGroupOwnerId and either EC2SecurityGroupName or EC2SecurityGroupId
	// must be provided.
	EC2SecurityGroupOwnerId *string `type:"string"`
}

// String returns the string representation
func (s AuthorizeDBSecurityGroupIngressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AuthorizeDBSecurityGroupIngressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AuthorizeDBSecurityGroupIngressInput"}

	if s.DBSecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSecurityGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/AuthorizeDBSecurityGroupIngressResult
type AuthorizeDBSecurityGroupIngressOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details for an Amazon RDS DB security group.
	//
	// This data type is used as a response element in the DescribeDBSecurityGroups
	// action.
	DBSecurityGroup *DBSecurityGroup `json:"rds:AuthorizeDBSecurityGroupIngressOutput:DBSecurityGroup" type:"structure"`
}

// String returns the string representation
func (s AuthorizeDBSecurityGroupIngressOutput) String() string {
	return awsutil.Prettify(s)
}

const opAuthorizeDBSecurityGroupIngress = "AuthorizeDBSecurityGroupIngress"

// AuthorizeDBSecurityGroupIngressRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Enables ingress to a DBSecurityGroup using one of two forms of authorization.
// First, EC2 or VPC security groups can be added to the DBSecurityGroup if
// the application using the database is running on EC2 or VPC instances. Second,
// IP ranges are available if the application accessing your database is running
// on the Internet. Required parameters for this API are one of CIDR range,
// EC2SecurityGroupId for VPC, or (EC2SecurityGroupOwnerId and either EC2SecurityGroupName
// or EC2SecurityGroupId for non-VPC).
//
// You can't authorize ingress from an EC2 security group in one AWS Region
// to an Amazon RDS DB instance in another. You can't authorize ingress from
// a VPC security group in one VPC to an Amazon RDS DB instance in another.
//
// For an overview of CIDR ranges, go to the Wikipedia Tutorial (http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).
//
//    // Example sending a request using AuthorizeDBSecurityGroupIngressRequest.
//    req := client.AuthorizeDBSecurityGroupIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/AuthorizeDBSecurityGroupIngress
func (c *Client) AuthorizeDBSecurityGroupIngressRequest(input *AuthorizeDBSecurityGroupIngressInput) AuthorizeDBSecurityGroupIngressRequest {
	op := &aws.Operation{
		Name:       opAuthorizeDBSecurityGroupIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AuthorizeDBSecurityGroupIngressInput{}
	}

	req := c.newRequest(op, input, &AuthorizeDBSecurityGroupIngressOutput{})
	return AuthorizeDBSecurityGroupIngressRequest{Request: req, Input: input, Copy: c.AuthorizeDBSecurityGroupIngressRequest}
}

// AuthorizeDBSecurityGroupIngressRequest is the request type for the
// AuthorizeDBSecurityGroupIngress API operation.
type AuthorizeDBSecurityGroupIngressRequest struct {
	*aws.Request
	Input *AuthorizeDBSecurityGroupIngressInput
	Copy  func(*AuthorizeDBSecurityGroupIngressInput) AuthorizeDBSecurityGroupIngressRequest
}

// Send marshals and sends the AuthorizeDBSecurityGroupIngress API request.
func (r AuthorizeDBSecurityGroupIngressRequest) Send(ctx context.Context) (*AuthorizeDBSecurityGroupIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AuthorizeDBSecurityGroupIngressResponse{
		AuthorizeDBSecurityGroupIngressOutput: r.Request.Data.(*AuthorizeDBSecurityGroupIngressOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AuthorizeDBSecurityGroupIngressResponse is the response type for the
// AuthorizeDBSecurityGroupIngress API operation.
type AuthorizeDBSecurityGroupIngressResponse struct {
	*AuthorizeDBSecurityGroupIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AuthorizeDBSecurityGroupIngress request.
func (r *AuthorizeDBSecurityGroupIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
