// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetImportJobsRequest
type GetImportJobsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	Token *string `location:"querystring" locationName:"token" type:"string"`
}

// String returns the string representation
func (s GetImportJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetImportJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetImportJobsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImportJobsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "page-size", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetImportJobsResponse
type GetImportJobsOutput struct {
	_ struct{} `type:"structure" payload:"ImportJobsResponse"`

	// Provides information about the status and settings of all the import jobs
	// that are associated with an application or segment. An import job is a job
	// that imports endpoint definitions from one or more files.
	//
	// ImportJobsResponse is a required field
	ImportJobsResponse *ImportJobsResponse `json:"pinpoint:GetImportJobsOutput:ImportJobsResponse" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetImportJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImportJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ImportJobsResponse != nil {
		v := s.ImportJobsResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ImportJobsResponse", v, metadata)
	}
	return nil
}

const opGetImportJobs = "GetImportJobs"

// GetImportJobsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status and settings of all the import jobs
// for an application.
//
//    // Example sending a request using GetImportJobsRequest.
//    req := client.GetImportJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetImportJobs
func (c *Client) GetImportJobsRequest(input *GetImportJobsInput) GetImportJobsRequest {
	op := &aws.Operation{
		Name:       opGetImportJobs,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/jobs/import",
	}

	if input == nil {
		input = &GetImportJobsInput{}
	}

	req := c.newRequest(op, input, &GetImportJobsOutput{})
	return GetImportJobsRequest{Request: req, Input: input, Copy: c.GetImportJobsRequest}
}

// GetImportJobsRequest is the request type for the
// GetImportJobs API operation.
type GetImportJobsRequest struct {
	*aws.Request
	Input *GetImportJobsInput
	Copy  func(*GetImportJobsInput) GetImportJobsRequest
}

// Send marshals and sends the GetImportJobs API request.
func (r GetImportJobsRequest) Send(ctx context.Context) (*GetImportJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetImportJobsResponse{
		GetImportJobsOutput: r.Request.Data.(*GetImportJobsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetImportJobsResponse is the response type for the
// GetImportJobs API operation.
type GetImportJobsResponse struct {
	*GetImportJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetImportJobs request.
func (r *GetImportJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
