package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/godbus/dbus"
	"github.com/whytheplatypus/lights"
	"golang.org/x/image/colornames"
)

var r, c int
var max = []int{0}
var min = 0
var device string
var conn *dbus.Conn

func init() {
	flag.StringVar(&device, "d", "", "The ID of your razer device from dbus")

	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Enable for verbose logging")

	flag.Parse()

	if verbose {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
	} else {
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	if !lights.ValidDevice(device) {
		log.Println(lights.DeviceList)
		fmt.Printf("Please set a device with -d : %s \n", lights.DeviceList)
		return
	}

	data, err := Asset("config")
	if err != nil {
		panic(err)
	}
	log.Println(string(data))
	if err := ioutil.WriteFile("/tmp/kava", data, 0644); err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() {
		<-done
		cancel()
	}()
	cava := exec.CommandContext(ctx, "cava", "-p", "/tmp/kava")
	in, err := cava.StdoutPipe()
	if err != nil {
		panic(err)
	}
	conn = lights.Conn
	m, err := lights.GetMatrixDimensions(device, conn)
	if err != nil {
		panic(err)
	}

	log.Println(m)

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
	defer lights.SetReactive(device, conn, &lights.RGBA{RGBA: colornames.Purple})
	cava.Start()
	cava.Wait()
}

func run(in io.Reader) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		cf := scanner.Text()
		log.Println(cf)
		sbars := strings.Split(cf, ";")
		ibars := make([]int, len(sbars))
		for i, b := range sbars {
			ib, err := strconv.Atoi(b)
			if err != nil {
				log.Printf("bad bar : %s : %s", b, err.Error())
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
	for i := range frame.Rows {
		frame.Rows[i] = &lights.Row{
			Num:    uint8(i),
			Colors: make([]lights.Color, c),
		}
		for k := range frame.Rows[i].Colors {
			frame.Rows[i].Colors[k] = &lights.RGBA{RGBA: colornames.Snow}
		}
	}
	for i, h := range bars {
		if h > max[len(max)-1] || (h < min && h > 0) {
			if h > max[len(max)-1] {
				log.Printf("update max %d", h)
				log.Println("")
				max = append(max, h)
			} else if h < min {
				log.Printf("update min %d", h)
				log.Println("")
				min = h
			}
		}
		h = int((float64(r) * float64(h-min)) / (float64(max[len(max)-1]) - float64(min)))
		for ii := r - h; ii < r; ii++ {
			if ii >= 0 && ii < len(frame.Rows) && i >= 0 && i < len(frame.Rows[ii].Colors) {
				frame.Rows[ii].Colors[i] = &lights.RGBA{RGBA: colornames.Purple}
			} else {
				log.Println(i, h, r, min, max)
				log.Printf("out of bounds : %d : %d", ii, i)
				break
			}
		}
	}
	return frame
}
