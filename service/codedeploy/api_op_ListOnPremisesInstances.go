// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a ListOnPremisesInstances operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListOnPremisesInstancesInput
type ListOnPremisesInstancesInput struct {
	_ struct{} `type:"structure"`

	// An identifier returned from the previous list on-premises instances call.
	// It can be used to return the next set of on-premises instances in the list.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The registration status of the on-premises instances:
	//
	//    * Deregistered: Include deregistered on-premises instances in the resulting
	//    list.
	//
	//    * Registered: Include registered on-premises instances in the resulting
	//    list.
	RegistrationStatus RegistrationStatus `locationName:"registrationStatus" type:"string" enum:"true"`

	// The on-premises instance tags that are used to restrict the on-premises instance
	// names returned.
	TagFilters []TagFilter `locationName:"tagFilters" type:"list"`
}

// String returns the string representation
func (s ListOnPremisesInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of the list on-premises instances operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListOnPremisesInstancesOutput
type ListOnPremisesInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The list of matching on-premises instance names.
	InstanceNames []string `json:"codedeploy:ListOnPremisesInstancesOutput:InstanceNames" locationName:"instanceNames" type:"list"`

	// If a large amount of information is returned, an identifier is also returned.
	// It can be used in a subsequent list on-premises instances call to return
	// the next set of on-premises instances in the list.
	NextToken *string `json:"codedeploy:ListOnPremisesInstancesOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListOnPremisesInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListOnPremisesInstances = "ListOnPremisesInstances"

// ListOnPremisesInstancesRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets a list of names for one or more on-premises instances.
//
// Unless otherwise specified, both registered and deregistered on-premises
// instance names are listed. To list only registered or deregistered on-premises
// instance names, use the registration status parameter.
//
//    // Example sending a request using ListOnPremisesInstancesRequest.
//    req := client.ListOnPremisesInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListOnPremisesInstances
func (c *Client) ListOnPremisesInstancesRequest(input *ListOnPremisesInstancesInput) ListOnPremisesInstancesRequest {
	op := &aws.Operation{
		Name:       opListOnPremisesInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListOnPremisesInstancesInput{}
	}

	req := c.newRequest(op, input, &ListOnPremisesInstancesOutput{})
	return ListOnPremisesInstancesRequest{Request: req, Input: input, Copy: c.ListOnPremisesInstancesRequest}
}

// ListOnPremisesInstancesRequest is the request type for the
// ListOnPremisesInstances API operation.
type ListOnPremisesInstancesRequest struct {
	*aws.Request
	Input *ListOnPremisesInstancesInput
	Copy  func(*ListOnPremisesInstancesInput) ListOnPremisesInstancesRequest
}

// Send marshals and sends the ListOnPremisesInstances API request.
func (r ListOnPremisesInstancesRequest) Send(ctx context.Context) (*ListOnPremisesInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListOnPremisesInstancesResponse{
		ListOnPremisesInstancesOutput: r.Request.Data.(*ListOnPremisesInstancesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListOnPremisesInstancesResponse is the response type for the
// ListOnPremisesInstances API operation.
type ListOnPremisesInstancesResponse struct {
	*ListOnPremisesInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListOnPremisesInstances request.
func (r *ListOnPremisesInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
