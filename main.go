package main

import "github.com/mkdirjava/dummy_domain/stringer"

func main() {
	service := stringer.NewGetStringerService()
	println(service.GetString())
}
