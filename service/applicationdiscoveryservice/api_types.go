// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Information about agents or connectors that were instructed to start collecting
// data. Information includes the agent/connector ID, a description of the operation,
// and whether the agent/connector configuration was updated.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/AgentConfigurationStatus
type AgentConfigurationStatus struct {
	_ struct{} `type:"structure"`

	// The agent/connector ID.
	AgentId *string `json:"discovery:AgentConfigurationStatus:AgentId" locationName:"agentId" type:"string"`

	// A description of the operation performed.
	Description *string `json:"discovery:AgentConfigurationStatus:Description" locationName:"description" type:"string"`

	// Information about the status of the StartDataCollection and StopDataCollection
	// operations. The system has recorded the data collection operation. The agent/connector
	// receives this command the next time it polls for a new command.
	OperationSucceeded *bool `json:"discovery:AgentConfigurationStatus:OperationSucceeded" locationName:"operationSucceeded" type:"boolean"`
}

// String returns the string representation
func (s AgentConfigurationStatus) String() string {
	return awsutil.Prettify(s)
}

// Information about agents or connectors associated with the user’s AWS account.
// Information includes agent or connector IDs, IP addresses, media access control
// (MAC) addresses, agent or connector health, hostname where the agent or connector
// resides, and agent version for each agent.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/AgentInfo
type AgentInfo struct {
	_ struct{} `type:"structure"`

	// The agent or connector ID.
	AgentId *string `json:"discovery:AgentInfo:AgentId" locationName:"agentId" type:"string"`

	// Network details about the host where the agent or connector resides.
	AgentNetworkInfoList []AgentNetworkInfo `json:"discovery:AgentInfo:AgentNetworkInfoList" locationName:"agentNetworkInfoList" type:"list"`

	// Type of agent.
	AgentType *string `json:"discovery:AgentInfo:AgentType" locationName:"agentType" type:"string"`

	// Status of the collection process for an agent or connector.
	CollectionStatus *string `json:"discovery:AgentInfo:CollectionStatus" locationName:"collectionStatus" type:"string"`

	// The ID of the connector.
	ConnectorId *string `json:"discovery:AgentInfo:ConnectorId" locationName:"connectorId" type:"string"`

	// The health of the agent or connector.
	Health AgentStatus `json:"discovery:AgentInfo:Health" locationName:"health" type:"string" enum:"true"`

	// The name of the host where the agent or connector resides. The host can be
	// a server or virtual machine.
	HostName *string `json:"discovery:AgentInfo:HostName" locationName:"hostName" type:"string"`

	// Time since agent or connector health was reported.
	LastHealthPingTime *string `json:"discovery:AgentInfo:LastHealthPingTime" locationName:"lastHealthPingTime" type:"string"`

	// Agent's first registration timestamp in UTC.
	RegisteredTime *string `json:"discovery:AgentInfo:RegisteredTime" locationName:"registeredTime" type:"string"`

	// The agent or connector version.
	Version *string `json:"discovery:AgentInfo:Version" locationName:"version" type:"string"`
}

// String returns the string representation
func (s AgentInfo) String() string {
	return awsutil.Prettify(s)
}

// Network details about the host where the agent/connector resides.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/AgentNetworkInfo
type AgentNetworkInfo struct {
	_ struct{} `type:"structure"`

	// The IP address for the host where the agent/connector resides.
	IpAddress *string `json:"discovery:AgentNetworkInfo:IpAddress" locationName:"ipAddress" type:"string"`

	// The MAC address for the host where the agent/connector resides.
	MacAddress *string `json:"discovery:AgentNetworkInfo:MacAddress" locationName:"macAddress" type:"string"`
}

// String returns the string representation
func (s AgentNetworkInfo) String() string {
	return awsutil.Prettify(s)
}

