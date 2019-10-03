// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListDirectoriesRequest
type ListDirectoriesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to retrieve.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The state of the directories in the list. Can be either Enabled, Disabled,
	// or Deleted.
	State DirectoryState `locationName:"state" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListDirectoriesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDirectoriesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDirectoriesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDirectoriesInput) MarshalFields(e protocol.FieldEncoder) error {

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
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListDirectoriesResponse
type ListDirectoriesOutput struct {
	_ struct{} `type:"structure"`

	// Lists all directories that are associated with your account in pagination
	// fashion.
	//
	// Directories is a required field
	Directories []Directory `json:"clouddirectory:ListDirectoriesOutput:Directories" type:"list" required:"true"`

	// The pagination token.
	NextToken *string `json:"clouddirectory:ListDirectoriesOutput:NextToken" type:"string"`
}

// String returns the string representation
func (s ListDirectoriesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDirectoriesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Directories != nil {
		v := s.Directories

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Directories", metadata)
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

const opListDirectories = "ListDirectories"

// ListDirectoriesRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Lists directories created within an account.
//
//    // Example sending a request using ListDirectoriesRequest.
//    req := client.ListDirectoriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListDirectories
func (c *Client) ListDirectoriesRequest(input *ListDirectoriesInput) ListDirectoriesRequest {
	op := &aws.Operation{
		Name:       opListDirectories,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/directory/list",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDirectoriesInput{}
	}

	req := c.newRequest(op, input, &ListDirectoriesOutput{})
	return ListDirectoriesRequest{Request: req, Input: input, Copy: c.ListDirectoriesRequest}
}

// ListDirectoriesRequest is the request type for the
// ListDirectories API operation.
type ListDirectoriesRequest struct {
	*aws.Request
	Input *ListDirectoriesInput
	Copy  func(*ListDirectoriesInput) ListDirectoriesRequest
}

// Send marshals and sends the ListDirectories API request.
func (r ListDirectoriesRequest) Send(ctx context.Context) (*ListDirectoriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDirectoriesResponse{
		ListDirectoriesOutput: r.Request.Data.(*ListDirectoriesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDirectoriesRequestPaginator returns a paginator for ListDirectories.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDirectoriesRequest(input)
//   p := clouddirectory.NewListDirectoriesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDirectoriesPaginator(req ListDirectoriesRequest) ListDirectoriesPaginator {
	return ListDirectoriesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDirectoriesInput
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

// ListDirectoriesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDirectoriesPaginator struct {
	aws.Pager
}

func (p *ListDirectoriesPaginator) CurrentPage() *ListDirectoriesOutput {
	return p.Pager.CurrentPage().(*ListDirectoriesOutput)
}

// ListDirectoriesResponse is the response type for the
// ListDirectories API operation.
type ListDirectoriesResponse struct {
	*ListDirectoriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDirectories request.
func (r *ListDirectoriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
