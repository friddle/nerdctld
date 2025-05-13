package socket

import (
	"encoding/json"
	"net"
)

// SendCommand 发送命令到socket
func SendCommand(conn net.Conn, args []string) error {
	return json.NewEncoder(conn).Encode(args)
}

// ReceiveCommand 从socket接收命令
func ReceiveCommand(conn net.Conn) ([]string, error) {
	var args []string
	err := json.NewDecoder(conn).Decode(&args)
	return args, err
}