// Error messages returned for each import task that you deleted as a response
// for this command.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/BatchDeleteImportDataError
type BatchDeleteImportDataError struct {
	_ struct{} `type:"structure"`

	// The type of error that occurred for a specific import task.
	ErrorCode BatchDeleteImportDataErrorCode `json:"discovery:BatchDeleteImportDataError:ErrorCode" locationName:"errorCode" type:"string" enum:"true"`

	// The description of the error that occurred for a specific import task.
	ErrorDescription *string `json:"discovery:BatchDeleteImportDataError:ErrorDescription" locationName:"errorDescription" type:"string"`

	// The unique import ID associated with the error that occurred.
	ImportTaskId *string `json:"discovery:BatchDeleteImportDataError:ImportTaskId" locationName:"importTaskId" type:"string"`
}

// String returns the string representation
func (s BatchDeleteImportDataError) String() string {
	return awsutil.Prettify(s)
}

// Tags for a configuration item. Tags are metadata that help you categorize
// IT assets.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/ConfigurationTag
type ConfigurationTag struct {
	_ struct{} `type:"structure"`

	// The configuration ID for the item to tag. You can specify a list of keys
	// and values.
	ConfigurationId *string `json:"discovery:ConfigurationTag:ConfigurationId" locationName:"configurationId" type:"string"`

	// A type of IT asset to tag.
	ConfigurationType ConfigurationItemType `json:"discovery:ConfigurationTag:ConfigurationType" locationName:"configurationType" type:"string" enum:"true"`

	// A type of tag on which to filter. For example, serverType.
	Key *string `json:"discovery:ConfigurationTag:Key" locationName:"key" type:"string"`

	// The time the configuration tag was created in Coordinated Universal Time
	// (UTC).
	TimeOfCreation *time.Time `json:"discovery:ConfigurationTag:TimeOfCreation" locationName:"timeOfCreation" type:"timestamp" timestampFormat:"unix"`

	// A value on which to filter. For example key = serverType and value = web
	// server.
	Value *string `json:"discovery:ConfigurationTag:Value" locationName:"value" type:"string"`
}

// String returns the string representation
func (s ConfigurationTag) String() string {
	return awsutil.Prettify(s)
}

