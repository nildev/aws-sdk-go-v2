// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to obtain deliverability metrics for a domain.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetDomainStatisticsReportRequest
type GetDomainStatisticsReportInput struct {
	_ struct{} `type:"structure"`

	// The domain that you want to obtain deliverability metrics for.
	//
	// Domain is a required field
	Domain *string `location:"uri" locationName:"Domain" type:"string" required:"true"`

	// The last day (in Unix time) that you want to obtain domain deliverability
	// metrics for. The EndDate that you specify has to be less than or equal to
	// 30 days after the StartDate.
	//
	// EndDate is a required field
	EndDate *time.Time `location:"querystring" locationName:"EndDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The first day (in Unix time) that you want to obtain domain deliverability
	// metrics for.
	//
	// StartDate is a required field
	StartDate *time.Time `location:"querystring" locationName:"StartDate" type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s GetDomainStatisticsReportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDomainStatisticsReportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDomainStatisticsReportInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}

	if s.EndDate == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndDate"))
	}

	if s.StartDate == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartDate"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainStatisticsReportInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Domain != nil {
		v := *s.Domain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Domain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndDate != nil {
		v := *s.EndDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "EndDate", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	if s.StartDate != nil {
		v := *s.StartDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "StartDate", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	return nil
}

// An object that includes statistics that are related to the domain that you
// specified.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetDomainStatisticsReportResponse
type GetDomainStatisticsReportOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains deliverability metrics for the domain that you specified.
	// This object contains data for each day, starting on the StartDate and ending
	// on the EndDate.
	//
	// DailyVolumes is a required field
	DailyVolumes []DailyVolume `json:"email:GetDomainStatisticsReportOutput:DailyVolumes" type:"list" required:"true"`

	// An object that contains deliverability metrics for the domain that you specified.
	// The data in this object is a summary of all of the data that was collected
	// from the StartDate to the EndDate.
	//
	// OverallVolume is a required field
	OverallVolume *OverallVolume `json:"email:GetDomainStatisticsReportOutput:OverallVolume" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetDomainStatisticsReportOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainStatisticsReportOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DailyVolumes != nil {
		v := s.DailyVolumes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "DailyVolumes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.OverallVolume != nil {
		v := s.OverallVolume

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "OverallVolume", v, metadata)
	}
	return nil
}

const opGetDomainStatisticsReport = "GetDomainStatisticsReport"

// GetDomainStatisticsReportRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Retrieve inbox placement and engagement rates for the domains that you use
// to send email.
//
//    // Example sending a request using GetDomainStatisticsReportRequest.
//    req := client.GetDomainStatisticsReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetDomainStatisticsReport
func (c *Client) GetDomainStatisticsReportRequest(input *GetDomainStatisticsReportInput) GetDomainStatisticsReportRequest {
	op := &aws.Operation{
		Name:       opGetDomainStatisticsReport,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/deliverability-dashboard/statistics-report/{Domain}",
	}

	if input == nil {
		input = &GetDomainStatisticsReportInput{}
	}

	req := c.newRequest(op, input, &GetDomainStatisticsReportOutput{})
	return GetDomainStatisticsReportRequest{Request: req, Input: input, Copy: c.GetDomainStatisticsReportRequest}
}

// GetDomainStatisticsReportRequest is the request type for the
// GetDomainStatisticsReport API operation.
type GetDomainStatisticsReportRequest struct {
	*aws.Request
	Input *GetDomainStatisticsReportInput
	Copy  func(*GetDomainStatisticsReportInput) GetDomainStatisticsReportRequest
}

// Send marshals and sends the GetDomainStatisticsReport API request.
func (r GetDomainStatisticsReportRequest) Send(ctx context.Context) (*GetDomainStatisticsReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainStatisticsReportResponse{
		GetDomainStatisticsReportOutput: r.Request.Data.(*GetDomainStatisticsReportOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDomainStatisticsReportResponse is the response type for the
// GetDomainStatisticsReport API operation.
type GetDomainStatisticsReportResponse struct {
	*GetDomainStatisticsReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainStatisticsReport request.
func (r *GetDomainStatisticsReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
