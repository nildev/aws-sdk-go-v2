// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/PutInventoryRequest
type PutInventoryInput struct {
	_ struct{} `type:"structure"`

	// One or more instance IDs where you want to add or update inventory items.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The inventory items that you want to add or update on instances.
	//
	// Items is a required field
	Items []InventoryItem `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s PutInventoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutInventoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutInventoryInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.Items == nil {
		invalidParams.Add(aws.NewErrParamRequired("Items"))
	}
	if s.Items != nil && len(s.Items) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Items", 1))
	}
	if s.Items != nil {
		for i, v := range s.Items {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Items", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/PutInventoryResult
type PutInventoryOutput struct {
	_ struct{} `type:"structure"`

	// Information about the request.
	Message *string `json:"ssm:PutInventoryOutput:Message" type:"string"`
}

// String returns the string representation
func (s PutInventoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutInventory = "PutInventory"

// PutInventoryRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Bulk update custom inventory items on one more instance. The request adds
// an inventory item, if it doesn't already exist, or updates an inventory item,
// if it does exist.
//
//    // Example sending a request using PutInventoryRequest.
//    req := client.PutInventoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/PutInventory
func (c *Client) PutInventoryRequest(input *PutInventoryInput) PutInventoryRequest {
	op := &aws.Operation{
		Name:       opPutInventory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutInventoryInput{}
	}

	req := c.newRequest(op, input, &PutInventoryOutput{})
	return PutInventoryRequest{Request: req, Input: input, Copy: c.PutInventoryRequest}
}

// PutInventoryRequest is the request type for the
// PutInventory API operation.
type PutInventoryRequest struct {
	*aws.Request
	Input *PutInventoryInput
	Copy  func(*PutInventoryInput) PutInventoryRequest
}

// Send marshals and sends the PutInventory API request.
func (r PutInventoryRequest) Send(ctx context.Context) (*PutInventoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutInventoryResponse{
		PutInventoryOutput: r.Request.Data.(*PutInventoryOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutInventoryResponse is the response type for the
// PutInventory API operation.
type PutInventoryResponse struct {
	*PutInventoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutInventory request.
func (r *PutInventoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
