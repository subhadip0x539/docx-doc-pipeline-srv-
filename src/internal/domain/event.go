package domain

type TEventType string

const (
	PDF_MERGE_EVENT_TYPE TEventType = "PDFMERGE"
	PDF_SPLIT_EVENT_TYPE TEventType = "PDFSPLIT"
)

type TEventKey string

const (
	PDF_MERGE_REQUEST_KEY TEventKey = "pdf.merge.request.#"
	PDF_SPLIT_REQUEST_KEY TEventKey = "pdf.split.request.#"
)

type TEvent struct {
	ID        string
	Type      TEventType
	Key       TEventKey
	Documents []TDocument
}
