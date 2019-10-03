// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package simpledb

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

type Attribute struct {
	_ struct{} `type:"structure"`

	AlternateNameEncoding *string `json:"sdb:Attribute:AlternateNameEncoding" type:"string"`

	AlternateValueEncoding *string `json:"sdb:Attribute:AlternateValueEncoding" type:"string"`

	// The name of the attribute.
	//
	// Name is a required field
	Name *string `json:"sdb:Attribute:Name" type:"string" required:"true"`

	// The value of the attribute.
	//
	// Value is a required field
	Value *string `json:"sdb:Attribute:Value" type:"string" required:"true"`
}

// String returns the string representation
func (s Attribute) String() string {
	return awsutil.Prettify(s)
}

type DeletableAttribute struct {
	_ struct{} `type:"structure"`

	// The name of the attribute.
	//
	// Name is a required field
	Name *string `json:"sdb:DeletableAttribute:Name" type:"string" required:"true"`

	// The value of the attribute.
	Value *string `json:"sdb:DeletableAttribute:Value" type:"string"`
}

// String returns the string representation
func (s DeletableAttribute) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletableAttribute) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletableAttribute"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletableItem struct {
	_ struct{} `type:"structure"`

	Attributes []DeletableAttribute `json:"sdb:DeletableItem:Attributes" locationNameList:"Attribute" type:"list" flattened:"true"`

	// Name is a required field
	Name *string `json:"sdb:DeletableItem:Name" locationName:"ItemName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletableItem) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletableItem) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletableItem"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type Item struct {
	_ struct{} `type:"structure"`

	AlternateNameEncoding *string `json:"sdb:Item:AlternateNameEncoding" type:"string"`

	// A list of attributes.
	//
	// Attributes is a required field
	Attributes []Attribute `json:"sdb:Item:Attributes" locationNameList:"Attribute" type:"list" flattened:"true" required:"true"`

	// The name of the item.
	//
	// Name is a required field
	Name *string `json:"sdb:Item:Name" type:"string" required:"true"`
}

// String returns the string representation
func (s Item) String() string {
	return awsutil.Prettify(s)
}

type ReplaceableAttribute struct {
	_ struct{} `type:"structure"`

	// The name of the replaceable attribute.
	//
	// Name is a required field
	Name *string `json:"sdb:ReplaceableAttribute:Name" type:"string" required:"true"`

	// A flag specifying whether or not to replace the attribute/value pair or to
	// add a new attribute/value pair. The default setting is
	//    false
	// .
	Replace *bool `json:"sdb:ReplaceableAttribute:Replace" type:"boolean"`

	// The value of the replaceable attribute.
	//
	// Value is a required field
	Value *string `json:"sdb:ReplaceableAttribute:Value" type:"string" required:"true"`
}

// String returns the string representation
func (s ReplaceableAttribute) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReplaceableAttribute) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReplaceableAttribute"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ReplaceableItem struct {
	_ struct{} `type:"structure"`

	// The list of attributes for a replaceable item.
	//
	// Attributes is a required field
	Attributes []ReplaceableAttribute `json:"sdb:ReplaceableItem:Attributes" locationNameList:"Attribute" type:"list" flattened:"true" required:"true"`

	// The name of the replaceable item.
	//
	// Name is a required field
	Name *string `json:"sdb:ReplaceableItem:Name" locationName:"ItemName" type:"string" required:"true"`
}

// String returns the string representation
func (s ReplaceableItem) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReplaceableItem) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReplaceableItem"}

	if s.Attributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attributes"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies the conditions under which data should be updated. If an update
// condition is specified for a request, the data will only be updated if the
// condition is satisfied. For example, if an attribute with a specific name
// and value exists, or if a specific attribute doesn't exist.
type UpdateCondition struct {
	_ struct{} `type:"structure"`

	// A value specifying whether or not the specified attribute must exist with
	// the specified value in order for the update condition to be satisfied. Specify
	// true if the attribute must exist for the update condition to be satisfied.
	// Specify false if the attribute should not exist in order for the update condition
	// to be satisfied.
	Exists *bool `json:"sdb:UpdateCondition:Exists" type:"boolean"`

	// The name of the attribute involved in the condition.
	Name *string `json:"sdb:UpdateCondition:Name" type:"string"`

	// The value of an attribute. This value can only be specified when the Exists
	// parameter is equal to true.
	Value *string `json:"sdb:UpdateCondition:Value" type:"string"`
}

// String returns the string representation
func (s UpdateCondition) String() string {
	return awsutil.Prettify(s)
}
