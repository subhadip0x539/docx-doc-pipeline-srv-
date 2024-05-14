package schema

type TResponseSeverity string

const (
	RESPONSE_SEVERITY_SUCCESS TResponseSeverity = "success"
	RESPONSE_SEVEITY_WARNING  TResponseSeverity = "waning"
	RESPONSE_SEVERITY_ERROR   TResponseSeverity = "error"
)

type TResponseBody interface {
	TPipelineDispatchResponseBody
}

type TResponse[T TResponseBody] struct {
	Body     *T                `json:"body"`
	Severity TResponseSeverity `json:"severity"`
	Message  string            `json:"message"`
}
