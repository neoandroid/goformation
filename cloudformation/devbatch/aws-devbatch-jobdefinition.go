// Code generated by "go generate". Please don't change this file directly.

package devbatch

import (
	"bytes"
	"encoding/json"

	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobDefinition AWS CloudFormation Resource (AWS::DevBatch::JobDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html
type JobDefinition struct {

	// ContainerProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-containerproperties
	ContainerProperties *JobDefinition_ContainerProperties `json:"ContainerProperties,omitempty"`

	// EksProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-eksproperties
	EksProperties *JobDefinition_EksProperties `json:"EksProperties,omitempty"`

	// JobDefinitionName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-jobdefinitionname
	JobDefinitionName *string `json:"JobDefinitionName,omitempty"`

	// NodeProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-nodeproperties
	NodeProperties *JobDefinition_NodeProperties `json:"NodeProperties,omitempty"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-parameters
	Parameters interface{} `json:"Parameters,omitempty"`

	// PlatformCapabilities AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-platformcapabilities
	PlatformCapabilities []string `json:"PlatformCapabilities,omitempty"`

	// PropagateTags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-propagatetags
	PropagateTags *bool `json:"PropagateTags,omitempty"`

	// RetryStrategy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-retrystrategy
	RetryStrategy *JobDefinition_RetryStrategy `json:"RetryStrategy,omitempty"`

	// SchedulingPriority AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-schedulingpriority
	SchedulingPriority *int `json:"SchedulingPriority,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-tags
	Tags interface{} `json:"Tags,omitempty"`

	// Timeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-timeout
	Timeout *JobDefinition_Timeout `json:"Timeout,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devbatch-jobdefinition.html#cfn-devbatch-jobdefinition-type
	Type string `json:"Type"`

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
func (r *JobDefinition) AWSCloudFormationType() string {
	return "AWS::DevBatch::JobDefinition"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r JobDefinition) MarshalJSON() ([]byte, error) {
	type Properties JobDefinition
	return json.Marshal(&struct {
		Type                string
		Properties          Properties
		DependsOn           []string                     `json:"DependsOn,omitempty"`
		Metadata            map[string]interface{}       `json:"Metadata,omitempty"`
		DeletionPolicy      policies.DeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy policies.UpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Condition           string                       `json:"Condition,omitempty"`
	}{
		Type:                r.AWSCloudFormationType(),
		Properties:          (Properties)(r),
		DependsOn:           r.AWSCloudFormationDependsOn,
		Metadata:            r.AWSCloudFormationMetadata,
		DeletionPolicy:      r.AWSCloudFormationDeletionPolicy,
		UpdateReplacePolicy: r.AWSCloudFormationUpdateReplacePolicy,
		Condition:           r.AWSCloudFormationCondition,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *JobDefinition) UnmarshalJSON(b []byte) error {
	type Properties JobDefinition
	res := &struct {
		Type                string
		Properties          *Properties
		DependsOn           interface{}
		Metadata            map[string]interface{}
		DeletionPolicy      string
		UpdateReplacePolicy string
		Condition           string
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = JobDefinition(*res.Properties)
	}
	if res.DependsOn != nil {
		switch obj := res.DependsOn.(type) {
		case string:
			r.AWSCloudFormationDependsOn = []string{obj}
		case []interface{}:
			s := make([]string, 0, len(obj))
			for _, v := range obj {
				if value, ok := v.(string); ok {
					s = append(s, value)
				}
			}
			r.AWSCloudFormationDependsOn = s
		}
	}
	if res.Metadata != nil {
		r.AWSCloudFormationMetadata = res.Metadata
	}
	if res.DeletionPolicy != "" {
		r.AWSCloudFormationDeletionPolicy = policies.DeletionPolicy(res.DeletionPolicy)
	}
	if res.UpdateReplacePolicy != "" {
		r.AWSCloudFormationUpdateReplacePolicy = policies.UpdateReplacePolicy(res.UpdateReplacePolicy)
	}
	if res.Condition != "" {
		r.AWSCloudFormationCondition = res.Condition
	}
	return nil
}