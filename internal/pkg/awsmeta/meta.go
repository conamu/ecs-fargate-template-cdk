package awsmeta

import (
	"context"
	metadata "github.com/brunoscheufler/aws-ecs-metadata-go"
	"log/slog"
	"net/http"
)

type awsmeta struct {
	logger *slog.Logger
}

var metaGetter *awsmeta

func Init(logger *slog.Logger) {
	metaGetter = &awsmeta{logger: logger}
}

func Get() *metadata.ContainerMetadataV4 {
	meta, err := metadata.GetContainerV4(context.Background(), &http.Client{})
	if err != nil {
		metaGetter.logger.Warn("Could not get container metadata", "err", err.Error())
	}
	return meta
}
