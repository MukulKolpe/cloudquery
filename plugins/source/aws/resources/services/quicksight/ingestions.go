// Code generated by codegen; DO NOT EDIT.

package quicksight

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ingestions() *schema.Table {
	return &schema.Table{
		Name:      "aws_quicksight_ingestions",
		Resolver:  fetchQuicksightIngestions,
		Multiplex: client.ServiceAccountRegionMultiplexer("quicksight"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveTags(),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "data_set_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "ingestion_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IngestionStatus"),
			},
			{
				Name:     "error_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ErrorInfo"),
			},
			{
				Name:     "ingestion_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IngestionId"),
			},
			{
				Name:     "ingestion_size_in_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("IngestionSizeInBytes"),
			},
			{
				Name:     "ingestion_time_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("IngestionTimeInSeconds"),
			},
			{
				Name:     "queue_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("QueueInfo"),
			},
			{
				Name:     "request_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RequestSource"),
			},
			{
				Name:     "request_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RequestType"),
			},
			{
				Name:     "row_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RowInfo"),
			},
		},
	}
}
