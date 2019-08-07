// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/TerminateProvisionedProductInput
type TerminateProvisionedProductInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// If set to true, AWS Service Catalog stops managing the specified provisioned
	// product even if it cannot delete the underlying resources.
	IgnoreErrors *bool `type:"boolean"`

	// The identifier of the provisioned product. You cannot specify both ProvisionedProductName
	// and ProvisionedProductId.
	ProvisionedProductId *string `min:"1" type:"string"`

	// The name of the provisioned product. You cannot specify both ProvisionedProductName
	// and ProvisionedProductId.
	ProvisionedProductName *string `min:"1" type:"string"`

	// An idempotency token that uniquely identifies the termination request. This
	// token is only valid during the termination process. After the provisioned
	// product is terminated, subsequent requests to terminate the same provisioned
	// product always return ResourceNotFound.
	//
	// TerminateToken is a required field
	TerminateToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`
}

// String returns the string representation
func (s TerminateProvisionedProductInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminateProvisionedProductInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminateProvisionedProductInput"}
	if s.ProvisionedProductId != nil && len(*s.ProvisionedProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisionedProductId", 1))
	}
	if s.ProvisionedProductName != nil && len(*s.ProvisionedProductName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisionedProductName", 1))
	}

	if s.TerminateToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("TerminateToken"))
	}
	if s.TerminateToken != nil && len(*s.TerminateToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TerminateToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/TerminateProvisionedProductOutput
type TerminateProvisionedProductOutput struct {
	_ struct{} `type:"structure"`

	// Information about the result of this request.
	RecordDetail *RecordDetail `json:"servicecatalog:TerminateProvisionedProductOutput:RecordDetail" type:"structure"`
}

// String returns the string representation
func (s TerminateProvisionedProductOutput) String() string {
	return awsutil.Prettify(s)
}

const opTerminateProvisionedProduct = "TerminateProvisionedProduct"

// TerminateProvisionedProductRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Terminates the specified provisioned product.
//
// This operation does not delete any records associated with the provisioned
// product.
//
// You can check the status of this request using DescribeRecord.
//
//    // Example sending a request using TerminateProvisionedProductRequest.
//    req := client.TerminateProvisionedProductRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/TerminateProvisionedProduct
func (c *Client) TerminateProvisionedProductRequest(input *TerminateProvisionedProductInput) TerminateProvisionedProductRequest {
	op := &aws.Operation{
		Name:       opTerminateProvisionedProduct,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TerminateProvisionedProductInput{}
	}

	req := c.newRequest(op, input, &TerminateProvisionedProductOutput{})
	return TerminateProvisionedProductRequest{Request: req, Input: input, Copy: c.TerminateProvisionedProductRequest}
}

// TerminateProvisionedProductRequest is the request type for the
// TerminateProvisionedProduct API operation.
type TerminateProvisionedProductRequest struct {
	*aws.Request
	Input *TerminateProvisionedProductInput
	Copy  func(*TerminateProvisionedProductInput) TerminateProvisionedProductRequest
}

// Send marshals and sends the TerminateProvisionedProduct API request.
func (r TerminateProvisionedProductRequest) Send(ctx context.Context) (*TerminateProvisionedProductResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TerminateProvisionedProductResponse{
		TerminateProvisionedProductOutput: r.Request.Data.(*TerminateProvisionedProductOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TerminateProvisionedProductResponse is the response type for the
// TerminateProvisionedProduct API operation.
type TerminateProvisionedProductResponse struct {
	*TerminateProvisionedProductOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TerminateProvisionedProduct request.
func (r *TerminateProvisionedProductResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
