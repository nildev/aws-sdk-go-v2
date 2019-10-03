// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a BatchGetApplications operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/BatchGetApplicationsInput
type BatchGetApplicationsInput struct {
	_ struct{} `type:"structure"`

	// A list of application names separated by spaces. The maximum number of application
	// names you can specify is 25.
	//
	// ApplicationNames is a required field
	ApplicationNames []string `locationName:"applicationNames" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetApplicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetApplicationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetApplicationsInput"}

	if s.ApplicationNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a BatchGetApplications operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/BatchGetApplicationsOutput
type BatchGetApplicationsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the applications.
	ApplicationsInfo []ApplicationInfo `json:"codedeploy:BatchGetApplicationsOutput:ApplicationsInfo" locationName:"applicationsInfo" type:"list"`
}

// String returns the string representation
func (s BatchGetApplicationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetApplications = "BatchGetApplications"

// BatchGetApplicationsRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets information about one or more applications. The maximum number of applications
// that can be returned is 25.
//
//    // Example sending a request using BatchGetApplicationsRequest.
//    req := client.BatchGetApplicationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/BatchGetApplications
func (c *Client) BatchGetApplicationsRequest(input *BatchGetApplicationsInput) BatchGetApplicationsRequest {
	op := &aws.Operation{
		Name:       opBatchGetApplications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetApplicationsInput{}
	}

	req := c.newRequest(op, input, &BatchGetApplicationsOutput{})
	return BatchGetApplicationsRequest{Request: req, Input: input, Copy: c.BatchGetApplicationsRequest}
}

// BatchGetApplicationsRequest is the request type for the
// BatchGetApplications API operation.
type BatchGetApplicationsRequest struct {
	*aws.Request
	Input *BatchGetApplicationsInput
	Copy  func(*BatchGetApplicationsInput) BatchGetApplicationsRequest
}

// Send marshals and sends the BatchGetApplications API request.
func (r BatchGetApplicationsRequest) Send(ctx context.Context) (*BatchGetApplicationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetApplicationsResponse{
		BatchGetApplicationsOutput: r.Request.Data.(*BatchGetApplicationsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetApplicationsResponse is the response type for the
// BatchGetApplications API operation.
type BatchGetApplicationsResponse struct {
	*BatchGetApplicationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetApplications request.
func (r *BatchGetApplicationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
