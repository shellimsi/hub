package main

import (
	"shellimsi/hub/server"
)

func main() {

	err := make(chan error, 1)
	go func() {
		err <- server.Serve(&server.Server{})
	}()

	go func() {
		for {
			// Todo: better error handing
			handleErr := <-err
			switch handleErr {
			default:
				panic(handleErr)

			}
		}
	}()
	select {}
}
