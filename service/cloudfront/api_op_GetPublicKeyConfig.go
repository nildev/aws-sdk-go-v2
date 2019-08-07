// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetPublicKeyConfigRequest
type GetPublicKeyConfigInput struct {
	_ struct{} `type:"structure"`

	// Request the ID for the public key configuration.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPublicKeyConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPublicKeyConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPublicKeyConfigInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPublicKeyConfigInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetPublicKeyConfigResult
type GetPublicKeyConfigOutput struct {
	_ struct{} `type:"structure" payload:"PublicKeyConfig"`

	// The current version of the public key configuration. For example: E2QWRUHAPOMQZL.
	ETag *string `json:"cloudfront:GetPublicKeyConfigOutput:ETag" location:"header" locationName:"ETag" type:"string"`

	// Return the result for the public key configuration.
	PublicKeyConfig *PublicKeyConfig `json:"cloudfront:GetPublicKeyConfigOutput:PublicKeyConfig" type:"structure"`
}

// String returns the string representation
func (s GetPublicKeyConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPublicKeyConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.PublicKeyConfig != nil {
		v := s.PublicKeyConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "PublicKeyConfig", v, metadata)
	}
	return nil
}

const opGetPublicKeyConfig = "GetPublicKeyConfig2019_03_26"

// GetPublicKeyConfigRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Return public key configuration informaation
//
//    // Example sending a request using GetPublicKeyConfigRequest.
//    req := client.GetPublicKeyConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetPublicKeyConfig
func (c *Client) GetPublicKeyConfigRequest(input *GetPublicKeyConfigInput) GetPublicKeyConfigRequest {
	op := &aws.Operation{
		Name:       opGetPublicKeyConfig,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/public-key/{Id}/config",
	}

	if input == nil {
		input = &GetPublicKeyConfigInput{}
	}

	req := c.newRequest(op, input, &GetPublicKeyConfigOutput{})
	return GetPublicKeyConfigRequest{Request: req, Input: input, Copy: c.GetPublicKeyConfigRequest}
}

// GetPublicKeyConfigRequest is the request type for the
// GetPublicKeyConfig API operation.
type GetPublicKeyConfigRequest struct {
	*aws.Request
	Input *GetPublicKeyConfigInput
	Copy  func(*GetPublicKeyConfigInput) GetPublicKeyConfigRequest
}

// Send marshals and sends the GetPublicKeyConfig API request.
func (r GetPublicKeyConfigRequest) Send(ctx context.Context) (*GetPublicKeyConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPublicKeyConfigResponse{
		GetPublicKeyConfigOutput: r.Request.Data.(*GetPublicKeyConfigOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPublicKeyConfigResponse is the response type for the
// GetPublicKeyConfig API operation.
type GetPublicKeyConfigResponse struct {
	*GetPublicKeyConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPublicKeyConfig request.
func (r *GetPublicKeyConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
