// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package simpledb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DomainMetadataInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain for which to display the metadata of.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DomainMetadataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DomainMetadataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DomainMetadataInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DomainMetadataOutput struct {
	_ struct{} `type:"structure"`

	// The number of unique attribute names in the domain.
	AttributeNameCount *int64 `json:"sdb:DomainMetadataOutput:AttributeNameCount" type:"integer"`

	// The total size of all unique attribute names in the domain, in bytes.
	AttributeNamesSizeBytes *int64 `json:"sdb:DomainMetadataOutput:AttributeNamesSizeBytes" type:"long"`

	// The number of all attribute name/value pairs in the domain.
	AttributeValueCount *int64 `json:"sdb:DomainMetadataOutput:AttributeValueCount" type:"integer"`

	// The total size of all attribute values in the domain, in bytes.
	AttributeValuesSizeBytes *int64 `json:"sdb:DomainMetadataOutput:AttributeValuesSizeBytes" type:"long"`

	// The number of all items in the domain.
	ItemCount *int64 `json:"sdb:DomainMetadataOutput:ItemCount" type:"integer"`

	// The total size of all item names in the domain, in bytes.
	ItemNamesSizeBytes *int64 `json:"sdb:DomainMetadataOutput:ItemNamesSizeBytes" type:"long"`

	// The data and time when metadata was calculated, in Epoch (UNIX) seconds.
	Timestamp *int64 `json:"sdb:DomainMetadataOutput:Timestamp" type:"integer"`
}

// String returns the string representation
func (s DomainMetadataOutput) String() string {
	return awsutil.Prettify(s)
}

const opDomainMetadata = "DomainMetadata"

// DomainMetadataRequest returns a request value for making API operation for
// Amazon SimpleDB.
//
// Returns information about the domain, including when the domain was created,
// the number of items and attributes in the domain, and the size of the attribute
// names and values.
//
//    // Example sending a request using DomainMetadataRequest.
//    req := client.DomainMetadataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DomainMetadataRequest(input *DomainMetadataInput) DomainMetadataRequest {
	op := &aws.Operation{
		Name:       opDomainMetadata,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DomainMetadataInput{}
	}

	req := c.newRequest(op, input, &DomainMetadataOutput{})
	return DomainMetadataRequest{Request: req, Input: input, Copy: c.DomainMetadataRequest}
}

// DomainMetadataRequest is the request type for the
// DomainMetadata API operation.
type DomainMetadataRequest struct {
	*aws.Request
	Input *DomainMetadataInput
	Copy  func(*DomainMetadataInput) DomainMetadataRequest
}

// Send marshals and sends the DomainMetadata API request.
func (r DomainMetadataRequest) Send(ctx context.Context) (*DomainMetadataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DomainMetadataResponse{
		DomainMetadataOutput: r.Request.Data.(*DomainMetadataOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DomainMetadataResponse is the response type for the
// DomainMetadata API operation.
type DomainMetadataResponse struct {
	*DomainMetadataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DomainMetadata request.
func (r *DomainMetadataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
