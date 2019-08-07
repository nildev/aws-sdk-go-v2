// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListEntitiesForPolicyRequest
type ListEntitiesForPolicyInput struct {
	_ struct{} `type:"structure"`

	// The entity type to use for filtering the results.
	//
	// For example, when EntityFilter is Role, only the roles that are attached
	// to the specified policy are returned. This parameter is optional. If it is
	// not included, all attached entities (users, groups, and roles) are returned.
	// The argument for this parameter must be one of the valid values listed below.
	EntityFilter EntityType `type:"string" enum:"true"`

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

	// The path prefix for filtering the results. This parameter is optional. If
	// it is not included, it defaults to a slash (/), listing all entities.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	PathPrefix *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the IAM policy for which you want the versions.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicyArn is a required field
	PolicyArn *string `min:"20" type:"string" required:"true"`

	// The policy usage method to use for filtering the results.
	//
	// To list only permissions policies, set PolicyUsageFilter to PermissionsPolicy.
	// To list only the policies used to set permissions boundaries, set the value
	// to PermissionsBoundary.
	//
	// This parameter is optional. If it is not included, all policies are returned.
	PolicyUsageFilter PolicyUsageType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListEntitiesForPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEntitiesForPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEntitiesForPolicyInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}
	if s.PathPrefix != nil && len(*s.PathPrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PathPrefix", 1))
	}

	if s.PolicyArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyArn"))
	}
	if s.PolicyArn != nil && len(*s.PolicyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListEntitiesForPolicy request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListEntitiesForPolicyResponse
type ListEntitiesForPolicyOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `json:"iam:ListEntitiesForPolicyOutput:IsTruncated" type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `json:"iam:ListEntitiesForPolicyOutput:Marker" type:"string"`

	// A list of IAM groups that the policy is attached to.
	PolicyGroups []PolicyGroup `json:"iam:ListEntitiesForPolicyOutput:PolicyGroups" type:"list"`

	// A list of IAM roles that the policy is attached to.
	PolicyRoles []PolicyRole `json:"iam:ListEntitiesForPolicyOutput:PolicyRoles" type:"list"`

	// A list of IAM users that the policy is attached to.
	PolicyUsers []PolicyUser `json:"iam:ListEntitiesForPolicyOutput:PolicyUsers" type:"list"`
}

// String returns the string representation
func (s ListEntitiesForPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opListEntitiesForPolicy = "ListEntitiesForPolicy"

// ListEntitiesForPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists all IAM users, groups, and roles that the specified managed policy
// is attached to.
//
// You can use the optional EntityFilter parameter to limit the results to a
// particular type of entity (users, groups, or roles). For example, to list
// only the roles that are attached to the specified policy, set EntityFilter
// to Role.
//
// You can paginate the results using the MaxItems and Marker parameters.
//
//    // Example sending a request using ListEntitiesForPolicyRequest.
//    req := client.ListEntitiesForPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListEntitiesForPolicy
func (c *Client) ListEntitiesForPolicyRequest(input *ListEntitiesForPolicyInput) ListEntitiesForPolicyRequest {
	op := &aws.Operation{
		Name:       opListEntitiesForPolicy,
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
		input = &ListEntitiesForPolicyInput{}
	}

	req := c.newRequest(op, input, &ListEntitiesForPolicyOutput{})
	return ListEntitiesForPolicyRequest{Request: req, Input: input, Copy: c.ListEntitiesForPolicyRequest}
}

// ListEntitiesForPolicyRequest is the request type for the
// ListEntitiesForPolicy API operation.
type ListEntitiesForPolicyRequest struct {
	*aws.Request
	Input *ListEntitiesForPolicyInput
	Copy  func(*ListEntitiesForPolicyInput) ListEntitiesForPolicyRequest
}

// Send marshals and sends the ListEntitiesForPolicy API request.
func (r ListEntitiesForPolicyRequest) Send(ctx context.Context) (*ListEntitiesForPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEntitiesForPolicyResponse{
		ListEntitiesForPolicyOutput: r.Request.Data.(*ListEntitiesForPolicyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEntitiesForPolicyRequestPaginator returns a paginator for ListEntitiesForPolicy.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEntitiesForPolicyRequest(input)
//   p := iam.NewListEntitiesForPolicyRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEntitiesForPolicyPaginator(req ListEntitiesForPolicyRequest) ListEntitiesForPolicyPaginator {
	return ListEntitiesForPolicyPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListEntitiesForPolicyInput
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

// ListEntitiesForPolicyPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEntitiesForPolicyPaginator struct {
	aws.Pager
}

func (p *ListEntitiesForPolicyPaginator) CurrentPage() *ListEntitiesForPolicyOutput {
	return p.Pager.CurrentPage().(*ListEntitiesForPolicyOutput)
}

// ListEntitiesForPolicyResponse is the response type for the
// ListEntitiesForPolicy API operation.
type ListEntitiesForPolicyResponse struct {
	*ListEntitiesForPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEntitiesForPolicy request.
func (r *ListEntitiesForPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
