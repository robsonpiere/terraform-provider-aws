package sagemaker_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/sagemaker"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfsagemaker "github.com/hashicorp/terraform-provider-aws/internal/service/sagemaker"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func TestAccSageMakerMonitoringSchedule_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_monitoring_schedule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckMonitoringScheduleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringScheduleConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitoringScheduleExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					acctest.CheckResourceAttrRegionalARN(resourceName, "arn", "sagemaker", fmt.Sprintf("monitoring-schedule/%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "monitoring_schedule_config.0.monitoring_job_definition_name", "aws_sagemaker_data_quality_job_definition.test", "name"),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.0.monitoring_type", "DataQuality"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSageMakerMonitoringSchedule_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_monitoring_schedule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckMonitoringScheduleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringScheduleConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataQualityJobDefinitionExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccMonitoringScheduleConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataQualityJobDefinitionExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccMonitoringScheduleConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataQualityJobDefinitionExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func TestAccSageMakerMonitoringSchedule_scheduleExpression(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_monitoring_schedule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckMonitoringScheduleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringScheduleConfig_scheduleExpressionHourly(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataQualityJobDefinitionExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.0.schedule_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.0.schedule_config.0.schedule_expression", "cron(0 * ? * * *)"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccMonitoringScheduleConfig_scheduleExpressionDaily(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataQualityJobDefinitionExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.0.schedule_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.0.schedule_config.0.schedule_expression", "cron(0 0 ? * * *)"),
				),
			},
			{
				Config: testAccMonitoringScheduleConfig_scheduleExpressionHourly(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataQualityJobDefinitionExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.0.schedule_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "monitoring_schedule_config.0.schedule_config.0.schedule_expression", "cron(0 * ? * * *)"),
				),
			},
		},
	})
}

func TestAccSageMakerMonitoringSchedule_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_sagemaker_monitoring_schedule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, sagemaker.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckMonitoringScheduleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringScheduleConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitoringScheduleExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceMonitoringSchedule(), resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfsagemaker.ResourceMonitoringSchedule(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckMonitoringScheduleDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn()

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_sagemaker_monitoring_schedule" {
				continue
			}

			_, err := tfsagemaker.FindMonitoringScheduleByName(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("SageMaker Monitoring Schedule (%s) still exists", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckMonitoringScheduleExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no SageMaker Monitoring Schedule ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).SageMakerConn()
		_, err := tfsagemaker.FindMonitoringScheduleByName(ctx, conn, rs.Primary.ID)

		return err
	}
}

func testAccMonitoringScheduleConfig_base(rName string) string {
	return fmt.Sprintf(`
data "aws_iam_policy_document" "access" {
  statement {
    effect = "Allow"

    actions = [
      "cloudwatch:PutMetricData",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "logs:CreateLogGroup",
      "logs:DescribeLogStreams",
      "ecr:GetAuthorizationToken",
      "ecr:BatchCheckLayerAvailability",
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "s3:GetObject",
    ]

    resources = ["*"]
  }
}

data "aws_partition" "current" {}

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["sagemaker.${data.aws_partition.current.dns_suffix}"]
    }
  }
}

resource "aws_iam_role" "test" {
  name               = %[1]q
  path               = "/"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy" "test" {
  role   = aws_iam_role.test.name
  policy = data.aws_iam_policy_document.access.json
}

resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_bucket_acl" "test" {
  bucket = aws_s3_bucket.test.id
  acl    = "private"
}

data "aws_sagemaker_prebuilt_ecr_image" "monitor" {
  repository_name = "sagemaker-model-monitor-analyzer"
  image_tag       = "latest"
}

resource "aws_sagemaker_data_quality_job_definition" "test" {
  name                 = %[1]q
  data_quality_app_specification {
    image_uri = data.aws_sagemaker_prebuilt_ecr_image.monitor.registry_path
  }
  data_quality_job_input {
    batch_transform_input {
      data_captured_destination_s3_uri = "https://${aws_s3_bucket.test.bucket_regional_domain_name}/captured"
      dataset_format {
        csv {}
      }
    }
  }
  data_quality_job_output_config {
    monitoring_outputs {
      s3_output {
	s3_uri = "https://${aws_s3_bucket.test.bucket_regional_domain_name}/output"
      }
    }
  }
  job_resources {
    cluster_config {
      instance_count = 1
      instance_type = "ml.t3.medium"
      volume_size_in_gb = 20
    }
  }
  role_arn = aws_iam_role.test.arn
}
`, rName)
}

func testAccMonitoringScheduleConfig_basic(rName string) string {
	return testAccMonitoringScheduleConfig_base(rName) + fmt.Sprintf(`
resource "aws_sagemaker_monitoring_schedule" "test" {
  name                 = %[1]q
  monitoring_schedule_config {
    monitoring_job_definition_name = aws_sagemaker_data_quality_job_definition.test.name
    monitoring_type = "DataQuality"
  }
}
`, rName)
}

func testAccMonitoringScheduleConfig_tags1(rName string, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccMonitoringScheduleConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_monitoring_schedule" "test" {
  name                 = %[1]q
  monitoring_schedule_config {
    monitoring_job_definition_name = aws_sagemaker_data_quality_job_definition.test.name
    monitoring_type = "DataQuality"
  }

  tags = {
    %[2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1))
}

func testAccMonitoringScheduleConfig_tags2(rName string, tagKey1, tagValue1 string, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccMonitoringScheduleConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_monitoring_schedule" "test" {
  name                 = %[1]q
  monitoring_schedule_config {
    monitoring_job_definition_name = aws_sagemaker_data_quality_job_definition.test.name
    monitoring_type = "DataQuality"
  }

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2))
}

func testAccMonitoringScheduleConfig_scheduleExpressionHourly(rName string) string {
	return acctest.ConfigCompose(testAccMonitoringScheduleConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_monitoring_schedule" "test" {
  name                 = %[1]q
  monitoring_schedule_config {
    monitoring_job_definition_name = aws_sagemaker_data_quality_job_definition.test.name
    monitoring_type = "DataQuality"
    schedule_config {
      schedule_expression = "cron(0 * ? * * *)"
    }
  }
}
`, rName))
}

func testAccMonitoringScheduleConfig_scheduleExpressionDaily(rName string) string {
	return acctest.ConfigCompose(testAccMonitoringScheduleConfig_base(rName), fmt.Sprintf(`
resource "aws_sagemaker_monitoring_schedule" "test" {
  name                 = %[1]q
  monitoring_schedule_config {
    monitoring_job_definition_name = aws_sagemaker_data_quality_job_definition.test.name
    monitoring_type = "DataQuality"
    schedule_config {
      schedule_expression = "cron(0 0 ? * * *)"
    }
  }
}
`, rName))
}
