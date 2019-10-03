// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CountOpenWorkflowExecutionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain containing the workflow executions to count.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// If specified, only workflow executions matching the WorkflowId in the filter
	// are counted.
	//
	// executionFilter, typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	ExecutionFilter *WorkflowExecutionFilter `locationName:"executionFilter" type:"structure"`

	// Specifies the start time criteria that workflow executions must meet in order
	// to be counted.
	//
	// StartTimeFilter is a required field
	StartTimeFilter *ExecutionTimeFilter `locationName:"startTimeFilter" type:"structure" required:"true"`

	// If specified, only executions that have a tag that matches the filter are
	// counted.
	//
	// executionFilter, typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	TagFilter *TagFilter `locationName:"tagFilter" type:"structure"`

	// Specifies the type of the workflow executions to be counted.
	//
	// executionFilter, typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	TypeFilter *WorkflowTypeFilter `locationName:"typeFilter" type:"structure"`
}

// String returns the string representation
func (s CountOpenWorkflowExecutionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CountOpenWorkflowExecutionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CountOpenWorkflowExecutionsInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.StartTimeFilter == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTimeFilter"))
	}
	if s.ExecutionFilter != nil {
		if err := s.ExecutionFilter.Validate(); err != nil {
			invalidParams.AddNested("ExecutionFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.StartTimeFilter != nil {
		if err := s.StartTimeFilter.Validate(); err != nil {
			invalidParams.AddNested("StartTimeFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.TagFilter != nil {
		if err := s.TagFilter.Validate(); err != nil {
			invalidParams.AddNested("TagFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.TypeFilter != nil {
		if err := s.TypeFilter.Validate(); err != nil {
			invalidParams.AddNested("TypeFilter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the count of workflow executions returned from CountOpenWorkflowExecutions
// or CountClosedWorkflowExecutions
type CountOpenWorkflowExecutionsOutput struct {
	_ struct{} `type:"structure"`

	// The number of workflow executions.
	//
	// Count is a required field
	Count *int64 `json:"swf:CountOpenWorkflowExecutionsOutput:Count" locationName:"count" type:"integer" required:"true"`

	// If set to true, indicates that the actual count was more than the maximum
	// supported by this API and the count returned is the truncated value.
	Truncated *bool `json:"swf:CountOpenWorkflowExecutionsOutput:Truncated" locationName:"truncated" type:"boolean"`
}

// String returns the string representation
func (s CountOpenWorkflowExecutionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCountOpenWorkflowExecutions = "CountOpenWorkflowExecutions"

// CountOpenWorkflowExecutionsRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Returns the number of open workflow executions within the given domain that
// meet the specified filtering criteria.
//
// This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * Constrain the following parameters by using a Condition element with
//    the appropriate keys. tagFilter.tag: String constraint. The key is swf:tagFilter.tag.
//    typeFilter.name: String constraint. The key is swf:typeFilter.name. typeFilter.version:
//    String constraint. The key is swf:typeFilter.version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using CountOpenWorkflowExecutionsRequest.
//    req := client.CountOpenWorkflowExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CountOpenWorkflowExecutionsRequest(input *CountOpenWorkflowExecutionsInput) CountOpenWorkflowExecutionsRequest {
	op := &aws.Operation{
		Name:       opCountOpenWorkflowExecutions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CountOpenWorkflowExecutionsInput{}
	}

	req := c.newRequest(op, input, &CountOpenWorkflowExecutionsOutput{})
	return CountOpenWorkflowExecutionsRequest{Request: req, Input: input, Copy: c.CountOpenWorkflowExecutionsRequest}
}

// CountOpenWorkflowExecutionsRequest is the request type for the
// CountOpenWorkflowExecutions API operation.
type CountOpenWorkflowExecutionsRequest struct {
	*aws.Request
	Input *CountOpenWorkflowExecutionsInput
	Copy  func(*CountOpenWorkflowExecutionsInput) CountOpenWorkflowExecutionsRequest
}

// Send marshals and sends the CountOpenWorkflowExecutions API request.
func (r CountOpenWorkflowExecutionsRequest) Send(ctx context.Context) (*CountOpenWorkflowExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CountOpenWorkflowExecutionsResponse{
		CountOpenWorkflowExecutionsOutput: r.Request.Data.(*CountOpenWorkflowExecutionsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CountOpenWorkflowExecutionsResponse is the response type for the
// CountOpenWorkflowExecutions API operation.
type CountOpenWorkflowExecutionsResponse struct {
	*CountOpenWorkflowExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CountOpenWorkflowExecutions request.
func (r *CountOpenWorkflowExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
