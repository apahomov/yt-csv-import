package config

import "fmt"

// Config хранит конфигурацию приложения.
type Config struct {
	Token    string
	OrgID    string
	Queue    string
	FilePath string
}

// Validate проверяет, что все обязательные поля конфигурации заполнены.
func (c *Config) Validate() error {
	if c.Token == "" {
		return fmt.Errorf("токен не указан")
	}
	if c.OrgID == "" {
		return fmt.Errorf("ID организации не указан")
	}
	if c.Queue == "" {
		return fmt.Errorf("ключ очереди не указан")
	}
	if c.FilePath == "" {
		return fmt.Errorf("путь к файлу не указан")
	}
	return nil
}
