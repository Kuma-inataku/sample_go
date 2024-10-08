package usecase

import (
	"gin-api-sample/internal/sample/domain/model"
	"gin-api-sample/internal/sample/infrastructure/db"
)

type SampleUsecase interface {
	GetSamples() ([]model.Sample, error)
	CreateSample(sample *model.Sample) error
	UpdateSample(sample *model.Sample) error
	DeleteSample(id string) error
}

type sampleUsecase struct {
	sampleRepository db.SampleRepository
}

func NewSampleUsecase() SampleUsecase {
	return &sampleUsecase{
		sampleRepository: db.NewSampleRepository(),
	}
}

func (u *sampleUsecase) GetSamples() ([]model.Sample, error) {
	return u.sampleRepository.GetSamples()
}

func (u *sampleUsecase) CreateSample(sample *model.Sample) error {
	return u.sampleRepository.CreateSample(sample)
}

func (u *sampleUsecase) UpdateSample(sample *model.Sample) error {
	return u.sampleRepository.UpdateSample(sample)
}

func (u *sampleUsecase) DeleteSample(id string) error {
	return u.sampleRepository.DeleteSample(id)
}
