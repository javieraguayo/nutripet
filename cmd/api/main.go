package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"nutripet/internal/user"
	"nutripet/internal/pet"
	"nutripet/internal/meal"
	"nutripet/internal/plan"
	"nutripet/pkg/database"
)

func main() {
	db, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := gin.Default()

	// Inicialización de los controladores
	userHandler := user.NewHandler(db)
	petHandler := pet.NewHandler(db)
	mealHandler := meal.NewHandler(db)
	planHandler := plan.NewHandler(db)

	// Rutas
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	router.POST("/pets", petHandler.CreatePet)
	router.GET("/pets", petHandler.ListPets)

	router.POST("/meals", mealHandler.CreateMeal)
	router.GET("/meals", mealHandler.ListMeals)
	router.GET("/next-meal/:petID", mealHandler.NextMealNotification)
	router.GET("/meals/:petID", mealHandler.GetMealsByPetIDAndDateRange)

	router.POST("/plans", planHandler.CreatePlan)
	router.GET("/plans/:petID", planHandler.GetPlansByPetID) // Nueva ruta para obtener planes por PetID

	router.Run(":8080")
}
