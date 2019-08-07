// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeInstanceAssociationsStatusRequest
type DescribeInstanceAssociationsStatusInput struct {
	_ struct{} `type:"structure"`

	// The instance IDs for which you want association status information.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceAssociationsStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstanceAssociationsStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInstanceAssociationsStatusInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeInstanceAssociationsStatusResult
type DescribeInstanceAssociationsStatusOutput struct {
	_ struct{} `type:"structure"`

	// Status information about the association.
	InstanceAssociationStatusInfos []InstanceAssociationStatusInfo `json:"ssm:DescribeInstanceAssociationsStatusOutput:InstanceAssociationStatusInfos" type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `json:"ssm:DescribeInstanceAssociationsStatusOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s DescribeInstanceAssociationsStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInstanceAssociationsStatus = "DescribeInstanceAssociationsStatus"

// DescribeInstanceAssociationsStatusRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// The status of the associations for the instance(s).
//
//    // Example sending a request using DescribeInstanceAssociationsStatusRequest.
//    req := client.DescribeInstanceAssociationsStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeInstanceAssociationsStatus
func (c *Client) DescribeInstanceAssociationsStatusRequest(input *DescribeInstanceAssociationsStatusInput) DescribeInstanceAssociationsStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeInstanceAssociationsStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstanceAssociationsStatusInput{}
	}

	req := c.newRequest(op, input, &DescribeInstanceAssociationsStatusOutput{})
	return DescribeInstanceAssociationsStatusRequest{Request: req, Input: input, Copy: c.DescribeInstanceAssociationsStatusRequest}
}

// DescribeInstanceAssociationsStatusRequest is the request type for the
// DescribeInstanceAssociationsStatus API operation.
type DescribeInstanceAssociationsStatusRequest struct {
	*aws.Request
	Input *DescribeInstanceAssociationsStatusInput
	Copy  func(*DescribeInstanceAssociationsStatusInput) DescribeInstanceAssociationsStatusRequest
}

// Send marshals and sends the DescribeInstanceAssociationsStatus API request.
func (r DescribeInstanceAssociationsStatusRequest) Send(ctx context.Context) (*DescribeInstanceAssociationsStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstanceAssociationsStatusResponse{
		DescribeInstanceAssociationsStatusOutput: r.Request.Data.(*DescribeInstanceAssociationsStatusOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInstanceAssociationsStatusResponse is the response type for the
// DescribeInstanceAssociationsStatus API operation.
type DescribeInstanceAssociationsStatusResponse struct {
	*DescribeInstanceAssociationsStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstanceAssociationsStatus request.
func (r *DescribeInstanceAssociationsStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
