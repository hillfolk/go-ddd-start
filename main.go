package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hillfolk/go-ddd-start/config"
	"github.com/hillfolk/go-ddd-start/db"
	"github.com/hillfolk/go-ddd-start/server"
)

func main() {
	fmt.Println("start go-ddd-start")
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
