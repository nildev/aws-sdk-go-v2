// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// The details of a DDoS attack.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/AttackDetail
type AttackDetail struct {
	_ struct{} `type:"structure"`

	// List of counters that describe the attack for the specified time period.
	AttackCounters []SummarizedCounter `type:"list"`

	// The unique identifier (ID) of the attack.
	AttackId *string `min:"1" type:"string"`

	// The array of AttackProperty objects.
	AttackProperties []AttackProperty `type:"list"`

	// The time the attack ended, in Unix time in seconds. For more information
	// see timestamp (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	EndTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// List of mitigation actions taken for the attack.
	Mitigations []Mitigation `type:"list"`

	// The ARN (Amazon Resource Name) of the resource that was attacked.
	ResourceArn *string `min:"1" type:"string"`

	// The time the attack started, in Unix time in seconds. For more information
	// see timestamp (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	StartTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// If applicable, additional detail about the resource being attacked, for example,
	// IP address or URL.
	SubResources []SubResourceSummary `type:"list"`
}

// String returns the string representation
func (s AttackDetail) String() string {
	return awsutil.Prettify(s)
}

// Details of the described attack.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/AttackProperty
type AttackProperty struct {
	_ struct{} `type:"structure"`

	// The type of distributed denial of service (DDoS) event that was observed.
	// NETWORK indicates layer 3 and layer 4 events and APPLICATION indicates layer
	// 7 events.
	AttackLayer AttackLayer `type:"string" enum:"true"`

	// Defines the DDoS attack property information that is provided. The WORDPRESS_PINGBACK_REFLECTOR
	// and WORDPRESS_PINGBACK_SOURCE values are valid only for WordPress reflective
	// pingback DDoS attacks.
	AttackPropertyIdentifier AttackPropertyIdentifier `type:"string" enum:"true"`

	// The array of Contributor objects that includes the top five contributors
	// to an attack.
	TopContributors []Contributor `type:"list"`

	// The total contributions made to this attack by all contributors, not just
	// the five listed in the TopContributors list.
	Total *int64 `type:"long"`

	// The unit of the Value of the contributions.
	Unit Unit `type:"string" enum:"true"`
}

// String returns the string representation
func (s AttackProperty) String() string {
	return awsutil.Prettify(s)
}

// Summarizes all DDoS attacks for a specified time period.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/AttackSummary
type AttackSummary struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the attack.
	AttackId *string `type:"string"`

	// The list of attacks for a specified time period.
	AttackVectors []AttackVectorDescription `type:"list"`

	// The end time of the attack, in Unix time in seconds. For more information
	// see timestamp (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	EndTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The ARN (Amazon Resource Name) of the resource that was attacked.
	ResourceArn *string `type:"string"`

	// The start time of the attack, in Unix time in seconds. For more information
	// see timestamp (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	StartTime *time.Time `type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s AttackSummary) String() string {
	return awsutil.Prettify(s)
}

// Describes the attack.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/AttackVectorDescription
type AttackVectorDescription struct {
	_ struct{} `type:"structure"`

	// The attack type. Valid values:
	//
	//    * UDP_TRAFFIC
	//
	//    * UDP_FRAGMENT
	//
	//    * GENERIC_UDP_REFLECTION
	//
	//    * DNS_REFLECTION
	//
	//    * NTP_REFLECTION
	//
	//    * CHARGEN_REFLECTION
	//
	//    * SSDP_REFLECTION
	//
	//    * PORT_MAPPER
	//
	//    * RIP_REFLECTION
	//
	//    * SNMP_REFLECTION
	//
	//    * MSSQL_REFLECTION
	//
	//    * NET_BIOS_REFLECTION
	//
	//    * SYN_FLOOD
	//
	//    * ACK_FLOOD
	//
	//    * REQUEST_FLOOD
	//
	//    * HTTP_REFLECTION
	//
	//    * UDS_REFLECTION
	//
	//    * MEMCACHED_REFLECTION
	//
	// VectorType is a required field
	VectorType *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AttackVectorDescription) String() string {
	return awsutil.Prettify(s)
}

// A contributor to the attack and their contribution.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/Contributor
type Contributor struct {
	_ struct{} `type:"structure"`

	// The name of the contributor. This is dependent on the AttackPropertyIdentifier.
	// For example, if the AttackPropertyIdentifier is SOURCE_COUNTRY, the Name
	// could be United States.
	Name *string `type:"string"`

	// The contribution of this contributor expressed in Protection units. For example
	// 10,000.
	Value *int64 `type:"long"`
}

// String returns the string representation
func (s Contributor) String() string {
	return awsutil.Prettify(s)
}

