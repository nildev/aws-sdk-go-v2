// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListUsersInGroupRequest
type ListUsersInGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the group.
	//
	// GroupName is a required field
	GroupName *string `min:"1" type:"string" required:"true"`

	// The limit of the request to list users.
	Limit *int64 `type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `min:"1" type:"string"`

	// The user pool ID for the user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListUsersInGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUsersInGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUsersInGroupInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListUsersInGroupResponse
type ListUsersInGroupOutput struct {
	_ struct{} `type:"structure"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `json:"cognito-idp:ListUsersInGroupOutput:NextToken" min:"1" type:"string"`

	// The users returned in the request to list users.
	Users []UserType `json:"cognito-idp:ListUsersInGroupOutput:Users" type:"list"`
}

// String returns the string representation
func (s ListUsersInGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opListUsersInGroup = "ListUsersInGroup"

// ListUsersInGroupRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Lists the users in the specified group.
//
// Requires developer credentials.
//
//    // Example sending a request using ListUsersInGroupRequest.
//    req := client.ListUsersInGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListUsersInGroup
func (c *Client) ListUsersInGroupRequest(input *ListUsersInGroupInput) ListUsersInGroupRequest {
	op := &aws.Operation{
		Name:       opListUsersInGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListUsersInGroupInput{}
	}

	req := c.newRequest(op, input, &ListUsersInGroupOutput{})
	return ListUsersInGroupRequest{Request: req, Input: input, Copy: c.ListUsersInGroupRequest}
}

// ListUsersInGroupRequest is the request type for the
// ListUsersInGroup API operation.
type ListUsersInGroupRequest struct {
	*aws.Request
	Input *ListUsersInGroupInput
	Copy  func(*ListUsersInGroupInput) ListUsersInGroupRequest
}

// Send marshals and sends the ListUsersInGroup API request.
func (r ListUsersInGroupRequest) Send(ctx context.Context) (*ListUsersInGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUsersInGroupResponse{
		ListUsersInGroupOutput: r.Request.Data.(*ListUsersInGroupOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListUsersInGroupRequestPaginator returns a paginator for ListUsersInGroup.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListUsersInGroupRequest(input)
//   p := cognitoidentityprovider.NewListUsersInGroupRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListUsersInGroupPaginator(req ListUsersInGroupRequest) ListUsersInGroupPaginator {
	return ListUsersInGroupPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListUsersInGroupInput
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

// ListUsersInGroupPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListUsersInGroupPaginator struct {
	aws.Pager
}

func (p *ListUsersInGroupPaginator) CurrentPage() *ListUsersInGroupOutput {
	return p.Pager.CurrentPage().(*ListUsersInGroupOutput)
}

// ListUsersInGroupResponse is the response type for the
// ListUsersInGroup API operation.
type ListUsersInGroupResponse struct {
	*ListUsersInGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUsersInGroup request.
func (r *ListUsersInGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
