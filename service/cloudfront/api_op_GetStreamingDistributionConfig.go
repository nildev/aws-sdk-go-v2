// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// To request to get a streaming distribution configuration.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetStreamingDistributionConfigRequest
type GetStreamingDistributionConfigInput struct {
	_ struct{} `type:"structure"`

	// The streaming distribution's ID.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetStreamingDistributionConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetStreamingDistributionConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetStreamingDistributionConfigInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetStreamingDistributionConfigInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// The returned result of the corresponding request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetStreamingDistributionConfigResult
type GetStreamingDistributionConfigOutput struct {
	_ struct{} `type:"structure" payload:"StreamingDistributionConfig"`

	// The current version of the configuration. For example: E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// The streaming distribution's configuration information.
	StreamingDistributionConfig *StreamingDistributionConfig `type:"structure"`
}

// String returns the string representation
func (s GetStreamingDistributionConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetStreamingDistributionConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.StreamingDistributionConfig != nil {
		v := s.StreamingDistributionConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "StreamingDistributionConfig", v, metadata)
	}
	return nil
}

const opGetStreamingDistributionConfig = "GetStreamingDistributionConfig2019_03_26"

// GetStreamingDistributionConfigRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Get the configuration information about a streaming distribution.
//
//    // Example sending a request using GetStreamingDistributionConfigRequest.
//    req := client.GetStreamingDistributionConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetStreamingDistributionConfig
func (c *Client) GetStreamingDistributionConfigRequest(input *GetStreamingDistributionConfigInput) GetStreamingDistributionConfigRequest {
	op := &aws.Operation{
		Name:       opGetStreamingDistributionConfig,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/streaming-distribution/{Id}/config",
	}

	if input == nil {
		input = &GetStreamingDistributionConfigInput{}
	}

	req := c.newRequest(op, input, &GetStreamingDistributionConfigOutput{})
	return GetStreamingDistributionConfigRequest{Request: req, Input: input, Copy: c.GetStreamingDistributionConfigRequest}
}

// GetStreamingDistributionConfigRequest is the request type for the
// GetStreamingDistributionConfig API operation.
type GetStreamingDistributionConfigRequest struct {
	*aws.Request
	Input *GetStreamingDistributionConfigInput
	Copy  func(*GetStreamingDistributionConfigInput) GetStreamingDistributionConfigRequest
}

// Send marshals and sends the GetStreamingDistributionConfig API request.
func (r GetStreamingDistributionConfigRequest) Send(ctx context.Context) (*GetStreamingDistributionConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetStreamingDistributionConfigResponse{
		GetStreamingDistributionConfigOutput: r.Request.Data.(*GetStreamingDistributionConfigOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetStreamingDistributionConfigResponse is the response type for the
// GetStreamingDistributionConfig API operation.
type GetStreamingDistributionConfigResponse struct {
	*GetStreamingDistributionConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetStreamingDistributionConfig request.
func (r *GetStreamingDistributionConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
