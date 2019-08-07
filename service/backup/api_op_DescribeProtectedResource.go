// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DescribeProtectedResourceInput
type DescribeProtectedResourceInput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format
	// of the ARN depends on the resource type.
	//
	// ResourceArn is a required field
	ResourceArn *string `location:"uri" locationName:"resourceArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeProtectedResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProtectedResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProtectedResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeProtectedResourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DescribeProtectedResourceOutput
type DescribeProtectedResourceOutput struct {
	_ struct{} `type:"structure"`

	// The date and time that a resource was last backed up, in Unix format and
	// Coordinated Universal Time (UTC). The value of LastBackupTime is accurate
	// to milliseconds. For example, the value 1516925490.087 represents Friday,
	// January 26, 2018 12:11:30.087 AM.
	LastBackupTime *time.Time `json:"backup:DescribeProtectedResourceOutput:LastBackupTime" type:"timestamp" timestampFormat:"unix"`

	// An ARN that uniquely identifies a resource. The format of the ARN depends
	// on the resource type.
	ResourceArn *string `json:"backup:DescribeProtectedResourceOutput:ResourceArn" type:"string"`

	// The type of AWS resource saved as a recovery point; for example, an EBS volume
	// or an Amazon RDS database.
	ResourceType *string `json:"backup:DescribeProtectedResourceOutput:ResourceType" type:"string"`
}

// String returns the string representation
func (s DescribeProtectedResourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeProtectedResourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LastBackupTime != nil {
		v := *s.LastBackupTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastBackupTime", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceType != nil {
		v := *s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeProtectedResource = "DescribeProtectedResource"

// DescribeProtectedResourceRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns information about a saved resource, including the last time it was
// backed-up, its Amazon Resource Name (ARN), and the AWS service type of the
// saved resource.
//
//    // Example sending a request using DescribeProtectedResourceRequest.
//    req := client.DescribeProtectedResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DescribeProtectedResource
func (c *Client) DescribeProtectedResourceRequest(input *DescribeProtectedResourceInput) DescribeProtectedResourceRequest {
	op := &aws.Operation{
		Name:       opDescribeProtectedResource,
		HTTPMethod: "GET",
		HTTPPath:   "/resources/{resourceArn}",
	}

	if input == nil {
		input = &DescribeProtectedResourceInput{}
	}

	req := c.newRequest(op, input, &DescribeProtectedResourceOutput{})
	return DescribeProtectedResourceRequest{Request: req, Input: input, Copy: c.DescribeProtectedResourceRequest}
}

// DescribeProtectedResourceRequest is the request type for the
// DescribeProtectedResource API operation.
type DescribeProtectedResourceRequest struct {
	*aws.Request
	Input *DescribeProtectedResourceInput
	Copy  func(*DescribeProtectedResourceInput) DescribeProtectedResourceRequest
}

// Send marshals and sends the DescribeProtectedResource API request.
func (r DescribeProtectedResourceRequest) Send(ctx context.Context) (*DescribeProtectedResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProtectedResourceResponse{
		DescribeProtectedResourceOutput: r.Request.Data.(*DescribeProtectedResourceOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProtectedResourceResponse is the response type for the
// DescribeProtectedResource API operation.
type DescribeProtectedResourceResponse struct {
	*DescribeProtectedResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProtectedResource request.
func (r *DescribeProtectedResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
