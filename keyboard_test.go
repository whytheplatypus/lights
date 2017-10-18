package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/whytheplatypus/lights/razer"
	"golang.org/x/image/colornames"
)

var keyboard bool

func TestMain(m *testing.M) {
	flag.BoolVar(&keyboard, "keyboard", false, "enable for interactive keyboard tests")
	flag.Parse()
	os.Exit(m.Run())
}

func TestKeys(t *testing.T) {
	if !keyboard {
		t.Skip("Run with -keyboard for interactive test")
	}
	dev := razer.DeviceList[0]
	evts := make(chan string, 1)
	go func(evts chan string, t *testing.T) {
		X, err := xgbutil.NewConn()
		if err != nil {
			t.Fatal(err)
		}
		keybind.Initialize(X)
		wID := X.RootWin()
		if err := keybind.GrabKeyboard(X, wID); err != nil {
			t.Fatalf("Could not grab keyboard: %s", err)
		}
		//win.Listen(xproto.EventMaskKeyPress, xproto.EventMaskKeyRelease)
		xevent.KeyPressFun(
			func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
				t.Logf("%+v", ev)
				if len(evts) < 1 {
					k := keybind.LookupString(X, ev.State, ev.Detail)
					t.Log("Key pressed:", k)
					evts <- k
				}
			}).Connect(X, wID)
		xevent.Main(X)
	}(evts, t)
	out, in := io.Pipe()

	go render(out)

	defer razer.SetBreathRandom(dev, razer.Conn)
	for _, key := range razer.Keys {
		razer.ClearCustom(dev, razer.Conn, &razer.RGBA{colornames.Black})
		_, err := fmt.Fprintln(in, fmt.Sprintf("%s:%s", key, "blue"))
		if err != nil {
			t.Fatal(err)
		}
		select {
		case pressed := <-evts:
			if pressed != key {
				//odd exception
				if pressed == "L2" && key == "F12" {
					continue
				}
				if pressed == "L1" && key == "F11" {
					continue
				}
				t.Errorf("Wrong key, expected %s got %s", key, pressed)
			}
			if pressed == "Caps_Lock" {
				//eat reset
				<-evts
			}
		case <-time.After(5 * time.Second):
			t.Error("Incomplete test")
			return
		}
	}
}
