package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Model struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Accuracy    float64 `json:"accuracy"`
	Author      string  `json:"author"`
	Downloads   int     `json:"downloads"`
	Rating      float64 `json:"rating"`
}

const filePath = "data/models.json"

func loadModels() []Model {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return []Model{}
	}

	var models []Model
	json.Unmarshal(data, &models)
	return models
}

func saveModels(models []Model) {
	data, _ := json.MarshalIndent(models, "", "  ")
	os.WriteFile(filePath, data, 0644)
}
func main() {

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.GET("/models", func(c *gin.Context) {
		c.JSON(200, loadModels())
	})
	r.POST("/models", func(c *gin.Context) {

		models := loadModels()

		var newModel Model
		c.BindJSON(&newModel)

		newModel.ID = len(models) + 1
		newModel.Downloads = 0
		newModel.Rating = 5.0

		models = append(models, newModel)
		saveModels(models)

		c.JSON(201, newModel)
	})
	r.DELETE("/models/:id", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))

		models := loadModels()
		var updated []Model

		for _, m := range models {
			if m.ID != id {
				updated = append(updated, m)
			}
		}

		saveModels(updated)

		c.JSON(200, gin.H{
			"message": "deleted",
		})
	})

	r.PATCH("/models/:id/download", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))

		models := loadModels()

		for i, m := range models {
			if m.ID == id {
				models[i].Downloads += 1
			}
		}

		saveModels(models)

		c.JSON(200, gin.H{
			"message": "downloaded",
		})
	})

	r.PATCH("/models/:id/rate", func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))

		models := loadModels()

		for i, m := range models {
			if m.ID == id {
				models[i].Rating += 0.1
			}
		}

		saveModels(models)

		c.JSON(200, gin.H{
			"message": "rated",
		})
	})

	r.Run(":8080")
}
