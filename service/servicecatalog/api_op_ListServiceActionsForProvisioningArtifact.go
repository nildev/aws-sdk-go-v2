// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListServiceActionsForProvisioningArtifactInput
type ListServiceActionsForProvisioningArtifactInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// The product identifier. For example, prod-abcdzk7xy33qa.
	//
	// ProductId is a required field
	ProductId *string `min:"1" type:"string" required:"true"`

	// The identifier of the provisioning artifact. For example, pa-4abcdjnxjj6ne.
	//
	// ProvisioningArtifactId is a required field
	ProvisioningArtifactId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListServiceActionsForProvisioningArtifactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListServiceActionsForProvisioningArtifactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListServiceActionsForProvisioningArtifactInput"}

	if s.ProductId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductId"))
	}
	if s.ProductId != nil && len(*s.ProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductId", 1))
	}

	if s.ProvisioningArtifactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisioningArtifactId"))
	}
	if s.ProvisioningArtifactId != nil && len(*s.ProvisioningArtifactId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisioningArtifactId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListServiceActionsForProvisioningArtifactOutput
type ListServiceActionsForProvisioningArtifactOutput struct {
	_ struct{} `type:"structure"`

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string `json:"servicecatalog:ListServiceActionsForProvisioningArtifactOutput:NextPageToken" type:"string"`

	// An object containing information about the self-service actions associated
	// with the provisioning artifact.
	ServiceActionSummaries []ServiceActionSummary `json:"servicecatalog:ListServiceActionsForProvisioningArtifactOutput:ServiceActionSummaries" type:"list"`
}

// String returns the string representation
func (s ListServiceActionsForProvisioningArtifactOutput) String() string {
	return awsutil.Prettify(s)
}

const opListServiceActionsForProvisioningArtifact = "ListServiceActionsForProvisioningArtifact"

// ListServiceActionsForProvisioningArtifactRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Returns a paginated list of self-service actions associated with the specified
// Product ID and Provisioning Artifact ID.
//
//    // Example sending a request using ListServiceActionsForProvisioningArtifactRequest.
//    req := client.ListServiceActionsForProvisioningArtifactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListServiceActionsForProvisioningArtifact
func (c *Client) ListServiceActionsForProvisioningArtifactRequest(input *ListServiceActionsForProvisioningArtifactInput) ListServiceActionsForProvisioningArtifactRequest {
	op := &aws.Operation{
		Name:       opListServiceActionsForProvisioningArtifact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"NextPageToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListServiceActionsForProvisioningArtifactInput{}
	}

	req := c.newRequest(op, input, &ListServiceActionsForProvisioningArtifactOutput{})
	return ListServiceActionsForProvisioningArtifactRequest{Request: req, Input: input, Copy: c.ListServiceActionsForProvisioningArtifactRequest}
}

// ListServiceActionsForProvisioningArtifactRequest is the request type for the
// ListServiceActionsForProvisioningArtifact API operation.
type ListServiceActionsForProvisioningArtifactRequest struct {
	*aws.Request
	Input *ListServiceActionsForProvisioningArtifactInput
	Copy  func(*ListServiceActionsForProvisioningArtifactInput) ListServiceActionsForProvisioningArtifactRequest
}

// Send marshals and sends the ListServiceActionsForProvisioningArtifact API request.
func (r ListServiceActionsForProvisioningArtifactRequest) Send(ctx context.Context) (*ListServiceActionsForProvisioningArtifactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServiceActionsForProvisioningArtifactResponse{
		ListServiceActionsForProvisioningArtifactOutput: r.Request.Data.(*ListServiceActionsForProvisioningArtifactOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListServiceActionsForProvisioningArtifactRequestPaginator returns a paginator for ListServiceActionsForProvisioningArtifact.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListServiceActionsForProvisioningArtifactRequest(input)
//   p := servicecatalog.NewListServiceActionsForProvisioningArtifactRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListServiceActionsForProvisioningArtifactPaginator(req ListServiceActionsForProvisioningArtifactRequest) ListServiceActionsForProvisioningArtifactPaginator {
	return ListServiceActionsForProvisioningArtifactPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListServiceActionsForProvisioningArtifactInput
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

// ListServiceActionsForProvisioningArtifactPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListServiceActionsForProvisioningArtifactPaginator struct {
	aws.Pager
}

func (p *ListServiceActionsForProvisioningArtifactPaginator) CurrentPage() *ListServiceActionsForProvisioningArtifactOutput {
	return p.Pager.CurrentPage().(*ListServiceActionsForProvisioningArtifactOutput)
}

// ListServiceActionsForProvisioningArtifactResponse is the response type for the
// ListServiceActionsForProvisioningArtifact API operation.
type ListServiceActionsForProvisioningArtifactResponse struct {
	*ListServiceActionsForProvisioningArtifactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServiceActionsForProvisioningArtifact request.
func (r *ListServiceActionsForProvisioningArtifactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
