// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/DescribeUserRequest
type DescribeUserInput struct {
	_ struct{} `type:"structure"`

	// A system-assigned unique identifier for an SFTP server that has this user
	// assigned.
	//
	// ServerId is a required field
	ServerId *string `type:"string" required:"true"`

	// The name of the user assigned to one or more servers. User names are part
	// of the sign-in credentials to use the AWS Transfer service and perform file
	// transfer tasks.
	//
	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUserInput"}

	if s.ServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerId"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/DescribeUserResponse
type DescribeUserOutput struct {
	_ struct{} `type:"structure"`

	// A system-assigned unique identifier for an SFTP server that has this user
	// assigned.
	//
	// ServerId is a required field
	ServerId *string `json:"transfer:DescribeUserOutput:ServerId" type:"string" required:"true"`

	// An array containing the properties of the user account for the ServerID value
	// that you specified.
	//
	// User is a required field
	User *DescribedUser `json:"transfer:DescribeUserOutput:User" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeUser = "DescribeUser"

// DescribeUserRequest returns a request value for making API operation for
// AWS Transfer for SFTP.
//
// Describes the user assigned to a specific server, as identified by its ServerId
// property.
//
// The response from this call returns the properties of the user associated
// with the ServerId value that was specified.
//
//    // Example sending a request using DescribeUserRequest.
//    req := client.DescribeUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/DescribeUser
func (c *Client) DescribeUserRequest(input *DescribeUserInput) DescribeUserRequest {
	op := &aws.Operation{
		Name:       opDescribeUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeUserInput{}
	}

	req := c.newRequest(op, input, &DescribeUserOutput{})
	return DescribeUserRequest{Request: req, Input: input, Copy: c.DescribeUserRequest}
}

// DescribeUserRequest is the request type for the
// DescribeUser API operation.
type DescribeUserRequest struct {
	*aws.Request
	Input *DescribeUserInput
	Copy  func(*DescribeUserInput) DescribeUserRequest
}

// Send marshals and sends the DescribeUser API request.
func (r DescribeUserRequest) Send(ctx context.Context) (*DescribeUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserResponse{
		DescribeUserOutput: r.Request.Data.(*DescribeUserOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserResponse is the response type for the
// DescribeUser API operation.
type DescribeUserResponse struct {
	*DescribeUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUser request.
func (r *DescribeUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
