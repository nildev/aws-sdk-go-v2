// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request object for UpdateBrokerStorage.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/UpdateBrokerStorageRequest
type UpdateBrokerStorageInput struct {
	_ struct{} `type:"structure"`

	// ClusterArn is a required field
	ClusterArn *string `location:"uri" locationName:"clusterArn" type:"string" required:"true"`

	// The version of cluster to update from. A successful operation will then generate
	// a new version.
	//
	// CurrentVersion is a required field
	CurrentVersion *string `locationName:"currentVersion" type:"string" required:"true"`

	// Describes the target volume size and the ID of the broker to apply the update
	// to.
	//
	// TargetBrokerEBSVolumeInfo is a required field
	TargetBrokerEBSVolumeInfo []BrokerEBSVolumeInfo `locationName:"targetBrokerEBSVolumeInfo" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateBrokerStorageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBrokerStorageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBrokerStorageInput"}

	if s.ClusterArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterArn"))
	}

	if s.CurrentVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentVersion"))
	}

	if s.TargetBrokerEBSVolumeInfo == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetBrokerEBSVolumeInfo"))
	}
	if s.TargetBrokerEBSVolumeInfo != nil {
		for i, v := range s.TargetBrokerEBSVolumeInfo {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TargetBrokerEBSVolumeInfo", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateBrokerStorageInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.CurrentVersion != nil {
		v := *s.CurrentVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "currentVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TargetBrokerEBSVolumeInfo != nil {
		v := s.TargetBrokerEBSVolumeInfo

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targetBrokerEBSVolumeInfo", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Response body for UpdateBrokerStorage.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/UpdateBrokerStorageResponse
type UpdateBrokerStorageOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string `json:"kafka:UpdateBrokerStorageOutput:ClusterArn" locationName:"clusterArn" type:"string"`

	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string `json:"kafka:UpdateBrokerStorageOutput:ClusterOperationArn" locationName:"clusterOperationArn" type:"string"`
}

// String returns the string representation
func (s UpdateBrokerStorageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateBrokerStorageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClusterOperationArn != nil {
		v := *s.ClusterOperationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterOperationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateBrokerStorage = "UpdateBrokerStorage"

// UpdateBrokerStorageRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Updates the EBS storage associated with MSK brokers.
//
//    // Example sending a request using UpdateBrokerStorageRequest.
//    req := client.UpdateBrokerStorageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/UpdateBrokerStorage
func (c *Client) UpdateBrokerStorageRequest(input *UpdateBrokerStorageInput) UpdateBrokerStorageRequest {
	op := &aws.Operation{
		Name:       opUpdateBrokerStorage,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/clusters/{clusterArn}/nodes/storage",
	}

	if input == nil {
		input = &UpdateBrokerStorageInput{}
	}

	req := c.newRequest(op, input, &UpdateBrokerStorageOutput{})
	return UpdateBrokerStorageRequest{Request: req, Input: input, Copy: c.UpdateBrokerStorageRequest}
}

// UpdateBrokerStorageRequest is the request type for the
// UpdateBrokerStorage API operation.
type UpdateBrokerStorageRequest struct {
	*aws.Request
	Input *UpdateBrokerStorageInput
	Copy  func(*UpdateBrokerStorageInput) UpdateBrokerStorageRequest
}

// Send marshals and sends the UpdateBrokerStorage API request.
func (r UpdateBrokerStorageRequest) Send(ctx context.Context) (*UpdateBrokerStorageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateBrokerStorageResponse{
		UpdateBrokerStorageOutput: r.Request.Data.(*UpdateBrokerStorageOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateBrokerStorageResponse is the response type for the
// UpdateBrokerStorage API operation.
type UpdateBrokerStorageResponse struct {
	*UpdateBrokerStorageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateBrokerStorage request.
func (r *UpdateBrokerStorageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
