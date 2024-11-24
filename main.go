package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"blind/tunnel"
)

func main() {
	// Client flags
	clientListen := flag.String("client-listen", "", "(e.g., 127.0.0.1:8080) Local TCP port to listen on")
	clientDest := flag.String("client-dest", "", "(e.g., 10.0.0.1:53) Remote DNS server address")

	// Server flags
	serverListen := flag.String("server-listen", "", "(e.g., 0.0.0.0:53) DNS listen address")
	serverDest := flag.String("server-dest", "", "(e.g., 127.0.0.1:80) Destination TCP address to forward to")

	debug := flag.Bool("debug", false, "Enable debug logging")
	flag.Parse()

	// Server mode if server flags are set
	if *serverListen != "" || *serverDest != "" {
		if *serverListen == "" || *serverDest == "" {
			fmt.Println("Error: both server-listen and server-dest are required for server mode")
			fmt.Println("Example: ./blind -server-listen 0.0.0.0:53 -server-dest 127.0.0.1:80")
			flag.Usage()
			os.Exit(1)
		}
		server := tunnel.NewDNSServer(*serverListen, *serverDest, *debug)
		log.Printf("Starting DNS tunnel server:")
		log.Printf("  DNS listening on: %s", *serverListen)
		log.Printf("  Forwarding to: %s", *serverDest)
		log.Fatal(server.Start())
	}

	// Client mode if client flags are set
	if *clientListen != "" || *clientDest != "" {
		if *clientListen == "" || *clientDest == "" {
			fmt.Println("Error: both client-listen and client-dest are required for client mode")
			fmt.Println("Example: ./blind -client-listen 127.0.0.1:8080 -client-dest 10.0.0.1:53")
			flag.Usage()
			os.Exit(1)
		}
		client, err := tunnel.NewDNSClient(*clientListen, *clientDest, *debug)
		if err != nil {
			log.Fatalf("Failed to create DNS client: %v", err)
		}
		log.Printf("Starting DNS tunnel client:")
		log.Printf("  TCP listening on: %s", *clientListen)
		log.Printf("  Tunneling to DNS server: %s", *clientDest)
		log.Fatal(client.Start())
	}

	// If no mode selected, show usage
	fmt.Println("Error: must specify either client or server mode")
	fmt.Println("\nServer mode example:")
	fmt.Println("  ./blind -server-listen 0.0.0.0:53 -server-dest 127.0.0.1:80")
	fmt.Println("\nClient mode example:")
	fmt.Println("  ./blind -client-listen 127.0.0.1:8080 -client-dest 10.0.0.1:53")
	flag.Usage()
	os.Exit(1)
}
