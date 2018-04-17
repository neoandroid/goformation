package cloudformation

// AWSEC2Instance_NoDevice AWS CloudFormation Resource (AWS::EC2::Instance.NoDevice)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-nodevice.html
type AWSEC2Instance_NoDevice struct {
	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_NoDevice) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.NoDevice"
}
