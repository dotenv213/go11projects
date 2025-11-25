package main

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/contrib/websocket"
)

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

var clients = make(map[*websocket.Conn]bool)

var broadcast = make(chan Message)

var mutex = sync.Mutex{}

func main() {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		mutex.Lock()
		clients[c] = true
		mutex.Unlock()
		
		log.Println("New client connected")
		defer func() {
			mutex.Lock()
			delete(clients, c)
			mutex.Unlock()
			c.Close()
			log.Println("Client disconnected")
		}()

		for {
			var msg Message
			if err := c.ReadJSON(&msg); err != nil {
				log.Println("Read error:", err)
				break 
			}

			broadcast <- msg
		}
	}))

	go handleMessages()

	app.Static("/", "./static")

	log.Fatal(app.Listen(":3000"))
}

func handleMessages() {
	for {
		msg := <-broadcast

		mutex.Lock()
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Printf("Write error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}