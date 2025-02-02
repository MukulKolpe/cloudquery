// Code generated by codegen; DO NOT EDIT.

package dms

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ReplicationInstances() *schema.Table {
	return &schema.Table{
		Name:        "aws_dms_replication_instances",
		Description: `https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html`,
		Resolver:    fetchDmsReplicationInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("dms"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationInstanceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "allocated_storage",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AllocatedStorage"),
			},
			{
				Name:     "auto_minor_version_upgrade",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoMinorVersionUpgrade"),
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "dns_name_servers",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DnsNameServers"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "free_until",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("FreeUntil"),
			},
			{
				Name:     "instance_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("InstanceCreateTime"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "multi_az",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MultiAZ"),
			},
			{
				Name:     "pending_modified_values",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingModifiedValues"),
			},
			{
				Name:     "preferred_maintenance_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredMaintenanceWindow"),
			},
			{
				Name:     "publicly_accessible",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PubliclyAccessible"),
			},
			{
				Name:     "replication_instance_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationInstanceClass"),
			},
			{
				Name:     "replication_instance_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationInstanceIdentifier"),
			},
			{
				Name:     "replication_instance_private_ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationInstancePrivateIpAddress"),
			},
			{
				Name:     "replication_instance_private_ip_addresses",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReplicationInstancePrivateIpAddresses"),
			},
			{
				Name:     "replication_instance_public_ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationInstancePublicIpAddress"),
			},
			{
				Name:     "replication_instance_public_ip_addresses",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReplicationInstancePublicIpAddresses"),
			},
			{
				Name:     "replication_instance_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationInstanceStatus"),
			},
			{
				Name:     "replication_subnet_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReplicationSubnetGroup"),
			},
			{
				Name:     "secondary_availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecondaryAvailabilityZone"),
			},
			{
				Name:     "vpc_security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcSecurityGroups"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}
