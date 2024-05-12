package adapter

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"docx-doc-pipeline-srv/src/internal/domain"
	"docx-doc-pipeline-srv/src/internal/port"
)

type TPipelineHandler struct {
	svc port.IPipelineService
}

func (h *TPipelineHandler) Dispatch(ctx *gin.Context) {
	var request TPipelineDispatchRequest

	if err := ctx.ShouldBind(&request); err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	event := domain.TEvent{
		ID:   uuid.New().String(),
		Type: domain.TEventType(request.Type),
	}
	for _, document := range request.Documents {
		event.Documents = append(event.Documents, domain.TDocument{
			ID: document.ID,
			Config: domain.TDocumentConfig{
				Order:    document.Config.Order,
				Rotation: document.Config.Rotation,
			},
		})
	}

	if err := h.svc.Dispatch(ctx, &event); err.Error != nil {
		slog.Error(err.Error.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Message})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, TResponse[TPipelineDispatchResponseBody]{
		Body:     nil,
		Severity: RESPONSE_SEVERITY_SUCCESS,
		Message:  "Event dispatched successfully",
	})
}

func NewPipelineHandler(svc port.IPipelineService) *TPipelineHandler {
	return &TPipelineHandler{svc: svc}
}
