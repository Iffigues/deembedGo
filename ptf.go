package main

func getPtf() map[string]func(string) {
	return map[string]func(string){
		"deembed": deembed,
	}
}
