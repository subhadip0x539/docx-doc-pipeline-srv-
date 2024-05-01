package domain

type TDocumentConfig struct {
	Order    int
	Rotation float64
}

type TDocument struct {
	ID     string
	Config TDocumentConfig
}
