// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Container for the parameters to the DescribeElasticsearchDomainConfig operation.
// Specifies the domain name for which you want configuration information.
type DescribeElasticsearchDomainConfigInput struct {
	_ struct{} `type:"structure"`

	// The Elasticsearch domain that you want to get information about.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"DomainName" min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeElasticsearchDomainConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeElasticsearchDomainConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeElasticsearchDomainConfigInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeElasticsearchDomainConfigInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DomainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a DescribeElasticsearchDomainConfig request. Contains the configuration
// information of the requested domain.
type DescribeElasticsearchDomainConfigOutput struct {
	_ struct{} `type:"structure"`

	// The configuration information of the domain requested in the DescribeElasticsearchDomainConfig
	// request.
	//
	// DomainConfig is a required field
	DomainConfig *ElasticsearchDomainConfig `json:"es:DescribeElasticsearchDomainConfigOutput:DomainConfig" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeElasticsearchDomainConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeElasticsearchDomainConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainConfig != nil {
		v := s.DomainConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DomainConfig", v, metadata)
	}
	return nil
}

const opDescribeElasticsearchDomainConfig = "DescribeElasticsearchDomainConfig"

// DescribeElasticsearchDomainConfigRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Provides cluster configuration information about the specified Elasticsearch
// domain, such as the state, creation date, update version, and update date
// for cluster options.
//
//    // Example sending a request using DescribeElasticsearchDomainConfigRequest.
//    req := client.DescribeElasticsearchDomainConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeElasticsearchDomainConfigRequest(input *DescribeElasticsearchDomainConfigInput) DescribeElasticsearchDomainConfigRequest {
	op := &aws.Operation{
		Name:       opDescribeElasticsearchDomainConfig,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/es/domain/{DomainName}/config",
	}

	if input == nil {
		input = &DescribeElasticsearchDomainConfigInput{}
	}

	req := c.newRequest(op, input, &DescribeElasticsearchDomainConfigOutput{})
	return DescribeElasticsearchDomainConfigRequest{Request: req, Input: input, Copy: c.DescribeElasticsearchDomainConfigRequest}
}

// DescribeElasticsearchDomainConfigRequest is the request type for the
// DescribeElasticsearchDomainConfig API operation.
type DescribeElasticsearchDomainConfigRequest struct {
	*aws.Request
	Input *DescribeElasticsearchDomainConfigInput
	Copy  func(*DescribeElasticsearchDomainConfigInput) DescribeElasticsearchDomainConfigRequest
}

// Send marshals and sends the DescribeElasticsearchDomainConfig API request.
func (r DescribeElasticsearchDomainConfigRequest) Send(ctx context.Context) (*DescribeElasticsearchDomainConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeElasticsearchDomainConfigResponse{
		DescribeElasticsearchDomainConfigOutput: r.Request.Data.(*DescribeElasticsearchDomainConfigOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeElasticsearchDomainConfigResponse is the response type for the
// DescribeElasticsearchDomainConfig API operation.
type DescribeElasticsearchDomainConfigResponse struct {
	*DescribeElasticsearchDomainConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeElasticsearchDomainConfig request.
func (r *DescribeElasticsearchDomainConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
