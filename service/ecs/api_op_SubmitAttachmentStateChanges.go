// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/SubmitAttachmentStateChangesRequest
type SubmitAttachmentStateChangesInput struct {
	_ struct{} `type:"structure"`

	// Any attachments associated with the state change request.
	//
	// Attachments is a required field
	Attachments []AttachmentStateChange `locationName:"attachments" type:"list" required:"true"`

	// The short name or full ARN of the cluster that hosts the container instance
	// the attachment belongs to.
	Cluster *string `locationName:"cluster" type:"string"`
}

// String returns the string representation
func (s SubmitAttachmentStateChangesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubmitAttachmentStateChangesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SubmitAttachmentStateChangesInput"}

	if s.Attachments == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attachments"))
	}
	if s.Attachments != nil {
		for i, v := range s.Attachments {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attachments", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/SubmitAttachmentStateChangesResponse
type SubmitAttachmentStateChangesOutput struct {
	_ struct{} `type:"structure"`

	// Acknowledgement of the state change.
	Acknowledgment *string `json:"ecs:SubmitAttachmentStateChangesOutput:Acknowledgment" locationName:"acknowledgment" type:"string"`
}

// String returns the string representation
func (s SubmitAttachmentStateChangesOutput) String() string {
	return awsutil.Prettify(s)
}

const opSubmitAttachmentStateChanges = "SubmitAttachmentStateChanges"

// SubmitAttachmentStateChangesRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
//
// This action is only used by the Amazon ECS agent, and it is not intended
// for use outside of the agent.
//
// Sent to acknowledge that an attachment changed states.
//
//    // Example sending a request using SubmitAttachmentStateChangesRequest.
//    req := client.SubmitAttachmentStateChangesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/SubmitAttachmentStateChanges
func (c *Client) SubmitAttachmentStateChangesRequest(input *SubmitAttachmentStateChangesInput) SubmitAttachmentStateChangesRequest {
	op := &aws.Operation{
		Name:       opSubmitAttachmentStateChanges,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SubmitAttachmentStateChangesInput{}
	}

	req := c.newRequest(op, input, &SubmitAttachmentStateChangesOutput{})
	return SubmitAttachmentStateChangesRequest{Request: req, Input: input, Copy: c.SubmitAttachmentStateChangesRequest}
}

// SubmitAttachmentStateChangesRequest is the request type for the
// SubmitAttachmentStateChanges API operation.
type SubmitAttachmentStateChangesRequest struct {
	*aws.Request
	Input *SubmitAttachmentStateChangesInput
	Copy  func(*SubmitAttachmentStateChangesInput) SubmitAttachmentStateChangesRequest
}

// Send marshals and sends the SubmitAttachmentStateChanges API request.
func (r SubmitAttachmentStateChangesRequest) Send(ctx context.Context) (*SubmitAttachmentStateChangesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SubmitAttachmentStateChangesResponse{
		SubmitAttachmentStateChangesOutput: r.Request.Data.(*SubmitAttachmentStateChangesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SubmitAttachmentStateChangesResponse is the response type for the
// SubmitAttachmentStateChanges API operation.
type SubmitAttachmentStateChangesResponse struct {
	*SubmitAttachmentStateChangesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SubmitAttachmentStateChanges request.
func (r *SubmitAttachmentStateChangesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
