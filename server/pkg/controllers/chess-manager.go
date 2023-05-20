package controllers

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Manager struct {
	clients ClientList

	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}

func (m *Manager) UseWebsocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (m *Manager) ServeWebsocket(c *websocket.Conn) {
	// c.Locals is added to the *websocket.Conn
	log.Println(c.Locals("allowed"))
	log.Println(c.Params("id"))
	log.Println(c.Query("v"))
	log.Println(c.Cookies("session"))

	var (
		mt  int
		msg []byte
		err error
	)
	if mt, msg, err = c.ReadMessage(); err != nil {
		log.Println("read:", err)
	}
	log.Printf("recv: %s", msg)

	if err = c.WriteMessage(mt, msg); err != nil {
		log.Println("write:", err)
	}
	c.Close()
}

func (m *Manager) addClient(client *Client) {

	m.Lock()
	defer m.Unlock()

	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()

		delete(m.clients, client)
	}
}
