package main

import (
	"os"
	"pikpo_exam/app/server"
)

func main() {
	if err := server.Run(); err != nil {
		os.Exit(1)
	}
}
