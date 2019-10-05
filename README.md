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

### sACN Usage
For use in an light/video/audio production enviroment this programm supports sACN (E1.31 protocol by the ESTA). This can be used to control the Barco via an lighting console.

Send an Unicast sACN Signal to the Computer. Multicast is not supported. 
Default Universe is 1, but can be changed via Flag --universe. 
- Set Channel 1 to 255 to close the Shutter. 
- Set Channel 2 to 255 to open the Shutter. 

## API
The Projects exposes a simple HTTP API under localhost:80/api/. Just send an HTTP Get to one of the following EndPoints.Caution, there's not authorization on the API. This Project is only intended to run in protected networks

Generic Call:
` curl localhost:80/api/ENDPOINT/FUNCTION`

Example Call: 
` curl localhost:80/api/shutteropen/fast `


```
# Shutter Functions
shutterclose /
        fast
        slow
shutteropen /
        fast
        slow
# Backlight of LCD
lcdbacklight /
        off
        on
# Freezing
freezeoff 
freezeon
# Focus of the Lens
lensfocus /
        near
        far
# Shift the Lens in multiple Directions
lensshift /
        up
        down
        left
        right
# Zoom of the Lens
lenszoom / 
        in
        out
# Exit Menu on LCD
menuexit /
        one
        all
# Select Video Source
source /
        1
        2
        3
        4
        
# Multiple Sample Patterns are displayed
pattern / 
        convergence-g
        convergence-rg
        convergence-gb
        hatch
        chars
        checkerboard
        colorbars
        multiburst
        outline
# Every Button from the IR Remote
infrared /
        7
        0
        1
        2
        arrowleft
        4
        8
        arrowdown
        6
        standby
        *
        3
        5
        enter
        exit
        9
        arrowup
        arrowright
```
## Dependencies

Infos about the Dependencies can be found in [go.mod](https://github.com/localleon/barco-slm-network/blob/master/go/go.mod) and [go.sum](https://github.com/localleon/barco-slm-network/blob/master/go/go.sum).

## Develop / Build
Pull Requests and Issues are always welcome. 

Clone Repo and execute build.sh to build your binarys for every tested plattfrom. 
