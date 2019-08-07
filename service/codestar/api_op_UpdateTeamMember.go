// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/UpdateTeamMemberRequest
type UpdateTeamMemberInput struct {
	_ struct{} `type:"structure"`

	// The ID of the project.
	//
	// ProjectId is a required field
	ProjectId *string `locationName:"projectId" min:"2" type:"string" required:"true"`

	// The role assigned to the user in the project. Project roles have different
	// levels of access. For more information, see Working with Teams (http://docs.aws.amazon.com/codestar/latest/userguide/working-with-teams.html)
	// in the AWS CodeStar User Guide.
	ProjectRole *string `locationName:"projectRole" type:"string"`

	// Whether a team member is allowed to remotely access project resources using
	// the SSH public key associated with the user's profile. Even if this is set
	// to True, the user must associate a public key with their profile before the
	// user can access resources.
	RemoteAccessAllowed *bool `locationName:"remoteAccessAllowed" type:"boolean"`

	// The Amazon Resource Name (ARN) of the user for whom you want to change team
	// membership attributes.
	//
	// UserArn is a required field
	UserArn *string `locationName:"userArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTeamMemberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTeamMemberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTeamMemberInput"}

	if s.ProjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectId"))
	}
	if s.ProjectId != nil && len(*s.ProjectId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectId", 2))
	}

	if s.UserArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserArn"))
	}
	if s.UserArn != nil && len(*s.UserArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("UserArn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/UpdateTeamMemberResult
type UpdateTeamMemberOutput struct {
	_ struct{} `type:"structure"`

	// The project role granted to the user.
	ProjectRole *string `json:"codestar:UpdateTeamMemberOutput:ProjectRole" locationName:"projectRole" type:"string"`

	// Whether a team member is allowed to remotely access project resources using
	// the SSH public key associated with the user's profile.
	RemoteAccessAllowed *bool `json:"codestar:UpdateTeamMemberOutput:RemoteAccessAllowed" locationName:"remoteAccessAllowed" type:"boolean"`

	// The Amazon Resource Name (ARN) of the user whose team membership attributes
	// were updated.
	UserArn *string `json:"codestar:UpdateTeamMemberOutput:UserArn" locationName:"userArn" min:"32" type:"string"`
}

// String returns the string representation
func (s UpdateTeamMemberOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTeamMember = "UpdateTeamMember"

// UpdateTeamMemberRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Updates a team member's attributes in an AWS CodeStar project. For example,
// you can change a team member's role in the project, or change whether they
// have remote access to project resources.
//
//    // Example sending a request using UpdateTeamMemberRequest.
//    req := client.UpdateTeamMemberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/UpdateTeamMember
func (c *Client) UpdateTeamMemberRequest(input *UpdateTeamMemberInput) UpdateTeamMemberRequest {
	op := &aws.Operation{
		Name:       opUpdateTeamMember,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTeamMemberInput{}
	}

	req := c.newRequest(op, input, &UpdateTeamMemberOutput{})
	return UpdateTeamMemberRequest{Request: req, Input: input, Copy: c.UpdateTeamMemberRequest}
}

// UpdateTeamMemberRequest is the request type for the
// UpdateTeamMember API operation.
type UpdateTeamMemberRequest struct {
	*aws.Request
	Input *UpdateTeamMemberInput
	Copy  func(*UpdateTeamMemberInput) UpdateTeamMemberRequest
}

// Send marshals and sends the UpdateTeamMember API request.
func (r UpdateTeamMemberRequest) Send(ctx context.Context) (*UpdateTeamMemberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTeamMemberResponse{
		UpdateTeamMemberOutput: r.Request.Data.(*UpdateTeamMemberOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTeamMemberResponse is the response type for the
// UpdateTeamMember API operation.
type UpdateTeamMemberResponse struct {
	*UpdateTeamMemberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTeamMember request.
func (r *UpdateTeamMemberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
