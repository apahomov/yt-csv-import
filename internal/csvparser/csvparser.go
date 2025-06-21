package csvparser

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Parse читает CSV-файл и возвращает слайс записей.
// Предполагается, что файл не содержит заголовка.
func Parse(filePath string) ([]Record, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s: %w", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// CSV файл имеет 3 колонки: Эпик,Заголовок,Описание
	reader.FieldsPerRecord = 3

	rawRecords, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать CSV файл: %w", err)
	}

	var records []Record
	for i, row := range rawRecords {
		if row[1] == "" {
			return nil, fmt.Errorf("заголовок задачи (вторая колонка) не может быть пустым в строке %d", i+1)
		}
		records = append(records, Record{
			Epic:        row[0],
			Summary:     row[1],
			Description: row[2],
		})
	}

	return records, nil
}
