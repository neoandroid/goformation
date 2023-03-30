// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_NumericRangeFilter AWS CloudFormation Resource (AWS::QuickSight::Analysis.NumericRangeFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefilter.html
type Analysis_NumericRangeFilter struct {

	// AggregationFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefilter.html#cfn-quicksight-analysis-numericrangefilter-aggregationfunction
	AggregationFunction *Analysis_AggregationFunction `json:"AggregationFunction,omitempty"`

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefilter.html#cfn-quicksight-analysis-numericrangefilter-column
	Column *Analysis_ColumnIdentifier `json:"Column"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefilter.html#cfn-quicksight-analysis-numericrangefilter-filterid
	FilterId string `json:"FilterId"`

	// IncludeMaximum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefilter.html#cfn-quicksight-analysis-numericrangefilter-includemaximum
	IncludeMaximum *bool `json:"IncludeMaximum,omitempty"`

	// IncludeMinimum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefilter.html#cfn-quicksight-analysis-numericrangefilter-includeminimum
	IncludeMinimum *bool `json:"IncludeMinimum,omitempty"`

	// NullOption AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefilter.html#cfn-quicksight-analysis-numericrangefilter-nulloption
	NullOption string `json:"NullOption"`

	// RangeMaximum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefilter.html#cfn-quicksight-analysis-numericrangefilter-rangemaximum
	RangeMaximum *Analysis_NumericRangeFilterValue `json:"RangeMaximum,omitempty"`

	// RangeMinimum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefilter.html#cfn-quicksight-analysis-numericrangefilter-rangeminimum
	RangeMinimum *Analysis_NumericRangeFilterValue `json:"RangeMinimum,omitempty"`

	// SelectAllOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-numericrangefilter.html#cfn-quicksight-analysis-numericrangefilter-selectalloptions
	SelectAllOptions *string `json:"SelectAllOptions,omitempty"`

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
func (r *Analysis_NumericRangeFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.NumericRangeFilter"
}
