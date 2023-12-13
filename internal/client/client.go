package internal

import (
	"log"
	"net"
	"sadensmol/go/test_faraway/pkg/service"
	"sadensmol/go/test_faraway/pkg/utils"
)

type Client struct {
	url        string
	powService *service.POWService
}

func NewClient(url string) Client {
	return Client{
		url:        url,
		powService: service.NewPOWService(),
	}
}
func (c Client) Start() {
	log.Printf("starting the client and connecting to server %s\n", c.url)

	// Connect to the server
	conn, err := net.Dial("tcp", c.url)
	if err != nil {
		log.Fatalf("Failed to connect to server: %s\n", err)
	}
	defer conn.Close()

	req, err := utils.ReadBytes(conn)
	if err != nil {
		log.Fatalf("failed to read request from server: %s\n", err)
	}

	h, err := c.powService.Mint(string(req))
	if err != nil {
		log.Fatalf("failed to mint a hash: %s\n", err)
	}

	err = utils.WriteBytes(conn, []byte(*h))
	if err != nil {
		log.Fatalf("failed to write response to the server: %s\n", err)
	}

	q, err := utils.ReadBytes(conn)
	if err != nil {
		log.Fatalf("failed to read quote from server: %s\n", err)
	}

	log.Println(string(q))
}
