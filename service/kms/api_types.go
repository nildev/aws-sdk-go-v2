// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Contains information about an alias.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/AliasListEntry
type AliasListEntry struct {
	_ struct{} `type:"structure"`

	// String that contains the key ARN.
	AliasArn *string `json:"kms:AliasListEntry:AliasArn" min:"20" type:"string"`

	// String that contains the alias. This value begins with alias/.
	AliasName *string `json:"kms:AliasListEntry:AliasName" min:"1" type:"string"`

	// String that contains the key identifier referred to by the alias.
	TargetKeyId *string `json:"kms:AliasListEntry:TargetKeyId" min:"1" type:"string"`
}

// String returns the string representation
func (s AliasListEntry) String() string {
	return awsutil.Prettify(s)
}

// Contains information about each custom key store in the custom key store
// list.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/CustomKeyStoresListEntry
type CustomKeyStoresListEntry struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the AWS CloudHSM cluster that is associated with
	// the custom key store.
	CloudHsmClusterId *string `json:"kms:CustomKeyStoresListEntry:CloudHsmClusterId" min:"19" type:"string"`

	// Describes the connection error. Valid values are:
	//
	//    * CLUSTER_NOT_FOUND - AWS KMS cannot find the AWS CloudHSM cluster with
	//    the specified cluster ID.
	//
	//    * INSUFFICIENT_CLOUDHSM_HSMS - The associated AWS CloudHSM cluster does
	//    not contain any active HSMs. To connect a custom key store to its AWS
	//    CloudHSM cluster, the cluster must contain at least one active HSM.
	//
	//    * INTERNAL_ERROR - AWS KMS could not complete the request due to an internal
	//    error. Retry the request. For ConnectCustomKeyStore requests, disconnect
	//    the custom key store before trying to connect again.
	//
	//    * INVALID_CREDENTIALS - AWS KMS does not have the correct password for
	//    the kmsuser crypto user in the AWS CloudHSM cluster.
	//
	//    * NETWORK_ERRORS - Network errors are preventing AWS KMS from connecting
	//    to the custom key store.
	//
	//    * USER_LOCKED_OUT - The kmsuser CU account is locked out of the associated
	//    AWS CloudHSM cluster due to too many failed password attempts. Before
	//    you can connect your custom key store to its AWS CloudHSM cluster, you
	//    must change the kmsuser account password and update the password value
	//    for the custom key store.
	//
	// For help with connection failures, see Troubleshooting Custom Key Stores
	// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html)
	// in the AWS Key Management Service Developer Guide.
	ConnectionErrorCode ConnectionErrorCodeType `json:"kms:CustomKeyStoresListEntry:ConnectionErrorCode" type:"string" enum:"true"`

	// Indicates whether the custom key store is connected to its AWS CloudHSM cluster.
	//
	// You can create and use CMKs in your custom key stores only when its connection
	// state is CONNECTED.
	//
	// The value is DISCONNECTED if the key store has never been connected or you
	// use the DisconnectCustomKeyStore operation to disconnect it. If the value
	// is CONNECTED but you are having trouble using the custom key store, make
	// sure that its associated AWS CloudHSM cluster is active and contains at least
	// one active HSM.
	//
	// A value of FAILED indicates that an attempt to connect was unsuccessful.
	// For help resolving a connection failure, see Troubleshooting a Custom Key
	// Store (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html)
	// in the AWS Key Management Service Developer Guide.
	ConnectionState ConnectionStateType `json:"kms:CustomKeyStoresListEntry:ConnectionState" type:"string" enum:"true"`

	// The date and time when the custom key store was created.
	CreationDate *time.Time `json:"kms:CustomKeyStoresListEntry:CreationDate" type:"timestamp" timestampFormat:"unix"`

	// A unique identifier for the custom key store.
	CustomKeyStoreId *string `json:"kms:CustomKeyStoresListEntry:CustomKeyStoreId" min:"1" type:"string"`

	// The user-specified friendly name for the custom key store.
	CustomKeyStoreName *string `json:"kms:CustomKeyStoresListEntry:CustomKeyStoreName" min:"1" type:"string"`

	// The trust anchor certificate of the associated AWS CloudHSM cluster. When
	// you initialize the cluster (https://docs.aws.amazon.com/cloudhsm/latest/userguide/initialize-cluster.html#sign-csr),
	// you create this certificate and save it in the customerCA.crt file.
	TrustAnchorCertificate *string `json:"kms:CustomKeyStoresListEntry:TrustAnchorCertificate" min:"1" type:"string"`
}

// String returns the string representation
func (s CustomKeyStoresListEntry) String() string {
	return awsutil.Prettify(s)
}

