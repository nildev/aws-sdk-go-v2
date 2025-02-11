// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Container for the parameters to the AddTags operation. Specify the tags that
// you want to attach to the Elasticsearch domain.
type AddTagsInput struct {
	_ struct{} `type:"structure"`

	// Specify the ARN for which you want to add the tags.
	//
	// ARN is a required field
	ARN *string `type:"string" required:"true"`

	// List of Tag that need to be added for the Elasticsearch domain.
	//
	// TagList is a required field
	TagList []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsInput"}

	if s.ARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ARN"))
	}

	if s.TagList == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagList"))
	}
	if s.TagList != nil {
		for i, v := range s.TagList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TagList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AddTagsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ARN != nil {
		v := *s.ARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TagList != nil {
		v := s.TagList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TagList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type AddTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AddTagsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAddTags = "AddTags"

// AddTagsRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Attaches tags to an existing Elasticsearch domain. Tags are a set of case-sensitive
// key value pairs. An Elasticsearch domain may have up to 10 tags. See Tagging
// Amazon Elasticsearch Service Domains for more information. (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-managedomains.html#es-managedomains-awsresorcetagging)
//
//    // Example sending a request using AddTagsRequest.
//    req := client.AddTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) AddTagsRequest(input *AddTagsInput) AddTagsRequest {
	op := &aws.Operation{
		Name:       opAddTags,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/tags",
	}

	if input == nil {
		input = &AddTagsInput{}
	}

	req := c.newRequest(op, input, &AddTagsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AddTagsRequest{Request: req, Input: input, Copy: c.AddTagsRequest}
}

// AddTagsRequest is the request type for the
// AddTags API operation.
type AddTagsRequest struct {
	*aws.Request
	Input *AddTagsInput
	Copy  func(*AddTagsInput) AddTagsRequest
}

// Send marshals and sends the AddTags API request.
func (r AddTagsRequest) Send(ctx context.Context) (*AddTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddTagsResponse{
		AddTagsOutput: r.Request.Data.(*AddTagsOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddTagsResponse is the response type for the
// AddTags API operation.
type AddTagsResponse struct {
	*AddTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddTags request.
func (r *AddTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
