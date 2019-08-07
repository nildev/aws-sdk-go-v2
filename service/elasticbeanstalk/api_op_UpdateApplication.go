// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to update an application.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/UpdateApplicationMessage
type UpdateApplicationInput struct {
	_ struct{} `type:"structure"`

	// The name of the application to update. If no such application is found, UpdateApplication
	// returns an InvalidParameterValue error.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// A new description for the application.
	//
	// Default: If not specified, AWS Elastic Beanstalk does not update the description.
	Description *string `type:"string"`
}

// String returns the string representation
func (s UpdateApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApplicationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result message containing a single description of an application.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/ApplicationDescriptionMessage
type UpdateApplicationOutput struct {
	_ struct{} `type:"structure"`

	// The ApplicationDescription of the application.
	Application *ApplicationDescription `json:"elasticbeanstalk:UpdateApplicationOutput:Application" type:"structure"`
}

// String returns the string representation
func (s UpdateApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateApplication = "UpdateApplication"

// UpdateApplicationRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Updates the specified application to have the specified properties.
//
// If a property (for example, description) is not provided, the value remains
// unchanged. To clear these properties, specify an empty string.
//
//    // Example sending a request using UpdateApplicationRequest.
//    req := client.UpdateApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/UpdateApplication
func (c *Client) UpdateApplicationRequest(input *UpdateApplicationInput) UpdateApplicationRequest {
	op := &aws.Operation{
		Name:       opUpdateApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateApplicationInput{}
	}

	req := c.newRequest(op, input, &UpdateApplicationOutput{})
	return UpdateApplicationRequest{Request: req, Input: input, Copy: c.UpdateApplicationRequest}
}

// UpdateApplicationRequest is the request type for the
// UpdateApplication API operation.
type UpdateApplicationRequest struct {
	*aws.Request
	Input *UpdateApplicationInput
	Copy  func(*UpdateApplicationInput) UpdateApplicationRequest
}

// Send marshals and sends the UpdateApplication API request.
func (r UpdateApplicationRequest) Send(ctx context.Context) (*UpdateApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApplicationResponse{
		UpdateApplicationOutput: r.Request.Data.(*UpdateApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApplicationResponse is the response type for the
// UpdateApplication API operation.
type UpdateApplicationResponse struct {
	*UpdateApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApplication request.
func (r *UpdateApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
