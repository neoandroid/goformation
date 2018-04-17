package cloudformation

// AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig AWS CloudFormation Resource (AWS::EMR::InstanceFleetConfig.EbsBlockDeviceConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig.html
type AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig struct {

	// VolumeSpecification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig.html#cfn-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig-volumespecification
	VolumeSpecification *AWSEMRInstanceFleetConfig_VolumeSpecification `json:"VolumeSpecification,omitempty"`

	// VolumesPerInstance AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig.html#cfn-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig-volumesperinstance
	VolumesPerInstance int `json:"VolumesPerInstance,omitempty"`

	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig.EbsBlockDeviceConfig"
}
