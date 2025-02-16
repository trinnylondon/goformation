// Code generated by "go generate". Please don't change this file directly.

package lakeformation

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// TagAssociation_Resource AWS CloudFormation Resource (AWS::LakeFormation::TagAssociation.Resource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-tagassociation-resource.html
type TagAssociation_Resource struct {

	// Catalog AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-tagassociation-resource.html#cfn-lakeformation-tagassociation-resource-catalog
	Catalog interface{} `json:"Catalog,omitempty"`

	// Database AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-tagassociation-resource.html#cfn-lakeformation-tagassociation-resource-database
	Database *TagAssociation_DatabaseResource `json:"Database,omitempty"`

	// Table AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-tagassociation-resource.html#cfn-lakeformation-tagassociation-resource-table
	Table *TagAssociation_TableResource `json:"Table,omitempty"`

	// TableWithColumns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-tagassociation-resource.html#cfn-lakeformation-tagassociation-resource-tablewithcolumns
	TableWithColumns *TagAssociation_TableWithColumnsResource `json:"TableWithColumns,omitempty"`

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
func (r *TagAssociation_Resource) AWSCloudFormationType() string {
	return "AWS::LakeFormation::TagAssociation.Resource"
}
