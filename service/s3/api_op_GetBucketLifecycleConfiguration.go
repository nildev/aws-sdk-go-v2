// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketLifecycleConfigurationRequest
type GetBucketLifecycleConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketLifecycleConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketLifecycleConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketLifecycleConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketLifecycleConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketLifecycleConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketLifecycleConfigurationOutput
type GetBucketLifecycleConfigurationOutput struct {
	_ struct{} `type:"structure"`

	Rules []LifecycleRule `json:"s3:GetBucketLifecycleConfigurationOutput:Rules" locationName:"Rule" type:"list" flattened:"true"`
}

// String returns the string representation
func (s GetBucketLifecycleConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketLifecycleConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Rules != nil {
		v := s.Rules

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "Rule", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetBucketLifecycleConfiguration = "GetBucketLifecycleConfiguration"

// GetBucketLifecycleConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the lifecycle configuration information set on the bucket.
//
//    // Example sending a request using GetBucketLifecycleConfigurationRequest.
//    req := client.GetBucketLifecycleConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketLifecycleConfiguration
func (c *Client) GetBucketLifecycleConfigurationRequest(input *GetBucketLifecycleConfigurationInput) GetBucketLifecycleConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetBucketLifecycleConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &GetBucketLifecycleConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetBucketLifecycleConfigurationOutput{})
	return GetBucketLifecycleConfigurationRequest{Request: req, Input: input, Copy: c.GetBucketLifecycleConfigurationRequest}
}

// GetBucketLifecycleConfigurationRequest is the request type for the
// GetBucketLifecycleConfiguration API operation.
type GetBucketLifecycleConfigurationRequest struct {
	*aws.Request
	Input *GetBucketLifecycleConfigurationInput
	Copy  func(*GetBucketLifecycleConfigurationInput) GetBucketLifecycleConfigurationRequest
}

// Send marshals and sends the GetBucketLifecycleConfiguration API request.
func (r GetBucketLifecycleConfigurationRequest) Send(ctx context.Context) (*GetBucketLifecycleConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketLifecycleConfigurationResponse{
		GetBucketLifecycleConfigurationOutput: r.Request.Data.(*GetBucketLifecycleConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketLifecycleConfigurationResponse is the response type for the
// GetBucketLifecycleConfiguration API operation.
type GetBucketLifecycleConfigurationResponse struct {
	*GetBucketLifecycleConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketLifecycleConfiguration request.
func (r *GetBucketLifecycleConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
