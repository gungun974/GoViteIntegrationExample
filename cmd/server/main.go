package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"gungun974.com/viteintegration/internal/logger"
	"gungun974.com/viteintegration/internal/routes"

	"github.com/joho/godotenv"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(
	w io.Writer,
	name string,
	data interface{},
) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	err := godotenv.Load()

	if err == nil {
		logger.MainLogger.Info("Loading .env file")
	}

	router := routes.MainRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	logger.MainLogger.Infof("🌍 WebServer is running at http://127.0.0.1:%s", port)

	err = http.ListenAndServe(":"+port, router)

	if err != nil {
		logger.MainLogger.Fatalf("Failed to start HTTP server on port %s : %v", port, err)
	}
}
