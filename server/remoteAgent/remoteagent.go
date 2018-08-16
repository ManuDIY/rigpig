package remoteAgent

import (
	"bufio"
	"log"
	"net"
	"time"
)

const (
	DEFAULT_CONNECTION_TIMEOUT time.Duration = 5 * time.Second
	DEFAULT_ADDRESS            string        = "127.0.0.1"
	DEFAULT_PORT               string        = "3000"
)

type RemoteAgentServer struct {
	Address         string
	Port            string
	Conn            net.Conn
	TimeoutDuration time.Duration
}

func NewRemoteAgentServer() *RemoteAgentServer {
	return &RemoteAgentServer{
		Address:         DEFAULT_ADDRESS,
		Port:            DEFAULT_PORT,
		TimeoutDuration: DEFAULT_CONNECTION_TIMEOUT,
	}
}

func (r *RemoteAgentServer) Listen() error {
	ipAddr := r.Address + ":" + r.Port

	log.Printf("Starting Remote Agent Server on %s:%s", r.Address, r.Port)

	listener, err := net.Listen("tcp", ipAddr)
	if err != nil {
		return err
	}

	defer func() {
		listener.Close()
		log.Println("API Server stopped")
	}()

	for {
		r.Conn, err = listener.Accept()
		if err != nil {
			log.Println(err)
			break
		}

		go r.HandleConnection()
	}
	return nil
}

func (r *RemoteAgentServer) HandleConnection() {
	log.Println("Handling new Remote Agent connection...")

	defer func() {
		log.Println("Closing Remote Agent connection")
		r.Conn.Close()
	}()

	timeoutDuration := r.TimeoutDuration

	bufReader := bufio.NewReader(r.Conn)

	for {
		r.Conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("%s", bytes)
	}
}
