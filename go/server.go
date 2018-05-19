package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Hundemeier/go-sacn/sacn"

	"github.com/gorilla/mux"
	"github.com/jacobsa/go-serial/serial"
)

var projAddr byte = 1

func main() {
	portName := flag.String("serial", "COM4", "the identifier for the serial port")
	baudRate := flag.Uint("baudrate", 115200, "the baudrate for the serial communication")
	universe := flag.Uint("universe", 1, "the sACN universe to listen on")
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
		//check if we have a lcdread and treat it differnet
		if params["cmd"] == "lcdread" {
			resp := readLCD(port)
			js, err := json.Marshal(resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			log.Printf("Got request: lcdread; responded accordingly")
		} else {
			writeCommand(port, params["cmd"], "")
			log.Printf("Got request: %v, %v", params["cmd"], params["data"])
		}
	}).Methods("GET")

	/*
	 This will serve files under /<filename> from path in var dir
	 Mux.Router is threating endpoints in order. So / has to be the last , else it will
	 get triggerd at every request
	*/
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	recv, err := sacn.Receive(uint16(*universe), "")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for p := range recv.DataChan {
			if p.Data()[0] == 255 {
				writeCommand(port, "shutterclose", "fast")
				log.Println("sACN: shutter closed")
			} else if p.Data()[0] == 0 {
				writeCommand(port, "shutteropen", "fast")
				log.Println("sACN: shutter opened")
			}
		}
	}()

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

//Help struct for the json data
type lcd struct {
	First  string //`json:"first"`
	Second string //`json:"second"`
}

var readLCDLock = false

func readLCD(port io.ReadWriteCloser) lcd {
	for readLCDLock {
		//wait unitl the lock is over
		time.Sleep(10 * time.Millisecond)
	}
	readLCDLock = true

	result := lcd{}
	//Attemp to read the first line
	writeBytes(port, createBytes(projAddr, []byte{0x7a, 0x02}, []byte{0, 0}))
	for {
		reader := bufio.NewReader(port)
		reply, _ := reader.ReadBytes('\xff')
		if reply[2] == 0x7a && reply[3] == 0x02 {
			result.First = cString(reply[6:])
			break
		}
	}
	writeBytes(port, createBytes(projAddr, []byte{0x7a, 0x02}, []byte{0, 1}))

	//second line
	for {
		reader := bufio.NewReader(port)
		reply, _ := reader.ReadBytes('\xff')
		if reply[2] == 0x7a && reply[3] == 0x02 {
			result.Second = cString(reply[6:])
			break
		}
	}
	readLCDLock = false
	return result
}

func cString(data []byte) string {
	index := 0
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			index = i
			break
		}
	}
	return string(data[:index])
}
