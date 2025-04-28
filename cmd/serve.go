/*
Copyright © 2025 Ahmed Mohamed <ahmedmohamed24.dev@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/config"
	"github.com/ahmedmohamed24/mentorship-crud-app/internal/pkg/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the server",
	Long:  `Run the server`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
func serve() {
	cfg, err := config.LoadConfig("./configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Infof("HTTP server running on %v", fmt.Sprintf(":%v", cfg.Server.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-sigChan
	log.Info("Received shutdown signal")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	select {
	case <-shutdownCtx.Done():
		log.Warn("Shutdown timeout exceeded")
	default:
		log.Info("Graceful shutdown completed")
	}
}
