package main

import (
  "github.com/gin-gonic/gin"
  "html/template"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*.tmpl")

  router.GET("/", func(c *gin.Context) {
    html := template.Must(template.ParseFiles("templates/base.tmpl", "templates/index.tmpl"))
    router.SetHTMLTemplate(html)
    c.HTML(200, "base.tmpl", gin.H{})
  })

  router.Run(":8080")
}
