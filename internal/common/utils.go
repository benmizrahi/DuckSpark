package common

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

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
func Deserialize(serializedData []byte) ([]interface{}, error) {
	var data []interface{}
	buf := bytes.NewBuffer(serializedData)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Serialize(data []interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ConvertToStrings(row []interface{}) []string {
	strings := make([]string, len(row))
	for i, value := range row {
		strings[i] = fmt.Sprintf("%v", value)
	}
	return strings
}

func CalculateTotalDataSize(data [][]interface{}) int {
	totalSize := 0

	for _, partition := range data {
		for _, element := range partition {
			elementSize := getElementSize(element)
			totalSize += elementSize
		}
	}

	return totalSize
}

func getElementSize(element interface{}) int {
	switch v := element.(type) {
	case string:
		return len(v)
	case []byte:
		return len(v)
	// Add more cases for other types as needed
	default:
		// Handle other data types or return an appropriate default size
		return 0
	}
}

func ConvertStringSliceToInterfaceSlice(data [][]string) [][]interface{} {
	convertedData := make([][]interface{}, len(data))

	for i, row := range data {
		convertedRow := make([]interface{}, len(row))
		for j, value := range row {
			convertedRow[j] = interface{}(value)
		}
		convertedData[i] = convertedRow
	}

	return convertedData
}
