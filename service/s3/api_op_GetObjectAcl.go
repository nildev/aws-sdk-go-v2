// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectAclRequest
type GetObjectAclInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// VersionId used to reference a specific version of the object.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s GetObjectAclInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectAclInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectAclInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetObjectAclInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectAclInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectAclOutput
type GetObjectAclOutput struct {
	_ struct{} `type:"structure"`

	// A list of grants.
	Grants []Grant `json:"s3:GetObjectAclOutput:Grants" locationName:"AccessControlList" locationNameList:"Grant" type:"list"`

	Owner *Owner `json:"s3:GetObjectAclOutput:Owner" type:"structure"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `json:"s3:GetObjectAclOutput:RequestCharged" location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetObjectAclOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectAclOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Grants != nil {
		v := s.Grants

		metadata := protocol.Metadata{ListLocationName: "Grant"}
		ls0 := e.List(protocol.BodyTarget, "AccessControlList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Owner != nil {
		v := s.Owner

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Owner", v, metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	return nil
}

const opGetObjectAcl = "GetObjectAcl"

// GetObjectAclRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the access control list (ACL) of an object.
//
//    // Example sending a request using GetObjectAclRequest.
//    req := client.GetObjectAclRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectAcl
func (c *Client) GetObjectAclRequest(input *GetObjectAclInput) GetObjectAclRequest {
	op := &aws.Operation{
		Name:       opGetObjectAcl,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}?acl",
	}

	if input == nil {
		input = &GetObjectAclInput{}
	}

	req := c.newRequest(op, input, &GetObjectAclOutput{})
	return GetObjectAclRequest{Request: req, Input: input, Copy: c.GetObjectAclRequest}
}

// GetObjectAclRequest is the request type for the
// GetObjectAcl API operation.
type GetObjectAclRequest struct {
	*aws.Request
	Input *GetObjectAclInput
	Copy  func(*GetObjectAclInput) GetObjectAclRequest
}

// Send marshals and sends the GetObjectAcl API request.
func (r GetObjectAclRequest) Send(ctx context.Context) (*GetObjectAclResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectAclResponse{
		GetObjectAclOutput: r.Request.Data.(*GetObjectAclOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectAclResponse is the response type for the
// GetObjectAcl API operation.
type GetObjectAclResponse struct {
	*GetObjectAclOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObjectAcl request.
func (r *GetObjectAclResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
