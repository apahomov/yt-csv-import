package csvparser

// Record представляет одну строку из CSV-файла.
type Record struct {
	Epic        string // Необязательно
	Summary     string // Обязательно
	Description string // Необязательно
}
