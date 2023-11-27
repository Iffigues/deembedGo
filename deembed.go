package main

import (
	"deembed/macho"
	"fmt"
	"os"
)

func getPlatform(a string) exec {
	switch a {
	case "windows":
		return nil
	case "darwin":
		return &macho.Macho{}
	case "linux":
		return nil
	default:
		return nil
	}
}

func start(e exec, a string) {
	e.Open(a)
}

func deembed(a string) {
	bi, err := getBuildInfo(a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os := getBuildInfoSetting(bi, "GOOS")
	exe := getPlatform(os)
	start(exe, a)

}
