// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListInstanceProfilesForRoleRequest
type ListInstanceProfilesForRoleInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The name of the role to list instance profiles for.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListInstanceProfilesForRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListInstanceProfilesForRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListInstanceProfilesForRoleInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListInstanceProfilesForRole request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListInstanceProfilesForRoleResponse
type ListInstanceProfilesForRoleOutput struct {
	_ struct{} `type:"structure"`

	// A list of instance profiles.
	//
	// InstanceProfiles is a required field
	InstanceProfiles []InstanceProfile `json:"iam:ListInstanceProfilesForRoleOutput:InstanceProfiles" type:"list" required:"true"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `json:"iam:ListInstanceProfilesForRoleOutput:IsTruncated" type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `json:"iam:ListInstanceProfilesForRoleOutput:Marker" type:"string"`
}

// String returns the string representation
func (s ListInstanceProfilesForRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opListInstanceProfilesForRole = "ListInstanceProfilesForRole"

// ListInstanceProfilesForRoleRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the instance profiles that have the specified associated IAM role.
// If there are none, the operation returns an empty list. For more information
// about instance profiles, go to About Instance Profiles (https://docs.aws.amazon.com/IAM/latest/UserGuide/AboutInstanceProfiles.html).
//
// You can paginate the results using the MaxItems and Marker parameters.
//
//    // Example sending a request using ListInstanceProfilesForRoleRequest.
//    req := client.ListInstanceProfilesForRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListInstanceProfilesForRole
func (c *Client) ListInstanceProfilesForRoleRequest(input *ListInstanceProfilesForRoleInput) ListInstanceProfilesForRoleRequest {
	op := &aws.Operation{
		Name:       opListInstanceProfilesForRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListInstanceProfilesForRoleInput{}
	}

	req := c.newRequest(op, input, &ListInstanceProfilesForRoleOutput{})
	return ListInstanceProfilesForRoleRequest{Request: req, Input: input, Copy: c.ListInstanceProfilesForRoleRequest}
}

// ListInstanceProfilesForRoleRequest is the request type for the
// ListInstanceProfilesForRole API operation.
type ListInstanceProfilesForRoleRequest struct {
	*aws.Request
	Input *ListInstanceProfilesForRoleInput
	Copy  func(*ListInstanceProfilesForRoleInput) ListInstanceProfilesForRoleRequest
}

// Send marshals and sends the ListInstanceProfilesForRole API request.
func (r ListInstanceProfilesForRoleRequest) Send(ctx context.Context) (*ListInstanceProfilesForRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListInstanceProfilesForRoleResponse{
		ListInstanceProfilesForRoleOutput: r.Request.Data.(*ListInstanceProfilesForRoleOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListInstanceProfilesForRoleRequestPaginator returns a paginator for ListInstanceProfilesForRole.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListInstanceProfilesForRoleRequest(input)
//   p := iam.NewListInstanceProfilesForRoleRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListInstanceProfilesForRolePaginator(req ListInstanceProfilesForRoleRequest) ListInstanceProfilesForRolePaginator {
	return ListInstanceProfilesForRolePaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListInstanceProfilesForRoleInput
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

// ListInstanceProfilesForRolePaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListInstanceProfilesForRolePaginator struct {
	aws.Pager
}

func (p *ListInstanceProfilesForRolePaginator) CurrentPage() *ListInstanceProfilesForRoleOutput {
	return p.Pager.CurrentPage().(*ListInstanceProfilesForRoleOutput)
}

// ListInstanceProfilesForRoleResponse is the response type for the
// ListInstanceProfilesForRole API operation.
type ListInstanceProfilesForRoleResponse struct {
	*ListInstanceProfilesForRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListInstanceProfilesForRole request.
func (r *ListInstanceProfilesForRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
