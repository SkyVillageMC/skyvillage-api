package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type ConnectedSocket struct {
	Connection *websocket.Conn
	UserId     string
}

var sockets map[string]ConnectedSocket

func InitWs(r *gin.Engine) {

	sockets = make(map[string]ConnectedSocket)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	r.GET("/ws", func(c *gin.Context) {

		if c.Query("id") == "" {
			c.String(401, "please provide user id")
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
		if err != nil {
			log.Println(err.Error())
			return
		}

		sockets[c.Query("id")] = ConnectedSocket{
			Connection: conn,
			UserId:     c.Query("id"),
		}

		defer conn.Close()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				conn.WriteJSON(gin.H{
					"error": "read error",
				})
				break
			}

			err = conn.WriteMessage(1, message)
			if err != nil {
				log.Println("Error during message writing:", err)
				break
			}
		}
	})
}

func SendToUser(id string, json gin.H) error {
	s := sockets[id]

	if s.Connection == nil {
		return nil
	}

	return s.Connection.WriteJSON(json)
}
