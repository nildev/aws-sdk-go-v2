// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ImportCertificateRequest
type ImportCertificateInput struct {
	_ struct{} `type:"structure"`

	// The certificate to import.
	//
	// Certificate is automatically base64 encoded/decoded by the SDK.
	//
	// Certificate is a required field
	Certificate []byte `min:"1" type:"blob" required:"true"`

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of an imported certificate to replace. To import a new certificate, omit
	// this field.
	CertificateArn *string `min:"20" type:"string"`

	// The PEM encoded certificate chain.
	//
	// CertificateChain is automatically base64 encoded/decoded by the SDK.
	CertificateChain []byte `min:"1" type:"blob"`

	// The private key that matches the public key in the certificate.
	//
	// PrivateKey is automatically base64 encoded/decoded by the SDK.
	//
	// PrivateKey is a required field
	PrivateKey []byte `min:"1" type:"blob" required:"true"`
}

// String returns the string representation
func (s ImportCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportCertificateInput"}

	if s.Certificate == nil {
		invalidParams.Add(aws.NewErrParamRequired("Certificate"))
	}
	if s.Certificate != nil && len(s.Certificate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Certificate", 1))
	}
	if s.CertificateArn != nil && len(*s.CertificateArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateArn", 20))
	}
	if s.CertificateChain != nil && len(s.CertificateChain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateChain", 1))
	}

	if s.PrivateKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("PrivateKey"))
	}
	if s.PrivateKey != nil && len(s.PrivateKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PrivateKey", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ImportCertificateResponse
type ImportCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the imported certificate.
	CertificateArn *string `json:"acm:ImportCertificateOutput:CertificateArn" min:"20" type:"string"`
}

// String returns the string representation
func (s ImportCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportCertificate = "ImportCertificate"

// ImportCertificateRequest returns a request value for making API operation for
// AWS Certificate Manager.
//
// Imports a certificate into AWS Certificate Manager (ACM) to use with services
// that are integrated with ACM. Note that integrated services (https://docs.aws.amazon.com/acm/latest/userguide/acm-services.html)
// allow only certificate types and keys they support to be associated with
// their resources. Further, their support differs depending on whether the
// certificate is imported into IAM or into ACM. For more information, see the
// documentation for each service. For more information about importing certificates
// into ACM, see Importing Certificates (https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html)
// in the AWS Certificate Manager User Guide.
//
// ACM does not provide managed renewal (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html)
// for certificates that you import.
//
// Note the following guidelines when importing third party certificates:
//
//    * You must enter the private key that matches the certificate you are
//    importing.
//
//    * The private key must be unencrypted. You cannot import a private key
//    that is protected by a password or a passphrase.
//
//    * If the certificate you are importing is not self-signed, you must enter
//    its certificate chain.
//
//    * If a certificate chain is included, the issuer must be the subject of
//    one of the certificates in the chain.
//
//    * The certificate, private key, and certificate chain must be PEM-encoded.
//
//    * The current time must be between the Not Before and Not After certificate
//    fields.
//
//    * The Issuer field must not be empty.
//
//    * The OCSP authority URL, if present, must not exceed 1000 characters.
//
//    * To import a new certificate, omit the CertificateArn argument. Include
//    this argument only when you want to replace a previously imported certificate.
//
//    * When you import a certificate by using the CLI, you must specify the
//    certificate, the certificate chain, and the private key by their file
//    names preceded by file://. For example, you can specify a certificate
//    saved in the C:\temp folder as file://C:\temp\certificate_to_import.pem.
//    If you are making an HTTP or HTTPS Query request, include these arguments
//    as BLOBs.
//
//    * When you import a certificate by using an SDK, you must specify the
//    certificate, the certificate chain, and the private key files in the manner
//    required by the programming language you're using.
//
// This operation returns the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
// of the imported certificate.
//
//    // Example sending a request using ImportCertificateRequest.
//    req := client.ImportCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/ImportCertificate
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
