// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ImportCertificateMessage
type ImportCertificateInput struct {
	_ struct{} `type:"structure"`

	// A customer-assigned name for the certificate. Identifiers must begin with
	// a letter; must contain only ASCII letters, digits, and hyphens; and must
	// not end with a hyphen or contain two consecutive hyphens.
	//
	// CertificateIdentifier is a required field
	CertificateIdentifier *string `type:"string" required:"true"`

	// The contents of a .pem file, which contains an X.509 certificate.
	CertificatePem *string `type:"string"`

	// The location of an imported Oracle Wallet certificate for use with SSL.
	//
	// CertificateWallet is automatically base64 encoded/decoded by the SDK.
	CertificateWallet []byte `type:"blob"`

	// The tags associated with the certificate.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s ImportCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportCertificateInput"}

	if s.CertificateIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ImportCertificateResponse
type ImportCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The certificate to be uploaded.
	Certificate *Certificate `json:"dms:ImportCertificateOutput:Certificate" type:"structure"`
}

// String returns the string representation
func (s ImportCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportCertificate = "ImportCertificate"

// ImportCertificateRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Uploads the specified certificate.
//
//    // Example sending a request using ImportCertificateRequest.
//    req := client.ImportCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/ImportCertificate
func (c *Client) ImportCertificateRequest(input *ImportCertificateInput) ImportCertificateRequest {
	op := &aws.Operation{
		Name:       opImportCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportCertificateInput{}
	}

	req := c.newRequest(op, input, &ImportCertificateOutput{})
	return ImportCertificateRequest{Request: req, Input: input, Copy: c.ImportCertificateRequest}
}

// ImportCertificateRequest is the request type for the
// ImportCertificate API operation.
type ImportCertificateRequest struct {
	*aws.Request
	Input *ImportCertificateInput
	Copy  func(*ImportCertificateInput) ImportCertificateRequest
}

// Send marshals and sends the ImportCertificate API request.
func (r ImportCertificateRequest) Send(ctx context.Context) (*ImportCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportCertificateResponse{
		ImportCertificateOutput: r.Request.Data.(*ImportCertificateOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportCertificateResponse is the response type for the
// ImportCertificate API operation.
type ImportCertificateResponse struct {
	*ImportCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportCertificate request.
func (r *ImportCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
