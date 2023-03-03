package main

import (
	"flag"
)

var (
	dirFlag  string
	fromFlag string
	toFlag   string
)

func init() {
	flag.StringVar(&dirFlag, "dir_name", "", "absolute directory name where you have pictures that you want to convert format")
	flag.StringVar(&fromFlag, "from", "jpeg", "picture format that you want to convert from")
	flag.StringVar(&toFlag, "to", "png", "picture format that you want to convert to")
}

func main() {
	flag.Parse()
}
