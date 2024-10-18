package tor

import (
	"log"

	"github.com/yawning/bulb"
)

type Client struct {
	conn *bulb.Conn
}

func NewClient(controlPort string) *Client {
	conn, err := bulb.Dial("tcp", controlPort)
	if err != nil {
		log.Fatalf("Failed to connect to Tor: %v", err)
	}

	if err := conn.Authenticate(""); err != nil {
		log.Fatalf("Failed to authenticate with Tor: %v", err)
	}

	return &Client{conn: conn}
}

func (c *Client) ChangeIP() {
	_, err := c.conn.Request("SIGNAL NEWNYM")
	if err != nil {
		log.Fatalf("Failed to change IP: %v", err)
	}
	log.Println("Successfully changed Tor IP")
}
func (c *Client) Close() {
	c.conn.Close()
}
