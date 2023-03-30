// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_SheetVisualScopingConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.SheetVisualScopingConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetvisualscopingconfiguration.html
type Dashboard_SheetVisualScopingConfiguration struct {

	// Scope AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetvisualscopingconfiguration.html#cfn-quicksight-dashboard-sheetvisualscopingconfiguration-scope
	Scope string `json:"Scope"`

	// SheetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetvisualscopingconfiguration.html#cfn-quicksight-dashboard-sheetvisualscopingconfiguration-sheetid
	SheetId string `json:"SheetId"`

	// VisualIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetvisualscopingconfiguration.html#cfn-quicksight-dashboard-sheetvisualscopingconfiguration-visualids
	VisualIds []string `json:"VisualIds,omitempty"`

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
func (r *Dashboard_SheetVisualScopingConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.SheetVisualScopingConfiguration"
}
