// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/ListAttributesRequest
type ListAttributesInput struct {
	_ struct{} `type:"structure"`

	// The name of the attribute with which to filter the results.
	AttributeName *string `locationName:"attributeName" type:"string"`

	// The value of the attribute with which to filter results. You must also specify
	// an attribute name to use this parameter.
	AttributeValue *string `locationName:"attributeValue" type:"string"`

	// The short name or full Amazon Resource Name (ARN) of the cluster to list
	// attributes. If you do not specify a cluster, the default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The maximum number of cluster results returned by ListAttributes in paginated
	// output. When this parameter is used, ListAttributes only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListAttributes
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If this parameter is not used, then ListAttributes returns up to 100
	// results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListAttributes request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The type of the target with which to list attributes.
	//
	// TargetType is a required field
	TargetType TargetType `locationName:"targetType" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ListAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAttributesInput"}
	if len(s.TargetType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TargetType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/ListAttributesResponse
type ListAttributesOutput struct {
	_ struct{} `type:"structure"`

	// A list of attribute objects that meet the criteria of the request.
	Attributes []Attribute `json:"ecs:ListAttributesOutput:Attributes" locationName:"attributes" type:"list"`

	// The nextToken value to include in a future ListAttributes request. When the
	// results of a ListAttributes request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string `json:"ecs:ListAttributesOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAttributes = "ListAttributes"

// ListAttributesRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Lists the attributes for Amazon ECS resources within a specified target type
// and cluster. When you specify a target type and cluster, ListAttributes returns
// a list of attribute objects, one for each attribute on each resource. You
// can filter the list of results to a single attribute name to only return
// results that have that name. You can also filter the results by attribute
// name and value, for example, to see which container instances in a cluster
// are running a Linux AMI (ecs.os-type=linux).
//
//    // Example sending a request using ListAttributesRequest.
//    req := client.ListAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/ListAttributes
func (c *Client) ListAttributesRequest(input *ListAttributesInput) ListAttributesRequest {
	op := &aws.Operation{
		Name:       opListAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAttributesInput{}
	}

	req := c.newRequest(op, input, &ListAttributesOutput{})
	return ListAttributesRequest{Request: req, Input: input, Copy: c.ListAttributesRequest}
}

// ListAttributesRequest is the request type for the
// ListAttributes API operation.
type ListAttributesRequest struct {
	*aws.Request
	Input *ListAttributesInput
	Copy  func(*ListAttributesInput) ListAttributesRequest
}

// Send marshals and sends the ListAttributes API request.
func (r ListAttributesRequest) Send(ctx context.Context) (*ListAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAttributesResponse{
		ListAttributesOutput: r.Request.Data.(*ListAttributesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAttributesResponse is the response type for the
// ListAttributes API operation.
type ListAttributesResponse struct {
	*ListAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAttributes request.
func (r *ListAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
