package main

import (
	"fmt"
	"net/http"
	"os"

	"prabha-logs/log-wrapper/internal/config/globals"
	"prabha-logs/log-wrapper/internal/handlers"
	"prabha-logs/log-wrapper/internal/services/queue"
)

// @title Log Wrapper API
// @version 1.0
// @description Log Wrapper service is used to send logs to kafka queue which is consumed by graylog.
// @securityDefinitions.basic BasicAuth
// @host staging-log.lynk.co.in
// @BasePath /
func main() {
	go queue.SendLogs()
	//globals.Logger.Info("Port: ", config.PORT)
	port := "9008"
	fmt.Println("Port: ", port)
	err := http.ListenAndServe(":"+port, handlers.GetRouter())
	if err != nil {
		os.Exit(1)
		globals.Logger.Fatalf("Unable to start the server. Err:", err)
	}
}
