package importer

import (
	"fmt"

	"github.com/user/yt-csv-importer/internal/config"
	"github.com/user/yt-csv-importer/internal/csvparser"
	"github.com/user/yt-csv-importer/internal/tracker"
)

// Importer координирует процесс импорта.
type Importer struct {
	cfg       config.Config
	tracker   *tracker.Client
	epicCache map[string]string // Кэш для epicName -> epicID
}

// New создает новый экземпляр импортера.
func New(cfg config.Config) (*Importer, error) {
	trackerClient, err := tracker.NewClient(cfg.Token, cfg.OrgID)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания клиента трекера: %w", err)
	}

	return &Importer{
		cfg:       cfg,
		tracker:   trackerClient,
		epicCache: make(map[string]string),
	}, nil
}

// Run запускает процесс импорта.
func (i *Importer) Run() error {
	records, err := csvparser.Parse(i.cfg.FilePath)
	if err != nil {
		return fmt.Errorf("ошибка парсинга CSV: %w", err)
	}

	fmt.Printf("Найдено %d записей для импорта.\n", len(records))

	for _, record := range records {
		createReq := tracker.CreateIssueRequest{
			Summary:     record.Summary,
			Description: record.Description,
			Queue:       map[string]string{"key": i.cfg.Queue},
		}

		// TODO: Реализовать поиск/создание эпика и добавление его в Parent

		resp, err := i.tracker.CreateIssue(createReq)
		if err != nil {
			fmt.Printf("Ошибка создания задачи '%s': %v\n", record.Summary, err)
			continue
		}
		fmt.Printf("Успешно создана задача: %s (%s)\n", resp.Key, record.Summary)
	}

	return nil
}
