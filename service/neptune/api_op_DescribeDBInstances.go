// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeDBInstancesMessage
type DescribeDBInstancesInput struct {
	_ struct{} `type:"structure"`

	// The user-supplied instance identifier. If this parameter is specified, information
	// from only the specific DB instance is returned. This parameter isn't case-sensitive.
	//
	// Constraints:
	//
	//    * If supplied, must match the identifier of an existing DBInstance.
	DBInstanceIdentifier *string `type:"string"`

	// A filter that specifies one or more DB instances to describe.
	//
	// Supported filters:
	//
	//    * db-cluster-id - Accepts DB cluster identifiers and DB cluster Amazon
	//    Resource Names (ARNs). The results list will only include information
	//    about the DB instances associated with the DB clusters identified by these
	//    ARNs.
	//
	//    * db-instance-id - Accepts DB instance identifiers and DB instance Amazon
	//    Resource Names (ARNs). The results list will only include information
	//    about the DB instances identified by these ARNs.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeDBInstances request.
	// If this parameter is specified, the response includes only records beyond
	// the marker, up to the value specified by MaxRecords.
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
func (s DescribeDBInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBInstancesInput"}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DBInstanceMessage
type DescribeDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	// A list of DBInstance instances.
	DBInstances []DBInstance `json:"rds:DescribeDBInstancesOutput:DBInstances" locationNameList:"DBInstance" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords .
	Marker *string `json:"rds:DescribeDBInstancesOutput:Marker" type:"string"`
}

// String returns the string representation
func (s DescribeDBInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBInstances = "DescribeDBInstances"

// DescribeDBInstancesRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Returns information about provisioned instances. This API supports pagination.
//
//    // Example sending a request using DescribeDBInstancesRequest.
//    req := client.DescribeDBInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeDBInstances
func (c *Client) DescribeDBInstancesRequest(input *DescribeDBInstancesInput) DescribeDBInstancesRequest {
	op := &aws.Operation{
		Name:       opDescribeDBInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDBInstancesInput{}
	}

	req := c.newRequest(op, input, &DescribeDBInstancesOutput{})
	return DescribeDBInstancesRequest{Request: req, Input: input, Copy: c.DescribeDBInstancesRequest}
}

// DescribeDBInstancesRequest is the request type for the
// DescribeDBInstances API operation.
type DescribeDBInstancesRequest struct {
	*aws.Request
	Input *DescribeDBInstancesInput
	Copy  func(*DescribeDBInstancesInput) DescribeDBInstancesRequest
}

// Send marshals and sends the DescribeDBInstances API request.
func (r DescribeDBInstancesRequest) Send(ctx context.Context) (*DescribeDBInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBInstancesResponse{
		DescribeDBInstancesOutput: r.Request.Data.(*DescribeDBInstancesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBInstancesRequestPaginator returns a paginator for DescribeDBInstances.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBInstancesRequest(input)
//   p := neptune.NewDescribeDBInstancesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBInstancesPaginator(req DescribeDBInstancesRequest) DescribeDBInstancesPaginator {
	return DescribeDBInstancesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBInstancesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeDBInstancesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBInstancesPaginator struct {
	aws.Pager
}

func (p *DescribeDBInstancesPaginator) CurrentPage() *DescribeDBInstancesOutput {
	return p.Pager.CurrentPage().(*DescribeDBInstancesOutput)
}

// DescribeDBInstancesResponse is the response type for the
// DescribeDBInstances API operation.
type DescribeDBInstancesResponse struct {
	*DescribeDBInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBInstances request.
func (r *DescribeDBInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