// Use this structure to allow cryptographic operations in the grant only when
// the operation request includes the specified encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context).
//
// AWS KMS applies the grant constraints only when the grant allows a cryptographic
// operation that accepts an encryption context as input, such as the following.
//
//    * Encrypt
//
//    * Decrypt
//
//    * GenerateDataKey
//
//    * GenerateDataKeyWithoutPlaintext
//
//    * ReEncrypt
//
// AWS KMS does not apply the grant constraints to other operations, such as
// DescribeKey or ScheduleKeyDeletion.
//
// In a cryptographic operation, the encryption context in the decryption operation
// must be an exact, case-sensitive match for the keys and values in the encryption
// context of the encryption operation. Only the order of the pairs can vary.
//
// However, in a grant constraint, the key in each key-value pair is not case
// sensitive, but the value is case sensitive.
//
// To avoid confusion, do not use multiple encryption context pairs that differ
// only by case. To require a fully case-sensitive encryption context, use the
// kms:EncryptionContext: and kms:EncryptionContextKeys conditions in an IAM
// or key policy. For details, see kms:EncryptionContext: (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-context)
// in the AWS Key Management Service Developer Guide .
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GrantConstraints
type GrantConstraints struct {
	_ struct{} `type:"structure"`

	// A list of key-value pairs that must match the encryption context in the cryptographic
	// operation request. The grant allows the operation only when the encryption
	// context in the request is the same as the encryption context specified in
	// this constraint.
	EncryptionContextEquals map[string]string `json:"kms:GrantConstraints:EncryptionContextEquals" type:"map"`

	// A list of key-value pairs that must be included in the encryption context
	// of the cryptographic operation request. The grant allows the cryptographic
	// operation only when the encryption context in the request includes the key-value
	// pairs specified in this constraint, although it can include additional key-value
	// pairs.
	EncryptionContextSubset map[string]string `json:"kms:GrantConstraints:EncryptionContextSubset" type:"map"`
}

// String returns the string representation
func (s GrantConstraints) String() string {
	return awsutil.Prettify(s)
}

// Contains information about an entry in a list of grants.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GrantListEntry
type GrantListEntry struct {
	_ struct{} `type:"structure"`

	// A list of key-value pairs that must be present in the encryption context
	// of certain subsequent operations that the grant allows.
	Constraints *GrantConstraints `json:"kms:GrantListEntry:Constraints" type:"structure"`

	// The date and time when the grant was created.
	CreationDate *time.Time `json:"kms:GrantListEntry:CreationDate" type:"timestamp" timestampFormat:"unix"`

	// The unique identifier for the grant.
	GrantId *string `json:"kms:GrantListEntry:GrantId" min:"1" type:"string"`

	// The principal that receives the grant's permissions.
	GranteePrincipal *string `json:"kms:GrantListEntry:GranteePrincipal" min:"1" type:"string"`

	// The AWS account under which the grant was issued.
	IssuingAccount *string `json:"kms:GrantListEntry:IssuingAccount" min:"1" type:"string"`

	// The unique identifier for the customer master key (CMK) to which the grant
	// applies.
	KeyId *string `json:"kms:GrantListEntry:KeyId" min:"1" type:"string"`

	// The friendly name that identifies the grant. If a name was provided in the
	// CreateGrant request, that name is returned. Otherwise this value is null.
	Name *string `json:"kms:GrantListEntry:Name" min:"1" type:"string"`

	// The list of operations permitted by the grant.
	Operations []GrantOperation `json:"kms:GrantListEntry:Operations" type:"list"`

	// The principal that can retire the grant.
	RetiringPrincipal *string `json:"kms:GrantListEntry:RetiringPrincipal" min:"1" type:"string"`
}

// String returns the string representation
func (s GrantListEntry) String() string {
	return awsutil.Prettify(s)
}

// Contains information about each entry in the key list.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/KeyListEntry
type KeyListEntry struct {
	_ struct{} `type:"structure"`

	// ARN of the key.
	KeyArn *string `json:"kms:KeyListEntry:KeyArn" min:"20" type:"string"`

	// Unique identifier of the key.
	KeyId *string `json:"kms:KeyListEntry:KeyId" min:"1" type:"string"`
}

// String returns the string representation
func (s KeyListEntry) String() string {
	return awsutil.Prettify(s)
}

