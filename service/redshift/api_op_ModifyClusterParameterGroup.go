// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifyClusterParameterGroupMessage
type ModifyClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the parameter group to be modified.
	//
	// ParameterGroupName is a required field
	ParameterGroupName *string `type:"string" required:"true"`

	// An array of parameters to be modified. A maximum of 20 parameters can be
	// modified in a single request.
	//
	// For each parameter to be modified, you must supply at least the parameter
	// name and parameter value; other name-value pairs of the parameter are optional.
	//
	// For the workload management (WLM) configuration, you must supply all the
	// name-value pairs in the wlm_json_configuration parameter.
	//
	// Parameters is a required field
	Parameters []Parameter `locationNameList:"Parameter" type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyClusterParameterGroupInput"}

	if s.ParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterGroupName"))
	}

	if s.Parameters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Parameters"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ClusterParameterGroupNameMessage
type ModifyClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// The name of the cluster parameter group.
	ParameterGroupName *string `json:"redshift:ModifyClusterParameterGroupOutput:ParameterGroupName" type:"string"`

	// The status of the parameter group. For example, if you made a change to a
	// parameter group name-value pair, then the change could be pending a reboot
	// of an associated cluster.
	ParameterGroupStatus *string `json:"redshift:ModifyClusterParameterGroupOutput:ParameterGroupStatus" type:"string"`
}

// String returns the string representation
func (s ModifyClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyClusterParameterGroup = "ModifyClusterParameterGroup"

// ModifyClusterParameterGroupRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Modifies the parameters of a parameter group.
//
// For more information about parameters and parameter groups, go to Amazon
// Redshift Parameter Groups (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using ModifyClusterParameterGroupRequest.
//    req := client.ModifyClusterParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifyClusterParameterGroup
func (c *Client) ModifyClusterParameterGroupRequest(input *ModifyClusterParameterGroupInput) ModifyClusterParameterGroupRequest {
	op := &aws.Operation{
		Name:       opModifyClusterParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyClusterParameterGroupInput{}
	}

	req := c.newRequest(op, input, &ModifyClusterParameterGroupOutput{})
	return ModifyClusterParameterGroupRequest{Request: req, Input: input, Copy: c.ModifyClusterParameterGroupRequest}
}

// ModifyClusterParameterGroupRequest is the request type for the
// ModifyClusterParameterGroup API operation.
type ModifyClusterParameterGroupRequest struct {
	*aws.Request
	Input *ModifyClusterParameterGroupInput
	Copy  func(*ModifyClusterParameterGroupInput) ModifyClusterParameterGroupRequest
}

// Send marshals and sends the ModifyClusterParameterGroup API request.
func (r ModifyClusterParameterGroupRequest) Send(ctx context.Context) (*ModifyClusterParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyClusterParameterGroupResponse{
		ModifyClusterParameterGroupOutput: r.Request.Data.(*ModifyClusterParameterGroupOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyClusterParameterGroupResponse is the response type for the
// ModifyClusterParameterGroup API operation.
type ModifyClusterParameterGroupResponse struct {
	*ModifyClusterParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyClusterParameterGroup request.
func (r *ModifyClusterParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
