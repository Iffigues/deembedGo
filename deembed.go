package main

import (
	"deembed/macho"
	"errors"
	"fmt"
	"os"
)

func onError(err error, e int, a ...string) {
	if err == nil {
		return
	}
	for _, val := range a {
		fmt.Print(val + "\n")
	}
	fmt.Println(err)
	os.Exit(e)
}

func getPlatform(a string) (exec, error) {
	switch a {
	case "windows":
		return nil, nil
	case "darwin":
		return &macho.Macho{}, nil
	case "linux":
		return nil, nil
	default:
		return nil, errors.New("operating system " + a + " is not gerer")
	}
}

func start(e exec, a string) {
	e.Open(a)
}

func deembed(a string) {
	bi, err := getBuildInfo(a)
	onError(err, 1)
	goos := getBuildInfoSetting(bi, "GOOS")
	exe, err := getPlatform(goos)
	onError(err, 1)
	start(exe, a)
}
