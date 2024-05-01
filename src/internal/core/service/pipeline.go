package service

import (
	"fmt"
	"net/http"

	"context"

	"docx-doc-pipeline-srv/src/internal/domain"
	"docx-doc-pipeline-srv/src/internal/port"
)

type TPipelineService struct {
	repo port.IPipelineRepo
}

func NewPipelineService(repo port.IPipelineRepo) *TPipelineService {
	return &TPipelineService{repo: repo}
}

func (s *TPipelineService) Dispatch(ctx context.Context, event *domain.TEvent) domain.TError {
	switch event.Type {
	case domain.PdfMergeEventTag:
		event.Key = domain.PdfMergeRequestEventKey
		return s.repo.PublishEvent(ctx, event)
	case domain.PdfSplitEventType:
		event.Key = domain.PdfSplitRequestEventKey
		return s.repo.PublishEvent(ctx, event)
	default:
		return domain.TError{
			Code:    http.StatusBadRequest,
			Message: domain.BadRequest,
			Error:   fmt.Errorf("unknown event type: %s", event.Type),
		}
	}
}
