package main

import (
	"example.com/m/internal/routers"
)

func main() {
	r := routers.NewRouter()
	err := r.Run(":4200")
	if err != nil {
		return
	}
}
