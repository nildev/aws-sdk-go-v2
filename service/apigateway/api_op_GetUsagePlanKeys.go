// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The GET request to get all the usage plan keys representing the API keys
// added to a specified usage plan.
type GetUsagePlanKeysInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// A query parameter specifying the name of the to-be-returned usage plan keys.
	NameQuery *string `location:"querystring" locationName:"name" type:"string"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`

	// [Required] The Id of the UsagePlan resource representing the usage plan containing
	// the to-be-retrieved UsagePlanKey resource representing a plan customer.
	//
	// UsagePlanId is a required field
	UsagePlanId *string `location:"uri" locationName:"usageplanId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetUsagePlanKeysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUsagePlanKeysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUsagePlanKeysInput"}

	if s.UsagePlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UsagePlanId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetUsagePlanKeysInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.UsagePlanId != nil {
		v := *s.UsagePlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "usageplanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.NameQuery != nil {
		v := *s.NameQuery

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents the collection of usage plan keys added to usage plans for the
// associated API keys and, possibly, other types of keys.
//
// Create and Use Usage Plans (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type GetUsagePlanKeysOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []UsagePlanKey `json:"apigateway:GetUsagePlanKeysOutput:Items" locationName:"item" type:"list"`

	Position *string `json:"apigateway:GetUsagePlanKeysOutput:Position" locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetUsagePlanKeysOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetUsagePlanKeysOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "item", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetUsagePlanKeys = "GetUsagePlanKeys"

// GetUsagePlanKeysRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets all the usage plan keys representing the API keys added to a specified
// usage plan.
//
//    // Example sending a request using GetUsagePlanKeysRequest.
//    req := client.GetUsagePlanKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetUsagePlanKeysRequest(input *GetUsagePlanKeysInput) GetUsagePlanKeysRequest {
	op := &aws.Operation{
		Name:       opGetUsagePlanKeys,
		HTTPMethod: "GET",
		HTTPPath:   "/usageplans/{usageplanId}/keys",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"position"},
			OutputTokens:    []string{"position"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetUsagePlanKeysInput{}
	}

	req := c.newRequest(op, input, &GetUsagePlanKeysOutput{})
	return GetUsagePlanKeysRequest{Request: req, Input: input, Copy: c.GetUsagePlanKeysRequest}
}

// GetUsagePlanKeysRequest is the request type for the
// GetUsagePlanKeys API operation.
type GetUsagePlanKeysRequest struct {
	*aws.Request
	Input *GetUsagePlanKeysInput
	Copy  func(*GetUsagePlanKeysInput) GetUsagePlanKeysRequest
}

// Send marshals and sends the GetUsagePlanKeys API request.
func (r GetUsagePlanKeysRequest) Send(ctx context.Context) (*GetUsagePlanKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUsagePlanKeysResponse{
		GetUsagePlanKeysOutput: r.Request.Data.(*GetUsagePlanKeysOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetUsagePlanKeysRequestPaginator returns a paginator for GetUsagePlanKeys.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetUsagePlanKeysRequest(input)
//   p := apigateway.NewGetUsagePlanKeysRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetUsagePlanKeysPaginator(req GetUsagePlanKeysRequest) GetUsagePlanKeysPaginator {
	return GetUsagePlanKeysPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetUsagePlanKeysInput
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

// GetUsagePlanKeysPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetUsagePlanKeysPaginator struct {
	aws.Pager
}

func (p *GetUsagePlanKeysPaginator) CurrentPage() *GetUsagePlanKeysOutput {
	return p.Pager.CurrentPage().(*GetUsagePlanKeysOutput)
}

// GetUsagePlanKeysResponse is the response type for the
// GetUsagePlanKeys API operation.
type GetUsagePlanKeysResponse struct {
	*GetUsagePlanKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUsagePlanKeys request.
func (r *GetUsagePlanKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
