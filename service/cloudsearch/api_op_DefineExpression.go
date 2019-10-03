// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the DefineExpression operation. Specifies
// the name of the domain you want to update and the expression you want to
// configure.
type DefineExpressionInput struct {
	_ struct{} `type:"structure"`

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start
	// with a letter or number and can contain the following characters: a-z (lowercase),
	// 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`

	// A named expression that can be evaluated at search time. Can be used to sort
	// the search results, define other expressions, or return computed information
	// in the search results.
	//
	// Expression is a required field
	Expression *Expression `type:"structure" required:"true"`
}

// String returns the string representation
func (s DefineExpressionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DefineExpressionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DefineExpressionInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if s.Expression == nil {
		invalidParams.Add(aws.NewErrParamRequired("Expression"))
	}
	if s.Expression != nil {
		if err := s.Expression.Validate(); err != nil {
			invalidParams.AddNested("Expression", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a DefineExpression request. Contains the status of the newly-configured
// expression.
type DefineExpressionOutput struct {
	_ struct{} `type:"structure"`

	// The value of an Expression and its current status.
	//
	// Expression is a required field
	Expression *ExpressionStatus `json:"cloudsearch:DefineExpressionOutput:Expression" type:"structure" required:"true"`
}

// String returns the string representation
func (s DefineExpressionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDefineExpression = "DefineExpression"

// DefineExpressionRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Configures an Expression for the search domain. Used to create new expressions
// and modify existing ones. If the expression exists, the new configuration
// replaces the old one. For more information, see Configuring Expressions (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-expressions.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using DefineExpressionRequest.
//    req := client.DefineExpressionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DefineExpressionRequest(input *DefineExpressionInput) DefineExpressionRequest {
	op := &aws.Operation{
		Name:       opDefineExpression,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DefineExpressionInput{}
	}

	req := c.newRequest(op, input, &DefineExpressionOutput{})
	return DefineExpressionRequest{Request: req, Input: input, Copy: c.DefineExpressionRequest}
}

// DefineExpressionRequest is the request type for the
// DefineExpression API operation.
type DefineExpressionRequest struct {
	*aws.Request
	Input *DefineExpressionInput
	Copy  func(*DefineExpressionInput) DefineExpressionRequest
}

// Send marshals and sends the DefineExpression API request.
func (r DefineExpressionRequest) Send(ctx context.Context) (*DefineExpressionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DefineExpressionResponse{
		DefineExpressionOutput: r.Request.Data.(*DefineExpressionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DefineExpressionResponse is the response type for the
// DefineExpression API operation.
type DefineExpressionResponse struct {
	*DefineExpressionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DefineExpression request.
func (r *DefineExpressionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
