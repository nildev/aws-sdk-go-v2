// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to return the Amazon SES custom MAIL FROM attributes
// for a list of identities. For information about using a custom MAIL FROM
// domain, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetIdentityMailFromDomainAttributesRequest
type GetIdentityMailFromDomainAttributesInput struct {
	_ struct{} `type:"structure"`

	// A list of one or more identities.
	//
	// Identities is a required field
	Identities []string `type:"list" required:"true"`
}

// String returns the string representation
func (s GetIdentityMailFromDomainAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIdentityMailFromDomainAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIdentityMailFromDomainAttributesInput"}

	if s.Identities == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identities"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the custom MAIL FROM attributes for a list of identities.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetIdentityMailFromDomainAttributesResponse
type GetIdentityMailFromDomainAttributesOutput struct {
	_ struct{} `type:"structure"`

	// A map of identities to custom MAIL FROM attributes.
	//
	// MailFromDomainAttributes is a required field
	MailFromDomainAttributes map[string]IdentityMailFromDomainAttributes `json:"email:GetIdentityMailFromDomainAttributesOutput:MailFromDomainAttributes" type:"map" required:"true"`
}

// String returns the string representation
func (s GetIdentityMailFromDomainAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetIdentityMailFromDomainAttributes = "GetIdentityMailFromDomainAttributes"

// GetIdentityMailFromDomainAttributesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns the custom MAIL FROM attributes for a list of identities (email addresses
// : domains).
//
// This operation is throttled at one request per second and can only get custom
// MAIL FROM attributes for up to 100 identities at a time.
//
//    // Example sending a request using GetIdentityMailFromDomainAttributesRequest.
//    req := client.GetIdentityMailFromDomainAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetIdentityMailFromDomainAttributes
func (c *Client) GetIdentityMailFromDomainAttributesRequest(input *GetIdentityMailFromDomainAttributesInput) GetIdentityMailFromDomainAttributesRequest {
	op := &aws.Operation{
		Name:       opGetIdentityMailFromDomainAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetIdentityMailFromDomainAttributesInput{}
	}

	req := c.newRequest(op, input, &GetIdentityMailFromDomainAttributesOutput{})
	return GetIdentityMailFromDomainAttributesRequest{Request: req, Input: input, Copy: c.GetIdentityMailFromDomainAttributesRequest}
}

// GetIdentityMailFromDomainAttributesRequest is the request type for the
// GetIdentityMailFromDomainAttributes API operation.
type GetIdentityMailFromDomainAttributesRequest struct {
	*aws.Request
	Input *GetIdentityMailFromDomainAttributesInput
	Copy  func(*GetIdentityMailFromDomainAttributesInput) GetIdentityMailFromDomainAttributesRequest
}

// Send marshals and sends the GetIdentityMailFromDomainAttributes API request.
func (r GetIdentityMailFromDomainAttributesRequest) Send(ctx context.Context) (*GetIdentityMailFromDomainAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIdentityMailFromDomainAttributesResponse{
		GetIdentityMailFromDomainAttributesOutput: r.Request.Data.(*GetIdentityMailFromDomainAttributesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIdentityMailFromDomainAttributesResponse is the response type for the
// GetIdentityMailFromDomainAttributes API operation.
type GetIdentityMailFromDomainAttributesResponse struct {
	*GetIdentityMailFromDomainAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIdentityMailFromDomainAttributes request.
func (r *GetIdentityMailFromDomainAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
