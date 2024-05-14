package service

import (
	"context"
	"fmt"

	"net/http"

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
	case domain.PDF_MERGE_EVENT_TYPE:
		event.Key = domain.PDF_MERGE_REQUEST_KEY
		return s.repo.PublishEvent(ctx, event)
	case domain.PDF_SPLIT_EVENT_TYPE:
		event.Key = domain.PDF_SPLIT_REQUEST_KEY
		return s.repo.PublishEvent(ctx, event)
	default:
		return domain.TError{
			Code:    http.StatusBadRequest,
			Message: domain.ERROR_MESSAGE_BAD_REQUEST,
			Error:   fmt.Errorf("unknown event type: %s", event.Type),
		}
	}
}
