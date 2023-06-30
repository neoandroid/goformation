// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_PivotTableFieldOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.PivotTableFieldOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldoptions.html
type Analysis_PivotTableFieldOptions struct {

	// CollapseStateOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldoptions.html#cfn-quicksight-analysis-pivottablefieldoptions-collapsestateoptions
	CollapseStateOptions []Analysis_PivotTableFieldCollapseStateOption `json:"CollapseStateOptions,omitempty"`

	// DataPathOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldoptions.html#cfn-quicksight-analysis-pivottablefieldoptions-datapathoptions
	DataPathOptions []Analysis_PivotTableDataPathOption `json:"DataPathOptions,omitempty"`

	// SelectedFieldOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldoptions.html#cfn-quicksight-analysis-pivottablefieldoptions-selectedfieldoptions
	SelectedFieldOptions []Analysis_PivotTableFieldOption `json:"SelectedFieldOptions,omitempty"`

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
func (r *Analysis_PivotTableFieldOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.PivotTableFieldOptions"
}
