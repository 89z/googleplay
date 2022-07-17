package googleplay

import (
	"fmt"
	"os"
	"testing"
)

func Test_Delivery(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	var head Header
	head.Open_Auth(home + "/googleplay/auth.txt")
	head.Open_Device(home + "/googleplay/x86.bin")
	del, err := head.Delivery("com.google.android.youtube", 1524221376)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", del)
}
