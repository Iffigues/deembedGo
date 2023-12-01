package utils

import (
	"debug/buildinfo"
	"debug/dwarf"
	"errors"
	"fmt"
	"runtime/debug"
)

type Exec interface {
	Open() error
	Close()
	BuildInfo() (*debug.BuildInfo, error)
	Dwarf() (*dwarf.Data, error)
}

func GetBuildInfoSetting(d *debug.BuildInfo, a string) (string, error) {
	for _, val := range d.Settings {
		if val.Key == a {
			return val.Value, nil
		}
	}
	return "", errors.New("don't exists")
}

func GetBuildInfo(a string) (*debug.BuildInfo, error) {
	return buildinfo.ReadFile(a)
}

func BuildInfo(a string) {
	aa, e := buildinfo.ReadFile(a)

	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(aa)
}
