// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request for meta data about a dataset (creation date, number of records,
// size) by owner and dataset name.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/DescribeDatasetRequest
type DescribeDatasetInput struct {
	_ struct{} `type:"structure"`

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	//
	// DatasetName is a required field
	DatasetName *string `location:"uri" locationName:"DatasetName" min:"1" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// IdentityId is a required field
	IdentityId *string `location:"uri" locationName:"IdentityId" min:"1" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `location:"uri" locationName:"IdentityPoolId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDatasetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDatasetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDatasetInput"}

	if s.DatasetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetName"))
	}
	if s.DatasetName != nil && len(*s.DatasetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetName", 1))
	}

	if s.IdentityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityId"))
	}
	if s.IdentityId != nil && len(*s.IdentityId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityId", 1))
	}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDatasetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DatasetName != nil {
		v := *s.DatasetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DatasetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityId != nil {
		v := *s.IdentityId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityPoolId != nil {
		v := *s.IdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Response to a successful DescribeDataset request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/DescribeDatasetResponse
type DescribeDatasetOutput struct {
	_ struct{} `type:"structure"`

	// Meta data for a collection of data for an identity. An identity can have
	// multiple datasets. A dataset can be general or associated with a particular
	// entity in an application (like a saved game). Datasets are automatically
	// created if they don't exist. Data is synced by dataset, and a dataset can
	// hold up to 1MB of key-value pairs.
	Dataset *Dataset `json:"cognito-sync:DescribeDatasetOutput:Dataset" type:"structure"`
}

// String returns the string representation
func (s DescribeDatasetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDatasetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Dataset != nil {
		v := s.Dataset

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Dataset", v, metadata)
	}
	return nil
}

const opDescribeDataset = "DescribeDataset"

// DescribeDatasetRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Gets meta data about a dataset by identity and dataset name. With Amazon
// Cognito Sync, each identity has access only to its own data. Thus, the credentials
// used to make this API call need to have access to the identity data.
//
// This API can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials. You should use Cognito Identity credentials
// to make this API call.
//
//    // Example sending a request using DescribeDatasetRequest.
//    req := client.DescribeDatasetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/DescribeDataset
func (c *Client) DescribeDatasetRequest(input *DescribeDatasetInput) DescribeDatasetRequest {
	op := &aws.Operation{
		Name:       opDescribeDataset,
		HTTPMethod: "GET",
		HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}",
	}

	if input == nil {
		input = &DescribeDatasetInput{}
	}

	req := c.newRequest(op, input, &DescribeDatasetOutput{})
	return DescribeDatasetRequest{Request: req, Input: input, Copy: c.DescribeDatasetRequest}
}

// DescribeDatasetRequest is the request type for the
// DescribeDataset API operation.
type DescribeDatasetRequest struct {
	*aws.Request
	Input *DescribeDatasetInput
	Copy  func(*DescribeDatasetInput) DescribeDatasetRequest
}

// Send marshals and sends the DescribeDataset API request.
func (r DescribeDatasetRequest) Send(ctx context.Context) (*DescribeDatasetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDatasetResponse{
		DescribeDatasetOutput: r.Request.Data.(*DescribeDatasetOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDatasetResponse is the response type for the
// DescribeDataset API operation.
type DescribeDatasetResponse struct {
	*DescribeDatasetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDataset request.
func (r *DescribeDatasetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
