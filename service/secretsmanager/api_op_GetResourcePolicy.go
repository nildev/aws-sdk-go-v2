// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/GetResourcePolicyRequest
type GetResourcePolicyInput struct {
	_ struct{} `type:"structure"`

	// Specifies the secret that you want to retrieve the attached resource-based
	// policy for. You can specify either the Amazon Resource Name (ARN) or the
	// friendly name of the secret.
	//
	// If you specify an ARN, we generally recommend that you specify a complete
	// ARN. You can specify a partial ARN too—for example, if you don’t include
	// the final hyphen and six random characters that Secrets Manager adds at the
	// end of the ARN when you created the secret. A partial ARN match can work
	// as long as it uniquely matches only one secret. However, if your secret has
	// a name that ends in a hyphen followed by six characters (before Secrets Manager
	// adds the hyphen and six characters to the ARN) and you try to use that as
	// a partial ARN, then those characters cause Secrets Manager to assume that
	// you’re specifying a complete ARN. This confusion can cause unexpected results.
	// To avoid this situation, we recommend that you don’t create secret names
	// that end with a hyphen followed by six characters.
	//
	// SecretId is a required field
	SecretId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetResourcePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourcePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourcePolicyInput"}

	if s.SecretId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretId"))
	}
	if s.SecretId != nil && len(*s.SecretId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/GetResourcePolicyResponse
type GetResourcePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the secret that the resource-based policy was retrieved for.
	ARN *string `json:"secretsmanager:GetResourcePolicyOutput:ARN" min:"20" type:"string"`

	// The friendly name of the secret that the resource-based policy was retrieved
	// for.
	Name *string `json:"secretsmanager:GetResourcePolicyOutput:Name" min:"1" type:"string"`

	// A JSON-formatted string that describes the permissions that are associated
	// with the attached secret. These permissions are combined with any permissions
	// that are associated with the user or role that attempts to access this secret.
	// The combined permissions specify who can access the secret and what actions
	// they can perform. For more information, see Authentication and Access Control
	// for AWS Secrets Manager (http://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html)
	// in the AWS Secrets Manager User Guide.
	ResourcePolicy *string `json:"secretsmanager:GetResourcePolicyOutput:ResourcePolicy" min:"1" type:"string"`
}

// String returns the string representation
func (s GetResourcePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetResourcePolicy = "GetResourcePolicy"

// GetResourcePolicyRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Retrieves the JSON text of the resource-based policy document that's attached
// to the specified secret. The JSON request string input and response output
// are shown formatted with white space and line breaks for better readability.
// Submit your input as a single line JSON string.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:GetResourcePolicy
//
// Related operations
//
//    * To attach a resource policy to a secret, use PutResourcePolicy.
//
//    * To delete the resource-based policy that's attached to a secret, use
//    DeleteResourcePolicy.
//
//    * To list all of the currently available secrets, use ListSecrets.
//
//    // Example sending a request using GetResourcePolicyRequest.
//    req := client.GetResourcePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/GetResourcePolicy
func (c *Client) GetResourcePolicyRequest(input *GetResourcePolicyInput) GetResourcePolicyRequest {
	op := &aws.Operation{
		Name:       opGetResourcePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetResourcePolicyInput{}
	}

	req := c.newRequest(op, input, &GetResourcePolicyOutput{})
	return GetResourcePolicyRequest{Request: req, Input: input, Copy: c.GetResourcePolicyRequest}
}

// GetResourcePolicyRequest is the request type for the
// GetResourcePolicy API operation.
type GetResourcePolicyRequest struct {
	*aws.Request
	Input *GetResourcePolicyInput
	Copy  func(*GetResourcePolicyInput) GetResourcePolicyRequest
}

// Send marshals and sends the GetResourcePolicy API request.
func (r GetResourcePolicyRequest) Send(ctx context.Context) (*GetResourcePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourcePolicyResponse{
		GetResourcePolicyOutput: r.Request.Data.(*GetResourcePolicyOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetResourcePolicyResponse is the response type for the
// GetResourcePolicy API operation.
type GetResourcePolicyResponse struct {
	*GetResourcePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResourcePolicy request.
func (r *GetResourcePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
