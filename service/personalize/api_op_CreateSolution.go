// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/CreateSolutionRequest
type CreateSolutionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset group that provides the training
	// data.
	//
	// DatasetGroupArn is a required field
	DatasetGroupArn *string `locationName:"datasetGroupArn" type:"string" required:"true"`

	// When your have multiple event types (using an EVENT_TYPE schema field), this
	// parameter specifies which event type (for example, 'click' or 'like') is
	// used for training the model.
	EventType *string `locationName:"eventType" type:"string"`

	// The name for the solution.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// Whether to perform automated machine learning (AutoML). The default is false.
	// For this case, you must specify recipeArn.
	//
	// When set to true, Amazon Personalize analyzes your training data and selects
	// the optimal USER_PERSONALIZATION recipe and hyperparameters. In this case,
	// you must omit recipeArn. Amazon Personalize determines the optimal recipe
	// by running tests with different values for the hyperparameters. AutoML lengthens
	// the training process as compared to selecting a specific recipe.
	PerformAutoML *bool `locationName:"performAutoML" type:"boolean"`

	// Whether to perform hyperparameter optimization (HPO) on the specified or
	// selected recipe. The default is false.
	//
	// When performing AutoML, this parameter is always true and you should not
	// set it to false.
	PerformHPO *bool `locationName:"performHPO" type:"boolean"`

	// The ARN of the recipe to use for model training. Only specified when performAutoML
	// is false.
	RecipeArn *string `locationName:"recipeArn" type:"string"`

	// The configuration to use with the solution. When performAutoML is set to
	// true, Amazon Personalize only evaluates the autoMLConfig section of the solution
	// configuration.
	SolutionConfig *SolutionConfig `locationName:"solutionConfig" type:"structure"`
}

// String returns the string representation
func (s CreateSolutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSolutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSolutionInput"}

	if s.DatasetGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetGroupArn"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.SolutionConfig != nil {
		if err := s.SolutionConfig.Validate(); err != nil {
			invalidParams.AddNested("SolutionConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/CreateSolutionResponse
type CreateSolutionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the solution.
	SolutionArn *string `locationName:"solutionArn" type:"string"`
}

// String returns the string representation
func (s CreateSolutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSolution = "CreateSolution"

// CreateSolutionRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Creates the configuration for training a model. A trained model is known
// as a solution. After the configuration is created, you train the model (create
// a solution) by calling the CreateSolutionVersion operation. Every time you
// call CreateSolutionVersion, a new version of the solution is created.
//
// After creating a solution version, you check its accuracy by calling GetSolutionMetrics.
// When you are satisfied with the version, you deploy it using CreateCampaign.
// The campaign provides recommendations to a client through the GetRecommendations
// (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html)
// API.
//
// To train a model, Amazon Personalize requires training data and a recipe.
// The training data comes from the dataset group that you provide in the request.
// A recipe specifies the training algorithm and a feature transformation. You
// can specify one of the predefined recipes provided by Amazon Personalize.
// Alternatively, you can specify performAutoML and Amazon Personalize will
// analyze your data and select the optimum USER_PERSONALIZATION recipe for
// you.
//
// Status
//
// A solution can be in one of the following states:
//
//    * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
//    * DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the solution, call DescribeSolution. Wait until the
// status shows as ACTIVE before calling CreateSolutionVersion.
//
// Related APIs
//
//    * ListSolutions
//
//    * CreateSolutionVersion
//
//    * DescribeSolution
//
//    * DeleteSolution
//
//    * ListSolutionVersions
//
//    * DescribeSolutionVersion
//
//    // Example sending a request using CreateSolutionRequest.
//    req := client.CreateSolutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/CreateSolution
func (c *Client) CreateSolutionRequest(input *CreateSolutionInput) CreateSolutionRequest {
	op := &aws.Operation{
		Name:       opCreateSolution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSolutionInput{}
	}

	req := c.newRequest(op, input, &CreateSolutionOutput{})
	return CreateSolutionRequest{Request: req, Input: input, Copy: c.CreateSolutionRequest}
}

// CreateSolutionRequest is the request type for the
// CreateSolution API operation.
type CreateSolutionRequest struct {
	*aws.Request
	Input *CreateSolutionInput
	Copy  func(*CreateSolutionInput) CreateSolutionRequest
}

// Send marshals and sends the CreateSolution API request.
func (r CreateSolutionRequest) Send(ctx context.Context) (*CreateSolutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSolutionResponse{
		CreateSolutionOutput: r.Request.Data.(*CreateSolutionOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSolutionResponse is the response type for the
// CreateSolution API operation.
type CreateSolutionResponse struct {
	*CreateSolutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSolution request.
func (r *CreateSolutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
