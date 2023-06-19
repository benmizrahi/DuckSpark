package common

import "fmt"

func ChunkSlice[T interface{}](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func ConvertToStrings(row []interface{}) []string {
	strings := make([]string, len(row))
	for i, value := range row {
		strings[i] = fmt.Sprintf("%v", value)
	}
	return strings
}
