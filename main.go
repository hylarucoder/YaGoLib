package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli"
)

var banner = `
  ___    ___ ________  ________  ________  ___       ___  ________
 |\  \  /  /|\   __  \|\   ____\|\   __  \|\  \     |\  \|\   __  \
 \ \  \/  / | \  \|\  \ \  \___|\ \  \|\  \ \  \    \ \  \ \  \|\ /_
  \ \    / / \ \   __  \ \  \  __\ \  \\\  \ \  \    \ \  \ \   __  \
   \/  /  /   \ \  \ \  \ \  \|\  \ \  \\\  \ \  \____\ \  \ \  \|\  \
 __/  / /      \ \__\ \__\ \_______\ \_______\ \_______\ \__\ \_______\
|\___/ /        \|__|\|__|\|_______|\|_______|\|_______|\|__|\|_______|
\|___|/
`

func main() {
	fmt.Println(banner)
	app := cli.NewApp()
	app.Name = "yagolib"
	app.Version = "0.0.1"
	app.Usage = "yet another go library for lazy guys"
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	resp := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("hello world")
	fmt.Println(resp)

}
