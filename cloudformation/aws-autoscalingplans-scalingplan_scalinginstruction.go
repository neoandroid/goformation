package cloudformation

// AWSAutoScalingPlansScalingPlan_ScalingInstruction AWS CloudFormation Resource (AWS::AutoScalingPlans::ScalingPlan.ScalingInstruction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html
type AWSAutoScalingPlansScalingPlan_ScalingInstruction struct {

	// MaxCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-maxcapacity
	MaxCapacity int `json:"MaxCapacity"`

	// MinCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-mincapacity
	MinCapacity int `json:"MinCapacity"`

	// ResourceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-resourceid
	ResourceId string `json:"ResourceId,omitempty"`

	// ScalableDimension AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-scalabledimension
	ScalableDimension string `json:"ScalableDimension,omitempty"`

	// ServiceNamespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-servicenamespace
	ServiceNamespace string `json:"ServiceNamespace,omitempty"`

	// TargetTrackingConfigurations AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-targettrackingconfigurations
	TargetTrackingConfigurations []AWSAutoScalingPlansScalingPlan_TargetTrackingConfiguration `json:"TargetTrackingConfigurations,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingPlansScalingPlan_ScalingInstruction) AWSCloudFormationType() string {
	return "AWS::AutoScalingPlans::ScalingPlan.ScalingInstruction"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAutoScalingPlansScalingPlan_ScalingInstruction) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAutoScalingPlansScalingPlan_ScalingInstruction) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAutoScalingPlansScalingPlan_ScalingInstruction) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAutoScalingPlansScalingPlan_ScalingInstruction) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAutoScalingPlansScalingPlan_ScalingInstruction) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
