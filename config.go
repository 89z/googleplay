package googleplay

import (
	"bytes"
	"github.com/89z/rosso/os"
	"github.com/89z/rosso/protobuf"
	"github.com/89z/rosso/strconv"
	"io"
	"net/http"
)

func (n Native_Platform) String() string {
	var buf strconv.Buffer
	buf.WriteString("nativePlatform")
	for key, val := range n {
		buf.WriteByte('\n')
		buf.AppendInt(key)
		buf.WriteString(": ")
		buf.WriteString(val)
	}
	return string(buf)
}

func (h *Header) Open_Device(name string) error {
	buf, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	h.Device.Message, err = protobuf.Unmarshal(buf)
	if err != nil {
		return err
	}
	return nil
}

type Native_Platform map[int64]string

var Platforms = Native_Platform{
	// com.google.android.youtube
	0: "x86",
	// com.miui.weather2
	1: "armeabi-v7a",
	// com.kakaogames.twodin
	2: "arm64-v8a",
}

func (d Device) Create(name string) error {
	data := d.Marshal()
	return os.WriteFile(name, data)
}

// These can use default values, but they must all be included
type Config struct {
	Device_Feature      []string
	Five_Way_Navigation uint64
	GL_ES_Version       uint64
	GL_Extension        []string
	Hard_Keyboard       uint64
	Keyboard            uint64
	Navigation          uint64
	Screen_Density      uint64
	Screen_Layout       uint64
	Shared_Library      []string
	Touch_Screen        uint64
}

var Phone = Config{
	Device_Feature: []string{
		// app.source.getcontact
		"android.hardware.location.gps",
		// br.com.rodrigokolb.realdrum
		"android.software.midi",
		// com.app.xt
		"android.hardware.camera.front",
		// com.clearchannel.iheartradio.controller
		"android.hardware.microphone",
		// com.google.android.apps.walletnfcrel
		"android.software.device_admin",
		// com.google.android.youtube
		"android.hardware.touchscreen",
		"android.hardware.wifi",
		// com.illumix.fnafar
		"android.hardware.sensor.gyroscope",
		// com.madhead.tos.zh
		"android.hardware.sensor.accelerometer",
		// com.miHoYo.GenshinImpact
		"android.hardware.opengles.aep",
		// com.pinterest
		"android.hardware.camera",
		"android.hardware.location",
		"android.hardware.screen.portrait",
		// com.supercell.brawlstars
		"android.hardware.touchscreen.multitouch",
		// com.sygic.aura
		"android.hardware.location.network",
		// com.xiaomi.smarthome
		"android.hardware.bluetooth",
		"android.hardware.bluetooth_le",
		"android.hardware.camera.autofocus",
		"android.hardware.usb.host",
		// kr.sira.metal
		"android.hardware.sensor.compass",
		// org.thoughtcrime.securesms
		"android.hardware.telephony",
		// org.videolan.vlc
		"android.hardware.screen.landscape",
	},
	Shared_Library: []string{
		// com.amctve.amcfullepisodes
		"org.apache.http.legacy",
		// com.binance.dev
		"android.test.runner",
		// com.miui.weather2
		"global-miui11-empty.jar",
	},
	// com.axis.drawingdesk.v3
	GL_ES_Version: 0x9_9999,
	GL_Extension: []string{
		// com.instagram.android
		"GL_OES_compressed_ETC1_RGB8_texture",
		// com.kakaogames.twodin
		"GL_KHR_texture_compression_astc_ldr",
	},
}

func (d Device) ID() (uint64, error) {
	return d.Get_Fixed64(7)
}

// A Sleep is needed after this.
func (c Config) Checkin(platform string) (*Device, error) {
	checkin := protobuf.Message{
		4: protobuf.Message{ // checkin
			1: protobuf.Message{ // build
				10: protobuf.Varint(29), // sdkVersion
			},
		},
		14: protobuf.Varint(3), // version
		18: protobuf.Message{ // deviceConfiguration
			1:  protobuf.Varint(c.Touch_Screen), // touchScreen
			2:  protobuf.Varint(c.Keyboard),
			3:  protobuf.Varint(c.Navigation),
			4:  protobuf.Varint(c.Screen_Layout),
			5:  protobuf.Varint(c.Hard_Keyboard),
			6:  protobuf.Varint(c.Five_Way_Navigation),
			7:  protobuf.Varint(c.Screen_Density),
			8:  protobuf.Varint(c.GL_ES_Version),
			11: protobuf.String(platform), // nativePlatform
		},
	}
	for _, library := range c.Shared_Library {
		// .deviceConfiguration.systemSharedLibrary
		checkin.Get(18).Add_String(9, library)
	}
	for _, extension := range c.GL_Extension {
		// .deviceConfiguration.glExtension
		checkin.Get(18).Add_String(15, extension)
	}
	for _, name := range c.Device_Feature {
		// .deviceConfiguration.deviceFeature
		checkin.Get(18).Add(26, protobuf.Message{
			1: protobuf.String(name),
		})
	}
	req, err := http.NewRequest(
		"POST", "https://android.googleapis.com/checkin",
		bytes.NewReader(checkin.Marshal()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-protobuffer")
	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var dev Device
	dev.Message, err = protobuf.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return &dev, nil
}

type Device struct {
	protobuf.Message
}
