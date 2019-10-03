// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/DescribeExportTasksRequest
type DescribeExportTasksInput struct {
	_ struct{} `type:"structure"`

	// One or more unique identifiers used to query the status of an export request.
	ExportIds []string `locationName:"exportIds" type:"list"`

	// One or more filters.
	//
	//    * AgentId - ID of the agent whose collected data will be exported
	Filters []ExportFilter `locationName:"filters" type:"list"`

	// The maximum number of volume results returned by DescribeExportTasks in paginated
	// output. When this parameter is used, DescribeExportTasks only returns maxResults
	// results in a single page along with a nextToken response element.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated DescribeExportTasks
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeExportTasksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeExportTasksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeExportTasksInput"}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/DescribeExportTasksResponse
type DescribeExportTasksOutput struct {
	_ struct{} `type:"structure"`

	// Contains one or more sets of export request details. When the status of a
	// request is SUCCEEDED, the response includes a URL for an Amazon S3 bucket
	// where you can view the data in a CSV file.
	ExportsInfo []ExportInfo `json:"discovery:DescribeExportTasksOutput:ExportsInfo" locationName:"exportsInfo" type:"list"`

	// The nextToken value to include in a future DescribeExportTasks request. When
	// the results of a DescribeExportTasks request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `json:"discovery:DescribeExportTasksOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeExportTasksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeExportTasks = "DescribeExportTasks"

// DescribeExportTasksRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Retrieve status of one or more export tasks. You can retrieve the status
// of up to 100 export tasks.
//
//    // Example sending a request using DescribeExportTasksRequest.
//    req := client.DescribeExportTasksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/DescribeExportTasks
func (c *Client) DescribeExportTasksRequest(input *DescribeExportTasksInput) DescribeExportTasksRequest {
	op := &aws.Operation{
		Name:       opDescribeExportTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeExportTasksInput{}
	}

	req := c.newRequest(op, input, &DescribeExportTasksOutput{})
	return DescribeExportTasksRequest{Request: req, Input: input, Copy: c.DescribeExportTasksRequest}
}

// DescribeExportTasksRequest is the request type for the
// DescribeExportTasks API operation.
type DescribeExportTasksRequest struct {
	*aws.Request
	Input *DescribeExportTasksInput
	Copy  func(*DescribeExportTasksInput) DescribeExportTasksRequest
}

// Send marshals and sends the DescribeExportTasks API request.
func (r DescribeExportTasksRequest) Send(ctx context.Context) (*DescribeExportTasksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeExportTasksResponse{
		DescribeExportTasksOutput: r.Request.Data.(*DescribeExportTasksOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeExportTasksResponse is the response type for the
// DescribeExportTasks API operation.
type DescribeExportTasksResponse struct {
	*DescribeExportTasksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeExportTasks request.
func (r *DescribeExportTasksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