// A list of continuous export descriptions.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/ContinuousExportDescription
type ContinuousExportDescription struct {
	_ struct{} `type:"structure"`

	// The type of data collector used to gather this data (currently only offered
	// for AGENT).
	DataSource DataSource `json:"discovery:ContinuousExportDescription:DataSource" locationName:"dataSource" type:"string" enum:"true"`

	// The unique ID assigned to this export.
	ExportId *string `json:"discovery:ContinuousExportDescription:ExportId" locationName:"exportId" type:"string"`

	// The name of the s3 bucket where the export data parquet files are stored.
	S3Bucket *string `json:"discovery:ContinuousExportDescription:S3Bucket" locationName:"s3Bucket" type:"string"`

	// An object which describes how the data is stored.
	//
	//    * databaseName - the name of the Glue database used to store the schema.
	SchemaStorageConfig map[string]string `json:"discovery:ContinuousExportDescription:SchemaStorageConfig" locationName:"schemaStorageConfig" type:"map"`

	// The timestamp representing when the continuous export was started.
	StartTime *time.Time `json:"discovery:ContinuousExportDescription:StartTime" locationName:"startTime" type:"timestamp" timestampFormat:"unix"`

	// Describes the status of the export. Can be one of the following values:
	//
	//    * START_IN_PROGRESS - setting up resources to start continuous export.
	//
	//    * START_FAILED - an error occurred setting up continuous export. To recover,
	//    call start-continuous-export again.
	//
	//    * ACTIVE - data is being exported to the customer bucket.
	//
	//    * ERROR - an error occurred during export. To fix the issue, call stop-continuous-export
	//    and start-continuous-export.
	//
	//    * STOP_IN_PROGRESS - stopping the export.
	//
	//    * STOP_FAILED - an error occurred stopping the export. To recover, call
	//    stop-continuous-export again.
	//
	//    * INACTIVE - the continuous export has been stopped. Data is no longer
	//    being exported to the customer bucket.
	Status ContinuousExportStatus `json:"discovery:ContinuousExportDescription:Status" locationName:"status" type:"string" enum:"true"`

	// Contains information about any errors that have occurred. This data type
	// can have the following values:
	//
	//    * ACCESS_DENIED - You don’t have permission to start Data Exploration
	//    in Amazon Athena. Contact your AWS administrator for help. For more information,
	//    see Setting Up AWS Application Discovery Service (http://docs.aws.amazon.com/application-discovery/latest/userguide/setting-up.html)
	//    in the Application Discovery Service User Guide.
	//
	//    * DELIVERY_STREAM_LIMIT_FAILURE - You reached the limit for Amazon Kinesis
	//    Data Firehose delivery streams. Reduce the number of streams or request
	//    a limit increase and try again. For more information, see Kinesis Data
	//    Streams Limits (http://docs.aws.amazon.com/streams/latest/dev/service-sizes-and-limits.html)
	//    in the Amazon Kinesis Data Streams Developer Guide.
	//
	//    * FIREHOSE_ROLE_MISSING - The Data Exploration feature is in an error
	//    state because your IAM User is missing the AWSApplicationDiscoveryServiceFirehose
	//    role. Turn on Data Exploration in Amazon Athena and try again. For more
	//    information, see Step 3: Provide Application Discovery Service Access
	//    to Non-Administrator Users by Attaching Policies (http://docs.aws.amazon.com/application-discovery/latest/userguide/setting-up.html#setting-up-user-policy)
	//    in the Application Discovery Service User Guide.
	//
	//    * FIREHOSE_STREAM_DOES_NOT_EXIST - The Data Exploration feature is in
	//    an error state because your IAM User is missing one or more of the Kinesis
	//    data delivery streams.
	//
	//    * INTERNAL_FAILURE - The Data Exploration feature is in an error state
	//    because of an internal failure. Try again later. If this problem persists,
	//    contact AWS Support.
	//
	//    * S3_BUCKET_LIMIT_FAILURE - You reached the limit for Amazon S3 buckets.
	//    Reduce the number of Amazon S3 buckets or request a limit increase and
	//    try again. For more information, see Bucket Restrictions and Limitations
	//    (http://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html)
	//    in the Amazon Simple Storage Service Developer Guide.
	//
	//    * S3_NOT_SIGNED_UP - Your account is not signed up for the Amazon S3 service.
	//    You must sign up before you can use Amazon S3. You can sign up at the
	//    following URL: https://aws.amazon.com/s3 (https://aws.amazon.com/s3).
	StatusDetail *string `json:"discovery:ContinuousExportDescription:StatusDetail" locationName:"statusDetail" min:"1" type:"string"`

	// The timestamp that represents when this continuous export was stopped.
	StopTime *time.Time `json:"discovery:ContinuousExportDescription:StopTime" locationName:"stopTime" type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s ContinuousExportDescription) String() string {
	return awsutil.Prettify(s)
}

// Inventory data for installed discovery agents.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/CustomerAgentInfo
type CustomerAgentInfo struct {
	_ struct{} `type:"structure"`

	// Number of active discovery agents.
	//
	// ActiveAgents is a required field
	ActiveAgents *int64 `json:"discovery:CustomerAgentInfo:ActiveAgents" locationName:"activeAgents" type:"integer" required:"true"`

	// Number of blacklisted discovery agents.
	//
	// BlackListedAgents is a required field
	BlackListedAgents *int64 `json:"discovery:CustomerAgentInfo:BlackListedAgents" locationName:"blackListedAgents" type:"integer" required:"true"`

	// Number of healthy discovery agents
	//
	// HealthyAgents is a required field
	HealthyAgents *int64 `json:"discovery:CustomerAgentInfo:HealthyAgents" locationName:"healthyAgents" type:"integer" required:"true"`

	// Number of discovery agents with status SHUTDOWN.
	//
	// ShutdownAgents is a required field
	ShutdownAgents *int64 `json:"discovery:CustomerAgentInfo:ShutdownAgents" locationName:"shutdownAgents" type:"integer" required:"true"`

	// Total number of discovery agents.
	//
	// TotalAgents is a required field
	TotalAgents *int64 `json:"discovery:CustomerAgentInfo:TotalAgents" locationName:"totalAgents" type:"integer" required:"true"`

	// Number of unhealthy discovery agents.
	//
	// UnhealthyAgents is a required field
	UnhealthyAgents *int64 `json:"discovery:CustomerAgentInfo:UnhealthyAgents" locationName:"unhealthyAgents" type:"integer" required:"true"`

	// Number of unknown discovery agents.
	//
	// UnknownAgents is a required field
	UnknownAgents *int64 `json:"discovery:CustomerAgentInfo:UnknownAgents" locationName:"unknownAgents" type:"integer" required:"true"`
}

// String returns the string representation
func (s CustomerAgentInfo) String() string {
	return awsutil.Prettify(s)
}

// Inventory data for installed discovery connectors.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/CustomerConnectorInfo
type CustomerConnectorInfo struct {
	_ struct{} `type:"structure"`

	// Number of active discovery connectors.
	//
	// ActiveConnectors is a required field
	ActiveConnectors *int64 `json:"discovery:CustomerConnectorInfo:ActiveConnectors" locationName:"activeConnectors" type:"integer" required:"true"`

	// Number of blacklisted discovery connectors.
	//
	// BlackListedConnectors is a required field
	BlackListedConnectors *int64 `json:"discovery:CustomerConnectorInfo:BlackListedConnectors" locationName:"blackListedConnectors" type:"integer" required:"true"`

	// Number of healthy discovery connectors.
	//
	// HealthyConnectors is a required field
	HealthyConnectors *int64 `json:"discovery:CustomerConnectorInfo:HealthyConnectors" locationName:"healthyConnectors" type:"integer" required:"true"`

	// Number of discovery connectors with status SHUTDOWN,
	//
	// ShutdownConnectors is a required field
	ShutdownConnectors *int64 `json:"discovery:CustomerConnectorInfo:ShutdownConnectors" locationName:"shutdownConnectors" type:"integer" required:"true"`

	// Total number of discovery connectors.
	//
	// TotalConnectors is a required field
	TotalConnectors *int64 `json:"discovery:CustomerConnectorInfo:TotalConnectors" locationName:"totalConnectors" type:"integer" required:"true"`

	// Number of unhealthy discovery connectors.
	//
	// UnhealthyConnectors is a required field
	UnhealthyConnectors *int64 `json:"discovery:CustomerConnectorInfo:UnhealthyConnectors" locationName:"unhealthyConnectors" type:"integer" required:"true"`

	// Number of unknown discovery connectors.
	//
	// UnknownConnectors is a required field
	UnknownConnectors *int64 `json:"discovery:CustomerConnectorInfo:UnknownConnectors" locationName:"unknownConnectors" type:"integer" required:"true"`
}

// String returns the string representation
func (s CustomerConnectorInfo) String() string {
	return awsutil.Prettify(s)
}

// Used to select which agent's data is to be exported. A single agent ID may
// be selected for export using the StartExportTask (http://docs.aws.amazon.com/application-discovery/latest/APIReference/API_StartExportTask.html)
// action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/ExportFilter
type ExportFilter struct {
	_ struct{} `type:"structure"`

	// Supported condition: EQUALS
	//
	// Condition is a required field
	Condition *string `json:"discovery:ExportFilter:Condition" locationName:"condition" type:"string" required:"true"`

	// A single ExportFilter name. Supported filters: agentId.
	//
	// Name is a required field
	Name *string `json:"discovery:ExportFilter:Name" locationName:"name" type:"string" required:"true"`

	// A single agentId for a Discovery Agent. An agentId can be found using the
	// DescribeAgents (http://docs.aws.amazon.com/application-discovery/latest/APIReference/API_DescribeExportTasks.html)
	// action. Typically an ADS agentId is in the form o-0123456789abcdef0.
	//
	// Values is a required field
	Values []string `json:"discovery:ExportFilter:Values" locationName:"values" type:"list" required:"true"`
}

// String returns the string representation
func (s ExportFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExportFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExportFilter"}

	if s.Condition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Condition"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Values == nil {
		invalidParams.Add(aws.NewErrParamRequired("Values"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information regarding the export status of discovered data. The value is
// an array of objects.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/ExportInfo
type ExportInfo struct {
	_ struct{} `type:"structure"`

	// A URL for an Amazon S3 bucket where you can review the exported data. The
	// URL is displayed only if the export succeeded.
	ConfigurationsDownloadUrl *string `json:"discovery:ExportInfo:ConfigurationsDownloadUrl" locationName:"configurationsDownloadUrl" type:"string"`

	// A unique identifier used to query an export.
	//
	// ExportId is a required field
	ExportId *string `json:"discovery:ExportInfo:ExportId" locationName:"exportId" type:"string" required:"true"`

	// The time that the data export was initiated.
	//
	// ExportRequestTime is a required field
	ExportRequestTime *time.Time `json:"discovery:ExportInfo:ExportRequestTime" locationName:"exportRequestTime" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The status of the data export job.
	//
	// ExportStatus is a required field
	ExportStatus ExportStatus `json:"discovery:ExportInfo:ExportStatus" locationName:"exportStatus" type:"string" required:"true" enum:"true"`

	// If true, the export of agent information exceeded the size limit for a single
	// export and the exported data is incomplete for the requested time range.
	// To address this, select a smaller time range for the export by using startDate
	// and endDate.
	IsTruncated *bool `json:"discovery:ExportInfo:IsTruncated" locationName:"isTruncated" type:"boolean"`

	// The endTime used in the StartExportTask request. If no endTime was requested,
	// this result does not appear in ExportInfo.
	RequestedEndTime *time.Time `json:"discovery:ExportInfo:RequestedEndTime" locationName:"requestedEndTime" type:"timestamp" timestampFormat:"unix"`

	// The value of startTime parameter in the StartExportTask request. If no startTime
	// was requested, this result does not appear in ExportInfo.
	RequestedStartTime *time.Time `json:"discovery:ExportInfo:RequestedStartTime" locationName:"requestedStartTime" type:"timestamp" timestampFormat:"unix"`

	// A status message provided for API callers.
	//
	// StatusMessage is a required field
	StatusMessage *string `json:"discovery:ExportInfo:StatusMessage" locationName:"statusMessage" type:"string" required:"true"`
}

