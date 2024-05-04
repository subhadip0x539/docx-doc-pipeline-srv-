package port

import (
	"context"

	"docx-doc-pipeline-srv/src/internal/domain"
)

type IPipelineRepo interface {
	PublishEvent(ctx context.Context, event *domain.TEvent) domain.TError
}

type IPipelineService interface {
	Dispatch(ctx context.Context, event *domain.TEvent) domain.TError
}
