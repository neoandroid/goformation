package cloudformation

// AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination AWS CloudFormation Resource (AWS::SES::ConfigurationSetEventDestination.KinesisFirehoseDestination)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-kinesisfirehosedestination.html
type AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination struct {

	// DeliveryStreamARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-kinesisfirehosedestination.html#cfn-ses-configurationseteventdestination-kinesisfirehosedestination-deliverystreamarn
	DeliveryStreamARN string `json:"DeliveryStreamARN,omitempty"`

	// IAMRoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-kinesisfirehosedestination.html#cfn-ses-configurationseteventdestination-kinesisfirehosedestination-iamrolearn
	IAMRoleARN string `json:"IAMRoleARN,omitempty"`

	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination) AWSCloudFormationType() string {
	return "AWS::SES::ConfigurationSetEventDestination.KinesisFirehoseDestination"
}
