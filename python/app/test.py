#!usr/bin/python3 
import serial

shutterclose = b'\xfe\x01\x23\x42\x00\x66\xff'
shutteropen =  b'\xfe\x01\x22\x42\x00\x65\xff'

ser = serial.Serial('/dev/ttyUSB0')  # open serial port
print("Used Port:" + ser.name)         # check which port was really used
ser.write(shutterclose)     # write a byte 

ser.close()             # close port
