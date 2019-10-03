// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBClusterParameterGroupsMessage
type DescribeDBClusterParameterGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific DB cluster parameter group to return details for.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBClusterParameterGroup.
	DBClusterParameterGroupName *string `type:"string"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeDBClusterParameterGroups
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDBClusterParameterGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBClusterParameterGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBClusterParameterGroupsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DBClusterParameterGroupsMessage
type DescribeDBClusterParameterGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of DB cluster parameter groups.
	DBClusterParameterGroups []DBClusterParameterGroup `json:"rds:DescribeDBClusterParameterGroupsOutput:DBClusterParameterGroups" locationNameList:"DBClusterParameterGroup" type:"list"`

	// An optional pagination token provided by a previous DescribeDBClusterParameterGroups
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `json:"rds:DescribeDBClusterParameterGroupsOutput:Marker" type:"string"`
}

// String returns the string representation
func (s DescribeDBClusterParameterGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBClusterParameterGroups = "DescribeDBClusterParameterGroups"

// DescribeDBClusterParameterGroupsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns a list of DBClusterParameterGroup descriptions. If a DBClusterParameterGroupName
// parameter is specified, the list will contain only the description of the
// specified DB cluster parameter group.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using DescribeDBClusterParameterGroupsRequest.
//    req := client.DescribeDBClusterParameterGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBClusterParameterGroups
func (c *Client) DescribeDBClusterParameterGroupsRequest(input *DescribeDBClusterParameterGroupsInput) DescribeDBClusterParameterGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBClusterParameterGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBClusterParameterGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBClusterParameterGroupsOutput{})
	return DescribeDBClusterParameterGroupsRequest{Request: req, Input: input, Copy: c.DescribeDBClusterParameterGroupsRequest}
}

// DescribeDBClusterParameterGroupsRequest is the request type for the
// DescribeDBClusterParameterGroups API operation.
type DescribeDBClusterParameterGroupsRequest struct {
	*aws.Request
	Input *DescribeDBClusterParameterGroupsInput
	Copy  func(*DescribeDBClusterParameterGroupsInput) DescribeDBClusterParameterGroupsRequest
}

// Send marshals and sends the DescribeDBClusterParameterGroups API request.
func (r DescribeDBClusterParameterGroupsRequest) Send(ctx context.Context) (*DescribeDBClusterParameterGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBClusterParameterGroupsResponse{
		DescribeDBClusterParameterGroupsOutput: r.Request.Data.(*DescribeDBClusterParameterGroupsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDBClusterParameterGroupsResponse is the response type for the
// DescribeDBClusterParameterGroups API operation.
type DescribeDBClusterParameterGroupsResponse struct {
	*DescribeDBClusterParameterGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBClusterParameterGroups request.
func (r *DescribeDBClusterParameterGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
