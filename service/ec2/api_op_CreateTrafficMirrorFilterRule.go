// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTrafficMirrorFilterRuleRequest
type CreateTrafficMirrorFilterRuleInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// The description of the Traffic Mirror rule.
	Description *string `type:"string"`

	// The destination CIDR block to assign to the Traffic Mirror rule.
	//
	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// The destination port range.
	DestinationPortRange *TrafficMirrorPortRangeRequest `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The protocol, for example UDP, to assign to the Traffic Mirror rule.
	//
	// For information about the protocol value, see Protocol Numbers (https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
	// on the Internet Assigned Numbers Authority (IANA) website.
	Protocol *int64 `type:"integer"`

	// The action to take (accept | reject) on the filtered traffic.
	//
	// RuleAction is a required field
	RuleAction TrafficMirrorRuleAction `type:"string" required:"true" enum:"true"`

	// The number of the Traffic Mirror rule. This number must be unique for each
	// Traffic Mirror rule in a given direction. The rules are processed in ascending
	// order by rule number.
	//
	// RuleNumber is a required field
	RuleNumber *int64 `type:"integer" required:"true"`

	// The source CIDR block to assign to the Traffic Mirror rule.
	//
	// SourceCidrBlock is a required field
	SourceCidrBlock *string `type:"string" required:"true"`

	// The source port range.
	SourcePortRange *TrafficMirrorPortRangeRequest `type:"structure"`

	// The type of traffic (ingress | egress).
	//
	// TrafficDirection is a required field
	TrafficDirection TrafficDirection `type:"string" required:"true" enum:"true"`

	// The ID of the filter that this rule is associated with.
	//
	// TrafficMirrorFilterId is a required field
	TrafficMirrorFilterId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTrafficMirrorFilterRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTrafficMirrorFilterRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTrafficMirrorFilterRuleInput"}

	if s.DestinationCidrBlock == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCidrBlock"))
	}
	if len(s.RuleAction) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RuleAction"))
	}

	if s.RuleNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleNumber"))
	}

	if s.SourceCidrBlock == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceCidrBlock"))
	}
	if len(s.TrafficDirection) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TrafficDirection"))
	}

	if s.TrafficMirrorFilterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficMirrorFilterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTrafficMirrorFilterRuleResult
type CreateTrafficMirrorFilterRuleOutput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `json:"ec2:CreateTrafficMirrorFilterRuleOutput:ClientToken" locationName:"clientToken" type:"string"`

	// The Traffic Mirror rule.
	TrafficMirrorFilterRule *TrafficMirrorFilterRule `json:"ec2:CreateTrafficMirrorFilterRuleOutput:TrafficMirrorFilterRule" locationName:"trafficMirrorFilterRule" type:"structure"`
}

// String returns the string representation
func (s CreateTrafficMirrorFilterRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTrafficMirrorFilterRule = "CreateTrafficMirrorFilterRule"

// CreateTrafficMirrorFilterRuleRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a Traffic Mirror rule.
//
// A Traffic Mirror rule defines the Traffic Mirror source traffic to mirror.
//
// You need the Traffic Mirror filter ID when you create the rule.
//
//    // Example sending a request using CreateTrafficMirrorFilterRuleRequest.
//    req := client.CreateTrafficMirrorFilterRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTrafficMirrorFilterRule
func (c *Client) CreateTrafficMirrorFilterRuleRequest(input *CreateTrafficMirrorFilterRuleInput) CreateTrafficMirrorFilterRuleRequest {
	op := &aws.Operation{
		Name:       opCreateTrafficMirrorFilterRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTrafficMirrorFilterRuleInput{}
	}

	req := c.newRequest(op, input, &CreateTrafficMirrorFilterRuleOutput{})
	return CreateTrafficMirrorFilterRuleRequest{Request: req, Input: input, Copy: c.CreateTrafficMirrorFilterRuleRequest}
}

// CreateTrafficMirrorFilterRuleRequest is the request type for the
// CreateTrafficMirrorFilterRule API operation.
type CreateTrafficMirrorFilterRuleRequest struct {
	*aws.Request
	Input *CreateTrafficMirrorFilterRuleInput
	Copy  func(*CreateTrafficMirrorFilterRuleInput) CreateTrafficMirrorFilterRuleRequest
}

// Send marshals and sends the CreateTrafficMirrorFilterRule API request.
func (r CreateTrafficMirrorFilterRuleRequest) Send(ctx context.Context) (*CreateTrafficMirrorFilterRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTrafficMirrorFilterRuleResponse{
		CreateTrafficMirrorFilterRuleOutput: r.Request.Data.(*CreateTrafficMirrorFilterRuleOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTrafficMirrorFilterRuleResponse is the response type for the
// CreateTrafficMirrorFilterRule API operation.
type CreateTrafficMirrorFilterRuleResponse struct {
	*CreateTrafficMirrorFilterRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTrafficMirrorFilterRule request.
func (r *CreateTrafficMirrorFilterRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
