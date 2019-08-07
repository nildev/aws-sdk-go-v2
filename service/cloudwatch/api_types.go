// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Represents the history of a specific alarm.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/AlarmHistoryItem
type AlarmHistoryItem struct {
	_ struct{} `type:"structure"`

	// The descriptive name for the alarm.
	AlarmName *string `json:"monitoring:AlarmHistoryItem:AlarmName" min:"1" type:"string"`

	// Data about the alarm, in JSON format.
	HistoryData *string `json:"monitoring:AlarmHistoryItem:HistoryData" min:"1" type:"string"`

	// The type of alarm history item.
	HistoryItemType HistoryItemType `json:"monitoring:AlarmHistoryItem:HistoryItemType" type:"string" enum:"true"`

	// A summary of the alarm history, in text format.
	HistorySummary *string `json:"monitoring:AlarmHistoryItem:HistorySummary" min:"1" type:"string"`

	// The time stamp for the alarm history item.
	Timestamp *time.Time `json:"monitoring:AlarmHistoryItem:Timestamp" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s AlarmHistoryItem) String() string {
	return awsutil.Prettify(s)
}

// An anomaly detection model associated with a particular CloudWatch metric
// athresnd statistic. You can use the model to display a band of expected normal
// values when the metric is graphed.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/AnomalyDetector
type AnomalyDetector struct {
	_ struct{} `type:"structure"`

	// The configuration specifies details about how the anomaly detection model
	// is to be trained, including time ranges to exclude from use for training
	// the model, and the time zone to use for the metric.
	Configuration *AnomalyDetectorConfiguration `json:"monitoring:AnomalyDetector:Configuration" type:"structure"`

	// The metric dimensions associated with the anomaly detection model.
	Dimensions []Dimension `json:"monitoring:AnomalyDetector:Dimensions" type:"list"`

	// The name of the metric associated with the anomaly detection model.
	MetricName *string `json:"monitoring:AnomalyDetector:MetricName" min:"1" type:"string"`

	// The namespace of the metric associated with the anomaly detection model.
	Namespace *string `json:"monitoring:AnomalyDetector:Namespace" min:"1" type:"string"`

	// The statistic associated with the anomaly detection model.
	Stat *string `json:"monitoring:AnomalyDetector:Stat" type:"string"`
}

// String returns the string representation
func (s AnomalyDetector) String() string {
	return awsutil.Prettify(s)
}

// The configuration specifies details about how the anomaly detection model
// is to be trained, including time ranges to exclude from use for training
// the model and the time zone to use for the metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/AnomalyDetectorConfiguration
type AnomalyDetectorConfiguration struct {
	_ struct{} `type:"structure"`

	// An array of time ranges to exclude from use when the anomaly detection model
	// is trained. Use this to make sure that events that could cause unusual values
	// for the metric, such as deployments, aren't used when CloudWatch creates
	// the model.
	ExcludedTimeRanges []Range `json:"monitoring:AnomalyDetectorConfiguration:ExcludedTimeRanges" type:"list"`

	// The time zone to use for the metric. This is useful to enable the model to
	// automatically account for daylight savings time changes if the metric is
	// sensitive to such time changes.
	//
	// To specify a time zone, use the name of the time zone as specified in the
	// standard tz database. For more information, see tz database (https://en.wikipedia.org/wiki/Tz_database).
	MetricTimezone *string `json:"monitoring:AnomalyDetectorConfiguration:MetricTimezone" type:"string"`
}

// String returns the string representation
func (s AnomalyDetectorConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AnomalyDetectorConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AnomalyDetectorConfiguration"}
	if s.ExcludedTimeRanges != nil {
		for i, v := range s.ExcludedTimeRanges {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ExcludedTimeRanges", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a specific dashboard.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DashboardEntry
type DashboardEntry struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string `json:"monitoring:DashboardEntry:DashboardArn" type:"string"`

	// The name of the dashboard.
	DashboardName *string `json:"monitoring:DashboardEntry:DashboardName" type:"string"`

	// The time stamp of when the dashboard was last modified, either by an API
	// call or through the console. This number is expressed as the number of milliseconds
	// since Jan 1, 1970 00:00:00 UTC.
	LastModified *time.Time `json:"monitoring:DashboardEntry:LastModified" type:"timestamp" timestampFormat:"iso8601"`

	// The size of the dashboard, in bytes.
	Size *int64 `json:"monitoring:DashboardEntry:Size" type:"long"`
}

// String returns the string representation
func (s DashboardEntry) String() string {
	return awsutil.Prettify(s)
}

// An error or warning for the operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DashboardValidationMessage
type DashboardValidationMessage struct {
	_ struct{} `type:"structure"`

	// The data path related to the message.
	DataPath *string `json:"monitoring:DashboardValidationMessage:DataPath" type:"string"`

	// A message describing the error or warning.
	Message *string `json:"monitoring:DashboardValidationMessage:Message" type:"string"`
}

// String returns the string representation
func (s DashboardValidationMessage) String() string {
	return awsutil.Prettify(s)
}

// Encapsulates the statistical data that CloudWatch computes from metric data.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/Datapoint
type Datapoint struct {
	_ struct{} `type:"structure"`

	// The average of the metric values that correspond to the data point.
	Average *float64 `json:"monitoring:Datapoint:Average" type:"double"`

	// The percentile statistic for the data point.
	ExtendedStatistics map[string]float64 `json:"monitoring:Datapoint:ExtendedStatistics" type:"map"`

	// The maximum metric value for the data point.
	Maximum *float64 `json:"monitoring:Datapoint:Maximum" type:"double"`

	// The minimum metric value for the data point.
	Minimum *float64 `json:"monitoring:Datapoint:Minimum" type:"double"`

	// The number of metric values that contributed to the aggregate value of this
	// data point.
	SampleCount *float64 `json:"monitoring:Datapoint:SampleCount" type:"double"`

	// The sum of the metric values for the data point.
	Sum *float64 `json:"monitoring:Datapoint:Sum" type:"double"`

	// The time stamp used for the data point.
	Timestamp *time.Time `json:"monitoring:Datapoint:Timestamp" type:"timestamp" timestampFormat:"iso8601"`

	// The standard unit for the data point.
	Unit StandardUnit `json:"monitoring:Datapoint:Unit" type:"string" enum:"true"`
}

// String returns the string representation
func (s Datapoint) String() string {
	return awsutil.Prettify(s)
}

// Expands the identity of a metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/Dimension
type Dimension struct {
	_ struct{} `type:"structure"`

	// The name of the dimension.
	//
	// Name is a required field
	Name *string `json:"monitoring:Dimension:Name" min:"1" type:"string" required:"true"`

	// The value representing the dimension measurement.
	//
	// Value is a required field
	Value *string `json:"monitoring:Dimension:Value" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s Dimension) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Dimension) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Dimension"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}
	if s.Value != nil && len(*s.Value) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Value", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents filters for a dimension.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DimensionFilter
