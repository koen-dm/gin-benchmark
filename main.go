package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func benchmark(c *gin.Context) {
	var nestedArray [][]int

	if err := c.BindJSON(&nestedArray); err != nil {
		return
	}

	for _, array := range nestedArray {
		go func(x []int) {
			bubbleSort(x)
		}(array)
	}

	c.IndentedJSON(http.StatusOK, nestedArray)
}

func bubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func main() {
	router := gin.Default()
	router.POST("/benchmark", benchmark)
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World"})

	})
	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"message": "method not allowed"})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "method not found"})
	})
	router.Run(":8080")
}
