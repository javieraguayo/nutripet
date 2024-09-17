package plan

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nutripet/pkg/models"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) CreatePlan(c *gin.Context) {
	var plan models.Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el plan de alimentación"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plan de alimentación creado con éxito"})
}

func (h *Handler) GetPlansByPetID(c *gin.Context) {
	var plans []models.Plan
	petID := c.Param("petID")

	if err := h.DB.Where("pet_id = ?", petID).Find(&plans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los planes"})
		return
	}

	c.JSON(http.StatusOK, plans)
}
