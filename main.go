package main

import (
  "log"
  "net/http"
  "html/template"
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
  "github.com/gin-gonic/contrib/sessions"
  //"github.com/gin-gonic/contrib/renders/multitemplate"
)

func main() {
  router := gin.Default()
  router.Static("/assets", "./assets")
  router.LoadHTMLGlob("templates/*.tmpl")
  m := melody.New()
  store := sessions.NewCookieStore([]byte("secret"))
  router.Use (sessions.Sessions("session", store))

  router.GET("/", func(c *gin.Context) {
    html := template.Must(template.ParseFiles("templates/base.tmpl", "templates/index.tmpl"))
    router.SetHTMLTemplate(html)
    c.HTML(200, "base.tmpl", gin.H{})
  })

  router.POST("/chat", func(c *gin.Context) {
    name := c.PostForm("userName")
    html := template.Must(template.ParseFiles("templates/base.tmpl", "templates/chat.tmpl"))
    router.SetHTMLTemplate(html)

    session := sessions.Default(c)
    session.Set("name", name)
    session.Save()

    log.Println("新しいユーザが参加しました. userName:", name)
    systemMsg := []byte("[システム] > "+name+"が参加しました. ")
    m.Broadcast(systemMsg)
    //templates := multitemplate.New()
    //templates.AddFromFiles("contact", "templates/base.tmpl", "templates/chat.tmpl")
    c.HTML(200, "chat.tmpl", gin.H{
      "userName": name,
    })
  })

  router.GET("/chat", func(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently,"/")
  })

  router.GET("/ws", func(c *gin.Context) {
    session := sessions.Default(c)
    name := session.Get("name")
    temp := name.(string)
    systemMsg := []byte("[システム] > "+temp+"が退出しました. ")
    m.HandleRequest(c.Writer, c.Request)
    log.Println("debug:", name)
    m.Broadcast(systemMsg)
  })

  m.HandleMessage(func(s *melody.Session, msg []byte) {
    log.Printf("messege:%s", msg)
    m.Broadcast(msg)
  })

  router.Run(":8080")
}
