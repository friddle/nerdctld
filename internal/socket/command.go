package socket

import (
	"encoding/json"
	"net"
)

type Command struct {
	Args []string `json:"args"`
	Env  []string `json:"env"`
	Pwd  string   `json:"pwd"`
}

// SendCommand 发送命令到socket
func SendCommand(conn net.Conn, command *Command) error {
	return json.NewEncoder(conn).Encode(command)
}

// ReceiveCommand 从socket接收命令
func ReceiveCommand(conn net.Conn) (*Command, error) {
	var command Command
	err := json.NewDecoder(conn).Decode(&command)
	return &command, err
}
