package fsx

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rs/zerolog/log"
)

func fetchFsxFileCaches(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Fsx
	input := fsx.DescribeFileCachesInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeFileCachesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			var ae smithy.APIError
			if errors.As(err, &ae) && ae.ErrorCode() == "BadRequest" && ae.ErrorMessage() == "The requested feature is not enabled for this AWS account." {
				log.Warn().Err(err).Msg("skipping resource")
				return nil
			}

			return err
		}
		res <- result.FileCaches
	}
	return nil
}

func resolveFileCacheTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.FileCache)
	cl := meta.(*client.Client)
	svc := cl.Services().Fsx
	out, err := svc.ListTagsForResource(ctx, &fsx.ListTagsForResourceInput{ResourceARN: item.ResourceARN})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.Tags))
}
