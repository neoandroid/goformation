// Code generated by "go generate". Please don't change this file directly.

package pcaconnectorad

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_ExtensionsV2 AWS CloudFormation Resource (AWS::PCAConnectorAD::Template.ExtensionsV2)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-extensionsv2.html
type Template_ExtensionsV2 struct {

	// ApplicationPolicies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-extensionsv2.html#cfn-pcaconnectorad-template-extensionsv2-applicationpolicies
	ApplicationPolicies *Template_ApplicationPolicies `json:"ApplicationPolicies,omitempty"`

	// KeyUsage AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-extensionsv2.html#cfn-pcaconnectorad-template-extensionsv2-keyusage
	KeyUsage *Template_KeyUsage `json:"KeyUsage"`

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
func (r *Template_ExtensionsV2) AWSCloudFormationType() string {
	return "AWS::PCAConnectorAD::Template.ExtensionsV2"
}
