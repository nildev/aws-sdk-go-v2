// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to create a new domain name.
type CreateDomainNameInput struct {
	_ struct{} `type:"structure"`

	// The reference to an AWS-managed certificate that will be used by edge-optimized
	// endpoint for this domain name. AWS Certificate Manager is the only supported
	// source.
	CertificateArn *string `locationName:"certificateArn" type:"string"`

	// [Deprecated] The body of the server certificate that will be used by edge-optimized
	// endpoint for this domain name provided by your certificate authority.
	CertificateBody *string `locationName:"certificateBody" type:"string"`

	// [Deprecated] The intermediate certificates and optionally the root certificate,
	// one after the other without any blank lines, used by an edge-optimized endpoint
	// for this domain name. If you include the root certificate, your certificate
	// chain must start with intermediate certificates and end with the root certificate.
	// Use the intermediate certificates that were provided by your certificate
	// authority. Do not include any intermediaries that are not in the chain of
	// trust path.
	CertificateChain *string `locationName:"certificateChain" type:"string"`

	// The user-friendly name of the certificate that will be used by edge-optimized
	// endpoint for this domain name.
	CertificateName *string `locationName:"certificateName" type:"string"`

	// [Deprecated] Your edge-optimized endpoint's domain name certificate's private
	// key.
	CertificatePrivateKey *string `locationName:"certificatePrivateKey" type:"string"`

	// [Required] The name of the DomainName resource.
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`

	// The endpoint configuration of this DomainName showing the endpoint types
	// of the domain name.
	EndpointConfiguration *EndpointConfiguration `locationName:"endpointConfiguration" type:"structure"`

	// The reference to an AWS-managed certificate that will be used by regional
	// endpoint for this domain name. AWS Certificate Manager is the only supported
	// source.
	RegionalCertificateArn *string `locationName:"regionalCertificateArn" type:"string"`

	// The user-friendly name of the certificate that will be used by regional endpoint
	// for this domain name.
	RegionalCertificateName *string `locationName:"regionalCertificateName" type:"string"`

	// The Transport Layer Security (TLS) version + cipher suite for this DomainName.
	// The valid values are TLS_1_0 and TLS_1_2.
	SecurityPolicy SecurityPolicy `locationName:"securityPolicy" type:"string" enum:"true"`

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/].
	// The tag key can be up to 128 characters and must not start with aws:. The
	// tag value can be up to 256 characters.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateDomainNameInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDomainNameInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDomainNameInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDomainNameInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.CertificateArn != nil {
		v := *s.CertificateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateBody != nil {
		v := *s.CertificateBody

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateBody", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateChain != nil {
		v := *s.CertificateChain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateChain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateName != nil {
		v := *s.CertificateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificatePrivateKey != nil {
		v := *s.CertificatePrivateKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificatePrivateKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndpointConfiguration != nil {
		v := s.EndpointConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "endpointConfiguration", v, metadata)
	}
	if s.RegionalCertificateArn != nil {
		v := *s.RegionalCertificateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "regionalCertificateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RegionalCertificateName != nil {
		v := *s.RegionalCertificateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "regionalCertificateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.SecurityPolicy) > 0 {
		v := s.SecurityPolicy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityPolicy", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Represents a custom domain name as a user-friendly host name of an API (RestApi).
//
// When you deploy an API, API Gateway creates a default host name for the API.
// This default API host name is of the {restapi-id}.execute-api.{region}.amazonaws.com
// format. With the default host name, you can access the API's root resource
// with the URL of https://{restapi-id}.execute-api.{region}.amazonaws.com/{stage}/.
// When you set up a custom domain name of apis.example.com for this API, you
// can then access the same resource using the URL of the https://apis.examples.com/myApi,
// where myApi is the base path mapping (BasePathMapping) of your API under
// the custom domain name.
//
// Set a Custom Host Name for an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type CreateDomainNameOutput struct {
	_ struct{} `type:"structure"`

	// The reference to an AWS-managed certificate that will be used by edge-optimized
	// endpoint for this domain name. AWS Certificate Manager is the only supported
	// source.
	CertificateArn *string `json:"apigateway:CreateDomainNameOutput:CertificateArn" locationName:"certificateArn" type:"string"`

	// The name of the certificate that will be used by edge-optimized endpoint
	// for this domain name.
	CertificateName *string `json:"apigateway:CreateDomainNameOutput:CertificateName" locationName:"certificateName" type:"string"`

	// The timestamp when the certificate that was used by edge-optimized endpoint
	// for this domain name was uploaded.
	CertificateUploadDate *time.Time `json:"apigateway:CreateDomainNameOutput:CertificateUploadDate" locationName:"certificateUploadDate" type:"timestamp" timestampFormat:"unix"`

	// The domain name of the Amazon CloudFront distribution associated with this
	// custom domain name for an edge-optimized endpoint. You set up this association
	// when adding a DNS record pointing the custom domain name to this distribution
	// name. For more information about CloudFront distributions, see the Amazon
	// CloudFront documentation (https://aws.amazon.com/documentation/cloudfront/).
	DistributionDomainName *string `json:"apigateway:CreateDomainNameOutput:DistributionDomainName" locationName:"distributionDomainName" type:"string"`

	// The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized
	// endpoint. The valid value is Z2FDTNDATAQYW2 for all the regions. For more
	// information, see Set up a Regional Custom Domain Name (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-create.html)
	// and AWS Regions and Endpoints for API Gateway (https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region).
	DistributionHostedZoneId *string `json:"apigateway:CreateDomainNameOutput:DistributionHostedZoneId" locationName:"distributionHostedZoneId" type:"string"`

	// The custom domain name as an API host name, for example, my-api.example.com.
	DomainName *string `json:"apigateway:CreateDomainNameOutput:DomainName" locationName:"domainName" type:"string"`

	// The status of the DomainName migration. The valid values are AVAILABLE and
	// UPDATING. If the status is UPDATING, the domain cannot be modified further
	// until the existing operation is complete. If it is AVAILABLE, the domain
	// can be updated.
	DomainNameStatus DomainNameStatus `json:"apigateway:CreateDomainNameOutput:DomainNameStatus" locationName:"domainNameStatus" type:"string" enum:"true"`

	// An optional text message containing detailed information about status of
	// the DomainName migration.
	DomainNameStatusMessage *string `json:"apigateway:CreateDomainNameOutput:DomainNameStatusMessage" locationName:"domainNameStatusMessage" type:"string"`

	// The endpoint configuration of this DomainName showing the endpoint types
	// of the domain name.
	EndpointConfiguration *EndpointConfiguration `json:"apigateway:CreateDomainNameOutput:EndpointConfiguration" locationName:"endpointConfiguration" type:"structure"`

	// The reference to an AWS-managed certificate that will be used for validating
	// the regional domain name. AWS Certificate Manager is the only supported source.
	RegionalCertificateArn *string `json:"apigateway:CreateDomainNameOutput:RegionalCertificateArn" locationName:"regionalCertificateArn" type:"string"`

	// The name of the certificate that will be used for validating the regional
	// domain name.
	RegionalCertificateName *string `json:"apigateway:CreateDomainNameOutput:RegionalCertificateName" locationName:"regionalCertificateName" type:"string"`

	// The domain name associated with the regional endpoint for this custom domain
	// name. You set up this association by adding a DNS record that points the
	// custom domain name to this regional domain name. The regional domain name
	// is returned by API Gateway when you create a regional endpoint.
	RegionalDomainName *string `json:"apigateway:CreateDomainNameOutput:RegionalDomainName" locationName:"regionalDomainName" type:"string"`

	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	// For more information, see Set up a Regional Custom Domain Name (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-create.html)
	// and AWS Regions and Endpoints for API Gateway (https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region).
	RegionalHostedZoneId *string `json:"apigateway:CreateDomainNameOutput:RegionalHostedZoneId" locationName:"regionalHostedZoneId" type:"string"`

	// The Transport Layer Security (TLS) version + cipher suite for this DomainName.
	// The valid values are TLS_1_0 and TLS_1_2.
	SecurityPolicy SecurityPolicy `json:"apigateway:CreateDomainNameOutput:SecurityPolicy" locationName:"securityPolicy" type:"string" enum:"true"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `json:"apigateway:CreateDomainNameOutput:Tags" locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateDomainNameOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDomainNameOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificateArn != nil {
		v := *s.CertificateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateName != nil {
		v := *s.CertificateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateUploadDate != nil {
		v := *s.CertificateUploadDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateUploadDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.DistributionDomainName != nil {
		v := *s.DistributionDomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "distributionDomainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DistributionHostedZoneId != nil {
		v := *s.DistributionHostedZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "distributionHostedZoneId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DomainNameStatus) > 0 {
		v := s.DomainNameStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainNameStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DomainNameStatusMessage != nil {
		v := *s.DomainNameStatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainNameStatusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndpointConfiguration != nil {
		v := s.EndpointConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "endpointConfiguration", v, metadata)
	}
	if s.RegionalCertificateArn != nil {
		v := *s.RegionalCertificateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "regionalCertificateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RegionalCertificateName != nil {
		v := *s.RegionalCertificateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "regionalCertificateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RegionalDomainName != nil {
		v := *s.RegionalDomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "regionalDomainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RegionalHostedZoneId != nil {
		v := *s.RegionalHostedZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "regionalHostedZoneId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.SecurityPolicy) > 0 {
		v := s.SecurityPolicy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityPolicy", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opCreateDomainName = "CreateDomainName"

// CreateDomainNameRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Creates a new domain name.
//
//    // Example sending a request using CreateDomainNameRequest.
//    req := client.CreateDomainNameRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateDomainNameRequest(input *CreateDomainNameInput) CreateDomainNameRequest {
	op := &aws.Operation{
		Name:       opCreateDomainName,
		HTTPMethod: "POST",
		HTTPPath:   "/domainnames",
	}

	if input == nil {
		input = &CreateDomainNameInput{}
	}

	req := c.newRequest(op, input, &CreateDomainNameOutput{})
	return CreateDomainNameRequest{Request: req, Input: input, Copy: c.CreateDomainNameRequest}
}

// CreateDomainNameRequest is the request type for the
// CreateDomainName API operation.
type CreateDomainNameRequest struct {
	*aws.Request
	Input *CreateDomainNameInput
	Copy  func(*CreateDomainNameInput) CreateDomainNameRequest
}

// Send marshals and sends the CreateDomainName API request.
func (r CreateDomainNameRequest) Send(ctx context.Context) (*CreateDomainNameResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDomainNameResponse{
		CreateDomainNameOutput: r.Request.Data.(*CreateDomainNameOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDomainNameResponse is the response type for the
// CreateDomainName API operation.
type CreateDomainNameResponse struct {
	*CreateDomainNameOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDomainName request.
func (r *CreateDomainNameResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
