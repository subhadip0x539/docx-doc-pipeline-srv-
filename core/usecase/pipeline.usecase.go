package usecase

import "docx-doc-pipeline-srv/core/repo"

type IPipelineUseCase interface {
}

type PipelineUseCase struct {
	mongoRepo  repo.IMongoRepo
	rabbitRepo repo.IRabbitRepo
}

func (uc *PipelineUseCase) Dispatch() {
}

func NewPipelineUseCase(mongoRepo repo.IMongoRepo, rabbitRepo repo.IRabbitRepo) IPipelineUseCase {
	return &PipelineUseCase{mongoRepo: mongoRepo, rabbitRepo: rabbitRepo}
}
