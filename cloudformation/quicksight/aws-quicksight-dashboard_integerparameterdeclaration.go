// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_IntegerParameterDeclaration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.IntegerParameterDeclaration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-integerparameterdeclaration.html
type Dashboard_IntegerParameterDeclaration struct {

	// DefaultValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-integerparameterdeclaration.html#cfn-quicksight-dashboard-integerparameterdeclaration-defaultvalues
	DefaultValues *Dashboard_IntegerDefaultValues `json:"DefaultValues,omitempty"`

	// MappedDataSetParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-integerparameterdeclaration.html#cfn-quicksight-dashboard-integerparameterdeclaration-mappeddatasetparameters
	MappedDataSetParameters []Dashboard_MappedDataSetParameter `json:"MappedDataSetParameters,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-integerparameterdeclaration.html#cfn-quicksight-dashboard-integerparameterdeclaration-name
	Name string `json:"Name"`

	// ParameterValueType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-integerparameterdeclaration.html#cfn-quicksight-dashboard-integerparameterdeclaration-parametervaluetype
	ParameterValueType string `json:"ParameterValueType"`

	// ValueWhenUnset AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-integerparameterdeclaration.html#cfn-quicksight-dashboard-integerparameterdeclaration-valuewhenunset
	ValueWhenUnset *Dashboard_IntegerValueWhenUnsetConfiguration `json:"ValueWhenUnset,omitempty"`

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
func (r *Dashboard_IntegerParameterDeclaration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.IntegerParameterDeclaration"
}
