// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreateProductInput
type CreateProductInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The description of the product.
	Description *string `type:"string"`

	// The distributor of the product.
	Distributor *string `type:"string"`

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The name of the product.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The owner of the product.
	//
	// Owner is a required field
	Owner *string `type:"string" required:"true"`

	// The type of product.
	//
	// ProductType is a required field
	ProductType ProductType `type:"string" required:"true" enum:"true"`

	// The configuration of the provisioning artifact.
	//
	// ProvisioningArtifactParameters is a required field
	ProvisioningArtifactParameters *ProvisioningArtifactProperties `type:"structure" required:"true"`

	// The support information about the product.
	SupportDescription *string `type:"string"`

	// The contact email for product support.
	SupportEmail *string `type:"string"`

	// The contact URL for product support.
	SupportUrl *string `type:"string"`

	// One or more tags.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateProductInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProductInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProductInput"}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Owner == nil {
		invalidParams.Add(aws.NewErrParamRequired("Owner"))
	}
	if len(s.ProductType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ProductType"))
	}

	if s.ProvisioningArtifactParameters == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisioningArtifactParameters"))
	}
	if s.ProvisioningArtifactParameters != nil {
		if err := s.ProvisioningArtifactParameters.Validate(); err != nil {
			invalidParams.AddNested("ProvisioningArtifactParameters", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreateProductOutput
type CreateProductOutput struct {
	_ struct{} `type:"structure"`

	// Information about the product view.
	ProductViewDetail *ProductViewDetail `json:"servicecatalog:CreateProductOutput:ProductViewDetail" type:"structure"`

	// Information about the provisioning artifact.
	ProvisioningArtifactDetail *ProvisioningArtifactDetail `json:"servicecatalog:CreateProductOutput:ProvisioningArtifactDetail" type:"structure"`

	// Information about the tags associated with the product.
	Tags []Tag `json:"servicecatalog:CreateProductOutput:Tags" type:"list"`
}

// String returns the string representation
func (s CreateProductOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateProduct = "CreateProduct"

// CreateProductRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Creates a product.
//
//    // Example sending a request using CreateProductRequest.
//    req := client.CreateProductRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreateProduct
func (c *Client) CreateProductRequest(input *CreateProductInput) CreateProductRequest {
	op := &aws.Operation{
		Name:       opCreateProduct,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateProductInput{}
	}

	req := c.newRequest(op, input, &CreateProductOutput{})
	return CreateProductRequest{Request: req, Input: input, Copy: c.CreateProductRequest}
}

// CreateProductRequest is the request type for the
// CreateProduct API operation.
type CreateProductRequest struct {
	*aws.Request
	Input *CreateProductInput
	Copy  func(*CreateProductInput) CreateProductRequest
}

// Send marshals and sends the CreateProduct API request.
func (r CreateProductRequest) Send(ctx context.Context) (*CreateProductResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProductResponse{
		CreateProductOutput: r.Request.Data.(*CreateProductOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProductResponse is the response type for the
// CreateProduct API operation.
type CreateProductResponse struct {
	*CreateProductOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProduct request.
func (r *CreateProductResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
