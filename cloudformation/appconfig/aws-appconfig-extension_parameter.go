// Code generated by "go generate". Please don't change this file directly.

package appconfig

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Extension_Parameter AWS CloudFormation Resource (AWS::AppConfig::Extension.Parameter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-parameter.html
type Extension_Parameter struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-parameter.html#cfn-appconfig-extension-parameter-description
	Description *string `json:"Description,omitempty"`

	// Required AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-parameter.html#cfn-appconfig-extension-parameter-required
	Required bool `json:"Required"`

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
func (r *Extension_Parameter) AWSCloudFormationType() string {
	return "AWS::AppConfig::Extension.Parameter"
}
