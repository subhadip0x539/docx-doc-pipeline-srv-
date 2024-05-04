package repo

import (
	"context"
	"fmt"

	"encoding/json"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"

	"docx-doc-pipeline-srv/src/internal/adapter"
	"docx-doc-pipeline-srv/src/internal/domain"
)

type TPipelineRepo struct {
	exchange string
	amqp     *amqp.Channel
}

func NewPipelineRepo(channel *amqp.Channel, exchange string) *TPipelineRepo {
	return &TPipelineRepo{amqp: channel, exchange: exchange}
}

func (r *TPipelineRepo) PublishEvent(ctx context.Context, event *domain.TEvent) domain.TError {
	message := adapter.TEventMessageBody{
		ID:   event.ID,
		Type: string(event.Type),
	}

	for _, document := range event.Documents {
		message.Documents = append(message.Documents, adapter.TEventMessageBodyDocument{
			ID: document.ID,
			Config: adapter.TEventMessageBodyDocumentConfig{
				Order:    document.Config.Order,
				Rotation: document.Config.Rotation,
			},
		})
	}

	body, err := json.Marshal(message)
	if err != nil {
		return domain.TError{
			Code:    http.StatusInternalServerError,
			Message: domain.INTERNAL_SERVER_ERROR,
			Error:   fmt.Errorf("failed to marshal message body: %w", err),
		}
	}

	err = r.amqp.PublishWithContext(ctx, r.exchange, string(event.Key), false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return domain.TError{
			Code:    http.StatusInternalServerError,
			Message: domain.INTERNAL_SERVER_ERROR,
			Error:   fmt.Errorf("failed to publish event: %w", err),
		}
	}

	return domain.TError{}
}
