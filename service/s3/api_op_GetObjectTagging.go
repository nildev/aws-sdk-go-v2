// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectTaggingRequest
type GetObjectTaggingInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s GetObjectTaggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectTaggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectTaggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetObjectTaggingInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectTaggingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectTaggingOutput
type GetObjectTaggingOutput struct {
	_ struct{} `type:"structure"`

	// TagSet is a required field
	TagSet []Tag `json:"s3:GetObjectTaggingOutput:TagSet" locationNameList:"Tag" type:"list" required:"true"`

	VersionId *string `json:"s3:GetObjectTaggingOutput:VersionId" location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s GetObjectTaggingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectTaggingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TagSet != nil {
		v := s.TagSet

		metadata := protocol.Metadata{ListLocationName: "Tag"}
		ls0 := e.List(protocol.BodyTarget, "TagSet", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-version-id", protocol.StringValue(v), metadata)
	}
	return nil
}

const opGetObjectTagging = "GetObjectTagging"

// GetObjectTaggingRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the tag-set of an object.
//
//    // Example sending a request using GetObjectTaggingRequest.
//    req := client.GetObjectTaggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectTagging
func (c *Client) GetObjectTaggingRequest(input *GetObjectTaggingInput) GetObjectTaggingRequest {
	op := &aws.Operation{
		Name:       opGetObjectTagging,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}?tagging",
	}

	if input == nil {
		input = &GetObjectTaggingInput{}
	}

	req := c.newRequest(op, input, &GetObjectTaggingOutput{})
	return GetObjectTaggingRequest{Request: req, Input: input, Copy: c.GetObjectTaggingRequest}
}

// GetObjectTaggingRequest is the request type for the
// GetObjectTagging API operation.
type GetObjectTaggingRequest struct {
	*aws.Request
	Input *GetObjectTaggingInput
	Copy  func(*GetObjectTaggingInput) GetObjectTaggingRequest
}

// Send marshals and sends the GetObjectTagging API request.
func (r GetObjectTaggingRequest) Send(ctx context.Context) (*GetObjectTaggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectTaggingResponse{
		GetObjectTaggingOutput: r.Request.Data.(*GetObjectTaggingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectTaggingResponse is the response type for the
// GetObjectTagging API operation.
type GetObjectTaggingResponse struct {
	*GetObjectTaggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObjectTagging request.
func (r *GetObjectTaggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
