// Code generated by "go generate". Please don't change this file directly.

package billingconductor

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// CustomLineItem_CustomLineItemFlatChargeDetails AWS CloudFormation Resource (AWS::BillingConductor::CustomLineItem.CustomLineItemFlatChargeDetails)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemflatchargedetails.html
type CustomLineItem_CustomLineItemFlatChargeDetails struct {

	// ChargeValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-customlineitemflatchargedetails.html#cfn-billingconductor-customlineitem-customlineitemflatchargedetails-chargevalue
	ChargeValue float64 `json:"ChargeValue"`

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
func (r *CustomLineItem_CustomLineItemFlatChargeDetails) AWSCloudFormationType() string {
	return "AWS::BillingConductor::CustomLineItem.CustomLineItemFlatChargeDetails"
}
