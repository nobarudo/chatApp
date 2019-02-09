package main

import (
  "log"
  "github.com/gin-gonic/gin"
  "html/template"
  "gopkg.in/olahol/melody.v1"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*.tmpl")
  m := melody.New()

  router.GET("/", func(c *gin.Context) {
    html := template.Must(template.ParseFiles("templates/base.tmpl", "templates/index.tmpl"))
    router.SetHTMLTemplate(html)
    c.HTML(200, "base.tmpl", gin.H{})
  })

  router.GET("/ws", func(c *gin.Context) {
    m.HandleRequest(c.Writer, c.Request)
  })

  m.HandleMessage(func(s *melody.Session, msg []byte) {
    m.Broadcast(msg)
  })

  router.Run(":8080")
}
