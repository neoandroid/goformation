package cloudformation

// AWSServerlessFunction_VpcConfig AWS CloudFormation Resource (AWS::Serverless::Function.VpcConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html
type AWSServerlessFunction_VpcConfig struct {

	// SecurityGroupIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html
	SubnetIds []string `json:"SubnetIds,omitempty"`

	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_VpcConfig) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.VpcConfig"
}
