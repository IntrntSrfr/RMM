package main

import "github.com/intrntsrfr/rmm-api/handler"

func main() {
	r := handler.NewHandler()
	r.Run(":4444")
}
