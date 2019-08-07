// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/DescribeUpdateRequest
type DescribeUpdateInput struct {
	_ struct{} `type:"structure"`

	// The name of the Amazon EKS cluster to update.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// The ID of the update to describe.
	//
	// UpdateId is a required field
	UpdateId *string `location:"uri" locationName:"updateId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUpdateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUpdateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUpdateInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.UpdateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UpdateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeUpdateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpdateId != nil {
		v := *s.UpdateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "updateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/DescribeUpdateResponse
type DescribeUpdateOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the specified update.
	Update *Update `json:"eks:DescribeUpdateOutput:Update" locationName:"update" type:"structure"`
}

// String returns the string representation
func (s DescribeUpdateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeUpdateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Update != nil {
		v := s.Update

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "update", v, metadata)
	}
	return nil
}

const opDescribeUpdate = "DescribeUpdate"

// DescribeUpdateRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Returns descriptive information about an update against your Amazon EKS cluster.
//
// When the status of the update is Succeeded, the update is complete. If an
// update fails, the status is Failed, and an error detail explains the reason
// for the failure.
//
//    // Example sending a request using DescribeUpdateRequest.
//    req := client.DescribeUpdateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/DescribeUpdate
func (c *Client) DescribeUpdateRequest(input *DescribeUpdateInput) DescribeUpdateRequest {
	op := &aws.Operation{
		Name:       opDescribeUpdate,
		HTTPMethod: "GET",
		HTTPPath:   "/clusters/{name}/updates/{updateId}",
	}

	if input == nil {
		input = &DescribeUpdateInput{}
	}

	req := c.newRequest(op, input, &DescribeUpdateOutput{})
	return DescribeUpdateRequest{Request: req, Input: input, Copy: c.DescribeUpdateRequest}
}

// DescribeUpdateRequest is the request type for the
// DescribeUpdate API operation.
type DescribeUpdateRequest struct {
	*aws.Request
	Input *DescribeUpdateInput
	Copy  func(*DescribeUpdateInput) DescribeUpdateRequest
}

// Send marshals and sends the DescribeUpdate API request.
func (r DescribeUpdateRequest) Send(ctx context.Context) (*DescribeUpdateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUpdateResponse{
		DescribeUpdateOutput: r.Request.Data.(*DescribeUpdateOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUpdateResponse is the response type for the
// DescribeUpdate API operation.
type DescribeUpdateResponse struct {
	*DescribeUpdateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUpdate request.
func (r *DescribeUpdateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
