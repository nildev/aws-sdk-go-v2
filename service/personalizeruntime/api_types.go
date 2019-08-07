// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalizeruntime

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// An object that identifies an item.
//
// The and APIs return a list of PredictedItems.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-runtime-2018-05-22/PredictedItem
type PredictedItem struct {
	_ struct{} `type:"structure"`

	// The recommended item ID.
	ItemId *string `json:"personalize-runtime:PredictedItem:ItemId" locationName:"itemId" type:"string"`
}

// String returns the string representation
func (s PredictedItem) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PredictedItem) MarshalFields(e protocol.FieldEncoder) error {
	if s.ItemId != nil {
		v := *s.ItemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "itemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
