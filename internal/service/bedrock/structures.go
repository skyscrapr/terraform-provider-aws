// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrock

import (
	"context"

	bedrock_types "github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
)

// func flattenTrainingMetrics(metrics *bedrock.TrainingMetrics) []map[string]interface{} {
// 	if metrics == nil {
// 		return []map[string]interface{}{}
// 	}

// 	m := map[string]interface{}{
// 		"training_loss": aws.Float64Value(metrics.TrainingLoss),
// 	}

// 	return []map[string]interface{}{m}
// }

// func flattenValidationDataConfig(config *bedrock_types.ValidationDataConfig) []map[string]interface{} {
// 	if config == nil {
// 		return []map[string]interface{}{}
// 	}

// 	l := make([]map[string]interface{}, 0, len(config.Validators))

// 	for _, validator := range config.Validators {
// 		m := map[string]interface{}{
// 			"validator": aws.StringValue(validator.S3Uri),
// 		}
// 		l = append(l, m)
// 	}

// 	return l
// }

// func flattenValidationMetrics(metrics []*bedrock.ValidatorMetric) []map[string]interface{} {
// 	if metrics == nil {
// 		return []map[string]interface{}{}
// 	}

// 	l := make([]map[string]interface{}, 0, len(metrics))

// 	for _, metric := range metrics {
// 		m := map[string]interface{}{
// 			"validation_loss": aws.Float64Value(metric.ValidationLoss),
// 		}
// 		l = append(l, m)
// 	}

// 	return l
// }

// func flattenCustomModelSummaries(models []*bedrock_types.CustomModelSummary) []map[string]interface{} {
// 	if len(models) == 0 {
// 		return []map[string]interface{}{}
// 	}

// 	l := make([]map[string]interface{}, 0, len(models))

// 	for _, model := range models {
// 		m := map[string]interface{}{
// 			"base_model_arn":  aws.StringValue(model.BaseModelArn),
// 			"base_model_name": aws.StringValue(model.BaseModelName),
// 			"model_arn":       aws.StringValue(model.ModelArn),
// 			"model_name":      aws.StringValue(model.ModelName),
// 			"creation_time":   aws.TimeValue(model.CreationTime).Format(time.RFC3339),
// 		}
// 		l = append(l, m)
// 	}

// 	return l
// }

// func flattenFoundationModelSummaries(models []*bedrock.FoundationModelSummary) []map[string]interface{} {
// 	if len(models) == 0 {
// 		return []map[string]interface{}{}
// 	}

// 	l := make([]map[string]interface{}, 0, len(models))

// 	for _, model := range models {
// 		m := map[string]interface{}{
// 			"model_arn":                    aws.StringValue(model.ModelArn),
// 			"model_id":                     aws.StringValue(model.ModelId),
// 			"model_name":                   aws.StringValue(model.ModelName),
// 			"provider_name":                aws.StringValue(model.ProviderName),
// 			"customizations_supported":     aws.StringValueSlice(model.CustomizationsSupported),
// 			"inference_types_supported":    aws.StringValueSlice(model.InferenceTypesSupported),
// 			"input_modalities":             aws.StringValueSlice(model.InputModalities),
// 			"output_modalities":            aws.StringValueSlice(model.OutputModalities),
// 			"response_streaming_supported": aws.BoolValue(model.ResponseStreamingSupported),
// 		}

// 		l = append(l, m)
// 	}

// 	return l
// }

func flattenLoggingConfig(ctx context.Context, apiObject *bedrock_types.LoggingConfig) *loggingConfigModel {
	if apiObject == nil {
		return nil
	}

	return &loggingConfigModel{
		EmbeddingDataDeliveryEnabled: flex.BoolToFramework(ctx, apiObject.EmbeddingDataDeliveryEnabled),
		ImageDataDeliveryEnabled:     flex.BoolToFramework(ctx, apiObject.ImageDataDeliveryEnabled),
		TextDataDeliveryEnabled:      flex.BoolToFramework(ctx, apiObject.TextDataDeliveryEnabled),
		CloudWatchConfig:             flattenCloudWatchConfig(ctx, apiObject.CloudWatchConfig),
		S3Config:                     flattenS3Config(ctx, apiObject.S3Config),
	}
}

