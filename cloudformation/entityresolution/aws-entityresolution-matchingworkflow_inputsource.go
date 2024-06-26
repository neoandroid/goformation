// Code generated by "go generate". Please don't change this file directly.

package entityresolution

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// MatchingWorkflow_InputSource AWS CloudFormation Resource (AWS::EntityResolution::MatchingWorkflow.InputSource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html
type MatchingWorkflow_InputSource struct {

	// ApplyNormalization AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-applynormalization
	ApplyNormalization *bool `json:"ApplyNormalization,omitempty"`

	// InputSourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-inputsourcearn
	InputSourceARN string `json:"InputSourceARN"`

	// SchemaArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-inputsource.html#cfn-entityresolution-matchingworkflow-inputsource-schemaarn
	SchemaArn string `json:"SchemaArn"`

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
func (r *MatchingWorkflow_InputSource) AWSCloudFormationType() string {
	return "AWS::EntityResolution::MatchingWorkflow.InputSource"
}
