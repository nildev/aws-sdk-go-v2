// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Requests API Gateway to change information about a Deployment resource.
type UpdateDeploymentInput struct {
	_ struct{} `type:"structure"`

	// The replacement identifier for the Deployment resource to change information
	// about.
	//
	// DeploymentId is a required field
	DeploymentId *string `location:"uri" locationName:"deployment_id" type:"string" required:"true"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDeploymentInput"}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.PatchOperations != nil {
		v := s.PatchOperations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "patchOperations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "deployment_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An immutable representation of a RestApi resource that can be called by users
// using Stages. A deployment must be associated with a Stage for it to be callable
// over the Internet.
//
// To create a deployment, call POST on the Deployments resource of a RestApi.
// To view, update, or delete a deployment, call GET, PATCH, or DELETE on the
// specified deployment resource (/restapis/{restapi_id}/deployments/{deployment_id}).
//
// RestApi, Deployments, Stage, AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-deployment.html),
// AWS SDKs (https://aws.amazon.com/tools/)
type UpdateDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// A summary of the RestApi at the date and time that the deployment resource
	// was created.
	ApiSummary map[string]map[string]MethodSnapshot `json:"apigateway:UpdateDeploymentOutput:ApiSummary" locationName:"apiSummary" type:"map"`

	// The date and time that the deployment resource was created.
	CreatedDate *time.Time `json:"apigateway:UpdateDeploymentOutput:CreatedDate" locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// The description for the deployment resource.
	Description *string `json:"apigateway:UpdateDeploymentOutput:Description" locationName:"description" type:"string"`

	// The identifier for the deployment resource.
	Id *string `json:"apigateway:UpdateDeploymentOutput:Id" locationName:"id" type:"string"`
}

// String returns the string representation
func (s UpdateDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiSummary != nil {
		v := s.ApiSummary

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "apiSummary", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms1 := ms0.Map(k1)
			ms1.Start()
			for k2, v2 := range v1 {
				ms1.MapSetFields(k2, v2)
			}
			ms1.End()
		}
		ms0.End()

	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateDeployment = "UpdateDeployment"

// UpdateDeploymentRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Changes information about a Deployment resource.
//
//    // Example sending a request using UpdateDeploymentRequest.
//    req := client.UpdateDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateDeploymentRequest(input *UpdateDeploymentInput) UpdateDeploymentRequest {
	op := &aws.Operation{
		Name:       opUpdateDeployment,
		HTTPMethod: "PATCH",
		HTTPPath:   "/restapis/{restapi_id}/deployments/{deployment_id}",
	}

	if input == nil {
		input = &UpdateDeploymentInput{}
	}

	req := c.newRequest(op, input, &UpdateDeploymentOutput{})
	return UpdateDeploymentRequest{Request: req, Input: input, Copy: c.UpdateDeploymentRequest}
}

// UpdateDeploymentRequest is the request type for the
// UpdateDeployment API operation.
type UpdateDeploymentRequest struct {
	*aws.Request
	Input *UpdateDeploymentInput
	Copy  func(*UpdateDeploymentInput) UpdateDeploymentRequest
}

// Send marshals and sends the UpdateDeployment API request.
func (r UpdateDeploymentRequest) Send(ctx context.Context) (*UpdateDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDeploymentResponse{
		UpdateDeploymentOutput: r.Request.Data.(*UpdateDeploymentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDeploymentResponse is the response type for the
// UpdateDeployment API operation.
type UpdateDeploymentResponse struct {
	*UpdateDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDeployment request.
func (r *UpdateDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
