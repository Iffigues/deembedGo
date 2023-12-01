package main

import (
	"deembed/macho"
	"deembed/utils"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func buildinfo(a string) {
	bi, err := utils.GetBuildInfo(a)
	if err != nil {
		log.Fatal(a, err)
	}
	fmt.Println(bi)
}

func buildinfosetting(a, b string) {
	bi, err := utils.GetBuildInfo(a)
	if err != nil {
		log.Fatal(a, err)
	}
	s, err := utils.GetBuildInfoSetting(bi, b)
	if err != nil {
		log.Fatal(a, err)
	}
	fmt.Println(b, "=", s)
}

func switchy(a, b string, c map[string]string) {
	if b == "" {
		return
	}
	switch a {
	case "deembed":
		if b == "" {
			return
		}
	case "buildinfo":
		if b == "" {
			return
		}
		if set, ok := c["settings"]; ok && set != "" {
			buildinfosetting(b, set)
		} else {
			buildinfo(b)
		}
	case "settings":
		return
	default:
		fmt.Println("not a good vibe")
	}
}

func begin(a map[string]string) {
	for key, val := range a {
		switchy(key, val, a)
	}
}

func getPlatform(a, p string) (utils.Exec, error) {
	switch a {
	case "windows":
		return nil, nil
	case "darwin":
		return macho.NewMacho(p), nil
	case "linux":
		return nil, nil
	default:
		return nil, errors.New("operating system " + a + " is not gerer")
	}
}

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

func Deembed(a string) (utils.Exec, error) {
	bi, err := utils.GetBuildInfo(a)
	onError(err, 1)
	goos, err := utils.GetBuildInfoSetting(bi, "GOOS")
	if err != nil {
		log.Fatal(err)
	}
	return getPlatform(goos, a)
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	var deembed string
	flag.StringVar(&deembed, "deembed", "", "Specify deembed")

	var buildinfo string
	flag.StringVar(&buildinfo, "buildinfo", "", "Specify buildinfo")

	var Settings string
	flag.StringVar(&Settings, "settings", "", "Specify buildinfo")

	flag.Parse()
	begin(map[string]string{"deembed": deembed, "buildinfo": buildinfo, "settings": Settings})
}
