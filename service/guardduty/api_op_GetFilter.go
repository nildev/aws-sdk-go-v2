// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetFilterRequest
type GetFilterInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector the filter is associated with.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The name of the filter you want to get.
	//
	// FilterName is a required field
	FilterName *string `location:"uri" locationName:"filterName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFilterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFilterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFilterInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.FilterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FilterName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFilterInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FilterName != nil {
		v := *s.FilterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "filterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetFilterResponse
type GetFilterOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the action that is to be applied to the findings that match the
	// filter.
	//
	// Action is a required field
	Action FilterAction `locationName:"action" min:"1" type:"string" required:"true" enum:"true"`

	// The description of the filter.
	Description *string `locationName:"description" type:"string"`

	// Represents the criteria to be used in the filter for querying findings.
	//
	// FindingCriteria is a required field
	FindingCriteria *FindingCriteria `locationName:"findingCriteria" type:"structure" required:"true"`

	// The name of the filter.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"3" type:"string" required:"true"`

	// Specifies the position of the filter in the list of current filters. Also
	// specifies the order in which this filter is applied to the findings.
	Rank *int64 `locationName:"rank" min:"1" type:"integer"`

	// The tags of the filter resource.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s GetFilterOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFilterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Action) > 0 {
		v := s.Action

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "action", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FindingCriteria != nil {
		v := s.FindingCriteria

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "findingCriteria", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Rank != nil {
		v := *s.Rank

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "rank", protocol.Int64Value(v), metadata)
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

const opGetFilter = "GetFilter"

// GetFilterRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Returns the details of the filter specified by the filter name.
//
//    // Example sending a request using GetFilterRequest.
//    req := client.GetFilterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetFilter
func (c *Client) GetFilterRequest(input *GetFilterInput) GetFilterRequest {
	op := &aws.Operation{
		Name:       opGetFilter,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}/filter/{filterName}",
	}

	if input == nil {
		input = &GetFilterInput{}
	}

	req := c.newRequest(op, input, &GetFilterOutput{})
	return GetFilterRequest{Request: req, Input: input, Copy: c.GetFilterRequest}
}

// GetFilterRequest is the request type for the
// GetFilter API operation.
type GetFilterRequest struct {
	*aws.Request
	Input *GetFilterInput
	Copy  func(*GetFilterInput) GetFilterRequest
}

// Send marshals and sends the GetFilter API request.
func (r GetFilterRequest) Send(ctx context.Context) (*GetFilterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFilterResponse{
		GetFilterOutput: r.Request.Data.(*GetFilterOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFilterResponse is the response type for the
// GetFilter API operation.
type GetFilterResponse struct {
	*GetFilterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFilter request.
func (r *GetFilterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
