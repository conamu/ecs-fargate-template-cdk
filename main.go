package main

import (
	"cdk-fargate/infra"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

// this is where the infra is built and deployed

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	infra.NewCdkFargateStack(app, "CdkFargateStack", &infra.CdkFargateStackProps{
		awscdk.StackProps{
			Env: infra.Env(),
		},
	})

	app.Synth(nil)
}
