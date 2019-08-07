// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/GetLogRecordRequest
type GetLogRecordInput struct {
	_ struct{} `type:"structure"`

	// The pointer corresponding to the log event record you want to retrieve. You
	// get this from the response of a GetQueryResults operation. In that response,
	// the value of the @ptr field for a log event is the value to use as logRecordPointer
	// to retrieve that complete log event record.
	//
	// LogRecordPointer is a required field
	LogRecordPointer *string `locationName:"logRecordPointer" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLogRecordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLogRecordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLogRecordInput"}

	if s.LogRecordPointer == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogRecordPointer"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/GetLogRecordResponse
type GetLogRecordOutput struct {
	_ struct{} `type:"structure"`

	// The requested log event, as a JSON string.
	LogRecord map[string]string `json:"logs:GetLogRecordOutput:LogRecord" locationName:"logRecord" type:"map"`
}

// String returns the string representation
func (s GetLogRecordOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetLogRecord = "GetLogRecord"

// GetLogRecordRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Retrieves all the fields and values of a single log event. All fields are
// retrieved, even if the original query that produced the logRecordPointer
// retrieved only a subset of fields. Fields are returned as field name/field
// value pairs.
//
// Additionally, the entire unparsed log event is returned within @message.
//
//    // Example sending a request using GetLogRecordRequest.
//    req := client.GetLogRecordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/GetLogRecord
func (c *Client) GetLogRecordRequest(input *GetLogRecordInput) GetLogRecordRequest {
	op := &aws.Operation{
		Name:       opGetLogRecord,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetLogRecordInput{}
	}

	req := c.newRequest(op, input, &GetLogRecordOutput{})
	return GetLogRecordRequest{Request: req, Input: input, Copy: c.GetLogRecordRequest}
}

// GetLogRecordRequest is the request type for the
// GetLogRecord API operation.
type GetLogRecordRequest struct {
	*aws.Request
	Input *GetLogRecordInput
	Copy  func(*GetLogRecordInput) GetLogRecordRequest
}

// Send marshals and sends the GetLogRecord API request.
func (r GetLogRecordRequest) Send(ctx context.Context) (*GetLogRecordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLogRecordResponse{
		GetLogRecordOutput: r.Request.Data.(*GetLogRecordOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLogRecordResponse is the response type for the
// GetLogRecord API operation.
type GetLogRecordResponse struct {
	*GetLogRecordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLogRecord request.
func (r *GetLogRecordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
