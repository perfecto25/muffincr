package main

import (
	"fmt"
	"net"
)

func main() {
	// Define the multicast group address and port
	multicastAddr := "239.0.0.1"
	port := "12345"

	// Define the local IP address of the interface you want to bind to
	localIP := "192.168.30.10"

	// Resolve the multicast address and port
	addr, err := net.ResolveUDPAddr("udp", multicastAddr+":"+port)
	if err != nil {
		fmt.Println("Failed to resolve address:", err)
		return
	}

	// Resolve the local IP address
	localAddr, err := net.ResolveIPAddr("ip", localIP)
	if err != nil {
		fmt.Println("Failed to resolve local address:", err)
		return
	}

	// Create a UDP connection
	conn, err := net.ListenPacket("udp", addr.String())
	if err != nil {
		fmt.Println("Failed to create socket:", err)
		return
	}
	defer conn.Close()

	// Set the outgoing interface
	err = conn.(*net.UDPConn).SetMulticastInterface(localAddr)
	if err != nil {
		fmt.Println("Failed to set multicast interface:", err)
		return
	}

	// Join the multicast group
	err = conn.(*net.UDPConn).JoinGroup(localAddr, &net.UDPAddr{IP: addr.IP})
	if err != nil {
		fmt.Println("Failed to join multicast group:", err)
		return
	}

	// Read from the socket
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFrom(buffer)
	if err != nil {
		fmt.Println("Failed to read from socket:", err)
		return
	}

	// Process the received data
	fmt.Println("Received:", string(buffer[:n]))
}
