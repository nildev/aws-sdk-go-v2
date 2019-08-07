// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/CreateNamedQueryInput
type CreateNamedQueryInput struct {
	_ struct{} `type:"structure"`

	// A unique case-sensitive string used to ensure the request to create the query
	// is idempotent (executes only once). If another CreateNamedQuery request is
	// received, the same response is returned and another query is not created.
	// If a parameter has changed, for example, the QueryString, an error is returned.
	//
	// This token is listed as not required because AWS SDKs (for example the AWS
	// SDK for Java) auto-generate the token for users. If you are not using the
	// AWS SDK or the AWS CLI, you must provide this token or the action will fail.
	ClientRequestToken *string `min:"32" type:"string" idempotencyToken:"true"`

	// The database to which the query belongs.
	//
	// Database is a required field
	Database *string `min:"1" type:"string" required:"true"`

	// The query description.
	Description *string `min:"1" type:"string"`

	// The query name.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The contents of the query with all query statements.
	//
	// QueryString is a required field
	QueryString *string `min:"1" type:"string" required:"true"`

	// The name of the workgroup in which the named query is being created.
	WorkGroup *string `type:"string"`
}

// String returns the string representation
func (s CreateNamedQueryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNamedQueryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNamedQueryInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 32))
	}

	if s.Database == nil {
		invalidParams.Add(aws.NewErrParamRequired("Database"))
	}
	if s.Database != nil && len(*s.Database) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Database", 1))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.QueryString == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryString"))
	}
	if s.QueryString != nil && len(*s.QueryString) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QueryString", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/CreateNamedQueryOutput
type CreateNamedQueryOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the query.
	NamedQueryId *string `json:"athena:CreateNamedQueryOutput:NamedQueryId" type:"string"`
}

// String returns the string representation
func (s CreateNamedQueryOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateNamedQuery = "CreateNamedQuery"

// CreateNamedQueryRequest returns a request value for making API operation for
// Amazon Athena.
//
// Creates a named query in the specified workgroup. Requires that you have
// access to the workgroup.
//
// For code samples using the AWS SDK for Java, see Examples and Code Samples
// (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in the Amazon
// Athena User Guide.
//
//    // Example sending a request using CreateNamedQueryRequest.
//    req := client.CreateNamedQueryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/CreateNamedQuery
func (c *Client) CreateNamedQueryRequest(input *CreateNamedQueryInput) CreateNamedQueryRequest {
	op := &aws.Operation{
		Name:       opCreateNamedQuery,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNamedQueryInput{}
	}

	req := c.newRequest(op, input, &CreateNamedQueryOutput{})
	return CreateNamedQueryRequest{Request: req, Input: input, Copy: c.CreateNamedQueryRequest}
}

// CreateNamedQueryRequest is the request type for the
// CreateNamedQuery API operation.
type CreateNamedQueryRequest struct {
	*aws.Request
	Input *CreateNamedQueryInput
	Copy  func(*CreateNamedQueryInput) CreateNamedQueryRequest
}

// Send marshals and sends the CreateNamedQuery API request.
func (r CreateNamedQueryRequest) Send(ctx context.Context) (*CreateNamedQueryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNamedQueryResponse{
		CreateNamedQueryOutput: r.Request.Data.(*CreateNamedQueryOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNamedQueryResponse is the response type for the
// CreateNamedQuery API operation.
type CreateNamedQueryResponse struct {
	*CreateNamedQueryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNamedQuery request.
func (r *CreateNamedQueryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
