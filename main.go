package main

import (
	"flag"
	"os"
)

func begin(a map[string]string) {
	e := getPtf()
	for key, val := range a {
		if val != "" {
			if ptf, ok := e[key]; ok {
				ptf(val)
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	var deembed string
	flag.StringVar(&deembed, "deembed", "", "Specify deembed")

	flag.Parse()
	begin(map[string]string{"deembed": deembed})
}
