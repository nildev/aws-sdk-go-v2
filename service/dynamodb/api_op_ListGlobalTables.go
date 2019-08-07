// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/ListGlobalTablesInput
type ListGlobalTablesInput struct {
	_ struct{} `type:"structure"`

	// The first global table name that this operation will evaluate.
	ExclusiveStartGlobalTableName *string `min:"3" type:"string"`

	// The maximum number of table names to return.
	Limit *int64 `min:"1" type:"integer"`

	// Lists the global tables in a specific Region.
	RegionName *string `type:"string"`
}

// String returns the string representation
func (s ListGlobalTablesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGlobalTablesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGlobalTablesInput"}
	if s.ExclusiveStartGlobalTableName != nil && len(*s.ExclusiveStartGlobalTableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("ExclusiveStartGlobalTableName", 3))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/ListGlobalTablesOutput
type ListGlobalTablesOutput struct {
	_ struct{} `type:"structure"`

	// List of global table names.
	GlobalTables []GlobalTable `json:"dynamodb:ListGlobalTablesOutput:GlobalTables" type:"list"`

	// Last evaluated global table name.
	LastEvaluatedGlobalTableName *string `json:"dynamodb:ListGlobalTablesOutput:LastEvaluatedGlobalTableName" min:"3" type:"string"`
}

// String returns the string representation
func (s ListGlobalTablesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListGlobalTables = "ListGlobalTables"

// ListGlobalTablesRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Lists all global tables that have a replica in the specified Region.
//
//    // Example sending a request using ListGlobalTablesRequest.
//    req := client.ListGlobalTablesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/ListGlobalTables
func (c *Client) ListGlobalTablesRequest(input *ListGlobalTablesInput) ListGlobalTablesRequest {
	op := &aws.Operation{
		Name:       opListGlobalTables,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListGlobalTablesInput{}
	}

	req := c.newRequest(op, input, &ListGlobalTablesOutput{})
	return ListGlobalTablesRequest{Request: req, Input: input, Copy: c.ListGlobalTablesRequest}
}

// ListGlobalTablesRequest is the request type for the
// ListGlobalTables API operation.
type ListGlobalTablesRequest struct {
	*aws.Request
	Input *ListGlobalTablesInput
	Copy  func(*ListGlobalTablesInput) ListGlobalTablesRequest
}

// Send marshals and sends the ListGlobalTables API request.
func (r ListGlobalTablesRequest) Send(ctx context.Context) (*ListGlobalTablesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGlobalTablesResponse{
		ListGlobalTablesOutput: r.Request.Data.(*ListGlobalTablesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListGlobalTablesResponse is the response type for the
// ListGlobalTables API operation.
type ListGlobalTablesResponse struct {
	*ListGlobalTablesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGlobalTables request.
func (r *ListGlobalTablesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
