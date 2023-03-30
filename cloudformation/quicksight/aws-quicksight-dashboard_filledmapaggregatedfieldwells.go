// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_FilledMapAggregatedFieldWells AWS CloudFormation Resource (AWS::QuickSight::Dashboard.FilledMapAggregatedFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapaggregatedfieldwells.html
type Dashboard_FilledMapAggregatedFieldWells struct {

	// Geospatial AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapaggregatedfieldwells.html#cfn-quicksight-dashboard-filledmapaggregatedfieldwells-geospatial
	Geospatial []Dashboard_DimensionField `json:"Geospatial,omitempty"`

	// Values AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapaggregatedfieldwells.html#cfn-quicksight-dashboard-filledmapaggregatedfieldwells-values
	Values []Dashboard_MeasureField `json:"Values,omitempty"`

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
func (r *Dashboard_FilledMapAggregatedFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.FilledMapAggregatedFieldWells"
}
