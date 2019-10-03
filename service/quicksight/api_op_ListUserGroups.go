// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListUserGroupsRequest
type ListUserGroupsInput struct {
	_ struct{} `type:"structure"`

	// The AWS Account ID that the user is in. Currently, you use the ID for the
	// AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The maximum number of results to return from this request.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`

	// A pagination token that can be used in a subsequent request.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`

	// The Amazon QuickSight user name that you want to list group memberships for.
	//
	// UserName is a required field
	UserName *string `location:"uri" locationName:"UserName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListUserGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUserGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUserGroupsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListUserGroupsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.0"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserName != nil {
		v := *s.UserName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "UserName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListUserGroupsResponse
type ListUserGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The list of groups the user is a member of.
	GroupList []Group `json:"quicksight:ListUserGroupsOutput:GroupList" type:"list"`

	// A pagination token that can be used in a subsequent request.
	NextToken *string `json:"quicksight:ListUserGroupsOutput:NextToken" type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `json:"quicksight:ListUserGroupsOutput:RequestId" type:"string"`

	// The HTTP status of the request.
	Status *int64 `json:"quicksight:ListUserGroupsOutput:Status" location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s ListUserGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListUserGroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GroupList != nil {
		v := s.GroupList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "GroupList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opListUserGroups = "ListUserGroups"

// ListUserGroupsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Lists the Amazon QuickSight groups that an Amazon QuickSight user is a member
// of.
//
// The permission resource is arn:aws:quicksight:us-east-1:<aws-account-id>:user/default/<user-name> .
//
// The response is a one or more group objects.
//
// CLI Sample:
//
// aws quicksight list-user-groups -\-user-name=Pat -\-aws-account-id=111122223333
// -\-namespace=default -\-region=us-east-1
//
//    // Example sending a request using ListUserGroupsRequest.
//    req := client.ListUserGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListUserGroups
func (c *Client) ListUserGroupsRequest(input *ListUserGroupsInput) ListUserGroupsRequest {
	op := &aws.Operation{
		Name:       opListUserGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/users/{UserName}/groups",
	}

	if input == nil {
		input = &ListUserGroupsInput{}
	}

	req := c.newRequest(op, input, &ListUserGroupsOutput{})
	return ListUserGroupsRequest{Request: req, Input: input, Copy: c.ListUserGroupsRequest}
}

// ListUserGroupsRequest is the request type for the
// ListUserGroups API operation.
type ListUserGroupsRequest struct {
	*aws.Request
	Input *ListUserGroupsInput
	Copy  func(*ListUserGroupsInput) ListUserGroupsRequest
}

// Send marshals and sends the ListUserGroups API request.
func (r ListUserGroupsRequest) Send(ctx context.Context) (*ListUserGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUserGroupsResponse{
		ListUserGroupsOutput: r.Request.Data.(*ListUserGroupsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListUserGroupsResponse is the response type for the
// ListUserGroups API operation.
type ListUserGroupsResponse struct {
	*ListUserGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUserGroups request.
func (r *ListUserGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
