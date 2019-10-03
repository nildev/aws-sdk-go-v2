// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for RegisterImage.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RegisterImageRequest
type RegisterImageInput struct {
	_ struct{} `type:"structure"`

	// The architecture of the AMI.
	//
	// Default: For Amazon EBS-backed AMIs, i386. For instance store-backed AMIs,
	// the architecture specified in the manifest file.
	Architecture ArchitectureValues `locationName:"architecture" type:"string" enum:"true"`

	// The billing product codes. Your account must be authorized to specify billing
	// product codes. Otherwise, you can use the AWS Marketplace to bill for the
	// use of an AMI.
	BillingProducts []string `locationName:"BillingProduct" locationNameList:"item" type:"list"`

	// The block device mapping entries.
	BlockDeviceMappings []BlockDeviceMapping `locationName:"BlockDeviceMapping" locationNameList:"BlockDeviceMapping" type:"list"`

	// A description for your AMI.
	Description *string `locationName:"description" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Set to true to enable enhanced networking with ENA for the AMI and any instances
	// that you launch from the AMI.
	//
	// This option is supported only for HVM AMIs. Specifying this option with a
	// PV AMI can make instances launched from the AMI unreachable.
	EnaSupport *bool `locationName:"enaSupport" type:"boolean"`

	// The full path to your AMI manifest in Amazon S3 storage. The specified bucket
	// must have the aws-exec-read canned access control list (ACL) to ensure that
	// it can be accessed by Amazon EC2. For more information, see Canned ACLs (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
	// in the Amazon S3 Service Developer Guide.
	ImageLocation *string `type:"string"`

	// The ID of the kernel.
	KernelId *string `locationName:"kernelId" type:"string"`

	// A name for your AMI.
	//
	// Constraints: 3-128 alphanumeric characters, parentheses (()), square brackets
	// ([]), spaces ( ), periods (.), slashes (/), dashes (-), single quotes ('),
	// at-signs (@), or underscores(_)
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The ID of the RAM disk.
	RamdiskId *string `locationName:"ramdiskId" type:"string"`

	// The device name of the root device volume (for example, /dev/sda1).
	RootDeviceName *string `locationName:"rootDeviceName" type:"string"`

	// Set to simple to enable enhanced networking with the Intel 82599 Virtual
	// Function interface for the AMI and any instances that you launch from the
	// AMI.
	//
	// There is no way to disable sriovNetSupport at this time.
	//
	// This option is supported only for HVM AMIs. Specifying this option with a
	// PV AMI can make instances launched from the AMI unreachable.
	SriovNetSupport *string `locationName:"sriovNetSupport" type:"string"`

	// The type of virtualization (hvm | paravirtual).
	//
	// Default: paravirtual
	VirtualizationType *string `locationName:"virtualizationType" type:"string"`
}

// String returns the string representation
func (s RegisterImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterImageInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of RegisterImage.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RegisterImageResult
type RegisterImageOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the newly registered AMI.
	ImageId *string `json:"ec2:RegisterImageOutput:ImageId" locationName:"imageId" type:"string"`
}

// String returns the string representation
func (s RegisterImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterImage = "RegisterImage"

// RegisterImageRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Registers an AMI. When you're creating an AMI, this is the final step you
// must complete before you can launch an instance from the AMI. For more information
// about creating AMIs, see Creating Your Own AMIs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// For Amazon EBS-backed instances, CreateImage creates and registers the AMI
// in a single request, so you don't have to register the AMI yourself.
//
// You can also use RegisterImage to create an Amazon EBS-backed Linux AMI from
// a snapshot of a root device volume. You specify the snapshot using the block
// device mapping. For more information, see Launching a Linux Instance from
// a Backup (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-launch-snapshot.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// You can't register an image where a secondary (non-root) snapshot has AWS
// Marketplace product codes.
//
// Some Linux distributions, such as Red Hat Enterprise Linux (RHEL) and SUSE
// Linux Enterprise Server (SLES), use the EC2 billing product code associated
// with an AMI to verify the subscription status for package updates. Creating
// an AMI from an EBS snapshot does not maintain this billing code, and instances
// launched from such an AMI are not able to connect to package update infrastructure.
// If you purchase a Reserved Instance offering for one of these Linux distributions
// and launch instances using an AMI that does not contain the required billing
// code, your Reserved Instance is not applied to these instances.
//
// To create an AMI for operating systems that require a billing code, see CreateImage.
//
// If needed, you can deregister an AMI at any time. Any modifications you make
// to an AMI backed by an instance store volume invalidates its registration.
// If you make changes to an image, deregister the previous image and register
// the new image.
//
//    // Example sending a request using RegisterImageRequest.
//    req := client.RegisterImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RegisterImage
func (c *Client) RegisterImageRequest(input *RegisterImageInput) RegisterImageRequest {
	op := &aws.Operation{
		Name:       opRegisterImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterImageInput{}
	}

	req := c.newRequest(op, input, &RegisterImageOutput{})
	return RegisterImageRequest{Request: req, Input: input, Copy: c.RegisterImageRequest}
}

// RegisterImageRequest is the request type for the
// RegisterImage API operation.
type RegisterImageRequest struct {
	*aws.Request
	Input *RegisterImageInput
	Copy  func(*RegisterImageInput) RegisterImageRequest
}

// Send marshals and sends the RegisterImage API request.
func (r RegisterImageRequest) Send(ctx context.Context) (*RegisterImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterImageResponse{
		RegisterImageOutput: r.Request.Data.(*RegisterImageOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterImageResponse is the response type for the
// RegisterImage API operation.
type RegisterImageResponse struct {
	*RegisterImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterImage request.
func (r *RegisterImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
