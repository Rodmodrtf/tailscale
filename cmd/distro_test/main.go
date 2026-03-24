package main

import (
	"fmt"
	"os"
)

func have(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func haveDir(file string) bool {
	fi, err := os.Stat(file)
	return err == nil && fi.IsDir()
}

func main() {
	fmt.Printf("haveDir(/jffs): %v\n", haveDir("/jffs"))
	fmt.Printf("have(/jffs/scripts/start_tailscale.sh): %v\n", have("/jffs/scripts/start_tailscale.sh"))
	fmt.Printf("have(/etc/openwrt_version): %v\n", have("/etc/openwrt_version"))
	fmt.Printf("have(/etc/alpine-release): %v\n", have("/etc/alpine-release"))
}
