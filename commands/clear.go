package commands

import (
	"github.com/whytheplatypus/lights/razer"
	"golang.org/x/image/colornames"
)

type Clear struct{}

func (r *Clear) Run(args []string) error {
	// TODO(getkeyboard) same as before why 0?
	// should be something like GetKeyboard
	dev := razer.DeviceList[0]
	razer.ClearCustom(dev, razer.Conn, &razer.RGBA{colornames.Black})
	return nil
}
