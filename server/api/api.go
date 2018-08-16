package api

import (
	"bufio"
	"log"
	"net"
	"time"
)

const (
	DEFAULT_CONNECTION_TIMEOUT time.Duration = 5 * time.Second
	DEFAULT_ADDRESS            string        = "127.0.0.1"
	DEFAULT_PORT               string        = "3001"
)

type ApiServer struct {
	Address         string
	Port            string
	Conn            net.Conn
	TimeoutDuration time.Duration
}

func NewApiServer() *ApiServer {
	return &ApiServer{
		Address:         DEFAULT_ADDRESS,
		Port:            DEFAULT_PORT,
		TimeoutDuration: DEFAULT_CONNECTION_TIMEOUT,
	}
}

func (a *ApiServer) Listen() error {
	ipAddr := a.Address + ":" + a.Port

	log.Printf("Starting API Server on %s:%s", a.Address, a.Port)

	listener, err := net.Listen("tcp", ipAddr)
	if err != nil {
		return err
	}

	defer func() {
		listener.Close()
		log.Println("API Server stopped")
	}()

	for {
		a.Conn, err = listener.Accept()
		if err != nil {
			log.Println(err)
			break
		}

		go a.HandleConnection()
	}
	return nil
}

func (a *ApiServer) HandleConnection() {
	log.Println("Handling new API connection...")

	defer func() {
		log.Println("Closing API connection")
		a.Conn.Close()
	}()

	timeoutDuration := a.TimeoutDuration

	bufReader := bufio.NewReader(a.Conn)

	for {
		a.Conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("%s", bytes)
	}
}
