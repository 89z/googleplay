package googleplay

import (
	"os"
	"testing"
	"time"
)

func checkin(id int64) error {
	platform := Platforms[id]
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	device, err := Phone.Checkin(platform)
	if err != nil {
		return err
	}
	platform += ".bin"
	if err := device.Create(home + "/googleplay/" + platform); err != nil {
		return err
	}
	time.Sleep(Sleep)
	return nil
}

func Test_Checkin_ARMEABI(t *testing.T) {
	err := checkin(1)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Checkin_ARM64(t *testing.T) {
	err := checkin(2)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Checkin_X86(t *testing.T) {
	err := checkin(0)
	if err != nil {
		t.Fatal(err)
	}
}
