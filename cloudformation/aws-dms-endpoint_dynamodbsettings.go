package cloudformation

// AWSDMSEndpoint_DynamoDbSettings AWS CloudFormation Resource (AWS::DMS::Endpoint.DynamoDbSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-dynamodbsettings.html
type AWSDMSEndpoint_DynamoDbSettings struct {

	// ServiceAccessRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-dynamodbsettings.html#cfn-dms-endpoint-dynamodbsettings-serviceaccessrolearn
	ServiceAccessRoleArn string `json:"ServiceAccessRoleArn,omitempty"`

	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSEndpoint_DynamoDbSettings) AWSCloudFormationType() string {
	return "AWS::DMS::Endpoint.DynamoDbSettings"
}
