package server

import (
	"net"
)

type TcpServer struct {
	listener *net.TCPListener
}

func New(config *ServerConfig) (*TcpServer, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", config.Addr)

	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		return nil, err
	}

	return &TcpServer{
		listener: listener,
	}, nil
}

func (s *TcpServer) Start() error {
	conn, err := s.listener.Accept()

	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.Write([]byte("hello"))

	if err != nil {
		return err
	}

	return nil
}
