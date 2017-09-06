package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/godbus/dbus"
	"github.com/whytheplatypus/lights"
	"golang.org/x/image/colornames"
)

var r, c int
var max = []int{0}
var min = 100000
var device string
var conn *dbus.Conn

func init() {
	flag.StringVar(&device, "d", "", "The ID of your razer device from dbus")
	flag.Parse()
}

func main() {
	if !lights.ValidDevice(device) {
		panic(errors.New("not a known device"))
	}
	data, err := Asset("config")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(data))
	if err := ioutil.WriteFile("/tmp/kava", data, 0644); err != nil {
		panic(err)
	}
	cava := exec.Command("cava", "-p", "/tmp/kava")
	in, err := cava.StdoutPipe()
	if err != nil {
		panic(err)
	}
	conn = lights.Conn
	var m []int32
	if err := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+device)).Call("razer.device.misc.getMatrixDimensions", 0).Store(&m); err != nil {
		panic(err)
	}

	//fmt.Println(m)
	r = int(m[0])
	c = int(m[1])

	go func() {
		for {
			if len(max) > 1 {
				max = max[:len(max)-1]
			}
			<-time.After(1 * time.Second)
		}
	}()

	go run(in)
	cava.Start()
	cava.Wait()
}

func run(in io.Reader) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		cf := scanner.Text()
		//fmt.Println(cf)
		sbars := strings.Split(cf, ";")
		ibars := make([]int, len(sbars))
		for i, b := range sbars {
			ib, err := strconv.Atoi(b)
			if err != nil {
				//log.Printf("bad bar : %s : %s", b, err.Error())
				continue
			}
			ibars[i] = ib
		}
		s := render(ibars)
		lights.Apply(device, conn, s)
		lights.SetCustom(device, conn)

	}
	return scanner.Err()
}

func render(bars []int) *lights.Set {
	frame := &lights.Set{
		Rows: make([]*lights.Row, r),
	}
	for i, _ := range frame.Rows {
		frame.Rows[i] = &lights.Row{
			Num:    uint8(i),
			Colors: make([]lights.Color, c),
		}
		for k, _ := range frame.Rows[i].Colors {
			frame.Rows[i].Colors[k] = &lights.RGBA{colornames.Snow}
		}
	}
	for i, h := range bars {
		if h > max[len(max)-1] || (h < min && h > 0) {
			if h > max[len(max)-1] {
				//fmt.Printf("update max %d", h)
				//fmt.Println("")
				max = append(max, h)
			} else if h < min {
				//fmt.Printf("update min %d", h)
				//fmt.Println("")
				min = h
			}
		}
		h = int((float64(r) * float64(h-min)) / (float64(max[len(max)-1]) - float64(min)))
		//fmt.Println(i, h, r, min, max)
		//frame.Rows[h].Colors[i] = &lights.RGBA{colornames.Purple}
		for ii := r - h; ii < r; ii++ {
			if ii >= 0 && ii < len(frame.Rows) && i >= 0 && i < len(frame.Rows[ii].Colors) {
				frame.Rows[ii].Colors[i] = &lights.RGBA{colornames.Purple}
			} else {
				//log.Printf("out of bounds : %d : %d", ii, i)
			}
		}
	}
	return frame
}
