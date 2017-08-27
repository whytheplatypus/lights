package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image/color"
	"io"
	"time"

	"golang.org/x/image/colornames"

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
		var m []int32
		if err := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+v)).Call("razer.device.misc.getMatrixDimensions", 0).Store(&m); err != nil {
			panic(err)
		}
		fmt.Println(m)
		reset := &Set{
			rows: make([]*Row, int(m[0])),
		}
		for i, _ := range reset.rows {
			reset.rows[i] = &Row{
				num:    uint8(i),
				start:  0,
				colors: make([]Color, int(m[1])),
			}
			for k, _ := range reset.rows[i].colors {
				reset.rows[i].colors[k] = &RGBA{colornames.Purple}
			}
		}
		Apply(v, conn, reset)
		/*
			Apply(v, conn, &Set{
				[]*Row{
					&Row{
						num: 3,
						colors: []Color{
							&RGB{
								10, 0, 0,
							},
							&RGB{
								10, 0, 10,
							},
							&RGB{
								0, 0, 100,
							},
							//&RGBA{colornames.Purple},
						},
					},
				},
			})
			/**/

		SetCustom(v, conn)
	}
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
	num uint8
	// can make this matter but requires updates to the driver I think?
	start  uint8
	colors []Color
}

func (r *Row) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, r.num); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, r.start); err != nil {
		return err
	}
	/**/
	offset := int(r.start)*2 + len(r.colors) - 1
	fmt.Println(offset, int(r.start))
	/**/
	stopcol := offset
	if err := binary.Write(w, binary.LittleEndian, uint8(stopcol)); err != nil {
		return err
	}
	for _, c := range r.colors {
		if err := binary.Write(w, binary.LittleEndian, append([]uint8{}, c.RGB()...)); err != nil {
			return err
		}
	}
	/**/
	if err := binary.Write(w, binary.LittleEndian, make([]uint8, int(r.start)*3)); err != nil {
		return err
	}
	/**/

	return nil
}

// TODO rename
type Set struct {
	rows []*Row
}

func (s *Set) Encode(w io.Writer) error {
	for _, r := range s.rows {
		if err := r.Encode(w); err != nil {
			return err
		}
	}
	return nil
}

func Apply(name string, conn *dbus.Conn, s *Set) {

	buf := new(bytes.Buffer)
	s.Encode(buf)

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
