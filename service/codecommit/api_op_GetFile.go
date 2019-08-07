// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetFileInput
type GetFileInput struct {
	_ struct{} `type:"structure"`

	// The fully-quaified reference that identifies the commit that contains the
	// file. For example, you could specify a full commit ID, a tag, a branch name,
	// or a reference such as refs/heads/master. If none is provided, then the head
	// commit will be used.
	CommitSpecifier *string `locationName:"commitSpecifier" type:"string"`

	// The fully-qualified path to the file, including the full name and extension
	// of the file. For example, /examples/file.md is the fully-qualified path to
	// a file named file.md in a folder named examples.
	//
	// FilePath is a required field
	FilePath *string `locationName:"filePath" type:"string" required:"true"`

	// The name of the repository that contains the file.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFileInput"}

	if s.FilePath == nil {
		invalidParams.Add(aws.NewErrParamRequired("FilePath"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetFileOutput
type GetFileOutput struct {
	_ struct{} `type:"structure"`

	// The blob ID of the object that represents the file content.
	//
	// BlobId is a required field
	BlobId *string `json:"codecommit:GetFileOutput:BlobId" locationName:"blobId" type:"string" required:"true"`

	// The full commit ID of the commit that contains the content returned by GetFile.
	//
	// CommitId is a required field
	CommitId *string `json:"codecommit:GetFileOutput:CommitId" locationName:"commitId" type:"string" required:"true"`

	// The base-64 encoded binary data object that represents the content of the
	// file.
	//
	// FileContent is automatically base64 encoded/decoded by the SDK.
	//
	// FileContent is a required field
	FileContent []byte `json:"codecommit:GetFileOutput:FileContent" locationName:"fileContent" type:"blob" required:"true"`

	// The extrapolated file mode permissions of the blob. Valid values include
	// strings such as EXECUTABLE and not numeric values.
	//
	// The file mode permissions returned by this API are not the standard file
	// mode permission values, such as 100644, but rather extrapolated values. See
	// below for a full list of supported return values.
	//
	// FileMode is a required field
	FileMode FileModeTypeEnum `json:"codecommit:GetFileOutput:FileMode" locationName:"fileMode" type:"string" required:"true" enum:"true"`

	// The fully qualified path to the specified file. This returns the name and
	// extension of the file.
	//
	// FilePath is a required field
	FilePath *string `json:"codecommit:GetFileOutput:FilePath" locationName:"filePath" type:"string" required:"true"`

	// The size of the contents of the file, in bytes.
	//
	// FileSize is a required field
	FileSize *int64 `json:"codecommit:GetFileOutput:FileSize" locationName:"fileSize" type:"long" required:"true"`
}

// String returns the string representation
func (s GetFileOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetFile = "GetFile"

// GetFileRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns the base-64 encoded contents of a specified file and its metadata.
//
//    // Example sending a request using GetFileRequest.
//    req := client.GetFileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetFile
func (c *Client) GetFileRequest(input *GetFileInput) GetFileRequest {
	op := &aws.Operation{
		Name:       opGetFile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetFileInput{}
	}

	req := c.newRequest(op, input, &GetFileOutput{})
	return GetFileRequest{Request: req, Input: input, Copy: c.GetFileRequest}
}

// GetFileRequest is the request type for the
// GetFile API operation.
type GetFileRequest struct {
	*aws.Request
	Input *GetFileInput
	Copy  func(*GetFileInput) GetFileRequest
}

// Send marshals and sends the GetFile API request.
func (r GetFileRequest) Send(ctx context.Context) (*GetFileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFileResponse{
		GetFileOutput: r.Request.Data.(*GetFileOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFileResponse is the response type for the
// GetFile API operation.
type GetFileResponse struct {
	*GetFileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFile request.
func (r *GetFileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
