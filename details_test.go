package googleplay

import (
   "os"
   "testing"
   "time"
)

var apps = []appType{
   {"2021-06-01", 0, "com.valvesoftware.android.steam.community"},
   {"2021-12-08", 0, "com.amctve.amcfullepisodes"},
   {"2022-02-14", 0, "org.videolan.vlc"},
   {"2022-03-01", 0, "kr.sira.metal"},
   {"2022-03-17", 0, "com.google.android.apps.walletnfcrel"},
   {"2022-03-24", 1, "com.miui.weather2"},
   {"2022-04-08", 1, "com.axis.drawingdesk.v3"},
   {"2022-05-05", 0, "com.clearchannel.iheartradio.controller"},
   {"2022-05-13", 0, "com.app.xt"},
   {"2022-05-16", 0, "com.instagram.android"},
   {"2022-05-17", 0, "com.pinterest"},
   {"2022-02-02", 2, "com.illumix.fnafar"},
   {"2022-03-16", 2, "com.miHoYo.GenshinImpact"},
   {"2022-03-24", 0, "app.source.getcontact"},
   {"2022-04-27", 1, "com.sygic.aura"},
   {"2022-04-29", 1, "com.xiaomi.smarthome"},
   {"2022-05-12", 0, "com.google.android.youtube"},
   {"2022-05-15", 2, "com.kakaogames.twodin"},
   {"2022-05-16", 0, "com.binance.dev"},
   {"2022-05-17", 0, "br.com.rodrigokolb.realdrum"},
   {"2022-05-17", 0, "org.thoughtcrime.securesms"},
   {"2022-05-18", 1, "com.madhead.tos.zh"},
}

type appType struct {
   date string
   platform int64 // X-DFE-Device-ID
   id string
}

func TestDetails(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   token, err := OpenToken(cache, "googleplay/token.json")
   if err != nil {
      t.Fatal(err)
   }
   head, err := token.Header(0, false)
   if err != nil {
      t.Fatal(err)
   }
   for _, app := range apps {
      platform := Platforms[app.platform]
      device, err := OpenDevice(cache, "googleplay", platform + ".json")
      if err != nil {
         t.Fatal(err)
      }
      head.AndroidID = uint64(device.AndroidID)
      det, err := head.Details(app.id)
      if err != nil {
         t.Fatal(err)
      }
      if det.CurrencyCode == "" {
         t.Fatal(det)
      }
      if det.NumDownloads == 0 {
         t.Fatal(det)
      }
      if det.Size == 0 {
         t.Fatal(det)
      }
      if det.Title == "" {
         t.Fatal(det)
      }
      if det.UploadDate == "" {
         t.Fatal(det)
      }
      if det.VersionCode == 0 {
         t.Fatal(det)
      }
      if det.VersionString == "" {
         t.Fatal(det)
      }
      time.Sleep(99 * time.Millisecond)
   }
}

