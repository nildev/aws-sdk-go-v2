// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/GetUploadStatusRequest
type GetUploadStatusInput struct {
	_ struct{} `type:"structure"`

	// The ID of the upload. This value is returned by the UploadEntityDefinitions
	// action.
	//
	// UploadId is a required field
	UploadId *string `locationName:"uploadId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetUploadStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUploadStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUploadStatusInput"}

	if s.UploadId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UploadId"))
	}
	if s.UploadId != nil && len(*s.UploadId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UploadId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/GetUploadStatusResponse
type GetUploadStatusOutput struct {
	_ struct{} `type:"structure"`

	// The date at which the upload was created.
	//
	// CreatedDate is a required field
	CreatedDate *time.Time `json:"iotthingsgraph:GetUploadStatusOutput:CreatedDate" locationName:"createdDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The reason for an upload failure.
	FailureReason []string `json:"iotthingsgraph:GetUploadStatusOutput:FailureReason" locationName:"failureReason" type:"list"`

	// The ARN of the upload.
	NamespaceArn *string `json:"iotthingsgraph:GetUploadStatusOutput:NamespaceArn" locationName:"namespaceArn" type:"string"`

	// The name of the upload's namespace.
	NamespaceName *string `json:"iotthingsgraph:GetUploadStatusOutput:NamespaceName" locationName:"namespaceName" type:"string"`

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace.
	NamespaceVersion *int64 `json:"iotthingsgraph:GetUploadStatusOutput:NamespaceVersion" locationName:"namespaceVersion" type:"long"`

	// The ID of the upload.
	//
	// UploadId is a required field
	UploadId *string `json:"iotthingsgraph:GetUploadStatusOutput:UploadId" locationName:"uploadId" min:"1" type:"string" required:"true"`

	// The status of the upload. The initial status is IN_PROGRESS. The response
	// show all validation failures if the upload fails.
	//
	// UploadStatus is a required field
	UploadStatus UploadStatus `json:"iotthingsgraph:GetUploadStatusOutput:UploadStatus" locationName:"uploadStatus" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetUploadStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetUploadStatus = "GetUploadStatus"

// GetUploadStatusRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Gets the status of the specified upload.
//
//    // Example sending a request using GetUploadStatusRequest.
//    req := client.GetUploadStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/GetUploadStatus
func (c *Client) GetUploadStatusRequest(input *GetUploadStatusInput) GetUploadStatusRequest {
	op := &aws.Operation{
		Name:       opGetUploadStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetUploadStatusInput{}
	}

	req := c.newRequest(op, input, &GetUploadStatusOutput{})
	return GetUploadStatusRequest{Request: req, Input: input, Copy: c.GetUploadStatusRequest}
}

// GetUploadStatusRequest is the request type for the
// GetUploadStatus API operation.
type GetUploadStatusRequest struct {
	*aws.Request
	Input *GetUploadStatusInput
	Copy  func(*GetUploadStatusInput) GetUploadStatusRequest
}

// Send marshals and sends the GetUploadStatus API request.
func (r GetUploadStatusRequest) Send(ctx context.Context) (*GetUploadStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUploadStatusResponse{
		GetUploadStatusOutput: r.Request.Data.(*GetUploadStatusOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUploadStatusResponse is the response type for the
// GetUploadStatus API operation.
type GetUploadStatusResponse struct {
	*GetUploadStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUploadStatus request.
func (r *GetUploadStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
