package main

import (
	"fmt"
	"os"
	"runtime"
)

func getPlatform() string {
	switch runtime.GOOS {
	case "windows":
		return "Windows"
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	default:
		return "Other"
	}
}

func deembed(a string) {
	bi, err := getBuildInfo(a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os := getBuildInfoSetting(bi, "GOOS")
	arch := getBuildInfoSetting(bi, "GOARCH")
	fmt.Println(os, arch)
}
