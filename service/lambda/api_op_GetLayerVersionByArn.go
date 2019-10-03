// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetLayerVersionByArnRequest
type GetLayerVersionByArnInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the layer version.
	//
	// Arn is a required field
	Arn *string `location:"querystring" locationName:"Arn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLayerVersionByArnInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLayerVersionByArnInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLayerVersionByArnInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLayerVersionByArnInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetLayerVersionResponse
type GetLayerVersionByArnOutput struct {
	_ struct{} `type:"structure"`

	// The layer's compatible runtimes.
	CompatibleRuntimes []Runtime `json:"lambda:GetLayerVersionByArnOutput:CompatibleRuntimes" type:"list"`

	// Details about the layer version.
	Content *LayerVersionContentOutput `json:"lambda:GetLayerVersionByArnOutput:Content" type:"structure"`

	// The date that the layer version was created, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	CreatedDate *string `json:"lambda:GetLayerVersionByArnOutput:CreatedDate" type:"string"`

	// The description of the version.
	Description *string `json:"lambda:GetLayerVersionByArnOutput:Description" type:"string"`

	// The ARN of the layer.
	LayerArn *string `json:"lambda:GetLayerVersionByArnOutput:LayerArn" min:"1" type:"string"`

	// The ARN of the layer version.
	LayerVersionArn *string `json:"lambda:GetLayerVersionByArnOutput:LayerVersionArn" min:"1" type:"string"`

	// The layer's software license.
	LicenseInfo *string `json:"lambda:GetLayerVersionByArnOutput:LicenseInfo" type:"string"`

	// The version number.
	Version *int64 `json:"lambda:GetLayerVersionByArnOutput:Version" type:"long"`
}

// String returns the string representation
func (s GetLayerVersionByArnOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLayerVersionByArnOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CompatibleRuntimes != nil {
		v := s.CompatibleRuntimes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "CompatibleRuntimes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Content != nil {
		v := s.Content

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Content", v, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatedDate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LayerArn != nil {
		v := *s.LayerArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LayerArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LayerVersionArn != nil {
		v := *s.LayerVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LayerVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LicenseInfo != nil {
		v := *s.LicenseInfo

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LicenseInfo", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opGetLayerVersionByArn = "GetLayerVersionByArn"

// GetLayerVersionByArnRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns information about a version of an AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html),
// with a link to download the layer archive that's valid for 10 minutes.
//
//    // Example sending a request using GetLayerVersionByArnRequest.
//    req := client.GetLayerVersionByArnRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetLayerVersionByArn
func (c *Client) GetLayerVersionByArnRequest(input *GetLayerVersionByArnInput) GetLayerVersionByArnRequest {
	op := &aws.Operation{
		Name:       opGetLayerVersionByArn,
		HTTPMethod: "GET",
		HTTPPath:   "/2018-10-31/layers?find=LayerVersion",
	}

	if input == nil {
		input = &GetLayerVersionByArnInput{}
	}

	req := c.newRequest(op, input, &GetLayerVersionByArnOutput{})
	return GetLayerVersionByArnRequest{Request: req, Input: input, Copy: c.GetLayerVersionByArnRequest}
}

// GetLayerVersionByArnRequest is the request type for the
// GetLayerVersionByArn API operation.
type GetLayerVersionByArnRequest struct {
	*aws.Request
	Input *GetLayerVersionByArnInput
	Copy  func(*GetLayerVersionByArnInput) GetLayerVersionByArnRequest
}

// Send marshals and sends the GetLayerVersionByArn API request.
func (r GetLayerVersionByArnRequest) Send(ctx context.Context) (*GetLayerVersionByArnResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLayerVersionByArnResponse{
		GetLayerVersionByArnOutput: r.Request.Data.(*GetLayerVersionByArnOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLayerVersionByArnResponse is the response type for the
// GetLayerVersionByArn API operation.
type GetLayerVersionByArnResponse struct {
	*GetLayerVersionByArnOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLayerVersionByArn request.
func (r *GetLayerVersionByArnResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
