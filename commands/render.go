package commands

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/whytheplatypus/lights/razer"
)

type Renderer struct{}

func defaultKeyboard() string {
	kbs := razer.GetKeyboards(razer.Conn)
	if len(kbs) > 0 {
		return kbs[0]
	}
	return ""
}

func (r *Renderer) Run(args []string) error {
	rFlags := flag.NewFlagSet("renderer", flag.ContinueOnError)

	var path string
	rFlags.StringVar(
		&path,
		"fifo",
		os.Getenv("LIGHTS_FIFO"),
		"specify path for a fifo, one will be created if it doesn't already exist at the path")

	var dev string
	rFlags.StringVar(
		&dev,
		"keyboard",
		defaultKeyboard(),
		fmt.Sprintf("specify the device id to render out to : \n %s", razer.DeviceList))

	var preset string
	rFlags.StringVar(
		&preset,
		"preset",
		"reactive",
		//TODO(preset-options) include options
		"specify the preset effect to return to when render exits")

	if err := rFlags.Parse(args); err != nil {
		return err
	}

	in, err := Input(path)
	if err != nil {
		return err
	}

	defer presets[preset](dev, razer.Conn)

	pipe := make(chan *razer.Set, 100)
	go renderLoop(dev, pipe)
	go render(dev, pipe, in)

	sig := Signal(syscall.SIGINT, syscall.SIGTERM)
	<-sig
	return nil
}

func Input(path string) (io.Reader, error) {
	if path != "" {
		os.Remove(path)
		if err := syscall.Mkfifo(path, 0600); err != nil {
			log.Println(err)
			return nil, err
		}
		var err error
		in, err := os.OpenFile(path, os.O_RDWR, 0600)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		// TODO(removeremove) can we not use these?
		defer os.Remove(path)
		return in, nil
	}
	return os.Stdin, nil
}

func renderLoop(dev string, pipe <-chan *razer.Set) {
	for s := range pipe {
		razer.Apply(dev, razer.Conn, s)
		razer.SetCustom(dev, razer.Conn)
	}
}

func render(dev string, pipe chan *razer.Set, in io.Reader) {

	dn, err := razer.GetDeviceName(dev, razer.Conn)
	if err != nil {
		log.Fatal(err)
	}

	keyboard, ok := razer.Keyboards[dn]
	if !ok {
		panic(fmt.Errorf("no keyboard found %s", dn))
	}

	// can configre to come from a fifo
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		// TODO(untangle) these 3 or 4 statements feel jumbled
		line := scanner.Text()
		if len(pipe) > 0 {
			log.Println("[DEBUG] skipping", line)
			continue
		}
		log.Println("[DEBUG]", line)
		s := &razer.Set{}
		if err := razer.UnmarshalString(line, keyboard, s); err != nil {
			log.Println("[ERROR] ", err)
			continue
		}
		pipe <- s
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func Signal(sigs ...os.Signal) chan os.Signal {
	// TODO(sigsetup) better for this to be a named function so it's clear what's happening
	sig := make(chan os.Signal)

	//Setup signal handling for clean shutdown
	signal.Notify(sig, sigs...)
	return sig
}
