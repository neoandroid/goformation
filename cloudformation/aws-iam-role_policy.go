package cloudformation

// AWSIAMRole_Policy AWS CloudFormation Resource (AWS::IAM::Role.Policy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type AWSIAMRole_Policy struct {

	// PolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`

	// PolicyName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policyname
	PolicyName string `json:"PolicyName,omitempty"`

	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMRole_Policy) AWSCloudFormationType() string {
	return "AWS::IAM::Role.Policy"
}
