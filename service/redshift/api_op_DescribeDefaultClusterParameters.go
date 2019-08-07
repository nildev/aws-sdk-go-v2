// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeDefaultClusterParametersMessage
type DescribeDefaultClusterParametersInput struct {
	_ struct{} `type:"structure"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeDefaultClusterParameters
	// request exceed the value specified in MaxRecords, AWS returns a value in
	// the Marker field of the response. You can retrieve the next set of response
	// records by providing the returned marker value in the Marker parameter and
	// retrying the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The name of the cluster parameter group family.
	//
	// ParameterGroupFamily is a required field
	ParameterGroupFamily *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDefaultClusterParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDefaultClusterParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDefaultClusterParametersInput"}

	if s.ParameterGroupFamily == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterGroupFamily"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeDefaultClusterParametersResult
type DescribeDefaultClusterParametersOutput struct {
	_ struct{} `type:"structure"`

	// Describes the default cluster parameters for a parameter group family.
	DefaultClusterParameters *DefaultClusterParameters `json:"redshift:DescribeDefaultClusterParametersOutput:DefaultClusterParameters" type:"structure"`
}

// String returns the string representation
func (s DescribeDefaultClusterParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDefaultClusterParameters = "DescribeDefaultClusterParameters"

// DescribeDefaultClusterParametersRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns a list of parameter settings for the specified parameter group family.
//
// For more information about parameters and parameter groups, go to Amazon
// Redshift Parameter Groups (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using DescribeDefaultClusterParametersRequest.
//    req := client.DescribeDefaultClusterParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeDefaultClusterParameters
func (c *Client) DescribeDefaultClusterParametersRequest(input *DescribeDefaultClusterParametersInput) DescribeDefaultClusterParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeDefaultClusterParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"DefaultClusterParameters.Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDefaultClusterParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeDefaultClusterParametersOutput{})
	return DescribeDefaultClusterParametersRequest{Request: req, Input: input, Copy: c.DescribeDefaultClusterParametersRequest}
}

// DescribeDefaultClusterParametersRequest is the request type for the
// DescribeDefaultClusterParameters API operation.
type DescribeDefaultClusterParametersRequest struct {
	*aws.Request
	Input *DescribeDefaultClusterParametersInput
	Copy  func(*DescribeDefaultClusterParametersInput) DescribeDefaultClusterParametersRequest
}

// Send marshals and sends the DescribeDefaultClusterParameters API request.
func (r DescribeDefaultClusterParametersRequest) Send(ctx context.Context) (*DescribeDefaultClusterParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDefaultClusterParametersResponse{
		DescribeDefaultClusterParametersOutput: r.Request.Data.(*DescribeDefaultClusterParametersOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDefaultClusterParametersRequestPaginator returns a paginator for DescribeDefaultClusterParameters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDefaultClusterParametersRequest(input)
//   p := redshift.NewDescribeDefaultClusterParametersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDefaultClusterParametersPaginator(req DescribeDefaultClusterParametersRequest) DescribeDefaultClusterParametersPaginator {
	return DescribeDefaultClusterParametersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDefaultClusterParametersInput
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

// DescribeDefaultClusterParametersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDefaultClusterParametersPaginator struct {
	aws.Pager
}

func (p *DescribeDefaultClusterParametersPaginator) CurrentPage() *DescribeDefaultClusterParametersOutput {
	return p.Pager.CurrentPage().(*DescribeDefaultClusterParametersOutput)
}

// DescribeDefaultClusterParametersResponse is the response type for the
// DescribeDefaultClusterParameters API operation.
type DescribeDefaultClusterParametersResponse struct {
	*DescribeDefaultClusterParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDefaultClusterParameters request.
func (r *DescribeDefaultClusterParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
