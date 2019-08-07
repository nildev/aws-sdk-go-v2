// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// AddTagsToResourceInput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/AddTagsToResourceInput
type AddTagsToResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource you want to add tags to.
	//
	// ResourceARN is a required field
	ResourceARN *string `min:"50" type:"string" required:"true"`

	// The key-value pair that represents the tag you want to add to the resource.
	// The value can be an empty string.
	//
	// Valid characters for key and value are letters, spaces, and numbers representable
	// in UTF-8 format, and the following special characters: + - = . _ : / @. The
	// maximum length of a tag's key is 128 characters, and the maximum length for
	// a tag's value is 256.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsToResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsToResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsToResourceInput"}

	if s.ResourceARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceARN"))
	}
	if s.ResourceARN != nil && len(*s.ResourceARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceARN", 50))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// AddTagsToResourceOutput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/AddTagsToResourceOutput
type AddTagsToResourceOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource you want to add tags to.
	ResourceARN *string `json:"storagegateway:AddTagsToResourceOutput:ResourceARN" min:"50" type:"string"`
}

// String returns the string representation
func (s AddTagsToResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddTagsToResource = "AddTagsToResource"

// AddTagsToResourceRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Adds one or more tags to the specified resource. You use tags to add metadata
// to resources, which you can use to categorize these resources. For example,
// you can categorize resources by purpose, owner, environment, or team. Each
// tag consists of a key and a value, which you define. You can add tags to
// the following AWS Storage Gateway resources:
//
//    * Storage gateways of all types
//
//    * Storage volumes
//
//    * Virtual tapes
//
//    * NFS and SMB file shares
//
// You can create a maximum of 50 tags for each resource. Virtual tapes and
// storage volumes that are recovered to a new gateway maintain their tags.
//
//    // Example sending a request using AddTagsToResourceRequest.
//    req := client.AddTagsToResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/AddTagsToResource
func (c *Client) AddTagsToResourceRequest(input *AddTagsToResourceInput) AddTagsToResourceRequest {
	op := &aws.Operation{
		Name:       opAddTagsToResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddTagsToResourceInput{}
	}

	req := c.newRequest(op, input, &AddTagsToResourceOutput{})
	return AddTagsToResourceRequest{Request: req, Input: input, Copy: c.AddTagsToResourceRequest}
}

// AddTagsToResourceRequest is the request type for the
// AddTagsToResource API operation.
type AddTagsToResourceRequest struct {
	*aws.Request
	Input *AddTagsToResourceInput
	Copy  func(*AddTagsToResourceInput) AddTagsToResourceRequest
}

// Send marshals and sends the AddTagsToResource API request.
func (r AddTagsToResourceRequest) Send(ctx context.Context) (*AddTagsToResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddTagsToResourceResponse{
		AddTagsToResourceOutput: r.Request.Data.(*AddTagsToResourceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddTagsToResourceResponse is the response type for the
// AddTagsToResource API operation.
type AddTagsToResourceResponse struct {
	*AddTagsToResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddTagsToResource request.
func (r *AddTagsToResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
