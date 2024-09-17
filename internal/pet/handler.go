package pet

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

func (h *Handler) CreatePet(c *gin.Context) {
	var pet models.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&pet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la mascota"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mascota creada con éxito"})
}

func (h *Handler) ListPets(c *gin.Context) {
	var pets []models.Pet
	if err := h.DB.Find(&pets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al listar las mascotas"})
		return
	}

	c.JSON(http.StatusOK, pets)
}
