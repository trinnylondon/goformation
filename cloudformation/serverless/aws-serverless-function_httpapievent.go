package serverless

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

type Function_HttpApiFunctionAuth struct {
	AuthorizationScopes []string `json:"AuthorizationScopes,omitempty"`
	Authorizer          string   `json:"Authorizer,omitempty"`
}

// Function_HttpApiEvent AWS CloudFormation Resource (AWS::Serverless::Function.HttpApiEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
type Function_HttpApiEvent struct {
	// Required: false
	APIID string `json:"ApiId,omitempty"`

	// Required: false
	Auth *Function_HttpApiFunctionAuth `json:"Auth,omitempty"`

	// Method AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
	Method string `json:"Method,omitempty"`

	// Path AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
	Path string `json:"Path,omitempty"`

	// PayloadFormatVersion is "1.0" or "2.0".  Default to "2.0".
	// Required: false
	PayloadFormatVersion string `json:"PayloadFormatVersion,omitempty"`

	// Required: false
	RouteSettings interface{} `json:"RouteSettings,omitempty"`

	// TimeoutInMillis is custom timeout between 50 and 29,000 milliseconds.  Default to 5000.
	// Required: false
	TimeoutInMillis int `json:"TimeoutInMillis,omitempty"`

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
func (r *Function_HttpApiEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.HttpApiEvent"
}
