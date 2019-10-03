// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ImportVolumeRequest
type ImportVolumeInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone for the resulting EBS volume.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `locationName:"availabilityZone" type:"string" required:"true"`

	// A description of the volume.
	Description *string `locationName:"description" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The disk image.
	//
	// Image is a required field
	Image *DiskImageDetail `locationName:"image" type:"structure" required:"true"`

	// The volume size.
	//
	// Volume is a required field
	Volume *VolumeDetail `locationName:"volume" type:"structure" required:"true"`
}

// String returns the string representation
func (s ImportVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportVolumeInput"}

	if s.AvailabilityZone == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZone"))
	}

	if s.Image == nil {
		invalidParams.Add(aws.NewErrParamRequired("Image"))
	}

	if s.Volume == nil {
		invalidParams.Add(aws.NewErrParamRequired("Volume"))
	}
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			invalidParams.AddNested("Image", err.(aws.ErrInvalidParams))
		}
	}
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			invalidParams.AddNested("Volume", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ImportVolumeResult
type ImportVolumeOutput struct {
	_ struct{} `type:"structure"`

	// Information about the conversion task.
	ConversionTask *ConversionTask `json:"ec2:ImportVolumeOutput:ConversionTask" locationName:"conversionTask" type:"structure"`
}

// String returns the string representation
func (s ImportVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opImportVolume = "ImportVolume"

// ImportVolumeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates an import volume task using metadata from the specified disk image.For
// more information, see Importing Disks to Amazon EBS (https://docs.aws.amazon.com/AWSEC2/latest/CommandLineReference/importing-your-volumes-into-amazon-ebs.html).
//
// For information about the import manifest referenced by this API action,
// see VM Import Manifest (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/manifest.html).
//
//    // Example sending a request using ImportVolumeRequest.
//    req := client.ImportVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ImportVolume
func (c *Client) ImportVolumeRequest(input *ImportVolumeInput) ImportVolumeRequest {
	op := &aws.Operation{
		Name:       opImportVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ImportVolumeInput{}
	}

	req := c.newRequest(op, input, &ImportVolumeOutput{})
	return ImportVolumeRequest{Request: req, Input: input, Copy: c.ImportVolumeRequest}
}

// ImportVolumeRequest is the request type for the
// ImportVolume API operation.
type ImportVolumeRequest struct {
	*aws.Request
	Input *ImportVolumeInput
	Copy  func(*ImportVolumeInput) ImportVolumeRequest
}

// Send marshals and sends the ImportVolume API request.
func (r ImportVolumeRequest) Send(ctx context.Context) (*ImportVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportVolumeResponse{
		ImportVolumeOutput: r.Request.Data.(*ImportVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportVolumeResponse is the response type for the
// ImportVolume API operation.
type ImportVolumeResponse struct {
	*ImportVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportVolume request.
func (r *ImportVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
