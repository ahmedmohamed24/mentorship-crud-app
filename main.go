/*
Copyright © 2025 Ahmed Mohamed <ahmedmohamed24.dev@gmail.com>
*/
package main

import (
	"os"

	"github.com/ahmedmohamed24/mentorship-crud-app/cmd"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
}
func main() {
	cmd.Execute()
}
