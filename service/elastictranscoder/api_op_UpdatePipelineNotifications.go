// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The UpdatePipelineNotificationsRequest structure.
type UpdatePipelineNotificationsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the pipeline for which you want to change notification
	// settings.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic
	// that you want to notify to report job status.
	//
	// To receive notifications, you must also subscribe to the new topic in the
	// Amazon SNS console.
	//
	//    * Progressing: The topic ARN for the Amazon Simple Notification Service
	//    (Amazon SNS) topic that you want to notify when Elastic Transcoder has
	//    started to process jobs that are added to this pipeline. This is the ARN
	//    that Amazon SNS returned when you created the topic.
	//
	//    * Complete: The topic ARN for the Amazon SNS topic that you want to notify
	//    when Elastic Transcoder has finished processing a job. This is the ARN
	//    that Amazon SNS returned when you created the topic.
	//
	//    * Warning: The topic ARN for the Amazon SNS topic that you want to notify
	//    when Elastic Transcoder encounters a warning condition. This is the ARN
	//    that Amazon SNS returned when you created the topic.
	//
	//    * Error: The topic ARN for the Amazon SNS topic that you want to notify
	//    when Elastic Transcoder encounters an error condition. This is the ARN
	//    that Amazon SNS returned when you created the topic.
	//
	// Notifications is a required field
	Notifications *Notifications `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdatePipelineNotificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePipelineNotificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePipelineNotificationsInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.Notifications == nil {
		invalidParams.Add(aws.NewErrParamRequired("Notifications"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePipelineNotificationsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Notifications != nil {
		v := s.Notifications

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Notifications", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The UpdatePipelineNotificationsResponse structure.
type UpdatePipelineNotificationsOutput struct {
	_ struct{} `type:"structure"`

	// A section of the response body that provides information about the pipeline
	// associated with this notification.
	Pipeline *Pipeline `json:"elastictranscoder:UpdatePipelineNotificationsOutput:Pipeline" type:"structure"`
}

// String returns the string representation
func (s UpdatePipelineNotificationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePipelineNotificationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Pipeline != nil {
		v := s.Pipeline

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Pipeline", v, metadata)
	}
	return nil
}

const opUpdatePipelineNotifications = "UpdatePipelineNotifications"

// UpdatePipelineNotificationsRequest returns a request value for making API operation for
// Amazon Elastic Transcoder.
//
// With the UpdatePipelineNotifications operation, you can update Amazon Simple
// Notification Service (Amazon SNS) notifications for a pipeline.
//
// When you update notifications for a pipeline, Elastic Transcoder returns
// the values that you specified in the request.
//
//    // Example sending a request using UpdatePipelineNotificationsRequest.
//    req := client.UpdatePipelineNotificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdatePipelineNotificationsRequest(input *UpdatePipelineNotificationsInput) UpdatePipelineNotificationsRequest {
	op := &aws.Operation{
		Name:       opUpdatePipelineNotifications,
		HTTPMethod: "POST",
		HTTPPath:   "/2012-09-25/pipelines/{Id}/notifications",
	}

	if input == nil {
		input = &UpdatePipelineNotificationsInput{}
	}

	req := c.newRequest(op, input, &UpdatePipelineNotificationsOutput{})
	return UpdatePipelineNotificationsRequest{Request: req, Input: input, Copy: c.UpdatePipelineNotificationsRequest}
}

// UpdatePipelineNotificationsRequest is the request type for the
// UpdatePipelineNotifications API operation.
type UpdatePipelineNotificationsRequest struct {
	*aws.Request
	Input *UpdatePipelineNotificationsInput
	Copy  func(*UpdatePipelineNotificationsInput) UpdatePipelineNotificationsRequest
}

// Send marshals and sends the UpdatePipelineNotifications API request.
func (r UpdatePipelineNotificationsRequest) Send(ctx context.Context) (*UpdatePipelineNotificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePipelineNotificationsResponse{
		UpdatePipelineNotificationsOutput: r.Request.Data.(*UpdatePipelineNotificationsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePipelineNotificationsResponse is the response type for the
// UpdatePipelineNotifications API operation.
type UpdatePipelineNotificationsResponse struct {
	*UpdatePipelineNotificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePipelineNotifications request.
func (r *UpdatePipelineNotificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
