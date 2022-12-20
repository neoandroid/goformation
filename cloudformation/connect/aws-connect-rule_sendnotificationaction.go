// Code generated by "go generate". Please don't change this file directly.

package connect

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Rule_SendNotificationAction AWS CloudFormation Resource (AWS::Connect::Rule.SendNotificationAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html
type Rule_SendNotificationAction struct {

	// Content AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html#cfn-connect-rule-sendnotificationaction-content
	Content string `json:"Content"`

	// ContentType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html#cfn-connect-rule-sendnotificationaction-contenttype
	ContentType string `json:"ContentType"`

	// DeliveryMethod AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html#cfn-connect-rule-sendnotificationaction-deliverymethod
	DeliveryMethod string `json:"DeliveryMethod"`

	// Recipient AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html#cfn-connect-rule-sendnotificationaction-recipient
	Recipient *Rule_NotificationRecipientType `json:"Recipient"`

	// Subject AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html#cfn-connect-rule-sendnotificationaction-subject
	Subject *string `json:"Subject,omitempty"`

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
func (r *Rule_SendNotificationAction) AWSCloudFormationType() string {
	return "AWS::Connect::Rule.SendNotificationAction"
}
