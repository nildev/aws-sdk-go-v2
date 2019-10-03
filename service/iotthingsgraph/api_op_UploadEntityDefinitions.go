// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/UploadEntityDefinitionsRequest
type UploadEntityDefinitionsInput struct {
	_ struct{} `type:"structure"`

	// A Boolean that specifies whether to deprecate all entities in the latest
	// version before uploading the new DefinitionDocument. If set to true, the
	// upload will create a new namespace version.
	DeprecateExistingEntities *bool `locationName:"deprecateExistingEntities" type:"boolean"`

	// The DefinitionDocument that defines the updated entities.
	Document *DefinitionDocument `locationName:"document" type:"structure"`

	// A Boolean that specifies whether to synchronize with the latest version of
	// the public namespace. If set to true, the upload will create a new namespace
	// version.
	SyncWithPublicNamespace *bool `locationName:"syncWithPublicNamespace" type:"boolean"`
}

// String returns the string representation
func (s UploadEntityDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadEntityDefinitionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UploadEntityDefinitionsInput"}
	if s.Document != nil {
		if err := s.Document.Validate(); err != nil {
			invalidParams.AddNested("Document", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/UploadEntityDefinitionsResponse
type UploadEntityDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	// The ID that specifies the upload action. You can use this to track the status
	// of the upload.
	//
	// UploadId is a required field
	UploadId *string `json:"iotthingsgraph:UploadEntityDefinitionsOutput:UploadId" locationName:"uploadId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UploadEntityDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opUploadEntityDefinitions = "UploadEntityDefinitions"

// UploadEntityDefinitionsRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Asynchronously uploads one or more entity definitions to the user's namespace.
// The document parameter is required if syncWithPublicNamespace and deleteExistingEntites
// are false. If the syncWithPublicNamespace parameter is set to true, the user's
// namespace will synchronize with the latest version of the public namespace.
// If deprecateExistingEntities is set to true, all entities in the latest version
// will be deleted before the new DefinitionDocument is uploaded.
//
// When a user uploads entity definitions for the first time, the service creates
// a new namespace for the user. The new namespace tracks the public namespace.
// Currently users can have only one namespace. The namespace version increments
// whenever a user uploads entity definitions that are backwards-incompatible
// and whenever a user sets the syncWithPublicNamespace parameter or the deprecateExistingEntities
// parameter to true.
//
// The IDs for all of the entities should be in URN format. Each entity must
// be in the user's namespace. Users can't create entities in the public namespace,
// but entity definitions can refer to entities in the public namespace.
//
// Valid entities are Device, DeviceModel, Service, Capability, State, Action,
// Event, Property, Mapping, Enum.
//
//    // Example sending a request using UploadEntityDefinitionsRequest.
//    req := client.UploadEntityDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/UploadEntityDefinitions
func (c *Client) UploadEntityDefinitionsRequest(input *UploadEntityDefinitionsInput) UploadEntityDefinitionsRequest {
	op := &aws.Operation{
		Name:       opUploadEntityDefinitions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UploadEntityDefinitionsInput{}
	}

	req := c.newRequest(op, input, &UploadEntityDefinitionsOutput{})
	return UploadEntityDefinitionsRequest{Request: req, Input: input, Copy: c.UploadEntityDefinitionsRequest}
}

// UploadEntityDefinitionsRequest is the request type for the
// UploadEntityDefinitions API operation.
type UploadEntityDefinitionsRequest struct {
	*aws.Request
	Input *UploadEntityDefinitionsInput
	Copy  func(*UploadEntityDefinitionsInput) UploadEntityDefinitionsRequest
}

// Send marshals and sends the UploadEntityDefinitions API request.
func (r UploadEntityDefinitionsRequest) Send(ctx context.Context) (*UploadEntityDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UploadEntityDefinitionsResponse{
		UploadEntityDefinitionsOutput: r.Request.Data.(*UploadEntityDefinitionsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UploadEntityDefinitionsResponse is the response type for the
// UploadEntityDefinitions API operation.
type UploadEntityDefinitionsResponse struct {
	*UploadEntityDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UploadEntityDefinitions request.
func (r *UploadEntityDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
