// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribePortfolioShareStatusInput
type DescribePortfolioShareStatusInput struct {
	_ struct{} `type:"structure"`

	// The token for the portfolio share operation. This token is returned either
	// by CreatePortfolioShare or by DeletePortfolioShare.
	//
	// PortfolioShareToken is a required field
	PortfolioShareToken *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribePortfolioShareStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePortfolioShareStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePortfolioShareStatusInput"}

	if s.PortfolioShareToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioShareToken"))
	}
	if s.PortfolioShareToken != nil && len(*s.PortfolioShareToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioShareToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribePortfolioShareStatusOutput
type DescribePortfolioShareStatusOutput struct {
	_ struct{} `type:"structure"`

	// Organization node identifier. It can be either account id, organizational
	// unit id or organization id.
	OrganizationNodeValue *string `json:"servicecatalog:DescribePortfolioShareStatusOutput:OrganizationNodeValue" type:"string"`

	// The portfolio identifier.
	PortfolioId *string `json:"servicecatalog:DescribePortfolioShareStatusOutput:PortfolioId" min:"1" type:"string"`

	// The token for the portfolio share operation. For example, share-6v24abcdefghi.
	PortfolioShareToken *string `json:"servicecatalog:DescribePortfolioShareStatusOutput:PortfolioShareToken" min:"1" type:"string"`

	// Information about the portfolio share operation.
	ShareDetails *ShareDetails `json:"servicecatalog:DescribePortfolioShareStatusOutput:ShareDetails" type:"structure"`

	// Status of the portfolio share operation.
	Status ShareStatus `json:"servicecatalog:DescribePortfolioShareStatusOutput:Status" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribePortfolioShareStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePortfolioShareStatus = "DescribePortfolioShareStatus"

// DescribePortfolioShareStatusRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets the status of the specified portfolio share operation. This API can
// only be called by the master account in the organization.
//
//    // Example sending a request using DescribePortfolioShareStatusRequest.
//    req := client.DescribePortfolioShareStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribePortfolioShareStatus
func (c *Client) DescribePortfolioShareStatusRequest(input *DescribePortfolioShareStatusInput) DescribePortfolioShareStatusRequest {
	op := &aws.Operation{
		Name:       opDescribePortfolioShareStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePortfolioShareStatusInput{}
	}

	req := c.newRequest(op, input, &DescribePortfolioShareStatusOutput{})
	return DescribePortfolioShareStatusRequest{Request: req, Input: input, Copy: c.DescribePortfolioShareStatusRequest}
}

// DescribePortfolioShareStatusRequest is the request type for the
// DescribePortfolioShareStatus API operation.
type DescribePortfolioShareStatusRequest struct {
	*aws.Request
	Input *DescribePortfolioShareStatusInput
	Copy  func(*DescribePortfolioShareStatusInput) DescribePortfolioShareStatusRequest
}

// Send marshals and sends the DescribePortfolioShareStatus API request.
func (r DescribePortfolioShareStatusRequest) Send(ctx context.Context) (*DescribePortfolioShareStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePortfolioShareStatusResponse{
		DescribePortfolioShareStatusOutput: r.Request.Data.(*DescribePortfolioShareStatusOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePortfolioShareStatusResponse is the response type for the
// DescribePortfolioShareStatus API operation.
type DescribePortfolioShareStatusResponse struct {
	*DescribePortfolioShareStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePortfolioShareStatus request.
func (r *DescribePortfolioShareStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
