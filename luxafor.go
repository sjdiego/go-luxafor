package main

import (
	"flag"
	"fmt"
	"github.com/karalabe/hid"
	"log"
)

const (
	vendorID uint16 = 0x04d8
	deviceID uint16 = 0xf372
)

var (
	red    int
	green  int
	blue   int
	device int
)

type Luxafor struct {
	deviceInfo hid.DeviceInfo
}

func init() {
	flag.IntVar(&red, "r", 0, "Set intensity of red color (0-255)")
	flag.IntVar(&green, "g", 0, "Set intensity of green color (0-255)")
	flag.IntVar(&blue, "b", 0, "Set intensity of blue color (0-255)")
	flag.IntVar(&device, "d", 1, "Select Luxafor device (1-n)")
}

func main() {
	flag.Parse()

	luxafor, err := OpenDevice()

	if err != nil {
		log.Fatalf("Error opening device. %s\n", err)
	}
	defer func() { _ = luxafor.Close() }()

	if _, err = luxafor.Write([]byte{1, 255, byte(red), byte(green), byte(blue)}); err != nil {
		log.Fatalln("Error sending command.")
	}

	fmt.Println("Command sent to device successfully.")
}

func OpenDevice() (*hid.Device, error) {
	found   := false
	devices := hid.Enumerate(vendorID, deviceID)

	for key, _ := range devices {
		if key == (device-1) {
			found = true
		}
	}

	if !found {
		log.Fatalf("Device %d not found.", device)
	}

	lux := Luxafor { deviceInfo: devices[device-1] }
	return lux.deviceInfo.Open()
}
