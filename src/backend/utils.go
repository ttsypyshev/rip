package backend

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Парсинг списка в карту
func ParseList(listStr string) map[string]string {
	result := make(map[string]string)
	lines := strings.Split(listStr, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || !strings.Contains(line, ":") {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		key := strings.TrimSpace(parts[0]) + ":"
		value := strings.TrimSpace(parts[1])
		result[key] = value
	}
	return result
}

// FromEnv собирает DSN строку из переменных окружения
func FromEnv() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return ""
	}

	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
}

// handleError обрабатывает и логирует ошибки, отправляет ответ
func handleError(c *gin.Context, statusCode int, err error, additionalErrs ...error) {
	if err == nil {
		log.Panic("Logging error: empty main error argument")
		return
	}

	c.String(statusCode, err.Error())

	var errorMessages strings.Builder
	errorMessages.WriteString(err.Error())
	for _, additionalErr := range additionalErrs {
		if additionalErr != nil {
			errorMessages.WriteString(": " + additionalErr.Error())
		}
	}

	log.Printf("Error: %s", errorMessages.String())
}
