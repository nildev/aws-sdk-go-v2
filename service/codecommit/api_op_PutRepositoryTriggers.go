// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input ofa put repository triggers operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PutRepositoryTriggersInput
type PutRepositoryTriggersInput struct {
	_ struct{} `type:"structure"`

	// The name of the repository where you want to create or update the trigger.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`

	// The JSON block of configuration information for each trigger.
	//
	// Triggers is a required field
	Triggers []RepositoryTrigger `locationName:"triggers" type:"list" required:"true"`
}

// String returns the string representation
func (s PutRepositoryTriggersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRepositoryTriggersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRepositoryTriggersInput"}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if s.Triggers == nil {
		invalidParams.Add(aws.NewErrParamRequired("Triggers"))
	}
	if s.Triggers != nil {
		for i, v := range s.Triggers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Triggers", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a put repository triggers operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PutRepositoryTriggersOutput
type PutRepositoryTriggersOutput struct {
	_ struct{} `type:"structure"`

	// The system-generated unique ID for the create or update operation.
	ConfigurationId *string `json:"codecommit:PutRepositoryTriggersOutput:ConfigurationId" locationName:"configurationId" type:"string"`
}

// String returns the string representation
func (s PutRepositoryTriggersOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutRepositoryTriggers = "PutRepositoryTriggers"

// PutRepositoryTriggersRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Replaces all triggers for a repository. This can be used to create or delete
// triggers.
//
//    // Example sending a request using PutRepositoryTriggersRequest.
//    req := client.PutRepositoryTriggersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PutRepositoryTriggers
func (c *Client) PutRepositoryTriggersRequest(input *PutRepositoryTriggersInput) PutRepositoryTriggersRequest {
	op := &aws.Operation{
		Name:       opPutRepositoryTriggers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutRepositoryTriggersInput{}
	}

	req := c.newRequest(op, input, &PutRepositoryTriggersOutput{})
	return PutRepositoryTriggersRequest{Request: req, Input: input, Copy: c.PutRepositoryTriggersRequest}
}

// PutRepositoryTriggersRequest is the request type for the
// PutRepositoryTriggers API operation.
type PutRepositoryTriggersRequest struct {
	*aws.Request
	Input *PutRepositoryTriggersInput
	Copy  func(*PutRepositoryTriggersInput) PutRepositoryTriggersRequest
}

// Send marshals and sends the PutRepositoryTriggers API request.
func (r PutRepositoryTriggersRequest) Send(ctx context.Context) (*PutRepositoryTriggersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRepositoryTriggersResponse{
		PutRepositoryTriggersOutput: r.Request.Data.(*PutRepositoryTriggersOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRepositoryTriggersResponse is the response type for the
// PutRepositoryTriggers API operation.
type PutRepositoryTriggersResponse struct {
	*PutRepositoryTriggersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRepositoryTriggers request.
func (r *PutRepositoryTriggersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
