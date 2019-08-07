// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeEngineDefaultParameters operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeEngineDefaultParametersMessage
type DescribeEngineDefaultParametersInput struct {
	_ struct{} `type:"structure"`

	// The name of the cache parameter group family.
	//
	// Valid values are: memcached1.4 | memcached1.5 | redis2.6 | redis2.8 | redis3.2
	// | redis4.0 | redis5.0 |
	//
	// CacheParameterGroupFamily is a required field
	CacheParameterGroupFamily *string `type:"string" required:"true"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeEngineDefaultParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEngineDefaultParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEngineDefaultParametersInput"}

	if s.CacheParameterGroupFamily == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheParameterGroupFamily"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeEngineDefaultParametersResult
type DescribeEngineDefaultParametersOutput struct {
	_ struct{} `type:"structure"`

	// Represents the output of a DescribeEngineDefaultParameters operation.
	EngineDefaults *EngineDefaults `json:"elasticache:DescribeEngineDefaultParametersOutput:EngineDefaults" type:"structure"`
}

// String returns the string representation
func (s DescribeEngineDefaultParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEngineDefaultParameters = "DescribeEngineDefaultParameters"

// DescribeEngineDefaultParametersRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns the default engine and system parameter information for the specified
// cache engine.
//
//    // Example sending a request using DescribeEngineDefaultParametersRequest.
//    req := client.DescribeEngineDefaultParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeEngineDefaultParameters
func (c *Client) DescribeEngineDefaultParametersRequest(input *DescribeEngineDefaultParametersInput) DescribeEngineDefaultParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeEngineDefaultParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"EngineDefaults.Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEngineDefaultParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeEngineDefaultParametersOutput{})
	return DescribeEngineDefaultParametersRequest{Request: req, Input: input, Copy: c.DescribeEngineDefaultParametersRequest}
}

// DescribeEngineDefaultParametersRequest is the request type for the
// DescribeEngineDefaultParameters API operation.
type DescribeEngineDefaultParametersRequest struct {
	*aws.Request
	Input *DescribeEngineDefaultParametersInput
	Copy  func(*DescribeEngineDefaultParametersInput) DescribeEngineDefaultParametersRequest
}

// Send marshals and sends the DescribeEngineDefaultParameters API request.
func (r DescribeEngineDefaultParametersRequest) Send(ctx context.Context) (*DescribeEngineDefaultParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEngineDefaultParametersResponse{
		DescribeEngineDefaultParametersOutput: r.Request.Data.(*DescribeEngineDefaultParametersOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEngineDefaultParametersRequestPaginator returns a paginator for DescribeEngineDefaultParameters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEngineDefaultParametersRequest(input)
//   p := elasticache.NewDescribeEngineDefaultParametersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEngineDefaultParametersPaginator(req DescribeEngineDefaultParametersRequest) DescribeEngineDefaultParametersPaginator {
	return DescribeEngineDefaultParametersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEngineDefaultParametersInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeEngineDefaultParametersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEngineDefaultParametersPaginator struct {
	aws.Pager
}

func (p *DescribeEngineDefaultParametersPaginator) CurrentPage() *DescribeEngineDefaultParametersOutput {
	return p.Pager.CurrentPage().(*DescribeEngineDefaultParametersOutput)
}

// DescribeEngineDefaultParametersResponse is the response type for the
// DescribeEngineDefaultParameters API operation.
type DescribeEngineDefaultParametersResponse struct {
	*DescribeEngineDefaultParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEngineDefaultParameters request.
func (r *DescribeEngineDefaultParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
