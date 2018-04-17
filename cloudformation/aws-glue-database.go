package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSGlueDatabase AWS CloudFormation Resource (AWS::Glue::Database)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html
type AWSGlueDatabase struct {

	// CatalogId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html#cfn-glue-database-catalogid
	CatalogId string `json:"CatalogId,omitempty"`

	// DatabaseInput AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html#cfn-glue-database-databaseinput
	DatabaseInput *AWSGlueDatabase_DatabaseInput `json:"DatabaseInput,omitempty"`

	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueDatabase) AWSCloudFormationType() string {
	return "AWS::Glue::Database"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSGlueDatabase) MarshalJSON() ([]byte, error) {
	type Properties AWSGlueDatabase
	return json.Marshal(&struct {
		Type string

		DependsOn  *[]string               `json:",omitempty"`
		Metadata   *map[string]interface{} `json:",omitempty"`
		Properties Properties
	}{
		Type: r.AWSCloudFormationType(),

		DependsOn:  r.DependsOn,
		Metadata:   r.Metadata,
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSGlueDatabase) UnmarshalJSON(b []byte) error {
	type Properties AWSGlueDatabase
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSGlueDatabase(*res.Properties)
	}

	return nil
}

// GetAllAWSGlueDatabaseResources retrieves all AWSGlueDatabase items from an AWS CloudFormation template
func (t *Template) GetAllAWSGlueDatabaseResources() map[string]AWSGlueDatabase {
	results := map[string]AWSGlueDatabase{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSGlueDatabase:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Glue::Database" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSGlueDatabase
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSGlueDatabaseWithName retrieves all AWSGlueDatabase items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSGlueDatabaseWithName(name string) (AWSGlueDatabase, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSGlueDatabase:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Glue::Database" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSGlueDatabase
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSGlueDatabase{}, errors.New("resource not found")
}
