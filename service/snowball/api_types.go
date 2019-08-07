// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// The address that you want the Snowball or Snowballs associated with a specific
// job to be shipped to. Addresses are validated at the time of creation. The
// address you provide must be located within the serviceable area of your region.
// Although no individual elements of the Address are required, if the address
// is invalid or unsupported, then an exception is thrown.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/Address
type Address struct {
	_ struct{} `type:"structure"`

	// The unique ID for an address.
	AddressId *string `json:"snowball:Address:AddressId" min:"40" type:"string"`

	// The city in an address that a Snowball is to be delivered to.
	City *string `json:"snowball:Address:City" min:"1" type:"string"`

	// The name of the company to receive a Snowball at an address.
	Company *string `json:"snowball:Address:Company" min:"1" type:"string"`

	// The country in an address that a Snowball is to be delivered to.
	Country *string `json:"snowball:Address:Country" min:"1" type:"string"`

	// If the address you are creating is a primary address, then set this option
	// to true. This field is not supported in most regions.
	IsRestricted *bool `json:"snowball:Address:IsRestricted" type:"boolean"`

	// This field is no longer used and the value is ignored.
	Landmark *string `json:"snowball:Address:Landmark" min:"1" type:"string"`

	// The name of a person to receive a Snowball at an address.
	Name *string `json:"snowball:Address:Name" min:"1" type:"string"`

	// The phone number associated with an address that a Snowball is to be delivered
	// to.
	PhoneNumber *string `json:"snowball:Address:PhoneNumber" min:"1" type:"string"`

	// The postal code in an address that a Snowball is to be delivered to.
	PostalCode *string `json:"snowball:Address:PostalCode" min:"1" type:"string"`

	// This field is no longer used and the value is ignored.
	PrefectureOrDistrict *string `json:"snowball:Address:PrefectureOrDistrict" min:"1" type:"string"`

	// The state or province in an address that a Snowball is to be delivered to.
	StateOrProvince *string `json:"snowball:Address:StateOrProvince" min:"1" type:"string"`

	// The first line in a street address that a Snowball is to be delivered to.
	Street1 *string `json:"snowball:Address:Street1" min:"1" type:"string"`

	// The second line in a street address that a Snowball is to be delivered to.
	Street2 *string `json:"snowball:Address:Street2" min:"1" type:"string"`

	// The third line in a street address that a Snowball is to be delivered to.
	Street3 *string `json:"snowball:Address:Street3" min:"1" type:"string"`
}

// String returns the string representation
func (s Address) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Address) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Address"}
	if s.AddressId != nil && len(*s.AddressId) < 40 {
		invalidParams.Add(aws.NewErrParamMinLen("AddressId", 40))
	}
	if s.City != nil && len(*s.City) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("City", 1))
	}
	if s.Company != nil && len(*s.Company) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Company", 1))
	}
	if s.Country != nil && len(*s.Country) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Country", 1))
	}
	if s.Landmark != nil && len(*s.Landmark) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Landmark", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.PhoneNumber != nil && len(*s.PhoneNumber) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PhoneNumber", 1))
	}
	if s.PostalCode != nil && len(*s.PostalCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PostalCode", 1))
	}
	if s.PrefectureOrDistrict != nil && len(*s.PrefectureOrDistrict) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PrefectureOrDistrict", 1))
	}
	if s.StateOrProvince != nil && len(*s.StateOrProvince) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StateOrProvince", 1))
	}
	if s.Street1 != nil && len(*s.Street1) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Street1", 1))
	}
	if s.Street2 != nil && len(*s.Street2) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Street2", 1))
	}
	if s.Street3 != nil && len(*s.Street3) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Street3", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains a cluster's state, a cluster's ID, and other important information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/ClusterListEntry
type ClusterListEntry struct {
	_ struct{} `type:"structure"`

	// The 39-character ID for the cluster that you want to list, for example CID123e4567-e89b-12d3-a456-426655440000.
	ClusterId *string `json:"snowball:ClusterListEntry:ClusterId" min:"1" type:"string"`

	// The current state of this cluster. For information about the state of a specific
	// node, see JobListEntry$JobState.
	ClusterState ClusterState `json:"snowball:ClusterListEntry:ClusterState" type:"string" enum:"true"`

	// The creation date for this cluster.
	CreationDate *time.Time `json:"snowball:ClusterListEntry:CreationDate" type:"timestamp" timestampFormat:"unix"`

	// Defines an optional description of the cluster, for example Environmental
	// Data Cluster-01.
	Description *string `json:"snowball:ClusterListEntry:Description" min:"1" type:"string"`
}

