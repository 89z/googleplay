package main

import (
	"flag"
	gp "github.com/89z/googleplay"
	"os"
)

type flags struct {
	app      string
	device   bool
	email    string
	password string
	platform int64
	purchase bool
	single   bool
	version  uint64
}

func main() {
	var f flags
	// a
	flag.StringVar(&f.app, "a", "", "app")
	// device
	flag.BoolVar(&f.device, "device", false, "create device")
	// email
	flag.StringVar(&f.email, "email", "", "your email")
	// p
	flag.Int64Var(&f.platform, "p", 0, gp.Platforms.String())
	// password
	flag.StringVar(&f.password, "password", "", "your password")
	// purchase
	flag.BoolVar(&f.purchase, "purchase", false, "purchase app")
	// s
	flag.BoolVar(&f.single, "s", false, "single APK")
	// v
	flag.Uint64Var(&f.version, "v", 0, "app version")
	flag.Parse()
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dir += "/googleplay"
	if f.email != "" {
		err := f.do_auth(dir)
		if err != nil {
			panic(err)
		}
	} else {
		platform := gp.Platforms[f.platform]
		if f.device {
			err := do_device(dir, platform)
			if err != nil {
				panic(err)
			}
		} else if f.app != "" {
			head, err := f.do_header(dir, platform)
			if err != nil {
				panic(err)
			}
			if f.purchase {
				err := head.Purchase(f.app)
				if err != nil {
					panic(err)
				}
			} else if f.version >= 1 {
				err := f.do_delivery(head)
				if err != nil {
					panic(err)
				}
			} else {
				detail, err := f.do_details(head)
				if err != nil {
					panic(err)
				}
				os.Stdout.Write(detail)
			}
		} else {
			flag.Usage()
		}
	}
}
