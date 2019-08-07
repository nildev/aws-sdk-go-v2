// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateRouteRequest
type UpdateRouteInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	ApiKeyRequired *bool `locationName:"apiKeyRequired" type:"boolean"`

	// A list of authorization scopes configured on a route. The scopes are used
	// with a COGNITO_USER_POOLS authorizer to authorize the method invocation.
	// The authorization works by matching the route scopes against the scopes parsed
	// from the access token in the incoming request. The method invocation is authorized
	// if any route scope matches a claimed scope in the access token. Otherwise,
	// the invocation is not authorized. When the route scope is configured, the
	// client must provide an access token instead of an identity token for authorization
	// purposes.
	AuthorizationScopes []string `locationName:"authorizationScopes" type:"list"`

	// The authorization type. Valid values are NONE for open access, AWS_IAM for
	// using AWS IAM permissions, and CUSTOM for using a Lambda authorizer.
	AuthorizationType AuthorizationType `locationName:"authorizationType" type:"string" enum:"true"`

	// The identifier.
	AuthorizerId *string `locationName:"authorizerId" type:"string"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ModelSelectionExpression *string `locationName:"modelSelectionExpression" type:"string"`

	// A string with a length between [1-64].
	OperationName *string `locationName:"operationName" type:"string"`

	// The route models.
	RequestModels map[string]string `locationName:"requestModels" type:"map"`

	// The route parameters.
	RequestParameters map[string]ParameterConstraints `locationName:"requestParameters" type:"map"`

	// RouteId is a required field
	RouteId *string `location:"uri" locationName:"routeId" type:"string" required:"true"`

	// After evaulating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	RouteKey *string `locationName:"routeKey" type:"string"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	RouteResponseSelectionExpression *string `locationName:"routeResponseSelectionExpression" type:"string"`

	// A string with a length between [1-128].
	Target *string `locationName:"target" type:"string"`
}

