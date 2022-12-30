package middlewares

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

// LogStruct - logger structure
type LogStruct struct {
	IP        string `json:"ip"`
	URL       string `json:"url"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Duration  int64  `json:"duration"`
	Agent     string `json:"agent"`
	Status    int    `json:"status"`
	Method    string `json:"method"`
}

// Log - logger will print JSON formatted logs onto STDOUT
func Log(ctx *fiber.Ctx) error {
	t := time.Now()
	logInfo := LogStruct{
		IP:        ctx.IP(),
		URL:       ctx.OriginalURL(),
		StartTime: t.String(),
		Method:    string(ctx.Request().Header.Method()),
		Agent:     string(ctx.Request().Header.UserAgent()),
	}
	ctx.Next()
	logInfo.Status = ctx.Response().StatusCode()
	logInfo.EndTime = time.Now().String()
	logInfo.Duration = time.Since(t).Milliseconds()
	logStr, _ := json.Marshal(logInfo)

	logger := log.New(os.Stdout, "", 0)
	logger.Printf("%s", string(logStr))

	return nil
}
