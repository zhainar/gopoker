package main

import (
	"flag"
	"fmt"
	"github.com/zhainar/gopoker/internal/messaging"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	log "github.com/sirupsen/logrus"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	r := gin.Default()

	r.GET("/echo", echo)
	r.GET("/", home)
	r.Static("/static", "web/static")

	err := r.Run(*addr)

	log.Fatal(err)
}

func home(c *gin.Context) {
	var homeTemplate = template.Must(template.ParseFiles("web/index.html"))
	homeTemplate.Execute(c.Writer, fmt.Sprintf("ws://%s/echo", c.Request.Host))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  128,
	WriteBufferSize: 128,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var dispatcher = messaging.NewDispatcher()

func echo(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("msg: %s", msg)

		message, err := messaging.NewMessage(msg)

		if err !=nil {
			conn.WriteJSON(map[string]string{
				"error": err.Error(),
			})
			break
		}

		dispatcher.Handle(conn, message)
	}
}
