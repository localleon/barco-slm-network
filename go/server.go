package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacobsa/go-serial/serial"
)

var projAddr byte = 1

func main() {
	portName := flag.String("name", "COM4", "the identifier for the serial port")
	baudRate := flag.Uint("baudrate", 9600, "the baudrate for the serial communication")

	flag.Parse()
	fmt.Printf("Trying to open %v \n", *portName)
	// Set up options.
	options := serial.OpenOptions{
		PortName:        *portName,
		BaudRate:        *baudRate,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("Could not open serial port: %v", err)
	}
	defer port.Close() //Make sure to close the port
	fmt.Printf("Successfully opened the serial port! Baudrate: %v", *baudRate)

	//Start REST-Api:
	router := mux.NewRouter()
	router.HandleFunc("/api/{cmd}/{data}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		writeCommand(port, params["cmd"], params["data"])
		log.Printf("Got request: %v, %v", params["cmd"], params["data"])
	}).Methods("GET")
	router.HandleFunc("/api/{cmd}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		writeCommand(port, params["cmd"], "")
		log.Printf("Got request: %v, %v", params["cmd"], params["data"])
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func writeCommand(port io.ReadWriteCloser, cmd string, data string) {
	writeBytes(port, createBytes(projAddr, m[cmd].cmd, m[cmd].datas[data]))
}

func writeBytes(port io.ReadWriteCloser, data []byte) {
	if len(data) > 3 { // Only write if the data has at least start, checksum and end byte and a cmd set
		_, err := port.Write(data)
		if err != nil {
			log.Printf("Error: port.Write: %v", err)
		}
	}
}
