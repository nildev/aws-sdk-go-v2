// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/ListUserProfilesRequest
type ListUserProfilesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a response.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The continuation token for the next set of results, if the results cannot
	// be returned in one response.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListUserProfilesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUserProfilesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUserProfilesInput"}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/ListUserProfilesResult
type ListUserProfilesOutput struct {
	_ struct{} `type:"structure"`

	// The continuation token to use when requesting the next set of results, if
	// there are more results to be returned.
	NextToken *string `json:"codestar:ListUserProfilesOutput:NextToken" locationName:"nextToken" min:"1" type:"string"`

	// All the user profiles configured in AWS CodeStar for an AWS account.
	//
	// UserProfiles is a required field
	UserProfiles []UserProfileSummary `json:"codestar:ListUserProfilesOutput:UserProfiles" locationName:"userProfiles" type:"list" required:"true"`
}

// String returns the string representation
func (s ListUserProfilesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListUserProfiles = "ListUserProfiles"

// ListUserProfilesRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Lists all the user profiles configured for your AWS account in AWS CodeStar.
//
//    // Example sending a request using ListUserProfilesRequest.
//    req := client.ListUserProfilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/ListUserProfiles
func (c *Client) ListUserProfilesRequest(input *ListUserProfilesInput) ListUserProfilesRequest {
	op := &aws.Operation{
		Name:       opListUserProfiles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListUserProfilesInput{}
	}

	req := c.newRequest(op, input, &ListUserProfilesOutput{})
	return ListUserProfilesRequest{Request: req, Input: input, Copy: c.ListUserProfilesRequest}
}

// ListUserProfilesRequest is the request type for the
// ListUserProfiles API operation.
type ListUserProfilesRequest struct {
	*aws.Request
	Input *ListUserProfilesInput
	Copy  func(*ListUserProfilesInput) ListUserProfilesRequest
}

// Send marshals and sends the ListUserProfiles API request.
func (r ListUserProfilesRequest) Send(ctx context.Context) (*ListUserProfilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUserProfilesResponse{
		ListUserProfilesOutput: r.Request.Data.(*ListUserProfilesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListUserProfilesResponse is the response type for the
// ListUserProfiles API operation.
type ListUserProfilesResponse struct {
	*ListUserProfilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUserProfiles request.
func (r *ListUserProfilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
