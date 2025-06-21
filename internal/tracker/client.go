package tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client для работы с API Yandex.Tracker.
type Client struct {
	httpClient *http.Client
	token      string
	orgID      string
}

// NewClient создает нового клиента.
func NewClient(token, orgID string) (*Client, error) {
	if token == "" || orgID == "" {
		return nil, fmt.Errorf("токен и ID организации не должны быть пустыми")
	}
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		token:      token,
		orgID:      orgID,
	}, nil
}

// FindEpic ищет эпик по имени в указанной очереди.
// Возвращает ID эпика или пустую строку, если не найден.
func (c *Client) FindEpic(name, queue string) (string, error) {
	// Заглушка. Реализация будет в следующей фазе.
	fmt.Printf("[DEBUG] Поиск эпика: Имя='%s', Очередь='%s'\n", name, queue)

	searchReq := SearchRequest{
		Filter: map[string]string{
			"queue":   queue,
			"summary": name,
			"type":    "epic",
		},
	}

	body, err := json.Marshal(searchReq)
	if err != nil {
		return "", fmt.Errorf("ошибка маршалинга запроса поиска эпика: %w", err)
	}

	// Логика выполнения HTTP запроса будет добавлена позже.
	_ = body

	// Пока возвращаем пустую строку для имитации.
	return "", nil
}

// CreateIssue создает задачу или эпик.
// Для эпика в issueData.Type нужно передать "epic".
func (c *Client) CreateIssue(issueData CreateIssueRequest) (*IssueResponse, error) {
	// Заглушка. Реализация будет в следующей фазе.
	fmt.Printf("[DEBUG] Создание задачи: %+v\n", issueData)

	body, err := json.Marshal(issueData)
	if err != nil {
		return nil, fmt.Errorf("ошибка маршалинга запроса на создание задачи: %w", err)
	}

	// Логика выполнения HTTP запроса будет добавлена позже.
	_ = body

	return &IssueResponse{
		ID:  fmt.Sprintf("mock-id-%d", time.Now().UnixNano()),
		Key: fmt.Sprintf("%s-MOCK", issueData.Queue.(map[string]string)["key"]),
	}, nil
}