// Contains metadata about a customer master key (CMK).
//
// This data type is used as a response element for the CreateKey and DescribeKey
// operations.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/KeyMetadata
type KeyMetadata struct {
	_ struct{} `type:"structure"`

	// The twelve-digit account ID of the AWS account that owns the CMK.
	AWSAccountId *string `json:"kms:KeyMetadata:AWSAccountId" type:"string"`

	// The Amazon Resource Name (ARN) of the CMK. For examples, see AWS Key Management
	// Service (AWS KMS) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms)
	// in the Example ARNs section of the AWS General Reference.
	Arn *string `json:"kms:KeyMetadata:Arn" min:"20" type:"string"`

	// The cluster ID of the AWS CloudHSM cluster that contains the key material
	// for the CMK. When you create a CMK in a custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html),
	// AWS KMS creates the key material for the CMK in the associated AWS CloudHSM
	// cluster. This value is present only when the CMK is created in a custom key
	// store.
	CloudHsmClusterId *string `json:"kms:KeyMetadata:CloudHsmClusterId" min:"19" type:"string"`

	// The date and time when the CMK was created.
	CreationDate *time.Time `json:"kms:KeyMetadata:CreationDate" type:"timestamp" timestampFormat:"unix"`

	// A unique identifier for the custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// that contains the CMK. This value is present only when the CMK is created
	// in a custom key store.
	CustomKeyStoreId *string `json:"kms:KeyMetadata:CustomKeyStoreId" min:"1" type:"string"`

	// The date and time after which AWS KMS deletes the CMK. This value is present
	// only when KeyState is PendingDeletion.
	DeletionDate *time.Time `json:"kms:KeyMetadata:DeletionDate" type:"timestamp" timestampFormat:"unix"`

	// The description of the CMK.
	Description *string `json:"kms:KeyMetadata:Description" type:"string"`

	// Specifies whether the CMK is enabled. When KeyState is Enabled this value
	// is true, otherwise it is false.
	Enabled *bool `json:"kms:KeyMetadata:Enabled" type:"boolean"`

	// Specifies whether the CMK's key material expires. This value is present only
	// when Origin is EXTERNAL, otherwise this value is omitted.
	ExpirationModel ExpirationModelType `json:"kms:KeyMetadata:ExpirationModel" type:"string" enum:"true"`

	// The globally unique identifier for the CMK.
	//
	// KeyId is a required field
	KeyId *string `json:"kms:KeyMetadata:KeyId" min:"1" type:"string" required:"true"`

	// The manager of the CMK. CMKs in your AWS account are either customer managed
	// or AWS managed. For more information about the difference, see Customer Master
	// Keys (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master_keys)
	// in the AWS Key Management Service Developer Guide.
	KeyManager KeyManagerType `json:"kms:KeyMetadata:KeyManager" type:"string" enum:"true"`

	// The state of the CMK.
	//
	// For more information about how key state affects the use of a CMK, see How
	// Key State Affects the Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
	// in the AWS Key Management Service Developer Guide.
	KeyState KeyState `json:"kms:KeyMetadata:KeyState" type:"string" enum:"true"`

	// The cryptographic operations for which you can use the CMK. The only valid
	// value is ENCRYPT_DECRYPT, which means you can use the CMK to encrypt and
	// decrypt data.
	KeyUsage KeyUsageType `json:"kms:KeyMetadata:KeyUsage" type:"string" enum:"true"`

	// The source of the CMK's key material. When this value is AWS_KMS, AWS KMS
	// created the key material. When this value is EXTERNAL, the key material was
	// imported from your existing key management infrastructure or the CMK lacks
	// key material. When this value is AWS_CLOUDHSM, the key material was created
	// in the AWS CloudHSM cluster associated with a custom key store.
	Origin OriginType `json:"kms:KeyMetadata:Origin" type:"string" enum:"true"`

	// The time at which the imported key material expires. When the key material
	// expires, AWS KMS deletes the key material and the CMK becomes unusable. This
	// value is present only for CMKs whose Origin is EXTERNAL and whose ExpirationModel
	// is KEY_MATERIAL_EXPIRES, otherwise this value is omitted.
	ValidTo *time.Time `json:"kms:KeyMetadata:ValidTo" type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s KeyMetadata) String() string {
	return awsutil.Prettify(s)
}

// A key-value pair. A tag consists of a tag key and a tag value. Tag keys and
// tag values are both required, but tag values can be empty (null) strings.
//
// For information about the rules that apply to tag keys and tag values, see
// User-Defined Tag Restrictions (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html)
// in the AWS Billing and Cost Management User Guide.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// The key of the tag.
	//
	// TagKey is a required field
	TagKey *string `json:"kms:Tag:TagKey" min:"1" type:"string" required:"true"`

	// The value of the tag.
	//
	// TagValue is a required field
	TagValue *string `json:"kms:Tag:TagValue" type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.TagKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKey"))
	}
	if s.TagKey != nil && len(*s.TagKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TagKey", 1))
	}

	if s.TagValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagValue"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
