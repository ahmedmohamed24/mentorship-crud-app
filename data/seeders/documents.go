package seeders

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ahmedmohamed24/mentorship-crud-app/internal/modules/document"
	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/config"
	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/db"
)

func SeedDocuments() error {
	cfg, err := config.LoadConfig("./configs/config.yaml")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	dbConn, err := db.NewDBClient(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}

	rand.NewSource(time.Now().UnixNano())
	for i := 1; i <= 1000; i++ {
		doc := document.Document{
			Title:   fmt.Sprintf("Document Title %d", i),
			Author:  fmt.Sprintf("Author %d", rand.Intn(100)+1),
			Content: randomString(100),
		}
		if err := dbConn.Create(&doc).Error; err != nil {
			return fmt.Errorf("failed to insert document #%d: %w", i, err)
		}
	}
	return nil
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
