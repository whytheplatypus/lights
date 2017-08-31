package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/godbus/dbus"
	"github.com/whytheplatypus/lights"
	"golang.org/x/image/colornames"
)

func main() {
	device := "BY1636A24100115"
	conn := lights.Conn
	var m []int32
	if err := conn.Object("org.razer", dbus.ObjectPath("/org/razer/device/"+device)).Call("razer.device.misc.getMatrixDimensions", 0).Store(&m); err != nil {
		panic(err)
	}

	fmt.Println(m)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cf := scanner.Text()
		fmt.Println(cf)
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
		s := render(ibars, int(m[0]), int(m[1]))
		lights.Apply(device, conn, s)
		lights.SetCustom(device, conn)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func render(bars []int, r, c int) *lights.Set {
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
		//fmt.Println(i, h)
		//frame.Rows[h].Colors[i] = &lights.RGBA{colornames.Purple}
		for ii := r - h; ii < r; ii++ {
			frame.Rows[ii].Colors[i] = &lights.RGBA{colornames.Purple}
		}
	}
	return frame
}
