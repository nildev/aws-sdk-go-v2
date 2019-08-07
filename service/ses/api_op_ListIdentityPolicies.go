// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to return a list of sending authorization policies that
// are attached to an identity. Sending authorization is an Amazon SES feature
// that enables you to authorize other senders to use your identities. For information,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListIdentityPoliciesRequest
type ListIdentityPoliciesInput struct {
	_ struct{} `type:"structure"`

	// The identity that is associated with the policy for which the policies will
	// be listed. You can specify an identity by using its name or by using its
	// Amazon Resource Name (ARN). Examples: user@example.com, example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	//
	// To successfully call this API, you must own the identity.
	//
	// Identity is a required field
	Identity *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListIdentityPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListIdentityPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListIdentityPoliciesInput"}

	if s.Identity == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identity"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A list of names of sending authorization policies that apply to an identity.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListIdentityPoliciesResponse
type ListIdentityPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// A list of names of policies that apply to the specified identity.
	//
	// PolicyNames is a required field
	PolicyNames []string `json:"email:ListIdentityPoliciesOutput:PolicyNames" type:"list" required:"true"`
}

// String returns the string representation
func (s ListIdentityPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListIdentityPolicies = "ListIdentityPolicies"

// ListIdentityPoliciesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns a list of sending authorization policies that are attached to the
// given identity (an email address or a domain). This API returns only a list.
// If you want the actual policy content, you can use GetIdentityPolicies.
//
// This API is for the identity owner only. If you have not verified the identity,
// this API will return an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using ListIdentityPoliciesRequest.
//    req := client.ListIdentityPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListIdentityPolicies
func (c *Client) ListIdentityPoliciesRequest(input *ListIdentityPoliciesInput) ListIdentityPoliciesRequest {
	op := &aws.Operation{
		Name:       opListIdentityPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListIdentityPoliciesInput{}
	}

	req := c.newRequest(op, input, &ListIdentityPoliciesOutput{})
	return ListIdentityPoliciesRequest{Request: req, Input: input, Copy: c.ListIdentityPoliciesRequest}
}

// ListIdentityPoliciesRequest is the request type for the
// ListIdentityPolicies API operation.
type ListIdentityPoliciesRequest struct {
	*aws.Request
	Input *ListIdentityPoliciesInput
	Copy  func(*ListIdentityPoliciesInput) ListIdentityPoliciesRequest
}

// Send marshals and sends the ListIdentityPolicies API request.
func (r ListIdentityPoliciesRequest) Send(ctx context.Context) (*ListIdentityPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIdentityPoliciesResponse{
		ListIdentityPoliciesOutput: r.Request.Data.(*ListIdentityPoliciesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListIdentityPoliciesResponse is the response type for the
// ListIdentityPolicies API operation.
type ListIdentityPoliciesResponse struct {
	*ListIdentityPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIdentityPolicies request.
func (r *ListIdentityPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
