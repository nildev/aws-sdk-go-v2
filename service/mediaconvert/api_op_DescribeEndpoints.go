// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Send an request with an empty body to the regional API endpoint to get your
// account API endpoint.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/DescribeEndpointsRequest
type DescribeEndpointsInput struct {
	_ struct{} `type:"structure"`

	// Optional. Max number of endpoints, up to twenty, that will be returned at
	// one time.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// Optional field, defaults to DEFAULT. Specify DEFAULT for this operation to
	// return your endpoints if any exist, or to create an endpoint for you and
	// return it if one doesn't already exist. Specify GET_ONLY to return your endpoints
	// if any exist, or an empty list if none exist.
	Mode DescribeEndpointsMode `locationName:"mode" type:"string" enum:"true"`

	// Use this string, provided with the response to a previous request, to request
	// the next batch of endpoints.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeEndpointsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if len(s.Mode) > 0 {
		v := s.Mode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "mode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Successful describe endpoints requests will return your account API endpoint.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/DescribeEndpointsResponse
type DescribeEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// List of endpoints
	Endpoints []Endpoint `json:"mediaconvert:DescribeEndpointsOutput:Endpoints" locationName:"endpoints" type:"list"`

	// Use this string to request the next batch of endpoints.
	NextToken *string `json:"mediaconvert:DescribeEndpointsOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeEndpointsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Endpoints != nil {
		v := s.Endpoints

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "endpoints", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeEndpoints = "DescribeEndpoints"

// DescribeEndpointsRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Send an request with an empty body to the regional API endpoint to get your
// account API endpoint.
//
//    // Example sending a request using DescribeEndpointsRequest.
//    req := client.DescribeEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/DescribeEndpoints
func (c *Client) DescribeEndpointsRequest(input *DescribeEndpointsInput) DescribeEndpointsRequest {
	op := &aws.Operation{
		Name:       opDescribeEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/2017-08-29/endpoints",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEndpointsInput{}
	}

	req := c.newRequest(op, input, &DescribeEndpointsOutput{})
	return DescribeEndpointsRequest{Request: req, Input: input, Copy: c.DescribeEndpointsRequest}
}

// DescribeEndpointsRequest is the request type for the
// DescribeEndpoints API operation.
type DescribeEndpointsRequest struct {
	*aws.Request
	Input *DescribeEndpointsInput
	Copy  func(*DescribeEndpointsInput) DescribeEndpointsRequest
}

// Send marshals and sends the DescribeEndpoints API request.
func (r DescribeEndpointsRequest) Send(ctx context.Context) (*DescribeEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEndpointsResponse{
		DescribeEndpointsOutput: r.Request.Data.(*DescribeEndpointsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEndpointsRequestPaginator returns a paginator for DescribeEndpoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEndpointsRequest(input)
//   p := mediaconvert.NewDescribeEndpointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEndpointsPaginator(req DescribeEndpointsRequest) DescribeEndpointsPaginator {
	return DescribeEndpointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEndpointsInput
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

// DescribeEndpointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEndpointsPaginator struct {
	aws.Pager
}

func (p *DescribeEndpointsPaginator) CurrentPage() *DescribeEndpointsOutput {
	return p.Pager.CurrentPage().(*DescribeEndpointsOutput)
}

// DescribeEndpointsResponse is the response type for the
// DescribeEndpoints API operation.
type DescribeEndpointsResponse struct {
	*DescribeEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEndpoints request.
func (r *DescribeEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
