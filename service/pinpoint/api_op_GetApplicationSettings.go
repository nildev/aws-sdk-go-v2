// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetApplicationSettingsRequest
type GetApplicationSettingsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetApplicationSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApplicationSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApplicationSettingsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApplicationSettingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetApplicationSettingsResponse
type GetApplicationSettingsOutput struct {
	_ struct{} `type:"structure" payload:"ApplicationSettingsResource"`

	// Provides information about an application, including the default settings
	// for an application.
	//
	// ApplicationSettingsResource is a required field
	ApplicationSettingsResource *ApplicationSettingsResource `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetApplicationSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApplicationSettingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationSettingsResource != nil {
		v := s.ApplicationSettingsResource

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ApplicationSettingsResource", v, metadata)
	}
	return nil
}

const opGetApplicationSettings = "GetApplicationSettings"

// GetApplicationSettingsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the settings for an application.
//
//    // Example sending a request using GetApplicationSettingsRequest.
//    req := client.GetApplicationSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetApplicationSettings
func (c *Client) GetApplicationSettingsRequest(input *GetApplicationSettingsInput) GetApplicationSettingsRequest {
	op := &aws.Operation{
		Name:       opGetApplicationSettings,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/settings",
	}

	if input == nil {
		input = &GetApplicationSettingsInput{}
	}

	req := c.newRequest(op, input, &GetApplicationSettingsOutput{})
	return GetApplicationSettingsRequest{Request: req, Input: input, Copy: c.GetApplicationSettingsRequest}
}

// GetApplicationSettingsRequest is the request type for the
// GetApplicationSettings API operation.
type GetApplicationSettingsRequest struct {
	*aws.Request
	Input *GetApplicationSettingsInput
	Copy  func(*GetApplicationSettingsInput) GetApplicationSettingsRequest
}

// Send marshals and sends the GetApplicationSettings API request.
func (r GetApplicationSettingsRequest) Send(ctx context.Context) (*GetApplicationSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApplicationSettingsResponse{
		GetApplicationSettingsOutput: r.Request.Data.(*GetApplicationSettingsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApplicationSettingsResponse is the response type for the
// GetApplicationSettings API operation.
type GetApplicationSettingsResponse struct {
	*GetApplicationSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApplicationSettings request.
func (r *GetApplicationSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
