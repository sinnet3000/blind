# Blind - DNS Tunnel

A high-performance DNS tunneling tool for TCP/IP traffic, written in Go.

Copyright (c) 2024 Barrett Lyon. All rights reserved.
MIT License

## Overview

Blind allows you to tunnel TCP traffic through DNS queries, enabling connectivity in restricted network environments. It consists of a client and server component that work together to establish a bidirectional communication channel using DNS protocols.

## Features

- TCP over DNS tunneling
- Support for both client and server modes
- High performance and low latency
- Automatic session management
- Resilient connection handling
- Debug logging

## Installation

```bash
go install github.com/blyon/blind@latest
```

Or build from source:

```bash
git clone https://github.com/blyon/blind.git
cd blind
go build
```

## Usage Examples

### Basic Examples

1. Simple SSH Tunnel:

```bash
# On DNS server (public internet)
sudo ./blind -server-listen 0.0.0.0:53 -server-dest 127.0.0.1:22

# On client machine (behind firewall)
./blind -client-listen 127.0.0.1:2222 -client-dest dns-server.com:53

# Connect via SSH
ssh -p 2222 user@127.0.0.1
```

2. Debug Logging:

```bash
./blind -client-listen 127.0.0.1:2222 \
        -client-dest dns.example.com:53 \
        -debug
```

### Advanced Examples

1. HTTP Proxy Tunnel:

```bash
# Server side (forwarding to local HTTP proxy)
sudo ./blind -server-listen 0.0.0.0:53 -server-dest 127.0.0.1:3128 -debug

# Client side
./blind -client-listen 127.0.0.1:8080 -client-dest dns.example.com:53

# Configure browser to use 127.0.0.1:8080 as HTTP proxy
```

2. Database Connection Tunnel:

```bash
# Server side (forwarding to PostgreSQL)
sudo ./blind -server-listen 0.0.0.0:53 -server-dest db.internal:5432

# Client side
./blind -client-listen 127.0.0.1:5432 -client-dest dns.example.com:53

# Connect to database
psql -h 127.0.0.1 -p 5432 -U dbuser dbname
```

### Systemd Service Example

Create a systemd service file for automatic startup:

```ini
# /etc/systemd/system/blind.service
[Unit]
Description=Blind DNS Tunnel Service
After=network.target

[Service]
Type=simple
User=root
ExecStart=/usr/local/bin/blind -server-listen 0.0.0.0:53 -server-dest 10.0.0.1:22
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

Enable and start the service:
```bash
sudo systemctl enable blind
sudo systemctl start blind
sudo systemctl status blind
```

### Docker Example

```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o blind

FROM alpine:latest
COPY --from=0 /app/blind /usr/local/bin/
EXPOSE 53/udp
ENTRYPOINT ["blind"]
```

Run the Docker container:
```bash
# Server mode
docker run -p 53:53/udp blind -server-listen 0.0.0.0:53 -server-dest target:22

# Client mode
docker run -p 2222:2222 blind -client-listen 0.0.0.0:2222 -client-dest dns.example.com:53
```

## License

MIT License - See LICENSE file for details

## Author

Barrett Lyon
```

</rewritten_file>
```
