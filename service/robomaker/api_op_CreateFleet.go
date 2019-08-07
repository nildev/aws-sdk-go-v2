// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateFleetRequest
type CreateFleetInput struct {
	_ struct{} `type:"structure"`

	// The name of the fleet.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// A map that contains tag keys and tag values that are attached to the fleet.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFleetInput"}

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
func (s CreateFleetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateFleetResponse
type CreateFleetOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the fleet.
	Arn *string `json:"robomaker:CreateFleetOutput:Arn" locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time `json:"robomaker:CreateFleetOutput:CreatedAt" locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The name of the fleet.
	Name *string `json:"robomaker:CreateFleetOutput:Name" locationName:"name" min:"1" type:"string"`

	// The list of all tags added to the fleet.
	Tags map[string]string `json:"robomaker:CreateFleetOutput:Tags" locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateFleetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFleetOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opCreateFleet = "CreateFleet"

// CreateFleetRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Creates a fleet, a logical group of robots running the same robot application.
//
//    // Example sending a request using CreateFleetRequest.
//    req := client.CreateFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateFleet
func (c *Client) CreateFleetRequest(input *CreateFleetInput) CreateFleetRequest {
	op := &aws.Operation{
		Name:       opCreateFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/createFleet",
	}

	if input == nil {
		input = &CreateFleetInput{}
	}

	req := c.newRequest(op, input, &CreateFleetOutput{})
	return CreateFleetRequest{Request: req, Input: input, Copy: c.CreateFleetRequest}
}

// CreateFleetRequest is the request type for the
// CreateFleet API operation.
type CreateFleetRequest struct {
	*aws.Request
	Input *CreateFleetInput
	Copy  func(*CreateFleetInput) CreateFleetRequest
}

// Send marshals and sends the CreateFleet API request.
func (r CreateFleetRequest) Send(ctx context.Context) (*CreateFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFleetResponse{
		CreateFleetOutput: r.Request.Data.(*CreateFleetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFleetResponse is the response type for the
// CreateFleet API operation.
type CreateFleetResponse struct {
	*CreateFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFleet request.
func (r *CreateFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
