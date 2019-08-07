// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchCreatePartitionRequest
type BatchCreatePartitionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the catalog in which the partion is to be created. Currently, this
	// should be the AWS account ID.
	CatalogId *string `min:"1" type:"string"`

	// The name of the metadata database in which the partition is to be created.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// A list of PartitionInput structures that define the partitions to be created.
	//
	// PartitionInputList is a required field
	PartitionInputList []PartitionInput `type:"list" required:"true"`

	// The name of the metadata table in which the partition is to be created.
	//
	// TableName is a required field
	TableName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s BatchCreatePartitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchCreatePartitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchCreatePartitionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.PartitionInputList == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartitionInputList"))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 1))
	}
	if s.PartitionInputList != nil {
		for i, v := range s.PartitionInputList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PartitionInputList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchCreatePartitionResponse
type BatchCreatePartitionOutput struct {
	_ struct{} `type:"structure"`

	// Errors encountered when trying to create the requested partitions.
	Errors []PartitionError `json:"glue:BatchCreatePartitionOutput:Errors" type:"list"`
}

// String returns the string representation
func (s BatchCreatePartitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchCreatePartition = "BatchCreatePartition"

// BatchCreatePartitionRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates one or more partitions in a batch operation.
//
//    // Example sending a request using BatchCreatePartitionRequest.
//    req := client.BatchCreatePartitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchCreatePartition
func (c *Client) BatchCreatePartitionRequest(input *BatchCreatePartitionInput) BatchCreatePartitionRequest {
	op := &aws.Operation{
		Name:       opBatchCreatePartition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchCreatePartitionInput{}
	}

	req := c.newRequest(op, input, &BatchCreatePartitionOutput{})
	return BatchCreatePartitionRequest{Request: req, Input: input, Copy: c.BatchCreatePartitionRequest}
}

// BatchCreatePartitionRequest is the request type for the
// BatchCreatePartition API operation.
type BatchCreatePartitionRequest struct {
	*aws.Request
	Input *BatchCreatePartitionInput
	Copy  func(*BatchCreatePartitionInput) BatchCreatePartitionRequest
}

// Send marshals and sends the BatchCreatePartition API request.
func (r BatchCreatePartitionRequest) Send(ctx context.Context) (*BatchCreatePartitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchCreatePartitionResponse{
		BatchCreatePartitionOutput: r.Request.Data.(*BatchCreatePartitionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchCreatePartitionResponse is the response type for the
// BatchCreatePartition API operation.
type BatchCreatePartitionResponse struct {
	*BatchCreatePartitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchCreatePartition request.
func (r *BatchCreatePartitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
