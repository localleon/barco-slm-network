# barco-slm-network
Network Interface for Barco SLM R6 Performer via the RS232 Port

## Introduction
This Project implements the most important functions of 
Barcos Serial Communication Protocol for LCD-DLP Projectors. 
Its written in Golang and provides you with an Rest API and an sACN Interface
for controlling it from Devices like an GrandMa2. 

## Usage 
1. Setup your Barco Beamer with RS232 Serial Communication and a Baudrate of 9600
2. Connect your Computer via a USB-to-RS232 Adapter or other serial Interface. 
3. Download on of the builds. Open your commandline and type ``` file -name=SERIALPORT ```

Example:
```
barco-slm-network-v1-0.exe -name=COM3
```

## Dependencies
```
import (
        "flag"
        "fmt"
        "io"
        "log"
        "net/http"
        "github.com/gorilla/mux"
        "github.com/jacobsa/go-serial/serial"
)
```
