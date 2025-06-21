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
	// Разрешаем переменное количество полей (2-3 колонки)
	reader.FieldsPerRecord = -1

	rawRecords, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать CSV файл: %w", err)
	}

	var records []Record
	for i, row := range rawRecords {
		if len(row) < 2 {
			return nil, fmt.Errorf("недостаточно колонок в строке %d (минимум 2: Эпик, Заголовок)", i+1)
		}
		if row[1] == "" {
			return nil, fmt.Errorf("заголовок задачи (вторая колонка) не может быть пустым в строке %d", i+1)
		}

		record := Record{
			Epic:    row[0],
			Summary: row[1],
		}

		// Описание опционально (третья колонка)
		if len(row) > 2 {
			record.Description = row[2]
		}

		records = append(records, record)
	}

	return records, nil
}
