package main

import (
	"encoding/csv"
	"net/http"
	"os"
	"sort"

	"github.com/gin-gonic/gin"
)

// Func membaca file csv
func readCSVFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	uniqueWords := make(map[string]bool)
	var words []string

	for _, record := range records {
		for _, word := range record {
			if _, exists := uniqueWords[word]; !exists {
				uniqueWords[word] = true
				words = append(words, word)
			}
		}
	}

	// Sort the words alphabetically
	sort.Strings(words)

	return words, nil
}

func main() {
	router := gin.Default()

	router.GET("/vocabulary", func(c *gin.Context) {
		filePath := "vocabulary.csv" // Path ke file CSV
		words, err := readCSVFile(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"vocabulary": words,
		})
	})

	router.Run(":8080")
}
