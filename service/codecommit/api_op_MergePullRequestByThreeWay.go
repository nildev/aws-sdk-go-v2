// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/MergePullRequestByThreeWayInput
type MergePullRequestByThreeWayInput struct {
	_ struct{} `type:"structure"`

	// The name of the author who created the commit. This information will be used
	// as both the author and committer for the commit.
	AuthorName *string `locationName:"authorName" type:"string"`

	// The commit message to include in the commit information for the merge.
	CommitMessage *string `locationName:"commitMessage" type:"string"`

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL
	// is used, which will return a not mergeable result if the same file has differences
	// in both branches. If LINE_LEVEL is specified, a conflict will be considered
	// not mergeable if the same file in both branches has differences on the same
	// line.
	ConflictDetailLevel ConflictDetailLevelTypeEnum `locationName:"conflictDetailLevel" type:"string" enum:"true"`

	// A list of inputs to use when resolving conflicts during a merge if AUTOMERGE
	// is chosen as the conflict resolution strategy.
	ConflictResolution *ConflictResolution `locationName:"conflictResolution" type:"structure"`

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation
	// will be successful.
	ConflictResolutionStrategy ConflictResolutionStrategyTypeEnum `locationName:"conflictResolutionStrategy" type:"string" enum:"true"`

	// The email address of the person merging the branches. This information will
	// be used in the commit information for the merge.
	Email *string `locationName:"email" type:"string"`

	// If the commit contains deletions, whether to keep a folder or folder structure
	// if the changes leave the folders empty. If this is specified as true, a .gitkeep
	// file will be created for empty folders. The default is false.
	KeepEmptyFolders *bool `locationName:"keepEmptyFolders" type:"boolean"`

	// The system-generated ID of the pull request. To get this ID, use ListPullRequests.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The name of the repository where the pull request was created.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`

	// The full commit ID of the original or updated commit in the pull request
	// source branch. Pass this value if you want an exception thrown if the current
	// commit ID of the tip of the source branch does not match this commit ID.
	SourceCommitId *string `locationName:"sourceCommitId" type:"string"`
}

// String returns the string representation
func (s MergePullRequestByThreeWayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MergePullRequestByThreeWayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MergePullRequestByThreeWayInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}
	if s.ConflictResolution != nil {
		if err := s.ConflictResolution.Validate(); err != nil {
			invalidParams.AddNested("ConflictResolution", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/MergePullRequestByThreeWayOutput
type MergePullRequestByThreeWayOutput struct {
	_ struct{} `type:"structure"`

	// Returns information about a pull request.
	PullRequest *PullRequest `json:"codecommit:MergePullRequestByThreeWayOutput:PullRequest" locationName:"pullRequest" type:"structure"`
}

// String returns the string representation
func (s MergePullRequestByThreeWayOutput) String() string {
	return awsutil.Prettify(s)
}

const opMergePullRequestByThreeWay = "MergePullRequestByThreeWay"

// MergePullRequestByThreeWayRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Attempts to merge the source commit of a pull request into the specified
// destination branch for that pull request at the specified commit using the
// three-way merge strategy. If the merge is successful, it closes the pull
// request.
//
//    // Example sending a request using MergePullRequestByThreeWayRequest.
//    req := client.MergePullRequestByThreeWayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/MergePullRequestByThreeWay
func (c *Client) MergePullRequestByThreeWayRequest(input *MergePullRequestByThreeWayInput) MergePullRequestByThreeWayRequest {
	op := &aws.Operation{
		Name:       opMergePullRequestByThreeWay,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MergePullRequestByThreeWayInput{}
	}

	req := c.newRequest(op, input, &MergePullRequestByThreeWayOutput{})
	return MergePullRequestByThreeWayRequest{Request: req, Input: input, Copy: c.MergePullRequestByThreeWayRequest}
}

// MergePullRequestByThreeWayRequest is the request type for the
// MergePullRequestByThreeWay API operation.
type MergePullRequestByThreeWayRequest struct {
	*aws.Request
	Input *MergePullRequestByThreeWayInput
	Copy  func(*MergePullRequestByThreeWayInput) MergePullRequestByThreeWayRequest
}

// Send marshals and sends the MergePullRequestByThreeWay API request.
func (r MergePullRequestByThreeWayRequest) Send(ctx context.Context) (*MergePullRequestByThreeWayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &MergePullRequestByThreeWayResponse{
		MergePullRequestByThreeWayOutput: r.Request.Data.(*MergePullRequestByThreeWayOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// MergePullRequestByThreeWayResponse is the response type for the
// MergePullRequestByThreeWay API operation.
type MergePullRequestByThreeWayResponse struct {
	*MergePullRequestByThreeWayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// MergePullRequestByThreeWay request.
func (r *MergePullRequestByThreeWayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
