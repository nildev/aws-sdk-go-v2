// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeOpsItemsRequest
type DescribeOpsItemsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `type:"string"`

	// One or more filters to limit the reponse.
	//
	//    * Key: CreatedTime Operations: GreaterThan, LessThan
	//
	//    * Key: LastModifiedBy Operations: Contains, Equals
	//
	//    * Key: LastModifiedTime Operations: GreaterThan, LessThan
	//
	//    * Key: Priority Operations: Equals
	//
	//    * Key: Source Operations: Contains, Equals
	//
	//    * Key: Status Operations: Equals
	//
	//    * Key: Title Operations: Contains
	//
	//    * Key: OperationalData* Operations: Equals
	//
	//    * Key: OperationalDataKey Operations: Equals
	//
	//    * Key: OperationalDataValue Operations: Equals, Contains
	//
	//    * Key: OpsItemId Operations: Equals
	//
	//    * Key: ResourceId Operations: Contains
	//
	//    * Key: AutomationId Operations: Equals
	//
	// *If you filter the response by using the OperationalData operator, specify
	// a key-value pair by using the following JSON format: {"key":"key_name","value":"a_value"}
	OpsItemFilters []OpsItemFilter `type:"list"`
}

// String returns the string representation
func (s DescribeOpsItemsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOpsItemsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeOpsItemsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.OpsItemFilters != nil {
		for i, v := range s.OpsItemFilters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "OpsItemFilters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeOpsItemsResponse
type DescribeOpsItemsOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `json:"ssm:DescribeOpsItemsOutput:NextToken" type:"string"`

	// A list of OpsItems.
	OpsItemSummaries []OpsItemSummary `json:"ssm:DescribeOpsItemsOutput:OpsItemSummaries" type:"list"`
}

// String returns the string representation
func (s DescribeOpsItemsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeOpsItems = "DescribeOpsItems"

// DescribeOpsItemsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Query a set of OpsItems. You must have permission in AWS Identity and Access
// Management (IAM) to query a list of OpsItems. For more information, see Getting
// Started with OpsCenter (http://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-getting-started.html)
// in the AWS Systems Manager User Guide.
//
// Operations engineers and IT professionals use OpsCenter to view, investigate,
// and remediate operational issues impacting the performance and health of
// their AWS resources. For more information, see AWS Systems Manager OpsCenter
// (http://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html)
// in the AWS Systems Manager User Guide.
//
//    // Example sending a request using DescribeOpsItemsRequest.
//    req := client.DescribeOpsItemsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeOpsItems
func (c *Client) DescribeOpsItemsRequest(input *DescribeOpsItemsInput) DescribeOpsItemsRequest {
	op := &aws.Operation{
		Name:       opDescribeOpsItems,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeOpsItemsInput{}
	}

	req := c.newRequest(op, input, &DescribeOpsItemsOutput{})
	return DescribeOpsItemsRequest{Request: req, Input: input, Copy: c.DescribeOpsItemsRequest}
}

// DescribeOpsItemsRequest is the request type for the
// DescribeOpsItems API operation.
type DescribeOpsItemsRequest struct {
	*aws.Request
	Input *DescribeOpsItemsInput
	Copy  func(*DescribeOpsItemsInput) DescribeOpsItemsRequest
}

// Send marshals and sends the DescribeOpsItems API request.
func (r DescribeOpsItemsRequest) Send(ctx context.Context) (*DescribeOpsItemsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOpsItemsResponse{
		DescribeOpsItemsOutput: r.Request.Data.(*DescribeOpsItemsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeOpsItemsResponse is the response type for the
// DescribeOpsItems API operation.
type DescribeOpsItemsResponse struct {
	*DescribeOpsItemsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOpsItems request.
func (r *DescribeOpsItemsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
