# barco-slm-network
Network Interface for Barco SLM R6 Performer via the RS232 Port
![barco-slm-network](https://user-images.githubusercontent.com/28186014/57984895-98c4e380-7a60-11e9-976b-bcf88d789216.png)

## Introduction
This Project implements the most important functions of 
Barcos Serial Communication Protocol for LCD-DLP Projectors. 
Its written in Golang and provides you with an Rest API and an sACN Interface
for controlling it from Devices like an GrandMa2. This was only tested on a Barco SLM R6 Performer, but should support more modells. 

## Usage 
### Setup on the Barco Beamer 
1. Turn on RS232 Communication on the Beamer. 
2. Set a Baudrate of 115200 (variable via Flag) and the Projector Adress to 1 (this is important!) 

### Setup on PC 
1. Connect your Computer via a USB-to-RS232 Adapter or other serial Interface to the Input RS232 Port of the Beamer
2. Download the Git Repo and one of binarys from the relase section.
3. Execute the file via CLI and set the "--serial" flag to your COM-Port 
4. Visit the Web-Interface under localhost:80 and enjoy!

**If everything worked there should be a text displayed on the LCD of the Beamer**

Example CLI Command:
```
barco-slm-network-v1-0.exe -serial=COM3
```
### CLI FLags 
This is also displayed via `--help` 
```
--serial ("the identifier for the serial port, default is COM4")
--baudrate ("the baudrate for the serial communication, default is 115200")
--universe ("the sACN universe to listen on, default is 1")
--showCmds ("if this flag is set, all possible CMD-DATA combinations are printed")
--verbose ("extra Output, default false")
```
### Web Interface
Visit http://localhost:80 to use the WebInterface. Here you can control all supported Functions and read the Text from the LCD. 
The following buttons can be held down to repeat the API calls: arrow up/down/left/right, lensshift up/down/left/right, focus near/far, zoom in/out.
The following keys on your keyboard are shortcuts to these API/Button functions:
```
arrow keys: arrow up/down/left/right
o: open shutter fast
c: close shutter fast
enter key: enter
ESC: cancel
```

### sACN Usage
For use in an light/video/audio production enviroment this programm supports sACN (E1.31 protocol by the ESTA). This can be used to control the Barco via an lighting console.

Send an Unicast sACN Signal to the Computer. Multicast is not supported. 
Default Universe is 1, but can be changed via Flag --universe. 
- Set Channel 1 to 255 to close the Shutter. 
- Set Channel 2 to 255 to open the Shutter. 

## API
The Projects exposes a simple HTTP API under localhost:80/api/. Just send an HTTP Get to one of the following EndPoints.Caution, there's **no authorization** on the API. This Project is only intended to run in protected networks. 

This API is documented with the [OpenAPI v3.0.0 Specification](https://github.com/OAI/OpenAPI-Specification). The Specification file (YAML) can be found under `go/api/openapi.yaml`. To view the file use something like Postman or [Swagger Editor](https://editor.swagger.io/)

```
Generic Call:
curl localhost:80/api/ENDPOINT/FUNCTION

Example Call: 
curl localhost:80/api/shutteropen/fast
```

## Dependencies

Infos about the Dependencies can be found in [go.mod](https://github.com/localleon/barco-slm-network/blob/master/go/go.mod) and [go.sum](https://github.com/localleon/barco-slm-network/blob/master/go/go.sum).

## Develop / Build
Pull Requests and Issues are always welcome. 

Clone Repo and execute build.sh to build your binarys for every tested plattfrom. 
