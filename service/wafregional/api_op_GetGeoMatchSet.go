// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetGeoMatchSetRequest
type GetGeoMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The GeoMatchSetId of the GeoMatchSet that you want to get. GeoMatchSetId
	// is returned by CreateGeoMatchSet and by ListGeoMatchSets.
	//
	// GeoMatchSetId is a required field
	GeoMatchSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetGeoMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGeoMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGeoMatchSetInput"}

	if s.GeoMatchSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GeoMatchSetId"))
	}
	if s.GeoMatchSetId != nil && len(*s.GeoMatchSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GeoMatchSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetGeoMatchSetResponse
type GetGeoMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the GeoMatchSet that you specified in the GetGeoMatchSet
	// request. This includes the Type, which for a GeoMatchContraint is always
	// Country, as well as the Value, which is the identifier for a specific country.
	GeoMatchSet *waf.GeoMatchSet `json:"waf-regional:GetGeoMatchSetOutput:GeoMatchSet" type:"structure"`
}

// String returns the string representation
func (s GetGeoMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetGeoMatchSet = "GetGeoMatchSet"

// GetGeoMatchSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the GeoMatchSet that is specified by GeoMatchSetId.
//
//    // Example sending a request using GetGeoMatchSetRequest.
//    req := client.GetGeoMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetGeoMatchSet
func (c *Client) GetGeoMatchSetRequest(input *GetGeoMatchSetInput) GetGeoMatchSetRequest {
	op := &aws.Operation{
		Name:       opGetGeoMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetGeoMatchSetInput{}
	}

	req := c.newRequest(op, input, &GetGeoMatchSetOutput{})
	return GetGeoMatchSetRequest{Request: req, Input: input, Copy: c.GetGeoMatchSetRequest}
}

// GetGeoMatchSetRequest is the request type for the
// GetGeoMatchSet API operation.
type GetGeoMatchSetRequest struct {
	*aws.Request
	Input *GetGeoMatchSetInput
	Copy  func(*GetGeoMatchSetInput) GetGeoMatchSetRequest
}

// Send marshals and sends the GetGeoMatchSet API request.
func (r GetGeoMatchSetRequest) Send(ctx context.Context) (*GetGeoMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGeoMatchSetResponse{
		GetGeoMatchSetOutput: r.Request.Data.(*GetGeoMatchSetOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetGeoMatchSetResponse is the response type for the
// GetGeoMatchSet API operation.
type GetGeoMatchSetResponse struct {
	*GetGeoMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGeoMatchSet request.
func (r *GetGeoMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
