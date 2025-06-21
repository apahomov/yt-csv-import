package tracker

import (
	"fmt"
	"strings"
)

// CreateIssueRequest структура для создания задачи через API.
type CreateIssueRequest struct {
	Summary     string      `json:"summary"`
	Queue       interface{} `json:"queue"`            // может быть string (ключ) или object
	Parent      interface{} `json:"parent,omitempty"` // может быть string (ключ) или object
	Description string      `json:"description,omitempty"`
	Type        interface{} `json:"type,omitempty"` // для создания эпика
}

// IssueResponse структура для ответа от API при создании задачи.
type IssueResponse struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

// SearchRequest структура для поиска задач (в частности, эпиков).
type SearchRequest struct {
	Filter map[string]string `json:"filter"`
}

// APIError представляет структуру ошибки от API Яндекс.Трекера.
type APIError struct {
	StatusCode    int               `json:"statusCode"`
	ErrorMessages []string          `json:"errorMessages"`
	Errors        map[string]string `json:"errors"`
}

// Error возвращает строковое представление ошибки API.
func (e *APIError) Error() string {
	if len(e.ErrorMessages) > 0 {
		return strings.Join(e.ErrorMessages, "; ")
	}
	if len(e.Errors) > 0 {
		var details []string
		for k, v := range e.Errors {
			details = append(details, fmt.Sprintf("%s: %s", k, v))
		}
		return strings.Join(details, ", ")
	}
	return fmt.Sprintf("неизвестная ошибка API со статусом %d", e.StatusCode)
}
