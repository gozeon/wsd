package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Log struct {
	Msg  string
	Time time.Time
}

var logs []Log
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
			logs = append(logs, Log{Msg: log, Time: time.Now()})
		}
		ctx.JSONP(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	r.GET("/slog", func(ctx *gin.Context) {
		if len(logs) == 0 {
			ctx.String(http.StatusOK, "")
			return
		}
		log := logs[0]
		ctx.String(http.StatusOK, fmt.Sprintf("<b>%2d:%2d:%2d</b>  %s", log.Time.Hour(), log.Time.Minute(), log.Time.Second(), log.Msg))
		logs = logs[1:]
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