// String returns the string representation
func (s ClusterListEntry) String() string {
	return awsutil.Prettify(s)
}

// Contains metadata about a specific cluster.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/ClusterMetadata
type ClusterMetadata struct {
	_ struct{} `type:"structure"`

	// The automatically generated ID for a specific address.
	AddressId *string `json:"snowball:ClusterMetadata:AddressId" min:"40" type:"string"`

	// The automatically generated ID for a cluster.
	ClusterId *string `json:"snowball:ClusterMetadata:ClusterId" min:"1" type:"string"`

	// The current status of the cluster.
	ClusterState ClusterState `json:"snowball:ClusterMetadata:ClusterState" type:"string" enum:"true"`

	// The creation date for this cluster.
	CreationDate *time.Time `json:"snowball:ClusterMetadata:CreationDate" type:"timestamp" timestampFormat:"unix"`

	// The optional description of the cluster.
	Description *string `json:"snowball:ClusterMetadata:Description" min:"1" type:"string"`

	// The ID of the address that you want a cluster shipped to, after it will be
	// shipped to its primary address. This field is not supported in most regions.
	ForwardingAddressId *string `json:"snowball:ClusterMetadata:ForwardingAddressId" min:"40" type:"string"`

	// The type of job for this cluster. Currently, the only job type supported
	// for clusters is LOCAL_USE.
	JobType JobType `json:"snowball:ClusterMetadata:JobType" type:"string" enum:"true"`

	// The KmsKeyARN Amazon Resource Name (ARN) associated with this cluster. This
	// ARN was created using the CreateKey (http://docs.aws.amazon.com/kms/latest/APIReference/API_CreateKey.html)
	// API action in AWS Key Management Service (AWS KMS).
	KmsKeyARN *string `json:"snowball:ClusterMetadata:KmsKeyARN" type:"string"`

	// The Amazon Simple Notification Service (Amazon SNS) notification settings
	// for this cluster.
	Notification *Notification `json:"snowball:ClusterMetadata:Notification" type:"structure"`

	// The arrays of JobResource objects that can include updated S3Resource objects
	// or LambdaResource objects.
	Resources *JobResource `json:"snowball:ClusterMetadata:Resources" type:"structure"`

	// The role ARN associated with this cluster. This ARN was created using the
	// CreateRole (http://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html)
	// API action in AWS Identity and Access Management (IAM).
	RoleARN *string `json:"snowball:ClusterMetadata:RoleARN" type:"string"`

	// The shipping speed for each node in this cluster. This speed doesn't dictate
	// how soon you'll get each device, rather it represents how quickly each device
	// moves to its destination while in transit. Regional shipping speeds are as
	// follows:
	//
	//    * In Australia, you have access to express shipping. Typically, devices
	//    shipped express are delivered in about a day.
	//
	//    * In the European Union (EU), you have access to express shipping. Typically,
	//    devices shipped express are delivered in about a day. In addition, most
	//    countries in the EU have access to standard shipping, which typically
	//    takes less than a week, one way.
	//
	//    * In India, devices are delivered in one to seven days.
	//
	//    * In the US, you have access to one-day shipping and two-day shipping.
	ShippingOption ShippingOption `json:"snowball:ClusterMetadata:ShippingOption" type:"string" enum:"true"`

	// The type of AWS Snowball device to use for this cluster. The only supported
	// device types for cluster jobs are EDGE, EDGE_C, and EDGE_CG.
	SnowballType SnowballType `json:"snowball:ClusterMetadata:SnowballType" type:"string" enum:"true"`
}

// String returns the string representation
func (s ClusterMetadata) String() string {
	return awsutil.Prettify(s)
}

// A JSON-formatted object that describes a compatible Amazon Machine Image
// (AMI). For more information on compatible AMIs, see Using Amazon EC2 Compute
// Instances (http://docs.aws.amazon.com/snowball/latest/developer-guide/using-ec2.html)
// in the AWS Snowball Developer Guide.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/CompatibleImage
type CompatibleImage struct {
	_ struct{} `type:"structure"`

	// The unique identifier for an individual Snowball Edge AMI.
	AmiId *string `json:"snowball:CompatibleImage:AmiId" min:"1" type:"string"`

	// The optional name of a compatible image.
	Name *string `json:"snowball:CompatibleImage:Name" min:"1" type:"string"`
}

