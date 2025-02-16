// Code generated by "go generate". Please don't change this file directly.

package emr

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Cluster_PlacementGroupConfig AWS CloudFormation Resource (AWS::EMR::Cluster.PlacementGroupConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-placementgroupconfig.html
type Cluster_PlacementGroupConfig struct {

	// InstanceRole AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-placementgroupconfig.html#cfn-elasticmapreduce-cluster-placementgroupconfig-instancerole
	InstanceRole string `json:"InstanceRole"`

	// PlacementStrategy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-placementgroupconfig.html#cfn-elasticmapreduce-cluster-placementgroupconfig-placementstrategy
	PlacementStrategy *string `json:"PlacementStrategy,omitempty"`

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
func (r *Cluster_PlacementGroupConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.PlacementGroupConfig"
}
