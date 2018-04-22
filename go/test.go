package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacobsa/go-serial/serial"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/{cmd}/{data}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		log.Printf("Got request: %v, %v", params["cmd"], params["data"])
	}).Methods("GET")
	router.HandleFunc("/api/{cmd}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		log.Printf("Got request: %v, %v", params["cmd"], params["data"])
	}).Methods("GET")
	fmt.Println("Press Ctrl+C to abort")
	log.Fatal(http.ListenAndServe(":8000", router))

	m := map[string][]byte{
		"test":  []byte{1, 2, 3},
		"test2": []byte{4, 5, 6},
	}
	fmt.Println(m["test"])
	fmt.Println(m[""])
	fmt.Println([]byte("test"))
	for _, v := range []byte("test") {
		fmt.Println(v)
	}
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

	b := []byte{0xfe, 0x01, 0x22, 0x42, 0x00, 0x65, 0xff}
	//b := []byte{0xfe, 0x01, 0x23, 0x42, 0x00, 0x66, 0xff}
	n, err := port.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")

}
