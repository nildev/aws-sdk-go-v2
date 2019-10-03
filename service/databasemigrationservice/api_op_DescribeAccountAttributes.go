// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeAccountAttributesMessage
type DescribeAccountAttributesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeAccountAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeAccountAttributesResponse
type DescribeAccountAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Account quota information.
	AccountQuotas []AccountQuota `json:"dms:DescribeAccountAttributesOutput:AccountQuotas" type:"list"`

	// A unique AWS DMS identifier for an account in a particular AWS Region. The
	// value of this identifier has the following format: c99999999999. DMS uses
	// this identifier to name artifacts. For example, DMS uses this identifier
	// to name the default Amazon S3 bucket for storing task assessment reports
	// in a given AWS Region. The format of this S3 bucket name is the following:
	// dms-AccountNumber-UniqueAccountIdentifier. Here is an example name for this
	// default S3 bucket: dms-111122223333-c44445555666.
	//
	// AWS DMS supports UniqueAccountIdentifier in versions 3.1.4 and later.
	UniqueAccountIdentifier *string `json:"dms:DescribeAccountAttributesOutput:UniqueAccountIdentifier" type:"string"`
}

// String returns the string representation
func (s DescribeAccountAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAccountAttributes = "DescribeAccountAttributes"

// DescribeAccountAttributesRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Lists all of the AWS DMS attributes for a customer account. These attributes
// include AWS DMS quotas for the account and a unique account identifier in
// a particular DMS region. DMS quotas include a list of resource quotas supported
// by the account, such as the number of replication instances allowed. The
// description for each resource quota, includes the quota name, current usage
// toward that quota, and the quota's maximum value. DMS uses the unique account
// identifier to name each artifact used by DMS in the given region.
//
// This command does not take any parameters.
//
//    // Example sending a request using DescribeAccountAttributesRequest.
//    req := client.DescribeAccountAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeAccountAttributes
func (c *Client) DescribeAccountAttributesRequest(input *DescribeAccountAttributesInput) DescribeAccountAttributesRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAccountAttributesInput{}
	}

	req := c.newRequest(op, input, &DescribeAccountAttributesOutput{})
	return DescribeAccountAttributesRequest{Request: req, Input: input, Copy: c.DescribeAccountAttributesRequest}
}

// DescribeAccountAttributesRequest is the request type for the
// DescribeAccountAttributes API operation.
type DescribeAccountAttributesRequest struct {
	*aws.Request
	Input *DescribeAccountAttributesInput
	Copy  func(*DescribeAccountAttributesInput) DescribeAccountAttributesRequest
}

// Send marshals and sends the DescribeAccountAttributes API request.
func (r DescribeAccountAttributesRequest) Send(ctx context.Context) (*DescribeAccountAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountAttributesResponse{
		DescribeAccountAttributesOutput: r.Request.Data.(*DescribeAccountAttributesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountAttributesResponse is the response type for the
// DescribeAccountAttributes API operation.
type DescribeAccountAttributesResponse struct {
	*DescribeAccountAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountAttributes request.
func (r *DescribeAccountAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
