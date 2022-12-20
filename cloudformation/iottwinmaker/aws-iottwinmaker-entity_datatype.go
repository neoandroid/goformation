// Code generated by "go generate". Please don't change this file directly.

package iottwinmaker

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Entity_DataType AWS CloudFormation Resource (AWS::IoTTwinMaker::Entity.DataType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html
type Entity_DataType struct {

	// AllowedValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html#cfn-iottwinmaker-entity-datatype-allowedvalues
	AllowedValues []Entity_DataValue `json:"AllowedValues,omitempty"`

	// NestedType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html#cfn-iottwinmaker-entity-datatype-nestedtype
	NestedType *Entity_DataType `json:"NestedType,omitempty"`

	// Relationship AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html#cfn-iottwinmaker-entity-datatype-relationship
	Relationship *Entity_Relationship `json:"Relationship,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html#cfn-iottwinmaker-entity-datatype-type
	Type *string `json:"Type,omitempty"`

	// UnitOfMeasure AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html#cfn-iottwinmaker-entity-datatype-unitofmeasure
	UnitOfMeasure *string `json:"UnitOfMeasure,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Entity_DataType) AWSCloudFormationType() string {
	return "AWS::IoTTwinMaker::Entity.DataType"
}
