// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Information about a deployment.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateDeploymentRequest
type CreateDeploymentInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	// The ID of the deployment if you wish to redeploy a previous deployment.
	DeploymentId *string `type:"string"`

	// The type of deployment. When used for ''CreateDeployment'', only ''NewDeployment''
	// and ''Redeployment'' are valid.
	DeploymentType DeploymentType `type:"string" enum:"true"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`

	// The ID of the group version to be deployed.
	GroupVersionId *string `type:"string"`
}

// String returns the string representation
func (s CreateDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeploymentInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DeploymentType) > 0 {
		v := s.DeploymentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.GroupVersionId != nil {
		v := *s.GroupVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GroupVersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AmznClientToken != nil {
		v := *s.AmznClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-Client-Token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateDeploymentResponse
type CreateDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the deployment.
	DeploymentArn *string `json:"greengrass:CreateDeploymentOutput:DeploymentArn" type:"string"`

	// The ID of the deployment.
	DeploymentId *string `json:"greengrass:CreateDeploymentOutput:DeploymentId" type:"string"`
}

// String returns the string representation
func (s CreateDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeploymentArn != nil {
		v := *s.DeploymentArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDeployment = "CreateDeployment"

// CreateDeploymentRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a deployment. ''CreateDeployment'' requests are idempotent with respect
// to the ''X-Amzn-Client-Token'' token and the request parameters.
//
//    // Example sending a request using CreateDeploymentRequest.
//    req := client.CreateDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateDeployment
func (c *Client) CreateDeploymentRequest(input *CreateDeploymentInput) CreateDeploymentRequest {
	op := &aws.Operation{
		Name:       opCreateDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/groups/{GroupId}/deployments",
	}

	if input == nil {
		input = &CreateDeploymentInput{}
	}

	req := c.newRequest(op, input, &CreateDeploymentOutput{})
	return CreateDeploymentRequest{Request: req, Input: input, Copy: c.CreateDeploymentRequest}
}

// CreateDeploymentRequest is the request type for the
// CreateDeployment API operation.
type CreateDeploymentRequest struct {
	*aws.Request
	Input *CreateDeploymentInput
	Copy  func(*CreateDeploymentInput) CreateDeploymentRequest
}

// Send marshals and sends the CreateDeployment API request.
func (r CreateDeploymentRequest) Send(ctx context.Context) (*CreateDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeploymentResponse{
		CreateDeploymentOutput: r.Request.Data.(*CreateDeploymentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeploymentResponse is the response type for the
// CreateDeployment API operation.
type CreateDeploymentResponse struct {
	*CreateDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeployment request.
func (r *CreateDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
