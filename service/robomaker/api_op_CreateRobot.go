// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateRobotRequest
type CreateRobotInput struct {
	_ struct{} `type:"structure"`

	// The target architecture of the robot.
	//
	// Architecture is a required field
	Architecture Architecture `locationName:"architecture" type:"string" required:"true" enum:"true"`

	// The Greengrass group id.
	//
	// GreengrassGroupId is a required field
	GreengrassGroupId *string `locationName:"greengrassGroupId" min:"1" type:"string" required:"true"`

	// The name for the robot.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// A map that contains tag keys and tag values that are attached to the robot.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateRobotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRobotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRobotInput"}
	if len(s.Architecture) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Architecture"))
	}

	if s.GreengrassGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GreengrassGroupId"))
	}
	if s.GreengrassGroupId != nil && len(*s.GreengrassGroupId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GreengrassGroupId", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRobotInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if len(s.Architecture) > 0 {
		v := s.Architecture

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "architecture", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.GreengrassGroupId != nil {
		v := *s.GreengrassGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "greengrassGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateRobotResponse
type CreateRobotOutput struct {
	_ struct{} `type:"structure"`

	// The target architecture of the robot.
	Architecture Architecture `json:"robomaker:CreateRobotOutput:Architecture" locationName:"architecture" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the robot.
	Arn *string `json:"robomaker:CreateRobotOutput:Arn" locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the robot was created.
	CreatedAt *time.Time `json:"robomaker:CreateRobotOutput:CreatedAt" locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The Amazon Resource Name (ARN) of the Greengrass group associated with the
	// robot.
	GreengrassGroupId *string `json:"robomaker:CreateRobotOutput:GreengrassGroupId" locationName:"greengrassGroupId" min:"1" type:"string"`

	// The name of the robot.
	Name *string `json:"robomaker:CreateRobotOutput:Name" locationName:"name" min:"1" type:"string"`

	// The list of all tags added to the robot.
	Tags map[string]string `json:"robomaker:CreateRobotOutput:Tags" locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateRobotOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRobotOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Architecture) > 0 {
		v := s.Architecture

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "architecture", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.GreengrassGroupId != nil {
		v := *s.GreengrassGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "greengrassGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opCreateRobot = "CreateRobot"

// CreateRobotRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Creates a robot.
//
//    // Example sending a request using CreateRobotRequest.
//    req := client.CreateRobotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateRobot
func (c *Client) CreateRobotRequest(input *CreateRobotInput) CreateRobotRequest {
	op := &aws.Operation{
		Name:       opCreateRobot,
		HTTPMethod: "POST",
		HTTPPath:   "/createRobot",
	}

	if input == nil {
		input = &CreateRobotInput{}
	}

	req := c.newRequest(op, input, &CreateRobotOutput{})
	return CreateRobotRequest{Request: req, Input: input, Copy: c.CreateRobotRequest}
}

// CreateRobotRequest is the request type for the
// CreateRobot API operation.
type CreateRobotRequest struct {
	*aws.Request
	Input *CreateRobotInput
	Copy  func(*CreateRobotInput) CreateRobotRequest
}

// Send marshals and sends the CreateRobot API request.
func (r CreateRobotRequest) Send(ctx context.Context) (*CreateRobotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRobotResponse{
		CreateRobotOutput: r.Request.Data.(*CreateRobotOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRobotResponse is the response type for the
// CreateRobot API operation.
type CreateRobotResponse struct {
	*CreateRobotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRobot request.
func (r *CreateRobotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
