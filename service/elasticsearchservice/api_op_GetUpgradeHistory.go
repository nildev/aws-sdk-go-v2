// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Container for request parameters to GetUpgradeHistory operation.
type GetUpgradeHistoryInput struct {
	_ struct{} `type:"structure"`

	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter
	// or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"DomainName" min:"3" type:"string" required:"true"`

	// Set this value to limit the number of results returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// Paginated APIs accepts NextToken input to returns next page results and provides
	// a NextToken output in the response which can be used by the client to retrieve
	// more results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetUpgradeHistoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUpgradeHistoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUpgradeHistoryInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetUpgradeHistoryInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DomainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Container for response returned by GetUpgradeHistory operation.
type GetUpgradeHistoryOutput struct {
	_ struct{} `type:"structure"`

	// Pagination token that needs to be supplied to the next call to get the next
	// page of results
	NextToken *string `json:"es:GetUpgradeHistoryOutput:NextToken" type:"string"`

	// A list of UpgradeHistory objects corresponding to each Upgrade or Upgrade
	// Eligibility Check performed on a domain returned as part of GetUpgradeHistoryResponse
	// object.
	UpgradeHistories []UpgradeHistory `json:"es:GetUpgradeHistoryOutput:UpgradeHistories" type:"list"`
}

// String returns the string representation
func (s GetUpgradeHistoryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetUpgradeHistoryOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpgradeHistories != nil {
		v := s.UpgradeHistories

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UpgradeHistories", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetUpgradeHistory = "GetUpgradeHistory"

// GetUpgradeHistoryRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Retrieves the complete history of the last 10 upgrades that were performed
// on the domain.
//
//    // Example sending a request using GetUpgradeHistoryRequest.
//    req := client.GetUpgradeHistoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetUpgradeHistoryRequest(input *GetUpgradeHistoryInput) GetUpgradeHistoryRequest {
	op := &aws.Operation{
		Name:       opGetUpgradeHistory,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/es/upgradeDomain/{DomainName}/history",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetUpgradeHistoryInput{}
	}

	req := c.newRequest(op, input, &GetUpgradeHistoryOutput{})
	return GetUpgradeHistoryRequest{Request: req, Input: input, Copy: c.GetUpgradeHistoryRequest}
}

// GetUpgradeHistoryRequest is the request type for the
// GetUpgradeHistory API operation.
type GetUpgradeHistoryRequest struct {
	*aws.Request
	Input *GetUpgradeHistoryInput
	Copy  func(*GetUpgradeHistoryInput) GetUpgradeHistoryRequest
}

// Send marshals and sends the GetUpgradeHistory API request.
func (r GetUpgradeHistoryRequest) Send(ctx context.Context) (*GetUpgradeHistoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUpgradeHistoryResponse{
		GetUpgradeHistoryOutput: r.Request.Data.(*GetUpgradeHistoryOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetUpgradeHistoryRequestPaginator returns a paginator for GetUpgradeHistory.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetUpgradeHistoryRequest(input)
//   p := elasticsearchservice.NewGetUpgradeHistoryRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetUpgradeHistoryPaginator(req GetUpgradeHistoryRequest) GetUpgradeHistoryPaginator {
	return GetUpgradeHistoryPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetUpgradeHistoryInput
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

// GetUpgradeHistoryPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetUpgradeHistoryPaginator struct {
	aws.Pager
}

func (p *GetUpgradeHistoryPaginator) CurrentPage() *GetUpgradeHistoryOutput {
	return p.Pager.CurrentPage().(*GetUpgradeHistoryOutput)
}

// GetUpgradeHistoryResponse is the response type for the
// GetUpgradeHistory API operation.
type GetUpgradeHistoryResponse struct {
	*GetUpgradeHistoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUpgradeHistory request.
func (r *GetUpgradeHistoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
