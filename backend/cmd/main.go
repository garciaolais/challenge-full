package main

import (
	"backend/app/http"
	"backend/version"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	versionPtr := flag.Bool("version", false, "get the version number")
	portPtr := flag.Int("port", 1986, "first number in the list")
	flag.Parse()

	if *versionPtr {
		fmt.Println(version.Version)
		return 0
	}

	port := strconv.Itoa(*portPtr)
	http.Serve(port)
	return 0
}
