// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/UpdateServerRequest
type UpdateServerInput struct {
	_ struct{} `type:"structure"`

	// Sets the number of automated backups that you want to keep.
	BackupRetentionCount *int64 `type:"integer"`

	// Setting DisableAutomatedBackup to true disables automated or scheduled backups.
	// Automated backups are enabled by default.
	DisableAutomatedBackup *bool `type:"boolean"`

	// DDD:HH:MM (weekly start time) or HH:MM (daily start time).
	//
	// Time windows always use coordinated universal time (UTC). Valid strings for
	// day of week (DDD) are: Mon, Tue, Wed, Thr, Fri, Sat, or Sun.
	PreferredBackupWindow *string `type:"string"`

	// DDD:HH:MM (weekly start time) or HH:MM (daily start time).
	//
	// Time windows always use coordinated universal time (UTC). Valid strings for
	// day of week (DDD) are: Mon, Tue, Wed, Thr, Fri, Sat, or Sun.
	PreferredMaintenanceWindow *string `type:"string"`

	// The name of the server to update.
	//
	// ServerName is a required field
	ServerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateServerInput"}

	if s.ServerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerName"))
	}
	if s.ServerName != nil && len(*s.ServerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/UpdateServerResponse
type UpdateServerOutput struct {
	_ struct{} `type:"structure"`

	// Contains the response to a UpdateServer request.
	Server *Server `json:"opsworks-cm:UpdateServerOutput:Server" type:"structure"`
}

// String returns the string representation
func (s UpdateServerOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateServer = "UpdateServer"

// UpdateServerRequest returns a request value for making API operation for
// AWS OpsWorks for Chef Automate.
//
// Updates settings for a server.
//
// This operation is synchronous.
//
//    // Example sending a request using UpdateServerRequest.
//    req := client.UpdateServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/UpdateServer
func (c *Client) UpdateServerRequest(input *UpdateServerInput) UpdateServerRequest {
	op := &aws.Operation{
		Name:       opUpdateServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateServerInput{}
	}

	req := c.newRequest(op, input, &UpdateServerOutput{})
	return UpdateServerRequest{Request: req, Input: input, Copy: c.UpdateServerRequest}
}

// UpdateServerRequest is the request type for the
// UpdateServer API operation.
type UpdateServerRequest struct {
	*aws.Request
	Input *UpdateServerInput
	Copy  func(*UpdateServerInput) UpdateServerRequest
}

// Send marshals and sends the UpdateServer API request.
func (r UpdateServerRequest) Send(ctx context.Context) (*UpdateServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateServerResponse{
		UpdateServerOutput: r.Request.Data.(*UpdateServerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateServerResponse is the response type for the
// UpdateServer API operation.
type UpdateServerResponse struct {
	*UpdateServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateServer request.
func (r *UpdateServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
