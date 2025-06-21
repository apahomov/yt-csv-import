package ui

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/user/yt-csv-importer/internal/config"
)

// AskForConfig запрашивает конфигурацию у пользователя в интерактивном режиме.
func AskForConfig() (config.Config, error) {
	var cfg config.Config

	questions := []*survey.Question{
		{
			Name:     "token",
			Prompt:   &survey.Password{Message: "Введите ваш OAuth или IAM токен:"},
			Validate: survey.Required,
		},
		{
			Name:     "orgID",
			Prompt:   &survey.Input{Message: "Введите ID вашей организации:"},
			Validate: survey.Required,
		},
		{
			Name:     "queue",
			Prompt:   &survey.Input{Message: "Введите ключ очереди в Трекере:"},
			Validate: survey.Required,
		},
		{
			Name:     "filePath",
			Prompt:   &survey.Input{Message: "Укажите путь к CSV файлу:"},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Token    string `survey:"token"`
		OrgID    string `survey:"orgID"`
		Queue    string `survey:"queue"`
		FilePath string `survey:"filePath"`
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		return cfg, err
	}

	return config.Config(answers), nil
}
