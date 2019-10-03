// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteLaunchTemplateVersionsRequest
type DeleteLaunchTemplateVersionsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateId *string `type:"string"`

	// The name of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateName *string `min:"3" type:"string"`

	// The version numbers of one or more launch template versions to delete.
	//
	// Versions is a required field
	Versions []string `locationName:"LaunchTemplateVersion" locationNameList:"item" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteLaunchTemplateVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLaunchTemplateVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLaunchTemplateVersionsInput"}
	if s.LaunchTemplateName != nil && len(*s.LaunchTemplateName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("LaunchTemplateName", 3))
	}

	if s.Versions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Versions"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteLaunchTemplateVersionsResult
type DeleteLaunchTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the launch template versions that were successfully deleted.
	SuccessfullyDeletedLaunchTemplateVersions []DeleteLaunchTemplateVersionsResponseSuccessItem `json:"ec2:DeleteLaunchTemplateVersionsOutput:SuccessfullyDeletedLaunchTemplateVersions" locationName:"successfullyDeletedLaunchTemplateVersionSet" locationNameList:"item" type:"list"`

	// Information about the launch template versions that could not be deleted.
	UnsuccessfullyDeletedLaunchTemplateVersions []DeleteLaunchTemplateVersionsResponseErrorItem `json:"ec2:DeleteLaunchTemplateVersionsOutput:UnsuccessfullyDeletedLaunchTemplateVersions" locationName:"unsuccessfullyDeletedLaunchTemplateVersionSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DeleteLaunchTemplateVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLaunchTemplateVersions = "DeleteLaunchTemplateVersions"

// DeleteLaunchTemplateVersionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes one or more versions of a launch template. You cannot delete the
// default version of a launch template; you must first assign a different version
// as the default. If the default version is the only version for the launch
// template, you must delete the entire launch template using DeleteLaunchTemplate.
//
//    // Example sending a request using DeleteLaunchTemplateVersionsRequest.
//    req := client.DeleteLaunchTemplateVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteLaunchTemplateVersions
func (c *Client) DeleteLaunchTemplateVersionsRequest(input *DeleteLaunchTemplateVersionsInput) DeleteLaunchTemplateVersionsRequest {
	op := &aws.Operation{
		Name:       opDeleteLaunchTemplateVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLaunchTemplateVersionsInput{}
	}

	req := c.newRequest(op, input, &DeleteLaunchTemplateVersionsOutput{})
	return DeleteLaunchTemplateVersionsRequest{Request: req, Input: input, Copy: c.DeleteLaunchTemplateVersionsRequest}
}

// DeleteLaunchTemplateVersionsRequest is the request type for the
// DeleteLaunchTemplateVersions API operation.
type DeleteLaunchTemplateVersionsRequest struct {
	*aws.Request
	Input *DeleteLaunchTemplateVersionsInput
	Copy  func(*DeleteLaunchTemplateVersionsInput) DeleteLaunchTemplateVersionsRequest
}

// Send marshals and sends the DeleteLaunchTemplateVersions API request.
func (r DeleteLaunchTemplateVersionsRequest) Send(ctx context.Context) (*DeleteLaunchTemplateVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLaunchTemplateVersionsResponse{
		DeleteLaunchTemplateVersionsOutput: r.Request.Data.(*DeleteLaunchTemplateVersionsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLaunchTemplateVersionsResponse is the response type for the
// DeleteLaunchTemplateVersions API operation.
type DeleteLaunchTemplateVersionsResponse struct {
	*DeleteLaunchTemplateVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLaunchTemplateVersions request.
func (r *DeleteLaunchTemplateVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