// Contact information that the DRT can use to contact you during a suspected
// attack.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/EmergencyContact
type EmergencyContact struct {
	_ struct{} `type:"structure"`

	// An email address that the DRT can use to contact you during a suspected attack.
	//
	// EmailAddress is a required field
	EmailAddress *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s EmergencyContact) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EmergencyContact) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EmergencyContact"}

	if s.EmailAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailAddress"))
	}
	if s.EmailAddress != nil && len(*s.EmailAddress) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EmailAddress", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies how many protections of a given type you can create.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/Limit
type Limit struct {
	_ struct{} `type:"structure"`

	// The maximum number of protections that can be created for the specified Type.
	Max *int64 `type:"long"`

	// The type of protection.
	Type *string `type:"string"`
}

// String returns the string representation
func (s Limit) String() string {
	return awsutil.Prettify(s)
}

// The mitigation applied to a DDoS attack.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/Mitigation
type Mitigation struct {
	_ struct{} `type:"structure"`

	// The name of the mitigation taken for this attack.
	MitigationName *string `type:"string"`
}

// String returns the string representation
func (s Mitigation) String() string {
	return awsutil.Prettify(s)
}

// An object that represents a resource that is under DDoS protection.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/Protection
type Protection struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the protection.
	Id *string `min:"1" type:"string"`

	// The friendly name of the protection. For example, My CloudFront distributions.
	Name *string `min:"1" type:"string"`

	// The ARN (Amazon Resource Name) of the AWS resource that is protected.
	ResourceArn *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Protection) String() string {
	return awsutil.Prettify(s)
}

// The attack information for the specified SubResource.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/SubResourceSummary
type SubResourceSummary struct {
	_ struct{} `type:"structure"`

	// The list of attack types and associated counters.
	AttackVectors []SummarizedAttackVector `type:"list"`

	// The counters that describe the details of the attack.
	Counters []SummarizedCounter `type:"list"`

	// The unique identifier (ID) of the SubResource.
	Id *string `type:"string"`

	// The SubResource type.
	Type SubResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s SubResourceSummary) String() string {
	return awsutil.Prettify(s)
}

// Information about the AWS Shield Advanced subscription for an account.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/Subscription
type Subscription struct {
	_ struct{} `type:"structure"`

	// If ENABLED, the subscription will be automatically renewed at the end of
	// the existing subscription period.
	//
	// When you initally create a subscription, AutoRenew is set to ENABLED. You
	// can change this by submitting an UpdateSubscription request. If the UpdateSubscription
	// request does not included a value for AutoRenew, the existing value for AutoRenew
	// remains unchanged.
	AutoRenew AutoRenew `type:"string" enum:"true"`

	// The date and time your subscription will end.
	EndTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Specifies how many protections of a given type you can create.
	Limits []Limit `type:"list"`

	// The start time of the subscription, in Unix time in seconds. For more information
	// see timestamp (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	StartTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The length, in seconds, of the AWS Shield Advanced subscription for the account.
	TimeCommitmentInSeconds *int64 `type:"long"`
}

// String returns the string representation
func (s Subscription) String() string {
	return awsutil.Prettify(s)
}

// A summary of information about the attack.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/SummarizedAttackVector
type SummarizedAttackVector struct {
	_ struct{} `type:"structure"`

	// The list of counters that describe the details of the attack.
	VectorCounters []SummarizedCounter `type:"list"`

	// The attack type, for example, SNMP reflection or SYN flood.
	//
	// VectorType is a required field
	VectorType *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SummarizedAttackVector) String() string {
	return awsutil.Prettify(s)
}

// The counter that describes a DDoS attack.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/SummarizedCounter
type SummarizedCounter struct {
	_ struct{} `type:"structure"`

	// The average value of the counter for a specified time period.
	Average *float64 `type:"double"`

	// The maximum value of the counter for a specified time period.
	Max *float64 `type:"double"`

	// The number of counters for a specified time period.
	N *int64 `type:"integer"`

	// The counter name.
	Name *string `type:"string"`

	// The total of counter values for a specified time period.
	Sum *float64 `type:"double"`

	// The unit of the counters.
	Unit *string `type:"string"`
}

// String returns the string representation
func (s SummarizedCounter) String() string {
	return awsutil.Prettify(s)
}

// The time range.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/TimeRange
type TimeRange struct {
	_ struct{} `type:"structure"`

	// The start time, in Unix time in seconds. For more information see timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	FromInclusive *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The end time, in Unix time in seconds. For more information see timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	ToExclusive *time.Time `type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s TimeRange) String() string {
	return awsutil.Prettify(s)
}
