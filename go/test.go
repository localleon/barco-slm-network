package main

import (
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

func main2() {
	// Set up options.
	options := serial.OpenOptions{
		PortName:        "COM4",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	b := []byte{0xfe, 0x01, 0xf4, 0x81, 0x00, 0x76, 0xff}
	//b := []byte{0xfe, 0x01, 0x22, 0x42, 0x00, 0x65, 0xff}
	//b := []byte{0xfe, 0x01, 0x23, 0x42, 0x00, 0x66, 0xff}
	n, err := port.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")

}