// String returns the string representation
func (s UpdateRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRouteInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.RouteId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateRouteInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApiKeyRequired != nil {
		v := *s.ApiKeyRequired

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeyRequired", protocol.BoolValue(v), metadata)
	}
	if s.AuthorizationScopes != nil {
		v := s.AuthorizationScopes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "authorizationScopes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.AuthorizationType) > 0 {
		v := s.AuthorizationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizationType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AuthorizerId != nil {
		v := *s.AuthorizerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ModelSelectionExpression != nil {
		v := *s.ModelSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "modelSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OperationName != nil {
		v := *s.OperationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "operationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestModels != nil {
		v := s.RequestModels

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestModels", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.RequestParameters != nil {
		v := s.RequestParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.RouteKey != nil {
		v := *s.RouteKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteResponseSelectionExpression != nil {
		v := *s.RouteResponseSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeResponseSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Target != nil {
		v := *s.Target

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "target", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteId != nil {
		v := *s.RouteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "routeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateRouteResult
type UpdateRouteOutput struct {
	_ struct{} `type:"structure"`

	ApiKeyRequired *bool `json:"apigateway:UpdateRouteOutput:ApiKeyRequired" locationName:"apiKeyRequired" type:"boolean"`

	// A list of authorization scopes configured on a route. The scopes are used
	// with a COGNITO_USER_POOLS authorizer to authorize the method invocation.
	// The authorization works by matching the route scopes against the scopes parsed
	// from the access token in the incoming request. The method invocation is authorized
	// if any route scope matches a claimed scope in the access token. Otherwise,
	// the invocation is not authorized. When the route scope is configured, the
	// client must provide an access token instead of an identity token for authorization
	// purposes.
	AuthorizationScopes []string `json:"apigateway:UpdateRouteOutput:AuthorizationScopes" locationName:"authorizationScopes" type:"list"`

	// The authorization type. Valid values are NONE for open access, AWS_IAM for
	// using AWS IAM permissions, and CUSTOM for using a Lambda authorizer.
	AuthorizationType AuthorizationType `json:"apigateway:UpdateRouteOutput:AuthorizationType" locationName:"authorizationType" type:"string" enum:"true"`

	// The identifier.
	AuthorizerId *string `json:"apigateway:UpdateRouteOutput:AuthorizerId" locationName:"authorizerId" type:"string"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ModelSelectionExpression *string `json:"apigateway:UpdateRouteOutput:ModelSelectionExpression" locationName:"modelSelectionExpression" type:"string"`

	// A string with a length between [1-64].
	OperationName *string `json:"apigateway:UpdateRouteOutput:OperationName" locationName:"operationName" type:"string"`

	// The route models.
	RequestModels map[string]string `json:"apigateway:UpdateRouteOutput:RequestModels" locationName:"requestModels" type:"map"`

	// The route parameters.
	RequestParameters map[string]ParameterConstraints `json:"apigateway:UpdateRouteOutput:RequestParameters" locationName:"requestParameters" type:"map"`

	// The identifier.
	RouteId *string `json:"apigateway:UpdateRouteOutput:RouteId" locationName:"routeId" type:"string"`

	// After evaulating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	RouteKey *string `json:"apigateway:UpdateRouteOutput:RouteKey" locationName:"routeKey" type:"string"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	RouteResponseSelectionExpression *string `json:"apigateway:UpdateRouteOutput:RouteResponseSelectionExpression" locationName:"routeResponseSelectionExpression" type:"string"`

	// A string with a length between [1-128].
	Target *string `json:"apigateway:UpdateRouteOutput:Target" locationName:"target" type:"string"`
}

// String returns the string representation
func (s UpdateRouteOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateRouteOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiKeyRequired != nil {
		v := *s.ApiKeyRequired

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeyRequired", protocol.BoolValue(v), metadata)
	}
	if s.AuthorizationScopes != nil {
		v := s.AuthorizationScopes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "authorizationScopes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.AuthorizationType) > 0 {
		v := s.AuthorizationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizationType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AuthorizerId != nil {
		v := *s.AuthorizerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ModelSelectionExpression != nil {
		v := *s.ModelSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "modelSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OperationName != nil {
		v := *s.OperationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "operationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestModels != nil {
		v := s.RequestModels

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestModels", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.RequestParameters != nil {
		v := s.RequestParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.RouteId != nil {
		v := *s.RouteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteKey != nil {
		v := *s.RouteKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteResponseSelectionExpression != nil {
		v := *s.RouteResponseSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeResponseSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Target != nil {
		v := *s.Target

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "target", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateRoute = "UpdateRoute"

// UpdateRouteRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Updates a Route.
//
//    // Example sending a request using UpdateRouteRequest.
//    req := client.UpdateRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateRoute
func (c *Client) UpdateRouteRequest(input *UpdateRouteInput) UpdateRouteRequest {
	op := &aws.Operation{
		Name:       opUpdateRoute,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v2/apis/{apiId}/routes/{routeId}",
	}

	if input == nil {
		input = &UpdateRouteInput{}
	}

	req := c.newRequest(op, input, &UpdateRouteOutput{})
	return UpdateRouteRequest{Request: req, Input: input, Copy: c.UpdateRouteRequest}
}

// UpdateRouteRequest is the request type for the
// UpdateRoute API operation.
type UpdateRouteRequest struct {
	*aws.Request
	Input *UpdateRouteInput
	Copy  func(*UpdateRouteInput) UpdateRouteRequest
}

// Send marshals and sends the UpdateRoute API request.
func (r UpdateRouteRequest) Send(ctx context.Context) (*UpdateRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRouteResponse{
		UpdateRouteOutput: r.Request.Data.(*UpdateRouteOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRouteResponse is the response type for the
// UpdateRoute API operation.
type UpdateRouteResponse struct {
	*UpdateRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRoute request.
func (r *UpdateRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
