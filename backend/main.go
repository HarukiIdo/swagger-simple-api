package main

import (
	"flag"

	"github.com/HarukiIdo/swagger-simple-api/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 1999, "Web server port")
	s := app.NewServer()
	s.Run(port)
}
