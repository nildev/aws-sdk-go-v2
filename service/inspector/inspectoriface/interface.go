// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package inspectoriface provides an interface to enable mocking the Amazon Inspector service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package inspectoriface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
)

// InspectorAPI provides an interface to enable mocking the
// inspector.Inspector service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Inspector.
//    func myFunc(svc inspectoriface.InspectorAPI) bool {
//        // Make svc.AddAttributesToFindings request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := inspector.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockInspectorClient struct {
//        inspectoriface.InspectorAPI
//    }
//    func (m *mockInspectorClient) AddAttributesToFindings(input *inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockInspectorClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type InspectorAPI interface {
	AddAttributesToFindingsRequest(*inspector.AddAttributesToFindingsInput) inspector.AddAttributesToFindingsRequest

	CreateAssessmentTargetRequest(*inspector.CreateAssessmentTargetInput) inspector.CreateAssessmentTargetRequest

	CreateAssessmentTemplateRequest(*inspector.CreateAssessmentTemplateInput) inspector.CreateAssessmentTemplateRequest

	CreateResourceGroupRequest(*inspector.CreateResourceGroupInput) inspector.CreateResourceGroupRequest

	DeleteAssessmentRunRequest(*inspector.DeleteAssessmentRunInput) inspector.DeleteAssessmentRunRequest

	DeleteAssessmentTargetRequest(*inspector.DeleteAssessmentTargetInput) inspector.DeleteAssessmentTargetRequest

	DeleteAssessmentTemplateRequest(*inspector.DeleteAssessmentTemplateInput) inspector.DeleteAssessmentTemplateRequest

	DescribeAssessmentRunsRequest(*inspector.DescribeAssessmentRunsInput) inspector.DescribeAssessmentRunsRequest

	DescribeAssessmentTargetsRequest(*inspector.DescribeAssessmentTargetsInput) inspector.DescribeAssessmentTargetsRequest

	DescribeAssessmentTemplatesRequest(*inspector.DescribeAssessmentTemplatesInput) inspector.DescribeAssessmentTemplatesRequest

	DescribeCrossAccountAccessRoleRequest(*inspector.DescribeCrossAccountAccessRoleInput) inspector.DescribeCrossAccountAccessRoleRequest

	DescribeFindingsRequest(*inspector.DescribeFindingsInput) inspector.DescribeFindingsRequest

	DescribeResourceGroupsRequest(*inspector.DescribeResourceGroupsInput) inspector.DescribeResourceGroupsRequest

	DescribeRulesPackagesRequest(*inspector.DescribeRulesPackagesInput) inspector.DescribeRulesPackagesRequest

	GetAssessmentReportRequest(*inspector.GetAssessmentReportInput) inspector.GetAssessmentReportRequest

	GetTelemetryMetadataRequest(*inspector.GetTelemetryMetadataInput) inspector.GetTelemetryMetadataRequest

	ListAssessmentRunAgentsRequest(*inspector.ListAssessmentRunAgentsInput) inspector.ListAssessmentRunAgentsRequest

	ListAssessmentRunAgentsPages(*inspector.ListAssessmentRunAgentsInput, func(*inspector.ListAssessmentRunAgentsOutput, bool) bool) error
	ListAssessmentRunAgentsPagesWithContext(aws.Context, *inspector.ListAssessmentRunAgentsInput, func(*inspector.ListAssessmentRunAgentsOutput, bool) bool, ...aws.Option) error

	ListAssessmentRunsRequest(*inspector.ListAssessmentRunsInput) inspector.ListAssessmentRunsRequest

	ListAssessmentRunsPages(*inspector.ListAssessmentRunsInput, func(*inspector.ListAssessmentRunsOutput, bool) bool) error
	ListAssessmentRunsPagesWithContext(aws.Context, *inspector.ListAssessmentRunsInput, func(*inspector.ListAssessmentRunsOutput, bool) bool, ...aws.Option) error

	ListAssessmentTargetsRequest(*inspector.ListAssessmentTargetsInput) inspector.ListAssessmentTargetsRequest

	ListAssessmentTargetsPages(*inspector.ListAssessmentTargetsInput, func(*inspector.ListAssessmentTargetsOutput, bool) bool) error
	ListAssessmentTargetsPagesWithContext(aws.Context, *inspector.ListAssessmentTargetsInput, func(*inspector.ListAssessmentTargetsOutput, bool) bool, ...aws.Option) error

	ListAssessmentTemplatesRequest(*inspector.ListAssessmentTemplatesInput) inspector.ListAssessmentTemplatesRequest

	ListAssessmentTemplatesPages(*inspector.ListAssessmentTemplatesInput, func(*inspector.ListAssessmentTemplatesOutput, bool) bool) error
	ListAssessmentTemplatesPagesWithContext(aws.Context, *inspector.ListAssessmentTemplatesInput, func(*inspector.ListAssessmentTemplatesOutput, bool) bool, ...aws.Option) error

	ListEventSubscriptionsRequest(*inspector.ListEventSubscriptionsInput) inspector.ListEventSubscriptionsRequest

	ListEventSubscriptionsPages(*inspector.ListEventSubscriptionsInput, func(*inspector.ListEventSubscriptionsOutput, bool) bool) error
	ListEventSubscriptionsPagesWithContext(aws.Context, *inspector.ListEventSubscriptionsInput, func(*inspector.ListEventSubscriptionsOutput, bool) bool, ...aws.Option) error

	ListFindingsRequest(*inspector.ListFindingsInput) inspector.ListFindingsRequest

	ListFindingsPages(*inspector.ListFindingsInput, func(*inspector.ListFindingsOutput, bool) bool) error
	ListFindingsPagesWithContext(aws.Context, *inspector.ListFindingsInput, func(*inspector.ListFindingsOutput, bool) bool, ...aws.Option) error

	ListRulesPackagesRequest(*inspector.ListRulesPackagesInput) inspector.ListRulesPackagesRequest

	ListRulesPackagesPages(*inspector.ListRulesPackagesInput, func(*inspector.ListRulesPackagesOutput, bool) bool) error
	ListRulesPackagesPagesWithContext(aws.Context, *inspector.ListRulesPackagesInput, func(*inspector.ListRulesPackagesOutput, bool) bool, ...aws.Option) error

	ListTagsForResourceRequest(*inspector.ListTagsForResourceInput) inspector.ListTagsForResourceRequest

	PreviewAgentsRequest(*inspector.PreviewAgentsInput) inspector.PreviewAgentsRequest

	PreviewAgentsPages(*inspector.PreviewAgentsInput, func(*inspector.PreviewAgentsOutput, bool) bool) error
	PreviewAgentsPagesWithContext(aws.Context, *inspector.PreviewAgentsInput, func(*inspector.PreviewAgentsOutput, bool) bool, ...aws.Option) error

	RegisterCrossAccountAccessRoleRequest(*inspector.RegisterCrossAccountAccessRoleInput) inspector.RegisterCrossAccountAccessRoleRequest

	RemoveAttributesFromFindingsRequest(*inspector.RemoveAttributesFromFindingsInput) inspector.RemoveAttributesFromFindingsRequest

	SetTagsForResourceRequest(*inspector.SetTagsForResourceInput) inspector.SetTagsForResourceRequest

	StartAssessmentRunRequest(*inspector.StartAssessmentRunInput) inspector.StartAssessmentRunRequest

	StopAssessmentRunRequest(*inspector.StopAssessmentRunInput) inspector.StopAssessmentRunRequest

	SubscribeToEventRequest(*inspector.SubscribeToEventInput) inspector.SubscribeToEventRequest

	UnsubscribeFromEventRequest(*inspector.UnsubscribeFromEventInput) inspector.UnsubscribeFromEventRequest

	UpdateAssessmentTargetRequest(*inspector.UpdateAssessmentTargetInput) inspector.UpdateAssessmentTargetRequest
}

var _ InspectorAPI = (*inspector.Inspector)(nil)
