package lights_test

import (
	"flag"
	"os"
	"testing"
	"time"

	"golang.org/x/image/colornames"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/whytheplatypus/lights"
)

var keyboard bool

func TestMain(m *testing.M) {
	flag.BoolVar(&keyboard, "keyboard", false, "enable for interactive keyboard tests")
	flag.Parse()
	os.Exit(m.Run())
}

func TestDevices(t *testing.T) {
	if len(lights.DeviceList) < 1 {
		t.Error("No devices present")
	}
	t.Log(lights.DeviceList)
}

func TestKeys(t *testing.T) {
	if !keyboard {
		t.Skip("Run with -i for interactive test")
	}
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
	dev := lights.DeviceList[0]
	defer lights.SetBreathRandom(dev, lights.Conn)
	s, err := lights.GetDeviceName(dev, lights.Conn)
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range lights.Keys {
		locs := lights.Keyboards[s][key]
		lights.ClearCustom(dev, lights.Conn, &lights.RGBA{colornames.Black})
		s := &lights.Set{
			Rows: []*lights.Row{},
		}
		for _, loc := range locs {
			s.Rows = append(s.Rows, &lights.Row{
				Num:   loc.Row,
				Start: loc.Col,
				Colors: []lights.Color{
					&lights.RGBA{colornames.Purple},
				},
			})
		}
		lights.Apply(dev, lights.Conn, s)
		lights.SetCustom(dev, lights.Conn)
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

func TestGetDeviceName(t *testing.T) {
	dev := lights.DeviceList[0]
	_, err := lights.GetDeviceName(dev, lights.Conn)
	if err != nil {
		t.Fatal(err)
	}
}
