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

		if record.Epic != "" {
			epicID, ok := i.epicCache[record.Epic]
			if !ok {
				// Эпик не в кэше, ищем или создаем
				foundEpicID, err := i.tracker.FindEpic(record.Epic, i.cfg.Queue)
				if err != nil {
					fmt.Printf("Ошибка поиска эпика '%s': %v\n", record.Epic, err)
					continue
				}

				if foundEpicID != "" {
					epicID = foundEpicID
				} else {
					// Создаем эпик
					fmt.Printf("Эпик '%s' не найден, создаем новый...\n", record.Epic)
					epicReq := tracker.CreateIssueRequest{
						Summary: record.Epic,
						Queue:   map[string]string{"key": i.cfg.Queue},
						Type:    "epic",
					}
					resp, err := i.tracker.CreateIssue(epicReq)
					if err != nil {
						fmt.Printf("Ошибка создания эпика '%s': %v\n", record.Epic, err)
						continue
					}
					epicID = resp.ID
					fmt.Printf("Успешно создан эпик: %s\n", resp.Key)
				}
				i.epicCache[record.Epic] = epicID
			}
			createReq.Parent = map[string]string{"id": epicID}
		}

		resp, err := i.tracker.CreateIssue(createReq)
		if err != nil {
			fmt.Printf("Ошибка создания задачи '%s': %v\n", record.Summary, err)
			continue
		}
		if record.Epic != "" {
			fmt.Printf("Успешно создана задача: %s (%s) в эпике '%s'\n", resp.Key, record.Summary, record.Epic)
		} else {
			fmt.Printf("Успешно создана задача: %s (%s)\n", resp.Key, record.Summary)
		}
	}

	return nil
}
