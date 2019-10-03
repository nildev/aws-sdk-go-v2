// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for create Domain Association request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/CreateDomainAssociationRequest
type CreateDomainAssociationInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Domain name for the Domain Association.
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`

	// Enables automated creation of Subdomains for branches.
	EnableAutoSubDomain *bool `locationName:"enableAutoSubDomain" type:"boolean"`

	// Setting structure for the Subdomain.
	//
	// SubDomainSettings is a required field
	SubDomainSettings []SubDomainSetting `locationName:"subDomainSettings" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateDomainAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDomainAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDomainAssociationInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.SubDomainSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubDomainSettings"))
	}
	if s.SubDomainSettings != nil {
		for i, v := range s.SubDomainSettings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SubDomainSettings", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDomainAssociationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnableAutoSubDomain != nil {
		v := *s.EnableAutoSubDomain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enableAutoSubDomain", protocol.BoolValue(v), metadata)
	}
	if s.SubDomainSettings != nil {
		v := s.SubDomainSettings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "subDomainSettings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for the create Domain Association request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/CreateDomainAssociationResult
type CreateDomainAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Domain Association structure.
	//
	// DomainAssociation is a required field
	DomainAssociation *DomainAssociation `json:"amplify:CreateDomainAssociationOutput:DomainAssociation" locationName:"domainAssociation" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateDomainAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDomainAssociationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainAssociation != nil {
		v := s.DomainAssociation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "domainAssociation", v, metadata)
	}
	return nil
}

const opCreateDomainAssociation = "CreateDomainAssociation"

// CreateDomainAssociationRequest returns a request value for making API operation for
// AWS Amplify.
//
// Create a new DomainAssociation on an App
//
//    // Example sending a request using CreateDomainAssociationRequest.
//    req := client.CreateDomainAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/CreateDomainAssociation
func (c *Client) CreateDomainAssociationRequest(input *CreateDomainAssociationInput) CreateDomainAssociationRequest {
	op := &aws.Operation{
		Name:       opCreateDomainAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/apps/{appId}/domains",
	}

	if input == nil {
		input = &CreateDomainAssociationInput{}
	}

	req := c.newRequest(op, input, &CreateDomainAssociationOutput{})
	return CreateDomainAssociationRequest{Request: req, Input: input, Copy: c.CreateDomainAssociationRequest}
}

// CreateDomainAssociationRequest is the request type for the
// CreateDomainAssociation API operation.
type CreateDomainAssociationRequest struct {
	*aws.Request
	Input *CreateDomainAssociationInput
	Copy  func(*CreateDomainAssociationInput) CreateDomainAssociationRequest
}

// Send marshals and sends the CreateDomainAssociation API request.
func (r CreateDomainAssociationRequest) Send(ctx context.Context) (*CreateDomainAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDomainAssociationResponse{
		CreateDomainAssociationOutput: r.Request.Data.(*CreateDomainAssociationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDomainAssociationResponse is the response type for the
// CreateDomainAssociation API operation.
type CreateDomainAssociationResponse struct {
	*CreateDomainAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDomainAssociation request.
func (r *CreateDomainAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
