// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// This input determines how the ListClusters action filters the list of clusters
// that it returns.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/ListClustersInput
type ListClustersInput struct {
	_ struct{} `type:"structure"`

	// The cluster state filters to apply when listing clusters.
	ClusterStates []ClusterState `type:"list"`

	// The creation date and time beginning value filter for listing clusters.
	CreatedAfter *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The creation date and time end value filter for listing clusters.
	CreatedBefore *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListClustersInput) String() string {
	return awsutil.Prettify(s)
}

// This contains a ClusterSummaryList with the cluster details; for example,
// the cluster IDs, names, and status.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/ListClustersOutput
type ListClustersOutput struct {
	_ struct{} `type:"structure"`

	// The list of clusters for the account based on the given filters.
	Clusters []ClusterSummary `json:"elasticmapreduce:ListClustersOutput:Clusters" type:"list"`

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string `json:"elasticmapreduce:ListClustersOutput:Marker" type:"string"`
}

// String returns the string representation
func (s ListClustersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListClusters = "ListClusters"

// ListClustersRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Provides the status of all clusters visible to this AWS account. Allows you
// to filter the list of clusters based on certain criteria; for example, filtering
// by cluster creation date and time or by status. This call returns a maximum
// of 50 clusters per call, but returns a marker to track the paging of the
// cluster list across multiple ListClusters calls.
//
//    // Example sending a request using ListClustersRequest.
//    req := client.ListClustersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/ListClusters
func (c *Client) ListClustersRequest(input *ListClustersInput) ListClustersRequest {
	op := &aws.Operation{
		Name:       opListClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListClustersInput{}
	}

	req := c.newRequest(op, input, &ListClustersOutput{})
	return ListClustersRequest{Request: req, Input: input, Copy: c.ListClustersRequest}
}

// ListClustersRequest is the request type for the
// ListClusters API operation.
type ListClustersRequest struct {
	*aws.Request
	Input *ListClustersInput
	Copy  func(*ListClustersInput) ListClustersRequest
}

// Send marshals and sends the ListClusters API request.
func (r ListClustersRequest) Send(ctx context.Context) (*ListClustersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListClustersResponse{
		ListClustersOutput: r.Request.Data.(*ListClustersOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListClustersRequestPaginator returns a paginator for ListClusters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListClustersRequest(input)
//   p := emr.NewListClustersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListClustersPaginator(req ListClustersRequest) ListClustersPaginator {
	return ListClustersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListClustersInput
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

// ListClustersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListClustersPaginator struct {
	aws.Pager
}

func (p *ListClustersPaginator) CurrentPage() *ListClustersOutput {
	return p.Pager.CurrentPage().(*ListClustersOutput)
}

// ListClustersResponse is the response type for the
// ListClusters API operation.
type ListClustersResponse struct {
	*ListClustersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListClusters request.
func (r *ListClustersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
