// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The UpdateDomainContactPrivacy request includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/UpdateDomainContactPrivacyRequest
type UpdateDomainContactPrivacyInput struct {
	_ struct{} `type:"structure"`

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the admin contact.
	AdminPrivacy *bool `type:"boolean"`

	// The name of the domain that you want to update the privacy setting for.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the registrant contact (domain
	// owner).
	RegistrantPrivacy *bool `type:"boolean"`

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the technical contact.
	TechPrivacy *bool `type:"boolean"`
}

// String returns the string representation
func (s UpdateDomainContactPrivacyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDomainContactPrivacyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDomainContactPrivacyInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The UpdateDomainContactPrivacy response includes the following element.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/UpdateDomainContactPrivacyResponse
type UpdateDomainContactPrivacyOutput struct {
	_ struct{} `type:"structure"`

	// Identifier for tracking the progress of the request. To use this ID to query
	// the operation status, use GetOperationDetail.
	//
	// OperationId is a required field
	OperationId *string `json:"route53domains:UpdateDomainContactPrivacyOutput:OperationId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateDomainContactPrivacyOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateDomainContactPrivacy = "UpdateDomainContactPrivacy"

// UpdateDomainContactPrivacyRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation updates the specified domain contact's privacy setting. When
// privacy protection is enabled, contact information such as email address
// is replaced either with contact information for Amazon Registrar (for .com,
// .net, and .org domains) or with contact information for our registrar associate,
// Gandi.
//
// This operation affects only the contact information for the specified contact
// type (registrant, administrator, or tech). If the request succeeds, Amazon
// Route 53 returns an operation ID that you can use with GetOperationDetail
// to track the progress and completion of the action. If the request doesn't
// complete successfully, the domain registrant will be notified by email.
//
//    // Example sending a request using UpdateDomainContactPrivacyRequest.
//    req := client.UpdateDomainContactPrivacyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/UpdateDomainContactPrivacy
func (c *Client) UpdateDomainContactPrivacyRequest(input *UpdateDomainContactPrivacyInput) UpdateDomainContactPrivacyRequest {
	op := &aws.Operation{
		Name:       opUpdateDomainContactPrivacy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateDomainContactPrivacyInput{}
	}

	req := c.newRequest(op, input, &UpdateDomainContactPrivacyOutput{})
	return UpdateDomainContactPrivacyRequest{Request: req, Input: input, Copy: c.UpdateDomainContactPrivacyRequest}
}

// UpdateDomainContactPrivacyRequest is the request type for the
// UpdateDomainContactPrivacy API operation.
type UpdateDomainContactPrivacyRequest struct {
	*aws.Request
	Input *UpdateDomainContactPrivacyInput
	Copy  func(*UpdateDomainContactPrivacyInput) UpdateDomainContactPrivacyRequest
}

// Send marshals and sends the UpdateDomainContactPrivacy API request.
func (r UpdateDomainContactPrivacyRequest) Send(ctx context.Context) (*UpdateDomainContactPrivacyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDomainContactPrivacyResponse{
		UpdateDomainContactPrivacyOutput: r.Request.Data.(*UpdateDomainContactPrivacyOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDomainContactPrivacyResponse is the response type for the
// UpdateDomainContactPrivacy API operation.
type UpdateDomainContactPrivacyResponse struct {
	*UpdateDomainContactPrivacyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDomainContactPrivacy request.
func (r *UpdateDomainContactPrivacyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
