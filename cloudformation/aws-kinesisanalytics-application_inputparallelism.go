package cloudformation

// AWSKinesisAnalyticsApplication_InputParallelism AWS CloudFormation Resource (AWS::KinesisAnalytics::Application.InputParallelism)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputparallelism.html
type AWSKinesisAnalyticsApplication_InputParallelism struct {

	// Count AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputparallelism.html#cfn-kinesisanalytics-application-inputparallelism-count
	Count int `json:"Count,omitempty"`

	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_InputParallelism) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.InputParallelism"
}
