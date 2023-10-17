// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
)

// @FrameworkDataSource
func newDataSourceFoundationModel(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &dataSourceFoundationModel{}, nil
}

type dataSourceFoundationModel struct {
	framework.DataSourceWithConfigure
}

// Metadata should return the full name of the data source, such as
// examplecloud_thing.
func (d *dataSourceFoundationModel) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = "aws_bedrock_foundation_model"
}

// Schema returns the schema for this data source.
func (d *dataSourceFoundationModel) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"model_id": schema.StringAttribute{
				Required: true,
			},
			"model_arn": schema.StringAttribute{
				Computed: true,
			},
			"model_name": schema.StringAttribute{
				Computed: true,
			},
			"provider_name": schema.StringAttribute{
				Computed: true,
			},
			"customizations_supported": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"inference_types_supported": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"input_modalities": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"output_modalities": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"response_streaming_supported": schema.BoolAttribute{
				Computed: true,
			},
		},
	}
}

// Read is called when the provider must read data source values in order to update state.
// Config values should be read from the ReadRequest and new state values set on the ReadResponse.
func (d *dataSourceFoundationModel) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data foundationModel

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	conn := d.Meta().BedrockClient(ctx)

	input := &bedrock.GetFoundationModelInput{
		ModelIdentifier: data.ModelID.ValueStringPointer(),
	}
	model, err := conn.GetFoundationModel(ctx, input)
	if err != nil {
		response.Diagnostics.AddError("reading Bedrock Foundation Model", err.Error())
		return
	}

	response.Diagnostics.Append(data.refreshFromOutput(ctx, model)...)
	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

type foundationModel struct {
	ID                         types.String `tfsdk:"id"`
	ModelID                    types.String `tfsdk:"model_id"`
	ModelArn                   types.String `tfsdk:"model_arn"`
	ModelName                  types.String `tfsdk:"model_name"`
	ProviderName               types.String `tfsdk:"provider_name"`
	CustomizationsSupported    types.Set    `tfsdk:"customizations_supported"`
	InferenceTypesSupported    types.Set    `tfsdk:"inference_types_supported"`
	InputModalities            types.Set    `tfsdk:"input_modalities"`
	OutputModalities           types.Set    `tfsdk:"output_modalities"`
	ResponseStreamingSupported types.Bool   `tfsdk:"response_streaming_supported"`
}

func (data *foundationModel) refreshFromOutput(ctx context.Context, model *bedrock.GetFoundationModelOutput) diag.Diagnostics {
	var diags diag.Diagnostics

	if model == nil {
		return diags
	}

	data.ID = flex.StringToFramework(ctx, model.ModelDetails.ModelId)
	data.ModelArn = flex.StringToFramework(ctx, model.ModelDetails.ModelArn)
	data.ModelID = flex.StringToFramework(ctx, model.ModelDetails.ModelId)
	data.ModelName = flex.StringToFramework(ctx, model.ModelDetails.ModelName)
	data.ProviderName = flex.StringToFramework(ctx, model.ModelDetails.ProviderName)
	customizationsSupported := make([]string, 0, len(model.ModelDetails.CustomizationsSupported))
	for _, r := range model.ModelDetails.CustomizationsSupported {
		customizationsSupported = append(customizationsSupported, string(r))
	}
	data.CustomizationsSupported = flex.FlattenFrameworkStringValueSet(ctx, customizationsSupported)

	inferenceTypesSupported := make([]string, 0, len(model.ModelDetails.InferenceTypesSupported))
	for _, r := range model.ModelDetails.InferenceTypesSupported {
		inferenceTypesSupported = append(inferenceTypesSupported, string(r))
	}
	data.InferenceTypesSupported = flex.FlattenFrameworkStringValueSet(ctx, inferenceTypesSupported)

	inputModalities := make([]string, 0, len(model.ModelDetails.InputModalities))
	for _, r := range model.ModelDetails.InputModalities {
		inputModalities = append(inputModalities, string(r))
	}
	data.InputModalities = flex.FlattenFrameworkStringValueSet(ctx, inputModalities)

	outputModalities := make([]string, 0, len(model.ModelDetails.OutputModalities))
	for _, r := range model.ModelDetails.OutputModalities {
		outputModalities = append(outputModalities, string(r))
	}
	data.OutputModalities = flex.FlattenFrameworkStringValueSet(ctx, outputModalities)

	data.ResponseStreamingSupported = flex.BoolToFramework(ctx, model.ModelDetails.ResponseStreamingSupported)

	return diags
}
