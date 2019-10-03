// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignVersionsRequest
type GetCampaignVersionsInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// CampaignId is a required field
	CampaignId *string `location:"uri" locationName:"campaign-id" type:"string" required:"true"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	Token *string `location:"querystring" locationName:"token" type:"string"`
}

// String returns the string representation
func (s GetCampaignVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCampaignVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCampaignVersionsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.CampaignId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CampaignId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCampaignVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CampaignId != nil {
		v := *s.CampaignId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "campaign-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "page-size", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignVersionsResponse
type GetCampaignVersionsOutput struct {
	_ struct{} `type:"structure" payload:"CampaignsResponse"`

	// Provides information about the configuration and other settings for all the
	// campaigns that are associated with an application.
	//
	// CampaignsResponse is a required field
	CampaignsResponse *CampaignsResponse `json:"pinpoint:GetCampaignVersionsOutput:CampaignsResponse" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetCampaignVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCampaignVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CampaignsResponse != nil {
		v := s.CampaignsResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CampaignsResponse", v, metadata)
	}
	return nil
}

const opGetCampaignVersions = "GetCampaignVersions"

// GetCampaignVersionsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status, configuration, and other settings
// for all versions of a specific campaign.
//
//    // Example sending a request using GetCampaignVersionsRequest.
//    req := client.GetCampaignVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetCampaignVersions
func (c *Client) GetCampaignVersionsRequest(input *GetCampaignVersionsInput) GetCampaignVersionsRequest {
	op := &aws.Operation{
		Name:       opGetCampaignVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/campaigns/{campaign-id}/versions",
	}

	if input == nil {
		input = &GetCampaignVersionsInput{}
	}

	req := c.newRequest(op, input, &GetCampaignVersionsOutput{})
	return GetCampaignVersionsRequest{Request: req, Input: input, Copy: c.GetCampaignVersionsRequest}
}

// GetCampaignVersionsRequest is the request type for the
// GetCampaignVersions API operation.
type GetCampaignVersionsRequest struct {
	*aws.Request
	Input *GetCampaignVersionsInput
	Copy  func(*GetCampaignVersionsInput) GetCampaignVersionsRequest
}

// Send marshals and sends the GetCampaignVersions API request.
func (r GetCampaignVersionsRequest) Send(ctx context.Context) (*GetCampaignVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCampaignVersionsResponse{
		GetCampaignVersionsOutput: r.Request.Data.(*GetCampaignVersionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCampaignVersionsResponse is the response type for the
// GetCampaignVersions API operation.
type GetCampaignVersionsResponse struct {
	*GetCampaignVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCampaignVersions request.
func (r *GetCampaignVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
