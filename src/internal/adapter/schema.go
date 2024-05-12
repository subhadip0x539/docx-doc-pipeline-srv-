package adapter

type TPipelineDocumentConfig struct {
	Order    int     `json:"order"`
	Rotation float64 `json:"rotation"`
}

type TPipelineDocument struct {
	ID     string                  `json:"id" binding:"required,uuid"`
	Config TPipelineDocumentConfig `json:"config"`
}

type TPipelineDispatchRequest struct {
	Type      string              `json:"type" binding:"required"`
	Documents []TPipelineDocument `json:"documents" binding:"required,min=1,dive"`
}

type TEventMessageBodyDocumentConfig struct {
	Order    int     `json:"order"`
	Rotation float64 `json:"rotation"`
}

type TEventMessageBodyDocument struct {
	ID     string                          `json:"id"`
	Config TEventMessageBodyDocumentConfig `json:"config"`
}

type TEventMessageBody struct {
	ID        string                      `json:"id"`
	Type      string                      `json:"type"`
	Documents []TEventMessageBodyDocument `json:"documents"`
}

type TResponseSeverity string

const (
	RESPONSE_SEVERITY_SUCCESS TResponseSeverity = "success"
	RESPONSE_SEVEITY_WARNING  TResponseSeverity = "waning"
	RESPONSE_SEVERITY_ERROR   TResponseSeverity = "error"
)

type TPipelineDispatchResponseBody struct {
	ID string `json:"id"`
}

type TResponseBody interface {
	TPipelineDispatchResponseBody
}

type TResponse[T TResponseBody] struct {
	Body     *T                `json:"body"`
	Severity TResponseSeverity `json:"severity"`
	Message  string            `json:"message"`
}
