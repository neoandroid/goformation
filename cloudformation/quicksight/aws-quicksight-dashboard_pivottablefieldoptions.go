// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_PivotTableFieldOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.PivotTableFieldOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoptions.html
type Dashboard_PivotTableFieldOptions struct {

	// CollapseStateOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoptions.html#cfn-quicksight-dashboard-pivottablefieldoptions-collapsestateoptions
	CollapseStateOptions []Dashboard_PivotTableFieldCollapseStateOption `json:"CollapseStateOptions,omitempty"`

	// DataPathOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoptions.html#cfn-quicksight-dashboard-pivottablefieldoptions-datapathoptions
	DataPathOptions []Dashboard_PivotTableDataPathOption `json:"DataPathOptions,omitempty"`

	// SelectedFieldOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldoptions.html#cfn-quicksight-dashboard-pivottablefieldoptions-selectedfieldoptions
	SelectedFieldOptions []Dashboard_PivotTableFieldOption `json:"SelectedFieldOptions,omitempty"`

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
func (r *Dashboard_PivotTableFieldOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.PivotTableFieldOptions"
}
