// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_DataLabelType AWS CloudFormation Resource (AWS::QuickSight::Dashboard.DataLabelType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html
type Dashboard_DataLabelType struct {

	// DataPathLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html#cfn-quicksight-dashboard-datalabeltype-datapathlabeltype
	DataPathLabelType *Dashboard_DataPathLabelType `json:"DataPathLabelType,omitempty"`

	// FieldLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html#cfn-quicksight-dashboard-datalabeltype-fieldlabeltype
	FieldLabelType *Dashboard_FieldLabelType `json:"FieldLabelType,omitempty"`

	// MaximumLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html#cfn-quicksight-dashboard-datalabeltype-maximumlabeltype
	MaximumLabelType *Dashboard_MaximumLabelType `json:"MaximumLabelType,omitempty"`

	// MinimumLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html#cfn-quicksight-dashboard-datalabeltype-minimumlabeltype
	MinimumLabelType *Dashboard_MinimumLabelType `json:"MinimumLabelType,omitempty"`

	// RangeEndsLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeltype.html#cfn-quicksight-dashboard-datalabeltype-rangeendslabeltype
	RangeEndsLabelType *Dashboard_RangeEndsLabelType `json:"RangeEndsLabelType,omitempty"`

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
func (r *Dashboard_DataLabelType) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.DataLabelType"
}
