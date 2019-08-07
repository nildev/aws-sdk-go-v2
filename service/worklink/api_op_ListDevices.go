// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/ListDevicesRequest
type ListDevicesInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`

	// The maximum number of results to be included in the next page.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token used to retrieve the next page of results for this operation.
	// If this value is null, it retrieves the first page.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDevicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDevicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDevicesInput"}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDevicesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/ListDevicesResponse
type ListDevicesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the devices.
	Devices []DeviceSummary `json:"worklink:ListDevicesOutput:Devices" type:"list"`

	// The pagination token used to retrieve the next page of results for this operation.
	// If there are no more pages, this value is null.
	NextToken *string `json:"worklink:ListDevicesOutput:NextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListDevicesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDevicesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Devices != nil {
		v := s.Devices

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Devices", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDevices = "ListDevices"

// ListDevicesRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Retrieves a list of devices registered with the specified fleet.
//
//    // Example sending a request using ListDevicesRequest.
//    req := client.ListDevicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/ListDevices
func (c *Client) ListDevicesRequest(input *ListDevicesInput) ListDevicesRequest {
	op := &aws.Operation{
		Name:       opListDevices,
		HTTPMethod: "POST",
		HTTPPath:   "/listDevices",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDevicesInput{}
	}

	req := c.newRequest(op, input, &ListDevicesOutput{})
	return ListDevicesRequest{Request: req, Input: input, Copy: c.ListDevicesRequest}
}

// ListDevicesRequest is the request type for the
// ListDevices API operation.
type ListDevicesRequest struct {
	*aws.Request
	Input *ListDevicesInput
	Copy  func(*ListDevicesInput) ListDevicesRequest
}

// Send marshals and sends the ListDevices API request.
func (r ListDevicesRequest) Send(ctx context.Context) (*ListDevicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDevicesResponse{
		ListDevicesOutput: r.Request.Data.(*ListDevicesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDevicesRequestPaginator returns a paginator for ListDevices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDevicesRequest(input)
//   p := worklink.NewListDevicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDevicesPaginator(req ListDevicesRequest) ListDevicesPaginator {
	return ListDevicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDevicesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListDevicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDevicesPaginator struct {
	aws.Pager
}

func (p *ListDevicesPaginator) CurrentPage() *ListDevicesOutput {
	return p.Pager.CurrentPage().(*ListDevicesOutput)
}

// ListDevicesResponse is the response type for the
// ListDevices API operation.
type ListDevicesResponse struct {
	*ListDevicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDevices request.
func (r *ListDevicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