// String returns the string representation
func (s CompatibleImage) String() string {
	return awsutil.Prettify(s)
}

// Defines the real-time status of a Snowball's data transfer while the device
// is at AWS. This data is only available while a job has a JobState value of
// InProgress, for both import and export jobs.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/DataTransfer
type DataTransfer struct {
	_ struct{} `type:"structure"`

	// The number of bytes transferred between a Snowball and Amazon S3.
	BytesTransferred *int64 `json:"snowball:DataTransfer:BytesTransferred" type:"long"`

	// The number of objects transferred between a Snowball and Amazon S3.
	ObjectsTransferred *int64 `json:"snowball:DataTransfer:ObjectsTransferred" type:"long"`

	// The total bytes of data for a transfer between a Snowball and Amazon S3.
	// This value is set to 0 (zero) until all the keys that will be transferred
	// have been listed.
	TotalBytes *int64 `json:"snowball:DataTransfer:TotalBytes" type:"long"`

	// The total number of objects for a transfer between a Snowball and Amazon
	// S3. This value is set to 0 (zero) until all the keys that will be transferred
	// have been listed.
	TotalObjects *int64 `json:"snowball:DataTransfer:TotalObjects" type:"long"`
}

// String returns the string representation
func (s DataTransfer) String() string {
	return awsutil.Prettify(s)
}

// A JSON-formatted object that contains the IDs for an Amazon Machine Image
// (AMI), including the Amazon EC2 AMI ID and the Snowball Edge AMI ID. Each
// AMI has these two IDs to simplify identifying the AMI in both the AWS Cloud
// and on the device.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/Ec2AmiResource
type Ec2AmiResource struct {
	_ struct{} `type:"structure"`

	// The ID of the AMI in Amazon EC2.
	//
	// AmiId is a required field
	AmiId *string `json:"snowball:Ec2AmiResource:AmiId" min:"12" type:"string" required:"true"`

	// The ID of the AMI on the supported device.
	SnowballAmiId *string `json:"snowball:Ec2AmiResource:SnowballAmiId" min:"1" type:"string"`
}

// String returns the string representation
func (s Ec2AmiResource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Ec2AmiResource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Ec2AmiResource"}

	if s.AmiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AmiId"))
	}
	if s.AmiId != nil && len(*s.AmiId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AmiId", 12))
	}
	if s.SnowballAmiId != nil && len(*s.SnowballAmiId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnowballAmiId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The container for the EventTriggerDefinition$EventResourceARN.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/EventTriggerDefinition
type EventTriggerDefinition struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for any local Amazon S3 resource that is an
	// AWS Lambda function's event trigger associated with this job.
	EventResourceARN *string `json:"snowball:EventTriggerDefinition:EventResourceARN" type:"string"`
}

// String returns the string representation
func (s EventTriggerDefinition) String() string {
	return awsutil.Prettify(s)
}

// Each JobListEntry object contains a job's state, a job's ID, and a value
// that indicates whether the job is a job part, in the case of an export job.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/JobListEntry
type JobListEntry struct {
	_ struct{} `type:"structure"`

	// The creation date for this job.
	CreationDate *time.Time `json:"snowball:JobListEntry:CreationDate" type:"timestamp" timestampFormat:"unix"`

	// The optional description of this specific job, for example Important Photos
	// 2016-08-11.
	Description *string `json:"snowball:JobListEntry:Description" min:"1" type:"string"`

	// A value that indicates that this job is a master job. A master job represents
	// a successful request to create an export job. Master jobs aren't associated
	// with any Snowballs. Instead, each master job will have at least one job part,
	// and each job part is associated with a Snowball. It might take some time
	// before the job parts associated with a particular master job are listed,
	// because they are created after the master job is created.
	IsMaster *bool `json:"snowball:JobListEntry:IsMaster" type:"boolean"`

	// The automatically generated ID for a job, for example JID123e4567-e89b-12d3-a456-426655440000.
	JobId *string `json:"snowball:JobListEntry:JobId" min:"1" type:"string"`

	// The current state of this job.
	JobState JobState `json:"snowball:JobListEntry:JobState" type:"string" enum:"true"`

	// The type of job.
	JobType JobType `json:"snowball:JobListEntry:JobType" type:"string" enum:"true"`

	// The type of device used with this job.
	SnowballType SnowballType `json:"snowball:JobListEntry:SnowballType" type:"string" enum:"true"`
}

