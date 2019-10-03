// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediatailor

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/GetPlaybackConfigurationRequest
type GetPlaybackConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Name is a required field
	Name *string `location:"uri" locationName:"Name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPlaybackConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPlaybackConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPlaybackConfigurationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPlaybackConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/GetPlaybackConfigurationResponse
type GetPlaybackConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The URL for the ad decision server (ADS). This includes the specification
	// of static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing, you can provide a
	// static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string `json:"api.mediatailor:GetPlaybackConfigurationOutput:AdDecisionServerUrl" type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `json:"api.mediatailor:GetPlaybackConfigurationOutput:CdnConfiguration" type:"structure"`

	// The configuration for DASH content.
	DashConfiguration *DashConfiguration `json:"api.mediatailor:GetPlaybackConfigurationOutput:DashConfiguration" type:"structure"`

	// The configuration for HLS content.
	HlsConfiguration *HlsConfiguration `json:"api.mediatailor:GetPlaybackConfigurationOutput:HlsConfiguration" type:"structure"`

	// The identifier for the playback configuration.
	Name *string `json:"api.mediatailor:GetPlaybackConfigurationOutput:Name" type:"string"`

	// The Amazon Resource Name (ARN) for the playback configuration.
	PlaybackConfigurationArn *string `json:"api.mediatailor:GetPlaybackConfigurationOutput:PlaybackConfigurationArn" type:"string"`

	// The URL that the player accesses to get a manifest from AWS Elemental MediaTailor.
	// This session will use server-side reporting.
	PlaybackEndpointPrefix *string `json:"api.mediatailor:GetPlaybackConfigurationOutput:PlaybackEndpointPrefix" type:"string"`

	// The URL that the player uses to initialize a session that uses client-side
	// reporting.
	SessionInitializationEndpointPrefix *string `json:"api.mediatailor:GetPlaybackConfigurationOutput:SessionInitializationEndpointPrefix" type:"string"`

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill
	// in gaps in media content. Configuring the slate is optional for non-VPAID
	// playback configurations. For VPAID, the slate is required because MediaTailor
	// provides it in the slots designated for dynamic ad content. The slate must
	// be a high-quality asset that contains both audio and video.
	SlateAdUrl *string `json:"api.mediatailor:GetPlaybackConfigurationOutput:SlateAdUrl" type:"string"`

	// The tags assigned to the playback configuration.
	Tags map[string]string `json:"api.mediatailor:GetPlaybackConfigurationOutput:Tags" locationName:"tags" type:"map"`

	// The name that is used to associate this playback configuration with a custom
	// transcode profile. This overrides the dynamic transcoding defaults of MediaTailor.
	// Use this only if you have already set up custom profiles with the help of
	// AWS Support.
	TranscodeProfileName *string `json:"api.mediatailor:GetPlaybackConfigurationOutput:TranscodeProfileName" type:"string"`

	// The URL prefix for the master playlist for the stream, minus the asset ID.
	// The maximum length is 512 characters.
	VideoContentSourceUrl *string `json:"api.mediatailor:GetPlaybackConfigurationOutput:VideoContentSourceUrl" type:"string"`
}

// String returns the string representation
func (s GetPlaybackConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPlaybackConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdDecisionServerUrl != nil {
		v := *s.AdDecisionServerUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AdDecisionServerUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CdnConfiguration != nil {
		v := s.CdnConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CdnConfiguration", v, metadata)
	}
	if s.DashConfiguration != nil {
		v := s.DashConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DashConfiguration", v, metadata)
	}
	if s.HlsConfiguration != nil {
		v := s.HlsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HlsConfiguration", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlaybackConfigurationArn != nil {
		v := *s.PlaybackConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PlaybackConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlaybackEndpointPrefix != nil {
		v := *s.PlaybackEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PlaybackEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SessionInitializationEndpointPrefix != nil {
		v := *s.SessionInitializationEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SessionInitializationEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlateAdUrl != nil {
		v := *s.SlateAdUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SlateAdUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.TranscodeProfileName != nil {
		v := *s.TranscodeProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TranscodeProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VideoContentSourceUrl != nil {
		v := *s.VideoContentSourceUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VideoContentSourceUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetPlaybackConfiguration = "GetPlaybackConfiguration"

// GetPlaybackConfigurationRequest returns a request value for making API operation for
// AWS MediaTailor.
//
// Returns the playback configuration for the specified name.
//
//    // Example sending a request using GetPlaybackConfigurationRequest.
//    req := client.GetPlaybackConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/GetPlaybackConfiguration
func (c *Client) GetPlaybackConfigurationRequest(input *GetPlaybackConfigurationInput) GetPlaybackConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetPlaybackConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/playbackConfiguration/{Name}",
	}

	if input == nil {
		input = &GetPlaybackConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetPlaybackConfigurationOutput{})
	return GetPlaybackConfigurationRequest{Request: req, Input: input, Copy: c.GetPlaybackConfigurationRequest}
}

// GetPlaybackConfigurationRequest is the request type for the
// GetPlaybackConfiguration API operation.
type GetPlaybackConfigurationRequest struct {
	*aws.Request
	Input *GetPlaybackConfigurationInput
	Copy  func(*GetPlaybackConfigurationInput) GetPlaybackConfigurationRequest
}

// Send marshals and sends the GetPlaybackConfiguration API request.
func (r GetPlaybackConfigurationRequest) Send(ctx context.Context) (*GetPlaybackConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPlaybackConfigurationResponse{
		GetPlaybackConfigurationOutput: r.Request.Data.(*GetPlaybackConfigurationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPlaybackConfigurationResponse is the response type for the
// GetPlaybackConfiguration API operation.
type GetPlaybackConfigurationResponse struct {
	*GetPlaybackConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPlaybackConfiguration request.
func (r *GetPlaybackConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
