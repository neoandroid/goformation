// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_TimeBasedForecastProperties AWS CloudFormation Resource (AWS::QuickSight::Template.TimeBasedForecastProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html
type Template_TimeBasedForecastProperties struct {

	// LowerBoundary AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-lowerboundary
	LowerBoundary *float64 `json:"LowerBoundary,omitempty"`

	// PeriodsBackward AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-periodsbackward
	PeriodsBackward *float64 `json:"PeriodsBackward,omitempty"`

	// PeriodsForward AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-periodsforward
	PeriodsForward *float64 `json:"PeriodsForward,omitempty"`

	// PredictionInterval AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-predictioninterval
	PredictionInterval *float64 `json:"PredictionInterval,omitempty"`

	// Seasonality AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-seasonality
	Seasonality *float64 `json:"Seasonality,omitempty"`

	// UpperBoundary AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-upperboundary
	UpperBoundary *float64 `json:"UpperBoundary,omitempty"`

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
func (r *Template_TimeBasedForecastProperties) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.TimeBasedForecastProperties"
}
