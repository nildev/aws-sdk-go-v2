// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/UpdateUserSecurityProfilesRequest
type UpdateUserSecurityProfilesInput struct {
	_ struct{} `type:"structure"`

	// The identifier for your Amazon Connect instance. To find the ID of your instance,
	// open the AWS console and select Amazon Connect. Select the alias of the instance
	// in the Instance alias column. The instance ID is displayed in the Overview
	// section of your instance settings. For example, the instance ID is the set
	// of characters at the end of the instance ARN, after instance/, such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The identifiers for the security profiles to assign to the user.
	//
	// SecurityProfileIds is a required field
	SecurityProfileIds []string `min:"1" type:"list" required:"true"`

	// The identifier of the user account to assign the security profiles.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"UserId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateUserSecurityProfilesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserSecurityProfilesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserSecurityProfilesInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if s.SecurityProfileIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityProfileIds"))
	}
	if s.SecurityProfileIds != nil && len(s.SecurityProfileIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityProfileIds", 1))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserSecurityProfilesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.SecurityProfileIds != nil {
		v := s.SecurityProfileIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SecurityProfileIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "UserId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/UpdateUserSecurityProfilesOutput
type UpdateUserSecurityProfilesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateUserSecurityProfilesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserSecurityProfilesOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateUserSecurityProfiles = "UpdateUserSecurityProfiles"

// UpdateUserSecurityProfilesRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Updates the security profiles assigned to the user.
//
//    // Example sending a request using UpdateUserSecurityProfilesRequest.
//    req := client.UpdateUserSecurityProfilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/UpdateUserSecurityProfiles
func (c *Client) UpdateUserSecurityProfilesRequest(input *UpdateUserSecurityProfilesInput) UpdateUserSecurityProfilesRequest {
	op := &aws.Operation{
		Name:       opUpdateUserSecurityProfiles,
		HTTPMethod: "POST",
		HTTPPath:   "/users/{InstanceId}/{UserId}/security-profiles",
	}

	if input == nil {
		input = &UpdateUserSecurityProfilesInput{}
	}

	req := c.newRequest(op, input, &UpdateUserSecurityProfilesOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateUserSecurityProfilesRequest{Request: req, Input: input, Copy: c.UpdateUserSecurityProfilesRequest}
}

// UpdateUserSecurityProfilesRequest is the request type for the
// UpdateUserSecurityProfiles API operation.
type UpdateUserSecurityProfilesRequest struct {
	*aws.Request
	Input *UpdateUserSecurityProfilesInput
	Copy  func(*UpdateUserSecurityProfilesInput) UpdateUserSecurityProfilesRequest
}

// Send marshals and sends the UpdateUserSecurityProfiles API request.
func (r UpdateUserSecurityProfilesRequest) Send(ctx context.Context) (*UpdateUserSecurityProfilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserSecurityProfilesResponse{
		UpdateUserSecurityProfilesOutput: r.Request.Data.(*UpdateUserSecurityProfilesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserSecurityProfilesResponse is the response type for the
// UpdateUserSecurityProfiles API operation.
type UpdateUserSecurityProfilesResponse struct {
	*UpdateUserSecurityProfilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUserSecurityProfiles request.
func (r *UpdateUserSecurityProfilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
