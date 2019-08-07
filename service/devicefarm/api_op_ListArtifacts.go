// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the list artifacts operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListArtifactsRequest
type ListArtifactsInput struct {
	_ struct{} `type:"structure"`

	// The Run, Job, Suite, or Test ARN.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// The artifacts' type.
	//
	// Allowed values include:
	//
	//    * FILE: The artifacts are files.
	//
	//    * LOG: The artifacts are logs.
	//
	//    * SCREENSHOT: The artifacts are screenshots.
	//
	// Type is a required field
	Type ArtifactCategory `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ListArtifactsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListArtifactsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListArtifactsInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a list artifacts operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListArtifactsResult
type ListArtifactsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the artifacts.
	Artifacts []Artifact `json:"devicefarm:ListArtifactsOutput:Artifacts" locationName:"artifacts" type:"list"`

	// If the number of items that are returned is significantly large, this is
	// an identifier that is also returned, which can be used in a subsequent call
	// to this operation to return the next set of items in the list.
	NextToken *string `json:"devicefarm:ListArtifactsOutput:NextToken" locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListArtifactsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListArtifacts = "ListArtifacts"

// ListArtifactsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about artifacts.
//
//    // Example sending a request using ListArtifactsRequest.
//    req := client.ListArtifactsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListArtifacts
func (c *Client) ListArtifactsRequest(input *ListArtifactsInput) ListArtifactsRequest {
	op := &aws.Operation{
		Name:       opListArtifacts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListArtifactsInput{}
	}

	req := c.newRequest(op, input, &ListArtifactsOutput{})
	return ListArtifactsRequest{Request: req, Input: input, Copy: c.ListArtifactsRequest}
}

// ListArtifactsRequest is the request type for the
// ListArtifacts API operation.
type ListArtifactsRequest struct {
	*aws.Request
	Input *ListArtifactsInput
	Copy  func(*ListArtifactsInput) ListArtifactsRequest
}

// Send marshals and sends the ListArtifacts API request.
func (r ListArtifactsRequest) Send(ctx context.Context) (*ListArtifactsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListArtifactsResponse{
		ListArtifactsOutput: r.Request.Data.(*ListArtifactsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListArtifactsRequestPaginator returns a paginator for ListArtifacts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListArtifactsRequest(input)
//   p := devicefarm.NewListArtifactsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListArtifactsPaginator(req ListArtifactsRequest) ListArtifactsPaginator {
	return ListArtifactsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListArtifactsInput
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

// ListArtifactsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListArtifactsPaginator struct {
	aws.Pager
}

func (p *ListArtifactsPaginator) CurrentPage() *ListArtifactsOutput {
	return p.Pager.CurrentPage().(*ListArtifactsOutput)
}

// ListArtifactsResponse is the response type for the
// ListArtifacts API operation.
type ListArtifactsResponse struct {
	*ListArtifactsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListArtifacts request.
func (r *ListArtifactsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
