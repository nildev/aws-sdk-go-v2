// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DeleteCommentContentInput
type DeleteCommentContentInput struct {
	_ struct{} `type:"structure"`

	// The unique, system-generated ID of the comment. To get this ID, use GetCommentsForComparedCommit
	// or GetCommentsForPullRequest.
	//
	// CommentId is a required field
	CommentId *string `locationName:"commentId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCommentContentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCommentContentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCommentContentInput"}

	if s.CommentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CommentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DeleteCommentContentOutput
type DeleteCommentContentOutput struct {
	_ struct{} `type:"structure"`

	// Information about the comment you just deleted.
	Comment *Comment `json:"codecommit:DeleteCommentContentOutput:Comment" locationName:"comment" type:"structure"`
}

// String returns the string representation
func (s DeleteCommentContentOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCommentContent = "DeleteCommentContent"

// DeleteCommentContentRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Deletes the content of a comment made on a change, file, or commit in a repository.
//
//    // Example sending a request using DeleteCommentContentRequest.
//    req := client.DeleteCommentContentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DeleteCommentContent
func (c *Client) DeleteCommentContentRequest(input *DeleteCommentContentInput) DeleteCommentContentRequest {
	op := &aws.Operation{
		Name:       opDeleteCommentContent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCommentContentInput{}
	}

	req := c.newRequest(op, input, &DeleteCommentContentOutput{})
	return DeleteCommentContentRequest{Request: req, Input: input, Copy: c.DeleteCommentContentRequest}
}

// DeleteCommentContentRequest is the request type for the
// DeleteCommentContent API operation.
type DeleteCommentContentRequest struct {
	*aws.Request
	Input *DeleteCommentContentInput
	Copy  func(*DeleteCommentContentInput) DeleteCommentContentRequest
}

// Send marshals and sends the DeleteCommentContent API request.
func (r DeleteCommentContentRequest) Send(ctx context.Context) (*DeleteCommentContentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCommentContentResponse{
		DeleteCommentContentOutput: r.Request.Data.(*DeleteCommentContentOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCommentContentResponse is the response type for the
// DeleteCommentContent API operation.
type DeleteCommentContentResponse struct {
	*DeleteCommentContentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCommentContent request.
func (r *DeleteCommentContentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
