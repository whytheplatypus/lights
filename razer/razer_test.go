package razer_test

import (
	"testing"

	"github.com/whytheplatypus/lights/razer"
)

func TestDevices(t *testing.T) {
	if len(razer.DeviceList) < 1 {
		t.Error("No devices present")
	}
	t.Log(razer.DeviceList)
}

func TestGetDeviceName(t *testing.T) {
	dev := razer.DeviceList[0]
	_, err := razer.GetDeviceName(dev, razer.Conn)
	if err != nil {
		t.Fatal(err)
	}
}
