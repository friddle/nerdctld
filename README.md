# nerdctl run with socket

## 中文说明

nerdctl 是一个兼容 Docker CLI 的 containerd 客户端工具。由于其需要访问系统级资源和特权操作，在容器化环境（如 Docker-in-Docker）中直接运行 nerdctl 存在限制。

本项目提供了一种解决方案，通过 socket 转发的方式，允许在容器内通过宿主机执行 nerdctl 命令，从而实现在受限环境中使用 nerdctl 的功能。

## English Description

nerdctl is a Docker-compatible CLI for containerd. Due to its requirements for system-level access and privileged operations, running nerdctl directly in containerized environments (such as Docker-in-Docker) has limitations.

This project provides a solution by forwarding commands through a socket to the host machine, enabling the execution of nerdctl commands from within containers while actually running them on the host system.

## Thanks
thanks cursor/claude 让我10分钟完成这个功能。

## 安装说明

1. 解压服务端和客户端二进制文件：
```bash
tar xvf nerdctl.tar.gz -C /usr/local/bin
tar xvf nerdctld.tar.gz -C /usr/local/bin
```

2. 配置 systemd 服务：
```bash
sudo cp nerdctld.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable nerdctld
sudo systemctl start nerdctld
```

## 使用方法

1. 在容器内使用 nerdctlc 命令（与 nerdctl 命令用法相同）：
```bash
nerdctlc run -ti alpine
```

2. 在 Docker 运行时挂载 socket：
```bash
nerdctl run -ti \
  -v /var/run/nerdctl.socket:/var/run/nerdctl.socket \
  -v /usr/local/bin/nerdctlc:/usr/local/bin/nerdctlc \
  alpine nerdctlc ps
```
