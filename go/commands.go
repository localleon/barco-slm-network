package main

type command struct {
	cmd   []byte
	datas map[string][]byte
}

var m map[string]command

func init() {
	m = make(map[string]command)
	m["infrared"] = command{
		cmd: []byte{0x30},
		datas: map[string][]byte{
			"*":          []byte{0x77},
			"0":          []byte{0x19},
			"1":          []byte{0x10},
			"2":          []byte{0x11},
			"3":          []byte{0x12},
			"4":          []byte{0x13},
			"5":          []byte{0x14},
			"6":          []byte{0x15},
			"7":          []byte{0x16},
			"8":          []byte{0x17},
			"9":          []byte{0x18},
			"arrowdown":  []byte{0x05},
			"arrowup":    []byte{0x04},
			"arrowright": []byte{0x06},
			"arrowleft":  []byte{0x07},
			"enter":      []byte{0x0a},
			"exit":       []byte{0x08},
		},
	}
	m["lcdbacklight"] = command{
		cmd: []byte{0x7a, 0x84},
		datas: map[string][]byte{
			"off": []byte{0x0},
			"on":  []byte{0x01},
		},
	}
	m["lcdclear"] = command{
		cmd: []byte{0x7a, 0x85},
	}
}

//Generates a list of commands to wirte to the lcd
func writeLCD(projAddr byte, first string, second string) [][]byte {
	list := make([][]byte, 3)
	list[0] = createBytes(projAddr, m["lcdclear"].cmd, []byte{})
	//The cmd for writing:
	cmdWrite := []byte{0x7a, 0x82}
	//Create the data array
	firstData := []byte{0x00, 0x00} //first line and column
	for _, v := range []byte(first) {
		firstData = append(firstData, v)
	}
	firstData = append(firstData, 0x00)
	list[1] = createBytes(projAddr, cmdWrite, firstData)
	secondData := []byte{0x00, 0x01}
	for _, v := range []byte(second) {
		secondData = append(secondData, v)
	}
	secondData = append(secondData, 0x00)
	list[2] = createBytes(projAddr, cmdWrite, secondData)
	return list
}
