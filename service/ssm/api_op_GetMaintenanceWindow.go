// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetMaintenanceWindowRequest
type GetMaintenanceWindowInput struct {
	_ struct{} `type:"structure"`

	// The ID of the maintenance window for which you want to retrieve information.
	//
	// WindowId is a required field
	WindowId *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMaintenanceWindowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMaintenanceWindowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMaintenanceWindowInput"}

	if s.WindowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowId"))
	}
	if s.WindowId != nil && len(*s.WindowId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowId", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetMaintenanceWindowResult
type GetMaintenanceWindowOutput struct {
	_ struct{} `type:"structure"`

	// Whether targets must be registered with the maintenance window before tasks
	// can be defined for those targets.
	AllowUnassociatedTargets *bool `type:"boolean"`

	// The date the maintenance window was created.
	CreatedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The number of hours before the end of the maintenance window that Systems
	// Manager stops scheduling new tasks for execution.
	Cutoff *int64 `type:"integer"`

	// The description of the maintenance window.
	Description *string `min:"1" type:"string"`

	// The duration of the maintenance window in hours.
	Duration *int64 `min:"1" type:"integer"`

	// Indicates whether the maintenance window is enabled.
	Enabled *bool `type:"boolean"`

	// The date and time, in ISO-8601 Extended format, for when the maintenance
	// window is scheduled to become inactive. The maintenance window will not run
	// after this specified time.
	EndDate *string `type:"string"`

	// The date the maintenance window was last modified.
	ModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The name of the maintenance window.
	Name *string `min:"3" type:"string"`

	// The next time the maintenance window will actually run, taking into account
	// any specified times for the maintenance window to become active or inactive.
	NextExecutionTime *string `type:"string"`

	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string `min:"1" type:"string"`

	// The time zone that the scheduled maintenance window executions are based
	// on, in Internet Assigned Numbers Authority (IANA) format. For example: "America/Los_Angeles",
	// "etc/UTC", or "Asia/Seoul". For more information, see the Time Zone Database
	// (https://www.iana.org/time-zones) on the IANA website.
	ScheduleTimezone *string `type:"string"`

	// The date and time, in ISO-8601 Extended format, for when the maintenance
	// window is scheduled to become active. The maintenance window will not run
	// before this specified time.
	StartDate *string `type:"string"`

	// The ID of the created maintenance window.
	WindowId *string `min:"20" type:"string"`
}

// String returns the string representation
func (s GetMaintenanceWindowOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetMaintenanceWindow = "GetMaintenanceWindow"

// GetMaintenanceWindowRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves a maintenance window.
//
//    // Example sending a request using GetMaintenanceWindowRequest.
//    req := client.GetMaintenanceWindowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetMaintenanceWindow
func (c *Client) GetMaintenanceWindowRequest(input *GetMaintenanceWindowInput) GetMaintenanceWindowRequest {
	op := &aws.Operation{
		Name:       opGetMaintenanceWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetMaintenanceWindowInput{}
	}

	req := c.newRequest(op, input, &GetMaintenanceWindowOutput{})
	return GetMaintenanceWindowRequest{Request: req, Input: input, Copy: c.GetMaintenanceWindowRequest}
}

// GetMaintenanceWindowRequest is the request type for the
// GetMaintenanceWindow API operation.
type GetMaintenanceWindowRequest struct {
	*aws.Request
	Input *GetMaintenanceWindowInput
	Copy  func(*GetMaintenanceWindowInput) GetMaintenanceWindowRequest
}

// Send marshals and sends the GetMaintenanceWindow API request.
func (r GetMaintenanceWindowRequest) Send(ctx context.Context) (*GetMaintenanceWindowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMaintenanceWindowResponse{
		GetMaintenanceWindowOutput: r.Request.Data.(*GetMaintenanceWindowOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMaintenanceWindowResponse is the response type for the
// GetMaintenanceWindow API operation.
type GetMaintenanceWindowResponse struct {
	*GetMaintenanceWindowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMaintenanceWindow request.
func (r *GetMaintenanceWindowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
