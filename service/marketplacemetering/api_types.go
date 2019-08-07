// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacemetering

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// A UsageRecord indicates a quantity of usage for a given product, customer,
// dimension and time.
//
// Multiple requests with the same UsageRecords as input will be deduplicated
// to prevent double charges.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/UsageRecord
type UsageRecord struct {
	_ struct{} `type:"structure"`

	// The CustomerIdentifier is obtained through the ResolveCustomer operation
	// and represents an individual buyer in your application.
	//
	// CustomerIdentifier is a required field
	CustomerIdentifier *string `json:"metering.marketplace:UsageRecord:CustomerIdentifier" min:"1" type:"string" required:"true"`

	// During the process of registering a product on AWS Marketplace, up to eight
	// dimensions are specified. These represent different units of value in your
	// application.
	//
	// Dimension is a required field
	Dimension *string `json:"metering.marketplace:UsageRecord:Dimension" min:"1" type:"string" required:"true"`

	// The quantity of usage consumed by the customer for the given dimension and
	// time. Defaults to 0 if not specified.
	Quantity *int64 `json:"metering.marketplace:UsageRecord:Quantity" type:"integer"`

	// Timestamp, in UTC, for which the usage is being reported.
	//
	// Your application can meter usage for up to one hour in the past. Make sure
	// the timestamp value is not before the start of the software usage.
	//
	// Timestamp is a required field
	Timestamp *time.Time `json:"metering.marketplace:UsageRecord:Timestamp" type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s UsageRecord) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UsageRecord) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UsageRecord"}

	if s.CustomerIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomerIdentifier"))
	}
	if s.CustomerIdentifier != nil && len(*s.CustomerIdentifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomerIdentifier", 1))
	}

	if s.Dimension == nil {
		invalidParams.Add(aws.NewErrParamRequired("Dimension"))
	}
	if s.Dimension != nil && len(*s.Dimension) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Dimension", 1))
	}

	if s.Timestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("Timestamp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A UsageRecordResult indicates the status of a given UsageRecord processed
// by BatchMeterUsage.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/UsageRecordResult
type UsageRecordResult struct {
	_ struct{} `type:"structure"`

	// The MeteringRecordId is a unique identifier for this metering event.
	MeteringRecordId *string `json:"metering.marketplace:UsageRecordResult:MeteringRecordId" type:"string"`

	// The UsageRecordResult Status indicates the status of an individual UsageRecord
	// processed by BatchMeterUsage.
	//
	//    * Success- The UsageRecord was accepted and honored by BatchMeterUsage.
	//
	//    * CustomerNotSubscribed- The CustomerIdentifier specified is not subscribed
	//    to your product. The UsageRecord was not honored. Future UsageRecords
	//    for this customer will fail until the customer subscribes to your product.
	//
	//    * DuplicateRecord- Indicates that the UsageRecord was invalid and not
	//    honored. A previously metered UsageRecord had the same customer, dimension,
	//    and time, but a different quantity.
	Status UsageRecordResultStatus `json:"metering.marketplace:UsageRecordResult:Status" type:"string" enum:"true"`

	// The UsageRecord that was part of the BatchMeterUsage request.
	UsageRecord *UsageRecord `json:"metering.marketplace:UsageRecordResult:UsageRecord" type:"structure"`
}

// String returns the string representation
func (s UsageRecordResult) String() string {
	return awsutil.Prettify(s)
}
