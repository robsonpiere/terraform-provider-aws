---
subcategory: "VPC IPAM (IP Address Manager)"
layout: "aws"
page_title: "AWS: aws_vpc_ipams"
description: |-
  Terraform data source for managing VPC IPAMs.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_vpc_ipams

Terraform data source for managing VPC IPAMs.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsVpcIpams } from "./.gen/providers/aws/";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsVpcIpams(this, "example", {
      ipam_ids: ["ipam-abcd1234"],
    });
  }
}

```

### Filter by `tags`

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsVpcIpams } from "./.gen/providers/aws/";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsVpcIpams(this, "example", {
      filter: [
        {
          name: "tags.Some",
          values: ["Value"],
        },
      ],
    });
  }
}

```

### Filter by `tier`

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsVpcIpams } from "./.gen/providers/aws/";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsVpcIpams(this, "example", {
      filter: [
        {
          name: "tier",
          values: ["free"],
        },
      ],
    });
  }
}

```

## Argument Reference

The arguments of this data source act as filters for querying the available IPAMs.

* `ipam_ids` - (Optional) IDs of the IPAM resources to query for.
* `filter` - (Optional) Custom filter block as described below.

More complex filters can be expressed using one or more `filter` sub-blocks,
which take the following arguments:

* `name` - (Required) Name of the field to filter by, as defined by
  [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeIpams.html).

* `values` - (Required) Set of values that are accepted for the given field.
  An IPAM resource will be selected if any one of the given values matches.

## Attribute Reference

All of the argument attributes except `filter` are also exported as result attributes.

* `ipams` - List of IPAM resources matching the provided arguments.

### ipams

* `arn` - ARN of the IPAM.
* `defaultResourceDiscoveryAssociationId` - The default resource discovery association ID.
* `defaultResourceDiscoveryId` - The default resource discovery ID.
* `description` - Description for the IPAM.
* `enablePrivateGua` - If private GUA is enabled.
* `id` - ID of the IPAM resource.
* `ipamRegion` - Region that the IPAM exists in.
* `operatingRegions` - Regions that the IPAM is configured to operate in.
* `ownerId` - ID of the account that owns this IPAM.
* `privateDefaultScopeId` - ID of the default private scope.
* `publicDefaultScopeId` - ID of the default public scope.
* `resource_discovery_association_count` - Number of resource discovery associations.
* `scopeCount` - Number of scopes on this IPAM.
* `state` - Current state of the IPAM.
* `state_message` - State message of the IPAM.
* `tier` - IPAM Tier.

<!-- cache-key: cdktf-0.20.8 input-ad363e749c8448183da9939f9c6580c3e1b883ebada50a6e6506ce515a9a96d3 -->