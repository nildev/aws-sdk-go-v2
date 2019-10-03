// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request to get the name of a DomainName resource.
type GetDomainNameInput struct {
	_ struct{} `type:"structure"`

	// [Required] The name of the DomainName resource.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domain_name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDomainNameInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDomainNameInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDomainNameInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainNameInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domain_name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
type GetDomainNameOutput struct {
	_ struct{} `type:"structure"`

	// The reference to an AWS-managed certificate that will be used by edge-optimized
	// endpoint for this domain name. AWS Certificate Manager is the only supported
	// source.
	CertificateArn *string `json:"apigateway:GetDomainNameOutput:CertificateArn" locationName:"certificateArn" type:"string"`

	// The name of the certificate that will be used by edge-optimized endpoint
	// for this domain name.
	CertificateName *string `json:"apigateway:GetDomainNameOutput:CertificateName" locationName:"certificateName" type:"string"`

	// The timestamp when the certificate that was used by edge-optimized endpoint
	// for this domain name was uploaded.
	CertificateUploadDate *time.Time `json:"apigateway:GetDomainNameOutput:CertificateUploadDate" locationName:"certificateUploadDate" type:"timestamp" timestampFormat:"unix"`

	// The domain name of the Amazon CloudFront distribution associated with this
	// custom domain name for an edge-optimized endpoint. You set up this association
	// when adding a DNS record pointing the custom domain name to this distribution
	// name. For more information about CloudFront distributions, see the Amazon
	// CloudFront documentation (https://aws.amazon.com/documentation/cloudfront/).
	DistributionDomainName *string `json:"apigateway:GetDomainNameOutput:DistributionDomainName" locationName:"distributionDomainName" type:"string"`

	// The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized
	// endpoint. The valid value is Z2FDTNDATAQYW2 for all the regions. For more
	// information, see Set up a Regional Custom Domain Name (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-create.html)
	// and AWS Regions and Endpoints for API Gateway (https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region).
	DistributionHostedZoneId *string `json:"apigateway:GetDomainNameOutput:DistributionHostedZoneId" locationName:"distributionHostedZoneId" type:"string"`

	// The custom domain name as an API host name, for example, my-api.example.com.
	DomainName *string `json:"apigateway:GetDomainNameOutput:DomainName" locationName:"domainName" type:"string"`

	// The status of the DomainName migration. The valid values are AVAILABLE and
	// UPDATING. If the status is UPDATING, the domain cannot be modified further
	// until the existing operation is complete. If it is AVAILABLE, the domain
	// can be updated.
	DomainNameStatus DomainNameStatus `json:"apigateway:GetDomainNameOutput:DomainNameStatus" locationName:"domainNameStatus" type:"string" enum:"true"`

	// An optional text message containing detailed information about status of
	// the DomainName migration.
	DomainNameStatusMessage *string `json:"apigateway:GetDomainNameOutput:DomainNameStatusMessage" locationName:"domainNameStatusMessage" type:"string"`

	// The endpoint configuration of this DomainName showing the endpoint types
	// of the domain name.
	EndpointConfiguration *EndpointConfiguration `json:"apigateway:GetDomainNameOutput:EndpointConfiguration" locationName:"endpointConfiguration" type:"structure"`

	// The reference to an AWS-managed certificate that will be used for validating
	// the regional domain name. AWS Certificate Manager is the only supported source.
	RegionalCertificateArn *string `json:"apigateway:GetDomainNameOutput:RegionalCertificateArn" locationName:"regionalCertificateArn" type:"string"`

	// The name of the certificate that will be used for validating the regional
	// domain name.
	RegionalCertificateName *string `json:"apigateway:GetDomainNameOutput:RegionalCertificateName" locationName:"regionalCertificateName" type:"string"`

	// The domain name associated with the regional endpoint for this custom domain
	// name. You set up this association by adding a DNS record that points the
	// custom domain name to this regional domain name. The regional domain name
	// is returned by API Gateway when you create a regional endpoint.
	RegionalDomainName *string `json:"apigateway:GetDomainNameOutput:RegionalDomainName" locationName:"regionalDomainName" type:"string"`

	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	// For more information, see Set up a Regional Custom Domain Name (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-create.html)
	// and AWS Regions and Endpoints for API Gateway (https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region).
	RegionalHostedZoneId *string `json:"apigateway:GetDomainNameOutput:RegionalHostedZoneId" locationName:"regionalHostedZoneId" type:"string"`

	// The Transport Layer Security (TLS) version + cipher suite for this DomainName.
	// The valid values are TLS_1_0 and TLS_1_2.
	SecurityPolicy SecurityPolicy `json:"apigateway:GetDomainNameOutput:SecurityPolicy" locationName:"securityPolicy" type:"string" enum:"true"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `json:"apigateway:GetDomainNameOutput:Tags" locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetDomainNameOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainNameOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opGetDomainName = "GetDomainName"

// GetDomainNameRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Represents a domain name that is contained in a simpler, more intuitive URL
// that can be called.
//
//    // Example sending a request using GetDomainNameRequest.
//    req := client.GetDomainNameRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetDomainNameRequest(input *GetDomainNameInput) GetDomainNameRequest {
	op := &aws.Operation{
		Name:       opGetDomainName,
		HTTPMethod: "GET",
		HTTPPath:   "/domainnames/{domain_name}",
	}

	if input == nil {
		input = &GetDomainNameInput{}
	}

	req := c.newRequest(op, input, &GetDomainNameOutput{})
	return GetDomainNameRequest{Request: req, Input: input, Copy: c.GetDomainNameRequest}
}

// GetDomainNameRequest is the request type for the
// GetDomainName API operation.
type GetDomainNameRequest struct {
	*aws.Request
	Input *GetDomainNameInput
	Copy  func(*GetDomainNameInput) GetDomainNameRequest
}

// Send marshals and sends the GetDomainName API request.
func (r GetDomainNameRequest) Send(ctx context.Context) (*GetDomainNameResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainNameResponse{
		GetDomainNameOutput: r.Request.Data.(*GetDomainNameOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDomainNameResponse is the response type for the
// GetDomainName API operation.
type GetDomainNameResponse struct {
	*GetDomainNameOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainName request.
func (r *GetDomainNameResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
