package serverless

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// Scaling_Config, see https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-sqs.html#sam-function-sqs-scalingconfig
// Ryan: created from latest AWS doc
type Scaling_Config struct {
	// Min 2, Max 1000 (according to doc)
	MaximumConcurrency int `json:"MaximumConcurrency"`
}

// Function_SQSEvent AWS CloudFormation Resource (AWS::Serverless::Function.SQSEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#sqs
type Function_SQSEvent struct {

	// BatchSize AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#sqs
	BatchSize int `json:"BatchSize,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#sqs
	Enabled bool `json:"Enabled,omitempty"`

	// Queue AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#sqs
	Queue string `json:"Queue,omitempty"`

	// ScalingConfig AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-function-sqs.html#sam-function-sqs-scalingconfig
	// Ryan: created from latest AWS doc
	ScalingConfig *Scaling_Config `json:"ScalingConfig,omitempty"`

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
func (r *Function_SQSEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.SQSEvent"
}
