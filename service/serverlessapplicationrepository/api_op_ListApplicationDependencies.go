// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/ListApplicationDependenciesRequest
type ListApplicationDependenciesInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"applicationId" type:"string" required:"true"`

	MaxItems *int64 `location:"querystring" locationName:"maxItems" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	SemanticVersion *string `location:"querystring" locationName:"semanticVersion" type:"string"`
}

// String returns the string representation
func (s ListApplicationDependenciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListApplicationDependenciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListApplicationDependenciesInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListApplicationDependenciesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxItems", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SemanticVersion != nil {
		v := *s.SemanticVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "semanticVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/ListApplicationDependenciesResponse
type ListApplicationDependenciesOutput struct {
	_ struct{} `type:"structure"`

	Dependencies []ApplicationDependencySummary `locationName:"dependencies" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListApplicationDependenciesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListApplicationDependenciesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Dependencies != nil {
		v := s.Dependencies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dependencies", metadata)
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

const opListApplicationDependencies = "ListApplicationDependencies"

// ListApplicationDependenciesRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Retrieves the list of applications nested in the containing application.
//
//    // Example sending a request using ListApplicationDependenciesRequest.
//    req := client.ListApplicationDependenciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/ListApplicationDependencies
func (c *Client) ListApplicationDependenciesRequest(input *ListApplicationDependenciesInput) ListApplicationDependenciesRequest {
	op := &aws.Operation{
		Name:       opListApplicationDependencies,
		HTTPMethod: "GET",
		HTTPPath:   "/applications/{applicationId}/dependencies",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListApplicationDependenciesInput{}
	}

	req := c.newRequest(op, input, &ListApplicationDependenciesOutput{})
	return ListApplicationDependenciesRequest{Request: req, Input: input, Copy: c.ListApplicationDependenciesRequest}
}

// ListApplicationDependenciesRequest is the request type for the
// ListApplicationDependencies API operation.
type ListApplicationDependenciesRequest struct {
	*aws.Request
	Input *ListApplicationDependenciesInput
	Copy  func(*ListApplicationDependenciesInput) ListApplicationDependenciesRequest
}

// Send marshals and sends the ListApplicationDependencies API request.
func (r ListApplicationDependenciesRequest) Send(ctx context.Context) (*ListApplicationDependenciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListApplicationDependenciesResponse{
		ListApplicationDependenciesOutput: r.Request.Data.(*ListApplicationDependenciesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListApplicationDependenciesRequestPaginator returns a paginator for ListApplicationDependencies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListApplicationDependenciesRequest(input)
//   p := serverlessapplicationrepository.NewListApplicationDependenciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListApplicationDependenciesPaginator(req ListApplicationDependenciesRequest) ListApplicationDependenciesPaginator {
	return ListApplicationDependenciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListApplicationDependenciesInput
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

// ListApplicationDependenciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListApplicationDependenciesPaginator struct {
	aws.Pager
}

func (p *ListApplicationDependenciesPaginator) CurrentPage() *ListApplicationDependenciesOutput {
	return p.Pager.CurrentPage().(*ListApplicationDependenciesOutput)
}

// ListApplicationDependenciesResponse is the response type for the
// ListApplicationDependencies API operation.
type ListApplicationDependenciesResponse struct {
	*ListApplicationDependenciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListApplicationDependencies request.
func (r *ListApplicationDependenciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
