// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a batch get repositories operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/BatchGetRepositoriesInput
type BatchGetRepositoriesInput struct {
	_ struct{} `type:"structure"`

	// The names of the repositories to get information about.
	//
	// RepositoryNames is a required field
	RepositoryNames []string `locationName:"repositoryNames" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetRepositoriesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetRepositoriesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetRepositoriesInput"}

	if s.RepositoryNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a batch get repositories operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/BatchGetRepositoriesOutput
type BatchGetRepositoriesOutput struct {
	_ struct{} `type:"structure"`

	// A list of repositories returned by the batch get repositories operation.
	Repositories []RepositoryMetadata `json:"codecommit:BatchGetRepositoriesOutput:Repositories" locationName:"repositories" type:"list"`

	// Returns a list of repository names for which information could not be found.
	RepositoriesNotFound []string `json:"codecommit:BatchGetRepositoriesOutput:RepositoriesNotFound" locationName:"repositoriesNotFound" type:"list"`
}

// String returns the string representation
func (s BatchGetRepositoriesOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetRepositories = "BatchGetRepositories"

// BatchGetRepositoriesRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about one or more repositories.
//
// The description field for a repository accepts all HTML characters and all
// valid Unicode characters. Applications that do not HTML-encode the description
// and display it in a web page could expose users to potentially malicious
// code. Make sure that you HTML-encode the description field in any application
// that uses this API to display the repository description on a web page.
//
//    // Example sending a request using BatchGetRepositoriesRequest.
//    req := client.BatchGetRepositoriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/BatchGetRepositories
func (c *Client) BatchGetRepositoriesRequest(input *BatchGetRepositoriesInput) BatchGetRepositoriesRequest {
	op := &aws.Operation{
		Name:       opBatchGetRepositories,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetRepositoriesInput{}
	}

	req := c.newRequest(op, input, &BatchGetRepositoriesOutput{})
	return BatchGetRepositoriesRequest{Request: req, Input: input, Copy: c.BatchGetRepositoriesRequest}
}

// BatchGetRepositoriesRequest is the request type for the
// BatchGetRepositories API operation.
type BatchGetRepositoriesRequest struct {
	*aws.Request
	Input *BatchGetRepositoriesInput
	Copy  func(*BatchGetRepositoriesInput) BatchGetRepositoriesRequest
}

// Send marshals and sends the BatchGetRepositories API request.
func (r BatchGetRepositoriesRequest) Send(ctx context.Context) (*BatchGetRepositoriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetRepositoriesResponse{
		BatchGetRepositoriesOutput: r.Request.Data.(*BatchGetRepositoriesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetRepositoriesResponse is the response type for the
// BatchGetRepositories API operation.
type BatchGetRepositoriesResponse struct {
	*BatchGetRepositoriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetRepositories request.
func (r *BatchGetRepositoriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
