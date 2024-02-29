package main

import "github.com/josenymad/hexagonal-go-api/internal/adapters/primary/http"

func main() {
	http := http.NewHttpServer("localhost", "8000")
	http.Run()
}
