// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodbstreams

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetShardIterator operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/GetShardIteratorInput
type GetShardIteratorInput struct {
	_ struct{} `type:"structure"`

	// The sequence number of a stream record in the shard from which to start reading.
	SequenceNumber *string `min:"21" type:"string"`

	// The identifier of the shard. The iterator will be returned for this shard
	// ID.
	//
	// ShardId is a required field
	ShardId *string `min:"28" type:"string" required:"true"`

	// Determines how the shard iterator is used to start reading stream records
	// from the shard:
	//
	//    * AT_SEQUENCE_NUMBER - Start reading exactly from the position denoted
	//    by a specific sequence number.
	//
	//    * AFTER_SEQUENCE_NUMBER - Start reading right after the position denoted
	//    by a specific sequence number.
	//
	//    * TRIM_HORIZON - Start reading at the last (untrimmed) stream record,
	//    which is the oldest record in the shard. In DynamoDB Streams, there is
	//    a 24 hour limit on data retention. Stream records whose age exceeds this
	//    limit are subject to removal (trimming) from the stream.
	//
	//    * LATEST - Start reading just after the most recent stream record in the
	//    shard, so that you always read the most recent data in the shard.
	//
	// ShardIteratorType is a required field
	ShardIteratorType ShardIteratorType `type:"string" required:"true" enum:"true"`

	// The Amazon Resource Name (ARN) for the stream.
	//
	// StreamArn is a required field
	StreamArn *string `min:"37" type:"string" required:"true"`
}

// String returns the string representation
func (s GetShardIteratorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetShardIteratorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetShardIteratorInput"}
	if s.SequenceNumber != nil && len(*s.SequenceNumber) < 21 {
		invalidParams.Add(aws.NewErrParamMinLen("SequenceNumber", 21))
	}

	if s.ShardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShardId"))
	}
	if s.ShardId != nil && len(*s.ShardId) < 28 {
		invalidParams.Add(aws.NewErrParamMinLen("ShardId", 28))
	}
	if len(s.ShardIteratorType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ShardIteratorType"))
	}

	if s.StreamArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamArn"))
	}
	if s.StreamArn != nil && len(*s.StreamArn) < 37 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamArn", 37))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetShardIterator operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/GetShardIteratorOutput
type GetShardIteratorOutput struct {
	_ struct{} `type:"structure"`

	// The position in the shard from which to start reading stream records sequentially.
	// A shard iterator specifies this position using the sequence number of a stream
	// record in a shard.
	ShardIterator *string `json:"streams.dynamodb:GetShardIteratorOutput:ShardIterator" min:"1" type:"string"`
}

// String returns the string representation
func (s GetShardIteratorOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetShardIterator = "GetShardIterator"

// GetShardIteratorRequest returns a request value for making API operation for
// Amazon DynamoDB Streams.
//
// Returns a shard iterator. A shard iterator provides information about how
// to retrieve the stream records from within a shard. Use the shard iterator
// in a subsequent GetRecords request to read the stream records from the shard.
//
// A shard iterator expires 15 minutes after it is returned to the requester.
//
//    // Example sending a request using GetShardIteratorRequest.
//    req := client.GetShardIteratorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/GetShardIterator
func (c *Client) GetShardIteratorRequest(input *GetShardIteratorInput) GetShardIteratorRequest {
	op := &aws.Operation{
		Name:       opGetShardIterator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetShardIteratorInput{}
	}

	req := c.newRequest(op, input, &GetShardIteratorOutput{})
	return GetShardIteratorRequest{Request: req, Input: input, Copy: c.GetShardIteratorRequest}
}

// GetShardIteratorRequest is the request type for the
// GetShardIterator API operation.
type GetShardIteratorRequest struct {
	*aws.Request
	Input *GetShardIteratorInput
	Copy  func(*GetShardIteratorInput) GetShardIteratorRequest
}

// Send marshals and sends the GetShardIterator API request.
func (r GetShardIteratorRequest) Send(ctx context.Context) (*GetShardIteratorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetShardIteratorResponse{
		GetShardIteratorOutput: r.Request.Data.(*GetShardIteratorOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetShardIteratorResponse is the response type for the
// GetShardIterator API operation.
type GetShardIteratorResponse struct {
	*GetShardIteratorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetShardIterator request.
func (r *GetShardIteratorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
