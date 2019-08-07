// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/CreateDatasetContentRequest
type CreateDatasetContentInput struct {
	_ struct{} `type:"structure"`

	// The name of the data set.
	//
	// DatasetName is a required field
	DatasetName *string `location:"uri" locationName:"datasetName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDatasetContentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDatasetContentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDatasetContentInput"}

	if s.DatasetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetName"))
	}
	if s.DatasetName != nil && len(*s.DatasetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDatasetContentInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DatasetName != nil {
		v := *s.DatasetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "datasetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/CreateDatasetContentResponse
type CreateDatasetContentOutput struct {
	_ struct{} `type:"structure"`

	// The version ID of the data set contents which are being created.
	VersionId *string `json:"iotanalytics:CreateDatasetContentOutput:VersionId" locationName:"versionId" min:"7" type:"string"`
}

// String returns the string representation
func (s CreateDatasetContentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDatasetContentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "versionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDatasetContent = "CreateDatasetContent"

// CreateDatasetContentRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Creates the content of a data set by applying a "queryAction" (a SQL query)
// or a "containerAction" (executing a containerized application).
//
//    // Example sending a request using CreateDatasetContentRequest.
//    req := client.CreateDatasetContentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/CreateDatasetContent
func (c *Client) CreateDatasetContentRequest(input *CreateDatasetContentInput) CreateDatasetContentRequest {
	op := &aws.Operation{
		Name:       opCreateDatasetContent,
		HTTPMethod: "POST",
		HTTPPath:   "/datasets/{datasetName}/content",
	}

	if input == nil {
		input = &CreateDatasetContentInput{}
	}

	req := c.newRequest(op, input, &CreateDatasetContentOutput{})
	return CreateDatasetContentRequest{Request: req, Input: input, Copy: c.CreateDatasetContentRequest}
}

// CreateDatasetContentRequest is the request type for the
// CreateDatasetContent API operation.
type CreateDatasetContentRequest struct {
	*aws.Request
	Input *CreateDatasetContentInput
	Copy  func(*CreateDatasetContentInput) CreateDatasetContentRequest
}

// Send marshals and sends the CreateDatasetContent API request.
func (r CreateDatasetContentRequest) Send(ctx context.Context) (*CreateDatasetContentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDatasetContentResponse{
		CreateDatasetContentOutput: r.Request.Data.(*CreateDatasetContentOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDatasetContentResponse is the response type for the
// CreateDatasetContent API operation.
type CreateDatasetContentResponse struct {
	*CreateDatasetContentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDatasetContent request.
func (r *CreateDatasetContentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
