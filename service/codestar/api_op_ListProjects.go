// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/ListProjectsRequest
type ListProjectsInput struct {
	_ struct{} `type:"structure"`

	// The maximum amount of data that can be contained in a single set of results.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The continuation token to be used to return the next set of results, if the
	// results cannot be returned in one response.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListProjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProjectsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/ListProjectsResult
type ListProjectsOutput struct {
	_ struct{} `type:"structure"`

	// The continuation token to use when requesting the next set of results, if
	// there are more results to be returned.
	NextToken *string `json:"codestar:ListProjectsOutput:NextToken" locationName:"nextToken" min:"1" type:"string"`

	// A list of projects.
	//
	// Projects is a required field
	Projects []ProjectSummary `json:"codestar:ListProjectsOutput:Projects" locationName:"projects" type:"list" required:"true"`
}

// String returns the string representation
func (s ListProjectsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListProjects = "ListProjects"

// ListProjectsRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Lists all projects in AWS CodeStar associated with your AWS account.
//
//    // Example sending a request using ListProjectsRequest.
//    req := client.ListProjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/ListProjects
func (c *Client) ListProjectsRequest(input *ListProjectsInput) ListProjectsRequest {
	op := &aws.Operation{
		Name:       opListProjects,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListProjectsInput{}
	}

	req := c.newRequest(op, input, &ListProjectsOutput{})
	return ListProjectsRequest{Request: req, Input: input, Copy: c.ListProjectsRequest}
}

// ListProjectsRequest is the request type for the
// ListProjects API operation.
type ListProjectsRequest struct {
	*aws.Request
	Input *ListProjectsInput
	Copy  func(*ListProjectsInput) ListProjectsRequest
}

// Send marshals and sends the ListProjects API request.
func (r ListProjectsRequest) Send(ctx context.Context) (*ListProjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProjectsResponse{
		ListProjectsOutput: r.Request.Data.(*ListProjectsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListProjectsResponse is the response type for the
// ListProjects API operation.
type ListProjectsResponse struct {
	*ListProjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProjects request.
func (r *ListProjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
