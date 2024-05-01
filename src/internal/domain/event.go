package domain

type TEventType string

const (
	PdfMergeEventTag  TEventType = "PDFMERGE"
	PdfSplitEventType TEventType = "PDFSPLIT"
)

type TEventKey string

const (
	PdfMergeRequestEventKey TEventKey = "pdf.merge.request.#"
	PdfSplitRequestEventKey TEventKey = "pdf.split.request.#"
)

type TEvent struct {
	ID        string
	Type      TEventType
	Key       TEventKey
	Documents []TDocument
}
