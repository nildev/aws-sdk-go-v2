// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetFieldLevelEncryptionProfileConfigRequest
type GetFieldLevelEncryptionProfileConfigInput struct {
	_ struct{} `type:"structure"`

	// Get the ID for the field-level encryption profile configuration information.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFieldLevelEncryptionProfileConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFieldLevelEncryptionProfileConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFieldLevelEncryptionProfileConfigInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFieldLevelEncryptionProfileConfigInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetFieldLevelEncryptionProfileConfigResult
type GetFieldLevelEncryptionProfileConfigOutput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryptionProfileConfig"`

	// The current version of the field-level encryption profile configuration result.
	// For example: E2QWRUHAPOMQZL.
	ETag *string `json:"cloudfront:GetFieldLevelEncryptionProfileConfigOutput:ETag" location:"header" locationName:"ETag" type:"string"`

	// Return the field-level encryption profile configuration information.
	FieldLevelEncryptionProfileConfig *FieldLevelEncryptionProfileConfig `json:"cloudfront:GetFieldLevelEncryptionProfileConfigOutput:FieldLevelEncryptionProfileConfig" type:"structure"`
}

// String returns the string representation
func (s GetFieldLevelEncryptionProfileConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFieldLevelEncryptionProfileConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.FieldLevelEncryptionProfileConfig != nil {
		v := s.FieldLevelEncryptionProfileConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryptionProfileConfig", v, metadata)
	}
	return nil
}

const opGetFieldLevelEncryptionProfileConfig = "GetFieldLevelEncryptionProfileConfig2019_03_26"

// GetFieldLevelEncryptionProfileConfigRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Get the field-level encryption profile configuration information.
//
//    // Example sending a request using GetFieldLevelEncryptionProfileConfigRequest.
//    req := client.GetFieldLevelEncryptionProfileConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetFieldLevelEncryptionProfileConfig
func (c *Client) GetFieldLevelEncryptionProfileConfigRequest(input *GetFieldLevelEncryptionProfileConfigInput) GetFieldLevelEncryptionProfileConfigRequest {
	op := &aws.Operation{
		Name:       opGetFieldLevelEncryptionProfileConfig,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/field-level-encryption-profile/{Id}/config",
	}

	if input == nil {
		input = &GetFieldLevelEncryptionProfileConfigInput{}
	}

	req := c.newRequest(op, input, &GetFieldLevelEncryptionProfileConfigOutput{})
	return GetFieldLevelEncryptionProfileConfigRequest{Request: req, Input: input, Copy: c.GetFieldLevelEncryptionProfileConfigRequest}
}

// GetFieldLevelEncryptionProfileConfigRequest is the request type for the
// GetFieldLevelEncryptionProfileConfig API operation.
type GetFieldLevelEncryptionProfileConfigRequest struct {
	*aws.Request
	Input *GetFieldLevelEncryptionProfileConfigInput
	Copy  func(*GetFieldLevelEncryptionProfileConfigInput) GetFieldLevelEncryptionProfileConfigRequest
}

// Send marshals and sends the GetFieldLevelEncryptionProfileConfig API request.
func (r GetFieldLevelEncryptionProfileConfigRequest) Send(ctx context.Context) (*GetFieldLevelEncryptionProfileConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFieldLevelEncryptionProfileConfigResponse{
		GetFieldLevelEncryptionProfileConfigOutput: r.Request.Data.(*GetFieldLevelEncryptionProfileConfigOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFieldLevelEncryptionProfileConfigResponse is the response type for the
// GetFieldLevelEncryptionProfileConfig API operation.
type GetFieldLevelEncryptionProfileConfigResponse struct {
	*GetFieldLevelEncryptionProfileConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFieldLevelEncryptionProfileConfig request.
func (r *GetFieldLevelEncryptionProfileConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
