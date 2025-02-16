// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_TimeRangeFilter AWS CloudFormation Resource (AWS::QuickSight::Template.TimeRangeFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangefilter.html
type Template_TimeRangeFilter struct {

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangefilter.html#cfn-quicksight-template-timerangefilter-column
	Column *Template_ColumnIdentifier `json:"Column"`

	// ExcludePeriodConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangefilter.html#cfn-quicksight-template-timerangefilter-excludeperiodconfiguration
	ExcludePeriodConfiguration *Template_ExcludePeriodConfiguration `json:"ExcludePeriodConfiguration,omitempty"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangefilter.html#cfn-quicksight-template-timerangefilter-filterid
	FilterId string `json:"FilterId"`

	// IncludeMaximum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangefilter.html#cfn-quicksight-template-timerangefilter-includemaximum
	IncludeMaximum *bool `json:"IncludeMaximum,omitempty"`

	// IncludeMinimum AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangefilter.html#cfn-quicksight-template-timerangefilter-includeminimum
	IncludeMinimum *bool `json:"IncludeMinimum,omitempty"`

	// NullOption AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangefilter.html#cfn-quicksight-template-timerangefilter-nulloption
	NullOption string `json:"NullOption"`

	// RangeMaximumValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangefilter.html#cfn-quicksight-template-timerangefilter-rangemaximumvalue
	RangeMaximumValue *Template_TimeRangeFilterValue `json:"RangeMaximumValue,omitempty"`

	// RangeMinimumValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangefilter.html#cfn-quicksight-template-timerangefilter-rangeminimumvalue
	RangeMinimumValue *Template_TimeRangeFilterValue `json:"RangeMinimumValue,omitempty"`

	// TimeGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangefilter.html#cfn-quicksight-template-timerangefilter-timegranularity
	TimeGranularity *string `json:"TimeGranularity,omitempty"`

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
func (r *Template_TimeRangeFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.TimeRangeFilter"
}
