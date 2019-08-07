// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetExportSnapshotRecordsRequest
type GetExportSnapshotRecordsInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to a specific page of results for your get export
	// snapshot records request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetExportSnapshotRecordsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetExportSnapshotRecordsResult
type GetExportSnapshotRecordsOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects describing the export snapshot records.
	ExportSnapshotRecords []ExportSnapshotRecord `json:"lightsail:GetExportSnapshotRecordsOutput:ExportSnapshotRecords" locationName:"exportSnapshotRecords" type:"list"`

	// A token used for advancing to the next page of results of your get relational
	// database bundles request.
	NextPageToken *string `json:"lightsail:GetExportSnapshotRecordsOutput:NextPageToken" locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetExportSnapshotRecordsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetExportSnapshotRecords = "GetExportSnapshotRecords"

// GetExportSnapshotRecordsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the export snapshot record created as a result of the export snapshot
// operation.
//
// An export snapshot record can be used to create a new Amazon EC2 instance
// and its related resources with the create cloud formation stack operation.
//
//    // Example sending a request using GetExportSnapshotRecordsRequest.
//    req := client.GetExportSnapshotRecordsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetExportSnapshotRecords
func (c *Client) GetExportSnapshotRecordsRequest(input *GetExportSnapshotRecordsInput) GetExportSnapshotRecordsRequest {
	op := &aws.Operation{
		Name:       opGetExportSnapshotRecords,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetExportSnapshotRecordsInput{}
	}

	req := c.newRequest(op, input, &GetExportSnapshotRecordsOutput{})
	return GetExportSnapshotRecordsRequest{Request: req, Input: input, Copy: c.GetExportSnapshotRecordsRequest}
}

// GetExportSnapshotRecordsRequest is the request type for the
// GetExportSnapshotRecords API operation.
type GetExportSnapshotRecordsRequest struct {
	*aws.Request
	Input *GetExportSnapshotRecordsInput
	Copy  func(*GetExportSnapshotRecordsInput) GetExportSnapshotRecordsRequest
}

// Send marshals and sends the GetExportSnapshotRecords API request.
func (r GetExportSnapshotRecordsRequest) Send(ctx context.Context) (*GetExportSnapshotRecordsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetExportSnapshotRecordsResponse{
		GetExportSnapshotRecordsOutput: r.Request.Data.(*GetExportSnapshotRecordsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetExportSnapshotRecordsResponse is the response type for the
// GetExportSnapshotRecords API operation.
type GetExportSnapshotRecordsResponse struct {
	*GetExportSnapshotRecordsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetExportSnapshotRecords request.
func (r *GetExportSnapshotRecordsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
