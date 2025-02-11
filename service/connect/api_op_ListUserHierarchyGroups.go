// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/ListUserHierarchyGroupsRequest
type ListUserHierarchyGroupsInput struct {
	_ struct{} `type:"structure"`

	// The identifier for your Amazon Connect instance. To find the ID of your instance,
	// open the AWS console and select Amazon Connect. Select the alias of the instance
	// in the Instance alias column. The instance ID is displayed in the Overview
	// section of your instance settings. For example, the instance ID is the set
	// of characters at the end of the instance ARN, after instance/, such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The maximum number of hierarchy groups to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListUserHierarchyGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUserHierarchyGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUserHierarchyGroupsInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListUserHierarchyGroupsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/ListUserHierarchyGroupsResponse
type ListUserHierarchyGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A string returned in the response. Use the value returned in the response
	// as the value of the NextToken in a subsequent request to retrieve the next
	// set of results.
	NextToken *string `type:"string"`

	// An array of HierarchyGroupSummary objects.
	UserHierarchyGroupSummaryList []HierarchyGroupSummary `type:"list"`
}

// String returns the string representation
func (s ListUserHierarchyGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListUserHierarchyGroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserHierarchyGroupSummaryList != nil {
		v := s.UserHierarchyGroupSummaryList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UserHierarchyGroupSummaryList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListUserHierarchyGroups = "ListUserHierarchyGroups"

// ListUserHierarchyGroupsRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Returns a UserHierarchyGroupSummaryList, which is an array of HierarchyGroupSummary
// objects that contain information about the hierarchy groups in your instance.
//
//    // Example sending a request using ListUserHierarchyGroupsRequest.
//    req := client.ListUserHierarchyGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/ListUserHierarchyGroups
func (c *Client) ListUserHierarchyGroupsRequest(input *ListUserHierarchyGroupsInput) ListUserHierarchyGroupsRequest {
	op := &aws.Operation{
		Name:       opListUserHierarchyGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/user-hierarchy-groups-summary/{InstanceId}",
	}

	if input == nil {
		input = &ListUserHierarchyGroupsInput{}
	}

	req := c.newRequest(op, input, &ListUserHierarchyGroupsOutput{})
	return ListUserHierarchyGroupsRequest{Request: req, Input: input, Copy: c.ListUserHierarchyGroupsRequest}
}

// ListUserHierarchyGroupsRequest is the request type for the
// ListUserHierarchyGroups API operation.
type ListUserHierarchyGroupsRequest struct {
	*aws.Request
	Input *ListUserHierarchyGroupsInput
	Copy  func(*ListUserHierarchyGroupsInput) ListUserHierarchyGroupsRequest
}

// Send marshals and sends the ListUserHierarchyGroups API request.
func (r ListUserHierarchyGroupsRequest) Send(ctx context.Context) (*ListUserHierarchyGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUserHierarchyGroupsResponse{
		ListUserHierarchyGroupsOutput: r.Request.Data.(*ListUserHierarchyGroupsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListUserHierarchyGroupsResponse is the response type for the
// ListUserHierarchyGroups API operation.
type ListUserHierarchyGroupsResponse struct {
	*ListUserHierarchyGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUserHierarchyGroups request.
func (r *ListUserHierarchyGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
