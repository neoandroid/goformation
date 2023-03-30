// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_DefaultNewSheetConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.DefaultNewSheetConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-defaultnewsheetconfiguration.html
type Template_DefaultNewSheetConfiguration struct {

	// InteractiveLayoutConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-defaultnewsheetconfiguration.html#cfn-quicksight-template-defaultnewsheetconfiguration-interactivelayoutconfiguration
	InteractiveLayoutConfiguration *Template_DefaultInteractiveLayoutConfiguration `json:"InteractiveLayoutConfiguration,omitempty"`

	// PaginatedLayoutConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-defaultnewsheetconfiguration.html#cfn-quicksight-template-defaultnewsheetconfiguration-paginatedlayoutconfiguration
	PaginatedLayoutConfiguration *Template_DefaultPaginatedLayoutConfiguration `json:"PaginatedLayoutConfiguration,omitempty"`

	// SheetContentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-defaultnewsheetconfiguration.html#cfn-quicksight-template-defaultnewsheetconfiguration-sheetcontenttype
	SheetContentType *string `json:"SheetContentType,omitempty"`

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
func (r *Template_DefaultNewSheetConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.DefaultNewSheetConfiguration"
}
