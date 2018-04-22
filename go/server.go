package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacobsa/go-serial/serial"
)

var projAddr byte = 1

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
	defer port.Close() //Make sure to close the port

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
