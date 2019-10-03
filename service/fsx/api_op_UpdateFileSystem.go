// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The request object for the UpdateFileSystem operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fsx-2018-03-01/UpdateFileSystemRequest
type UpdateFileSystemInput struct {
	_ struct{} `type:"structure"`

	// (Optional) A string of up to 64 ASCII characters that Amazon FSx uses to
	// ensure idempotent updates. This string is automatically filled on your behalf
	// when you use the AWS Command Line Interface (AWS CLI) or an AWS SDK.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// The globally unique ID of the file system, assigned by Amazon FSx.
	//
	// FileSystemId is a required field
	FileSystemId *string `min:"11" type:"string" required:"true"`

	// The configuration object for Amazon FSx for Lustre file systems used in the
	// UpdateFileSystem operation.
	LustreConfiguration *UpdateFileSystemLustreConfiguration `type:"structure"`

	// The configuration update for this Microsoft Windows file system. The only
	// supported options are for backup and maintenance and for self-managed Active
	// Directory configuration.
	WindowsConfiguration *UpdateFileSystemWindowsConfiguration `type:"structure"`
}

// String returns the string representation
func (s UpdateFileSystemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFileSystemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFileSystemInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.FileSystemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemId"))
	}
	if s.FileSystemId != nil && len(*s.FileSystemId) < 11 {
		invalidParams.Add(aws.NewErrParamMinLen("FileSystemId", 11))
	}
	if s.LustreConfiguration != nil {
		if err := s.LustreConfiguration.Validate(); err != nil {
			invalidParams.AddNested("LustreConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.WindowsConfiguration != nil {
		if err := s.WindowsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("WindowsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response object for the UpdateFileSystem operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fsx-2018-03-01/UpdateFileSystemResponse
type UpdateFileSystemOutput struct {
	_ struct{} `type:"structure"`

	// A description of the file system that was updated.
	FileSystem *FileSystem `json:"fsx:UpdateFileSystemOutput:FileSystem" type:"structure"`
}

// String returns the string representation
func (s UpdateFileSystemOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateFileSystem = "UpdateFileSystem"

// UpdateFileSystemRequest returns a request value for making API operation for
// Amazon FSx.
//
// Updates a file system configuration.
//
//    // Example sending a request using UpdateFileSystemRequest.
//    req := client.UpdateFileSystemRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fsx-2018-03-01/UpdateFileSystem
func (c *Client) UpdateFileSystemRequest(input *UpdateFileSystemInput) UpdateFileSystemRequest {
	op := &aws.Operation{
		Name:       opUpdateFileSystem,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateFileSystemInput{}
	}

	req := c.newRequest(op, input, &UpdateFileSystemOutput{})
	return UpdateFileSystemRequest{Request: req, Input: input, Copy: c.UpdateFileSystemRequest}
}

// UpdateFileSystemRequest is the request type for the
// UpdateFileSystem API operation.
type UpdateFileSystemRequest struct {
	*aws.Request
	Input *UpdateFileSystemInput
	Copy  func(*UpdateFileSystemInput) UpdateFileSystemRequest
}

// Send marshals and sends the UpdateFileSystem API request.
func (r UpdateFileSystemRequest) Send(ctx context.Context) (*UpdateFileSystemResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFileSystemResponse{
		UpdateFileSystemOutput: r.Request.Data.(*UpdateFileSystemOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFileSystemResponse is the response type for the
// UpdateFileSystem API operation.
type UpdateFileSystemResponse struct {
	*UpdateFileSystemOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFileSystem request.
func (r *UpdateFileSystemResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
