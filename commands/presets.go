package commands

import (
	"github.com/godbus/dbus"
	"github.com/whytheplatypus/lights/razer"
	"golang.org/x/image/colornames"
)

var presets = map[string]func(name string, conn *dbus.Conn){
	"breath": razer.SetBreathRandom,
	"reactive": func(name string, conn *dbus.Conn) {
		razer.SetReactive(name, conn, &razer.RGBA{colornames.Purple})
	},
}
