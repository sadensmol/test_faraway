package internal

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"sadensmol/go/test_faraway/internal/server/storage"
	"sadensmol/go/test_faraway/pkg/service"
	"sadensmol/go/test_faraway/pkg/utils"
	"sync"
	"syscall"
)

type Server struct {
	port         int
	powService   *service.POWService
	quoteStorage *storage.QuoteStorage
	wg           *sync.WaitGroup
}

func NewServer(port int) Server {
	return Server{
		port:         port,
		powService:   service.NewPOWService(),
		quoteStorage: storage.NewQuoteStorage(),
		wg:           &sync.WaitGroup{},
	}
}
func (s Server) Start() {
	log.Printf("starting the server on port %d\n", s.port)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Printf("cannot read from connection %v\n", err)
				continue
			}

			select {
			//stop accepting connections if context is done
			case <-ctx.Done():
				return
			default:
				s.wg.Add(1)
				go s.handleConnection(conn)
			}
		}
	}()

	<-ctx.Done()
	// wait until all connections got processed
	s.wg.Wait()
	log.Println("server gracefully stopped!")
}

func (s Server) handleConnection(conn net.Conn) {
	log.Printf("handling connection from %v\n", conn.RemoteAddr())
	defer s.wg.Done()
	defer conn.Close()

	req, err := s.powService.GenerateRequest()
	if err != nil {
		log.Printf("cannot generate request %v", err)
		return
	}

	err = utils.WriteBytes(conn, []byte(*req))
	if err != nil {
		log.Printf("cannot send request to the client %v", err)
		return
	}

	res, err := utils.ReadBytes(conn)
	if err != nil {
		log.Printf("cannot send request to the client %v\n", err)
		return
	}

	if err = s.powService.Check(string(res)); err != nil {
		log.Printf("hash %s is not valid for %v\n", string(res), conn.RemoteAddr())
		return
	}

	q, err := s.quoteStorage.GetRandom()
	if err != nil {
		log.Printf("cannot fetch quote from storage %v\n", err)
		return
	}

	utils.WriteBytes(conn, []byte(*q))
	if err != nil {
		log.Printf("cannot send quote to the client %v\n", err)
		return
	}

}
