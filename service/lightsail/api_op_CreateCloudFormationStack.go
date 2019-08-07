// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateCloudFormationStackRequest
type CreateCloudFormationStackInput struct {
	_ struct{} `type:"structure"`

	// An array of parameters that will be used to create the new Amazon EC2 instance.
	// You can only pass one instance entry at a time in this array. You will get
	// an invalid parameter error if you pass more than one instance entry in this
	// array.
	//
	// Instances is a required field
	Instances []InstanceEntry `locationName:"instances" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateCloudFormationStackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCloudFormationStackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCloudFormationStackInput"}

	if s.Instances == nil {
		invalidParams.Add(aws.NewErrParamRequired("Instances"))
	}
	if s.Instances != nil {
		for i, v := range s.Instances {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Instances", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateCloudFormationStackResult
type CreateCloudFormationStackOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects describing the API operation.
	Operations []Operation `json:"lightsail:CreateCloudFormationStackOutput:Operations" locationName:"operations" type:"list"`
}

// String returns the string representation
func (s CreateCloudFormationStackOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCloudFormationStack = "CreateCloudFormationStack"

// CreateCloudFormationStackRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Creates an AWS CloudFormation stack, which creates a new Amazon EC2 instance
// from an exported Amazon Lightsail snapshot. This operation results in a CloudFormation
// stack record that can be used to track the AWS CloudFormation stack created.
// Use the get cloud formation stack records operation to get a list of the
// CloudFormation stacks created.
//
// Wait until after your new Amazon EC2 instance is created before running the
// create cloud formation stack operation again with the same export snapshot
// record.
//
//    // Example sending a request using CreateCloudFormationStackRequest.
//    req := client.CreateCloudFormationStackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateCloudFormationStack
func (c *Client) CreateCloudFormationStackRequest(input *CreateCloudFormationStackInput) CreateCloudFormationStackRequest {
	op := &aws.Operation{
		Name:       opCreateCloudFormationStack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCloudFormationStackInput{}
	}

	req := c.newRequest(op, input, &CreateCloudFormationStackOutput{})
	return CreateCloudFormationStackRequest{Request: req, Input: input, Copy: c.CreateCloudFormationStackRequest}
}

// CreateCloudFormationStackRequest is the request type for the
// CreateCloudFormationStack API operation.
type CreateCloudFormationStackRequest struct {
	*aws.Request
	Input *CreateCloudFormationStackInput
	Copy  func(*CreateCloudFormationStackInput) CreateCloudFormationStackRequest
}

// Send marshals and sends the CreateCloudFormationStack API request.
func (r CreateCloudFormationStackRequest) Send(ctx context.Context) (*CreateCloudFormationStackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCloudFormationStackResponse{
		CreateCloudFormationStackOutput: r.Request.Data.(*CreateCloudFormationStackOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCloudFormationStackResponse is the response type for the
// CreateCloudFormationStack API operation.
type CreateCloudFormationStackResponse struct {
	*CreateCloudFormationStackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCloudFormationStack request.
func (r *CreateCloudFormationStackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