func flattenCloudWatchConfig(ctx context.Context, apiObject *bedrock_types.CloudWatchConfig) *cloudWatchConfigModel {
	if apiObject == nil {
		return nil
	}

	return &cloudWatchConfigModel{
		LogGroupName:              flex.StringToFramework(ctx, apiObject.LogGroupName),
		RoleArn:                   flex.StringToFramework(ctx, apiObject.RoleArn),
		LargeDataDeliveryS3Config: flattenS3Config(ctx, apiObject.LargeDataDeliveryS3Config),
	}
}

func flattenS3Config(ctx context.Context, apiObject *bedrock_types.S3Config) *s3ConfigModel {
	if apiObject == nil {
		return nil
	}

	return &s3ConfigModel{
		BucketName: flex.StringToFramework(ctx, apiObject.BucketName),
		KeyPrefix:  flex.StringToFramework(ctx, apiObject.KeyPrefix),
	}
}

func expandLoggingConfig(model *loggingConfigModel) *bedrock_types.LoggingConfig {
	if model == nil {
		return nil
	}

	apiObject := &bedrock_types.LoggingConfig{
		EmbeddingDataDeliveryEnabled: model.EmbeddingDataDeliveryEnabled.ValueBoolPointer(),
		ImageDataDeliveryEnabled:     model.ImageDataDeliveryEnabled.ValueBoolPointer(),
		TextDataDeliveryEnabled:      model.TextDataDeliveryEnabled.ValueBoolPointer(),
	}
	if model.CloudWatchConfig != nil {
		apiObject.CloudWatchConfig = expandCloudWatchConfig(model.CloudWatchConfig)
	}
	if model.S3Config != nil {
		apiObject.S3Config = expandS3Config(model.S3Config)
	}

	return apiObject
}

func expandCloudWatchConfig(model *cloudWatchConfigModel) *bedrock_types.CloudWatchConfig {
	if model == nil {
		return nil
	}

	apiObject := &bedrock_types.CloudWatchConfig{
		LogGroupName:              model.LogGroupName.ValueStringPointer(),
		RoleArn:                   model.RoleArn.ValueStringPointer(),
		LargeDataDeliveryS3Config: expandS3Config(model.LargeDataDeliveryS3Config),
	}

	return apiObject
}

func expandS3Config(model *s3ConfigModel) *bedrock_types.S3Config {
	if model == nil {
		return nil
	}

	apiObject := &bedrock_types.S3Config{
		BucketName: model.BucketName.ValueStringPointer(),
		KeyPrefix:  model.KeyPrefix.ValueStringPointer(),
	}

	return apiObject
}

func expandVpcConfig(ctx context.Context, tfList []vpcConfigModel) []bedrock_types.VpcConfig {
	if len(tfList) == 0 {
		return nil
	}
	var vpcConfigs []bedrock_types.VpcConfig

	for _, item := range tfList {
		vpcConfig := bedrock_types.VpcConfig{
			SecurityGroupIds: flex.ExpandFrameworkStringValueSet(ctx, item.SecurityGroupIds),
			SubnetIds:        flex.ExpandFrameworkStringValueSet(ctx, item.SubnetIds),
		}
		vpcConfigs = append(vpcConfigs, vpcConfig)
	}

	return vpcConfigs
}

func expandValidationDataConfig(ctx context.Context, model *validationDataConfigModel) *bedrock_types.ValidationDataConfig {
	if model == nil {
		return nil
	}

	apiObject := &bedrock_types.ValidationDataConfig{}
	for _, validator := range flex.ExpandFrameworkStringValueSet(ctx, model.Validators) {
		apiObject.Validators = append(apiObject.Validators, bedrock_types.Validator{
			S3Uri: &validator,
		})
	}

	return apiObject
}
