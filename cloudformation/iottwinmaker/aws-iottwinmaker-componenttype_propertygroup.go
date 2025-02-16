// Code generated by "go generate". Please don't change this file directly.

package iottwinmaker

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ComponentType_PropertyGroup AWS CloudFormation Resource (AWS::IoTTwinMaker::ComponentType.PropertyGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-propertygroup.html
type ComponentType_PropertyGroup struct {

	// GroupType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-propertygroup.html#cfn-iottwinmaker-componenttype-propertygroup-grouptype
	GroupType *string `json:"GroupType,omitempty"`

	// PropertyNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-propertygroup.html#cfn-iottwinmaker-componenttype-propertygroup-propertynames
	PropertyNames []string `json:"PropertyNames,omitempty"`

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
func (r *ComponentType_PropertyGroup) AWSCloudFormationType() string {
	return "AWS::IoTTwinMaker::ComponentType.PropertyGroup"
}
