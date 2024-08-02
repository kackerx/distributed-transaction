package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}))
	slog.SetDefault(logger)
}

func main() {
	g := gin.Default()
	g.LoadHTMLGlob("template/*")

	g.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", map[string]string{
			"name": "kacker",
		})
	})

	g.GET("/api/greeting", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"age": "12"})
	})

	g.Run(":9999")
}
