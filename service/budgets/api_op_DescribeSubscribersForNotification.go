// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request of DescribeSubscribersForNotification
type DescribeSubscribersForNotificationInput struct {
	_ struct{} `type:"structure"`

	// The accountId that is associated with the budget whose subscribers you want
	// descriptions of.
	//
	// AccountId is a required field
	AccountId *string `min:"12" type:"string" required:"true"`

	// The name of the budget whose subscribers you want descriptions of.
	//
	// BudgetName is a required field
	BudgetName *string `min:"1" type:"string" required:"true"`

	// An optional integer that represents how many entries a paginated response
	// contains. The maximum is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token that you include in your request to indicate the next
	// set of results that you want to retrieve.
	NextToken *string `type:"string"`

	// The notification whose subscribers you want to list.
	//
	// Notification is a required field
	Notification *Notification `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeSubscribersForNotificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSubscribersForNotificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSubscribersForNotificationInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.AccountId != nil && len(*s.AccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 12))
	}

	if s.BudgetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BudgetName"))
	}
	if s.BudgetName != nil && len(*s.BudgetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BudgetName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Notification == nil {
		invalidParams.Add(aws.NewErrParamRequired("Notification"))
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			invalidParams.AddNested("Notification", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response of DescribeSubscribersForNotification
type DescribeSubscribersForNotificationOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token in the service response that indicates the next set
	// of results that you can retrieve.
	NextToken *string `json:"budgets:DescribeSubscribersForNotificationOutput:NextToken" type:"string"`

	// A list of subscribers that are associated with a notification.
	Subscribers []Subscriber `json:"budgets:DescribeSubscribersForNotificationOutput:Subscribers" min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeSubscribersForNotificationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSubscribersForNotification = "DescribeSubscribersForNotification"

// DescribeSubscribersForNotificationRequest returns a request value for making API operation for
// AWS Budgets.
//
// Lists the subscribers that are associated with a notification.
//
//    // Example sending a request using DescribeSubscribersForNotificationRequest.
//    req := client.DescribeSubscribersForNotificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeSubscribersForNotificationRequest(input *DescribeSubscribersForNotificationInput) DescribeSubscribersForNotificationRequest {
	op := &aws.Operation{
		Name:       opDescribeSubscribersForNotification,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSubscribersForNotificationInput{}
	}

	req := c.newRequest(op, input, &DescribeSubscribersForNotificationOutput{})
	return DescribeSubscribersForNotificationRequest{Request: req, Input: input, Copy: c.DescribeSubscribersForNotificationRequest}
}

// DescribeSubscribersForNotificationRequest is the request type for the
// DescribeSubscribersForNotification API operation.
type DescribeSubscribersForNotificationRequest struct {
	*aws.Request
	Input *DescribeSubscribersForNotificationInput
	Copy  func(*DescribeSubscribersForNotificationInput) DescribeSubscribersForNotificationRequest
}

// Send marshals and sends the DescribeSubscribersForNotification API request.
func (r DescribeSubscribersForNotificationRequest) Send(ctx context.Context) (*DescribeSubscribersForNotificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSubscribersForNotificationResponse{
		DescribeSubscribersForNotificationOutput: r.Request.Data.(*DescribeSubscribersForNotificationOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSubscribersForNotificationResponse is the response type for the
// DescribeSubscribersForNotification API operation.
type DescribeSubscribersForNotificationResponse struct {
	*DescribeSubscribersForNotificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSubscribersForNotification request.
func (r *DescribeSubscribersForNotificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
