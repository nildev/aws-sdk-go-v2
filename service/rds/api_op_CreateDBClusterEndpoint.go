// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBClusterEndpointMessage
type CreateDBClusterEndpointInput struct {
	_ struct{} `type:"structure"`

	// The identifier to use for the new endpoint. This parameter is stored as a
	// lowercase string.
	//
	// DBClusterEndpointIdentifier is a required field
	DBClusterEndpointIdentifier *string `type:"string" required:"true"`

	// The DB cluster identifier of the DB cluster associated with the endpoint.
	// This parameter is stored as a lowercase string.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The type of the endpoint. One of: READER, ANY.
	//
	// EndpointType is a required field
	EndpointType *string `type:"string" required:"true"`

	// List of DB instance identifiers that aren't part of the custom endpoint group.
	// All other eligible instances are reachable through the custom endpoint. Only
	// relevant if the list of static members is empty.
	ExcludedMembers []string `type:"list"`

	// List of DB instance identifiers that are part of the custom endpoint group.
	StaticMembers []string `type:"list"`
}

// String returns the string representation
func (s CreateDBClusterEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBClusterEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBClusterEndpointInput"}

	if s.DBClusterEndpointIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterEndpointIdentifier"))
	}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.EndpointType == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This data type represents the information you need to connect to an Amazon
// Aurora DB cluster. This data type is used as a response element in the following
// actions:
//
//    * CreateDBClusterEndpoint
//
//    * DescribeDBClusterEndpoints
//
//    * ModifyDBClusterEndpoint
//
//    * DeleteDBClusterEndpoint
//
// For the data structure that represents Amazon RDS DB instance endpoints,
// see Endpoint.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DBClusterEndpoint
type CreateDBClusterEndpointOutput struct {
	_ struct{} `type:"structure"`

	// The type associated with a custom endpoint. One of: READER, ANY.
	CustomEndpointType *string `json:"rds:CreateDBClusterEndpointOutput:CustomEndpointType" type:"string"`

	// The Amazon Resource Name (ARN) for the endpoint.
	DBClusterEndpointArn *string `json:"rds:CreateDBClusterEndpointOutput:DBClusterEndpointArn" type:"string"`

	// The identifier associated with the endpoint. This parameter is stored as
	// a lowercase string.
	DBClusterEndpointIdentifier *string `json:"rds:CreateDBClusterEndpointOutput:DBClusterEndpointIdentifier" type:"string"`

	// A unique system-generated identifier for an endpoint. It remains the same
	// for the whole life of the endpoint.
	DBClusterEndpointResourceIdentifier *string `json:"rds:CreateDBClusterEndpointOutput:DBClusterEndpointResourceIdentifier" type:"string"`

	// The DB cluster identifier of the DB cluster associated with the endpoint.
	// This parameter is stored as a lowercase string.
	DBClusterIdentifier *string `json:"rds:CreateDBClusterEndpointOutput:DBClusterIdentifier" type:"string"`

	// The DNS address of the endpoint.
	Endpoint *string `json:"rds:CreateDBClusterEndpointOutput:Endpoint" type:"string"`

	// The type of the endpoint. One of: READER, WRITER, CUSTOM.
	EndpointType *string `json:"rds:CreateDBClusterEndpointOutput:EndpointType" type:"string"`

	// List of DB instance identifiers that aren't part of the custom endpoint group.
	// All other eligible instances are reachable through the custom endpoint. Only
	// relevant if the list of static members is empty.
	ExcludedMembers []string `json:"rds:CreateDBClusterEndpointOutput:ExcludedMembers" type:"list"`

	// List of DB instance identifiers that are part of the custom endpoint group.
	StaticMembers []string `json:"rds:CreateDBClusterEndpointOutput:StaticMembers" type:"list"`

	// The current status of the endpoint. One of: creating, available, deleting,
	// modifying.
	Status *string `json:"rds:CreateDBClusterEndpointOutput:Status" type:"string"`
}

// String returns the string representation
func (s CreateDBClusterEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBClusterEndpoint = "CreateDBClusterEndpoint"

// CreateDBClusterEndpointRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new custom endpoint and associates it with an Amazon Aurora DB
// cluster.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using CreateDBClusterEndpointRequest.
//    req := client.CreateDBClusterEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBClusterEndpoint
func (c *Client) CreateDBClusterEndpointRequest(input *CreateDBClusterEndpointInput) CreateDBClusterEndpointRequest {
	op := &aws.Operation{
		Name:       opCreateDBClusterEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBClusterEndpointInput{}
	}

	req := c.newRequest(op, input, &CreateDBClusterEndpointOutput{})
	return CreateDBClusterEndpointRequest{Request: req, Input: input, Copy: c.CreateDBClusterEndpointRequest}
}

// CreateDBClusterEndpointRequest is the request type for the
// CreateDBClusterEndpoint API operation.
type CreateDBClusterEndpointRequest struct {
	*aws.Request
	Input *CreateDBClusterEndpointInput
	Copy  func(*CreateDBClusterEndpointInput) CreateDBClusterEndpointRequest
}

// Send marshals and sends the CreateDBClusterEndpoint API request.
func (r CreateDBClusterEndpointRequest) Send(ctx context.Context) (*CreateDBClusterEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBClusterEndpointResponse{
		CreateDBClusterEndpointOutput: r.Request.Data.(*CreateDBClusterEndpointOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBClusterEndpointResponse is the response type for the
// CreateDBClusterEndpoint API operation.
type CreateDBClusterEndpointResponse struct {
	*CreateDBClusterEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBClusterEndpoint request.
func (r *CreateDBClusterEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
