// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Provides options for deleting a vault notification configuration from an
// Amazon Glacier vault.
type DeleteVaultNotificationsInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVaultNotificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVaultNotificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVaultNotificationsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVaultNotificationsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteVaultNotificationsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVaultNotificationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVaultNotificationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteVaultNotifications = "DeleteVaultNotifications"

// DeleteVaultNotificationsRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation deletes the notification configuration set for a vault. The
// operation is eventually consistent; that is, it might take some time for
// Amazon S3 Glacier to completely disable the notifications and you might still
// receive some notifications for a short time after you send the delete request.
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For conceptual information and underlying REST API, see Configuring Vault
// Notifications in Amazon S3 Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/configuring-notifications.html)
// and Delete Vault Notification Configuration (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-delete.html)
// in the Amazon S3 Glacier Developer Guide.
//
//    // Example sending a request using DeleteVaultNotificationsRequest.
//    req := client.DeleteVaultNotificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteVaultNotificationsRequest(input *DeleteVaultNotificationsInput) DeleteVaultNotificationsRequest {
	op := &aws.Operation{
		Name:       opDeleteVaultNotifications,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/notification-configuration",
	}

	if input == nil {
		input = &DeleteVaultNotificationsInput{}
	}

	req := c.newRequest(op, input, &DeleteVaultNotificationsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVaultNotificationsRequest{Request: req, Input: input, Copy: c.DeleteVaultNotificationsRequest}
}

// DeleteVaultNotificationsRequest is the request type for the
// DeleteVaultNotifications API operation.
type DeleteVaultNotificationsRequest struct {
	*aws.Request
	Input *DeleteVaultNotificationsInput
	Copy  func(*DeleteVaultNotificationsInput) DeleteVaultNotificationsRequest
}

// Send marshals and sends the DeleteVaultNotifications API request.
func (r DeleteVaultNotificationsRequest) Send(ctx context.Context) (*DeleteVaultNotificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVaultNotificationsResponse{
		DeleteVaultNotificationsOutput: r.Request.Data.(*DeleteVaultNotificationsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVaultNotificationsResponse is the response type for the
// DeleteVaultNotifications API operation.
type DeleteVaultNotificationsResponse struct {
	*DeleteVaultNotificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVaultNotifications request.
func (r *DeleteVaultNotificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
