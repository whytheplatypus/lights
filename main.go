package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/godbus/dbus"
)

func main() {
	var list []string

	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = conn.Object("org.razer", "/org/razer").Call("razer.devices.getDevices", 0).Store(&list)
	if err != nil {
		panic(err)
	}
	for _, v := range list {
		fmt.Println(v)
		var d float64
		if err := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+v)).Call("razer.device.lighting.brightness.getBrightness", 0).Store(&d); err != nil {
			panic(err)
		}
		fmt.Println(d)

		SetRow(v, conn)
		SetCustom(v, conn)
	}
}

func SetRow(name string, conn *dbus.Conn) {
	test := []uint8{10, 0, 0, 0, 10, 0, 0, 0, 10, 10, 0, 0}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, []uint8{0})
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	err = binary.Write(buf, binary.BigEndian, []uint8{0})
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	err = binary.Write(buf, binary.BigEndian, uint8(len(test)/3-1))
	//err = binary.Write(buf, binary.LittleEndian, []int16{})
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	err = binary.Write(buf, binary.LittleEndian, test)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
	// test := []byte{3, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 00, 00}
	c := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+name)).Call("razer.device.lighting.chroma.setKeyRow", 0, buf.Bytes())
	if c.Err != nil {
		panic(c.Err)
	}
	fmt.Println(c.Body)
	time.Sleep(1 * time.Second)
}

func SetCustom(name string, conn *dbus.Conn) {

	c := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+name)).Call("razer.device.lighting.chroma.setCustom", 0)
	if c.Err != nil {
		panic(c.Err)
	}
	fmt.Println(c.Body)
}
