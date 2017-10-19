package commands

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/whytheplatypus/lights/razer"
	"golang.org/x/image/colornames"
)

type Renderer struct{}

func (r *Renderer) Run(args []string) error {
	// TODO(getkeyboard) what does this mean that it's the first device?
	dev := razer.DeviceList[0]
	// TODO(defaultpreset) why set it back to this? what's going on here?
	defer razer.SetBreathRandom(dev, razer.Conn)

	// TODO(sigsetup) better for this to be a named function so it's clear what's happening
	sig := make(chan os.Signal)

	//Setup signal handling for clean shutdown
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	in := os.Stdin

	// TODO(flagooo) this should be the first thing that happens
	rFlags := flag.NewFlagSet("renderer", flag.ContinueOnError)
	var path string
	rFlags.StringVar(
		&path,
		"fifo",
		os.Getenv("LIGHTS_FIFO"),
		"specify path for a fifo, one will be created if it doesn't already exist at the path")

	if err := rFlags.Parse(args); err != nil {
		return err
	}

	//TODO(fifocreate) named function to be clear we're making a fifo
	if path != "" {
		os.Remove(path)
		if err := syscall.Mkfifo(path, 0600); err != nil {
			log.Println(err)
			return err
		}
		var err error
		in, err = os.OpenFile(path, os.O_RDWR, 0600)
		if err != nil {
			log.Println(err)
			return err
		}
		// TODO(removeremove) can we not use these?
		defer os.Remove(path)
	}

	go render(in)
	<-sig
	return nil
}

func render(in io.Reader) {
	//TODO(dupgetkeyboard) I've seen ths before, that's no good
	dev := razer.DeviceList[0]
	// TODO(getkeyboard) encapsulate to be clear we're getting a keyboard
	dn, err := razer.GetDeviceName(dev, razer.Conn)
	if err != nil {
		log.Fatal(err)
	}
	//TODO(nokeyboardcase) what if there's no keyboard
	keyboard := razer.Keyboards[dn]

	// TODO(renderloop) what is this? a render loop?
	pipe := make(chan *razer.Set, 100)
	go func(pipe <-chan *razer.Set) {
		for s := range pipe {
			razer.Apply(dev, razer.Conn, s)
			razer.SetCustom(dev, razer.Conn)
		}
	}(pipe)
	// can configre to come from a fifo
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		// TODO(untangle) these 3 or 4 statements feel jumbled
		line := scanner.Text()
		if len(pipe) > 0 {
			log.Println("[DEBUG] skipping", line)
			continue
		}
		s := &razer.Set{
			Rows: []*razer.Row{},
		}
		log.Println("[DEBUG]", line)
		// TODO(setparser) this is a parser/unmarshler, whtever
		keys := strings.Split(line, ",")
		for _, key := range keys {
			parts := strings.Split(key, ":")
			c, parts := parts[len(parts)-1], parts[:len(parts)-1]
			crgba, ok := colornames.Map[c]
			if !ok {
				crgba, err = Hex(c, 3)
				if err != nil {
					log.Println("Could not render color", c)
					log.Println(err)
					continue
				}
			}
			color := &razer.RGBA{crgba}

			switch l := len(parts); l {
			case 1:
				if locs, ok := keyboard[parts[0]]; ok {
					//set key
					for _, loc := range locs {
						s.Rows = append(s.Rows, &razer.Row{
							Num:   loc.Row,
							Start: loc.Col,
							Colors: []razer.Color{
								color,
							},
						})
					}
				}
			case 2:
				c, err := strconv.Atoi(parts[0])
				if err != nil {
					continue
				}
				r, err := strconv.Atoi(parts[1])
				if err != nil {
					continue
				}

				s.Rows = append(s.Rows, &razer.Row{
					Num:   uint8(r),
					Start: uint8(c),
					Colors: []razer.Color{
						color,
					},
				})
			}
		}
		// TODO(readablecode) it's rendering into s and sending s to a pipe, this should be made clear via code and function names
		pipe <- s
	}

	if err := scanner.Err(); err != nil {
		//TODO(dontpanicandreturnerr) return don't panic
		panic(err)
	}
}

func Hex(scol string, contrast uint8) (color.RGBA, error) {
	format := "#%02x%02x%02x"
	if len(scol) == 4 {
		format = "#%1x%1x%1x"
	}

	var r, g, b uint8
	n, err := fmt.Sscanf(scol, format, &r, &g, &b)
	if err != nil {
		return color.RGBA{}, err
	}
	if n != 3 {
		return color.RGBA{}, fmt.Errorf("color: %v is not a hex-color", scol)
	}

	return color.RGBA{r * contrast, g * contrast, b * contrast, 0x00}, nil
}
