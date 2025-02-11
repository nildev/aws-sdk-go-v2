// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to get an invalidation's information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetInvalidationRequest
type GetInvalidationInput struct {
	_ struct{} `type:"structure"`

	// The distribution's ID.
	//
	// DistributionId is a required field
	DistributionId *string `location:"uri" locationName:"DistributionId" type:"string" required:"true"`

	// The identifier for the invalidation request, for example, IDFDVBD632BHDS5.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetInvalidationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInvalidationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInvalidationInput"}

	if s.DistributionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DistributionId"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetInvalidationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DistributionId != nil {
		v := *s.DistributionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DistributionId", protocol.StringValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// The returned result of the corresponding request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetInvalidationResult
type GetInvalidationOutput struct {
	_ struct{} `type:"structure" payload:"Invalidation"`

	// The invalidation's information. For more information, see Invalidation Complex
	// Type (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/InvalidationDatatype.html).
	Invalidation *Invalidation `type:"structure"`
}

// String returns the string representation
func (s GetInvalidationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetInvalidationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Invalidation != nil {
		v := s.Invalidation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "Invalidation", v, metadata)
	}
	return nil
}

const opGetInvalidation = "GetInvalidation2019_03_26"

// GetInvalidationRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Get the information about an invalidation.
//
//    // Example sending a request using GetInvalidationRequest.
//    req := client.GetInvalidationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/GetInvalidation
func (c *Client) GetInvalidationRequest(input *GetInvalidationInput) GetInvalidationRequest {
	op := &aws.Operation{
		Name:       opGetInvalidation,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/distribution/{DistributionId}/invalidation/{Id}",
	}

	if input == nil {
		input = &GetInvalidationInput{}
	}

	req := c.newRequest(op, input, &GetInvalidationOutput{})
	return GetInvalidationRequest{Request: req, Input: input, Copy: c.GetInvalidationRequest}
}

// GetInvalidationRequest is the request type for the
// GetInvalidation API operation.
type GetInvalidationRequest struct {
	*aws.Request
	Input *GetInvalidationInput
	Copy  func(*GetInvalidationInput) GetInvalidationRequest
}

// Send marshals and sends the GetInvalidation API request.
func (r GetInvalidationRequest) Send(ctx context.Context) (*GetInvalidationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInvalidationResponse{
		GetInvalidationOutput: r.Request.Data.(*GetInvalidationOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInvalidationResponse is the response type for the
// GetInvalidation API operation.
type GetInvalidationResponse struct {
	*GetInvalidationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInvalidation request.
func (r *GetInvalidationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
