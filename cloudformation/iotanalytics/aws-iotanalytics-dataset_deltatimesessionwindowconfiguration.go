// Code generated by "go generate". Please don't change this file directly.

package iotanalytics

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dataset_DeltaTimeSessionWindowConfiguration AWS CloudFormation Resource (AWS::IoTAnalytics::Dataset.DeltaTimeSessionWindowConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-deltatimesessionwindowconfiguration.html
type Dataset_DeltaTimeSessionWindowConfiguration struct {

	// TimeoutInMinutes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-deltatimesessionwindowconfiguration.html#cfn-iotanalytics-dataset-deltatimesessionwindowconfiguration-timeoutinminutes
	TimeoutInMinutes int `json:"TimeoutInMinutes"`

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
func (r *Dataset_DeltaTimeSessionWindowConfiguration) AWSCloudFormationType() string {
	return "AWS::IoTAnalytics::Dataset.DeltaTimeSessionWindowConfiguration"
}
