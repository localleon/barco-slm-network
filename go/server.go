package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jacobsa/go-serial/serial"
)

var projAddr byte = 1

func main() {
	portName := flag.String("serial", "COM4", "the identifier for the serial port")
	baudRate := flag.Uint("baudrate", 9600, "the baudrate for the serial communication")
	showKeys := flag.Bool("showCmds", false, "if this flag is set, all possible CMD-DATA combinations are printed")

	flag.Parse()

	if *showKeys {
		for i := range m {
			fmt.Println(i)
			for k := range m[i].datas {
				fmt.Println("\t" + k)
			}
		}
		os.Exit(0)
	}

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
	writeLCD(port, "FMT-Barco-Remote", "by Leon and Moesby")

	router := mux.NewRouter()

	//Start REST-Api:
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

	/*
	 This will serve files under /<filename> from path in var dir
	 Mux.Router is threating endpoints in order. So / has to be the last , else it will
	 get triggerd at every request
	*/
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

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

func writeLCD(port io.ReadWriteCloser, text1, text2 string) {
	data := calcLcdWriteBytes(projAddr, text1, text2)
	writeBytes(port, data[0]) // Clear LCD
	writeBytes(port, data[1]) // Write first line
	writeBytes(port, data[2]) // Write second line
}
