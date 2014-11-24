package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/jessevdk/go-flags"
)

const apiServer = "https://gd2.line.naver.jp:443"

var parser = flags.NewParser(&common, flags.Default)

type Common struct {
	User            string `short:"u" long:"user" description:"User name" required:"yes"`
	Device          string `short:"d" long:"device" description:"Device name" required:"yes"`
	ReportingServer string `short:"r" long:"reportingServer" description:"URL of the reporting server" default:"http://itwill.be/tsum"`
}

var common Common

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
	fmt.Printf("\nCompleted.\n")
}

type prog struct {
	str  string
	args []interface{}
}

func printProgress(p prog) {
	fmt.Print(strings.Repeat(" ", 79), "\r")
	fmt.Printf(p.str+"\r", p.args...)
}
