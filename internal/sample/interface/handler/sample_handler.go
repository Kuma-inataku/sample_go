package handler

import (
	"gin-api-sample/internal/sample/domain/model"
	"gin-api-sample/internal/sample/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct {
	sampleUsecase usecase.SampleUsecase
}

func NewSampleHandler() *SampleHandler {
	return &SampleHandler{
		sampleUsecase: usecase.NewSampleUsecase(),
	}
}

func (h *SampleHandler) GetSamples(c *gin.Context) {
	samples, err := h.sampleUsecase.GetSamples()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, samples)
}

func (h *SampleHandler) CreateSample(c *gin.Context) {
	var sample model.Sample
	if err := c.ShouldBindJSON(&sample); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.sampleUsecase.CreateSample(&sample); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sample)
}

func (h *SampleHandler) UpdateSample(c *gin.Context) {
	var sample model.Sample
	if err := c.ShouldBindJSON(&sample); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.sampleUsecase.UpdateSample(&sample); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sample)
}

func (h *SampleHandler) DeleteSample(c *gin.Context) {
	id := c.Param("id")
	if err := h.sampleUsecase.DeleteSample(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
