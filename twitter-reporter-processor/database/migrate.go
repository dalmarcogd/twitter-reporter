package database

import (
	"context"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/monitoring"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/utils"
	"go.elastic.co/apm"
)

func Migrate() error {
	tx := monitoring.GetTracer().StartTransaction("Migrate Database", "database")
	defer tx.End()
	ctx := apm.ContextWithTransaction(context.Background(), tx)
	defer ctx.Done()
	return utils.SpanTracer(ctx, "Database Migrate", "database.migrate", func(cx context.Context, span *apm.Span) error {
		return GetConnection(ctx).AutoMigrate(&ReporterModel{}, &TwitterUserModel{}, &TwitterTweetModel{}).Error
	})
}
