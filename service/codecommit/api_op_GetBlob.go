// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a get blob operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetBlobInput
type GetBlobInput struct {
	_ struct{} `type:"structure"`

	// The ID of the blob, which is its SHA-1 pointer.
	//
	// BlobId is a required field
	BlobId *string `locationName:"blobId" type:"string" required:"true"`

	// The name of the repository that contains the blob.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBlobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBlobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBlobInput"}

	if s.BlobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BlobId"))
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

// Represents the output of a get blob operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetBlobOutput
type GetBlobOutput struct {
	_ struct{} `type:"structure"`

	// The content of the blob, usually a file.
	//
	// Content is automatically base64 encoded/decoded by the SDK.
	//
	// Content is a required field
	Content []byte `json:"codecommit:GetBlobOutput:Content" locationName:"content" type:"blob" required:"true"`
}

// String returns the string representation
func (s GetBlobOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetBlob = "GetBlob"

// GetBlobRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns the base-64 encoded content of an individual blob within a repository.
//
//    // Example sending a request using GetBlobRequest.
//    req := client.GetBlobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetBlob
func (c *Client) GetBlobRequest(input *GetBlobInput) GetBlobRequest {
	op := &aws.Operation{
		Name:       opGetBlob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetBlobInput{}
	}

	req := c.newRequest(op, input, &GetBlobOutput{})
	return GetBlobRequest{Request: req, Input: input, Copy: c.GetBlobRequest}
}

// GetBlobRequest is the request type for the
// GetBlob API operation.
type GetBlobRequest struct {
	*aws.Request
	Input *GetBlobInput
	Copy  func(*GetBlobInput) GetBlobRequest
}

// Send marshals and sends the GetBlob API request.
func (r GetBlobRequest) Send(ctx context.Context) (*GetBlobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBlobResponse{
		GetBlobOutput: r.Request.Data.(*GetBlobOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBlobResponse is the response type for the
// GetBlob API operation.
type GetBlobResponse struct {
	*GetBlobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBlob request.
func (r *GetBlobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
