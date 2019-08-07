// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/CreateLicenseConfigurationRequest
type CreateLicenseConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Human-friendly description of the license configuration.
	Description *string `type:"string"`

	// Number of licenses managed by the license configuration.
	LicenseCount *int64 `type:"long"`

	// Flag indicating whether hard or soft license enforcement is used. Exceeding
	// a hard limit results in the blocked deployment of new instances.
	LicenseCountHardLimit *bool `type:"boolean"`

	// Dimension to use to track the license inventory.
	//
	// LicenseCountingType is a required field
	LicenseCountingType LicenseCountingType `type:"string" required:"true" enum:"true"`

	// Array of configured License Manager rules.
	LicenseRules []string `type:"list"`

	// Name of the license configuration.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The tags to apply to the resources during launch. You can only tag instances
	// and volumes on launch. The specified tags are applied to all instances or
	// volumes that are created during launch. To tag a resource after it has been
	// created, see CreateTags .
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateLicenseConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLicenseConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLicenseConfigurationInput"}
	if len(s.LicenseCountingType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LicenseCountingType"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/CreateLicenseConfigurationResponse
type CreateLicenseConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// ARN of the license configuration object after its creation.
	LicenseConfigurationArn *string `json:"license-manager:CreateLicenseConfigurationOutput:LicenseConfigurationArn" type:"string"`
}

// String returns the string representation
func (s CreateLicenseConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLicenseConfiguration = "CreateLicenseConfiguration"

// CreateLicenseConfigurationRequest returns a request value for making API operation for
// AWS License Manager.
//
// Creates a new license configuration object. A license configuration is an
// abstraction of a customer license agreement that can be consumed and enforced
// by License Manager. Components include specifications for the license type
// (licensing by instance, socket, CPU, or VCPU), tenancy (shared tenancy, Amazon
// EC2 Dedicated Instance, Amazon EC2 Dedicated Host, or any of these), host
// affinity (how long a VM must be associated with a host), the number of licenses
// purchased and used.
//
//    // Example sending a request using CreateLicenseConfigurationRequest.
//    req := client.CreateLicenseConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/CreateLicenseConfiguration
func (c *Client) CreateLicenseConfigurationRequest(input *CreateLicenseConfigurationInput) CreateLicenseConfigurationRequest {
	op := &aws.Operation{
		Name:       opCreateLicenseConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLicenseConfigurationInput{}
	}

	req := c.newRequest(op, input, &CreateLicenseConfigurationOutput{})
	return CreateLicenseConfigurationRequest{Request: req, Input: input, Copy: c.CreateLicenseConfigurationRequest}
}

// CreateLicenseConfigurationRequest is the request type for the
// CreateLicenseConfiguration API operation.
type CreateLicenseConfigurationRequest struct {
	*aws.Request
	Input *CreateLicenseConfigurationInput
	Copy  func(*CreateLicenseConfigurationInput) CreateLicenseConfigurationRequest
}

// Send marshals and sends the CreateLicenseConfiguration API request.
func (r CreateLicenseConfigurationRequest) Send(ctx context.Context) (*CreateLicenseConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLicenseConfigurationResponse{
		CreateLicenseConfigurationOutput: r.Request.Data.(*CreateLicenseConfigurationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLicenseConfigurationResponse is the response type for the
// CreateLicenseConfiguration API operation.
type CreateLicenseConfigurationResponse struct {
	*CreateLicenseConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLicenseConfiguration request.
func (r *CreateLicenseConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
