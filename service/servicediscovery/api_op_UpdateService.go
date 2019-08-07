// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/UpdateServiceRequest
type UpdateServiceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the service that you want to update.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// A complex type that contains the new settings for the service.
	//
	// Service is a required field
	Service *ServiceChange `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateServiceInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.Service == nil {
		invalidParams.Add(aws.NewErrParamRequired("Service"))
	}
	if s.Service != nil {
		if err := s.Service.Validate(); err != nil {
			invalidParams.AddNested("Service", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/UpdateServiceResponse
type UpdateServiceOutput struct {
	_ struct{} `type:"structure"`

	// A value that you can use to determine whether the request completed successfully.
	// To get the status of the operation, see GetOperation.
	OperationId *string `json:"servicediscovery:UpdateServiceOutput:OperationId" type:"string"`
}

// String returns the string representation
func (s UpdateServiceOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateService = "UpdateService"

// UpdateServiceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Submits a request to perform the following operations:
//
//    * Add or delete DnsRecords configurations
//
//    * Update the TTL setting for existing DnsRecords configurations
//
//    * Add, update, or delete HealthCheckConfig for a specified service
//
// For public and private DNS namespaces, you must specify all DnsRecords configurations
// (and, optionally, HealthCheckConfig) that you want to appear in the updated
// service. Any current configurations that don't appear in an UpdateService
// request are deleted.
//
// When you update the TTL setting for a service, AWS Cloud Map also updates
// the corresponding settings in all the records and health checks that were
// created by using the specified service.
//
//    // Example sending a request using UpdateServiceRequest.
//    req := client.UpdateServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/UpdateService
func (c *Client) UpdateServiceRequest(input *UpdateServiceInput) UpdateServiceRequest {
	op := &aws.Operation{
		Name:       opUpdateService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateServiceInput{}
	}

	req := c.newRequest(op, input, &UpdateServiceOutput{})
	return UpdateServiceRequest{Request: req, Input: input, Copy: c.UpdateServiceRequest}
}

// UpdateServiceRequest is the request type for the
// UpdateService API operation.
type UpdateServiceRequest struct {
	*aws.Request
	Input *UpdateServiceInput
	Copy  func(*UpdateServiceInput) UpdateServiceRequest
}

// Send marshals and sends the UpdateService API request.
func (r UpdateServiceRequest) Send(ctx context.Context) (*UpdateServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateServiceResponse{
		UpdateServiceOutput: r.Request.Data.(*UpdateServiceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateServiceResponse is the response type for the
// UpdateService API operation.
type UpdateServiceResponse struct {
	*UpdateServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateService request.
func (r *UpdateServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
