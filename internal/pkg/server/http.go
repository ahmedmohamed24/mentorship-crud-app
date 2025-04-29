package server

import (
	"fmt"
	"net/http"

	"github.com/ahmedmohamed24/mentorship-crud-app/internal/modules/document"
	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/config"
	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg *config.Config) (*http.Server, error) {
	db, err := db.NewDBClient(cfg)
	if err != nil {
		return nil, err
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	documentService := document.NewService(document.NewRepository(db))
	document.RegisterHandlers(router.Group("/api/v1/document"), documentService)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.Server.Port),
		Handler:      router.Handler(),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}
	return srv, nil
}