type DimensionFilter struct {
	_ struct{} `type:"structure"`

	// The dimension name to be matched.
	//
	// Name is a required field
	Name *string `json:"monitoring:DimensionFilter:Name" min:"1" type:"string" required:"true"`

	// The value of the dimension to be matched.
	Value *string `json:"monitoring:DimensionFilter:Value" min:"1" type:"string"`
}

// String returns the string representation
func (s DimensionFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DimensionFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DimensionFilter"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Value != nil && len(*s.Value) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Value", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A message returned by the GetMetricDataAPI, including a code and a description.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/MessageData
type MessageData struct {
	_ struct{} `type:"structure"`

	// The error code or status code associated with the message.
	Code *string `json:"monitoring:MessageData:Code" type:"string"`

	// The message text.
	Value *string `json:"monitoring:MessageData:Value" type:"string"`
}

// String returns the string representation
func (s MessageData) String() string {
	return awsutil.Prettify(s)
}

// Represents a specific metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/Metric
type Metric struct {
	_ struct{} `type:"structure"`

	// The dimensions for the metric.
	Dimensions []Dimension `json:"monitoring:Metric:Dimensions" type:"list"`

	// The name of the metric. This is a required field.
	MetricName *string `json:"monitoring:Metric:MetricName" min:"1" type:"string"`

	// The namespace of the metric.
	Namespace *string `json:"monitoring:Metric:Namespace" min:"1" type:"string"`
}

// String returns the string representation
func (s Metric) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Metric) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Metric"}
	if s.MetricName != nil && len(*s.MetricName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricName", 1))
	}
	if s.Namespace != nil && len(*s.Namespace) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Namespace", 1))
	}
	if s.Dimensions != nil {
		for i, v := range s.Dimensions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Dimensions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents an alarm.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/MetricAlarm
type MetricAlarm struct {
	_ struct{} `type:"structure"`

	// Indicates whether actions should be executed during any changes to the alarm
	// state.
	ActionsEnabled *bool `json:"monitoring:MetricAlarm:ActionsEnabled" type:"boolean"`

	// The actions to execute when this alarm transitions to the ALARM state from
	// any other state. Each action is specified as an Amazon Resource Name (ARN).
	AlarmActions []string `json:"monitoring:MetricAlarm:AlarmActions" type:"list"`

	// The Amazon Resource Name (ARN) of the alarm.
	AlarmArn *string `json:"monitoring:MetricAlarm:AlarmArn" min:"1" type:"string"`

	// The time stamp of the last update to the alarm configuration.
	AlarmConfigurationUpdatedTimestamp *time.Time `json:"monitoring:MetricAlarm:AlarmConfigurationUpdatedTimestamp" type:"timestamp" timestampFormat:"iso8601"`

	// The description of the alarm.
	AlarmDescription *string `json:"monitoring:MetricAlarm:AlarmDescription" type:"string"`

	// The name of the alarm.
	AlarmName *string `json:"monitoring:MetricAlarm:AlarmName" min:"1" type:"string"`

	// The arithmetic operation to use when comparing the specified statistic and
	// threshold. The specified statistic value is used as the first operand.
	ComparisonOperator ComparisonOperator `json:"monitoring:MetricAlarm:ComparisonOperator" type:"string" enum:"true"`

	// The number of datapoints that must be breaching to trigger the alarm.
	DatapointsToAlarm *int64 `json:"monitoring:MetricAlarm:DatapointsToAlarm" min:"1" type:"integer"`

	// The dimensions for the metric associated with the alarm.
	Dimensions []Dimension `json:"monitoring:MetricAlarm:Dimensions" type:"list"`

	// Used only for alarms based on percentiles. If ignore, the alarm state does
	// not change during periods with too few data points to be statistically significant.
	// If evaluate or this parameter is not used, the alarm is always evaluated
	// and possibly changes state no matter how many data points are available.
	EvaluateLowSampleCountPercentile *string `json:"monitoring:MetricAlarm:EvaluateLowSampleCountPercentile" min:"1" type:"string"`

	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *int64 `json:"monitoring:MetricAlarm:EvaluationPeriods" min:"1" type:"integer"`

	// The percentile statistic for the metric associated with the alarm. Specify
	// a value between p0.0 and p100.
	ExtendedStatistic *string `json:"monitoring:MetricAlarm:ExtendedStatistic" type:"string"`

	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA
	// state from any other state. Each action is specified as an Amazon Resource
	// Name (ARN).
	InsufficientDataActions []string `json:"monitoring:MetricAlarm:InsufficientDataActions" type:"list"`

	// The name of the metric associated with the alarm, if this is an alarm based
	// on a single metric.
	MetricName *string `json:"monitoring:MetricAlarm:MetricName" min:"1" type:"string"`

	// An array of MetricDataQuery structures, used in an alarm based on a metric
	// math expression. Each structure either retrieves a metric or performs a math
	// expression. One item in the Metrics array is the math expression that the
	// alarm watches. This expression by designated by having ReturnValue set to
	// true.
	Metrics []MetricDataQuery `json:"monitoring:MetricAlarm:Metrics" type:"list"`

	// The namespace of the metric associated with the alarm.
	Namespace *string `json:"monitoring:MetricAlarm:Namespace" min:"1" type:"string"`

	// The actions to execute when this alarm transitions to the OK state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN).
	OKActions []string `json:"monitoring:MetricAlarm:OKActions" type:"list"`

	// The period, in seconds, over which the statistic is applied.
	Period *int64 `json:"monitoring:MetricAlarm:Period" min:"1" type:"integer"`

	// An explanation for the alarm state, in text format.
	StateReason *string `json:"monitoring:MetricAlarm:StateReason" type:"string"`

	// An explanation for the alarm state, in JSON format.
	StateReasonData *string `json:"monitoring:MetricAlarm:StateReasonData" type:"string"`

	// The time stamp of the last update to the alarm state.
	StateUpdatedTimestamp *time.Time `json:"monitoring:MetricAlarm:StateUpdatedTimestamp" type:"timestamp" timestampFormat:"iso8601"`

	// The state value for the alarm.
	StateValue StateValue `json:"monitoring:MetricAlarm:StateValue" type:"string" enum:"true"`

	// The statistic for the metric associated with the alarm, other than percentile.
	// For percentile statistics, use ExtendedStatistic.
	Statistic Statistic `json:"monitoring:MetricAlarm:Statistic" type:"string" enum:"true"`

	// The value to compare with the specified statistic.
	Threshold *float64 `json:"monitoring:MetricAlarm:Threshold" type:"double"`

	// In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND
	// function used as the threshold for the alarm.
	ThresholdMetricId *string `json:"monitoring:MetricAlarm:ThresholdMetricId" min:"1" type:"string"`

	// Sets how this alarm is to handle missing data points. If this parameter is
	// omitted, the default behavior of missing is used.
	TreatMissingData *string `json:"monitoring:MetricAlarm:TreatMissingData" min:"1" type:"string"`

	// The unit of the metric associated with the alarm.
	Unit StandardUnit `json:"monitoring:MetricAlarm:Unit" type:"string" enum:"true"`
}

// String returns the string representation
func (s MetricAlarm) String() string {
	return awsutil.Prettify(s)
}

// This structure is used in both GetMetricData and PutMetricAlarm. The supported
// use of this structure is different for those two operations.
//
// When used in GetMetricData, it indicates the metric data to return, and whether
// this call is just retrieving a batch set of data for one metric, or is performing
// a math expression on metric data. A single GetMetricData call can include
// up to 100 MetricDataQuery structures.
//
// When used in PutMetricAlarm, it enables you to create an alarm based on a
// metric math expression. Each MetricDataQuery in the array specifies either
// a metric to retrieve, or a math expression to be performed on retrieved metrics.
// A single PutMetricAlarm call can include up to 20 MetricDataQuery structures
// in the array. The 20 structures can include as many as 10 structures that
// contain a MetricStat parameter to retrieve a metric, and as many as 10 structures
// that contain the Expression parameter to perform a math expression. Of those
// Expression structures, one must have True as the value for ReturnData. The
// result of this expression is the value the alarm watches.
//
// Any expression used in a PutMetricAlarm operation must return a single time
// series. For more information, see Metric Math Syntax and Functions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax)
// in the Amazon CloudWatch User Guide.
//
// Some of the parameters of this structure also have different uses whether
// you are using this structure in a GetMetricData operation or a PutMetricAlarm
// operation. These differences are explained in the following parameter list.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/MetricDataQuery
type MetricDataQuery struct {
	_ struct{} `type:"structure"`

	// The math expression to be performed on the returned data, if this object
	// is performing a math expression. This expression can use the Id of the other
	// metrics to refer to those metrics, and can also use the Id of other expressions
	// to use the result of those expressions. For more information about metric
	// math expressions, see Metric Math Syntax and Functions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax)
	// in the Amazon CloudWatch User Guide.
	//
	// Within each MetricDataQuery object, you must specify either Expression or
	// MetricStat but not both.
	Expression *string `json:"monitoring:MetricDataQuery:Expression" min:"1" type:"string"`

	// A short name used to tie this object to the results in the response. This
	// name must be unique within a single call to GetMetricData. If you are performing
	// math expressions on this set of data, this name represents that data and
	// can serve as a variable in the mathematical expression. The valid characters
	// are letters, numbers, and underscore. The first character must be a lowercase
	// letter.
	//
	// Id is a required field
	Id *string `json:"monitoring:MetricDataQuery:Id" min:"1" type:"string" required:"true"`

	// A human-readable label for this metric or expression. This is especially
	// useful if this is an expression, so that you know what the value represents.
	// If the metric or expression is shown in a CloudWatch dashboard widget, the
	// label is shown. If Label is omitted, CloudWatch generates a default.
	Label *string `json:"monitoring:MetricDataQuery:Label" type:"string"`

	// The metric to be returned, along with statistics, period, and units. Use
	// this parameter only if this object is retrieving a metric and not performing
	// a math expression on returned data.
	//
	// Within one MetricDataQuery object, you must specify either Expression or
	// MetricStat but not both.
	MetricStat *MetricStat `json:"monitoring:MetricDataQuery:MetricStat" type:"structure"`

	// When used in GetMetricData, this option indicates whether to return the timestamps
	// and raw data values of this metric. If you are performing this call just
	// to do math expressions and do not also need the raw data returned, you can
	// specify False. If you omit this, the default of True is used.
	//
	// When used in PutMetricAlarm, specify True for the one expression result to
	// use as the alarm. For all other metrics and expressions in the same PutMetricAlarm
	// operation, specify ReturnData as False.
	ReturnData *bool `json:"monitoring:MetricDataQuery:ReturnData" type:"boolean"`
}

// String returns the string representation
func (s MetricDataQuery) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricDataQuery) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricDataQuery"}
	if s.Expression != nil && len(*s.Expression) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Expression", 1))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if s.MetricStat != nil {
		if err := s.MetricStat.Validate(); err != nil {
			invalidParams.AddNested("MetricStat", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A GetMetricData call returns an array of MetricDataResult structures. Each
// of these structures includes the data points for that metric, along with
// the timestamps of those data points and other identifying information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/MetricDataResult
type MetricDataResult struct {
	_ struct{} `type:"structure"`

	// The short name you specified to represent this metric.
	Id *string `json:"monitoring:MetricDataResult:Id" min:"1" type:"string"`

	// The human-readable label associated with the data.
	Label *string `json:"monitoring:MetricDataResult:Label" type:"string"`

	// A list of messages with additional information about the data returned.
	Messages []MessageData `json:"monitoring:MetricDataResult:Messages" type:"list"`

	// The status of the returned data. Complete indicates that all data points
	// in the requested time range were returned. PartialData means that an incomplete
	// set of data points were returned. You can use the NextToken value that was
	// returned and repeat your request to get more data points. NextToken is not
	// returned if you are performing a math expression. InternalError indicates
	// that an error occurred. Retry your request using NextToken, if present.
	StatusCode StatusCode `json:"monitoring:MetricDataResult:StatusCode" type:"string" enum:"true"`

	// The timestamps for the data points, formatted in Unix timestamp format. The
	// number of timestamps always matches the number of values and the value for
	// Timestamps[x] is Values[x].
	Timestamps []time.Time `json:"monitoring:MetricDataResult:Timestamps" type:"list"`

	// The data points for the metric corresponding to Timestamps. The number of
	// values always matches the number of timestamps and the timestamp for Values[x]
	// is Timestamps[x].
	Values []float64 `json:"monitoring:MetricDataResult:Values" type:"list"`
}

// String returns the string representation
func (s MetricDataResult) String() string {
	return awsutil.Prettify(s)
}

// Encapsulates the information sent to either create a metric or add new values
// to be aggregated into an existing metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/MetricDatum
type MetricDatum struct {
	_ struct{} `type:"structure"`

	// Array of numbers that is used along with the Values array. Each number in
	// the Count array is the number of times the corresponding value in the Values
	// array occurred during the period.
	//
	// If you omit the Counts array, the default of 1 is used as the value for each
	// count. If you include a Counts array, it must include the same amount of
	// values as the Values array.
	Counts []float64 `json:"monitoring:MetricDatum:Counts" type:"list"`

	// The dimensions associated with the metric.
	Dimensions []Dimension `json:"monitoring:MetricDatum:Dimensions" type:"list"`

	// The name of the metric.
	//
	// MetricName is a required field
	MetricName *string `json:"monitoring:MetricDatum:MetricName" min:"1" type:"string" required:"true"`

	// The statistical values for the metric.
	StatisticValues *StatisticSet `json:"monitoring:MetricDatum:StatisticValues" type:"structure"`

	// Valid values are 1 and 60. Setting this to 1 specifies this metric as a high-resolution
	// metric, so that CloudWatch stores the metric with sub-minute resolution down
	// to one second. Setting this to 60 specifies this metric as a regular-resolution
	// metric, which CloudWatch stores at 1-minute resolution. Currently, high resolution
	// is available only for custom metrics. For more information about high-resolution
	// metrics, see High-Resolution Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html#high-resolution-metrics)
	// in the Amazon CloudWatch User Guide.
	//
	// This field is optional, if you do not specify it the default of 60 is used.
	StorageResolution *int64 `json:"monitoring:MetricDatum:StorageResolution" min:"1" type:"integer"`

	// The time the metric data was received, expressed as the number of milliseconds
	// since Jan 1, 1970 00:00:00 UTC.
	Timestamp *time.Time `json:"monitoring:MetricDatum:Timestamp" type:"timestamp" timestampFormat:"iso8601"`

	// When you are using a Put operation, this defines what unit you want to use
	// when storing the metric. In a Get operation, this displays the unit that
	// is used for the metric.
	Unit StandardUnit `json:"monitoring:MetricDatum:Unit" type:"string" enum:"true"`

	// The value for the metric.
	//
	// Although the parameter accepts numbers of type Double, CloudWatch rejects
	// values that are either too small or too large. Values must be in the range
	// of 8.515920e-109 to 1.174271e+108 (Base 10) or 2e-360 to 2e360 (Base 2).
	// In addition, special values (for example, NaN, +Infinity, -Infinity) are
	// not supported.
	Value *float64 `json:"monitoring:MetricDatum:Value" type:"double"`

	// Array of numbers representing the values for the metric during the period.
	// Each unique value is listed just once in this array, and the corresponding
	// number in the Counts array specifies the number of times that value occurred
	// during the period. You can include up to 150 unique values in each PutMetricData
	// action that specifies a Values array.
	//
	// Although the Values array accepts numbers of type Double, CloudWatch rejects
	// values that are either too small or too large. Values must be in the range
	// of 8.515920e-109 to 1.174271e+108 (Base 10) or 2e-360 to 2e360 (Base 2).
	// In addition, special values (for example, NaN, +Infinity, -Infinity) are
	// not supported.
	Values []float64 `json:"monitoring:MetricDatum:Values" type:"list"`
}

// String returns the string representation
func (s MetricDatum) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricDatum) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricDatum"}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}
	if s.MetricName != nil && len(*s.MetricName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricName", 1))
	}
	if s.StorageResolution != nil && *s.StorageResolution < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("StorageResolution", 1))
	}
	if s.Dimensions != nil {
		for i, v := range s.Dimensions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Dimensions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.StatisticValues != nil {
		if err := s.StatisticValues.Validate(); err != nil {
			invalidParams.AddNested("StatisticValues", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This structure defines the metric to be returned, along with the statistics,
// period, and units.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/MetricStat
type MetricStat struct {
	_ struct{} `type:"structure"`

	// The metric to return, including the metric name, namespace, and dimensions.
	//
	// Metric is a required field
	Metric *Metric `json:"monitoring:MetricStat:Metric" type:"structure" required:"true"`

	// The period, in seconds, to use when retrieving the metric.
	//
	// Period is a required field
	Period *int64 `json:"monitoring:MetricStat:Period" min:"1" type:"integer" required:"true"`

	// The statistic to return. It can include any CloudWatch statistic or extended
	// statistic.
	//
	// Stat is a required field
	Stat *string `json:"monitoring:MetricStat:Stat" type:"string" required:"true"`

	// When you are using a Put operation, this defines what unit you want to use
	// when storing the metric. In a Get operation, this displays the unit that
	// is used for the metric.
	Unit StandardUnit `json:"monitoring:MetricStat:Unit" type:"string" enum:"true"`
}

// String returns the string representation
func (s MetricStat) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricStat) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricStat"}

	if s.Metric == nil {
		invalidParams.Add(aws.NewErrParamRequired("Metric"))
	}

	if s.Period == nil {
		invalidParams.Add(aws.NewErrParamRequired("Period"))
	}
	if s.Period != nil && *s.Period < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Period", 1))
	}

	if s.Stat == nil {
		invalidParams.Add(aws.NewErrParamRequired("Stat"))
	}
	if s.Metric != nil {
		if err := s.Metric.Validate(); err != nil {
			invalidParams.AddNested("Metric", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies one range of days or times to exclude from use for training an
// anomaly detection model.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/Range
type Range struct {
	_ struct{} `type:"structure"`

	// The end time of the range to exclude. The format is yyyy-MM-dd'T'HH:mm:ss.
	// For example, 2019-07-01T23:59:59.
	//
	// EndTime is a required field
	EndTime *time.Time `json:"monitoring:Range:EndTime" type:"timestamp" timestampFormat:"iso8601" required:"true"`

	// The start time of the range to exclude. The format is yyyy-MM-dd'T'HH:mm:ss.
	// For example, 2019-07-01T23:59:59.
	//
	// StartTime is a required field
	StartTime *time.Time `json:"monitoring:Range:StartTime" type:"timestamp" timestampFormat:"iso8601" required:"true"`
}

// String returns the string representation
func (s Range) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Range) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Range"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a set of statistics that describes a specific metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/StatisticSet
type StatisticSet struct {
	_ struct{} `type:"structure"`

	// The maximum value of the sample set.
	//
	// Maximum is a required field
	Maximum *float64 `json:"monitoring:StatisticSet:Maximum" type:"double" required:"true"`

	// The minimum value of the sample set.
	//
	// Minimum is a required field
	Minimum *float64 `json:"monitoring:StatisticSet:Minimum" type:"double" required:"true"`

	// The number of samples used for the statistic set.
	//
	// SampleCount is a required field
	SampleCount *float64 `json:"monitoring:StatisticSet:SampleCount" type:"double" required:"true"`

	// The sum of values for the sample set.
	//
	// Sum is a required field
	Sum *float64 `json:"monitoring:StatisticSet:Sum" type:"double" required:"true"`
}

// String returns the string representation
func (s StatisticSet) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StatisticSet) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StatisticSet"}

	if s.Maximum == nil {
		invalidParams.Add(aws.NewErrParamRequired("Maximum"))
	}

	if s.Minimum == nil {
		invalidParams.Add(aws.NewErrParamRequired("Minimum"))
	}

	if s.SampleCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("SampleCount"))
	}

	if s.Sum == nil {
		invalidParams.Add(aws.NewErrParamRequired("Sum"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A key-value pair associated with a CloudWatch resource.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// A string that you can use to assign a value. The combination of tag keys
	// and values can help you organize and categorize your resources.
	//
	// Key is a required field
	Key *string `json:"monitoring:Tag:Key" min:"1" type:"string" required:"true"`

	// The value for the specified tag key.
	//
	// Value is a required field
	Value *string `json:"monitoring:Tag:Value" type:"string" required:"true"`
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
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
