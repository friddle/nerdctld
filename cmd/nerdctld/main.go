package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/friddle/nerdctld/internal/socket"
)

const socketPath = "/var/run/nerdctl.socket"

func main() {
	// 确保socket文件不存在
	os.Remove(socketPath)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("无法创建Unix Domain Socket: %v", err)
	}
	defer listener.Close()

	// 设置socket权限
	if err := os.Chmod(socketPath, 0666); err != nil {
		log.Fatalf("无法设置socket权限: %v", err)
	}

	// 处理信号以优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		os.Remove(socketPath)
		os.Exit(0)
	}()

	log.Printf("nerdctld 服务启动在 %s", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("接受连接错误: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	cmd, err := socket.ReceiveCommand(conn)
	if err != nil {
		log.Printf("读取命令错误: %v", err)
		return
	}

	// 执行nerdctl命令
	os.Chdir(cmd.Pwd)
	nerdctlCmd := exec.Command("nerdctl", cmd.Args...)
	nerdctlCmd.Env = append(nerdctlCmd.Env, cmd.Env...)
	nerdctlCmd.Stdout = conn
	nerdctlCmd.Stderr = conn

	if err := nerdctlCmd.Run(); err != nil {
		log.Printf("执行命令错误: %v", err)
	}
}
