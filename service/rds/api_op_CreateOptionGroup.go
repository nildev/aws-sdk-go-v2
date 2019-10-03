// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateOptionGroupMessage
type CreateOptionGroupInput struct {
	_ struct{} `type:"structure"`

	// Specifies the name of the engine that this option group should be associated
	// with.
	//
	// EngineName is a required field
	EngineName *string `type:"string" required:"true"`

	// Specifies the major version of the engine that this option group should be
	// associated with.
	//
	// MajorEngineVersion is a required field
	MajorEngineVersion *string `type:"string" required:"true"`

	// The description of the option group.
	//
	// OptionGroupDescription is a required field
	OptionGroupDescription *string `type:"string" required:"true"`

	// Specifies the name of the option group to be created.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: myoptiongroup
	//
	// OptionGroupName is a required field
	OptionGroupName *string `type:"string" required:"true"`

	// Tags to assign to the option group.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateOptionGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOptionGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateOptionGroupInput"}

	if s.EngineName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EngineName"))
	}

	if s.MajorEngineVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("MajorEngineVersion"))
	}

	if s.OptionGroupDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("OptionGroupDescription"))
	}

	if s.OptionGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OptionGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateOptionGroupResult
type CreateOptionGroupOutput struct {
	_ struct{} `type:"structure"`

	OptionGroup *OptionGroup `json:"rds:CreateOptionGroupOutput:OptionGroup" type:"structure"`
}

// String returns the string representation
func (s CreateOptionGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateOptionGroup = "CreateOptionGroup"

// CreateOptionGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new option group. You can create up to 20 option groups.
//
//    // Example sending a request using CreateOptionGroupRequest.
//    req := client.CreateOptionGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateOptionGroup
func (c *Client) CreateOptionGroupRequest(input *CreateOptionGroupInput) CreateOptionGroupRequest {
	op := &aws.Operation{
		Name:       opCreateOptionGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateOptionGroupInput{}
	}

	req := c.newRequest(op, input, &CreateOptionGroupOutput{})
	return CreateOptionGroupRequest{Request: req, Input: input, Copy: c.CreateOptionGroupRequest}
}

// CreateOptionGroupRequest is the request type for the
// CreateOptionGroup API operation.
type CreateOptionGroupRequest struct {
	*aws.Request
	Input *CreateOptionGroupInput
	Copy  func(*CreateOptionGroupInput) CreateOptionGroupRequest
}

// Send marshals and sends the CreateOptionGroup API request.
func (r CreateOptionGroupRequest) Send(ctx context.Context) (*CreateOptionGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateOptionGroupResponse{
		CreateOptionGroupOutput: r.Request.Data.(*CreateOptionGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateOptionGroupResponse is the response type for the
// CreateOptionGroup API operation.
type CreateOptionGroupResponse struct {
	*CreateOptionGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateOptionGroup request.
func (r *CreateOptionGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
