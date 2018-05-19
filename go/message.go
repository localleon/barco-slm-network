package main

//Message : This is a message struct with the needed parameters of a simple message
type Message struct {
	projectorAddr byte
	cmd           []byte
	data          []byte
}

func (m *Message) calculateChecksum() byte {
	chk := int(m.projectorAddr)
	//Add all cmd bytes
	for _, v := range m.cmd {
		chk += int(v)
	}
	for _, v := range m.data {
		chk += int(v)
	}
	return byte(chk % 256)
}

func (m *Message) getByteSlice() []byte {
	arr := make([]byte, 4)
	arr = append(arr, 0xfe)
	for _, h := range convertByte(m.projectorAddr) {
		arr = append(arr, h)
	}
	for _, v := range m.cmd {
		for _, h := range convertByte(v) {
			arr = append(arr, h)
		}
	}
	for _, v := range m.data {
		for _, h := range convertByte(v) {
			arr = append(arr, h)
		}
	}
	for _, h := range convertByte(m.calculateChecksum()) {
		arr = append(arr, h)
	}
	arr = append(arr, 0xff)
	return arr
}

//Creates a byte slice generated from the given data
func createBytes(projAddr byte, cmd []byte, data []byte) []byte {
	msg := Message{
		projectorAddr: projAddr,
		cmd:           cmd,
		data:          data,
	}
	return msg.getByteSlice()
}

//Converts a byte if its x80 xfe or xff like the documentation wants to
//Otherwise it returns the given byte again
func convertByte(val byte) []byte {
	if val == 0x80 {
		return []byte{0x80, 0x00}
	} else if val == 0xfe {
		return []byte{0x80, 0x7e}
	} else if val == 0xff {
		return []byte{0x80, 0x7f}
	}
	return []byte{val}
}

//Returns a nil pointer if the byt1 is not 0x80 or the bytes could not be converted
func convertBytes(byt1 byte, byt2 byte) *byte {
	bytpt := new(byte)
	if byt1 == 0x80 {
		if byt2 == 0x00 {
			*bytpt = 0x80
		} else if byt2 == 0x7e {
			*bytpt = 0xfe
		} else if byt2 == 0x7f {
			*bytpt = 0xff
		} else {
			bytpt = nil
		}
	} else {
		bytpt = nil
	}
	return bytpt
}

//Generates a list of commands to wirte to the lcd
func calcLcdWriteBytes(projAddr byte, first string, second string) [][]byte {
	list := make([][]byte, 3)
	//LCD clear:
	list[0] = createBytes(projAddr, []byte{0x7a, 0x85}, []byte{})
	//The cmd for writing:
	cmdWrite := []byte{0x7a, 0x82}
	//Create the data array
	firstData := []byte{0x00, 0x00} //first line and column
	for _, v := range []byte(first) {
		firstData = append(firstData, v)
	}
	// Convert first string to bytearray
	firstData = append(firstData, 0x00)
	list[1] = createBytes(projAddr, cmdWrite, firstData)
	secondData := []byte{0x00, 0x01}
	for _, v := range []byte(second) {
		secondData = append(secondData, v)
	}
	// Convert second String to bytearry
	secondData = append(secondData, 0x00)
	list[2] = createBytes(projAddr, cmdWrite, secondData)
	return list
}
