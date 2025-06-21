package tracker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const BaseURL = "https://api.tracker.yandex.net"

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
	body, err := json.Marshal(issueData)
	if err != nil {
		return nil, fmt.Errorf("ошибка маршалинга запроса на создание задачи: %w", err)
	}

	url := fmt.Sprintf("%s/v3/issues", BaseURL)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания HTTP-запроса: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("X-Cloud-Org-ID", c.orgID)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения HTTP-запроса: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения тела ответа: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		var apiErr APIError
		if json.Unmarshal(respBody, &apiErr) == nil {
			apiErr.StatusCode = resp.StatusCode
			return nil, &apiErr
		}
		return nil, fmt.Errorf("ошибка API (статус %d): %s", resp.StatusCode, string(respBody))
	}

	var issueResp IssueResponse
	if err := json.Unmarshal(respBody, &issueResp); err != nil {
		return nil, fmt.Errorf("ошибка десериализации ответа на создание задачи: %w", err)
	}

	return &issueResp, nil
}