// String returns the string representation
func (s ExportInfo) String() string {
	return awsutil.Prettify(s)
}

// A filter that can use conditional operators.
//
// For more information about filters, see Querying Discovered Configuration
// Items (http://docs.aws.amazon.com/application-discovery/latest/APIReference/discovery-api-queries.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/Filter
type Filter struct {
	_ struct{} `type:"structure"`

	// A conditional operator. The following operators are valid: EQUALS, NOT_EQUALS,
	// CONTAINS, NOT_CONTAINS. If you specify multiple filters, the system utilizes
	// all filters as though concatenated by AND. If you specify multiple values
	// for a particular filter, the system differentiates the values using OR. Calling
	// either DescribeConfigurations or ListConfigurations returns attributes of
	// matching configuration items.
	//
	// Condition is a required field
	Condition *string `json:"discovery:Filter:Condition" locationName:"condition" type:"string" required:"true"`

	// The name of the filter.
	//
	// Name is a required field
	Name *string `json:"discovery:Filter:Name" locationName:"name" type:"string" required:"true"`

	// A string value on which to filter. For example, if you choose the destinationServer.osVersion
	// filter name, you could specify Ubuntu for the value.
	//
	// Values is a required field
	Values []string `json:"discovery:Filter:Values" locationName:"values" type:"list" required:"true"`
}

