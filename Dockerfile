FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY . .

# 编译二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -o nerdctld cmd/nerdctld/main.go && \
    CGO_ENABLED=0 GOOS=linux go build -o nerdctlc cmd/nerdctlc/main.go

# 使用 Alpine 制作最终打包镜像
FROM alpine:latest

WORKDIR /dist
COPY --from=builder /build/nerdctld /dist/
COPY --from=builder /build/nerdctlc /dist/
COPY --from=builder /build/systemd/nerdctld.service /dist/

# 打包文件
RUN tar czf nerdctld.tar.gz nerdctld nerdctld.service nerdctlc