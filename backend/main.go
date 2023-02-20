package main

import (
	"backend/router"
)

func main() {
	r := router.Router()
	err := r.Run()
	if err != nil {
		return
	}
}
