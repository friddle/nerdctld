package main

import (
	"fmt"
	"net"
	"os"

	"github.com/yourusername/nerdctl-socket/internal/socket"
)

const socketPath = "/var/run/nerdctl.socket"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: nerdctlc <command> [args...]")
		os.Exit(1)
	}

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Printf("无法连接到socket: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 发送命令
	if err := socket.SendCommand(conn, os.Args[1:]); err != nil {
		fmt.Printf("发送命令错误: %v\n", err)
		os.Exit(1)
	}

	// 接收并打印结果
	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			break
		}
		os.Stdout.Write(buffer[:n])
	}
}
