// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dlm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dlm-2018-01-12/GetLifecyclePoliciesRequest
type GetLifecyclePoliciesInput struct {
	_ struct{} `type:"structure"`

	// The identifiers of the data lifecycle policies.
	PolicyIds []string `location:"querystring" locationName:"policyIds" type:"list"`

	// The resource type.
	ResourceTypes []ResourceTypeValues `location:"querystring" locationName:"resourceTypes" min:"1" type:"list"`

	// The activation state.
	State GettablePolicyStateValues `location:"querystring" locationName:"state" type:"string" enum:"true"`

	// The tags to add to objects created by the policy.
	//
	// Tags are strings in the format key=value.
	//
	// These user-defined tags are added in addition to the AWS-added lifecycle
	// tags.
	TagsToAdd []string `location:"querystring" locationName:"tagsToAdd" type:"list"`

	// The target tag for a policy.
	//
	// Tags are strings in the format key=value.
	TargetTags []string `location:"querystring" locationName:"targetTags" min:"1" type:"list"`
}

// String returns the string representation
func (s GetLifecyclePoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLifecyclePoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLifecyclePoliciesInput"}
	if s.ResourceTypes != nil && len(s.ResourceTypes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceTypes", 1))
	}
	if s.TargetTags != nil && len(s.TargetTags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetTags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLifecyclePoliciesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.PolicyIds != nil {
		v := s.PolicyIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "policyIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ResourceTypes != nil {
		v := s.ResourceTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "resourceTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.TagsToAdd != nil {
		v := s.TagsToAdd

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "tagsToAdd", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.TargetTags != nil {
		v := s.TargetTags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "targetTags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dlm-2018-01-12/GetLifecyclePoliciesResponse
type GetLifecyclePoliciesOutput struct {
	_ struct{} `type:"structure"`

	// Summary information about the lifecycle policies.
	Policies []LifecyclePolicySummary `json:"dlm:GetLifecyclePoliciesOutput:Policies" type:"list"`
}

// String returns the string representation
func (s GetLifecyclePoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLifecyclePoliciesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Policies != nil {
		v := s.Policies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Policies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetLifecyclePolicies = "GetLifecyclePolicies"

// GetLifecyclePoliciesRequest returns a request value for making API operation for
// Amazon Data Lifecycle Manager.
//
// Gets summary information about all or the specified data lifecycle policies.
//
// To get complete information about a policy, use GetLifecyclePolicy.
//
//    // Example sending a request using GetLifecyclePoliciesRequest.
//    req := client.GetLifecyclePoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dlm-2018-01-12/GetLifecyclePolicies
func (c *Client) GetLifecyclePoliciesRequest(input *GetLifecyclePoliciesInput) GetLifecyclePoliciesRequest {
	op := &aws.Operation{
		Name:       opGetLifecyclePolicies,
		HTTPMethod: "GET",
		HTTPPath:   "/policies",
	}

	if input == nil {
		input = &GetLifecyclePoliciesInput{}
	}

	req := c.newRequest(op, input, &GetLifecyclePoliciesOutput{})
	return GetLifecyclePoliciesRequest{Request: req, Input: input, Copy: c.GetLifecyclePoliciesRequest}
}

// GetLifecyclePoliciesRequest is the request type for the
// GetLifecyclePolicies API operation.
type GetLifecyclePoliciesRequest struct {
	*aws.Request
	Input *GetLifecyclePoliciesInput
	Copy  func(*GetLifecyclePoliciesInput) GetLifecyclePoliciesRequest
}

// Send marshals and sends the GetLifecyclePolicies API request.
func (r GetLifecyclePoliciesRequest) Send(ctx context.Context) (*GetLifecyclePoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLifecyclePoliciesResponse{
		GetLifecyclePoliciesOutput: r.Request.Data.(*GetLifecyclePoliciesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLifecyclePoliciesResponse is the response type for the
// GetLifecyclePolicies API operation.
type GetLifecyclePoliciesResponse struct {
	*GetLifecyclePoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLifecyclePolicies request.
func (r *GetLifecyclePoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
