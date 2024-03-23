// Code generated by "go generate". Please don't change this file directly.

package amplify

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Domain_Certificate AWS CloudFormation Resource (AWS::Amplify::Domain.Certificate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificate.html
type Domain_Certificate struct {

	// CertificateArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificate.html#cfn-amplify-domain-certificate-certificatearn
	CertificateArn *string `json:"CertificateArn,omitempty"`

	// CertificateType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificate.html#cfn-amplify-domain-certificate-certificatetype
	CertificateType *string `json:"CertificateType,omitempty"`

	// CertificateVerificationDNSRecord AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificate.html#cfn-amplify-domain-certificate-certificateverificationdnsrecord
	CertificateVerificationDNSRecord *string `json:"CertificateVerificationDNSRecord,omitempty"`

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
func (r *Domain_Certificate) AWSCloudFormationType() string {
	return "AWS::Amplify::Domain.Certificate"
}