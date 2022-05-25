package main

import (
   "fmt"
   "github.com/89z/format"
   "io"
   "net/http"
   "os"
   "time"
   gp "github.com/89z/googleplay"
)

func doDetails(head *gp.Header, app string, parse bool) error {
   detail, err := head.Details(app)
   if err != nil {
      return err
   }
   if parse {
      date, err := detail.Time()
      if err != nil {
         return err
      }
      detail.UploadDate = date.String()
   }
   fmt.Println(detail)
   return nil
}

func doHeader(platform string, single bool) (*gp.Header, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return nil, err
   }
   token, err := gp.OpenToken(cache, "googleplay/token.json")
   if err != nil {
      return nil, err
   }
   device, err := gp.OpenDevice(cache, "googleplay", platform + ".json")
   if err != nil {
      return nil, err
   }
   return token.Header(device.AndroidID, single)
}

func doDelivery(head *gp.Header, app string, ver uint64) error {
   download := func(addr, name string) error {
      fmt.Println("GET", addr)
      res, err := http.Get(addr)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      file, err := os.Create(name)
      if err != nil {
         return err
      }
      defer file.Close()
      pro := format.ProgressBytes(file, res.ContentLength)
      if _, err := io.Copy(pro, res.Body); err != nil {
         return err
      }
      return nil
   }
   del, err := head.Delivery(app, ver)
   if err != nil {
      return err
   }
   for _, split := range del.SplitDeliveryData {
      err := download(split.DownloadURL, del.Split(split.ID))
      if err != nil {
         return err
      }
   }
   for _, file := range del.AdditionalFile {
      err := download(file.DownloadURL, del.Additional(file.FileType))
      if err != nil {
         return err
      }
   }
   return download(del.DownloadURL, del.Download())
}

func doToken(email, password string) error {
   tok, err := gp.NewToken(email, password)
   if err != nil {
      return err
   }
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   return tok.Create(cache, "googleplay/token.json")
}

func doDevice(platform string) error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   device, err := gp.Phone.Checkin(platform)
   if err != nil {
      return err
   }
   fmt.Printf("Sleeping %v for server to process\n", gp.Sleep)
   time.Sleep(gp.Sleep)
   return device.Create(cache, "googleplay", platform + ".json")
}
