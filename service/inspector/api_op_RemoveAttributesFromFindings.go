// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/RemoveAttributesFromFindingsRequest
type RemoveAttributesFromFindingsInput struct {
	_ struct{} `type:"structure"`

	// The array of attribute keys that you want to remove from specified findings.
	//
	// AttributeKeys is a required field
	AttributeKeys []string `locationName:"attributeKeys" type:"list" required:"true"`

	// The ARNs that specify the findings that you want to remove attributes from.
	//
	// FindingArns is a required field
	FindingArns []string `locationName:"findingArns" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveAttributesFromFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveAttributesFromFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveAttributesFromFindingsInput"}

	if s.AttributeKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttributeKeys"))
	}

	if s.FindingArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("FindingArns"))
	}
	if s.FindingArns != nil && len(s.FindingArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FindingArns", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/RemoveAttributesFromFindingsResponse
type RemoveAttributesFromFindingsOutput struct {
	_ struct{} `type:"structure"`

	// Attributes details that cannot be described. An error code is provided for
	// each failed item.
	//
	// FailedItems is a required field
	FailedItems map[string]FailedItemDetails `json:"inspector:RemoveAttributesFromFindingsOutput:FailedItems" locationName:"failedItems" type:"map" required:"true"`
}

// String returns the string representation
func (s RemoveAttributesFromFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveAttributesFromFindings = "RemoveAttributesFromFindings"

// RemoveAttributesFromFindingsRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Removes entire attributes (key and value pairs) from the findings that are
// specified by the ARNs of the findings where an attribute with the specified
// key exists.
//
//    // Example sending a request using RemoveAttributesFromFindingsRequest.
//    req := client.RemoveAttributesFromFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/RemoveAttributesFromFindings
func (c *Client) RemoveAttributesFromFindingsRequest(input *RemoveAttributesFromFindingsInput) RemoveAttributesFromFindingsRequest {
	op := &aws.Operation{
		Name:       opRemoveAttributesFromFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveAttributesFromFindingsInput{}
	}

	req := c.newRequest(op, input, &RemoveAttributesFromFindingsOutput{})
	return RemoveAttributesFromFindingsRequest{Request: req, Input: input, Copy: c.RemoveAttributesFromFindingsRequest}
}

// RemoveAttributesFromFindingsRequest is the request type for the
// RemoveAttributesFromFindings API operation.
type RemoveAttributesFromFindingsRequest struct {
	*aws.Request
	Input *RemoveAttributesFromFindingsInput
	Copy  func(*RemoveAttributesFromFindingsInput) RemoveAttributesFromFindingsRequest
}

// Send marshals and sends the RemoveAttributesFromFindings API request.
func (r RemoveAttributesFromFindingsRequest) Send(ctx context.Context) (*RemoveAttributesFromFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveAttributesFromFindingsResponse{
		RemoveAttributesFromFindingsOutput: r.Request.Data.(*RemoveAttributesFromFindingsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveAttributesFromFindingsResponse is the response type for the
// RemoveAttributesFromFindings API operation.
type RemoveAttributesFromFindingsResponse struct {
	*RemoveAttributesFromFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveAttributesFromFindings request.
func (r *RemoveAttributesFromFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
