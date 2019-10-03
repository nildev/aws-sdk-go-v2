// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the DeleteSnapshot operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DeleteSnapshotRequest
type DeleteSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory snapshot to be deleted.
	//
	// SnapshotId is a required field
	SnapshotId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSnapshotInput"}

	if s.SnapshotId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the results of the DeleteSnapshot operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DeleteSnapshotResult
type DeleteSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory snapshot that was deleted.
	SnapshotId *string `json:"ds:DeleteSnapshotOutput:SnapshotId" type:"string"`
}

// String returns the string representation
func (s DeleteSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSnapshot = "DeleteSnapshot"

// DeleteSnapshotRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Deletes a directory snapshot.
//
//    // Example sending a request using DeleteSnapshotRequest.
//    req := client.DeleteSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DeleteSnapshot
func (c *Client) DeleteSnapshotRequest(input *DeleteSnapshotInput) DeleteSnapshotRequest {
	op := &aws.Operation{
		Name:       opDeleteSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSnapshotInput{}
	}

	req := c.newRequest(op, input, &DeleteSnapshotOutput{})
	return DeleteSnapshotRequest{Request: req, Input: input, Copy: c.DeleteSnapshotRequest}
}

// DeleteSnapshotRequest is the request type for the
// DeleteSnapshot API operation.
type DeleteSnapshotRequest struct {
	*aws.Request
	Input *DeleteSnapshotInput
	Copy  func(*DeleteSnapshotInput) DeleteSnapshotRequest
}

// Send marshals and sends the DeleteSnapshot API request.
func (r DeleteSnapshotRequest) Send(ctx context.Context) (*DeleteSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSnapshotResponse{
		DeleteSnapshotOutput: r.Request.Data.(*DeleteSnapshotOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSnapshotResponse is the response type for the
// DeleteSnapshot API operation.
type DeleteSnapshotResponse struct {
	*DeleteSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSnapshot request.
func (r *DeleteSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
