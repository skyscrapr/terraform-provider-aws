package bedrock

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
)

func TestResourceCustomModelRefresh(t *testing.T) {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		t.Errorf(err.Error())
	}

	client := bedrock.NewFromConfig(cfg)

	modelArn := "arn:aws:bedrock:us-east-1:050412605402:custom-model/amazon.titan-text-express-v1:0:8k/t9iht281qhts"
	data := resourceCustomModelModel{
		ID: flex.StringToFramework(ctx, &modelArn),
	}

	diags := data.refresh(ctx, client)
	if diags.HasError() {
		t.Errorf(diags.Errors()[0].Detail())
	}
}
