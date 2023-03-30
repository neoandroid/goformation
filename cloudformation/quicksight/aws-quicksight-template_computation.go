// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_Computation AWS CloudFormation Resource (AWS::QuickSight::Template.Computation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html
type Template_Computation struct {

	// Forecast AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html#cfn-quicksight-template-computation-forecast
	Forecast *Template_ForecastComputation `json:"Forecast,omitempty"`

	// GrowthRate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html#cfn-quicksight-template-computation-growthrate
	GrowthRate *Template_GrowthRateComputation `json:"GrowthRate,omitempty"`

	// MaximumMinimum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html#cfn-quicksight-template-computation-maximumminimum
	MaximumMinimum *Template_MaximumMinimumComputation `json:"MaximumMinimum,omitempty"`

	// MetricComparison AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html#cfn-quicksight-template-computation-metriccomparison
	MetricComparison *Template_MetricComparisonComputation `json:"MetricComparison,omitempty"`

	// PeriodOverPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html#cfn-quicksight-template-computation-periodoverperiod
	PeriodOverPeriod *Template_PeriodOverPeriodComputation `json:"PeriodOverPeriod,omitempty"`

	// PeriodToDate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html#cfn-quicksight-template-computation-periodtodate
	PeriodToDate *Template_PeriodToDateComputation `json:"PeriodToDate,omitempty"`

	// TopBottomMovers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html#cfn-quicksight-template-computation-topbottommovers
	TopBottomMovers *Template_TopBottomMoversComputation `json:"TopBottomMovers,omitempty"`

	// TopBottomRanked AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html#cfn-quicksight-template-computation-topbottomranked
	TopBottomRanked *Template_TopBottomRankedComputation `json:"TopBottomRanked,omitempty"`

	// TotalAggregation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html#cfn-quicksight-template-computation-totalaggregation
	TotalAggregation *Template_TotalAggregationComputation `json:"TotalAggregation,omitempty"`

	// UniqueValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-computation.html#cfn-quicksight-template-computation-uniquevalues
	UniqueValues *Template_UniqueValuesComputation `json:"UniqueValues,omitempty"`

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
func (r *Template_Computation) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.Computation"
}
