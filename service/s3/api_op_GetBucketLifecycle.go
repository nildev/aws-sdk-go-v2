// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketLifecycleRequest
type GetBucketLifecycleInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketLifecycleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketLifecycleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketLifecycleInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketLifecycleInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketLifecycleInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketLifecycleOutput
type GetBucketLifecycleOutput struct {
	_ struct{} `type:"structure"`

	Rules []Rule `json:"s3:GetBucketLifecycleOutput:Rules" locationName:"Rule" type:"list" flattened:"true"`
}

// String returns the string representation
func (s GetBucketLifecycleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketLifecycleOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opGetBucketLifecycle = "GetBucketLifecycle"

// GetBucketLifecycleRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// No longer used, see the GetBucketLifecycleConfiguration operation.
//
//    // Example sending a request using GetBucketLifecycleRequest.
//    req := client.GetBucketLifecycleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketLifecycle
func (c *Client) GetBucketLifecycleRequest(input *GetBucketLifecycleInput) GetBucketLifecycleRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, GetBucketLifecycle, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opGetBucketLifecycle,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &GetBucketLifecycleInput{}
	}

	req := c.newRequest(op, input, &GetBucketLifecycleOutput{})
	return GetBucketLifecycleRequest{Request: req, Input: input, Copy: c.GetBucketLifecycleRequest}
}

// GetBucketLifecycleRequest is the request type for the
// GetBucketLifecycle API operation.
type GetBucketLifecycleRequest struct {
	*aws.Request
	Input *GetBucketLifecycleInput
	Copy  func(*GetBucketLifecycleInput) GetBucketLifecycleRequest
}

// Send marshals and sends the GetBucketLifecycle API request.
func (r GetBucketLifecycleRequest) Send(ctx context.Context) (*GetBucketLifecycleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketLifecycleResponse{
		GetBucketLifecycleOutput: r.Request.Data.(*GetBucketLifecycleOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketLifecycleResponse is the response type for the
// GetBucketLifecycle API operation.
type GetBucketLifecycleResponse struct {
	*GetBucketLifecycleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketLifecycle request.
func (r *GetBucketLifecycleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
