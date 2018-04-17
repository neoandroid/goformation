package cloudformation

// AWSEMRInstanceGroupConfig_ScalingTrigger AWS CloudFormation Resource (AWS::EMR::InstanceGroupConfig.ScalingTrigger)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html
type AWSEMRInstanceGroupConfig_ScalingTrigger struct {

	// CloudWatchAlarmDefinition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html#cfn-elasticmapreduce-instancegroupconfig-scalingtrigger-cloudwatchalarmdefinition
	CloudWatchAlarmDefinition *AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition `json:"CloudWatchAlarmDefinition,omitempty"`

	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_ScalingTrigger) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.ScalingTrigger"
}
