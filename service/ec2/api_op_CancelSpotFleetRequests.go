// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CancelSpotFleetRequests.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CancelSpotFleetRequestsRequest
type CancelSpotFleetRequestsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The IDs of the Spot Fleet requests.
	//
	// SpotFleetRequestIds is a required field
	SpotFleetRequestIds []string `locationName:"spotFleetRequestId" locationNameList:"item" type:"list" required:"true"`

	// Indicates whether to terminate instances for a Spot Fleet request if it is
	// canceled successfully.
	//
	// TerminateInstances is a required field
	TerminateInstances *bool `locationName:"terminateInstances" type:"boolean" required:"true"`
}

// String returns the string representation
func (s CancelSpotFleetRequestsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelSpotFleetRequestsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelSpotFleetRequestsInput"}

	if s.SpotFleetRequestIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SpotFleetRequestIds"))
	}

	if s.TerminateInstances == nil {
		invalidParams.Add(aws.NewErrParamRequired("TerminateInstances"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CancelSpotFleetRequests.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CancelSpotFleetRequestsResponse
type CancelSpotFleetRequestsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Spot Fleet requests that are successfully canceled.
	SuccessfulFleetRequests []CancelSpotFleetRequestsSuccessItem `json:"ec2:CancelSpotFleetRequestsOutput:SuccessfulFleetRequests" locationName:"successfulFleetRequestSet" locationNameList:"item" type:"list"`

	// Information about the Spot Fleet requests that are not successfully canceled.
	UnsuccessfulFleetRequests []CancelSpotFleetRequestsErrorItem `json:"ec2:CancelSpotFleetRequestsOutput:UnsuccessfulFleetRequests" locationName:"unsuccessfulFleetRequestSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CancelSpotFleetRequestsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelSpotFleetRequests = "CancelSpotFleetRequests"

// CancelSpotFleetRequestsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Cancels the specified Spot Fleet requests.
//
// After you cancel a Spot Fleet request, the Spot Fleet launches no new Spot
// Instances. You must specify whether the Spot Fleet should also terminate
// its Spot Instances. If you terminate the instances, the Spot Fleet request
// enters the cancelled_terminating state. Otherwise, the Spot Fleet request
// enters the cancelled_running state and the instances continue to run until
// they are interrupted or you terminate them manually.
//
//    // Example sending a request using CancelSpotFleetRequestsRequest.
//    req := client.CancelSpotFleetRequestsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CancelSpotFleetRequests
func (c *Client) CancelSpotFleetRequestsRequest(input *CancelSpotFleetRequestsInput) CancelSpotFleetRequestsRequest {
	op := &aws.Operation{
		Name:       opCancelSpotFleetRequests,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelSpotFleetRequestsInput{}
	}

	req := c.newRequest(op, input, &CancelSpotFleetRequestsOutput{})
	return CancelSpotFleetRequestsRequest{Request: req, Input: input, Copy: c.CancelSpotFleetRequestsRequest}
}

// CancelSpotFleetRequestsRequest is the request type for the
// CancelSpotFleetRequests API operation.
type CancelSpotFleetRequestsRequest struct {
	*aws.Request
	Input *CancelSpotFleetRequestsInput
	Copy  func(*CancelSpotFleetRequestsInput) CancelSpotFleetRequestsRequest
}

// Send marshals and sends the CancelSpotFleetRequests API request.
func (r CancelSpotFleetRequestsRequest) Send(ctx context.Context) (*CancelSpotFleetRequestsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelSpotFleetRequestsResponse{
		CancelSpotFleetRequestsOutput: r.Request.Data.(*CancelSpotFleetRequestsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelSpotFleetRequestsResponse is the response type for the
// CancelSpotFleetRequests API operation.
type CancelSpotFleetRequestsResponse struct {
	*CancelSpotFleetRequestsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelSpotFleetRequests request.
func (r *CancelSpotFleetRequestsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
