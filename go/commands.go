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
			"standby":    []byte{0x0e},
		},
	}
	m["lcdbacklight"] = command{
		cmd: []byte{0x7a, 0x84},
		datas: map[string][]byte{
			"off": []byte{0x0},
			"on":  []byte{0x01},
		},
	}
	m["shutterclose"] = command{
		cmd: []byte{0x23, 0x42},
		datas: map[string][]byte{
			"fast": []byte{0x00},
			"slow": []byte{0x01},
		},
	}
	m["shutteropen"] = command{
		cmd: []byte{0x22, 0x42},
		datas: map[string][]byte{
			"fast": []byte{0x00},
			"slow": []byte{0x01},
		},
	}
	m["lensfocus"] = command{
		cmd: []byte{0xf4, 0x83},
		datas: map[string][]byte{
			"near": []byte{0x00},
			"far":  []byte{0x01},
		},
	}
	m["lensshift"] = command{
		cmd: []byte{0xf4, 0x81},
		datas: map[string][]byte{
			"up":    []byte{0x00},
			"down":  []byte{0x01},
			"left":  []byte{0x02},
			"right": []byte{0x03},
		},
	}
	m["lenszoom"] = command{
		cmd: []byte{0xf4, 0x82},
		datas: map[string][]byte{
			"in":  []byte{0x00},
			"out": []byte{0x01},
		},
	}
	m["freezeoff"] = command{
		cmd: []byte{0x26, 0x23},
	}
	m["freezeon"] = command{
		cmd: []byte{0x27, 0x23},
	}
	m["menuexit"] = command{
		cmd: []byte{0x42, 0x01},
		datas: map[string][]byte{
			"one": []byte{0x01},
			"all": []byte{0xff},
		},
	}
	m["source"] = command{
		cmd: []byte{0x33},
		datas: map[string][]byte{
			"1": []byte{0x01},
			"2": []byte{0x02},
			"3": []byte{0x03},
			"4": []byte{0x04},
		},
	}
	m["pattern"] = command{
		cmd: []byte{0x41},
		datas: map[string][]byte{
			"convergence-g":  []byte{0x01},
			"convergence-rg": []byte{0x02},
			"convergence-gb": []byte{0x03},
			"hatch":          []byte{0x04},
			"checkerboard":   []byte{0x19},
			"colorbars":      []byte{0x1a},
			"multiburst":     []byte{0x1b},
			"outline":        []byte{0x1c},
			"chars":          []byte{0x23},
		},
	}
}
