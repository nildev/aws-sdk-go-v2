// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateGroupRequest
type UpdateGroupInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the group is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The description for the group that you want to update.
	Description *string `min:"1" type:"string"`

	// The name of the group that you want to update.
	//
	// GroupName is a required field
	GroupName *string `location:"uri" locationName:"GroupName" min:"1" type:"string" required:"true"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGroupInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.0"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupName != nil {
		v := *s.GroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateGroupResponse
type UpdateGroupOutput struct {
	_ struct{} `type:"structure"`

	// The name of the group.
	Group *Group `json:"quicksight:UpdateGroupOutput:Group" type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `json:"quicksight:UpdateGroupOutput:RequestId" type:"string"`

	// The http status of the request.
	Status *int64 `json:"quicksight:UpdateGroupOutput:Status" location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s UpdateGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Group != nil {
		v := s.Group

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Group", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateGroup = "UpdateGroup"

// UpdateGroupRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Changes a group description.
//
// The permissions resource is arn:aws:quicksight:us-east-1:<aws-account-id>:group/default/<group-name> .
//
// The response is a group object.
//
// CLI Sample:
//
// aws quicksight update-group --aws-account-id=111122223333 --namespace=default
// --group-name=Sales --description="Sales BI Dashboards"
//
//    // Example sending a request using UpdateGroupRequest.
//    req := client.UpdateGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateGroup
func (c *Client) UpdateGroupRequest(input *UpdateGroupInput) UpdateGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateGroup,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/groups/{GroupName}",
	}

	if input == nil {
		input = &UpdateGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateGroupOutput{})
	return UpdateGroupRequest{Request: req, Input: input, Copy: c.UpdateGroupRequest}
}

// UpdateGroupRequest is the request type for the
// UpdateGroup API operation.
type UpdateGroupRequest struct {
	*aws.Request
	Input *UpdateGroupInput
	Copy  func(*UpdateGroupInput) UpdateGroupRequest
}

// Send marshals and sends the UpdateGroup API request.
func (r UpdateGroupRequest) Send(ctx context.Context) (*UpdateGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateGroupResponse{
		UpdateGroupOutput: r.Request.Data.(*UpdateGroupOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateGroupResponse is the response type for the
// UpdateGroup API operation.
type UpdateGroupResponse struct {
	*UpdateGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateGroup request.
func (r *UpdateGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
