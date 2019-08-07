// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/CreateAccountRequest
type CreateAccountInput struct {
	_ struct{} `type:"structure"`

	// The friendly name of the member account.
	//
	// AccountName is a required field
	AccountName *string `min:"1" type:"string" required:"true"`

	// The email address of the owner to assign to the new member account. This
	// email address must not already be associated with another AWS account. You
	// must use a valid email address to complete account creation. You can't access
	// the root user of the account or remove an account that was created with an
	// invalid email address.
	//
	// Email is a required field
	Email *string `min:"6" type:"string" required:"true"`

	// If set to ALLOW, the new account enables IAM users to access account billing
	// information if they have the required permissions. If set to DENY, only the
	// root user of the new account can access account billing information. For
	// more information, see Activating Access to the Billing and Cost Management
	// Console (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html#ControllingAccessWebsite-Activate)
	// in the AWS Billing and Cost Management User Guide.
	//
	// If you don't specify this parameter, the value defaults to ALLOW, and IAM
	// users and roles with the required permissions can access billing information
	// for the new account.
	IamUserAccessToBilling IAMUserAccessToBilling `type:"string" enum:"true"`

	// (Optional)
	//
	// The name of an IAM role that AWS Organizations automatically preconfigures
	// in the new member account. This role trusts the master account, allowing
	// users in the master account to assume the role, as permitted by the master
	// account administrator. The role has administrator permissions in the new
	// member account.
	//
	// If you don't specify this parameter, the role name defaults to OrganizationAccountAccessRole.
	//
	// For more information about how to use this role to access the member account,
	// see Accessing and Administering the Member Accounts in Your Organization
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html#orgs_manage_accounts_create-cross-account-role)
	// in the AWS Organizations User Guide, and steps 2 and 3 in Tutorial: Delegate
	// Access Across AWS Accounts Using IAM Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html)
	// in the IAM User Guide.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter is a string of characters that can consist of uppercase letters,
	// lowercase letters, digits with no spaces, and any of the following characters:
	// =,.@-
	RoleName *string `type:"string"`
}

// String returns the string representation
func (s CreateAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAccountInput"}

	if s.AccountName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountName"))
	}
	if s.AccountName != nil && len(*s.AccountName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountName", 1))
	}

	if s.Email == nil {
		invalidParams.Add(aws.NewErrParamRequired("Email"))
	}
	if s.Email != nil && len(*s.Email) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("Email", 6))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/CreateAccountResponse
type CreateAccountOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the request to create an account.
	// This response structure might not be fully populated when you first receive
	// it because account creation is an asynchronous process. You can pass the
	// returned CreateAccountStatus ID as a parameter to DescribeCreateAccountStatus
	// to get status about the progress of the request at later times. You can also
	// check the AWS CloudTrail log for the CreateAccountResult event. For more
	// information, see Monitoring the Activity in Your Organization (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_monitoring.html)
	// in the AWS Organizations User Guide.
	CreateAccountStatus *CreateAccountStatus `json:"organizations:CreateAccountOutput:CreateAccountStatus" type:"structure"`
}

// String returns the string representation
func (s CreateAccountOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateAccount = "CreateAccount"

// CreateAccountRequest returns a request value for making API operation for
// AWS Organizations.
//
// Creates an AWS account that is automatically a member of the organization
// whose credentials made the request. This is an asynchronous request that
// AWS performs in the background. Because CreateAccount operates asynchronously,
// it can return a successful completion message even though account initialization
// might still be in progress. You might need to wait a few minutes before you
// can successfully access the account. To check the status of the request,
// do one of the following:
//
//    * Use the OperationId response element from this operation to provide
//    as a parameter to the DescribeCreateAccountStatus operation.
//
//    * Check the AWS CloudTrail log for the CreateAccountResult event. For
//    information on using AWS CloudTrail with AWS Organizations, see Monitoring
//    the Activity in Your Organization (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_monitoring.html)
//    in the AWS Organizations User Guide.
//
// The user who calls the API to create an account must have the organizations:CreateAccount
// permission. If you enabled all features in the organization, AWS Organizations
// will create the required service-linked role named AWSServiceRoleForOrganizations.
// For more information, see AWS Organizations and Service-Linked Roles (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html#orgs_integrate_services-using_slrs)
// in the AWS Organizations User Guide.
//
// AWS Organizations preconfigures the new member account with a role (named
// OrganizationAccountAccessRole by default) that grants users in the master
// account administrator permissions in the new member account. Principals in
// the master account can assume the role. AWS Organizations clones the company
// name and address information for the new account from the organization's
// master account.
//
// This operation can be called only from the organization's master account.
//
// For more information about creating accounts, see Creating an AWS Account
// in Your Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_create.html)
// in the AWS Organizations User Guide.
//
//    * When you create an account in an organization using the AWS Organizations
//    console, API, or CLI commands, the information required for the account
//    to operate as a standalone account, such as a payment method and signing
//    the end user license agreement (EULA) is not automatically collected.
//    If you must remove an account from your organization later, you can do
//    so only after you provide the missing information. Follow the steps at
//    To leave an organization as a member account (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_remove.html#leave-without-all-info)
//    in the AWS Organizations User Guide.
//
//    * If you get an exception that indicates that you exceeded your account
//    limits for the organization, contact AWS Support (https://console.aws.amazon.com/support/home#/).
//
//    * If you get an exception that indicates that the operation failed because
//    your organization is still initializing, wait one hour and then try again.
//    If the error persists, contact AWS Support (https://console.aws.amazon.com/support/home#/).
//
//    * Using CreateAccount to create multiple temporary accounts isn't recommended.
//    You can only close an account from the Billing and Cost Management Console,
//    and you must be signed in as the root user. For information on the requirements
//    and process for closing an account, see Closing an AWS Account (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_close.html)
//    in the AWS Organizations User Guide.
//
// When you create a member account with this operation, you can choose whether
// to create the account with the IAM User and Role Access to Billing Information
// switch enabled. If you enable it, IAM users and roles that have appropriate
// permissions can view billing information for the account. If you disable
// it, only the account root user can access billing information. For information
// about how to disable this switch for an account, see Granting Access to Your
// Billing Information and Tools (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html).
//
//    // Example sending a request using CreateAccountRequest.
//    req := client.CreateAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/CreateAccount
func (c *Client) CreateAccountRequest(input *CreateAccountInput) CreateAccountRequest {
	op := &aws.Operation{
		Name:       opCreateAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAccountInput{}
	}

	req := c.newRequest(op, input, &CreateAccountOutput{})
	return CreateAccountRequest{Request: req, Input: input, Copy: c.CreateAccountRequest}
}

// CreateAccountRequest is the request type for the
// CreateAccount API operation.
type CreateAccountRequest struct {
	*aws.Request
	Input *CreateAccountInput
	Copy  func(*CreateAccountInput) CreateAccountRequest
}

// Send marshals and sends the CreateAccount API request.
func (r CreateAccountRequest) Send(ctx context.Context) (*CreateAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAccountResponse{
		CreateAccountOutput: r.Request.Data.(*CreateAccountOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAccountResponse is the response type for the
// CreateAccount API operation.
type CreateAccountResponse struct {
	*CreateAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAccount request.
func (r *CreateAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
