package algorithms

import "github.com/MathisBurger/AnalAsia/internal/database/models"

// This is an bubbles implementation for a
// specific type
func BubbleSortWordModels(arr []models.WordModel) []models.WordModel {
	n := len(arr)

	for i := 0; i < (n - 1); i++ {
		for j := 0; j < (n - i - 1); j++ {
			if arr[j].Counter > arr[j+1].Counter {
				cache := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = cache
			}
		}
	}

	return arr
}
