package lights

import (
	"bytes"
	"encoding/binary"
	"image/color"
	"io"
	"log"

	"github.com/godbus/dbus"
)

var DeviceList []string
var validDevices = map[string]struct{}{}
var Conn *dbus.Conn

func init() {
	var err error
	Conn, err = dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	err = Conn.Object("org.razer", "/org/razer").Call("razer.devices.getDevices", 0).Store(&DeviceList)
	if err != nil {
		panic(err)
	}
	for _, d := range DeviceList {
		validDevices[d] = struct{}{}
	}
}

func ValidDevice(d string) bool {
	_, ok := validDevices[d]
	return ok
}

type RGBA struct {
	color.RGBA
}

func (rgba *RGBA) RGB() []uint8 {
	return []uint8{rgba.R, rgba.G, rgba.B}
}

type RGB struct {
	r, g, b uint8
}

func (rgb *RGB) RGB() []uint8 {
	return []uint8{rgb.r, rgb.g, rgb.b}
}

type Color interface {
	RGB() []uint8
}

// TODO rename
type Row struct {
	Num uint8
	// can make this matter but requires updates to the driver I think?
	Start  uint8
	Colors []Color
}

func (r *Row) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, r.Num); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, r.Start); err != nil {
		return err
	}
	/**/
	offset := int(r.Start) + len(r.Colors) - 1
	/**/
	stopcol := offset
	if err := binary.Write(w, binary.LittleEndian, uint8(stopcol)); err != nil {
		return err
	}
	for _, c := range r.Colors {
		if err := binary.Write(w, binary.LittleEndian, append([]uint8{}, c.RGB()...)); err != nil {
			return err
		}
	}

	return nil
}

// TODO rename
type Set struct {
	Rows []*Row
}

func (s *Set) Encode(w io.Writer) error {
	for _, r := range s.Rows {
		if err := r.Encode(w); err != nil {
			return err
		}
	}
	return nil
}

func Apply(name string, conn *dbus.Conn, s *Set) {

	buf := new(bytes.Buffer)
	s.Encode(buf)

	log.Printf("%x", buf.Bytes())
	// test := []byte{3, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 255, 00, 255, 00, 00}
	c := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+name)).Call("razer.device.lighting.chroma.setKeyRow", 0, buf.Bytes())
	if c.Err != nil {
		panic(c.Err)
	}
}

func SetCustom(name string, conn *dbus.Conn) {

	c := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+name)).Call("razer.device.lighting.chroma.setCustom", 0)
	if c.Err != nil {
		panic(c.Err)
	}
}

func SetReactive(name string, conn *dbus.Conn, color Color) {
	rgb := color.RGB()
	c := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+name)).Call("razer.device.lighting.chroma.setReactive", 0, rgb[0], rgb[1], rgb[2], uint8(1))
	if c.Err != nil {
		panic(c.Err)
	}
}

func ClearCustom(name string, conn *dbus.Conn, color Color) {
	var m []int32
	if err := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+name)).Call("razer.device.misc.getMatrixDimensions", 0).Store(&m); err != nil {
		panic(err)
	}

	//fmt.Println(m)
	r := int(m[0])
	c := int(m[1])

	frame := &Set{
		Rows: make([]*Row, r),
	}
	for i, _ := range frame.Rows {
		frame.Rows[i] = &Row{
			Num:    uint8(i),
			Colors: make([]Color, c),
		}
		for k, _ := range frame.Rows[i].Colors {
			frame.Rows[i].Colors[k] = color
		}
	}
	Apply(name, conn, frame)
	SetCustom(name, conn)
}

func SetBreathRandom(name string, conn *dbus.Conn) {

	c := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+name)).Call("razer.device.lighting.chroma.setBreathRandom", 0)
	if c.Err != nil {
		panic(c.Err)
	}
}

func GetDeviceName(name string, conn *dbus.Conn) (string, error) {
	var s string
	c := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+name)).Call("razer.device.misc.getDeviceName", 0)
	c.Store(&s)
	return s, c.Err
}

func GetMatrixDimensions(name string, conn *dbus.Conn) ([]int32, error) {
	var m []int32
	err := conn.Object(
		"org.razer",
		dbus.ObjectPath("/org/razer/device/"+name),
	).Call(
		"razer.device.misc.getMatrixDimensions",
		0,
	).Store(&m)
	return m, err
}
