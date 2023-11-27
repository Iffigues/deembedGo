package main

import (
	"debug/buildinfo"
	"runtime/debug"
)

type exec interface {
	Open(string)
	Close()
}

func getBuildInfoSetting(d *debug.BuildInfo, a string) string {
	for _, val := range d.Settings {
		if val.Key == a {
			return val.Value
		}
	}
	return ""
}

func getBuildInfo(a string) (*debug.BuildInfo, error) {
	return buildinfo.ReadFile(a)
}
