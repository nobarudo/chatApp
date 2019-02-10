package main

import (
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  "html/template"
  "gopkg.in/olahol/melody.v1"
  //"github.com/gin-gonic/contrib/renders/multitemplate"
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

  router.POST("/chat", func(c *gin.Context) {
    test := c.PostForm("userName")
    log.Printf("debug:", test)
    html := template.Must(template.ParseFiles("templates/base.tmpl", "templates/chat.tmpl"))
    router.SetHTMLTemplate(html)
    //templates := multitemplate.New()
    //templates.AddFromFiles("contact", "templates/base.tmpl", "templates/chat.tmpl")
    c.HTML(200, "chat.tmpl", gin.H{
      "userName": test,
    })
  })

  router.GET("/chat", func(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently,"/")
  })

  router.GET("/ws", func(c *gin.Context) {
    m.HandleRequest(c.Writer, c.Request)
  })

  m.HandleMessage(func(s *melody.Session, msg []byte) {
    m.Broadcast(msg)
  })

  router.Run(":8080")
}