// String returns the string representation
func (s JobListEntry) String() string {
	return awsutil.Prettify(s)
}

// Contains job logs. Whenever Snowball is used to import data into or export
// data out of Amazon S3, you'll have the option of downloading a PDF job report.
// Job logs are returned as a part of the response syntax of the DescribeJob
// action in the JobMetadata data type. The job logs can be accessed for up
// to 60 minutes after this request has been made. To access any of the job
// logs after 60 minutes have passed, you'll have to make another call to the
// DescribeJob action.
//
// For import jobs, the PDF job report becomes available at the end of the import
// process. For export jobs, your job report typically becomes available while
// the Snowball for your job part is being delivered to you.
//
// The job report provides you insight into the state of your Amazon S3 data
// transfer. The report includes details about your job or job part for your
// records.
//
// For deeper visibility into the status of your transferred objects, you can
// look at the two associated logs: a success log and a failure log. The logs
// are saved in comma-separated value (CSV) format, and the name of each log
// includes the ID of the job or job part that the log describes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/JobLogs
type JobLogs struct {
	_ struct{} `type:"structure"`

	// A link to an Amazon S3 presigned URL where the job completion report is located.
	JobCompletionReportURI *string `json:"snowball:JobLogs:JobCompletionReportURI" min:"1" type:"string"`

	// A link to an Amazon S3 presigned URL where the job failure log is located.
	JobFailureLogURI *string `json:"snowball:JobLogs:JobFailureLogURI" min:"1" type:"string"`

	// A link to an Amazon S3 presigned URL where the job success log is located.
	JobSuccessLogURI *string `json:"snowball:JobLogs:JobSuccessLogURI" min:"1" type:"string"`
}

// String returns the string representation
func (s JobLogs) String() string {
	return awsutil.Prettify(s)
}

