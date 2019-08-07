// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloud9

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DescribeEnvironmentMembershipsRequest
type DescribeEnvironmentMembershipsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the environment to get environment member information about.
	EnvironmentId *string `locationName:"environmentId" type:"string"`

	// The maximum number of environment members to get information about.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// During a previous call, if there are more than 25 items in the list, only
	// the first 25 items are returned, along with a unique string called a next
	// token. To get the next batch of items in the list, call this operation again,
	// adding the next token to the call. To get all of the items in the list, keep
	// calling this operation with each subsequent next token that is returned,
	// until no more next tokens are returned.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The type of environment member permissions to get information about. Available
	// values include:
	//
	//    * owner: Owns the environment.
	//
	//    * read-only: Has read-only access to the environment.
	//
	//    * read-write: Has read-write access to the environment.
	//
	// If no value is specified, information about all environment members are returned.
	Permissions []Permissions `locationName:"permissions" type:"list"`

	// The Amazon Resource Name (ARN) of an individual environment member to get
	// information about. If no value is specified, information about all environment
	// members are returned.
	UserArn *string `locationName:"userArn" type:"string"`
}

// String returns the string representation
func (s DescribeEnvironmentMembershipsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DescribeEnvironmentMembershipsResult
type DescribeEnvironmentMembershipsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the environment members for the environment.
	Memberships []EnvironmentMember `json:"cloud9:DescribeEnvironmentMembershipsOutput:Memberships" locationName:"memberships" type:"list"`

	// If there are more than 25 items in the list, only the first 25 items are
	// returned, along with a unique string called a next token. To get the next
	// batch of items in the list, call this operation again, adding the next token
	// to the call.
	NextToken *string `json:"cloud9:DescribeEnvironmentMembershipsOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeEnvironmentMembershipsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEnvironmentMemberships = "DescribeEnvironmentMemberships"

// DescribeEnvironmentMembershipsRequest returns a request value for making API operation for
// AWS Cloud9.
//
// Gets information about environment members for an AWS Cloud9 development
// environment.
//
//    // Example sending a request using DescribeEnvironmentMembershipsRequest.
//    req := client.DescribeEnvironmentMembershipsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DescribeEnvironmentMemberships
func (c *Client) DescribeEnvironmentMembershipsRequest(input *DescribeEnvironmentMembershipsInput) DescribeEnvironmentMembershipsRequest {
	op := &aws.Operation{
		Name:       opDescribeEnvironmentMemberships,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEnvironmentMembershipsInput{}
	}

	req := c.newRequest(op, input, &DescribeEnvironmentMembershipsOutput{})
	return DescribeEnvironmentMembershipsRequest{Request: req, Input: input, Copy: c.DescribeEnvironmentMembershipsRequest}
}

// DescribeEnvironmentMembershipsRequest is the request type for the
// DescribeEnvironmentMemberships API operation.
type DescribeEnvironmentMembershipsRequest struct {
	*aws.Request
	Input *DescribeEnvironmentMembershipsInput
	Copy  func(*DescribeEnvironmentMembershipsInput) DescribeEnvironmentMembershipsRequest
}

// Send marshals and sends the DescribeEnvironmentMemberships API request.
func (r DescribeEnvironmentMembershipsRequest) Send(ctx context.Context) (*DescribeEnvironmentMembershipsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEnvironmentMembershipsResponse{
		DescribeEnvironmentMembershipsOutput: r.Request.Data.(*DescribeEnvironmentMembershipsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEnvironmentMembershipsRequestPaginator returns a paginator for DescribeEnvironmentMemberships.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEnvironmentMembershipsRequest(input)
//   p := cloud9.NewDescribeEnvironmentMembershipsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEnvironmentMembershipsPaginator(req DescribeEnvironmentMembershipsRequest) DescribeEnvironmentMembershipsPaginator {
	return DescribeEnvironmentMembershipsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEnvironmentMembershipsInput
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

// DescribeEnvironmentMembershipsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEnvironmentMembershipsPaginator struct {
	aws.Pager
}

func (p *DescribeEnvironmentMembershipsPaginator) CurrentPage() *DescribeEnvironmentMembershipsOutput {
	return p.Pager.CurrentPage().(*DescribeEnvironmentMembershipsOutput)
}

// DescribeEnvironmentMembershipsResponse is the response type for the
// DescribeEnvironmentMemberships API operation.
type DescribeEnvironmentMembershipsResponse struct {
	*DescribeEnvironmentMembershipsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEnvironmentMemberships request.
func (r *DescribeEnvironmentMembershipsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