// String returns the string representation
func (s Filter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Filter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Filter"}

	if s.Condition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Condition"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Values == nil {
		invalidParams.Add(aws.NewErrParamRequired("Values"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An array of information related to the import task request that includes
// status information, times, IDs, the Amazon S3 Object URL for the import file,
// and more.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/ImportTask
type ImportTask struct {
	_ struct{} `type:"structure"`

	// The total number of application records in the import file that failed to
	// be imported.
	ApplicationImportFailure *int64 `json:"discovery:ImportTask:ApplicationImportFailure" locationName:"applicationImportFailure" type:"integer"`

	// The total number of application records in the import file that were successfully
	// imported.
	ApplicationImportSuccess *int64 `json:"discovery:ImportTask:ApplicationImportSuccess" locationName:"applicationImportSuccess" type:"integer"`

	// A unique token used to prevent the same import request from occurring more
	// than once. If you didn't provide a token, a token was automatically generated
	// when the import task request was sent.
	ClientRequestToken *string `json:"discovery:ImportTask:ClientRequestToken" locationName:"clientRequestToken" min:"1" type:"string"`

	// A link to a compressed archive folder (in the ZIP format) that contains an
	// error log and a file of failed records. You can use these two files to quickly
	// identify records that failed, why they failed, and correct those records.
	// Afterward, you can upload the corrected file to your Amazon S3 bucket and
	// create another import task request.
	//
	// This field also includes authorization information so you can confirm the
	// authenticity of the compressed archive before you download it.
	//
	// If some records failed to be imported we recommend that you correct the records
	// in the failed entries file and then imports that failed entries file. This
	// prevents you from having to correct and update the larger original file and
	// attempt importing it again.
	ErrorsAndFailedEntriesZip *string `json:"discovery:ImportTask:ErrorsAndFailedEntriesZip" locationName:"errorsAndFailedEntriesZip" type:"string"`

	// The time that the import task request finished, presented in the Unix time
	// stamp format.
	ImportCompletionTime *time.Time `json:"discovery:ImportTask:ImportCompletionTime" locationName:"importCompletionTime" type:"timestamp" timestampFormat:"unix"`

	// The time that the import task request was deleted, presented in the Unix
	// time stamp format.
	ImportDeletedTime *time.Time `json:"discovery:ImportTask:ImportDeletedTime" locationName:"importDeletedTime" type:"timestamp" timestampFormat:"unix"`

	// The time that the import task request was made, presented in the Unix time
	// stamp format.
	ImportRequestTime *time.Time `json:"discovery:ImportTask:ImportRequestTime" locationName:"importRequestTime" type:"timestamp" timestampFormat:"unix"`

	// The unique ID for a specific import task. These IDs aren't globally unique,
	// but they are unique within an AWS account.
	ImportTaskId *string `json:"discovery:ImportTask:ImportTaskId" locationName:"importTaskId" type:"string"`

	// The URL for your import file that you've uploaded to Amazon S3.
	ImportUrl *string `json:"discovery:ImportTask:ImportUrl" locationName:"importUrl" min:"1" type:"string"`

	// A descriptive name for an import task. You can use this name to filter future
	// requests related to this import task, such as identifying applications and
	// servers that were included in this import task. We recommend that you use
	// a meaningful name for each import task.
	Name *string `json:"discovery:ImportTask:Name" locationName:"name" min:"1" type:"string"`

	// The total number of server records in the import file that failed to be imported.
	ServerImportFailure *int64 `json:"discovery:ImportTask:ServerImportFailure" locationName:"serverImportFailure" type:"integer"`

	// The total number of server records in the import file that were successfully
	// imported.
	ServerImportSuccess *int64 `json:"discovery:ImportTask:ServerImportSuccess" locationName:"serverImportSuccess" type:"integer"`

	// The status of the import task. An import can have the status of IMPORT_COMPLETE
	// and still have some records fail to import from the overall request. More
	// information can be found in the downloadable archive defined in the errorsAndFailedEntriesZip
	// field, or in the Migration Hub management console.
	Status ImportStatus `json:"discovery:ImportTask:Status" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ImportTask) String() string {
	return awsutil.Prettify(s)
}

// A name-values pair of elements you can use to filter the results when querying
// your import tasks. Currently, wildcards are not supported for filters.
//
// When filtering by import status, all other filter values are ignored.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/ImportTaskFilter
type ImportTaskFilter struct {
	_ struct{} `type:"structure"`

	// The name, status, or import task ID for a specific import task.
	Name ImportTaskFilterName `json:"discovery:ImportTaskFilter:Name" locationName:"name" type:"string" enum:"true"`

	// An array of strings that you can provide to match against a specific name,
	// status, or import task ID to filter the results for your import task queries.
	Values []string `json:"discovery:ImportTaskFilter:Values" locationName:"values" min:"1" type:"list"`
}

// String returns the string representation
func (s ImportTaskFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportTaskFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportTaskFilter"}
	if s.Values != nil && len(s.Values) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Values", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Details about neighboring servers.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/NeighborConnectionDetail
type NeighborConnectionDetail struct {
	_ struct{} `type:"structure"`

	// The number of open network connections with the neighboring server.
	//
	// ConnectionsCount is a required field
	ConnectionsCount *int64 `json:"discovery:NeighborConnectionDetail:ConnectionsCount" locationName:"connectionsCount" type:"long" required:"true"`

	// The destination network port for the connection.
	DestinationPort *int64 `json:"discovery:NeighborConnectionDetail:DestinationPort" locationName:"destinationPort" type:"integer"`

	// The ID of the server that accepted the network connection.
	//
	// DestinationServerId is a required field
	DestinationServerId *string `json:"discovery:NeighborConnectionDetail:DestinationServerId" locationName:"destinationServerId" type:"string" required:"true"`

	// The ID of the server that opened the network connection.
	//
	// SourceServerId is a required field
	SourceServerId *string `json:"discovery:NeighborConnectionDetail:SourceServerId" locationName:"sourceServerId" type:"string" required:"true"`

	// The network protocol used for the connection.
	TransportProtocol *string `json:"discovery:NeighborConnectionDetail:TransportProtocol" locationName:"transportProtocol" type:"string"`
}

// String returns the string representation
func (s NeighborConnectionDetail) String() string {
	return awsutil.Prettify(s)
}

// A field and direction for ordered output.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/OrderByElement
type OrderByElement struct {
	_ struct{} `type:"structure"`

	// The field on which to order.
	//
	// FieldName is a required field
	FieldName *string `json:"discovery:OrderByElement:FieldName" locationName:"fieldName" type:"string" required:"true"`

	// Ordering direction.
	SortOrder OrderString `json:"discovery:OrderByElement:SortOrder" locationName:"sortOrder" type:"string" enum:"true"`
}

// String returns the string representation
func (s OrderByElement) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *OrderByElement) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "OrderByElement"}

	if s.FieldName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FieldName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Metadata that help you categorize IT assets.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// The type of tag on which to filter.
	//
	// Key is a required field
	Key *string `json:"discovery:Tag:Key" locationName:"key" type:"string" required:"true"`

	// A value for a tag key on which to filter.
	//
	// Value is a required field
	Value *string `json:"discovery:Tag:Value" locationName:"value" type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The tag filter. Valid names are: tagKey, tagValue, configurationId.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/TagFilter
type TagFilter struct {
	_ struct{} `type:"structure"`

	// A name of the tag filter.
	//
	// Name is a required field
	Name *string `json:"discovery:TagFilter:Name" locationName:"name" type:"string" required:"true"`

	// Values for the tag filter.
	//
	// Values is a required field
	Values []string `json:"discovery:TagFilter:Values" locationName:"values" type:"list" required:"true"`
}

// String returns the string representation
func (s TagFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagFilter"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Values == nil {
		invalidParams.Add(aws.NewErrParamRequired("Values"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
