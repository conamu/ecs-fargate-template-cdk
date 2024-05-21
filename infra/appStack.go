package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

var vpc awsec2.Vpc

func AppStack(scope constructs.Construct, id string, props *awscdk.NestedStackProps) awscdk.Stack {

	stack := awscdk.NewNestedStack(scope, &id, props)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("CdkFargateQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	appCluster := awsecs.NewCluster(stack, jsii.String("template-app-cluster"), &awsecs.ClusterProps{
		Vpc: vpc,
	})

	awsecspatterns.NewApplicationLoadBalancedFargateService(
		stack,
		jsii.String("template-fargate-service"),
		&awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
			Cluster: appCluster,
			RuntimePlatform: &awsecs.RuntimePlatform{
				CpuArchitecture:       awsecs.CpuArchitecture_ARM64(),
				OperatingSystemFamily: awsecs.OperatingSystemFamily_LINUX(),
			},
			Cpu:          jsii.Number(256),
			DesiredCount: jsii.Number(3),
			TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
				Image: awsecs.ContainerImage_FromAsset(jsii.String("."), &awsecs.AssetImageProps{
					Platform: awsecrassets.Platform_LINUX_ARM64(),
				})},
			MemoryLimitMiB:     jsii.Number(512),
			PublicLoadBalancer: jsii.Bool(true),
		},
	)

	return stack
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func Env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	//return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String("038796470268"),
		Region:  jsii.String("eu-west-1"),
	}

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
