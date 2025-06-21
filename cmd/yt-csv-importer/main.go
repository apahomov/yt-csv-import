package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/yt-csv-importer/internal/config"
	"github.com/user/yt-csv-importer/internal/importer"
	"github.com/user/yt-csv-importer/internal/ui"
)

var cfg config.Config

var rootCmd = &cobra.Command{
	Use:   "yt-csv-importer",
	Short: "Импорт задач в Яндекс.Трекер из CSV файла",
	Long:  `Консольная утилита для автоматизации создания задач в Яндекс.Трекере путем импорта данных из CSV-файла.`,
	Run: func(cmd *cobra.Command, args []string) {
		isInteractive := !cmd.Flags().Changed("token") &&
			!cmd.Flags().Changed("org-id") &&
			!cmd.Flags().Changed("queue") &&
			!cmd.Flags().Changed("file")

		var err error
		if isInteractive {
			cfg, err = ui.AskForConfig()
			if err != nil {
				fmt.Printf("Ошибка получения конфигурации в интерактивном режиме: %v\n", err)
				os.Exit(1)
			}
		}

		if err := cfg.Validate(); err != nil {
			fmt.Printf("Ошибка валидации конфигурации: %v\n", err)
			os.Exit(1)
		}

		imp, err := importer.New(cfg)
		if err != nil {
			fmt.Printf("Ошибка инициализации импортера: %v\n", err)
			os.Exit(1)
		}

		if err := imp.Run(); err != nil {
			fmt.Printf("Ошибка во время импорта: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&cfg.Token, "token", "t", "", "OAuth или IAM токен для доступа к API Яндекс.Трекера")
	rootCmd.Flags().StringVarP(&cfg.OrgID, "org-id", "o", "", "ID организации в Яндекс.Облаке или Яндекс 360")
	rootCmd.Flags().StringVarP(&cfg.Queue, "queue", "q", "", "Ключ очереди в Яндекс.Трекере")
	rootCmd.Flags().StringVarP(&cfg.FilePath, "file", "f", "", "Путь к CSV файлу для импорта")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
