// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to return information about the remote access session.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListRemoteAccessSessionsRequest
type ListRemoteAccessSessionsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the remote access session about which you
	// are requesting information.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListRemoteAccessSessionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRemoteAccessSessionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRemoteAccessSessionsInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server after AWS Device Farm makes a request
// to return information about the remote access session.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListRemoteAccessSessionsResult
type ListRemoteAccessSessionsOutput struct {
	_ struct{} `type:"structure"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `json:"devicefarm:ListRemoteAccessSessionsOutput:NextToken" locationName:"nextToken" min:"4" type:"string"`

	// A container representing the metadata from the service about each remote
	// access session you are requesting.
	RemoteAccessSessions []RemoteAccessSession `json:"devicefarm:ListRemoteAccessSessionsOutput:RemoteAccessSessions" locationName:"remoteAccessSessions" type:"list"`
}

// String returns the string representation
func (s ListRemoteAccessSessionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRemoteAccessSessions = "ListRemoteAccessSessions"

// ListRemoteAccessSessionsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Returns a list of all currently running remote access sessions.
//
//    // Example sending a request using ListRemoteAccessSessionsRequest.
//    req := client.ListRemoteAccessSessionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListRemoteAccessSessions
func (c *Client) ListRemoteAccessSessionsRequest(input *ListRemoteAccessSessionsInput) ListRemoteAccessSessionsRequest {
	op := &aws.Operation{
		Name:       opListRemoteAccessSessions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRemoteAccessSessionsInput{}
	}

	req := c.newRequest(op, input, &ListRemoteAccessSessionsOutput{})
	return ListRemoteAccessSessionsRequest{Request: req, Input: input, Copy: c.ListRemoteAccessSessionsRequest}
}

// ListRemoteAccessSessionsRequest is the request type for the
// ListRemoteAccessSessions API operation.
type ListRemoteAccessSessionsRequest struct {
	*aws.Request
	Input *ListRemoteAccessSessionsInput
	Copy  func(*ListRemoteAccessSessionsInput) ListRemoteAccessSessionsRequest
}

// Send marshals and sends the ListRemoteAccessSessions API request.
func (r ListRemoteAccessSessionsRequest) Send(ctx context.Context) (*ListRemoteAccessSessionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRemoteAccessSessionsResponse{
		ListRemoteAccessSessionsOutput: r.Request.Data.(*ListRemoteAccessSessionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRemoteAccessSessionsResponse is the response type for the
// ListRemoteAccessSessions API operation.
type ListRemoteAccessSessionsResponse struct {
	*ListRemoteAccessSessionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRemoteAccessSessions request.
func (r *ListRemoteAccessSessionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
