// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ConfirmProductInstanceRequest
type ConfirmProductInstanceInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The product code. This must be a product code that you own.
	//
	// ProductCode is a required field
	ProductCode *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ConfirmProductInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfirmProductInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfirmProductInstanceInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.ProductCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductCode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ConfirmProductInstanceResult
type ConfirmProductInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID of the instance owner. This is only present if the product
	// code is attached to the instance.
	OwnerId *string `json:"ec2:ConfirmProductInstanceOutput:OwnerId" locationName:"ownerId" type:"string"`

	// The return value of the request. Returns true if the specified product code
	// is owned by the requester and associated with the specified instance.
	Return *bool `json:"ec2:ConfirmProductInstanceOutput:Return" locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ConfirmProductInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opConfirmProductInstance = "ConfirmProductInstance"

// ConfirmProductInstanceRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Determines whether a product code is associated with an instance. This action
// can only be used by the owner of the product code. It is useful when a product
// code owner must verify whether another user's instance is eligible for support.
//
//    // Example sending a request using ConfirmProductInstanceRequest.
//    req := client.ConfirmProductInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ConfirmProductInstance
func (c *Client) ConfirmProductInstanceRequest(input *ConfirmProductInstanceInput) ConfirmProductInstanceRequest {
	op := &aws.Operation{
		Name:       opConfirmProductInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ConfirmProductInstanceInput{}
	}

	req := c.newRequest(op, input, &ConfirmProductInstanceOutput{})
	return ConfirmProductInstanceRequest{Request: req, Input: input, Copy: c.ConfirmProductInstanceRequest}
}

// ConfirmProductInstanceRequest is the request type for the
// ConfirmProductInstance API operation.
type ConfirmProductInstanceRequest struct {
	*aws.Request
	Input *ConfirmProductInstanceInput
	Copy  func(*ConfirmProductInstanceInput) ConfirmProductInstanceRequest
}

// Send marshals and sends the ConfirmProductInstance API request.
func (r ConfirmProductInstanceRequest) Send(ctx context.Context) (*ConfirmProductInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ConfirmProductInstanceResponse{
		ConfirmProductInstanceOutput: r.Request.Data.(*ConfirmProductInstanceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ConfirmProductInstanceResponse is the response type for the
// ConfirmProductInstance API operation.
type ConfirmProductInstanceResponse struct {
	*ConfirmProductInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ConfirmProductInstance request.
func (r *ConfirmProductInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
