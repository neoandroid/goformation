// Code generated by "go generate". Please don't change this file directly.

package ses

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// EmailIdentity_ConfigurationSetAttributes AWS CloudFormation Resource (AWS::SES::EmailIdentity.ConfigurationSetAttributes)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-emailidentity-configurationsetattributes.html
type EmailIdentity_ConfigurationSetAttributes struct {

	// ConfigurationSetName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-emailidentity-configurationsetattributes.html#cfn-ses-emailidentity-configurationsetattributes-configurationsetname
	ConfigurationSetName *string `json:"ConfigurationSetName,omitempty"`

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
func (r *EmailIdentity_ConfigurationSetAttributes) AWSCloudFormationType() string {
	return "AWS::SES::EmailIdentity.ConfigurationSetAttributes"
}