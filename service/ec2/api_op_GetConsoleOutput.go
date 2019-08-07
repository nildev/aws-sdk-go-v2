// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetConsoleOutputRequest
type GetConsoleOutputInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// When enabled, retrieves the latest console output for the instance.
	//
	// Default: disabled (false)
	Latest *bool `type:"boolean"`
}

// String returns the string representation
func (s GetConsoleOutputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConsoleOutputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConsoleOutputInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetConsoleOutputResult
type GetConsoleOutputOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the instance.
	InstanceId *string `json:"ec2:GetConsoleOutputOutput:InstanceId" locationName:"instanceId" type:"string"`

	// The console output, base64-encoded. If you are using a command line tool,
	// the tool decodes the output for you.
	Output *string `json:"ec2:GetConsoleOutputOutput:Output" locationName:"output" type:"string"`

	// The time at which the output was last updated.
	Timestamp *time.Time `json:"ec2:GetConsoleOutputOutput:Timestamp" locationName:"timestamp" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s GetConsoleOutputOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetConsoleOutput = "GetConsoleOutput"

// GetConsoleOutputRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Gets the console output for the specified instance. For Linux instances,
// the instance console output displays the exact console output that would
// normally be displayed on a physical monitor attached to a computer. For Windows
// instances, the instance console output includes the last three system event
// log errors.
//
// By default, the console output returns buffered information that was posted
// shortly after an instance transition state (start, stop, reboot, or terminate).
// This information is available for at least one hour after the most recent
// post. Only the most recent 64 KB of console output is available.
//
// You can optionally retrieve the latest serial console output at any time
// during the instance lifecycle. This option is supported on instance types
// that use the Nitro hypervisor.
//
// For more information, see Instance Console Output (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-console.html#instance-console-console-output)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using GetConsoleOutputRequest.
//    req := client.GetConsoleOutputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetConsoleOutput
func (c *Client) GetConsoleOutputRequest(input *GetConsoleOutputInput) GetConsoleOutputRequest {
	op := &aws.Operation{
		Name:       opGetConsoleOutput,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConsoleOutputInput{}
	}

	req := c.newRequest(op, input, &GetConsoleOutputOutput{})
	return GetConsoleOutputRequest{Request: req, Input: input, Copy: c.GetConsoleOutputRequest}
}

// GetConsoleOutputRequest is the request type for the
// GetConsoleOutput API operation.
type GetConsoleOutputRequest struct {
	*aws.Request
	Input *GetConsoleOutputInput
	Copy  func(*GetConsoleOutputInput) GetConsoleOutputRequest
}

// Send marshals and sends the GetConsoleOutput API request.
func (r GetConsoleOutputRequest) Send(ctx context.Context) (*GetConsoleOutputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConsoleOutputResponse{
		GetConsoleOutputOutput: r.Request.Data.(*GetConsoleOutputOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConsoleOutputResponse is the response type for the
// GetConsoleOutput API operation.
type GetConsoleOutputResponse struct {
	*GetConsoleOutputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConsoleOutput request.
func (r *GetConsoleOutputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
