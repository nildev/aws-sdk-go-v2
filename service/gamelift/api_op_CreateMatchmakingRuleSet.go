// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateMatchmakingRuleSetInput
type CreateMatchmakingRuleSetInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a matchmaking rule set. A matchmaking configuration
	// identifies the rule set it uses by this name value. (Note: The rule set name
	// is different from the optional "name" field in the rule set body.)
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// Collection of matchmaking rules, formatted as a JSON string. Comments are
	// not allowed in JSON, but most elements support a description field.
	//
	// RuleSetBody is a required field
	RuleSetBody *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateMatchmakingRuleSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMatchmakingRuleSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMatchmakingRuleSetInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.RuleSetBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleSetBody"))
	}
	if s.RuleSetBody != nil && len(*s.RuleSetBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleSetBody", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateMatchmakingRuleSetOutput
type CreateMatchmakingRuleSetOutput struct {
	_ struct{} `type:"structure"`

	// Object that describes the newly created matchmaking rule set.
	//
	// RuleSet is a required field
	RuleSet *MatchmakingRuleSet `json:"gamelift:CreateMatchmakingRuleSetOutput:RuleSet" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateMatchmakingRuleSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateMatchmakingRuleSet = "CreateMatchmakingRuleSet"

// CreateMatchmakingRuleSetRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Creates a new rule set for FlexMatch matchmaking. A rule set describes the
// type of match to create, such as the number and size of teams, and sets the
// parameters for acceptable player matches, such as minimum skill level or
// character type. A rule set is used by a MatchmakingConfiguration.
//
// To create a matchmaking rule set, provide unique rule set name and the rule
// set body in JSON format. Rule sets must be defined in the same region as
// the matchmaking configuration they are used with.
//
// Since matchmaking rule sets cannot be edited, it is a good idea to check
// the rule set syntax using ValidateMatchmakingRuleSet before creating a new
// rule set.
//
// Learn more
//
//    * Build a Rule Set (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-rulesets.html)
//
//    * Design a Matchmaker (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-configuration.html)
//
//    * Matchmaking with FlexMatch (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-intro.html)
//
// Related operations
//
//    * CreateMatchmakingConfiguration
//
//    * DescribeMatchmakingConfigurations
//
//    * UpdateMatchmakingConfiguration
//
//    * DeleteMatchmakingConfiguration
//
//    * CreateMatchmakingRuleSet
//
//    * DescribeMatchmakingRuleSets
//
//    * ValidateMatchmakingRuleSet
//
//    * DeleteMatchmakingRuleSet
//
//    // Example sending a request using CreateMatchmakingRuleSetRequest.
//    req := client.CreateMatchmakingRuleSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateMatchmakingRuleSet
func (c *Client) CreateMatchmakingRuleSetRequest(input *CreateMatchmakingRuleSetInput) CreateMatchmakingRuleSetRequest {
	op := &aws.Operation{
		Name:       opCreateMatchmakingRuleSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateMatchmakingRuleSetInput{}
	}

	req := c.newRequest(op, input, &CreateMatchmakingRuleSetOutput{})
	return CreateMatchmakingRuleSetRequest{Request: req, Input: input, Copy: c.CreateMatchmakingRuleSetRequest}
}

// CreateMatchmakingRuleSetRequest is the request type for the
// CreateMatchmakingRuleSet API operation.
type CreateMatchmakingRuleSetRequest struct {
	*aws.Request
	Input *CreateMatchmakingRuleSetInput
	Copy  func(*CreateMatchmakingRuleSetInput) CreateMatchmakingRuleSetRequest
}

// Send marshals and sends the CreateMatchmakingRuleSet API request.
func (r CreateMatchmakingRuleSetRequest) Send(ctx context.Context) (*CreateMatchmakingRuleSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMatchmakingRuleSetResponse{
		CreateMatchmakingRuleSetOutput: r.Request.Data.(*CreateMatchmakingRuleSetOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMatchmakingRuleSetResponse is the response type for the
// CreateMatchmakingRuleSet API operation.
type CreateMatchmakingRuleSetResponse struct {
	*CreateMatchmakingRuleSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMatchmakingRuleSet request.
func (r *CreateMatchmakingRuleSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
