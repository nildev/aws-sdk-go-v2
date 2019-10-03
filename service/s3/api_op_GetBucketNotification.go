// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketNotificationConfigurationRequest
type GetBucketNotificationInput struct {
	_ struct{} `type:"structure"`

	// Name of the bucket to get the notification configuration for.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketNotificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketNotificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketNotificationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketNotificationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketNotificationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/NotificationConfigurationDeprecated
type GetBucketNotificationOutput struct {
	_ struct{} `type:"structure"`

	CloudFunctionConfiguration *CloudFunctionConfiguration `json:"s3:GetBucketNotificationOutput:CloudFunctionConfiguration" type:"structure"`

	QueueConfiguration *QueueConfigurationDeprecated `json:"s3:GetBucketNotificationOutput:QueueConfiguration" type:"structure"`

	TopicConfiguration *TopicConfigurationDeprecated `json:"s3:GetBucketNotificationOutput:TopicConfiguration" type:"structure"`
}

// String returns the string representation
func (s GetBucketNotificationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketNotificationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CloudFunctionConfiguration != nil {
		v := s.CloudFunctionConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CloudFunctionConfiguration", v, metadata)
	}
	if s.QueueConfiguration != nil {
		v := s.QueueConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "QueueConfiguration", v, metadata)
	}
	if s.TopicConfiguration != nil {
		v := s.TopicConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TopicConfiguration", v, metadata)
	}
	return nil
}

const opGetBucketNotification = "GetBucketNotification"

// GetBucketNotificationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// No longer used, see the GetBucketNotificationConfiguration operation.
//
//    // Example sending a request using GetBucketNotificationRequest.
//    req := client.GetBucketNotificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketNotification
func (c *Client) GetBucketNotificationRequest(input *GetBucketNotificationInput) GetBucketNotificationRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, GetBucketNotification, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opGetBucketNotification,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?notification",
	}

	if input == nil {
		input = &GetBucketNotificationInput{}
	}

	req := c.newRequest(op, input, &GetBucketNotificationOutput{})
	return GetBucketNotificationRequest{Request: req, Input: input, Copy: c.GetBucketNotificationRequest}
}

// GetBucketNotificationRequest is the request type for the
// GetBucketNotification API operation.
type GetBucketNotificationRequest struct {
	*aws.Request
	Input *GetBucketNotificationInput
	Copy  func(*GetBucketNotificationInput) GetBucketNotificationRequest
}

// Send marshals and sends the GetBucketNotification API request.
func (r GetBucketNotificationRequest) Send(ctx context.Context) (*GetBucketNotificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketNotificationResponse{
		GetBucketNotificationOutput: r.Request.Data.(*GetBucketNotificationOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketNotificationResponse is the response type for the
// GetBucketNotification API operation.
type GetBucketNotificationResponse struct {
	*GetBucketNotificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketNotification request.
func (r *GetBucketNotificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
