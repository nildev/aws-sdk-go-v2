// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codebuildiface provides an interface to enable mocking the AWS CodeBuild service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codebuildiface

import (
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

// CodeBuildAPI provides an interface to enable mocking the
// codebuild.CodeBuild service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CodeBuild.
//    func myFunc(svc codebuildiface.CodeBuildAPI) bool {
//        // Make svc.BatchDeleteBuilds request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := codebuild.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodeBuildClient struct {
//        codebuildiface.CodeBuildAPI
//    }
//    func (m *mockCodeBuildClient) BatchDeleteBuilds(input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodeBuildClient{}
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
type CodeBuildAPI interface {
	BatchDeleteBuildsRequest(*codebuild.BatchDeleteBuildsInput) codebuild.BatchDeleteBuildsRequest

	BatchGetBuildsRequest(*codebuild.BatchGetBuildsInput) codebuild.BatchGetBuildsRequest

	BatchGetProjectsRequest(*codebuild.BatchGetProjectsInput) codebuild.BatchGetProjectsRequest

	CreateProjectRequest(*codebuild.CreateProjectInput) codebuild.CreateProjectRequest

	CreateWebhookRequest(*codebuild.CreateWebhookInput) codebuild.CreateWebhookRequest

	DeleteProjectRequest(*codebuild.DeleteProjectInput) codebuild.DeleteProjectRequest

	DeleteWebhookRequest(*codebuild.DeleteWebhookInput) codebuild.DeleteWebhookRequest

	ListBuildsRequest(*codebuild.ListBuildsInput) codebuild.ListBuildsRequest

	ListBuildsForProjectRequest(*codebuild.ListBuildsForProjectInput) codebuild.ListBuildsForProjectRequest

	ListCuratedEnvironmentImagesRequest(*codebuild.ListCuratedEnvironmentImagesInput) codebuild.ListCuratedEnvironmentImagesRequest

	ListProjectsRequest(*codebuild.ListProjectsInput) codebuild.ListProjectsRequest

	StartBuildRequest(*codebuild.StartBuildInput) codebuild.StartBuildRequest

	StopBuildRequest(*codebuild.StopBuildInput) codebuild.StopBuildRequest

	UpdateProjectRequest(*codebuild.UpdateProjectInput) codebuild.UpdateProjectRequest
}

var _ CodeBuildAPI = (*codebuild.CodeBuild)(nil)