// Contains information about a specific job including shipping information,
// job status, and other important metadata. This information is returned as
// a part of the response syntax of the DescribeJob action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/JobMetadata
type JobMetadata struct {
	_ struct{} `type:"structure"`

	// The ID for the address that you want the Snowball shipped to.
	AddressId *string `json:"snowball:JobMetadata:AddressId" min:"40" type:"string"`

	// The 39-character ID for the cluster, for example CID123e4567-e89b-12d3-a456-426655440000.
	ClusterId *string `json:"snowball:JobMetadata:ClusterId" min:"1" type:"string"`

	// The creation date for this job.
	CreationDate *time.Time `json:"snowball:JobMetadata:CreationDate" type:"timestamp" timestampFormat:"unix"`

	// A value that defines the real-time status of a Snowball's data transfer while
	// the device is at AWS. This data is only available while a job has a JobState
	// value of InProgress, for both import and export jobs.
	DataTransferProgress *DataTransfer `json:"snowball:JobMetadata:DataTransferProgress" type:"structure"`

	// The description of the job, provided at job creation.
	Description *string `json:"snowball:JobMetadata:Description" min:"1" type:"string"`

	// The ID of the address that you want a job shipped to, after it will be shipped
	// to its primary address. This field is not supported in most regions.
	ForwardingAddressId *string `json:"snowball:JobMetadata:ForwardingAddressId" min:"40" type:"string"`

	// The automatically generated ID for a job, for example JID123e4567-e89b-12d3-a456-426655440000.
	JobId *string `json:"snowball:JobMetadata:JobId" min:"1" type:"string"`

	// Links to Amazon S3 presigned URLs for the job report and logs. For import
	// jobs, the PDF job report becomes available at the end of the import process.
	// For export jobs, your job report typically becomes available while the Snowball
	// for your job part is being delivered to you.
	JobLogInfo *JobLogs `json:"snowball:JobMetadata:JobLogInfo" type:"structure"`

	// The current status of the jobs.
	JobState JobState `json:"snowball:JobMetadata:JobState" type:"string" enum:"true"`

	// The type of job.
	JobType JobType `json:"snowball:JobMetadata:JobType" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) for the AWS Key Management Service (AWS KMS)
	// key associated with this job. This ARN was created using the CreateKey (http://docs.aws.amazon.com/kms/latest/APIReference/API_CreateKey.html)
	// API action in AWS KMS.
	KmsKeyARN *string `json:"snowball:JobMetadata:KmsKeyARN" type:"string"`

	// The Amazon Simple Notification Service (Amazon SNS) notification settings
	// associated with a specific job. The Notification object is returned as a
	// part of the response syntax of the DescribeJob action in the JobMetadata
	// data type.
	Notification *Notification `json:"snowball:JobMetadata:Notification" type:"structure"`

	// An array of S3Resource objects. Each S3Resource object represents an Amazon
	// S3 bucket that your transferred data will be exported from or imported into.
	Resources *JobResource `json:"snowball:JobMetadata:Resources" type:"structure"`

	// The role ARN associated with this job. This ARN was created using the CreateRole
	// (http://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html)
	// API action in AWS Identity and Access Management (IAM).
	RoleARN *string `json:"snowball:JobMetadata:RoleARN" type:"string"`

	// A job's shipping information, including inbound and outbound tracking numbers
	// and shipping speed options.
	ShippingDetails *ShippingDetails `json:"snowball:JobMetadata:ShippingDetails" type:"structure"`

	// The Snowball capacity preference for this job, specified at job creation.
	// In US regions, you can choose between 50 TB and 80 TB Snowballs. All other
	// regions use 80 TB capacity Snowballs.
	SnowballCapacityPreference SnowballCapacity `json:"snowball:JobMetadata:SnowballCapacityPreference" type:"string" enum:"true"`

	// The type of device used with this job.
	SnowballType SnowballType `json:"snowball:JobMetadata:SnowballType" type:"string" enum:"true"`
}

// String returns the string representation
func (s JobMetadata) String() string {
	return awsutil.Prettify(s)
}

// Contains an array of AWS resource objects. Each object represents an Amazon
// S3 bucket, an AWS Lambda function, or an Amazon Machine Image (AMI) based
// on Amazon EC2 that is associated with a particular job.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/JobResource
type JobResource struct {
	_ struct{} `type:"structure"`

	// The Amazon Machine Images (AMIs) associated with this job.
	Ec2AmiResources []Ec2AmiResource `json:"snowball:JobResource:Ec2AmiResources" type:"list"`

	// The Python-language Lambda functions for this job.
	LambdaResources []LambdaResource `json:"snowball:JobResource:LambdaResources" type:"list"`

	// An array of S3Resource objects.
	S3Resources []S3Resource `json:"snowball:JobResource:S3Resources" type:"list"`
}

// String returns the string representation
func (s JobResource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *JobResource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "JobResource"}
	if s.Ec2AmiResources != nil {
		for i, v := range s.Ec2AmiResources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Ec2AmiResources", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.S3Resources != nil {
		for i, v := range s.S3Resources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "S3Resources", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains a key range. For export jobs, a S3Resource object can have an optional
// KeyRange value. The length of the range is defined at job creation, and has
// either an inclusive BeginMarker, an inclusive EndMarker, or both. Ranges
// are UTF-8 binary sorted.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/KeyRange
type KeyRange struct {
	_ struct{} `type:"structure"`

	// The key that starts an optional key range for an export job. Ranges are inclusive
	// and UTF-8 binary sorted.
	BeginMarker *string `json:"snowball:KeyRange:BeginMarker" min:"1" type:"string"`

	// The key that ends an optional key range for an export job. Ranges are inclusive
	// and UTF-8 binary sorted.
	EndMarker *string `json:"snowball:KeyRange:EndMarker" min:"1" type:"string"`
}

// String returns the string representation
func (s KeyRange) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *KeyRange) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "KeyRange"}
	if s.BeginMarker != nil && len(*s.BeginMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BeginMarker", 1))
	}
	if s.EndMarker != nil && len(*s.EndMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EndMarker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Identifies
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/LambdaResource
type LambdaResource struct {
	_ struct{} `type:"structure"`

	// The array of ARNs for S3Resource objects to trigger the LambdaResource objects
	// associated with this job.
	EventTriggers []EventTriggerDefinition `json:"snowball:LambdaResource:EventTriggers" type:"list"`

	// An Amazon Resource Name (ARN) that represents an AWS Lambda function to be
	// triggered by PUT object actions on the associated local Amazon S3 resource.
	LambdaArn *string `json:"snowball:LambdaResource:LambdaArn" type:"string"`
}

// String returns the string representation
func (s LambdaResource) String() string {
	return awsutil.Prettify(s)
}

// The Amazon Simple Notification Service (Amazon SNS) notification settings
// associated with a specific job. The Notification object is returned as a
// part of the response syntax of the DescribeJob action in the JobMetadata
// data type.
//
// When the notification settings are defined during job creation, you can choose
// to notify based on a specific set of job states using the JobStatesToNotify
// array of strings, or you can specify that you want to have Amazon SNS notifications
// sent out for all job states with NotifyAll set to true.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/Notification
type Notification struct {
	_ struct{} `type:"structure"`

	// The list of job states that will trigger a notification for this job.
	JobStatesToNotify []JobState `json:"snowball:Notification:JobStatesToNotify" type:"list"`

	// Any change in job state will trigger a notification for this job.
	NotifyAll *bool `json:"snowball:Notification:NotifyAll" type:"boolean"`

	// The new SNS TopicArn that you want to associate with this job. You can create
	// Amazon Resource Names (ARNs) for topics by using the CreateTopic (http://docs.aws.amazon.com/sns/latest/api/API_CreateTopic.html)
	// Amazon SNS API action.
	//
	// You can subscribe email addresses to an Amazon SNS topic through the AWS
	// Management Console, or by using the Subscribe (http://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)
	// AWS Simple Notification Service (SNS) API action.
	SnsTopicARN *string `json:"snowball:Notification:SnsTopicARN" type:"string"`
}

// String returns the string representation
func (s Notification) String() string {
	return awsutil.Prettify(s)
}

// Each S3Resource object represents an Amazon S3 bucket that your transferred
// data will be exported from or imported into. For export jobs, this object
// can have an optional KeyRange value. The length of the range is defined at
// job creation, and has either an inclusive BeginMarker, an inclusive EndMarker,
// or both. Ranges are UTF-8 binary sorted.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/S3Resource
type S3Resource struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of an Amazon S3 bucket.
	BucketArn *string `json:"snowball:S3Resource:BucketArn" type:"string"`

	// For export jobs, you can provide an optional KeyRange within a specific Amazon
	// S3 bucket. The length of the range is defined at job creation, and has either
	// an inclusive BeginMarker, an inclusive EndMarker, or both. Ranges are UTF-8
	// binary sorted.
	KeyRange *KeyRange `json:"snowball:S3Resource:KeyRange" type:"structure"`
}

// String returns the string representation
func (s S3Resource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *S3Resource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "S3Resource"}
	if s.KeyRange != nil {
		if err := s.KeyRange.Validate(); err != nil {
			invalidParams.AddNested("KeyRange", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The Status and TrackingNumber information for an inbound or outbound shipment.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/Shipment
type Shipment struct {
	_ struct{} `type:"structure"`

	// Status information for a shipment.
	Status *string `json:"snowball:Shipment:Status" min:"1" type:"string"`

	// The tracking number for this job. Using this tracking number with your region's
	// carrier's website, you can track a Snowball as the carrier transports it.
	//
	// For India, the carrier is Amazon Logistics. For all other regions, UPS is
	// the carrier.
	TrackingNumber *string `json:"snowball:Shipment:TrackingNumber" min:"1" type:"string"`
}

// String returns the string representation
func (s Shipment) String() string {
	return awsutil.Prettify(s)
}

// A job's shipping information, including inbound and outbound tracking numbers
// and shipping speed options.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/ShippingDetails
type ShippingDetails struct {
	_ struct{} `type:"structure"`

	// The Status and TrackingNumber values for a Snowball being returned to AWS
	// for a particular job.
	InboundShipment *Shipment `json:"snowball:ShippingDetails:InboundShipment" type:"structure"`

	// The Status and TrackingNumber values for a Snowball being delivered to the
	// address that you specified for a particular job.
	OutboundShipment *Shipment `json:"snowball:ShippingDetails:OutboundShipment" type:"structure"`

	// The shipping speed for a particular job. This speed doesn't dictate how soon
	// you'll get the Snowball from the job's creation date. This speed represents
	// how quickly it moves to its destination while in transit. Regional shipping
	// speeds are as follows:
	//
	//    * In Australia, you have access to express shipping. Typically, Snowballs
	//    shipped express are delivered in about a day.
	//
	//    * In the European Union (EU), you have access to express shipping. Typically,
	//    Snowballs shipped express are delivered in about a day. In addition, most
	//    countries in the EU have access to standard shipping, which typically
	//    takes less than a week, one way.
	//
	//    * In India, Snowballs are delivered in one to seven days.
	//
	//    * In the United States of America (US), you have access to one-day shipping
	//    and two-day shipping.
	ShippingOption ShippingOption `json:"snowball:ShippingDetails:ShippingOption" type:"string" enum:"true"`
}

// String returns the string representation
func (s ShippingDetails) String() string {
	return awsutil.Prettify(s)
}
