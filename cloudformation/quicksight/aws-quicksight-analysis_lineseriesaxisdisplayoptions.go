// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_LineSeriesAxisDisplayOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.LineSeriesAxisDisplayOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-lineseriesaxisdisplayoptions.html
type Analysis_LineSeriesAxisDisplayOptions struct {

	// AxisOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-lineseriesaxisdisplayoptions.html#cfn-quicksight-analysis-lineseriesaxisdisplayoptions-axisoptions
	AxisOptions *Analysis_AxisDisplayOptions `json:"AxisOptions,omitempty"`

	// MissingDataConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-lineseriesaxisdisplayoptions.html#cfn-quicksight-analysis-lineseriesaxisdisplayoptions-missingdataconfigurations
	MissingDataConfigurations []Analysis_MissingDataConfiguration `json:"MissingDataConfigurations,omitempty"`

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
func (r *Analysis_LineSeriesAxisDisplayOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.LineSeriesAxisDisplayOptions"
}
