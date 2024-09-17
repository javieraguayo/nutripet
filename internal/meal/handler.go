package meal

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

func (h *Handler) CreateMeal(c *gin.Context) {
	var meal models.Meal
	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&meal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la comida"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comida registrada con éxito"})
}

func (h *Handler) ListMeals(c *gin.Context) {
	var meals []models.Meal
	if err := h.DB.Find(&meals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al listar las comidas"})
		return
	}

	c.JSON(http.StatusOK, meals)
}

func (h *Handler) NextMealNotification(c *gin.Context) {
	var meal models.Meal
	petID := c.Param("petID")

	if err := h.DB.Where("pet_id = ?", petID).Order("time ASC").First(&meal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se encontró la próxima comida"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"next_meal": meal.Time,
		"quantity":  meal.Quantity,
	})
}

func (h *Handler) GetMealsByPetIDAndDateRange(c *gin.Context) {
	petID := c.Param("petID")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	var startTime, endTime time.Time
	var err error

	if startDate != "" {
		startTime, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha de inicio inválida"})
			return
		}
	} else {
		startTime = time.Time{} // Fecha mínima
	}

	if endDate != "" {
		endTime, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha de fin inválida"})
			return
		}
	} else {
		endTime = time.Now() // Fecha actual
	}

	// Llamar al método del modelo para obtener las comidas
	meals, err := models.GetMealsByPetIDAndDateRange(h.DB, petID, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las comidas"})
		return
	}

	c.JSON(http.StatusOK, meals)
}
