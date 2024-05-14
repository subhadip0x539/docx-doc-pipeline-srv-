package schema

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

type TPipelineDispatchResponseBody struct {
	ID string `json:"id"`
}
