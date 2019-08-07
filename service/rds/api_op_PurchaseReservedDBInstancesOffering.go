// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/PurchaseReservedDBInstancesOfferingMessage
type PurchaseReservedDBInstancesOfferingInput struct {
	_ struct{} `type:"structure"`

	// The number of instances to reserve.
	//
	// Default: 1
	DBInstanceCount *int64 `type:"integer"`

	// Customer-specified identifier to track this reservation.
	//
	// Example: myreservationID
	ReservedDBInstanceId *string `type:"string"`

	// The ID of the Reserved DB instance offering to purchase.
	//
	// Example: 438012d3-4052-4cc7-b2e3-8d3372e0e706
	//
	// ReservedDBInstancesOfferingId is a required field
	ReservedDBInstancesOfferingId *string `type:"string" required:"true"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s PurchaseReservedDBInstancesOfferingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PurchaseReservedDBInstancesOfferingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PurchaseReservedDBInstancesOfferingInput"}

	if s.ReservedDBInstancesOfferingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedDBInstancesOfferingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/PurchaseReservedDBInstancesOfferingResult
type PurchaseReservedDBInstancesOfferingOutput struct {
	_ struct{} `type:"structure"`

	// This data type is used as a response element in the DescribeReservedDBInstances
	// and PurchaseReservedDBInstancesOffering actions.
	ReservedDBInstance *ReservedDBInstance `json:"rds:PurchaseReservedDBInstancesOfferingOutput:ReservedDBInstance" type:"structure"`
}

// String returns the string representation
func (s PurchaseReservedDBInstancesOfferingOutput) String() string {
	return awsutil.Prettify(s)
}

const opPurchaseReservedDBInstancesOffering = "PurchaseReservedDBInstancesOffering"

// PurchaseReservedDBInstancesOfferingRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Purchases a reserved DB instance offering.
//
//    // Example sending a request using PurchaseReservedDBInstancesOfferingRequest.
//    req := client.PurchaseReservedDBInstancesOfferingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/PurchaseReservedDBInstancesOffering
func (c *Client) PurchaseReservedDBInstancesOfferingRequest(input *PurchaseReservedDBInstancesOfferingInput) PurchaseReservedDBInstancesOfferingRequest {
	op := &aws.Operation{
		Name:       opPurchaseReservedDBInstancesOffering,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PurchaseReservedDBInstancesOfferingInput{}
	}

	req := c.newRequest(op, input, &PurchaseReservedDBInstancesOfferingOutput{})
	return PurchaseReservedDBInstancesOfferingRequest{Request: req, Input: input, Copy: c.PurchaseReservedDBInstancesOfferingRequest}
}

// PurchaseReservedDBInstancesOfferingRequest is the request type for the
// PurchaseReservedDBInstancesOffering API operation.
type PurchaseReservedDBInstancesOfferingRequest struct {
	*aws.Request
	Input *PurchaseReservedDBInstancesOfferingInput
	Copy  func(*PurchaseReservedDBInstancesOfferingInput) PurchaseReservedDBInstancesOfferingRequest
}

// Send marshals and sends the PurchaseReservedDBInstancesOffering API request.
func (r PurchaseReservedDBInstancesOfferingRequest) Send(ctx context.Context) (*PurchaseReservedDBInstancesOfferingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseReservedDBInstancesOfferingResponse{
		PurchaseReservedDBInstancesOfferingOutput: r.Request.Data.(*PurchaseReservedDBInstancesOfferingOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseReservedDBInstancesOfferingResponse is the response type for the
// PurchaseReservedDBInstancesOffering API operation.
type PurchaseReservedDBInstancesOfferingResponse struct {
	*PurchaseReservedDBInstancesOfferingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseReservedDBInstancesOffering request.
func (r *PurchaseReservedDBInstancesOfferingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
