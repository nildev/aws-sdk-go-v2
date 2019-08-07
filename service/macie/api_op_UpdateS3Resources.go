// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie-2017-12-19/UpdateS3ResourcesRequest
type UpdateS3ResourcesInput struct {
	_ struct{} `type:"structure"`

	// The AWS ID of the Amazon Macie member account whose S3 resources' classification
	// types you want to update.
	MemberAccountId *string `locationName:"memberAccountId" type:"string"`

	// The S3 resources whose classification types you want to update.
	//
	// S3ResourcesUpdate is a required field
	S3ResourcesUpdate []S3ResourceClassificationUpdate `locationName:"s3ResourcesUpdate" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateS3ResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateS3ResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateS3ResourcesInput"}

	if s.S3ResourcesUpdate == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3ResourcesUpdate"))
	}
	if s.S3ResourcesUpdate != nil {
		for i, v := range s.S3ResourcesUpdate {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "S3ResourcesUpdate", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie-2017-12-19/UpdateS3ResourcesResult
type UpdateS3ResourcesOutput struct {
	_ struct{} `type:"structure"`

	// The S3 resources whose classification types can't be updated. An error code
	// and an error message are provided for each failed item.
	FailedS3Resources []FailedS3Resource `json:"macie:UpdateS3ResourcesOutput:FailedS3Resources" locationName:"failedS3Resources" type:"list"`
}

// String returns the string representation
func (s UpdateS3ResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateS3Resources = "UpdateS3Resources"

// UpdateS3ResourcesRequest returns a request value for making API operation for
// Amazon Macie.
//
// Updates the classification types for the specified S3 resources. If memberAccountId
// isn't specified, the action updates the classification types of the S3 resources
// associated with Amazon Macie for the current master account. If memberAccountId
// is specified, the action updates the classification types of the S3 resources
// associated with Amazon Macie for the specified member account.
//
//    // Example sending a request using UpdateS3ResourcesRequest.
//    req := client.UpdateS3ResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie-2017-12-19/UpdateS3Resources
func (c *Client) UpdateS3ResourcesRequest(input *UpdateS3ResourcesInput) UpdateS3ResourcesRequest {
	op := &aws.Operation{
		Name:       opUpdateS3Resources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateS3ResourcesInput{}
	}

	req := c.newRequest(op, input, &UpdateS3ResourcesOutput{})
	return UpdateS3ResourcesRequest{Request: req, Input: input, Copy: c.UpdateS3ResourcesRequest}
}

// UpdateS3ResourcesRequest is the request type for the
// UpdateS3Resources API operation.
type UpdateS3ResourcesRequest struct {
	*aws.Request
	Input *UpdateS3ResourcesInput
	Copy  func(*UpdateS3ResourcesInput) UpdateS3ResourcesRequest
}

// Send marshals and sends the UpdateS3Resources API request.
func (r UpdateS3ResourcesRequest) Send(ctx context.Context) (*UpdateS3ResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateS3ResourcesResponse{
		UpdateS3ResourcesOutput: r.Request.Data.(*UpdateS3ResourcesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateS3ResourcesResponse is the response type for the
// UpdateS3Resources API operation.
type UpdateS3ResourcesResponse struct {
	*UpdateS3ResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateS3Resources request.
func (r *UpdateS3ResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
