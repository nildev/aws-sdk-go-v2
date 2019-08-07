// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DescribeBrokerRequest
type DescribeBrokerInput struct {
	_ struct{} `type:"structure"`

	// BrokerId is a required field
	BrokerId *string `location:"uri" locationName:"broker-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBrokerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBrokerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBrokerInput"}

	if s.BrokerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BrokerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeBrokerInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BrokerId != nil {
		v := *s.BrokerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "broker-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DescribeBrokerResponse
type DescribeBrokerOutput struct {
	_ struct{} `type:"structure"`

	AutoMinorVersionUpgrade *bool `json:"mq:DescribeBrokerOutput:AutoMinorVersionUpgrade" locationName:"autoMinorVersionUpgrade" type:"boolean"`

	BrokerArn *string `json:"mq:DescribeBrokerOutput:BrokerArn" locationName:"brokerArn" type:"string"`

	BrokerId *string `json:"mq:DescribeBrokerOutput:BrokerId" locationName:"brokerId" type:"string"`

	BrokerInstances []BrokerInstance `json:"mq:DescribeBrokerOutput:BrokerInstances" locationName:"brokerInstances" type:"list"`

	BrokerName *string `json:"mq:DescribeBrokerOutput:BrokerName" locationName:"brokerName" type:"string"`

	// The status of the broker.
	BrokerState BrokerState `json:"mq:DescribeBrokerOutput:BrokerState" locationName:"brokerState" type:"string" enum:"true"`

	// Broker configuration information
	Configurations *Configurations `json:"mq:DescribeBrokerOutput:Configurations" locationName:"configurations" type:"structure"`

	Created *time.Time `json:"mq:DescribeBrokerOutput:Created" locationName:"created" type:"timestamp" timestampFormat:"unix"`

	// The deployment mode of the broker.
	DeploymentMode DeploymentMode `json:"mq:DescribeBrokerOutput:DeploymentMode" locationName:"deploymentMode" type:"string" enum:"true"`

	// Encryption options for the broker.
	EncryptionOptions *EncryptionOptions `json:"mq:DescribeBrokerOutput:EncryptionOptions" locationName:"encryptionOptions" type:"structure"`

	// The type of broker engine. Note: Currently, Amazon MQ supports only ActiveMQ.
	EngineType EngineType `json:"mq:DescribeBrokerOutput:EngineType" locationName:"engineType" type:"string" enum:"true"`

	EngineVersion *string `json:"mq:DescribeBrokerOutput:EngineVersion" locationName:"engineVersion" type:"string"`

	HostInstanceType *string `json:"mq:DescribeBrokerOutput:HostInstanceType" locationName:"hostInstanceType" type:"string"`

	// The list of information about logs currently enabled and pending to be deployed
	// for the specified broker.
	Logs *LogsSummary `json:"mq:DescribeBrokerOutput:Logs" locationName:"logs" type:"structure"`

	// The scheduled time period relative to UTC during which Amazon MQ begins to
	// apply pending updates or patches to the broker.
	MaintenanceWindowStartTime *WeeklyStartTime `json:"mq:DescribeBrokerOutput:MaintenanceWindowStartTime" locationName:"maintenanceWindowStartTime" type:"structure"`

	PendingEngineVersion *string `json:"mq:DescribeBrokerOutput:PendingEngineVersion" locationName:"pendingEngineVersion" type:"string"`

	PubliclyAccessible *bool `json:"mq:DescribeBrokerOutput:PubliclyAccessible" locationName:"publiclyAccessible" type:"boolean"`

	SecurityGroups []string `json:"mq:DescribeBrokerOutput:SecurityGroups" locationName:"securityGroups" type:"list"`

	SubnetIds []string `json:"mq:DescribeBrokerOutput:SubnetIds" locationName:"subnetIds" type:"list"`

	Tags map[string]string `json:"mq:DescribeBrokerOutput:Tags" locationName:"tags" type:"map"`

	Users []UserSummary `json:"mq:DescribeBrokerOutput:Users" locationName:"users" type:"list"`
}

// String returns the string representation
func (s DescribeBrokerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeBrokerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AutoMinorVersionUpgrade != nil {
		v := *s.AutoMinorVersionUpgrade

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "autoMinorVersionUpgrade", protocol.BoolValue(v), metadata)
	}
	if s.BrokerArn != nil {
		v := *s.BrokerArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BrokerId != nil {
		v := *s.BrokerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BrokerInstances != nil {
		v := s.BrokerInstances

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "brokerInstances", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.BrokerName != nil {
		v := *s.BrokerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.BrokerState) > 0 {
		v := s.BrokerState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Configurations != nil {
		v := s.Configurations

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "configurations", v, metadata)
	}
	if s.Created != nil {
		v := *s.Created

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "created", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if len(s.DeploymentMode) > 0 {
		v := s.DeploymentMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deploymentMode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.EncryptionOptions != nil {
		v := s.EncryptionOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryptionOptions", v, metadata)
	}
	if len(s.EngineType) > 0 {
		v := s.EngineType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "engineType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.EngineVersion != nil {
		v := *s.EngineVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "engineVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HostInstanceType != nil {
		v := *s.HostInstanceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "hostInstanceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Logs != nil {
		v := s.Logs

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "logs", v, metadata)
	}
	if s.MaintenanceWindowStartTime != nil {
		v := s.MaintenanceWindowStartTime

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "maintenanceWindowStartTime", v, metadata)
	}
	if s.PendingEngineVersion != nil {
		v := *s.PendingEngineVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "pendingEngineVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PubliclyAccessible != nil {
		v := *s.PubliclyAccessible

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "publiclyAccessible", protocol.BoolValue(v), metadata)
	}
	if s.SecurityGroups != nil {
		v := s.SecurityGroups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "securityGroups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SubnetIds != nil {
		v := s.SubnetIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "subnetIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Users != nil {
		v := s.Users

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "users", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribeBroker = "DescribeBroker"

// DescribeBrokerRequest returns a request value for making API operation for
// AmazonMQ.
//
// Returns information about the specified broker.
//
//    // Example sending a request using DescribeBrokerRequest.
//    req := client.DescribeBrokerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DescribeBroker
func (c *Client) DescribeBrokerRequest(input *DescribeBrokerInput) DescribeBrokerRequest {
	op := &aws.Operation{
		Name:       opDescribeBroker,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/brokers/{broker-id}",
	}

	if input == nil {
		input = &DescribeBrokerInput{}
	}

	req := c.newRequest(op, input, &DescribeBrokerOutput{})
	return DescribeBrokerRequest{Request: req, Input: input, Copy: c.DescribeBrokerRequest}
}

// DescribeBrokerRequest is the request type for the
// DescribeBroker API operation.
type DescribeBrokerRequest struct {
	*aws.Request
	Input *DescribeBrokerInput
	Copy  func(*DescribeBrokerInput) DescribeBrokerRequest
}

// Send marshals and sends the DescribeBroker API request.
func (r DescribeBrokerRequest) Send(ctx context.Context) (*DescribeBrokerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBrokerResponse{
		DescribeBrokerOutput: r.Request.Data.(*DescribeBrokerOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeBrokerResponse is the response type for the
// DescribeBroker API operation.
type DescribeBrokerResponse struct {
	*DescribeBrokerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBroker request.
func (r *DescribeBrokerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
