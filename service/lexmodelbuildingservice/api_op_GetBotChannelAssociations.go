// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotChannelAssociationsRequest
type GetBotChannelAssociationsInput struct {
	_ struct{} `type:"structure"`

	// An alias pointing to the specific version of the Amazon Lex bot to which
	// this association is being made.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"aliasName" min:"1" type:"string" required:"true"`

	// The name of the Amazon Lex bot in the association.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" min:"2" type:"string" required:"true"`

	// The maximum number of associations to return in the response. The default
	// is 50.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Substring to match in channel association names. An association will be returned
	// if any part of its name matches the substring. For example, "xyz" matches
	// both "xyzabc" and "abcxyz." To return all bot channel associations, use a
	// hyphen ("-") as the nameContains parameter.
	NameContains *string `location:"querystring" locationName:"nameContains" min:"1" type:"string"`

	// A pagination token for fetching the next page of associations. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the response.
	// To fetch the next page of associations, specify the pagination token in the
	// next request.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetBotChannelAssociationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBotChannelAssociationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBotChannelAssociationsInput"}

	if s.BotAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotAlias"))
	}
	if s.BotAlias != nil && len(*s.BotAlias) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BotAlias", 1))
	}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}
	if s.BotName != nil && len(*s.BotName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("BotName", 2))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBotChannelAssociationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BotAlias != nil {
		v := *s.BotAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "aliasName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NameContains != nil {
		v := *s.NameContains

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nameContains", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotChannelAssociationsResponse
type GetBotChannelAssociationsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects, one for each association, that provides information
	// about the Amazon Lex bot and its association with the channel.
	BotChannelAssociations []BotChannelAssociation `json:"models.lex:GetBotChannelAssociationsOutput:BotChannelAssociations" locationName:"botChannelAssociations" type:"list"`

	// A pagination token that fetches the next page of associations. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the response.
	// To fetch the next page of associations, specify the pagination token in the
	// next request.
	NextToken *string `json:"models.lex:GetBotChannelAssociationsOutput:NextToken" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetBotChannelAssociationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBotChannelAssociationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BotChannelAssociations != nil {
		v := s.BotChannelAssociations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "botChannelAssociations", metadata)
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

const opGetBotChannelAssociations = "GetBotChannelAssociations"

// GetBotChannelAssociationsRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Returns a list of all of the channels associated with the specified bot.
//
// The GetBotChannelAssociations operation requires permissions for the lex:GetBotChannelAssociations
// action.
//
//    // Example sending a request using GetBotChannelAssociationsRequest.
//    req := client.GetBotChannelAssociationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotChannelAssociations
func (c *Client) GetBotChannelAssociationsRequest(input *GetBotChannelAssociationsInput) GetBotChannelAssociationsRequest {
	op := &aws.Operation{
		Name:       opGetBotChannelAssociations,
		HTTPMethod: "GET",
		HTTPPath:   "/bots/{botName}/aliases/{aliasName}/channels/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetBotChannelAssociationsInput{}
	}

	req := c.newRequest(op, input, &GetBotChannelAssociationsOutput{})
	return GetBotChannelAssociationsRequest{Request: req, Input: input, Copy: c.GetBotChannelAssociationsRequest}
}

// GetBotChannelAssociationsRequest is the request type for the
// GetBotChannelAssociations API operation.
type GetBotChannelAssociationsRequest struct {
	*aws.Request
	Input *GetBotChannelAssociationsInput
	Copy  func(*GetBotChannelAssociationsInput) GetBotChannelAssociationsRequest
}

// Send marshals and sends the GetBotChannelAssociations API request.
func (r GetBotChannelAssociationsRequest) Send(ctx context.Context) (*GetBotChannelAssociationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBotChannelAssociationsResponse{
		GetBotChannelAssociationsOutput: r.Request.Data.(*GetBotChannelAssociationsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetBotChannelAssociationsRequestPaginator returns a paginator for GetBotChannelAssociations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetBotChannelAssociationsRequest(input)
//   p := lexmodelbuildingservice.NewGetBotChannelAssociationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetBotChannelAssociationsPaginator(req GetBotChannelAssociationsRequest) GetBotChannelAssociationsPaginator {
	return GetBotChannelAssociationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetBotChannelAssociationsInput
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

// GetBotChannelAssociationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetBotChannelAssociationsPaginator struct {
	aws.Pager
}

func (p *GetBotChannelAssociationsPaginator) CurrentPage() *GetBotChannelAssociationsOutput {
	return p.Pager.CurrentPage().(*GetBotChannelAssociationsOutput)
}

// GetBotChannelAssociationsResponse is the response type for the
// GetBotChannelAssociations API operation.
type GetBotChannelAssociationsResponse struct {
	*GetBotChannelAssociationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBotChannelAssociations request.
func (r *GetBotChannelAssociationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
