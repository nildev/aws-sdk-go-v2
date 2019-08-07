// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DescribeEntitiesDetectionJobRequest
type DescribeEntitiesDetectionJobInput struct {
	_ struct{} `type:"structure"`

	// The identifier that Amazon Comprehend generated for the job. The operation
	// returns this identifier in its response.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEntitiesDetectionJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEntitiesDetectionJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEntitiesDetectionJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DescribeEntitiesDetectionJobResponse
type DescribeEntitiesDetectionJobOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains the properties associated with an entities detection
	// job.
	EntitiesDetectionJobProperties *EntitiesDetectionJobProperties `json:"comprehend:DescribeEntitiesDetectionJobOutput:EntitiesDetectionJobProperties" type:"structure"`
}

// String returns the string representation
func (s DescribeEntitiesDetectionJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEntitiesDetectionJob = "DescribeEntitiesDetectionJob"

// DescribeEntitiesDetectionJobRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets the properties associated with an entities detection job. Use this operation
// to get the status of a detection job.
//
//    // Example sending a request using DescribeEntitiesDetectionJobRequest.
//    req := client.DescribeEntitiesDetectionJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DescribeEntitiesDetectionJob
func (c *Client) DescribeEntitiesDetectionJobRequest(input *DescribeEntitiesDetectionJobInput) DescribeEntitiesDetectionJobRequest {
	op := &aws.Operation{
		Name:       opDescribeEntitiesDetectionJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEntitiesDetectionJobInput{}
	}

	req := c.newRequest(op, input, &DescribeEntitiesDetectionJobOutput{})
	return DescribeEntitiesDetectionJobRequest{Request: req, Input: input, Copy: c.DescribeEntitiesDetectionJobRequest}
}

// DescribeEntitiesDetectionJobRequest is the request type for the
// DescribeEntitiesDetectionJob API operation.
type DescribeEntitiesDetectionJobRequest struct {
	*aws.Request
	Input *DescribeEntitiesDetectionJobInput
	Copy  func(*DescribeEntitiesDetectionJobInput) DescribeEntitiesDetectionJobRequest
}

// Send marshals and sends the DescribeEntitiesDetectionJob API request.
func (r DescribeEntitiesDetectionJobRequest) Send(ctx context.Context) (*DescribeEntitiesDetectionJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEntitiesDetectionJobResponse{
		DescribeEntitiesDetectionJobOutput: r.Request.Data.(*DescribeEntitiesDetectionJobOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEntitiesDetectionJobResponse is the response type for the
// DescribeEntitiesDetectionJob API operation.
type DescribeEntitiesDetectionJobResponse struct {
	*DescribeEntitiesDetectionJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEntitiesDetectionJob request.
func (r *DescribeEntitiesDetectionJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
