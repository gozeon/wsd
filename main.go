package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var logs []string
var js string

//go:embed ui
var f embed.FS

func main() {
	r := gin.Default()
	ui := template.Must(template.New("").ParseFS(f, "ui/*.html"))
	r.SetHTMLTemplate(ui)
	// r.StaticFS("/ui", http.FS(f))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/dash", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "dash.html", nil)
	})

	r.GET("/log", func(ctx *gin.Context) {
		log, hasLog := ctx.GetQuery("log")
		if hasLog {
			logs = append(logs, log)
		}
		ctx.JSONP(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	r.GET("/slog", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"logs": logs,
		})
	})

	r.POST("/code", func(ctx *gin.Context) {
		fmt.Println(js)
		if js != "" {
			ctx.AbortWithStatus(http.StatusAlreadyReported)
			return
		}
		var codeBody map[string]interface{}
		if err := ctx.ShouldBindJSON(&codeBody); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		js = codeBody["code"].(string)

		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	r.GET("/code", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, gin.H{
			"js": js,
		})
		js = ""
	})

	r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
