// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/DeleteServiceQuotaIncreaseRequestFromTemplateRequest
type DeleteServiceQuotaIncreaseRequestFromTemplateInput struct {
	_ struct{} `type:"structure"`

	// Specifies the AWS Region for the quota that you want to delete.
	//
	// AwsRegion is a required field
	AwsRegion *string `min:"1" type:"string" required:"true"`

	// Specifies the code for the quota that you want to delete.
	//
	// QuotaCode is a required field
	QuotaCode *string `min:"1" type:"string" required:"true"`

	// Specifies the code for the service that you want to delete.
	//
	// ServiceCode is a required field
	ServiceCode *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteServiceQuotaIncreaseRequestFromTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteServiceQuotaIncreaseRequestFromTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteServiceQuotaIncreaseRequestFromTemplateInput"}

	if s.AwsRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsRegion"))
	}
	if s.AwsRegion != nil && len(*s.AwsRegion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsRegion", 1))
	}

	if s.QuotaCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("QuotaCode"))
	}
	if s.QuotaCode != nil && len(*s.QuotaCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QuotaCode", 1))
	}

	if s.ServiceCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceCode"))
	}
	if s.ServiceCode != nil && len(*s.ServiceCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/DeleteServiceQuotaIncreaseRequestFromTemplateResponse
type DeleteServiceQuotaIncreaseRequestFromTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteServiceQuotaIncreaseRequestFromTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteServiceQuotaIncreaseRequestFromTemplate = "DeleteServiceQuotaIncreaseRequestFromTemplate"

// DeleteServiceQuotaIncreaseRequestFromTemplateRequest returns a request value for making API operation for
// Service Quotas.
//
// Removes a service quota increase request from the Service Quotas template.
//
//    // Example sending a request using DeleteServiceQuotaIncreaseRequestFromTemplateRequest.
//    req := client.DeleteServiceQuotaIncreaseRequestFromTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/DeleteServiceQuotaIncreaseRequestFromTemplate
func (c *Client) DeleteServiceQuotaIncreaseRequestFromTemplateRequest(input *DeleteServiceQuotaIncreaseRequestFromTemplateInput) DeleteServiceQuotaIncreaseRequestFromTemplateRequest {
	op := &aws.Operation{
		Name:       opDeleteServiceQuotaIncreaseRequestFromTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteServiceQuotaIncreaseRequestFromTemplateInput{}
	}

	req := c.newRequest(op, input, &DeleteServiceQuotaIncreaseRequestFromTemplateOutput{})
	return DeleteServiceQuotaIncreaseRequestFromTemplateRequest{Request: req, Input: input, Copy: c.DeleteServiceQuotaIncreaseRequestFromTemplateRequest}
}

// DeleteServiceQuotaIncreaseRequestFromTemplateRequest is the request type for the
// DeleteServiceQuotaIncreaseRequestFromTemplate API operation.
type DeleteServiceQuotaIncreaseRequestFromTemplateRequest struct {
	*aws.Request
	Input *DeleteServiceQuotaIncreaseRequestFromTemplateInput
	Copy  func(*DeleteServiceQuotaIncreaseRequestFromTemplateInput) DeleteServiceQuotaIncreaseRequestFromTemplateRequest
}

// Send marshals and sends the DeleteServiceQuotaIncreaseRequestFromTemplate API request.
func (r DeleteServiceQuotaIncreaseRequestFromTemplateRequest) Send(ctx context.Context) (*DeleteServiceQuotaIncreaseRequestFromTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteServiceQuotaIncreaseRequestFromTemplateResponse{
		DeleteServiceQuotaIncreaseRequestFromTemplateOutput: r.Request.Data.(*DeleteServiceQuotaIncreaseRequestFromTemplateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteServiceQuotaIncreaseRequestFromTemplateResponse is the response type for the
// DeleteServiceQuotaIncreaseRequestFromTemplate API operation.
type DeleteServiceQuotaIncreaseRequestFromTemplateResponse struct {
	*DeleteServiceQuotaIncreaseRequestFromTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteServiceQuotaIncreaseRequestFromTemplate request.
func (r *DeleteServiceQuotaIncreaseRequestFromTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
