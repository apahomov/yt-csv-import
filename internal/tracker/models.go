package tracker

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
