// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceApplication,
			TypeName: "aws_elastic_beanstalk_application",
			Name:     "Application",
		},
		{
			Factory:  dataSourceHostedZone,
			TypeName: "aws_elastic_beanstalk_hosted_zone",
			Name:     "Hosted Zone",
		},
		{
			Factory:  dataSourceSolutionStack,
			TypeName: "aws_elastic_beanstalk_solution_stack",
			Name:     "Solution Stack",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceApplication,
			TypeName: "aws_elastic_beanstalk_application",
			Name:     "Application",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceApplicationVersion,
			TypeName: "aws_elastic_beanstalk_application_version",
			Name:     "Application Version",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceConfigurationTemplate,
			TypeName: "aws_elastic_beanstalk_configuration_template",
			Name:     "Configuration Template",
		},
		{
			Factory:  resourceEnvironment,
			TypeName: "aws_elastic_beanstalk_environment",
			Name:     "Environment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ElasticBeanstalk
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*elasticbeanstalk.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*elasticbeanstalk.Options){
		elasticbeanstalk.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return elasticbeanstalk.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*elasticbeanstalk.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*elasticbeanstalk.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *elasticbeanstalk.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*elasticbeanstalk.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
