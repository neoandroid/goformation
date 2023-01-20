// Code generated by "go generate". Please don't change this file directly.

package serverless

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Function_HttpApiEvent AWS CloudFormation Resource (AWS::Serverless::Function.HttpApiEvent)
// See: https://github.com/aws/serverless-application-model/blob/master/versions/2016-10-31.md#httpapi
type Function_HttpApiEvent struct {

	// ApiId AWS CloudFormation Property
	// Required: false
	// See: https://github.com/aws/serverless-application-model/blob/master/versions/2016-10-31.md#httpapi
	ApiId *string `json:"ApiId,omitempty"`

	// Auth AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-httpapi.html
	Auth *Function_HttpApiFunctionAuth `json:"Auth,omitempty"`

	// Method AWS CloudFormation Property
	// Required: false
	// See: https://github.com/aws/serverless-application-model/blob/master/versions/2016-10-31.md#httpapi
	Method *string `json:"Method,omitempty"`

	// Path AWS CloudFormation Property
	// Required: false
	// See: https://github.com/aws/serverless-application-model/blob/master/versions/2016-10-31.md#httpapi
	Path *string `json:"Path,omitempty"`

	// PayloadFormatVersion AWS CloudFormation Property
	// Required: false
	// See: https://github.com/aws/serverless-application-model/blob/master/versions/2016-10-31.md#httpapi
	PayloadFormatVersion *string `json:"PayloadFormatVersion,omitempty"`

	// RouteSettings AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-routesettings
	RouteSettings *Function_RouteSettings `json:"RouteSettings,omitempty"`

	// TimeoutInMillis AWS CloudFormation Property
	// Required: false
	// See: https://github.com/aws/serverless-application-model/blob/master/versions/2016-10-31.md#httpapi
	TimeoutInMillis *int `json:"TimeoutInMillis,omitempty"`

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
func (r *Function_HttpApiEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.HttpApiEvent"
}
