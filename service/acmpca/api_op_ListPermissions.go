// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/ListPermissionsRequest
type ListPermissionsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Number (ARN) of the private CA to inspect. You can find
	// the ARN by calling the ListCertificateAuthorities action. This must be of
	// the form: arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// You can get a private CA's ARN by running the ListCertificateAuthorities
	// action.
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`

	// When paginating results, use this parameter to specify the maximum number
	// of items to return in the response. If additional items exist beyond the
	// number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	MaxResults *int64 `min:"1" type:"integer"`

	// When paginating results, use this parameter in a subsequent request after
	// you receive a response with truncated results. Set it to the value of NextToken
	// from the response you just received.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListPermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPermissionsInput"}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/ListPermissionsResponse
type ListPermissionsOutput struct {
	_ struct{} `type:"structure"`

	// When the list is truncated, this value is present and should be used for
	// the NextToken parameter in a subsequent pagination request.
	NextToken *string `min:"1" type:"string"`

	// Summary information about each permission assigned by the specified private
	// CA, including the action enabled, the policy provided, and the time of creation.
	Permissions []Permission `type:"list"`
}

// String returns the string representation
func (s ListPermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPermissions = "ListPermissions"

// ListPermissionsRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Lists all the permissions, if any, that have been assigned by a private CA.
// Permissions can be granted with the CreatePermission action and revoked with
// the DeletePermission action.
//
//    // Example sending a request using ListPermissionsRequest.
//    req := client.ListPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/ListPermissions
func (c *Client) ListPermissionsRequest(input *ListPermissionsInput) ListPermissionsRequest {
	op := &aws.Operation{
		Name:       opListPermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPermissionsInput{}
	}

	req := c.newRequest(op, input, &ListPermissionsOutput{})
	return ListPermissionsRequest{Request: req, Input: input, Copy: c.ListPermissionsRequest}
}

// ListPermissionsRequest is the request type for the
// ListPermissions API operation.
type ListPermissionsRequest struct {
	*aws.Request
	Input *ListPermissionsInput
	Copy  func(*ListPermissionsInput) ListPermissionsRequest
}

// Send marshals and sends the ListPermissions API request.
func (r ListPermissionsRequest) Send(ctx context.Context) (*ListPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPermissionsResponse{
		ListPermissionsOutput: r.Request.Data.(*ListPermissionsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPermissionsRequestPaginator returns a paginator for ListPermissions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPermissionsRequest(input)
//   p := acmpca.NewListPermissionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPermissionsPaginator(req ListPermissionsRequest) ListPermissionsPaginator {
	return ListPermissionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPermissionsInput
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

// ListPermissionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPermissionsPaginator struct {
	aws.Pager
}

func (p *ListPermissionsPaginator) CurrentPage() *ListPermissionsOutput {
	return p.Pager.CurrentPage().(*ListPermissionsOutput)
}

// ListPermissionsResponse is the response type for the
// ListPermissions API operation.
type ListPermissionsResponse struct {
	*ListPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPermissions request.
func (r *ListPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
